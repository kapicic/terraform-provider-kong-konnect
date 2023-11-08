package route

import (
    "context"
    "fmt"
    "github.com/hashicorp/terraform-plugin-framework/resource"
    "github.com/hashicorp/terraform-plugin-framework/resource/schema"
    "github.com/hashicorp/terraform-plugin-framework/types"
    "github.com/hashicorp/terraform-plugin-framework/path"
    "github.com/kong-sdk/pkg/client"
    "github.com/kong-sdk/pkg/shared"
        "github.com/kong-sdk/pkg/routes"
        "github.com/kong/internal/provider/route/models/headers"
        "github.com/kong/internal/provider/route/models/service"
    "github.com/kong/internal/utils"
        "github.com/hashicorp/terraform-plugin-framework/diag"
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
    CreatedAt types.Int64 `tfsdk:"created_at"`
    Headers *headers.Headers `tfsdk:"headers"`
    Hosts types.List `tfsdk:"hosts"`
    HttpsRedirectStatusCode types.Int64 `tfsdk:"https_redirect_status_code"`
    Id types.String `tfsdk:"id"`
    Methods types.List `tfsdk:"methods"`
    Name types.String `tfsdk:"name"`
    PathHandling types.String `tfsdk:"path_handling"`
    Paths types.List `tfsdk:"paths"`
    PreserveHost types.Bool `tfsdk:"preserve_host"`
    Protocols types.List `tfsdk:"protocols"`
    RegexPriority types.Int64 `tfsdk:"regex_priority"`
    RequestBuffering types.Bool `tfsdk:"request_buffering"`
    ResponseBuffering types.Bool `tfsdk:"response_buffering"`
    Service *service.Service `tfsdk:"service"`
    Snis types.List `tfsdk:"snis"`
    StripPath types.Bool `tfsdk:"strip_path"`
    Tags types.List `tfsdk:"tags"`
    UpdatedAt types.Int64 `tfsdk:"updated_at"`
    RuntimeGroupId types.String `tfsdk:"runtime_group_id"`
    RouteId types.String `tfsdk:"route_id"`
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
    Optional: true,

},

            "headers": schema.SingleNestedAttribute{
        Description: "One or more lists of values indexed by header name that will cause this route to match if present in the request. The `Host` header cannot be used with this attribute: hosts should be specified using the `hosts` attribute. When `headers` contains only one value and that value starts with the special prefix `~*`, the value is interpreted as a regular expression.",
    Optional: true,

    Attributes: map[string]schema.Attribute{
            "key": schema.StringAttribute{
        Description: "key",
    Optional: true,

},

},

},

                "hosts": schema.ListAttribute{
        Description: "A list of domain names that match this route. Note that the hosts value is case sensitive.",
    Optional: true,

    ElementType: types.StringType,
},



            "https_redirect_status_code": schema.Int64Attribute{
        Description: "The status code Kong responds with when all properties of a route match except the protocol i.e. if the protocol of the request is `HTTP` instead of `HTTPS`. `Location` header is injected by Kong if the field is set to 301, 302, 307 or 308. Note: This config applies only if the route is configured to only accept the `https` protocol.",
    Optional: true,

},

            "id": schema.StringAttribute{
        Description: "id",
    Optional: true,

},

                "methods": schema.ListAttribute{
        Description: "A list of HTTP methods that match this route.",
    Optional: true,

    ElementType: types.StringType,
},



            "name": schema.StringAttribute{
        Description: "The name of the route. Route names must be unique, and they are case sensitive. For example, there can be two different routes named test and Test.",
    Optional: true,

},

            "path_handling": schema.StringAttribute{
        Description: "Controls how the service path, route path and requested path are combined when sending a request to the upstream. See above for a detailed description of each behavior.",
    Optional: true,

},

                "paths": schema.ListAttribute{
        Description: "A list of paths that match this route.",
    Optional: true,

    ElementType: types.StringType,
},



            "preserve_host": schema.BoolAttribute{
        Description: "When matching a route via one of the `hosts` domain names, use the request `Host` header in the upstream request headers. If set to `false`, the upstream `Host` header will be that of the service's `host`.",
    Optional: true,

},

                "protocols": schema.ListAttribute{
        Description: "An array of the protocols this route should allow. See the [route Object](#route-object) section for a list of accepted protocols. When set to only `https`, HTTP requests are answered with an upgrade error. When set to only `http`, HTTPS requests are answered with an error.",
    Optional: true,

    ElementType: types.StringType,
},



            "regex_priority": schema.Int64Attribute{
        Description: "A number used to choose which route resolves a given request when several routes match it using regexes simultaneously. When two routes match the path and have the same `regex_priority`, the older one (lowest `created_at`) is used. Note that the priority for non-regex routes is different (longer non-regex routes are matched before shorter ones).",
    Optional: true,

},

            "request_buffering": schema.BoolAttribute{
        Description: "Whether to enable request body buffering or not. With HTTP 1.1, it may make sense to turn this off on services that receive data with chunked transfer encoding.",
    Optional: true,

},

            "response_buffering": schema.BoolAttribute{
        Description: "Whether to enable response body buffering or not. With HTTP 1.1, it may make sense to turn this off on services that send data with chunked transfer encoding.",
    Optional: true,

},

            "service": schema.SingleNestedAttribute{
        Description: "The service this route is associated to. This is where the route proxies traffic to.",
    Optional: true,

    Attributes: map[string]schema.Attribute{
            "id": schema.StringAttribute{
        Description: "id",
    Optional: true,

},

},

},

                "snis": schema.ListAttribute{
        Description: "A list of SNIs that match this route when using stream routing.",
    Optional: true,

    ElementType: types.StringType,
},



            "strip_path": schema.BoolAttribute{
        Description: "When matching a route via one of the `paths`, strip the matching prefix from the upstream request URL.",
    Optional: true,

},

                "tags": schema.ListAttribute{
        Description: "An optional set of strings associated with the route for grouping and filtering.",
    Optional: true,

    ElementType: types.StringType,
},



            "updated_at": schema.Int64Attribute{
        Description: "Unix epoch when the resource was last updated.",
    Optional: true,

},

            "runtime_group_id": schema.StringAttribute{
        Description: "The ID of your runtime group. This variable is available in the Konnect manager",
    Optional: true,

},

            "route_id": schema.StringAttribute{
        Description: "The unique identifier or the name of the route to retrieve.",
    Optional: true,

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
	utils.PopulateModelData(ctx, &data, resp.Diagnostics, req.State.Get)

	if resp.Diagnostics.HasError() {
		return
	}

			RuntimeGroupId := data.RuntimeGroupId.ValueString()
			RouteId := data.RouteId.ValueString()

	requestOptions := shared.RequestOptions{}


	route, err := r.client.Routes.GetRoute(RuntimeGroupId, RouteId, requestOptions)

	if err != nil {
		resp.Diagnostics.AddError(
			"Unexpected error calling Routes.GetRoute",
			err.Error(),
		)

		return
	}

	        data.CreatedAt = utils.NullableInt64(route.GetCreatedAt())

	        if route.Headers != nil {
        data.Headers = utils.NullableObject(route.Headers, headers.Headers{
                    Key: utils.NullableString(route.GetHeaders().GetKey()),

        })
    }

	        var HostsDiags diag.Diagnostics

    data.Hosts, HostsDiags = types.ListValueFrom(ctx, types.StringType, route.Hosts)
    if HostsDiags.HasError() {
        resp.Diagnostics.Append(HostsDiags...)
    }

	        data.HttpsRedirectStatusCode = utils.NullableInt64(route.GetHttpsRedirectStatusCode())

	        data.Id = utils.NullableString(route.GetId())

	        var MethodsDiags diag.Diagnostics

    data.Methods, MethodsDiags = types.ListValueFrom(ctx, types.StringType, route.Methods)
    if MethodsDiags.HasError() {
        resp.Diagnostics.Append(MethodsDiags...)
    }

	        data.Name = utils.NullableString(route.GetName())

	        data.PathHandling = utils.NullableString(route.GetPathHandling())

	        var PathsDiags diag.Diagnostics

    data.Paths, PathsDiags = types.ListValueFrom(ctx, types.StringType, route.Paths)
    if PathsDiags.HasError() {
        resp.Diagnostics.Append(PathsDiags...)
    }

	        data.PreserveHost = utils.NullableBool(route.GetPreserveHost())

	        var ProtocolsDiags diag.Diagnostics

    data.Protocols, ProtocolsDiags = types.ListValueFrom(ctx, types.StringType, route.Protocols)
    if ProtocolsDiags.HasError() {
        resp.Diagnostics.Append(ProtocolsDiags...)
    }

	        data.RegexPriority = utils.NullableInt64(route.GetRegexPriority())

	        data.RequestBuffering = utils.NullableBool(route.GetRequestBuffering())

	        data.ResponseBuffering = utils.NullableBool(route.GetResponseBuffering())

	        if route.Service != nil {
        data.Service = utils.NullableObject(route.Service, service.Service{
                    Id: utils.NullableString(route.GetService().GetId()),

        })
    }

	        var SnisDiags diag.Diagnostics

    data.Snis, SnisDiags = types.ListValueFrom(ctx, types.StringType, route.Snis)
    if SnisDiags.HasError() {
        resp.Diagnostics.Append(SnisDiags...)
    }

	        data.StripPath = utils.NullableBool(route.GetStripPath())

	        var TagsDiags diag.Diagnostics

    data.Tags, TagsDiags = types.ListValueFrom(ctx, types.StringType, route.Tags)
    if TagsDiags.HasError() {
        resp.Diagnostics.Append(TagsDiags...)
    }

	        data.UpdatedAt = utils.NullableInt64(route.GetUpdatedAt())


	if (resp.Diagnostics.HasError()) {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}


func (r *RouteResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
    var data RouteResourceModel
    utils.PopulateModelData(ctx, &data, resp.Diagnostics, req.Plan.Get)

    if resp.Diagnostics.HasError() {
        return
    }

            RuntimeGroupId := data.RuntimeGroupId.ValueString()

    requestOptions := shared.RequestOptions{}


        createRequest := routes.Route{
            CreatedAt: data.CreatedAt.ValueInt64Pointer(),
            
            Headers: utils.NullableTfStateObject(data.Headers, func(from *headers.Headers) routes.Headers {
                return routes.Headers{
                            Key: from.Key.ValueStringPointer(),
                }
            }),
                Hosts: utils.FromListToPrimitiveSlice[string](ctx, data.Hosts, types.StringType, &resp.Diagnostics),
            HttpsRedirectStatusCode: data.HttpsRedirectStatusCode.ValueInt64Pointer(),
            Id: data.Id.ValueStringPointer(),
                Methods: utils.FromListToPrimitiveSlice[string](ctx, data.Methods, types.StringType, &resp.Diagnostics),
            Name: data.Name.ValueStringPointer(),
            PathHandling: data.PathHandling.ValueStringPointer(),
                Paths: utils.FromListToPrimitiveSlice[string](ctx, data.Paths, types.StringType, &resp.Diagnostics),
            PreserveHost: data.PreserveHost.ValueBoolPointer(),
                Protocols: utils.FromListToPrimitiveSlice[string](ctx, data.Protocols, types.StringType, &resp.Diagnostics),
            RegexPriority: data.RegexPriority.ValueInt64Pointer(),
            RequestBuffering: data.RequestBuffering.ValueBoolPointer(),
            ResponseBuffering: data.ResponseBuffering.ValueBoolPointer(),
            
            Service: utils.NullableTfStateObject(data.Service, func(from *service.Service) routes.Service {
                return routes.Service{
                            Id: from.Id.ValueStringPointer(),
                }
            }),
                Snis: utils.FromListToPrimitiveSlice[string](ctx, data.Snis, types.StringType, &resp.Diagnostics),
            StripPath: data.StripPath.ValueBoolPointer(),
                Tags: utils.FromListToPrimitiveSlice[string](ctx, data.Tags, types.StringType, &resp.Diagnostics),
            UpdatedAt: data.UpdatedAt.ValueInt64Pointer(),
}
  
        route, err := r.client.Routes.CreateRoute(RuntimeGroupId, createRequest, requestOptions)

    if err != nil {
        resp.Diagnostics.AddError(
            "Error Creating Route",
            err.Error(),
        )

        return
    }

            data.CreatedAt = utils.NullableInt64(route.GetCreatedAt())

            if route.Headers != nil {
        data.Headers = utils.NullableObject(route.Headers, headers.Headers{
                    Key: utils.NullableString(route.GetHeaders().GetKey()),

        })
    }

            var HostsDiags diag.Diagnostics

    data.Hosts, HostsDiags = types.ListValueFrom(ctx, types.StringType, route.Hosts)
    if HostsDiags.HasError() {
        resp.Diagnostics.Append(HostsDiags...)
    }

            data.HttpsRedirectStatusCode = utils.NullableInt64(route.GetHttpsRedirectStatusCode())

            data.Id = utils.NullableString(route.GetId())

            var MethodsDiags diag.Diagnostics

    data.Methods, MethodsDiags = types.ListValueFrom(ctx, types.StringType, route.Methods)
    if MethodsDiags.HasError() {
        resp.Diagnostics.Append(MethodsDiags...)
    }

            data.Name = utils.NullableString(route.GetName())

            data.PathHandling = utils.NullableString(route.GetPathHandling())

            var PathsDiags diag.Diagnostics

    data.Paths, PathsDiags = types.ListValueFrom(ctx, types.StringType, route.Paths)
    if PathsDiags.HasError() {
        resp.Diagnostics.Append(PathsDiags...)
    }

            data.PreserveHost = utils.NullableBool(route.GetPreserveHost())

            var ProtocolsDiags diag.Diagnostics

    data.Protocols, ProtocolsDiags = types.ListValueFrom(ctx, types.StringType, route.Protocols)
    if ProtocolsDiags.HasError() {
        resp.Diagnostics.Append(ProtocolsDiags...)
    }

            data.RegexPriority = utils.NullableInt64(route.GetRegexPriority())

            data.RequestBuffering = utils.NullableBool(route.GetRequestBuffering())

            data.ResponseBuffering = utils.NullableBool(route.GetResponseBuffering())

            if route.Service != nil {
        data.Service = utils.NullableObject(route.Service, service.Service{
                    Id: utils.NullableString(route.GetService().GetId()),

        })
    }

            var SnisDiags diag.Diagnostics

    data.Snis, SnisDiags = types.ListValueFrom(ctx, types.StringType, route.Snis)
    if SnisDiags.HasError() {
        resp.Diagnostics.Append(SnisDiags...)
    }

            data.StripPath = utils.NullableBool(route.GetStripPath())

            var TagsDiags diag.Diagnostics

    data.Tags, TagsDiags = types.ListValueFrom(ctx, types.StringType, route.Tags)
    if TagsDiags.HasError() {
        resp.Diagnostics.Append(TagsDiags...)
    }

            data.UpdatedAt = utils.NullableInt64(route.GetUpdatedAt())


    if (resp.Diagnostics.HasError()) {
        return
    }

    resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}


func (r *RouteResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
    var data = &RouteResourceModel{}
    utils.PopulateModelData(ctx, &data, resp.Diagnostics, req.State.Get)

    if resp.Diagnostics.HasError() {
        return
    }

    requestOptions := shared.RequestOptions{}


            RuntimeGroupId := data.RuntimeGroupId.ValueString()
            RouteId := data.RouteId.ValueString()

    err := r.client.Routes.DeleteRoute(RuntimeGroupId, RouteId, requestOptions)

    if err != nil {
        resp.Diagnostics.AddError(
            "Error Deleting Route",
            err.Error(),
        )
    }
}


func (r *RouteResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
        var data = &RouteResourceModel{}
        utils.PopulateModelData(ctx, &data, resp.Diagnostics, req.Plan.Get)

        if resp.Diagnostics.HasError() {
            return
        }

        requestOptions := shared.RequestOptions{}


                RuntimeGroupId := data.RuntimeGroupId.ValueString()
                RouteId := data.RouteId.ValueString()

        updateRequest := routes.Route{
            CreatedAt: data.CreatedAt.ValueInt64Pointer(),
            
            Headers: utils.NullableTfStateObject(data.Headers, func(from *headers.Headers) routes.Headers {
                return routes.Headers{
                            Key: from.Key.ValueStringPointer(),
                }
            }),
                Hosts: utils.FromListToPrimitiveSlice[string](ctx, data.Hosts, types.StringType, &resp.Diagnostics),
            HttpsRedirectStatusCode: data.HttpsRedirectStatusCode.ValueInt64Pointer(),
            Id: data.Id.ValueStringPointer(),
                Methods: utils.FromListToPrimitiveSlice[string](ctx, data.Methods, types.StringType, &resp.Diagnostics),
            Name: data.Name.ValueStringPointer(),
            PathHandling: data.PathHandling.ValueStringPointer(),
                Paths: utils.FromListToPrimitiveSlice[string](ctx, data.Paths, types.StringType, &resp.Diagnostics),
            PreserveHost: data.PreserveHost.ValueBoolPointer(),
                Protocols: utils.FromListToPrimitiveSlice[string](ctx, data.Protocols, types.StringType, &resp.Diagnostics),
            RegexPriority: data.RegexPriority.ValueInt64Pointer(),
            RequestBuffering: data.RequestBuffering.ValueBoolPointer(),
            ResponseBuffering: data.ResponseBuffering.ValueBoolPointer(),
            
            Service: utils.NullableTfStateObject(data.Service, func(from *service.Service) routes.Service {
                return routes.Service{
                            Id: from.Id.ValueStringPointer(),
                }
            }),
                Snis: utils.FromListToPrimitiveSlice[string](ctx, data.Snis, types.StringType, &resp.Diagnostics),
            StripPath: data.StripPath.ValueBoolPointer(),
                Tags: utils.FromListToPrimitiveSlice[string](ctx, data.Tags, types.StringType, &resp.Diagnostics),
            UpdatedAt: data.UpdatedAt.ValueInt64Pointer(),
}
  

        route, err := r.client.Routes.UpsertRoute(RuntimeGroupId, RouteId, updateRequest, requestOptions)

        if err != nil {
            resp.Diagnostics.AddError(
                "Error updating Route",
                err.Error(),
            )

            return
        }

                data.CreatedAt = utils.NullableInt64(route.GetCreatedAt())

                if route.Headers != nil {
        data.Headers = utils.NullableObject(route.Headers, headers.Headers{
                    Key: utils.NullableString(route.GetHeaders().GetKey()),

        })
    }

                var HostsDiags diag.Diagnostics

    data.Hosts, HostsDiags = types.ListValueFrom(ctx, types.StringType, route.Hosts)
    if HostsDiags.HasError() {
        resp.Diagnostics.Append(HostsDiags...)
    }

                data.HttpsRedirectStatusCode = utils.NullableInt64(route.GetHttpsRedirectStatusCode())

                data.Id = utils.NullableString(route.GetId())

                var MethodsDiags diag.Diagnostics

    data.Methods, MethodsDiags = types.ListValueFrom(ctx, types.StringType, route.Methods)
    if MethodsDiags.HasError() {
        resp.Diagnostics.Append(MethodsDiags...)
    }

                data.Name = utils.NullableString(route.GetName())

                data.PathHandling = utils.NullableString(route.GetPathHandling())

                var PathsDiags diag.Diagnostics

    data.Paths, PathsDiags = types.ListValueFrom(ctx, types.StringType, route.Paths)
    if PathsDiags.HasError() {
        resp.Diagnostics.Append(PathsDiags...)
    }

                data.PreserveHost = utils.NullableBool(route.GetPreserveHost())

                var ProtocolsDiags diag.Diagnostics

    data.Protocols, ProtocolsDiags = types.ListValueFrom(ctx, types.StringType, route.Protocols)
    if ProtocolsDiags.HasError() {
        resp.Diagnostics.Append(ProtocolsDiags...)
    }

                data.RegexPriority = utils.NullableInt64(route.GetRegexPriority())

                data.RequestBuffering = utils.NullableBool(route.GetRequestBuffering())

                data.ResponseBuffering = utils.NullableBool(route.GetResponseBuffering())

                if route.Service != nil {
        data.Service = utils.NullableObject(route.Service, service.Service{
                    Id: utils.NullableString(route.GetService().GetId()),

        })
    }

                var SnisDiags diag.Diagnostics

    data.Snis, SnisDiags = types.ListValueFrom(ctx, types.StringType, route.Snis)
    if SnisDiags.HasError() {
        resp.Diagnostics.Append(SnisDiags...)
    }

                data.StripPath = utils.NullableBool(route.GetStripPath())

                var TagsDiags diag.Diagnostics

    data.Tags, TagsDiags = types.ListValueFrom(ctx, types.StringType, route.Tags)
    if TagsDiags.HasError() {
        resp.Diagnostics.Append(TagsDiags...)
    }

                data.UpdatedAt = utils.NullableInt64(route.GetUpdatedAt())


        if (resp.Diagnostics.HasError()) {
            return
        }

        resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}


func (r *RouteResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
    // Retrieve import ID and save to id attribute
    resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
