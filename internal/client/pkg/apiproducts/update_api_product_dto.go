package apiproducts

type UpdateApiProductDto struct {
	Name        *string   `json:"name,omitempty"`
	Description *string   `json:"description,omitempty"`
	Labels      *any      `json:"labels,omitempty"`
	PortalIds   *[]string `json:"portal_ids,omitempty"`
}

func (u *UpdateApiProductDto) SetName(name string) {
	u.Name = &name
}

func (u *UpdateApiProductDto) SetDescription(description string) {
	u.Description = &description
}

func (u *UpdateApiProductDto) SetLabels(labels any) {
	u.Labels = &labels
}

func (u *UpdateApiProductDto) SetPortalIds(portalIds []string) {
	u.PortalIds = &portalIds
}
