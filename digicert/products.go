package digicert

type Product struct {
	GroupName            string
	NameID               string `json:"name_id"`
	Name                 string
	Type                 string
	AllowedContainerIDs  []int `json:"allowed_container_ids"`
	AllowedValidityYears []int `json:"allowed_validity_years"`
}

type productList struct {
	Products []*Product
}

type ProductsService service

func (s *ProductsService) List() ([]*Product, *Response, error) {
	list := new(productList)
	resp, err := executeAction(s.client, "GET", "product", nil, list)
	if err != nil {
		return nil, resp, err
	}

	return list.Products, resp, nil
}
