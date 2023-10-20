package apiproducts

type CreateApiProductDto struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Labels      *any    `json:"labels,omitempty"`
}

func (c *CreateApiProductDto) SetName(name string) {
	c.Name = &name
}

func (c *CreateApiProductDto) SetDescription(description string) {
	c.Description = &description
}

func (c *CreateApiProductDto) SetLabels(labels any) {
	c.Labels = &labels
}
