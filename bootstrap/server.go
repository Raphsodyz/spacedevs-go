package bootstrap

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	controllers "github.com/Raphsodyz/spacedevs-go/cmd/api/controllers"
	"github.com/Raphsodyz/spacedevs-go/config"
	rediscache "github.com/Raphsodyz/spacedevs-go/pkg/cache"
	postgres "github.com/Raphsodyz/spacedevs-go/pkg/db/postgres"
	"github.com/Raphsodyz/spacedevs-go/repository"
	"github.com/Raphsodyz/spacedevs-go/usecase"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type Server struct {
	Engine *gin.Engine
	Config *config.Config
	Db     *pgxpool.Pool
	Redis  *redis.Client
}

func NewServer(configName string) (*Server, error) {
	v, err := config.LoadConfig(configName)
	if err != nil {
		return nil, fmt.Errorf("bootstrap.NewServer: failed to load config: %w", err)
	}

	cfg, err := config.ParseConfig(v)
	if err != nil {
		return nil, fmt.Errorf("bootstrap.NewServer: failed to parse config: %w", err)
	}

	db, err := postgres.NewPostgresPool(configName)
	if err != nil {
		return nil, fmt.Errorf("bootstrap.NewServer: failed to init postgres: %w", err)
	}

	redisClient, err := rediscache.NewRedisClient(cfg)
	if err != nil {
		return nil, fmt.Errorf("bootstrap.NewServer: failed to init redis: %w", err)
	}

	if cfg.Server.Mode == "Development" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	engine := gin.New()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())

	return &Server{
		Engine: engine,
		Config: cfg,
		Db:     db,
		Redis:  redisClient,
	}, nil
}

func (s *Server) Run() error {
	s.registerRoutes()

	srv := &http.Server{
		Addr:         s.Config.Server.Port,
		Handler:      s.Engine,
		ReadTimeout:  s.Config.Server.ReadTimeout * time.Second,
		WriteTimeout: s.Config.Server.WriteTimeout * time.Second,
	}

	errCh := make(chan error, 1)
	go func() {
		log.Printf("Server listening on %s", s.Config.Server.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errCh <- fmt.Errorf("bootstrap.Server.Run: %w", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-errCh:
		return err
	case <-quit:
		log.Println("Shutting down server...")
	}

	ctx, cancel := context.WithTimeout(
		context.Background(),
		s.Config.Server.CtxDefaultTimeout*time.Second,
	)

	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		return fmt.Errorf("bootstrap.Server.Run: forced shutdown: %w", err)
	}

	s.Db.Close()

	if err := s.Redis.Close(); err != nil {
		log.Printf("bootstrap.Server.Run: redis close warning: %v", err)
	}

	log.Println("Server exited cleanly.")
	return nil
}

func (s *Server) registerRoutes() {
	launchRepo := repository.NewLaunchRepository(s.Db)
	redisRepo := repository.NewRedisRepository(s.Redis)
	missionRepo := repository.NewMissionRepository(s.Db)
	configRepo := repository.NewConfigurationRepository(s.Db)
	locationRepo := repository.NewLocationRepository(s.Db)
	padRepo := repository.NewPadRepository(s.Db)

	searchUseCase := usecase.NewSearchLaunchUseCase(missionRepo, configRepo, locationRepo, padRepo, launchRepo, redisRepo)

	launchController := controllers.NewLaunchController(*searchUseCase)
	controllers.RegisterRoutes(s.Engine, launchController)
}
