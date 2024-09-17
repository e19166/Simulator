package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/e19166/Simulator/db"
	"github.com/e19166/Simulator/models"
)

// CreateProject handles the creation of a new project.
// Expects a JSON payload in the request body and inserts the project data into the database.
func CreateProject(w http.ResponseWriter, r *http.Request) {
	var project models.Project

	// Decode the JSON request body into the Project struct
	err := json.NewDecoder(r.Body).Decode(&project)
	if err != nil {
		// If decoding fails, respond with a bad request status and error message
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Log the received project data for debugging purposes
	fmt.Printf("Received Project Data: %+v\n", project)

	// Execute an SQL command to insert the project data into the database
	_, err = db.DB.Exec(`
        INSERT INTO projects 
        (name, environment, type, monitoring_type, deletion_timestamp, reason, message, api_version, kind, involved_object_name, involved_object_uuid, involved_object_version, action, event_time, component, host, count, outcome, current_status, correlation_id, user_idp_id, org_uuid, first_timestamp, last_timestamp, state) 
        VALUES 
        ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25)`,
		project.Metadata.Name,
		project.Metadata.Labels.Environment,
		project.Metadata.Labels.Type,
		project.Metadata.Annotations.MonitoringType,
		project.Metadata.DeletionTimestamp,
		project.Metadata.Reason,
		project.Metadata.Message,
		project.ApiVersion,
		project.Kind,
		project.InvolvedObject.Name,
		project.InvolvedObject.Uuid,
		project.InvolvedObject.Version,
		project.Action,
		project.EventTime,
		project.Source.Component,
		project.Source.Host,
		project.Count,
		project.Outcome,
		project.CurrentStatus,
		project.CorrelationID,
		project.UserIdpId,
		project.OrgUuid,
		project.Series.FirstTimestamp,
		project.Series.LastTimestamp,
		project.Series.State,
	)

	if err != nil {
		// If SQL execution fails, log the error and respond with an internal server error status
		fmt.Printf("SQL Execution Error: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with a successful creation status and the inserted project data
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(project)
}
