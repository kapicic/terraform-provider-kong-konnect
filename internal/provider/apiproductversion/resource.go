package apiproductversion

import (
    "context"
    "fmt"
    "github.com/hashicorp/terraform-plugin-framework/resource"
    "github.com/hashicorp/terraform-plugin-framework/resource/schema"
    "github.com/hashicorp/terraform-plugin-framework/types"
    "github.com/hashicorp/terraform-plugin-framework/path"
    "github.com/kong-sdk/pkg/client"
    "github.com/kong-sdk/pkg/shared"
        "github.com/kong-sdk/pkg/apiproductversions"
        "github.com/kong/internal/provider/apiproductversion/models/gateway_service_payload"
    "github.com/kong/internal/utils"
)

// ensure we implement the needed interfaces
var _ resource.Resource = &ApiProductVersionResource{}
var _ resource.ResourceWithImportState = &ApiProductVersionResource{}

// constructor
func NewApiProductVersionResource() resource.Resource {
    return &ApiProductVersionResource{}
}

// client wrapper
type ApiProductVersionResource struct {
    client *client.Client
}

type ApiProductVersionResourceModel struct {
    Id types.String `tfsdk:"id"`
    Name types.String `tfsdk:"name"`
    GatewayService *gateway_service_payload.GatewayServicePayload `tfsdk:"gateway_service"`
    PublishStatus types.String `tfsdk:"publish_status"`
    Deprecated types.Bool `tfsdk:"deprecated"`
    CreatedAt types.String `tfsdk:"created_at"`
    UpdatedAt types.String `tfsdk:"updated_at"`
    Notify types.Bool `tfsdk:"notify"`
    ApiProductId types.String `tfsdk:"api_product_id"`
}



func (r *ApiProductVersionResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
    resp.TypeName = req.ProviderTypeName + "_api_product_version"
}

func (r *ApiProductVersionResource) Schema(_ context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
    resp.Schema = schema.Schema{
        Attributes: map[string]schema.Attribute{
            "id": schema.StringAttribute{
        Description: "The API product version identifier.",
    Computed: true,
    Optional: true,

},

            "name": schema.StringAttribute{
        Description: "The version of the API product",
    Required: true,

},

            "gateway_service": schema.SingleNestedAttribute{
        Description: "gateway_service",
    Optional: true,

    Attributes: map[string]schema.Attribute{
            "id": schema.StringAttribute{
        Description: "The identifier of a gateway service associated with the version of the API product.",
    Required: true,

},

            "control_plane_id": schema.StringAttribute{
        Description: "The identifier of the control plane that the gateway service resides in",
    Required: true,

},

},

},

            "publish_status": schema.StringAttribute{
        Description: "The publish status of the API product version",
    Optional: true,

},

            "deprecated": schema.BoolAttribute{
        Description: "Indicates if this API product version is deprecated",
    Optional: true,

},

            "created_at": schema.StringAttribute{
        Description: "An ISO-8601 timestamp representation of entity creation date.",
    Computed: true,
    Optional: true,

},

            "updated_at": schema.StringAttribute{
        Description: "An ISO-8601 timestamp representation of entity update date.",
    Computed: true,
    Optional: true,

},

            "notify": schema.BoolAttribute{
        Description: "When set to `true`, and all the following conditions are true:- version of the API product deprecation has changed from `false` -> `true`- version of the API product is publishedthen consumers of the now deprecated verion of the API product will be notified.",
    Optional: true,

},

            "api_product_id": schema.StringAttribute{
        Description: "The API product identifier",
    Optional: true,

},

},

    }
}


func (r *ApiProductVersionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *ApiProductVersionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ApiProductVersionResourceModel
	utils.PopulateModelData(ctx, &data, resp.Diagnostics, req.State.Get)

	if resp.Diagnostics.HasError() {
		return
	}

			ApiProductId := data.ApiProductId.ValueString()
			Id := data.Id.ValueString()

	requestOptions := shared.RequestOptions{}


	apiProductVersion, err := r.client.ApiProductVersions.GetApiProductVersion(ApiProductId, Id, requestOptions)

	if err != nil {
		resp.Diagnostics.AddError(
			"Unexpected error calling ApiProductVersions.GetApiProductVersion",
			err.Error(),
		)

		return
	}

	        data.Id = utils.NullableString(apiProductVersion.GetId())

	        data.Name = utils.NullableString(apiProductVersion.GetName())

	        if apiProductVersion.GatewayService != nil {
        data.GatewayService = utils.NullableObject(apiProductVersion.GatewayService, gateway_service_payload.GatewayServicePayload{
                    Id: utils.NullableString(apiProductVersion.GetGatewayService().GetId()),

                    ControlPlaneId: utils.NullableString(apiProductVersion.GetGatewayService().GetControlPlaneId()),

        })
    }

	        // TODO: add nullable enum support
    data.PublishStatus = types.StringValue(string(*apiProductVersion.PublishStatus))

	        data.Deprecated = utils.NullableBool(apiProductVersion.GetDeprecated())

	        data.CreatedAt = utils.NullableString(apiProductVersion.GetCreatedAt())

	        data.UpdatedAt = utils.NullableString(apiProductVersion.GetUpdatedAt())


	if (resp.Diagnostics.HasError()) {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}


func (r *ApiProductVersionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
    var data ApiProductVersionResourceModel
    utils.PopulateModelData(ctx, &data, resp.Diagnostics, req.Plan.Get)

    if resp.Diagnostics.HasError() {
        return
    }

            ApiProductId := data.ApiProductId.ValueString()

    requestOptions := shared.RequestOptions{}


        createRequest := apiproductversions.CreateApiProductVersionDto{
            Name: data.Name.ValueStringPointer(),
            PublishStatus: utils.Pointer(apiproductversions.PublishStatus(data.PublishStatus.ValueString())),
            Deprecated: data.Deprecated.ValueBoolPointer(),
            
            GatewayService: utils.NullableTfStateObject(data.GatewayService, func(from *gateway_service_payload.GatewayServicePayload) apiproductversions.GatewayServicePayload {
                return apiproductversions.GatewayServicePayload{
                            Id: from.Id.ValueStringPointer(),
                            ControlPlaneId: from.ControlPlaneId.ValueStringPointer(),
                }
            }),
}
  
        apiProductVersion, err := r.client.ApiProductVersions.CreateApiProductVersion(ApiProductId, createRequest, requestOptions)

    if err != nil {
        resp.Diagnostics.AddError(
            "Error Creating ApiProductVersion",
            err.Error(),
        )

        return
    }

            data.Id = utils.NullableString(apiProductVersion.GetId())

            data.Name = utils.NullableString(apiProductVersion.GetName())

            if apiProductVersion.GatewayService != nil {
        data.GatewayService = utils.NullableObject(apiProductVersion.GatewayService, gateway_service_payload.GatewayServicePayload{
                    Id: utils.NullableString(apiProductVersion.GetGatewayService().GetId()),

                    ControlPlaneId: utils.NullableString(apiProductVersion.GetGatewayService().GetControlPlaneId()),

        })
    }

            // TODO: add nullable enum support
    data.PublishStatus = types.StringValue(string(*apiProductVersion.PublishStatus))

            data.Deprecated = utils.NullableBool(apiProductVersion.GetDeprecated())

            data.CreatedAt = utils.NullableString(apiProductVersion.GetCreatedAt())

            data.UpdatedAt = utils.NullableString(apiProductVersion.GetUpdatedAt())


    if (resp.Diagnostics.HasError()) {
        return
    }

    resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}


func (r *ApiProductVersionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
    var data = &ApiProductVersionResourceModel{}
    utils.PopulateModelData(ctx, &data, resp.Diagnostics, req.State.Get)

    if resp.Diagnostics.HasError() {
        return
    }

    requestOptions := shared.RequestOptions{}


            ApiProductId := data.ApiProductId.ValueString()
            Id := data.Id.ValueString()

    err := r.client.ApiProductVersions.DeleteApiProductVersion(ApiProductId, Id, requestOptions)

    if err != nil {
        resp.Diagnostics.AddError(
            "Error Deleting ApiProductVersion",
            err.Error(),
        )
    }
}


func (r *ApiProductVersionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
        var data = &ApiProductVersionResourceModel{}
        utils.PopulateModelData(ctx, &data, resp.Diagnostics, req.Plan.Get)

        if resp.Diagnostics.HasError() {
            return
        }

        requestOptions := shared.RequestOptions{}


                ApiProductId := data.ApiProductId.ValueString()
                Id := data.Id.ValueString()

        updateRequest := apiproductversions.UpdateApiProductVersionDto{
            Name: data.Name.ValueStringPointer(),
            PublishStatus: utils.Pointer(apiproductversions.PublishStatus2(data.PublishStatus.ValueString())),
            Deprecated: data.Deprecated.ValueBoolPointer(),
            Notify: data.Notify.ValueBoolPointer(),
            
            GatewayService: utils.NullableTfStateObject(data.GatewayService, func(from *gateway_service_payload.GatewayServicePayload) apiproductversions.GatewayServicePayload {
                return apiproductversions.GatewayServicePayload{
                            Id: from.Id.ValueStringPointer(),
                            ControlPlaneId: from.ControlPlaneId.ValueStringPointer(),
                }
            }),
}
  

        apiProductVersion, err := r.client.ApiProductVersions.UpdateApiProductVersion(ApiProductId, Id, updateRequest, requestOptions)

        if err != nil {
            resp.Diagnostics.AddError(
                "Error updating ApiProductVersion",
                err.Error(),
            )

            return
        }

                data.Id = utils.NullableString(apiProductVersion.GetId())

                data.Name = utils.NullableString(apiProductVersion.GetName())

                if apiProductVersion.GatewayService != nil {
        data.GatewayService = utils.NullableObject(apiProductVersion.GatewayService, gateway_service_payload.GatewayServicePayload{
                    Id: utils.NullableString(apiProductVersion.GetGatewayService().GetId()),

                    ControlPlaneId: utils.NullableString(apiProductVersion.GetGatewayService().GetControlPlaneId()),

        })
    }

                // TODO: add nullable enum support
    data.PublishStatus = types.StringValue(string(*apiProductVersion.PublishStatus))

                data.Deprecated = utils.NullableBool(apiProductVersion.GetDeprecated())

                data.CreatedAt = utils.NullableString(apiProductVersion.GetCreatedAt())

                data.UpdatedAt = utils.NullableString(apiProductVersion.GetUpdatedAt())


        if (resp.Diagnostics.HasError()) {
            return
        }

        resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}


func (r *ApiProductVersionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
    // Retrieve import ID and save to id attribute
    resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
