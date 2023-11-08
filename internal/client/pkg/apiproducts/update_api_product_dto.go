package apiproducts

type UpdateApiProductDto struct {
	Name        *string           `json:"name,omitempty"`
	Description *string           `json:"description,omitempty"`
	Labels      *ApiProductLabels `json:"labels,omitempty"`
	PortalIds   *[]string         `json:"portal_ids,omitempty"`
}

func (u *UpdateApiProductDto) SetName(name string) {
	u.Name = &name
}

func (u *UpdateApiProductDto) GetName() *string {
	if u == nil {
		return nil
	}
	return u.Name
}

func (u *UpdateApiProductDto) SetDescription(description string) {
	u.Description = &description
}

func (u *UpdateApiProductDto) GetDescription() *string {
	if u == nil {
		return nil
	}
	return u.Description
}

func (u *UpdateApiProductDto) SetLabels(labels ApiProductLabels) {
	u.Labels = &labels
}

func (u *UpdateApiProductDto) GetLabels() *ApiProductLabels {
	if u == nil {
		return nil
	}
	return u.Labels
}

func (u *UpdateApiProductDto) SetPortalIds(portalIds []string) {
	u.PortalIds = &portalIds
}

func (u *UpdateApiProductDto) GetPortalIds() *[]string {
	if u == nil {
		return nil
	}
	return u.PortalIds
}
