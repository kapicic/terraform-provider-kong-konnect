package service

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type Service struct {
	Id types.String `tfsdk:"id"`
}
