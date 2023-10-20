package apiproductversion

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/kong-sdk/pkg/apiproductversions"
	"github.com/kong-sdk/pkg/client"
	"github.com/kong-sdk/pkg/shared"
	"github.com/kong/internal/provider/apiproductversion/models/gateway_service_payload"
)

// ensure we implement the needed interfaces
var _ resource.Resource = &ApiproductversionResource{}
var _ resource.ResourceWithImportState = &ApiproductversionResource{}

// constructor
func NewApiproductversionResource() resource.Resource {
	return &ApiproductversionResource{}
}

// client wrapper
type ApiproductversionResource struct {
	client *client.Client
}

type ApiproductversionResourceModel struct {
	Id             types.String                                  `tfsdk:"id"`
	Name           types.String                                  `tfsdk:"name"`
	GatewayService gateway_service_payload.GatewayServicePayload `tfsdk:"gateway_service"`
	PublishStatus  types.String                                  `tfsdk:"publish_status"`
	Deprecated     types.Bool                                    `tfsdk:"deprecated"`
	CreatedAt      types.String                                  `tfsdk:"created_at"`
	UpdatedAt      types.String                                  `tfsdk:"updated_at"`
	Notify         types.Bool                                    `tfsdk:"notify"`
	ApiProductId   types.String                                  `tfsdk:"api_product_id"`
}

func (r *ApiproductversionResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_apiproductversion"
}

func (r *ApiproductversionResource) Schema(_ context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "The API product version identifier.",
				Optional:    true,
			},

			"name": schema.StringAttribute{
				Description: "The version of the API product",
				Required:    true,
			},

			"gateway_service": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Description: "The identifier of a gateway service associated with the version of the API product.",
						Required:    true,
					},

					"control_plane_id": schema.StringAttribute{
						Description: "The identifier of the control plane that the gateway service resides in",
						Required:    true,
					},
				},
			},

			"publish_status": schema.StringAttribute{
				Description: "The publish status of the API product version",
				Required:    true,
			},

			"deprecated": schema.BoolAttribute{
				Description: "Indicates if this API product version is deprecated",
				Optional:    true,
			},

			"created_at": schema.StringAttribute{
				Description: "An ISO-8601 timestamp representation of entity creation date.",
				Optional:    true,
			},

			"updated_at": schema.StringAttribute{
				Description: "An ISO-8601 timestamp representation of entity update date.",
				Optional:    true,
			},

			"notify": schema.BoolAttribute{
				Description: "When set to `true`, and all the following conditions are true:- version of the API product deprecation has changed from `false` -> `true`- version of the API product is publishedthen consumers of the now deprecated verion of the API product will be notified.",
				Optional:    true,
			},

			"api_product_id": schema.StringAttribute{
				Description: "The API product identifier",
				Required:    true,
			},
		},
	}
}

func (r *ApiproductversionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *ApiproductversionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ApiproductversionResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	ApiProductId := data.ApiProductId.ValueString()
	Id := data.Id.ValueString()

	Apiproductversion, err := r.client.ApiProductVersions.GetApiProductVersion(ApiProductId, Id, shared.RequestOptions{})

	if err != nil {
		resp.Diagnostics.AddError(
			"Unexpected error calling ApiProductVersions.GetApiProductVersion",
			err.Error(),
		)

		return
	}

	data.Id = types.StringValue(*Apiproductversion.Id)

	data.Name = types.StringValue(*Apiproductversion.Name)

	data.GatewayService = gateway_service_payload.GatewayServicePayload{
		Id: types.StringValue(*Apiproductversion.GatewayService.Id),

		ControlPlaneId: types.StringValue(*Apiproductversion.GatewayService.ControlPlaneId),
	}

	data.PublishStatus = types.StringValue(string(*Apiproductversion.PublishStatus))

	data.Deprecated = types.BoolValue(*Apiproductversion.Deprecated)

	data.CreatedAt = types.StringValue(*Apiproductversion.CreatedAt)

	data.UpdatedAt = types.StringValue(*Apiproductversion.UpdatedAt)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ApiproductversionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ApiproductversionResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	ApiProductId := data.ApiProductId.ValueString()

	// TODO: figure out struct name of createRequest
	createRequest := apiproductversions.CreateApiProductVersionDto{
		Name:          data.Name.ValueStringPointer(),
		PublishStatus: pointer(apiproductversions.PublishStatus(data.PublishStatus.ValueString())),
		Deprecated:    data.Deprecated.ValueBoolPointer(),
		GatewayService: &apiproductversions.GatewayServicePayload{
			Id: data.GatewayService.Id.ValueStringPointer(),

			ControlPlaneId: data.GatewayService.ControlPlaneId.ValueStringPointer(),
		},
	}

	// make request
	Apiproductversion, err := r.client.ApiProductVersions.CreateApiProductVersion(ApiProductId, createRequest, shared.RequestOptions{})

	if err != nil {
		resp.Diagnostics.AddError(
			"Error Creating Apiproductversion",
			err.Error(),
		)

		return
	}

	// TODO: this can probably be a function using reflection
	data.Id = types.StringValue(*Apiproductversion.Id)

	data.Name = types.StringValue(*Apiproductversion.Name)

	data.GatewayService = gateway_service_payload.GatewayServicePayload{
		Id: types.StringValue(*Apiproductversion.GatewayService.Id),

		ControlPlaneId: types.StringValue(*Apiproductversion.GatewayService.ControlPlaneId),
	}

	data.PublishStatus = types.StringValue(string(*Apiproductversion.PublishStatus))

	data.Deprecated = types.BoolValue(*Apiproductversion.Deprecated)

	data.CreatedAt = types.StringValue(*Apiproductversion.CreatedAt)

	data.UpdatedAt = types.StringValue(*Apiproductversion.UpdatedAt)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ApiproductversionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data = &ApiproductversionResourceModel{}

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	ApiProductId := data.ApiProductId.ValueString()
	Id := data.Id.ValueString()

	err := r.client.ApiProductVersions.DeleteApiProductVersion(ApiProductId, Id, shared.RequestOptions{})

	if err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting Apiproductversion",
			err.Error(),
		)
	}
}

func (r *ApiproductversionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data = &ApiproductversionResourceModel{}

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// TODO: add query params
	ApiProductId := data.ApiProductId.ValueString()
	Id := data.Id.ValueString()

	updateRequest := apiproductversions.UpdateApiProductVersionDto{
		Name:          data.Name.ValueStringPointer(),
		PublishStatus: pointer(apiproductversions.PublishStatus2(data.PublishStatus.ValueString())),
		Deprecated:    data.Deprecated.ValueBoolPointer(),
		Notify:        data.Notify.ValueBoolPointer(),
		GatewayService: &apiproductversions.GatewayServicePayload{
			Id: data.GatewayService.Id.ValueStringPointer(),

			ControlPlaneId: data.GatewayService.ControlPlaneId.ValueStringPointer(),
		},
	}

	Apiproductversion, err := r.client.ApiProductVersions.UpdateApiProductVersion(ApiProductId, Id, updateRequest, shared.RequestOptions{})

	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating Apiproductversion",
			err.Error(),
		)

		return
	}

	// TODO: this can probably be a function using reflection
	data.Id = types.StringValue(*Apiproductversion.Id)

	data.Name = types.StringValue(*Apiproductversion.Name)

	data.GatewayService = gateway_service_payload.GatewayServicePayload{
		Id: types.StringValue(*Apiproductversion.GatewayService.Id),

		ControlPlaneId: types.StringValue(*Apiproductversion.GatewayService.ControlPlaneId),
	}

	data.PublishStatus = types.StringValue(string(*Apiproductversion.PublishStatus))

	data.Deprecated = types.BoolValue(*Apiproductversion.Deprecated)

	data.CreatedAt = types.StringValue(*Apiproductversion.CreatedAt)

	data.UpdatedAt = types.StringValue(*Apiproductversion.UpdatedAt)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ApiproductversionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
