package apiproduct

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/kong-sdk/pkg/apiproducts"
	"github.com/kong-sdk/pkg/client"
	"github.com/kong-sdk/pkg/shared"
	"github.com/kong/internal/provider/apiproduct/models/api_product_labels"
	"github.com/kong/internal/utils"
)

// ensure we implement the needed interfaces
var _ resource.Resource = &ApiProductResource{}
var _ resource.ResourceWithImportState = &ApiProductResource{}

// constructor
func NewApiProductResource() resource.Resource {
	return &ApiProductResource{}
}

// client wrapper
type ApiProductResource struct {
	client *client.Client
}

type ApiProductResourceModel struct {
	Name        types.String                         `tfsdk:"name"`
	Description types.String                         `tfsdk:"description"`
	Labels      *api_product_labels.ApiProductLabels `tfsdk:"labels"`
	PortalIds   types.List                           `tfsdk:"portal_ids"`
	Id          types.String                         `tfsdk:"id"`
	CreatedAt   types.String                         `tfsdk:"created_at"`
	UpdatedAt   types.String                         `tfsdk:"updated_at"`
}

func (r *ApiProductResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_api_product"
}

func (r *ApiProductResource) Schema(_ context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				Description: "The name of the API product.",
				Required:    true,
			},

			"description": schema.StringAttribute{
				Description: "The description of the API product.",
				Optional:    true,
			},

			"labels": schema.SingleNestedAttribute{
				Description: "description: A maximum of 5 user-defined labels are allowed on this resource.Keys must not start with kong, konnect, insomnia, mesh, kic or _, which are reserved for Kong.Keys are case-sensitive.",
				Optional:    true,

				Attributes: map[string]schema.Attribute{
					"name": schema.StringAttribute{
						Description: "name",
						Optional:    true,
					},
				},
			},

			"portal_ids": schema.ListAttribute{
				Description: "The list of portal identifiers which this API product should be published to",
				Computed:    true,
				Optional:    true,

				ElementType: types.StringType,
			},

			"id": schema.StringAttribute{
				Description: "API product identifier",
				Computed:    true,
				Optional:    true,
			},

			"created_at": schema.StringAttribute{
				Description: "An ISO-8601 timestamp representation of entity creation date.",
				Computed:    true,
				Optional:    true,
			},

			"updated_at": schema.StringAttribute{
				Description: "An ISO-8601 timestamp representation of entity update date.",
				Computed:    true,
				Optional:    true,
			},
		},
	}
}

func (r *ApiProductResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *ApiProductResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ApiProductResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	Id := data.Id.ValueString()

	requestOptions := shared.RequestOptions{}

	apiProduct, err := r.client.ApiProducts.GetApiProduct(Id, requestOptions)

	if err != nil {
		resp.Diagnostics.AddError(
			"Unexpected error calling ApiProducts.GetApiProduct",
			err.Error(),
		)

		return
	}

	data.Id = utils.NullableString(apiProduct.Id)

	data.Name = utils.NullableString(apiProduct.Name)

	data.Description = utils.NullableString(apiProduct.Description)

	var PortalIdsDiags diag.Diagnostics

	data.PortalIds, PortalIdsDiags = types.ListValueFrom(ctx, types.StringType, apiProduct.PortalIds)
	if PortalIdsDiags.HasError() {
		resp.Diagnostics.Append(PortalIdsDiags...)
	}

	data.CreatedAt = utils.NullableString(apiProduct.CreatedAt)

	data.UpdatedAt = utils.NullableString(apiProduct.UpdatedAt)

	data.Labels = utils.NullableObject(apiProduct.Labels, api_product_labels.ApiProductLabels{
		Name: utils.NullableString(apiProduct.Labels.Name),
	})

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ApiProductResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ApiProductResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	requestOptions := shared.RequestOptions{}

	createRequest := apiproducts.CreateApiProductDto{
		Name:        data.Name.ValueStringPointer(),
		Description: data.Description.ValueStringPointer(),
		Labels: &apiproducts.ApiProductLabels{
			Name: data.Labels.Name.ValueStringPointer(),
		},
	}

	apiProduct, err := r.client.ApiProducts.CreateApiProduct(createRequest, requestOptions)

	if err != nil {
		resp.Diagnostics.AddError(
			"Error Creating ApiProduct",
			err.Error(),
		)

		return
	}

	data.Id = utils.NullableString(apiProduct.Id)

	data.Name = utils.NullableString(apiProduct.Name)

	data.Description = utils.NullableString(apiProduct.Description)

	var PortalIdsDiags diag.Diagnostics

	data.PortalIds, PortalIdsDiags = types.ListValueFrom(ctx, types.StringType, apiProduct.PortalIds)
	if PortalIdsDiags.HasError() {
		resp.Diagnostics.Append(PortalIdsDiags...)
	}

	data.CreatedAt = utils.NullableString(apiProduct.CreatedAt)

	data.UpdatedAt = utils.NullableString(apiProduct.UpdatedAt)

	data.Labels = utils.NullableObject(apiProduct.Labels, api_product_labels.ApiProductLabels{
		Name: utils.NullableString(apiProduct.Labels.Name),
	})

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ApiProductResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data = &ApiProductResourceModel{}

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	requestOptions := shared.RequestOptions{}

	Id := data.Id.ValueString()

	err := r.client.ApiProducts.DeleteApiProduct(Id, requestOptions)

	if err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting ApiProduct",
			err.Error(),
		)
	}
}

func (r *ApiProductResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	var data = &ApiProductResourceModel{}

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	requestOptions := shared.RequestOptions{}

	Id := data.Id.ValueString()

	updateRequest := apiproducts.UpdateApiProductDto{
		Name:        data.Name.ValueStringPointer(),
		Description: data.Description.ValueStringPointer(),
		Labels: &apiproducts.ApiProductLabels{
			Name: data.Labels.Name.ValueStringPointer(),
		},
		PortalIds: utils.FromListToPrimitiveSlice[string](ctx, data.PortalIds, types.StringType, &resp.Diagnostics),
	}

	apiProduct, err := r.client.ApiProducts.UpdateApiProduct(Id, updateRequest, requestOptions)

	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating ApiProduct",
			err.Error(),
		)

		return
	}

	data.Id = utils.NullableString(apiProduct.Id)

	data.Name = utils.NullableString(apiProduct.Name)

	data.Description = utils.NullableString(apiProduct.Description)

	var PortalIdsDiags diag.Diagnostics

	data.PortalIds, PortalIdsDiags = types.ListValueFrom(ctx, types.StringType, apiProduct.PortalIds)
	if PortalIdsDiags.HasError() {
		resp.Diagnostics.Append(PortalIdsDiags...)
	}

	data.CreatedAt = utils.NullableString(apiProduct.CreatedAt)

	data.UpdatedAt = utils.NullableString(apiProduct.UpdatedAt)

	data.Labels = utils.NullableObject(apiProduct.Labels, api_product_labels.ApiProductLabels{
		Name: utils.NullableString(apiProduct.Labels.Name),
	})

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ApiProductResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
