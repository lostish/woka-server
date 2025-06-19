package schemas

import (
	"net/url"
	"time"

	"github.com/surrealdb/surrealdb.go/pkg/models"
)

const (
	CRT_URL               = "url"
	CRT_IMAGE             = "image"
	RESOURCE_TYPE_PRODUCT = "product"
	RESOURCE_TYPE_TASK    = "task"
)

type _CommentResources struct {
	Name string
	Type string // CRT_URL | CRT_IMAGE | ...
	Src  url.URL
}

type Comment struct {
	ID        *models.RecordID                        `json:"id,omitempty"`
	By        *models.RecordID                        `json:"by"`
	Body      string                                  `json:"body"`
	Resources map[_CommentResources]_CommentResources `json:"resources"`
	CreatedAt time.Time                               `json:"created_at"`
}

// Polymorphic association (generic intermediate table)
type CommentLink struct {
	ID           *models.RecordID `json:"id,omitempty"`
	CommentId    *models.RecordID `json:"comment_id"`
	ResourceType string           `json:"resource_type"` // e.g. "product", "task", "order"
	ResourceId   *models.RecordID `json:"resource_id"`
	CreatedAt    time.Time        `json:"created_at"`
}
