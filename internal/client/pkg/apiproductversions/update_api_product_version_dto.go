package apiproductversions

type UpdateApiProductVersionDto struct {
	Name           *string                `json:"name,omitempty"`
	PublishStatus  *PublishStatus2        `json:"publish_status,omitempty"`
	Deprecated     *bool                  `json:"deprecated,omitempty"`
	Notify         *bool                  `json:"notify,omitempty"`
	GatewayService *GatewayServicePayload `json:"gateway_service,omitempty"`
}

func (u *UpdateApiProductVersionDto) SetName(name string) {
	u.Name = &name
}

func (u *UpdateApiProductVersionDto) SetPublishStatus(publishStatus PublishStatus2) {
	u.PublishStatus = &publishStatus
}

func (u *UpdateApiProductVersionDto) SetDeprecated(deprecated bool) {
	u.Deprecated = &deprecated
}

func (u *UpdateApiProductVersionDto) SetNotify(notify bool) {
	u.Notify = &notify
}

func (u *UpdateApiProductVersionDto) SetGatewayService(gatewayService GatewayServicePayload) {
	u.GatewayService = &gatewayService
}

func (u *UpdateApiProductVersionDto) GetGatewayService() *GatewayServicePayload {
	if u.GatewayService == nil {
		return nil
	}
	return u.GatewayService
}

type PublishStatus2 string

const (
	PUBLISH_STATUS2_UNPUBLISHED PublishStatus2 = "unpublished"
	PUBLISH_STATUS2_PUBLISHED   PublishStatus2 = "published"
)
