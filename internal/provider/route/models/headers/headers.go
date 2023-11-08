package headers

import (
        "github.com/hashicorp/terraform-plugin-framework/types"
)

type Headers struct {
    Key types.String `tfsdk:"key"`
}


