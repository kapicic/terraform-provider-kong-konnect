package api_product_labels

import (
        "github.com/hashicorp/terraform-plugin-framework/types"
)

type ApiProductLabels struct {
    Name types.String `tfsdk:"name"`
}


