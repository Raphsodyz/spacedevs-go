package audit

import (
	"time"
)

type Audit struct {
	UserInclusion string    `db:"user_inclusion" validate:"required,lte=20"`
	DateInclusion time.Time `db:"date_inclusion" validate:"required"`
	UserChange    *string   `db:"user_change" validate:"lte=20"`
	DateChange    time.Time `db:"date_change"`
	EffectiveDate time.Time `db:"date_change" validate:"required"`
}
