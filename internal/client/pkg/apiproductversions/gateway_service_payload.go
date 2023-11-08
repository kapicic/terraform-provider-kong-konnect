package apiproductversions

type GatewayServicePayload struct {
	Id             *string `json:"id,omitempty"`
	ControlPlaneId *string `json:"control_plane_id,omitempty"`
}

func (g *GatewayServicePayload) SetId(id string) {
	g.Id = &id
}

func (g *GatewayServicePayload) GetId() *string {
	if g == nil {
		return nil
	}
	return g.Id
}

func (g *GatewayServicePayload) SetControlPlaneId(controlPlaneId string) {
	g.ControlPlaneId = &controlPlaneId
}

func (g *GatewayServicePayload) GetControlPlaneId() *string {
	if g == nil {
		return nil
	}
	return g.ControlPlaneId
}
