package apiproduct

import (
    "context"
    "fmt"
    "github.com/hashicorp/terraform-plugin-framework/resource"
    "github.com/hashicorp/terraform-plugin-framework/resource/schema"
    "github.com/hashicorp/terraform-plugin-framework/types"
    "github.com/hashicorp/terraform-plugin-framework/path"
    "github.com/kong-sdk/pkg/client"
    "github.com/kong-sdk/pkg/shared"
        "github.com/kong-sdk/pkg/apiproducts"
        "github.com/kong/internal/provider/apiproduct/models/api_product_labels"
    "github.com/kong/internal/utils"
        "github.com/hashicorp/terraform-plugin-framework/diag"
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
    Id types.String `tfsdk:"id"`
    Name types.String `tfsdk:"name"`
    Description types.String `tfsdk:"description"`
    PortalIds types.List `tfsdk:"portal_ids"`
    CreatedAt types.String `tfsdk:"created_at"`
    UpdatedAt types.String `tfsdk:"updated_at"`
    Labels *api_product_labels.ApiProductLabels `tfsdk:"labels"`
}



func (r *ApiProductResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
    resp.TypeName = req.ProviderTypeName + "_api_product"
}

func (r *ApiProductResource) Schema(_ context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
    resp.Schema = schema.Schema{
        Attributes: map[string]schema.Attribute{
            "id": schema.StringAttribute{
        Description: "The API product ID.",
    Computed: true,
    Optional: true,

},

            "name": schema.StringAttribute{
        Description: "The name of the API product",
    Required: true,

},

            "description": schema.StringAttribute{
        Description: "The description of the API product",
    Optional: true,

},

                "portal_ids": schema.ListAttribute{
        Description: "The list of portal identifiers which this API product is published to",
    Computed: true,
    Optional: true,

    ElementType: types.StringType,
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

            "labels": schema.SingleNestedAttribute{
        Description: "description: A maximum of 5 user-defined labels are allowed on this resource.Keys must not start with kong, konnect, insomnia, mesh, kic or _, which are reserved for Kong.Keys are case-sensitive.",
    Optional: true,

    Attributes: map[string]schema.Attribute{
            "name": schema.StringAttribute{
        Description: "name",
    Optional: true,

},

},

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
	utils.PopulateModelData(ctx, &data, resp.Diagnostics, req.State.Get)

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

	        data.Id = utils.NullableString(apiProduct.GetId())

	        data.Name = utils.NullableString(apiProduct.GetName())

	        data.Description = utils.NullableString(apiProduct.GetDescription())

	        var PortalIdsDiags diag.Diagnostics

    data.PortalIds, PortalIdsDiags = types.ListValueFrom(ctx, types.StringType, apiProduct.PortalIds)
    if PortalIdsDiags.HasError() {
        resp.Diagnostics.Append(PortalIdsDiags...)
    }

	        data.CreatedAt = utils.NullableString(apiProduct.GetCreatedAt())

	        data.UpdatedAt = utils.NullableString(apiProduct.GetUpdatedAt())

	        if apiProduct.Labels != nil {
        data.Labels = utils.NullableObject(apiProduct.Labels, api_product_labels.ApiProductLabels{
                    Name: utils.NullableString(apiProduct.GetLabels().GetName()),

        })
    }


	if (resp.Diagnostics.HasError()) {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}


func (r *ApiProductResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
    var data ApiProductResourceModel
    utils.PopulateModelData(ctx, &data, resp.Diagnostics, req.Plan.Get)

    if resp.Diagnostics.HasError() {
        return
    }


    requestOptions := shared.RequestOptions{}


        createRequest := apiproducts.CreateApiProductDto{
            Name: data.Name.ValueStringPointer(),
            Description: data.Description.ValueStringPointer(),
            
            Labels: utils.NullableTfStateObject(data.Labels, func(from *api_product_labels.ApiProductLabels) apiproducts.ApiProductLabels {
                return apiproducts.ApiProductLabels{
                            Name: from.Name.ValueStringPointer(),
                }
            }),
}
  
        apiProduct, err := r.client.ApiProducts.CreateApiProduct(createRequest, requestOptions)

    if err != nil {
        resp.Diagnostics.AddError(
            "Error Creating ApiProduct",
            err.Error(),
        )

        return
    }

            data.Id = utils.NullableString(apiProduct.GetId())

            data.Name = utils.NullableString(apiProduct.GetName())

            data.Description = utils.NullableString(apiProduct.GetDescription())

            var PortalIdsDiags diag.Diagnostics

    data.PortalIds, PortalIdsDiags = types.ListValueFrom(ctx, types.StringType, apiProduct.PortalIds)
    if PortalIdsDiags.HasError() {
        resp.Diagnostics.Append(PortalIdsDiags...)
    }

            data.CreatedAt = utils.NullableString(apiProduct.GetCreatedAt())

            data.UpdatedAt = utils.NullableString(apiProduct.GetUpdatedAt())

            if apiProduct.Labels != nil {
        data.Labels = utils.NullableObject(apiProduct.Labels, api_product_labels.ApiProductLabels{
                    Name: utils.NullableString(apiProduct.GetLabels().GetName()),

        })
    }


    if (resp.Diagnostics.HasError()) {
        return
    }

    resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}


func (r *ApiProductResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
    var data = &ApiProductResourceModel{}
    utils.PopulateModelData(ctx, &data, resp.Diagnostics, req.State.Get)

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
        utils.PopulateModelData(ctx, &data, resp.Diagnostics, req.Plan.Get)

        if resp.Diagnostics.HasError() {
            return
        }

        requestOptions := shared.RequestOptions{}


                Id := data.Id.ValueString()

        updateRequest := apiproducts.UpdateApiProductDto{
            Name: data.Name.ValueStringPointer(),
            Description: data.Description.ValueStringPointer(),
            
            Labels: utils.NullableTfStateObject(data.Labels, func(from *api_product_labels.ApiProductLabels) apiproducts.ApiProductLabels {
                return apiproducts.ApiProductLabels{
                            Name: from.Name.ValueStringPointer(),
                }
            }),
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

                data.Id = utils.NullableString(apiProduct.GetId())

                data.Name = utils.NullableString(apiProduct.GetName())

                data.Description = utils.NullableString(apiProduct.GetDescription())

                var PortalIdsDiags diag.Diagnostics

    data.PortalIds, PortalIdsDiags = types.ListValueFrom(ctx, types.StringType, apiProduct.PortalIds)
    if PortalIdsDiags.HasError() {
        resp.Diagnostics.Append(PortalIdsDiags...)
    }

                data.CreatedAt = utils.NullableString(apiProduct.GetCreatedAt())

                data.UpdatedAt = utils.NullableString(apiProduct.GetUpdatedAt())

                if apiProduct.Labels != nil {
        data.Labels = utils.NullableObject(apiProduct.Labels, api_product_labels.ApiProductLabels{
                    Name: utils.NullableString(apiProduct.GetLabels().GetName()),

        })
    }


        if (resp.Diagnostics.HasError()) {
            return
        }

        resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}


func (r *ApiProductResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
    // Retrieve import ID and save to id attribute
    resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
