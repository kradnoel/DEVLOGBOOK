package types

import (
	"time"

	"github.com/google/uuid"
)

type Subtask struct {
	Guid   uuid.UUID `json:"guid"`
	From   time.Time `json:"from"`
	To     time.Time `json:"to"`
	TaskId uuid.UUID `json:"task_id"`
}

type FormSubtask struct {
	From   string `json:"from"`
	To     string `json:"to"`
	TaskId string `json:"task_id"`
}
