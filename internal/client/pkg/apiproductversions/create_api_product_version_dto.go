package apiproductversions

type CreateApiProductVersionDto struct {
	Name           *string                `json:"name,omitempty"`
	PublishStatus  *PublishStatus         `json:"publish_status,omitempty"`
	Deprecated     *bool                  `json:"deprecated,omitempty"`
	GatewayService *GatewayServicePayload `json:"gateway_service,omitempty"`
}

func (c *CreateApiProductVersionDto) SetName(name string) {
	c.Name = &name
}

func (c *CreateApiProductVersionDto) GetName() *string {
	if c == nil {
		return nil
	}
	return c.Name
}

func (c *CreateApiProductVersionDto) SetPublishStatus(publishStatus PublishStatus) {
	c.PublishStatus = &publishStatus
}

func (c *CreateApiProductVersionDto) GetPublishStatus() *PublishStatus {
	if c == nil {
		return nil
	}
	return c.PublishStatus
}

func (c *CreateApiProductVersionDto) SetDeprecated(deprecated bool) {
	c.Deprecated = &deprecated
}

func (c *CreateApiProductVersionDto) GetDeprecated() *bool {
	if c == nil {
		return nil
	}
	return c.Deprecated
}

func (c *CreateApiProductVersionDto) SetGatewayService(gatewayService GatewayServicePayload) {
	c.GatewayService = &gatewayService
}

func (c *CreateApiProductVersionDto) GetGatewayService() *GatewayServicePayload {
	if c == nil {
		return nil
	}
	return c.GatewayService
}

type PublishStatus string

const (
	PUBLISH_STATUS_UNPUBLISHED PublishStatus = "unpublished"
	PUBLISH_STATUS_PUBLISHED   PublishStatus = "published"
)
