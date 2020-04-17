package digicert

import (
	"fmt"
)

type OrdersService service

type Order struct {
	ID                          int64        `json:"id"`
	Certificate                 *Certificate `json:"certificate,omitempty"`
	ValidityYears               int          `json:"validity_years"`
	CustomExpirationDate        string       `json:"custom_expiration_date,omitempty"`
	Comments                    string       `json:"comments,omitempty"`
	DisableRenewalNotifications bool         `json:"disable_renewal_notifications,omitempty"`
	RenewalOfOrderID            int          `json:"renewal_of_order_id,omitempty"`
	PaymentMethod               string       `json:"payment_method,omitempty"`
	Status                      string       `json:"status,omitempty"`
	// These two fields are perhaps candidates for their own structs. The
	// Digicert API is not consistent in what properties it returns!
	Requests     *[]OrderRequest `json:"requests,omitempty"`
	CertficateID int64           `json:"certificate_id,omitempty"`
}

type OrderRequest struct {
	ID     int64  `json:"id"`
	Status string `json:"status"`
}

type orderList struct {
	Orders *[]Order
}

func (o Order) String() string {
	return Stringify(o)
}

func (s *OrdersService) Get(order_id int64) (*Order, *Response, error) {
	order := &Order{}
	resp, err := s.reqHelper("GET", fmt.Sprintf("order/certificate/%d", order_id), nil, order)
	if err != nil {
		return nil, resp, err
	}
	return order, resp, nil
}

func (s *OrdersService) List() (*[]Order, *Response, error) {
	list := new(orderList)
	resp, err := s.reqHelper("GET", "order/certificate", nil, list)
	if err != nil {
		return nil, resp, err
	}

	return list.Orders, resp, nil
}

func (s *OrdersService) reqHelper(method, path string, body, v interface{}) (*Response, error) {
	req, err := s.client.NewRequest(method, path, body)
	if err != nil {
		return nil, err
	}
	return s.client.Do(req, v)
}
