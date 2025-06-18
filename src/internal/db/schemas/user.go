package schemas

import (
	"net/url"
	"time"

	"github.com/surrealdb/surrealdb.go/pkg/models"
)

const (
	FRIENDSHIP_PENDING  = "pending"
	FRIENDSHIP_ACCEPTED = "accepted"
	FRIENDSHIP_BLOCKED  = "blocked"
)

type User struct {
	ID            *models.RecordID `json:"id,omitempty"`
	Name          string           `json:"name"`
	Email         string           `json:"email"`
	EmailVerified time.Duration    `json:"email_verified"`
	Pwd           string           `json:"pwd"`
	CreatedAt     time.Time        `json:"created_at"`
	UpdatedAt     time.Time        `json:"updated_at"`
}

type UserProfile struct {
	ID        *models.RecordID `json:"id,omitempty"`
	UserId    *models.RecordID `json:"user_id"`
	Avatar    url.URL          `json:"avatar"`
	Banner    url.URL          `json:"banner"`
	CreatedAt time.Time        `json:"created_at"`
	UpdatedAt time.Time        `json:"updated_at"`
}

type UserPhone struct {
	ID             *models.RecordID `json:"id,omitempty"`
	UserID         *models.RecordID `json:"user_id"`
	PhoneHash      string           `json:"phone_hash"` // hash o encryption
	Label          string           `json:"label"`      // "mobile", "work", ...
	Verified       bool             `json:"verified"`
	IsPrimaryPhone bool             `json:"is_primary_phone"`
	CreatedAt      time.Time        `json:"created_at"`
	UpdatedAt      time.Time        `json:"updated_at"`
}

type Friendship struct {
	ID        *models.RecordID `json:"id,omitempty"`
	UserID    *models.RecordID `json:"user_id"`
	FriendID  *models.RecordID `json:"friend_id"`
	Status    string           `json:"status"` // pending, accepted, blocked, ...
	CreatedAt time.Time        `json:"created_at"`
	UpdatedAt time.Time        `json:"updated_at"`
}
