package apiproducts

type ApiProduct struct {
	Id          *string           `json:"id,omitempty"`
	Name        *string           `json:"name,omitempty"`
	Description *string           `json:"description,omitempty"`
	PortalIds   *[]string         `json:"portal_ids,omitempty"`
	CreatedAt   *string           `json:"created_at,omitempty"`
	UpdatedAt   *string           `json:"updated_at,omitempty"`
	Labels      *ApiProductLabels `json:"labels,omitempty"`
}

func (a *ApiProduct) SetId(id string) {
	a.Id = &id
}

func (a *ApiProduct) SetName(name string) {
	a.Name = &name
}

func (a *ApiProduct) SetDescription(description string) {
	a.Description = &description
}

func (a *ApiProduct) SetPortalIds(portalIds []string) {
	a.PortalIds = &portalIds
}

func (a *ApiProduct) SetCreatedAt(createdAt string) {
	a.CreatedAt = &createdAt
}

func (a *ApiProduct) SetUpdatedAt(updatedAt string) {
	a.UpdatedAt = &updatedAt
}

func (a *ApiProduct) SetLabels(labels ApiProductLabels) {
	a.Labels = &labels
}

func (a *ApiProduct) GetLabels() *ApiProductLabels {
	if a.Labels == nil {
		return nil
	}
	return a.Labels
}
