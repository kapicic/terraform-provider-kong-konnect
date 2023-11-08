package runtimegroup

import (
    "context"
    "fmt"
    "github.com/hashicorp/terraform-plugin-framework/resource"
    "github.com/hashicorp/terraform-plugin-framework/resource/schema"
    "github.com/hashicorp/terraform-plugin-framework/types"
    "github.com/hashicorp/terraform-plugin-framework/path"
    "github.com/kong-sdk/pkg/client"
    "github.com/kong-sdk/pkg/shared"
        "github.com/kong-sdk/pkg/runtimegroups"
        "github.com/kong/internal/provider/runtimegroup/models/labels"
        "github.com/kong/internal/provider/runtimegroup/models/config"
    "github.com/kong/internal/utils"
)

// ensure we implement the needed interfaces
var _ resource.Resource = &RunTimeGroupResource{}
var _ resource.ResourceWithImportState = &RunTimeGroupResource{}

// constructor
func NewRunTimeGroupResource() resource.Resource {
    return &RunTimeGroupResource{}
}

// client wrapper
type RunTimeGroupResource struct {
    client *client.Client
}

type RunTimeGroupResourceModel struct {
    Id types.String `tfsdk:"id"`
    Name types.String `tfsdk:"name"`
    Description types.String `tfsdk:"description"`
    Labels *labels.Labels `tfsdk:"labels"`
    Config *config.Config `tfsdk:"config"`
    CreatedAt types.String `tfsdk:"created_at"`
    UpdatedAt types.String `tfsdk:"updated_at"`
    ClusterType types.String `tfsdk:"cluster_type"`
    AuthType types.String `tfsdk:"auth_type"`
}



func (r *RunTimeGroupResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
    resp.TypeName = req.ProviderTypeName + "_run_time_group"
}

func (r *RunTimeGroupResource) Schema(_ context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
    resp.Schema = schema.Schema{
        Attributes: map[string]schema.Attribute{
            "id": schema.StringAttribute{
        Description: "The runtime group ID.",
    Computed: true,
    Optional: true,

},

            "name": schema.StringAttribute{
        Description: "The name of the runtime group.",
    Required: true,

},

            "description": schema.StringAttribute{
        Description: "The description of the runtime group in Konnect.",
    Optional: true,

},

            "labels": schema.SingleNestedAttribute{
        Description: "Labels to facilitate tagged search on runtime groups. Keys must be of length 1-63 characters, and cannot start with 'kong', 'konnect', 'mesh', 'kic', or '_'.",
    Optional: true,

    Attributes: map[string]schema.Attribute{
            "name": schema.StringAttribute{
        Description: "name",
    Optional: true,

},

},

},

            "config": schema.SingleNestedAttribute{
        Description: "CP configuration object for related access endpoints.",
    Computed: true,
    Optional: true,

    Attributes: map[string]schema.Attribute{
            "control_plane_endpoint": schema.StringAttribute{
        Description: "Control Plane Endpoint.",
    Optional: true,

},

            "telemetry_endpoint": schema.StringAttribute{
        Description: "Telemetry Endpoint.",
    Optional: true,

},

},

},

            "created_at": schema.StringAttribute{
        Description: "An ISO-8604 timestamp representation of runtime group creation date.",
    Computed: true,
    Optional: true,

},

            "updated_at": schema.StringAttribute{
        Description: "An ISO-8604 timestamp representation of runtime group update date.",
    Computed: true,
    Optional: true,

},

            "cluster_type": schema.StringAttribute{
        Description: "The ClusterType value of the cluster associated with the Runtime Group.",
    Optional: true,

},

            "auth_type": schema.StringAttribute{
        Description: "The auth type value of the cluster associated with the Runtime Group.",
    Optional: true,

},

},

    }
}


func (r *RunTimeGroupResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *RunTimeGroupResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data RunTimeGroupResourceModel
	utils.PopulateModelData(ctx, &data, resp.Diagnostics, req.State.Get)

	if resp.Diagnostics.HasError() {
		return
	}

			Id := data.Id.ValueString()

	requestOptions := shared.RequestOptions{}


	runTimeGroup, err := r.client.RuntimeGroups.GetRuntimeGroup(Id, requestOptions)

	if err != nil {
		resp.Diagnostics.AddError(
			"Unexpected error calling RuntimeGroups.GetRuntimeGroup",
			err.Error(),
		)

		return
	}

	        data.Id = utils.NullableString(runTimeGroup.GetId())

	        data.Name = utils.NullableString(runTimeGroup.GetName())

	        data.Description = utils.NullableString(runTimeGroup.GetDescription())

	        if runTimeGroup.Labels != nil {
        data.Labels = utils.NullableObject(runTimeGroup.Labels, labels.Labels{
                    Name: utils.NullableString(runTimeGroup.GetLabels().GetName()),

        })
    }

	        if runTimeGroup.Config != nil {
        data.Config = utils.NullableObject(runTimeGroup.Config, config.Config{
                    ControlPlaneEndpoint: utils.NullableString(runTimeGroup.GetConfig().GetControlPlaneEndpoint()),

                    TelemetryEndpoint: utils.NullableString(runTimeGroup.GetConfig().GetTelemetryEndpoint()),

        })
    }

	        data.CreatedAt = utils.NullableString(runTimeGroup.GetCreatedAt())

	        data.UpdatedAt = utils.NullableString(runTimeGroup.GetUpdatedAt())


	if (resp.Diagnostics.HasError()) {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}


func (r *RunTimeGroupResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
    var data RunTimeGroupResourceModel
    utils.PopulateModelData(ctx, &data, resp.Diagnostics, req.Plan.Get)

    if resp.Diagnostics.HasError() {
        return
    }


    requestOptions := shared.RequestOptions{}


        createRequest := runtimegroups.CreateRuntimeGroupRequest{
            Name: data.Name.ValueStringPointer(),
            Description: data.Description.ValueStringPointer(),
            ClusterType: utils.Pointer(runtimegroups.ClusterType(data.ClusterType.ValueString())),
            AuthType: utils.Pointer(runtimegroups.AuthType(data.AuthType.ValueString())),
            
            Labels: utils.NullableTfStateObject(data.Labels, func(from *labels.Labels) runtimegroups.Labels {
                return runtimegroups.Labels{
                            Name: from.Name.ValueStringPointer(),
                }
            }),
}
  
        runTimeGroup, err := r.client.RuntimeGroups.CreateRuntimeGroup(createRequest, requestOptions)

    if err != nil {
        resp.Diagnostics.AddError(
            "Error Creating RunTimeGroup",
            err.Error(),
        )

        return
    }

            data.Id = utils.NullableString(runTimeGroup.GetId())

            data.Name = utils.NullableString(runTimeGroup.GetName())

            data.Description = utils.NullableString(runTimeGroup.GetDescription())

            if runTimeGroup.Labels != nil {
        data.Labels = utils.NullableObject(runTimeGroup.Labels, labels.Labels{
                    Name: utils.NullableString(runTimeGroup.GetLabels().GetName()),

        })
    }

            if runTimeGroup.Config != nil {
        data.Config = utils.NullableObject(runTimeGroup.Config, config.Config{
                    ControlPlaneEndpoint: utils.NullableString(runTimeGroup.GetConfig().GetControlPlaneEndpoint()),

                    TelemetryEndpoint: utils.NullableString(runTimeGroup.GetConfig().GetTelemetryEndpoint()),

        })
    }

            data.CreatedAt = utils.NullableString(runTimeGroup.GetCreatedAt())

            data.UpdatedAt = utils.NullableString(runTimeGroup.GetUpdatedAt())


    if (resp.Diagnostics.HasError()) {
        return
    }

    resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}


func (r *RunTimeGroupResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
    var data = &RunTimeGroupResourceModel{}
    utils.PopulateModelData(ctx, &data, resp.Diagnostics, req.State.Get)

    if resp.Diagnostics.HasError() {
        return
    }

    requestOptions := shared.RequestOptions{}


            Id := data.Id.ValueString()

    err := r.client.RuntimeGroups.DeleteRuntimeGroup(Id, requestOptions)

    if err != nil {
        resp.Diagnostics.AddError(
            "Error Deleting RunTimeGroup",
            err.Error(),
        )
    }
}


func (r *RunTimeGroupResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
        var data = &RunTimeGroupResourceModel{}
        utils.PopulateModelData(ctx, &data, resp.Diagnostics, req.Plan.Get)

        if resp.Diagnostics.HasError() {
            return
        }

        requestOptions := shared.RequestOptions{}


                Id := data.Id.ValueString()

        updateRequest := runtimegroups.UpdateRuntimeGroupRequest{
            Name: data.Name.ValueStringPointer(),
            Description: data.Description.ValueStringPointer(),
            AuthType: utils.Pointer(runtimegroups.AuthType1(data.AuthType.ValueString())),
            
            Labels: utils.NullableTfStateObject(data.Labels, func(from *labels.Labels) runtimegroups.Labels {
                return runtimegroups.Labels{
                            Name: from.Name.ValueStringPointer(),
                }
            }),
}
  

        runTimeGroup, err := r.client.RuntimeGroups.UpdateRuntimeGroup(Id, updateRequest, requestOptions)

        if err != nil {
            resp.Diagnostics.AddError(
                "Error updating RunTimeGroup",
                err.Error(),
            )

            return
        }

                data.Id = utils.NullableString(runTimeGroup.GetId())

                data.Name = utils.NullableString(runTimeGroup.GetName())

                data.Description = utils.NullableString(runTimeGroup.GetDescription())

                if runTimeGroup.Labels != nil {
        data.Labels = utils.NullableObject(runTimeGroup.Labels, labels.Labels{
                    Name: utils.NullableString(runTimeGroup.GetLabels().GetName()),

        })
    }

                if runTimeGroup.Config != nil {
        data.Config = utils.NullableObject(runTimeGroup.Config, config.Config{
                    ControlPlaneEndpoint: utils.NullableString(runTimeGroup.GetConfig().GetControlPlaneEndpoint()),

                    TelemetryEndpoint: utils.NullableString(runTimeGroup.GetConfig().GetTelemetryEndpoint()),

        })
    }

                data.CreatedAt = utils.NullableString(runTimeGroup.GetCreatedAt())

                data.UpdatedAt = utils.NullableString(runTimeGroup.GetUpdatedAt())


        if (resp.Diagnostics.HasError()) {
            return
        }

        resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}


func (r *RunTimeGroupResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
    // Retrieve import ID and save to id attribute
    resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
