package config

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type Config struct {
	ControlPlaneEndpoint types.String `tfsdk:"control_plane_endpoint"`
	TelemetryEndpoint    types.String `tfsdk:"telemetry_endpoint"`
}
