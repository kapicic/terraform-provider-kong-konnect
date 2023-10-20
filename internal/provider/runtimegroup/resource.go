package runtimegroup

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/kong-sdk/pkg/client"
	"github.com/kong-sdk/pkg/runtimegroups"
	"github.com/kong-sdk/pkg/shared"
	"github.com/kong/internal/provider/runtimegroup/models/config"
)

// ensure we implement the needed interfaces
var _ resource.Resource = &RuntimegroupResource{}
var _ resource.ResourceWithImportState = &RuntimegroupResource{}

// constructor
func NewRuntimegroupResource() resource.Resource {
	return &RuntimegroupResource{}
}

// client wrapper
type RuntimegroupResource struct {
	client *client.Client
}

type RuntimegroupResourceModel struct {
	Id          types.String  `tfsdk:"id"`
	Name        types.String  `tfsdk:"name"`
	Description types.String  `tfsdk:"description"`
	Config      config.Config `tfsdk:"config"`
	CreatedAt   types.String  `tfsdk:"created_at"`
	UpdatedAt   types.String  `tfsdk:"updated_at"`
	ClusterType types.String  `tfsdk:"cluster_type"`
	AuthType    types.String  `tfsdk:"auth_type"`
}

func (r *RuntimegroupResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_runtimegroup"
}

func (r *RuntimegroupResource) Schema(_ context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "The runtime group ID.",
				Optional:    true,
			},

			"name": schema.StringAttribute{
				Description: "The name of the runtime group.",
				Optional:    true,
			},

			"description": schema.StringAttribute{
				Description: "The description of the runtime group in Konnect.",
				Optional:    true,
			},

			"config": schema.SingleNestedAttribute{
				Optional:    true,
				Description: "CP configuration object for related access endpoints.",
				Attributes: map[string]schema.Attribute{
					"control_plane_endpoint": schema.StringAttribute{
						Description: "Control Plane Endpoint.",
						Optional:    true,
					},

					"telemetry_endpoint": schema.StringAttribute{
						Description: "Telemetry Endpoint.",
						Optional:    true,
					},
				},
			},

			"created_at": schema.StringAttribute{
				Description: "An ISO-8604 timestamp representation of runtime group creation date.",
				Optional:    true,
			},

			"updated_at": schema.StringAttribute{
				Description: "An ISO-8604 timestamp representation of runtime group update date.",
				Optional:    true,
			},

			"cluster_type": schema.StringAttribute{
				Description: "The ClusterType value of the cluster associated with the Runtime Group.",
				Optional:    true,
			},

			"auth_type": schema.StringAttribute{
				Description: "The auth type value of the cluster associated with the Runtime Group.",
				Optional:    true,
			},
		},
	}
}

func (r *RuntimegroupResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *RuntimegroupResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data RuntimegroupResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	Id := data.Id.ValueString()

	Runtimegroup, err := r.client.RuntimeGroups.GetRuntimeGroup(Id, shared.RequestOptions{})

	if err != nil {
		resp.Diagnostics.AddError(
			"Unexpected error calling RuntimeGroups.GetRuntimeGroup",
			err.Error(),
		)

		return
	}

	data.Id = types.StringValue(*Runtimegroup.Id)

	data.Name = types.StringValue(*Runtimegroup.Name)

	data.Description = types.StringValue(*Runtimegroup.Description)

	data.Config = config.Config{
		ControlPlaneEndpoint: types.StringValue(*Runtimegroup.Config.ControlPlaneEndpoint),

		TelemetryEndpoint: types.StringValue(*Runtimegroup.Config.TelemetryEndpoint),
	}

	data.CreatedAt = types.StringValue(*Runtimegroup.CreatedAt)

	data.UpdatedAt = types.StringValue(*Runtimegroup.UpdatedAt)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RuntimegroupResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data RuntimegroupResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// TODO: figure out struct name of createRequest
	createRequest := runtimegroups.CreateRuntimeGroupRequest{
		Name:        data.Name.ValueStringPointer(),
		Description: data.Description.ValueStringPointer(),
		ClusterType: pointer(runtimegroups.ClusterType(data.ClusterType.ValueString())),
		AuthType:    pointer(runtimegroups.AuthType(data.AuthType.ValueString())),
	}

	// make request
	Runtimegroup, err := r.client.RuntimeGroups.CreateRuntimeGroup(createRequest, shared.RequestOptions{})

	if err != nil {
		resp.Diagnostics.AddError(
			"Error Creating Runtimegroup",
			err.Error(),
		)

		return
	}

	// TODO: this can probably be a function using reflection
	data.Id = types.StringValue(*Runtimegroup.Id)

	data.Name = types.StringValue(*Runtimegroup.Name)

	data.Description = types.StringValue(*Runtimegroup.Description)

	data.Config = config.Config{
		ControlPlaneEndpoint: types.StringValue(*Runtimegroup.Config.ControlPlaneEndpoint),

		TelemetryEndpoint: types.StringValue(*Runtimegroup.Config.TelemetryEndpoint),
	}

	data.CreatedAt = types.StringValue(*Runtimegroup.CreatedAt)

	data.UpdatedAt = types.StringValue(*Runtimegroup.UpdatedAt)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RuntimegroupResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data = &RuntimegroupResourceModel{}

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	Id := data.Id.ValueString()

	err := r.client.RuntimeGroups.DeleteRuntimeGroup(Id, shared.RequestOptions{})

	if err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting Runtimegroup",
			err.Error(),
		)
	}
}

func (r *RuntimegroupResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data = &RuntimegroupResourceModel{}

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// TODO: add query params
	Id := data.Id.ValueString()

	updateRequest := runtimegroups.UpdateRuntimeGroupRequest{
		Name:        data.Name.ValueStringPointer(),
		Description: data.Description.ValueStringPointer(),
		AuthType:    pointer(runtimegroups.AuthType1(data.AuthType.ValueString())),
	}

	Runtimegroup, err := r.client.RuntimeGroups.UpdateRuntimeGroup(Id, updateRequest, shared.RequestOptions{})

	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating Runtimegroup",
			err.Error(),
		)

		return
	}

	// TODO: this can probably be a function using reflection
	data.Id = types.StringValue(*Runtimegroup.Id)

	data.Name = types.StringValue(*Runtimegroup.Name)

	data.Description = types.StringValue(*Runtimegroup.Description)

	data.Config = config.Config{
		ControlPlaneEndpoint: types.StringValue(*Runtimegroup.Config.ControlPlaneEndpoint),

		TelemetryEndpoint: types.StringValue(*Runtimegroup.Config.TelemetryEndpoint),
	}

	data.CreatedAt = types.StringValue(*Runtimegroup.CreatedAt)

	data.UpdatedAt = types.StringValue(*Runtimegroup.UpdatedAt)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RuntimegroupResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
