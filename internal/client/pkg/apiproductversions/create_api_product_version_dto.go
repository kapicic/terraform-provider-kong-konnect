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

func (c *CreateApiProductVersionDto) SetPublishStatus(publishStatus PublishStatus) {
	c.PublishStatus = &publishStatus
}

func (c *CreateApiProductVersionDto) SetDeprecated(deprecated bool) {
	c.Deprecated = &deprecated
}

func (c *CreateApiProductVersionDto) SetGatewayService(gatewayService GatewayServicePayload) {
	c.GatewayService = &gatewayService
}

func (c *CreateApiProductVersionDto) GetGatewayService() *GatewayServicePayload {
	if c.GatewayService == nil {
		return nil
	}
	return c.GatewayService
}

type PublishStatus string

const (
	PUBLISH_STATUS_UNPUBLISHED PublishStatus = "unpublished"
	PUBLISH_STATUS_PUBLISHED   PublishStatus = "published"
)
