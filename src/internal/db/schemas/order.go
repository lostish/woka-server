package schemas

import (
	"time"

	"github.com/surrealdb/surrealdb.go/pkg/models"
)

const (
	PENDING_ORDER   = "pending"
	PAID_ORDER      = "paid"
	SHIPPED_ORDER   = "shipped"
	DELIVERED_ORDER = "delivered"
	CANCELED_ORDER  = "canceled"
	REFUNDED_ORDER  = "refunded"
	//
	RETURN_REQUEST  = "requested"
	RETURN_APPROVED = "approved"
	RETURN_REJECTED = "rejected"
	RETURN_REFUNDED = "refunded"
)

type Order struct {
	ID            *models.RecordID `json:"id,omitempty"`
	BusinessId    *models.RecordID `json:"business_id"`
	CustomerId    *models.RecordID `json:"customer_id"`
	Status        string           `json:"status"`
	TotalAmount   float64          `json:"total_amount"`
	Currency      string           `json:"currency"`       // ISO: "USD", "PEN", etc.
	PaymentMethod string           `json:"payment_method"` // E.g. card token, "paypal", etc.
	ShippingAddr  string           `json:"shipping_addr"`
	Notes         string           `json:"notes"`
	CreatedAt     time.Time        `json:"created_at"`
	UpdatedAt     time.Time        `json:"updated_at"`
}

type OrderItem struct {
	ID         *models.RecordID `json:"id,omitempty"`
	OrderId    *models.RecordID `json:"order_id"`
	ProductId  *models.RecordID `json:"product_id"`
	Quantity   uint16           `json:"quantity"`
	UnitPrice  float64          `json:"unit_price"`  // current price
	TotalPrice float64          `json:"total_price"` // unitPrice * quantity
	MetadataId *models.RecordID `json:"metadata_id"` // relation to `ProductMetadata`
	CreatedAt  time.Time        `json:"created_at"`
}

type OrderStatusHistory struct {
	ID        *models.RecordID `json:"id,omitempty"`
	OrderId   *models.RecordID `json:"order_id"`
	Status    string           `json:"status"`     // ORDER.status
	ChangedBy *models.RecordID `json:"changed_by"` // business member or client
	Timestamp time.Time        `json:"timestamp"`
	Notes     string           `json:"notes"`
}

type OrderReturn struct {
	ID          *models.RecordID `json:"id,omitempty"`
	OrderId     *models.RecordID `json:"order_id"`
	Reason      string           `json:"reason"`
	Status      string           `json:"status"` // requested, approved, rejected, refunded
	RequestedAt time.Time        `json:"requested_at"`
}
