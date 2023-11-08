package apiproductversions


type ApiProductVersion struct {
  Id *string `json:"id,omitempty"`
  Name *string `json:"name,omitempty"`
  GatewayService *GatewayServicePayload `json:"gateway_service,omitempty"`
  PublishStatus *PublishStatus1 `json:"publish_status,omitempty"`
  Deprecated *bool `json:"deprecated,omitempty"`
  CreatedAt *string `json:"created_at,omitempty"`
  UpdatedAt *string `json:"updated_at,omitempty"`
}

func (a *ApiProductVersion) SetId(id string) {
  a.Id = &id
}

func (a *ApiProductVersion) GetId() *string {
  if a == nil {
    return nil
  }
  return a.Id
}

func (a *ApiProductVersion) SetName(name string) {
  a.Name = &name
}

func (a *ApiProductVersion) GetName() *string {
  if a == nil {
    return nil
  }
  return a.Name
}

func (a *ApiProductVersion) SetGatewayService(gatewayService GatewayServicePayload) {
  a.GatewayService = &gatewayService
}

func (a *ApiProductVersion) GetGatewayService() *GatewayServicePayload {
  if a == nil {
    return nil
  }
  return a.GatewayService
}

func (a *ApiProductVersion) SetPublishStatus(publishStatus PublishStatus1) {
  a.PublishStatus = &publishStatus
}

func (a *ApiProductVersion) GetPublishStatus() *PublishStatus1 {
  if a == nil {
    return nil
  }
  return a.PublishStatus
}

func (a *ApiProductVersion) SetDeprecated(deprecated bool) {
  a.Deprecated = &deprecated
}

func (a *ApiProductVersion) GetDeprecated() *bool {
  if a == nil {
    return nil
  }
  return a.Deprecated
}

func (a *ApiProductVersion) SetCreatedAt(createdAt string) {
  a.CreatedAt = &createdAt
}

func (a *ApiProductVersion) GetCreatedAt() *string {
  if a == nil {
    return nil
  }
  return a.CreatedAt
}

func (a *ApiProductVersion) SetUpdatedAt(updatedAt string) {
  a.UpdatedAt = &updatedAt
}

func (a *ApiProductVersion) GetUpdatedAt() *string {
  if a == nil {
    return nil
  }
  return a.UpdatedAt
}

type PublishStatus1 string

const (
  PUBLISH_STATUS1_UNPUBLISHED PublishStatus1 = "unpublished"
  PUBLISH_STATUS1_PUBLISHED PublishStatus1 = "published"
)




