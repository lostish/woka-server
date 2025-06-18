package schemas

import (
	"time"

	"github.com/surrealdb/surrealdb.go/pkg/models"
)

const (
	PREPARING_TASK = "preparing"
	REJECTED_TASK  = "rejected"
	READY_TASK     = "ready"
	STARTING_TASK  = "starting"
	PENDING_TASK   = "pending"
	FAILED_TASK    = "failed"
	REMOVE_TASK    = "remove"
	COMPLETE_TASK  = "complete"
)

type Tasks struct {
	ID          *models.RecordID `json:"id,omitempty"`
	BusinessId  *models.RecordID `json:"business_id,"`
	By          *models.RecordID `json:"by"`
	Name        string           `json:"name"`
	Description string           `json:"description"`
	Status      string           `json:"state"`
	CreatedAt   time.Time        `json:"created_at"`
	UpdatedAt   time.Time        `json:"updated_at"`
}
