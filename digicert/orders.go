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
	ID                          int          `json:"id"`
	Certificate                 Certificate  `json:"certificate,omitempty"`
	ValidityYears               int          `json:"validity_years"`
	ValidityDays                int          `json:"validity_days"`
	CustomExpirationDate        string       `json:"custom_expiration_date,omitempty"`
	Comments                    string       `json:"comments,omitempty"`
	AutoRenew                   int          `json:"auto_renew,omitempty"`
	DisableRenewalNotifications bool         `json:"disable_renewal_notifications,omitempty"`
	RenewedOrderID              int          `json:"renewed_order_id,omitempty"`
	RenewalOfOrderID            int          `json:"renewal_of_order_id,omitempty"`
	PaymentMethod               string       `json:"payment_method,omitempty"`
	Status                      string       `json:"status,omitempty"`
	Organization                Organization `json:"organization"`
	OrganizationContact         *Contact     `json:"organization_contact,omitempty"`
	TechnicalContact            *Contact     `json:"technical_contact,omitempty"`
	// These two fields are perhaps candidates for their own structs. The
	// Digicert API is not consistent in what properties it returns!
	Requests     []*OrderRequest `json:"requests,omitempty"`
	CertficateID int             `json:"certificate_id,omitempty"`
	Product      Product
	Container    Container `json:"container,omitempty"`
}

type OrderRequest struct {
	ID     int    `json:"id"`
	Status string `json:"status"`
}

type NewOrder struct {
	Certificate                 Certificate     `json:"certificate"`
	RenewedThumbprint           string          `json:"renewed_thumbprint,omitempty"`
	ValidityYears               int             `json:"validity_years,omitempty"`
	ValidityDays                int             `json:"validity_days,omitempty"`
	CustomExpirationDate        string          `json:"custom_expiration_date,omitempty"`
	Comments                    string          `json:"comments,omitempty"`
	AutoRenew                   int             `json:"auto_renew,omitempty"`
	CustomRenewalMessage        string          `json:"custom_renewal_message,omitempty"`
	DisableRenewalNotifications bool            `json:"disable_renewal_notifications"`
	AdditionalEmails            []string        `json:"additional_emails,omitempty"`
	RenewalOrderID              int             `json:"renewal_order_id,omitempty"`
	PaymentMethod               string          `json:"payment_method,omitempty"`
	PaymentProfile              *PaymentProfile `json:"payment_profile,omitempty"`
	DCVMethod                   string          `json:"dcv_method,omitempty"`
	SkipApproval                bool            `json:"skip_approval"`
	DisableCT                   bool            `json:"disable_ct"`
	Organization                Organization    `json:"organization"`
	Container                   *Container      `json:"container,omitempty"`
}

type PaymentProfile struct {
	ID           int    `json:"id,omitempty"`
	IsDefault    bool   `json:"is_default"`
	Status       string `json:"status,omitempty"`
	BillingEmail string `json:"billing_email,omitempty"`
}

func InitializeOrder() *NewOrder {
	order := new(NewOrder)
	order.setNewOrderDefaults()
	return order
}

func (order *NewOrder) setNewOrderDefaults() {
	if order.ValidityYears == 0 {
		order.ValidityYears = defaultCertificateTTL
	}

	if order.PaymentMethod == "" {
		order.PaymentMethod = "balance"
	}
}

type orderOrganizationAttrs struct {
	ID int `json:"id"`
}

type orderList struct {
	Orders []*Order
}

func (o Order) String() string {
	return Stringify(o)
}

func (s *OrdersService) Get(order_id int) (*Order, *Response, error) {
	order := &Order{}
	resp, err := executeAction(s.client, "GET", fmt.Sprintf("order/certificate/%d", order_id), nil, order)
	if err != nil {
		return nil, resp, err
	}
	return order, resp, nil
}

func (s *OrdersService) List() ([]*Order, *Response, error) {
	list := new(orderList)
	resp, err := executeAction(s.client, "GET", "order/certificate", nil, list)
	if err != nil {
		return nil, resp, err
	}

	return list.Orders, resp, nil
}

func (s *OrdersService) Create(orderReq *NewOrder, orderType string) (*Order, *Response, error) {
	order := new(Order)
	resp, err := executeAction(s.client, "POST", fmt.Sprintf("order/certificate/%s", orderType), orderReq, order)
	if err != nil {
		return nil, resp, err
	}
	return order, resp, nil
}

func (s *OrdersService) CreateWildcard(orderReq *NewOrder) (*Order, *Response, error) {
	return s.Create(orderReq, "ssl_wildcard")
}

func (s *OrdersService) CreateStandard(orderReq *NewOrder) (*Order, *Response, error) {
	return s.Create(orderReq, "ssl_plus")
}
