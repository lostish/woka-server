package schemas

import (
	"time"

	"github.com/surrealdb/surrealdb.go/pkg/models"
)

const (
	ActionRead              = "read"
	ActionWrite             = "write"
	ActionDelete            = "delete"
	ActionApproved          = "approved"
	ActionShare             = "share"
	EffectAllow             = "allow"
	EffectDeny              = "deny"
	SeniorityJunior         = "junior"
	SeniorityMid            = "mid"
	SenioritySenior         = "senior"
	PublicSensitivity       = "public"
	InternalSensitivity     = "internal"
	ConfidentialSensitivity = "confidential"
	ResourceDarfStatus      = "darf"
	ResourceActiveStatus    = "active"
	ResourceArchivedStatus  = "archived"
) // Add more if necessary

type SubjectAttributes struct {
	Role              string    `json:"role"`
	Department        string    `json:"department"`
	Seniority         string    `json:"seniority"`
	JoinDate          time.Time `json:"join_date"`
	SecurityClearance uint8     `json:"security_clearance"` // (0-5)
	OwnedResources    []string  `json:"owned_resources"`
	RestrictedActions []string  `json:"restricted_actions"`
}

type ResourceAttributes struct {
	Type         string           `json:"type"`
	Sensitivity  string           `json:"sensitivity"`
	OwnerID      *models.RecordID `json:"owner_id"`
	Department   string           `json:"department"`
	CreationDate time.Time        `json:"creation_date"`
	BusinessUnit string           `json:"business_unit"`
	Location     string           `json:"location"`
	Status       string           `json:"status"`
}

type ContextAttributes struct {
	Time            time.Time `json:"time"`
	DayOfWeek       string    `json:"day_of_week"`
	IPRange         string    `json:"ip_range"`
	DeviceType      string    `json:"device_type"`
	Location        string    `json:"location"`
	Authentication  string    `json:"authentication"`
	SessionDuration float64   `json:"session_duration"`
	BusinessHours   bool      `json:"business_hours"`
}

type ABACPolicy struct {
	ID         *models.RecordID   `json:"id,omitempty"`
	BusinessId *models.RecordID   `json:"business_id"`
	MemberId   *models.RecordID   `json:"member_id"`
	Subject    SubjectAttributes  `json:"subject"`
	Resources  ResourceAttributes `json:"resources"`
	Action     string             `json:"action"`
	Context    ContextAttributes  `json:"context"`
	Effect     string             `json:"effect"`
	CreatedAt  time.Time          `json:"created_at"`
	UpdatedAt  time.Time          `json:"updated_at"`
	ExpiresAt  *time.Time         `json:"expires_at"`
}
