package digicert

import (
	"fmt"
)

const (
	defaultCertificateTTL = 1
	defaultPaymentMethod  = "balance"
)

type OrdersService service

type Order struct {
	ID                          int64         `json:"id"`
	Certificate                 *Certificate  `json:"certificate,omitempty"`
	ValidityYears               int           `json:"validity_years"`
	CustomExpirationDate        string        `json:"custom_expiration_date,omitempty"`
	Comments                    string        `json:"comments,omitempty"`
	DisableRenewalNotifications bool          `json:"disable_renewal_notifications,omitempty"`
	RenewalOfOrderID            int           `json:"renewal_of_order_id,omitempty"`
	PaymentMethod               string        `json:"payment_method,omitempty"`
	Status                      string        `json:"status,omitempty"`
	Organization                *Organization `json:"organization"`
	// These two fields are perhaps candidates for their own structs. The
	// Digicert API is not consistent in what properties it returns!
	Requests     []*OrderRequest `json:"requests,omitempty"`
	CertficateID int64           `json:"certificate_id,omitempty"`
}

type OrderRequest struct {
	ID     int64  `json:"id"`
	Status string `json:"status"`
}

type NewOrder struct {
	Certificate                 Certificate            `json:"certificate"`
	RenewedThumbprint           string                 `json:"renewed_thumbprint,omitempty"`
	ValidityYears               int                    `json:"validity_years"`
	CustomExpirationDate        string                 `json:"custom_expiration_date,omitempty"`
	Comments                    string                 `json:"comments,omitempty"`
	AutoRenew                   int                    `json:"auto_renew,omitempty"`
	CustomRenewalMessage        string                 `json:"custom_renewal_message,omitempty"`
	DisableRenewalNotifications bool                   `json:"disable_renewal_notifications"`
	AdditionalEmails            []string               `json:"additional_emails,omitempty"`
	RenewalOrderID              string                 `json:"renewal_order_id,omitempty"`
	PaymentMethod               string                 `json:"payment_method"`
	DCVMethod                   string                 `json:"dcv_method,omitempty"`
	SkipApproval                bool                   `json:"skip_approval"`
	DisableCT                   bool                   `json:"disable_ct"`
	Organization                orderOrganizationAttrs `json:"organization"`
	Container                   struct {
		ID int `json:"id"`
	} `json:"container"`
}

func InitializeOrder() *NewOrder {
	order := new(NewOrder)
	//order.Organization = new(orderOrganizationAttrs)
	//order.Container = new(Container)
	order.setNewOrderDefaults()
	return order
}

func (order *NewOrder) setNewOrderDefaults() {
	order.SkipApproval = false // FIXME force this for now
	if order.ValidityYears == 0 {
		order.ValidityYears = defaultCertificateTTL
	}

	if order.PaymentMethod == "" {
		order.PaymentMethod = "balance"
	}
}

type orderOrganizationAttrs struct {
	ID       int64              `json:"id"`
	Contacts *[]newOrderContact `json:"contacts"`
}

type newOrderContact struct {
	UserID      int64  `json:"user_id,omitempty"`
	ContactType string `json:"contact_type,omitempty"`
	FirstName   string `json:"first_name,omitempty"`
	LastName    string `json:"last_name,omitempty"`
	JobTitle    string `json:"job_title,omitempty"`
	Telephone   string `json:"telephone,omitempty"`
	Email       string `json:"email,omitempty"`
}

type orderList struct {
	Orders *[]Order
}

func (o Order) String() string {
	return Stringify(o)
}

func (s *OrdersService) Get(order_id int64) (*Order, *Response, error) {
	order := &Order{}
	resp, err := executeAction(s.client, "GET", fmt.Sprintf("order/certificate/%d", order_id), nil, order)
	if err != nil {
		return nil, resp, err
	}
	return order, resp, nil
}

func (s *OrdersService) List() (*[]Order, *Response, error) {
	list := new(orderList)
	resp, err := executeAction(s.client, "GET", "order/certificate", nil, list)
	if err != nil {
		return nil, resp, err
	}

	return list.Orders, resp, nil
}

func (s *OrdersService) Create(orderReq *NewOrder) (*Order, *Response, error) {
	order := new(Order)
	resp, err := executeAction(s.client, "POST", "order/certificate/ssl_plus", orderReq, order)
	if err != nil {
		return nil, resp, err
	}
	return order, resp, nil
}
