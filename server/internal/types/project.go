package types

import (
	"github.com/google/uuid"
)

type Project struct {
	Guid        uuid.UUID `json:"guid"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	TimeSpent   float64   `json:"time_spent"`
}

type ProjectStatus struct {
	Total       int     `json:"total"`
	AverageTime float64 `json:"average_time"`
	Completed   int     `json:"completed"`
}
