package handlers

import (
	"fmt"
	"net/http"

	"github.com/e19166/Simulator/db"
	"github.com/e19166/Simulator/models"
	"github.com/gin-gonic/gin"
)

// CreateProject handles the creation of a new project.
func CreateProject(c *gin.Context) {
	var project models.Project

	// Bind the JSON request body to the Project struct
	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Log the received project data for debugging purposes
	fmt.Printf("Received Project Data: %+v\n", project)

	// Execute an SQL command to insert the project data into the database
	_, err := db.DB.Exec(`
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create project"})
		return
	}

	// Respond with a successful creation status and the inserted project data
	c.JSON(http.StatusCreated, project)
}