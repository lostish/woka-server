package schemas

import (
	"net/url"
	"time"

	"github.com/surrealdb/surrealdb.go/pkg/models"
)

const (
	BusinessOwner           = "owner"
	BusinessManager         = "manager"
	BusinessFinancer        = "financer"
	BusinessAnalyst         = "analyst"
	BusinessCustomerService = "customer-service"
	BusinessViewer          = "viewer"
) // Add more if necessary

type _BusinessAddress struct {
	LocalName string
	Address   string
}

type Business struct {
	ID          *models.RecordID                      `json:"id,omitempty"`
	UserId      *models.RecordID                      `json:"user_id"`
	Name        string                                `json:"name"`
	Description string                                `json:"description"`
	Locals      map[_BusinessAddress]_BusinessAddress `json:"address"`
	CreatedAt   time.Time                             `json:"created_at"`
	UpdatedAt   time.Time                             `json:"updated_at"`
}

type BusinessMembers struct {
	ID         *models.RecordID `json:"id,omitempty"`
	BusinessId *models.RecordID `json:"business_id"`
	MemberId   *models.RecordID `json:"member_id"`
	Role       string           `json:"role_id"`
	Status     string           `json:"status"`
	CreatedAt  time.Time        `json:"created_at"`
	UpdatedAt  time.Time        `json:"updated_at"`
}

type Categories struct {
	ID          *models.RecordID `json:"id,omitempty"`
	BusinessId  *models.RecordID `json:"business_id"`
	Name        string           `json:"name"`
	Description string           `json:"description"`
	Status      string           `json:"status"`
	CreatedAt   time.Time        `json:"created_at"`
	UpdatedAt   time.Time        `json:"updated_at"`
}

type _PreviewVariants struct {
	Name string
	Url  url.URL
}

type Product struct {
	ID              *models.RecordID                      `json:"id,omitempty"`
	BusinessId      *models.RecordID                      `json:"business_id"`
	CategoryId      *models.RecordID                      `json:"category_id"`
	Name            string                                `json:"name"`
	Description     string                                `json:"description"`
	Tags            string                                `json:"tags"`
	ShortThumbnail  url.URL                               `json:"short_thumbnail"`
	HighThumbnail   url.URL                               `json:"high_thumbnail"`
	PreviewVariants map[_PreviewVariants]_PreviewVariants `json:"preview_variants"`
	CreatedAt       time.Time                             `json:"created_at"`
	UpdatedAt       time.Time                             `json:"updated_at"`
}

type _DiscountsData struct {
	Value    float32
	StartsIn time.Time
	EndsIn   time.Time
}

type _Discounts struct {
	Name string
	Data _DiscountsData
}

type ProductMetadata struct {
	ID           *models.RecordID          `json:"id,omitempty"`
	ProductId    *models.RecordID          `json:"product_id"`
	Price        uint32                    `json:"price"`
	Discounts    map[_Discounts]_Discounts `json:"discounts"`
	CurrentStock uint16                    `json:"current_stock"`
	MaximumStock uint16                    `json:"maximum_stock"`
	CreatedAt    time.Time                 `json:"created_at"`
	UpdatedAt    time.Time                 `json:"updated_at"`
}

type ProductFeedback struct {
	ID        *models.RecordID `json:"id,omitempty"`
	ProductId *models.RecordID `json:"product_id"`
}
