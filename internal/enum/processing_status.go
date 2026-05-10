package enum

import "fmt"

type ProcessingStatus string

const (
	StatusDraft     ProcessingStatus = "DRAFT"
	StatusTrash     ProcessingStatus = "TRASH"
	StatusPublished ProcessingStatus = "PUBLISHED"
)

func (s ProcessingStatus) String() string {
	return string(s)
}

func (s ProcessingStatus) IsValid() bool {
	switch s {
	case StatusDraft, StatusTrash, StatusPublished:
		return true
	default:
		return false
	}
}

func ParseProcessingStatus(value string) (ProcessingStatus, error) {
	status := ProcessingStatus(value)

	if !status.IsValid() {
		return "", fmt.Errorf("invalid processing status: %s", value)
	}

	return status, nil
}
