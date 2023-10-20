package apiproduct

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/liblab-sdk/pkg/apiproducts"
	"github.com/liblab-sdk/pkg/client"
	"github.com/liblab-sdk/pkg/shared"
)

// ensure we implement the needed interfaces
var _ resource.Resource = &ApiproductResource{}
var _ resource.ResourceWithImportState = &ApiproductResource{}

// constructor
func NewApiproductResource() resource.Resource {
	return &ApiproductResource{}
}

// client wrapper
type ApiproductResource struct {
	client *client.Client
}

type ApiproductResourceModel struct {
	Id          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
	PortalIds   types.List   `tfsdk:"portal_ids"`
	CreatedAt   types.String `tfsdk:"created_at"`
	UpdatedAt   types.String `tfsdk:"updated_at"`
}

func (r *ApiproductResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_apiproduct"
}

func (r *ApiproductResource) Schema(_ context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "The API product ID.",
				Required:    true,
			},

			"name": schema.StringAttribute{
				Description: "The name of the API product",
				Required:    true,
			},

			"description": schema.StringAttribute{
				Description: "The description of the API product",
				Required:    true,
			},

			"portal_ids": schema.ListAttribute{
				Description: "The list of portal identifiers which this API product is published to",
				Required:    true,
				ElementType: types.StringType,
			},

			"created_at": schema.StringAttribute{
				Description: "An ISO-8601 timestamp representation of entity creation date.",
				Required:    true,
			},

			"updated_at": schema.StringAttribute{
				Description: "An ISO-8601 timestamp representation of entity update date.",
				Required:    true,
			},
		},
	}
}

func (r *ApiproductResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	apiClient, ok := req.ProviderData.(*client.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *client.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = apiClient
}

func (r *ApiproductResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ApiproductResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	Id := data.Id.ValueString()

	Apiproduct, err := r.client.ApiProducts.GetApiProduct(Id, shared.RequestOptions{})

	if err != nil {
		resp.Diagnostics.AddError(
			"Unexpected error calling ApiProducts.GetApiProduct",
			err.Error(),
		)

		return
	}

	data.Id = types.StringValue(*Apiproduct.Id)

	data.Name = types.StringValue(*Apiproduct.Name)

	data.Description = types.StringValue(*Apiproduct.Description)

	var PortalIdsDiags diag.Diagnostics

	data.PortalIds, PortalIdsDiags = types.ListValueFrom(ctx, types.StringType, Apiproduct.PortalIds)

	if PortalIdsDiags.HasError() {
		resp.Diagnostics.Append(PortalIdsDiags...)

		return
	}

	data.CreatedAt = types.StringValue(*Apiproduct.CreatedAt)

	data.UpdatedAt = types.StringValue(*Apiproduct.UpdatedAt)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ApiproductResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ApiproductResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// TODO: figure out struct name of createRequest
	createRequest := apiproducts.CreateApiProductDto{
		Name:        data.Name.ValueStringPointer(),
		Description: data.Description.ValueStringPointer(),
	}

	// make request
	Apiproduct, err := r.client.ApiProducts.CreateApiProduct(createRequest, shared.RequestOptions{})

	if err != nil {
		resp.Diagnostics.AddError(
			"Error Creating Apiproduct",
			err.Error(),
		)

		return
	}

	// TODO: this can probably be a function using reflection
	data.Id = types.StringValue(*Apiproduct.Id)

	data.Name = types.StringValue(*Apiproduct.Name)

	data.Description = types.StringValue(*Apiproduct.Description)

	var PortalIdsDiags diag.Diagnostics

	data.PortalIds, PortalIdsDiags = types.ListValueFrom(ctx, types.StringType, Apiproduct.PortalIds)

	if PortalIdsDiags.HasError() {
		resp.Diagnostics.Append(PortalIdsDiags...)

		return
	}

	data.CreatedAt = types.StringValue(*Apiproduct.CreatedAt)

	data.UpdatedAt = types.StringValue(*Apiproduct.UpdatedAt)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ApiproductResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data = &ApiproductResourceModel{}

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	Id := data.Id.ValueString()

	err := r.client.ApiProducts.DeleteApiProduct(Id, shared.RequestOptions{})

	if err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting Apiproduct",
			err.Error(),
		)
	}
}

func (r *ApiproductResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data = &ApiproductResourceModel{}

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// TODO: add query params
	Id := data.Id.ValueString()

	updateRequest := apiproducts.UpdateApiProductDto{
		Name:        data.Name.ValueStringPointer(),
		Description: data.Description.ValueStringPointer(),
	}

	Apiproduct, err := r.client.ApiProducts.UpdateApiProduct(Id, updateRequest, shared.RequestOptions{})

	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating Apiproduct",
			err.Error(),
		)

		return
	}

	// TODO: this can probably be a function using reflection
	data.Id = types.StringValue(*Apiproduct.Id)

	data.Name = types.StringValue(*Apiproduct.Name)

	data.Description = types.StringValue(*Apiproduct.Description)

	var PortalIdsDiags diag.Diagnostics

	data.PortalIds, PortalIdsDiags = types.ListValueFrom(ctx, types.StringType, Apiproduct.PortalIds)

	if PortalIdsDiags.HasError() {
		resp.Diagnostics.Append(PortalIdsDiags...)

		return
	}

	data.CreatedAt = types.StringValue(*Apiproduct.CreatedAt)

	data.UpdatedAt = types.StringValue(*Apiproduct.UpdatedAt)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ApiproductResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func Map[T, R any](from *[]T, f func(T) R) []R {
	to := make([]R, len(*from))
	for i, v := range *from {
		to[i] = f(v)
	}
	return to
}

func pointer[T any](v T) *T {
	return &v
}
