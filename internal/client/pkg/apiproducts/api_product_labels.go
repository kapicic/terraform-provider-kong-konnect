package apiproducts

type ApiProductLabels struct {
	Name *string `json:"name,omitempty"`
}

func (a *ApiProductLabels) SetName(name string) {
	a.Name = &name
}

func (a *ApiProductLabels) GetName() *string {
	if a == nil {
		return nil
	}
	return a.Name
}
