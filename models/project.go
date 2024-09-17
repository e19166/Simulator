package models

import (
	"github.com/google/uuid"
	"time"
)

// Project represents a project entity with various metadata and attributes.
type Project struct {
	// Metadata contains information about the project metadata
	Metadata struct {
		// Name is the name of the project
		Name string

		// Labels holds key-value pairs for additional categorization of the project
		Labels struct {
			// Environment indicates the environment in which the project is running (e.g., production, staging)
			Environment string

			// Type represents the type of the project
			Type string
		}

		// Annotations holds additional metadata about the project
		Annotations struct {
			// MonitoringType describes the type of monitoring applied to the project
			MonitoringType string
		}

		// DeletionTimestamp is the time when the project was marked for deletion
		DeletionTimestamp time.Time

		// Reason explains why the project was created or modified
		Reason string

		// Message contains a message related to the project
		Message string
	}

	// ApiVersion indicates the version of the API used
	ApiVersion string

	// Kind specifies the kind of the project or object
	Kind string

	// InvolvedObject represents an object related to the project
	InvolvedObject struct {
		// Name is the name of the involved object
		Name string

		// Uuid is the unique identifier of the involved object
		Uuid uuid.UUID

		// Version represents the version of the involved object
		Version string
	}

	// Action describes the action performed related to the project
	Action string

	// EventTime is the time when the event occurred
	EventTime time.Time

	// Source represents the source of the project event
	Source struct {
		// Component is the component from which the event originated
		Component string

		// Host is the host where the event occurred
		Host string
	}

	// Count represents the count of occurrences related to the project
	Count int

	// Outcome describes the result or outcome of the project event
	Outcome string

	// CurrentStatus indicates the current status of the project
	CurrentStatus string

	// CorrelationID is a unique identifier used to correlate related events
	CorrelationID string

	// UserIdpId is the identifier for the user in the identity provider system
	UserIdpId string

	// OrgUuid is the unique identifier for the organization
	OrgUuid string

	// Series holds information about the time series related to the project
	Series struct {
		// FirstTimestamp is the first timestamp of the series
		FirstTimestamp time.Time

		// LastTimestamp is the last timestamp of the series
		LastTimestamp time.Time

		// State represents the state of the series
		State string
	}
}
