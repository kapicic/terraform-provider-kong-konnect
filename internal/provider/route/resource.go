package route

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/kong-sdk/pkg/client"
	"github.com/kong-sdk/pkg/routes"
	"github.com/kong-sdk/pkg/shared"
	"github.com/kong/internal/provider/route/models/service"
)

// ensure we implement the needed interfaces
var _ resource.Resource = &RouteResource{}
var _ resource.ResourceWithImportState = &RouteResource{}

// constructor
func NewRouteResource() resource.Resource {
	return &RouteResource{}
}

// client wrapper
type RouteResource struct {
	client *client.Client
}

type RouteResourceModel struct {
	CreatedAt               types.Int64     `tfsdk:"created_at"`
	Hosts                   types.List      `tfsdk:"hosts"`
	HttpsRedirectStatusCode types.Int64     `tfsdk:"https_redirect_status_code"`
	Id                      types.String    `tfsdk:"id"`
	Methods                 types.List      `tfsdk:"methods"`
	Name                    types.String    `tfsdk:"name"`
	PathHandling            types.String    `tfsdk:"path_handling"`
	Paths                   types.List      `tfsdk:"paths"`
	PreserveHost            types.Bool      `tfsdk:"preserve_host"`
	Protocols               types.List      `tfsdk:"protocols"`
	RegexPriority           types.Int64     `tfsdk:"regex_priority"`
	RequestBuffering        types.Bool      `tfsdk:"request_buffering"`
	ResponseBuffering       types.Bool      `tfsdk:"response_buffering"`
	Service                 service.Service `tfsdk:"service"`
	Snis                    types.List      `tfsdk:"snis"`
	StripPath               types.Bool      `tfsdk:"strip_path"`
	Tags                    types.List      `tfsdk:"tags"`
	UpdatedAt               types.Int64     `tfsdk:"updated_at"`
	RuntimeGroupId          types.String    `tfsdk:"runtime_group_id"`
	RouteId                 types.String    `tfsdk:"route_id"`
}

func (r *RouteResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_route"
}

func (r *RouteResource) Schema(_ context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Route entities define rules to match client requests. Every request matching a given route will be proxied to its associated service.",
		Attributes: map[string]schema.Attribute{
			"created_at": schema.Int64Attribute{
				Description: "Unix epoch when the resource was created.",
				Optional:    true,
			},

			"hosts": schema.ListAttribute{
				Description: "A list of domain names that match this route. Note that the hosts value is case sensitive.",
				ElementType: types.StringType,
				Optional:    true,
			},

			"https_redirect_status_code": schema.Int64Attribute{
				Description: "The status code Kong responds with when all properties of a route match except the protocol i.e. if the protocol of the request is `HTTP` instead of `HTTPS`. `Location` header is injected by Kong if the field is set to 301, 302, 307 or 308. Note: This config applies only if the route is configured to only accept the `https` protocol.",
				Optional:    true,
			},

			"id": schema.StringAttribute{
				Optional: true,
			},

			"methods": schema.ListAttribute{
				Description: "A list of HTTP methods that match this route.",
				ElementType: types.StringType,
				Optional:    true,
			},

			"name": schema.StringAttribute{
				Description: "The name of the route. Route names must be unique, and they are case sensitive. For example, there can be two different routes named test and Test.",
				Optional:    true,
			},

			"path_handling": schema.StringAttribute{
				Description: "Controls how the service path, route path and requested path are combined when sending a request to the upstream. See above for a detailed description of each behavior.",
				Optional:    true,
			},

			"paths": schema.ListAttribute{
				Description: "A list of paths that match this route.",
				ElementType: types.StringType,
				Optional:    true,
			},

			"preserve_host": schema.BoolAttribute{
				Description: "When matching a route via one of the `hosts` domain names, use the request `Host` header in the upstream request headers. If set to `false`, the upstream `Host` header will be that of the service's `host`.",
				Optional:    true,
			},

			"protocols": schema.ListAttribute{
				Description: "An array of the protocols this route should allow. See the [route Object](#route-object) section for a list of accepted protocols. When set to only `https`, HTTP requests are answered with an upgrade error. When set to only `http`, HTTPS requests are answered with an error.",
				ElementType: types.StringType,
				Optional:    true,
			},

			"regex_priority": schema.Int64Attribute{
				Description: "A number used to choose which route resolves a given request when several routes match it using regexes simultaneously. When two routes match the path and have the same `regex_priority`, the older one (lowest `created_at`) is used. Note that the priority for non-regex routes is different (longer non-regex routes are matched before shorter ones).",
				Optional:    true,
			},

			"request_buffering": schema.BoolAttribute{
				Description: "Whether to enable request body buffering or not. With HTTP 1.1, it may make sense to turn this off on services that receive data with chunked transfer encoding.",
				Optional:    true,
			},

			"response_buffering": schema.BoolAttribute{
				Description: "Whether to enable response body buffering or not. With HTTP 1.1, it may make sense to turn this off on services that send data with chunked transfer encoding.",
				Optional:    true,
			},

			"service": schema.SingleNestedAttribute{
				Description: "The service this route is associated to. This is where the route proxies traffic to.",
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Optional: true,
					},
				},
				Optional: true,
			},

			"snis": schema.ListAttribute{
				Description: "A list of SNIs that match this route when using stream routing.",
				Optional:    true,
				ElementType: types.StringType,
			},

			"strip_path": schema.BoolAttribute{
				Optional:    true,
				Description: "When matching a route via one of the `paths`, strip the matching prefix from the upstream request URL.",
			},

			"tags": schema.ListAttribute{
				Optional:    true,
				Description: "An optional set of strings associated with the route for grouping and filtering.",
				ElementType: types.StringType,
			},

			"updated_at": schema.Int64Attribute{
				Description: "Unix epoch when the resource was last updated.",
				Optional:    true,
			},

			"runtime_group_id": schema.StringAttribute{
				Description: "The ID of your runtime group. This variable is available in the Konnect manager",
				Required:    true,
			},

			"route_id": schema.StringAttribute{
				Description: "The unique identifier or the name of the route to retrieve.",
				Required:    true,
			},
		},
	}
}

func (r *RouteResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *RouteResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data RouteResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	RuntimeGroupId := data.RuntimeGroupId.ValueString()
	RouteId := data.RouteId.ValueString()

	Route, err := r.client.Routes.GetRoute(RuntimeGroupId, RouteId, shared.RequestOptions{})

	if err != nil {
		resp.Diagnostics.AddError(
			"Unexpected error calling Routes.GetRoute",
			err.Error(),
		)

		return
	}

	data.CreatedAt = types.Int64Value(*Route.CreatedAt)

	var HostsDiags diag.Diagnostics

	data.Hosts, HostsDiags = types.ListValueFrom(ctx, types.StringType, Route.Hosts)

	if HostsDiags.HasError() {
		resp.Diagnostics.Append(HostsDiags...)

		return
	}

	data.HttpsRedirectStatusCode = types.Int64Value(*Route.HttpsRedirectStatusCode)

	data.Id = types.StringValue(*Route.Id)

	var MethodsDiags diag.Diagnostics

	data.Methods, MethodsDiags = types.ListValueFrom(ctx, types.StringType, Route.Methods)

	if MethodsDiags.HasError() {
		resp.Diagnostics.Append(MethodsDiags...)

		return
	}

	data.Name = types.StringValue(*Route.Name)

	data.PathHandling = types.StringValue(*Route.PathHandling)

	var PathsDiags diag.Diagnostics

	data.Paths, PathsDiags = types.ListValueFrom(ctx, types.StringType, Route.Paths)

	if PathsDiags.HasError() {
		resp.Diagnostics.Append(PathsDiags...)

		return
	}

	data.PreserveHost = types.BoolValue(*Route.PreserveHost)

	var ProtocolsDiags diag.Diagnostics

	data.Protocols, ProtocolsDiags = types.ListValueFrom(ctx, types.StringType, Route.Protocols)

	if ProtocolsDiags.HasError() {
		resp.Diagnostics.Append(ProtocolsDiags...)

		return
	}

	data.RegexPriority = types.Int64Value(*Route.RegexPriority)

	data.RequestBuffering = types.BoolValue(*Route.RequestBuffering)

	data.ResponseBuffering = types.BoolValue(*Route.ResponseBuffering)

	data.Service = service.Service{
		Id: types.StringValue(*Route.Service.Id),
	}

	var SnisDiags diag.Diagnostics

	data.Snis, SnisDiags = types.ListValueFrom(ctx, types.StringType, Route.Snis)

	if SnisDiags.HasError() {
		resp.Diagnostics.Append(SnisDiags...)

		return
	}

	data.StripPath = types.BoolValue(*Route.StripPath)

	var TagsDiags diag.Diagnostics

	data.Tags, TagsDiags = types.ListValueFrom(ctx, types.StringType, Route.Tags)

	if TagsDiags.HasError() {
		resp.Diagnostics.Append(TagsDiags...)

		return
	}

	data.UpdatedAt = types.Int64Value(*Route.UpdatedAt)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RouteResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data RouteResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	RuntimeGroupId := data.RuntimeGroupId.ValueString()

	// TODO: figure out struct name of createRequest
	createRequest := routes.Route{
		CreatedAt:               data.CreatedAt.ValueInt64Pointer(),
		HttpsRedirectStatusCode: data.HttpsRedirectStatusCode.ValueInt64Pointer(),
		Id:                      data.Id.ValueStringPointer(),
		Name:                    data.Name.ValueStringPointer(),
		PathHandling:            data.PathHandling.ValueStringPointer(),
		PreserveHost:            data.PreserveHost.ValueBoolPointer(),
		RegexPriority:           data.RegexPriority.ValueInt64Pointer(),
		RequestBuffering:        data.RequestBuffering.ValueBoolPointer(),
		ResponseBuffering:       data.ResponseBuffering.ValueBoolPointer(),
		Service: &routes.Service{
			Id: data.Service.Id.ValueStringPointer(),
		},
		StripPath: data.StripPath.ValueBoolPointer(),
		UpdatedAt: data.UpdatedAt.ValueInt64Pointer(),
	}

	// make request
	Route, err := r.client.Routes.CreateRoute(RuntimeGroupId, createRequest, shared.RequestOptions{})

	if err != nil {
		resp.Diagnostics.AddError(
			"Error Creating Route",
			err.Error(),
		)

		return
	}

	// TODO: this can probably be a function using reflection
	data.CreatedAt = types.Int64Value(*Route.CreatedAt)

	var HostsDiags diag.Diagnostics

	data.Hosts, HostsDiags = types.ListValueFrom(ctx, types.StringType, Route.Hosts)

	if HostsDiags.HasError() {
		resp.Diagnostics.Append(HostsDiags...)

		return
	}

	data.HttpsRedirectStatusCode = types.Int64Value(*Route.HttpsRedirectStatusCode)

	data.Id = types.StringValue(*Route.Id)

	var MethodsDiags diag.Diagnostics

	data.Methods, MethodsDiags = types.ListValueFrom(ctx, types.StringType, Route.Methods)

	if MethodsDiags.HasError() {
		resp.Diagnostics.Append(MethodsDiags...)

		return
	}

	data.Name = types.StringValue(*Route.Name)

	data.PathHandling = types.StringValue(*Route.PathHandling)

	var PathsDiags diag.Diagnostics

	data.Paths, PathsDiags = types.ListValueFrom(ctx, types.StringType, Route.Paths)

	if PathsDiags.HasError() {
		resp.Diagnostics.Append(PathsDiags...)

		return
	}

	data.PreserveHost = types.BoolValue(*Route.PreserveHost)

	var ProtocolsDiags diag.Diagnostics

	data.Protocols, ProtocolsDiags = types.ListValueFrom(ctx, types.StringType, Route.Protocols)

	if ProtocolsDiags.HasError() {
		resp.Diagnostics.Append(ProtocolsDiags...)

		return
	}

	data.RegexPriority = types.Int64Value(*Route.RegexPriority)

	data.RequestBuffering = types.BoolValue(*Route.RequestBuffering)

	data.ResponseBuffering = types.BoolValue(*Route.ResponseBuffering)

	data.Service = service.Service{
		Id: types.StringValue(*Route.Service.Id),
	}

	var SnisDiags diag.Diagnostics

	data.Snis, SnisDiags = types.ListValueFrom(ctx, types.StringType, Route.Snis)

	if SnisDiags.HasError() {
		resp.Diagnostics.Append(SnisDiags...)

		return
	}

	data.StripPath = types.BoolValue(*Route.StripPath)

	var TagsDiags diag.Diagnostics

	data.Tags, TagsDiags = types.ListValueFrom(ctx, types.StringType, Route.Tags)

	if TagsDiags.HasError() {
		resp.Diagnostics.Append(TagsDiags...)

		return
	}

	data.UpdatedAt = types.Int64Value(*Route.UpdatedAt)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RouteResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data = &RouteResourceModel{}

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	RuntimeGroupId := data.RuntimeGroupId.ValueString()
	RouteId := data.RouteId.ValueString()

	err := r.client.Routes.DeleteRoute(RuntimeGroupId, RouteId, shared.RequestOptions{})

	if err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting Route",
			err.Error(),
		)
	}
}

func (r *RouteResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data = &RouteResourceModel{}

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// TODO: add query params
	RuntimeGroupId := data.RuntimeGroupId.ValueString()
	RouteId := data.RouteId.ValueString()

	updateRequest := routes.Route{
		CreatedAt:               data.CreatedAt.ValueInt64Pointer(),
		HttpsRedirectStatusCode: data.HttpsRedirectStatusCode.ValueInt64Pointer(),
		Id:                      data.Id.ValueStringPointer(),
		Name:                    data.Name.ValueStringPointer(),
		PathHandling:            data.PathHandling.ValueStringPointer(),
		PreserveHost:            data.PreserveHost.ValueBoolPointer(),
		RegexPriority:           data.RegexPriority.ValueInt64Pointer(),
		RequestBuffering:        data.RequestBuffering.ValueBoolPointer(),
		ResponseBuffering:       data.ResponseBuffering.ValueBoolPointer(),
		Service: &routes.Service{
			Id: data.Service.Id.ValueStringPointer(),
		},
		StripPath: data.StripPath.ValueBoolPointer(),
		UpdatedAt: data.UpdatedAt.ValueInt64Pointer(),
	}

	Route, err := r.client.Routes.UpsertRoute(RuntimeGroupId, RouteId, updateRequest, shared.RequestOptions{})

	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating Route",
			err.Error(),
		)

		return
	}

	// TODO: this can probably be a function using reflection
	data.CreatedAt = types.Int64Value(*Route.CreatedAt)

	var HostsDiags diag.Diagnostics

	data.Hosts, HostsDiags = types.ListValueFrom(ctx, types.StringType, Route.Hosts)

	if HostsDiags.HasError() {
		resp.Diagnostics.Append(HostsDiags...)

		return
	}

	data.HttpsRedirectStatusCode = types.Int64Value(*Route.HttpsRedirectStatusCode)

	data.Id = types.StringValue(*Route.Id)

	var MethodsDiags diag.Diagnostics

	data.Methods, MethodsDiags = types.ListValueFrom(ctx, types.StringType, Route.Methods)

	if MethodsDiags.HasError() {
		resp.Diagnostics.Append(MethodsDiags...)

		return
	}

	data.Name = types.StringValue(*Route.Name)

	data.PathHandling = types.StringValue(*Route.PathHandling)

	var PathsDiags diag.Diagnostics

	data.Paths, PathsDiags = types.ListValueFrom(ctx, types.StringType, Route.Paths)

	if PathsDiags.HasError() {
		resp.Diagnostics.Append(PathsDiags...)

		return
	}

	data.PreserveHost = types.BoolValue(*Route.PreserveHost)

	var ProtocolsDiags diag.Diagnostics

	data.Protocols, ProtocolsDiags = types.ListValueFrom(ctx, types.StringType, Route.Protocols)

	if ProtocolsDiags.HasError() {
		resp.Diagnostics.Append(ProtocolsDiags...)

		return
	}

	data.RegexPriority = types.Int64Value(*Route.RegexPriority)

	data.RequestBuffering = types.BoolValue(*Route.RequestBuffering)

	data.ResponseBuffering = types.BoolValue(*Route.ResponseBuffering)

	data.Service = service.Service{
		Id: types.StringValue(*Route.Service.Id),
	}

	var SnisDiags diag.Diagnostics

	data.Snis, SnisDiags = types.ListValueFrom(ctx, types.StringType, Route.Snis)

	if SnisDiags.HasError() {
		resp.Diagnostics.Append(SnisDiags...)

		return
	}

	data.StripPath = types.BoolValue(*Route.StripPath)

	var TagsDiags diag.Diagnostics

	data.Tags, TagsDiags = types.ListValueFrom(ctx, types.StringType, Route.Tags)

	if TagsDiags.HasError() {
		resp.Diagnostics.Append(TagsDiags...)

		return
	}

	data.UpdatedAt = types.Int64Value(*Route.UpdatedAt)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RouteResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
