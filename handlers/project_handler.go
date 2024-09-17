package handlers

import (
	"encoding/json"
	"net/http"
	"fmt"

	"github.com/e19166/Simulator/db"
	"github.com/e19166/Simulator/models"
)

func CreateProject(w http.ResponseWriter, r *http.Request) {
    var project models.Project
    err := json.NewDecoder(r.Body).Decode(&project)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Log project data to verify it is being parsed correctly
    fmt.Printf("Received Project Data: %+v\n", project)

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
        fmt.Printf("SQL Execution Error: %v\n", err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(project)
}
