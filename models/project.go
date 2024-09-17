package models

import (
	"github.com/google/uuid"
	"time"
)

type Project struct {
	Metadata struct {
		Name string

		Labels struct {
			Environment string

			Type string
		}

		Annotations struct {
			MonitoringType string
		}

		DeletionTimestamp time.Time

		Reason string

		Message string
	}

	ApiVersion string

	Kind string

	InvolvedObject struct {
		Name string

		Uuid uuid.UUID

		Version string
	}

	Action string

	EventTime time.Time

	Source struct {
		Component string

		Host string
	}

	Count int

	Outcome string

	CurrentStatus string

	CorrelationID string

	UserIdpId string

	OrgUuid string

	Series struct {
		FirstTimestamp time.Time

		LastTimestamp time.Time

		State string
	}
}
