package apiproducts


type ApiProduct struct {
  Id *string `json:"id,omitempty"`
  Name *string `json:"name,omitempty"`
  Description *string `json:"description,omitempty"`
  PortalIds *[]string `json:"portal_ids,omitempty"`
  CreatedAt *string `json:"created_at,omitempty"`
  UpdatedAt *string `json:"updated_at,omitempty"`
  Labels *ApiProductLabels `json:"labels,omitempty"`
}

func (a *ApiProduct) SetId(id string) {
  a.Id = &id
}

func (a *ApiProduct) GetId() *string {
  if a == nil {
    return nil
  }
  return a.Id
}

func (a *ApiProduct) SetName(name string) {
  a.Name = &name
}

func (a *ApiProduct) GetName() *string {
  if a == nil {
    return nil
  }
  return a.Name
}

func (a *ApiProduct) SetDescription(description string) {
  a.Description = &description
}

func (a *ApiProduct) GetDescription() *string {
  if a == nil {
    return nil
  }
  return a.Description
}

func (a *ApiProduct) SetPortalIds(portalIds []string) {
  a.PortalIds = &portalIds
}

func (a *ApiProduct) GetPortalIds() *[]string {
  if a == nil {
    return nil
  }
  return a.PortalIds
}

func (a *ApiProduct) SetCreatedAt(createdAt string) {
  a.CreatedAt = &createdAt
}

func (a *ApiProduct) GetCreatedAt() *string {
  if a == nil {
    return nil
  }
  return a.CreatedAt
}

func (a *ApiProduct) SetUpdatedAt(updatedAt string) {
  a.UpdatedAt = &updatedAt
}

func (a *ApiProduct) GetUpdatedAt() *string {
  if a == nil {
    return nil
  }
  return a.UpdatedAt
}

func (a *ApiProduct) SetLabels(labels ApiProductLabels) {
  a.Labels = &labels
}

func (a *ApiProduct) GetLabels() *ApiProductLabels {
  if a == nil {
    return nil
  }
  return a.Labels
}



