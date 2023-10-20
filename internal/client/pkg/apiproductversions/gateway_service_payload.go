package apiproductversions

type GatewayServicePayload struct {
	Id             *string `json:"id,omitempty"`
	ControlPlaneId *string `json:"control_plane_id,omitempty"`
}

func (g *GatewayServicePayload) SetId(id string) {
	g.Id = &id
}

func (g *GatewayServicePayload) SetControlPlaneId(controlPlaneId string) {
	g.ControlPlaneId = &controlPlaneId
}
