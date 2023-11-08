package apiproducts

type CreateApiProductDto struct {
	Name        *string           `json:"name,omitempty"`
	Description *string           `json:"description,omitempty"`
	Labels      *ApiProductLabels `json:"labels,omitempty"`
}

func (c *CreateApiProductDto) SetName(name string) {
	c.Name = &name
}

func (c *CreateApiProductDto) GetName() *string {
	if c == nil {
		return nil
	}
	return c.Name
}

func (c *CreateApiProductDto) SetDescription(description string) {
	c.Description = &description
}

func (c *CreateApiProductDto) GetDescription() *string {
	if c == nil {
		return nil
	}
	return c.Description
}

func (c *CreateApiProductDto) SetLabels(labels ApiProductLabels) {
	c.Labels = &labels
}

func (c *CreateApiProductDto) GetLabels() *ApiProductLabels {
	if c == nil {
		return nil
	}
	return c.Labels
}
