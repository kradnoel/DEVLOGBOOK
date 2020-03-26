package types

import (
	"github.com/google/uuid"
)

type Task struct {
	Guid        uuid.UUID `json:"guid"`
	Description string    `json:"description"`
	TimeSpent   float64   `json:"time_spent"`
	ProjectId   uuid.UUID `json:"project_id"`
}

type FormTask struct {
	Description string `json:"description"`
	ProjectId   string `json:"project_id"`
}
