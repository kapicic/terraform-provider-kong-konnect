package client_certificate

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ClientCertificate struct {
	Id types.String `tfsdk:"id"`
}
