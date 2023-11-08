package labels

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type Labels struct {
	Name types.String `tfsdk:"name"`
}
