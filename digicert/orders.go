package digicert

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
	Status                      string
}

type orderList struct {
	Orders *[]Order
}

func (o Order) String() string {
	return Stringify(o)
}

func (s *OrdersService) List() (*[]Order, *Response, error) {
	req, err := s.client.NewRequest("GET", "order/certificate", nil)
	if err != nil {
		return nil, nil, err
	}

	list := new(orderList)
	resp, err := s.client.Do(req, list)
	if err != nil {
		return nil, resp, err
	}

	return list.Orders, resp, nil
}
