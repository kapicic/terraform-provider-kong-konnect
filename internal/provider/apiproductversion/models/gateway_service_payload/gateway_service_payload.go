package gateway_service_payload

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type GatewayServicePayload struct {
	Id             types.String `tfsdk:"id"`
	ControlPlaneId types.String `tfsdk:"control_plane_id"`
}
