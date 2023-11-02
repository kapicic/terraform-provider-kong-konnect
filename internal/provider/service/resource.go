package service

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/kong-sdk/pkg/client"
	"github.com/kong-sdk/pkg/services"
	"github.com/kong-sdk/pkg/shared"
	"github.com/kong/internal/provider/service/models/client_certificate"
	"github.com/kong/internal/utils"
)

// ensure we implement the needed interfaces
var _ resource.Resource = &ServiceResource{}
var _ resource.ResourceWithImportState = &ServiceResource{}

// constructor
func NewServiceResource() resource.Resource {
	return &ServiceResource{}
}

// client wrapper
type ServiceResource struct {
	client *client.Client
}

type ServiceResourceModel struct {
	CaCertificates    types.List                            `tfsdk:"ca_certificates"`
	ClientCertificate *client_certificate.ClientCertificate `tfsdk:"client_certificate"`
	ConnectTimeout    types.Int64                           `tfsdk:"connect_timeout"`
	CreatedAt         types.Int64                           `tfsdk:"created_at"`
	Enabled           types.Bool                            `tfsdk:"enabled"`
	Host              types.String                          `tfsdk:"host"`
	Id                types.String                          `tfsdk:"id"`
	Name              types.String                          `tfsdk:"name"`
	Path              types.String                          `tfsdk:"path"`
	Port              types.Int64                           `tfsdk:"port"`
	Protocol          types.String                          `tfsdk:"protocol"`
	ReadTimeout       types.Int64                           `tfsdk:"read_timeout"`
	Retries           types.Int64                           `tfsdk:"retries"`
	Tags              types.List                            `tfsdk:"tags"`
	TlsVerify         types.Bool                            `tfsdk:"tls_verify"`
	TlsVerifyDepth    types.Int64                           `tfsdk:"tls_verify_depth"`
	UpdatedAt         types.Int64                           `tfsdk:"updated_at"`
	Url               types.String                          `tfsdk:"url"`
	WriteTimeout      types.Int64                           `tfsdk:"write_timeout"`
	RuntimeGroupId    types.String                          `tfsdk:"runtime_group_id"`
	ServiceId         types.String                          `tfsdk:"service_id"`
}

func (r *ServiceResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_service"
}

func (r *ServiceResource) Schema(_ context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"ca_certificates": schema.ListAttribute{
				Description: "Array of `CA Certificate` object UUIDs that are used to build the trust store while verifying upstream server's TLS certificate. If set to `null` when Nginx default is respected. If default CA list in Nginx are not specified and TLS verification is enabled, then handshake with upstream server will always fail (because no CA are trusted).",
				Optional:    true,

				ElementType: types.StringType,
			},

			"client_certificate": schema.SingleNestedAttribute{
				Description: "Certificate to be used as client certificate while TLS handshaking to the upstream server.",
				Optional:    true,

				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Description: "id",
						Optional:    true,
					},
				},
			},

			"connect_timeout": schema.Int64Attribute{
				Description: "The timeout in milliseconds for establishing a connection to the upstream server.",
				Optional:    true,
			},

			"created_at": schema.Int64Attribute{
				Description: "Unix epoch when the resource was created.",
				Optional:    true,
			},

			"enabled": schema.BoolAttribute{
				Description: "Whether the service is active. If set to `false`, the proxy behavior will be as if any routes attached to it do not exist (404). Default: `true`.",
				Optional:    true,
			},

			"host": schema.StringAttribute{
				Description: "The host of the upstream server. Note that the host value is case sensitive.",
				Optional:    true,
			},

			"id": schema.StringAttribute{
				Description: "id",
				Optional:    true,
			},

			"name": schema.StringAttribute{
				Description: "The service name.",
				Optional:    true,
			},

			"path": schema.StringAttribute{
				Description: "The path to be used in requests to the upstream server.",
				Optional:    true,
			},

			"port": schema.Int64Attribute{
				Description: "The upstream server port.",
				Optional:    true,
			},

			"protocol": schema.StringAttribute{
				Description: "The protocol used to communicate with the upstream.",
				Optional:    true,
			},

			"read_timeout": schema.Int64Attribute{
				Description: "The timeout in milliseconds between two successive read operations for transmitting a request to the upstream server.",
				Optional:    true,
			},

			"retries": schema.Int64Attribute{
				Description: "The number of retries to execute upon failure to proxy.",
				Optional:    true,
			},

			"tags": schema.ListAttribute{
				Description: "An optional set of strings associated with the service for grouping and filtering.",
				Optional:    true,

				ElementType: types.StringType,
			},

			"tls_verify": schema.BoolAttribute{
				Description: "Whether to enable verification of upstream server TLS certificate. If set to `null`, then the Nginx default is respected.",
				Optional:    true,
			},

			"tls_verify_depth": schema.Int64Attribute{
				Description: "Maximum depth of chain while verifying Upstream server's TLS certificate. If set to `null`, then the Nginx default is respected.",
				Optional:    true,
			},

			"updated_at": schema.Int64Attribute{
				Description: "Unix epoch when the resource was last updated.",
				Optional:    true,
			},

			"url": schema.StringAttribute{
				Description: "Helper field to set `protocol`, `host`, `port` and `path` using a URL. This field is write-only and is not returned in responses.",
				Optional:    true,
			},

			"write_timeout": schema.Int64Attribute{
				Description: "The timeout in milliseconds between two successive write operations for transmitting a request to the upstream server.",
				Optional:    true,
			},

			"runtime_group_id": schema.StringAttribute{
				Description: "The ID of your runtime group. This variable is available in the Konnect manager",
				Optional:    true,
			},

			"service_id": schema.StringAttribute{
				Description: "ID **or** name of the service to lookup",
				Optional:    true,
			},
		},
	}
}

func (r *ServiceResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *ServiceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ServiceResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	RuntimeGroupId := data.RuntimeGroupId.ValueString()
	ServiceId := data.ServiceId.ValueString()

	requestOptions := shared.RequestOptions{}

	service, err := r.client.Services.GetService(RuntimeGroupId, ServiceId, requestOptions)

	if err != nil {
		resp.Diagnostics.AddError(
			"Unexpected error calling Services.GetService",
			err.Error(),
		)

		return
	}

	data.ConnectTimeout = utils.NullableInt64(service.ConnectTimeout)

	data.CreatedAt = utils.NullableInt64(service.CreatedAt)

	data.Enabled = utils.NullableBool(service.Enabled)

	data.Host = utils.NullableString(service.Host)

	data.Id = utils.NullableString(service.Id)

	data.Name = utils.NullableString(service.Name)

	data.Path = utils.NullableString(service.Path)

	data.Port = utils.NullableInt64(service.Port)

	data.Protocol = utils.NullableString(service.Protocol)

	data.ReadTimeout = utils.NullableInt64(service.ReadTimeout)

	data.Retries = utils.NullableInt64(service.Retries)

	data.UpdatedAt = utils.NullableInt64(service.UpdatedAt)

	data.WriteTimeout = utils.NullableInt64(service.WriteTimeout)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ServiceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ServiceResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	RuntimeGroupId := data.RuntimeGroupId.ValueString()

	requestOptions := shared.RequestOptions{}

	createRequest := services.Service{
		CaCertificates: utils.FromListToPrimitiveSlice[string](ctx, data.CaCertificates, types.StringType, &resp.Diagnostics),
		ClientCertificate: &services.ClientCertificate{
			Id: data.ClientCertificate.Id.ValueStringPointer(),
		},
		ConnectTimeout: data.ConnectTimeout.ValueInt64Pointer(),
		CreatedAt:      data.CreatedAt.ValueInt64Pointer(),
		Enabled:        data.Enabled.ValueBoolPointer(),
		Host:           data.Host.ValueStringPointer(),
		Id:             data.Id.ValueStringPointer(),
		Name:           data.Name.ValueStringPointer(),
		Path:           data.Path.ValueStringPointer(),
		Port:           data.Port.ValueInt64Pointer(),
		Protocol:       data.Protocol.ValueStringPointer(),
		ReadTimeout:    data.ReadTimeout.ValueInt64Pointer(),
		Retries:        data.Retries.ValueInt64Pointer(),
		Tags:           utils.FromListToPrimitiveSlice[string](ctx, data.Tags, types.StringType, &resp.Diagnostics),
		TlsVerify:      data.TlsVerify.ValueBoolPointer(),
		TlsVerifyDepth: data.TlsVerifyDepth.ValueInt64Pointer(),
		UpdatedAt:      data.UpdatedAt.ValueInt64Pointer(),
		Url:            data.Url.ValueStringPointer(),
		WriteTimeout:   data.WriteTimeout.ValueInt64Pointer(),
	}

	service, err := r.client.Services.CreateService(RuntimeGroupId, createRequest, requestOptions)

	if err != nil {
		resp.Diagnostics.AddError(
			"Error Creating Service",
			err.Error(),
		)

		return
	}

	data.ConnectTimeout = utils.NullableInt64(service.ConnectTimeout)

	data.CreatedAt = utils.NullableInt64(service.CreatedAt)

	data.Enabled = utils.NullableBool(service.Enabled)

	data.Host = utils.NullableString(service.Host)

	data.Id = utils.NullableString(service.Id)

	data.Name = utils.NullableString(service.Name)

	data.Path = utils.NullableString(service.Path)

	data.Port = utils.NullableInt64(service.Port)

	data.Protocol = utils.NullableString(service.Protocol)

	data.ReadTimeout = utils.NullableInt64(service.ReadTimeout)

	data.Retries = utils.NullableInt64(service.Retries)

	data.UpdatedAt = utils.NullableInt64(service.UpdatedAt)

	data.WriteTimeout = utils.NullableInt64(service.WriteTimeout)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ServiceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data = &ServiceResourceModel{}

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	requestOptions := shared.RequestOptions{}

	RuntimeGroupId := data.RuntimeGroupId.ValueString()
	ServiceId := data.ServiceId.ValueString()

	err := r.client.Services.DeleteService(RuntimeGroupId, ServiceId, requestOptions)

	if err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting Service",
			err.Error(),
		)
	}
}

func (r *ServiceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	var data = &ServiceResourceModel{}

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	requestOptions := shared.RequestOptions{}

	RuntimeGroupId := data.RuntimeGroupId.ValueString()
	ServiceId := data.ServiceId.ValueString()

	updateRequest := services.Service{
		CaCertificates: utils.FromListToPrimitiveSlice[string](ctx, data.CaCertificates, types.StringType, &resp.Diagnostics),
		ClientCertificate: &services.ClientCertificate{
			Id: data.ClientCertificate.Id.ValueStringPointer(),
		},
		ConnectTimeout: data.ConnectTimeout.ValueInt64Pointer(),
		CreatedAt:      data.CreatedAt.ValueInt64Pointer(),
		Enabled:        data.Enabled.ValueBoolPointer(),
		Host:           data.Host.ValueStringPointer(),
		Id:             data.Id.ValueStringPointer(),
		Name:           data.Name.ValueStringPointer(),
		Path:           data.Path.ValueStringPointer(),
		Port:           data.Port.ValueInt64Pointer(),
		Protocol:       data.Protocol.ValueStringPointer(),
		ReadTimeout:    data.ReadTimeout.ValueInt64Pointer(),
		Retries:        data.Retries.ValueInt64Pointer(),
		Tags:           utils.FromListToPrimitiveSlice[string](ctx, data.Tags, types.StringType, &resp.Diagnostics),
		TlsVerify:      data.TlsVerify.ValueBoolPointer(),
		TlsVerifyDepth: data.TlsVerifyDepth.ValueInt64Pointer(),
		UpdatedAt:      data.UpdatedAt.ValueInt64Pointer(),
		Url:            data.Url.ValueStringPointer(),
		WriteTimeout:   data.WriteTimeout.ValueInt64Pointer(),
	}

	service, err := r.client.Services.UpsertService(RuntimeGroupId, ServiceId, updateRequest, requestOptions)

	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating Service",
			err.Error(),
		)

		return
	}

	data.ConnectTimeout = utils.NullableInt64(service.ConnectTimeout)

	data.CreatedAt = utils.NullableInt64(service.CreatedAt)

	data.Enabled = utils.NullableBool(service.Enabled)

	data.Host = utils.NullableString(service.Host)

	data.Id = utils.NullableString(service.Id)

	data.Name = utils.NullableString(service.Name)

	data.Path = utils.NullableString(service.Path)

	data.Port = utils.NullableInt64(service.Port)

	data.Protocol = utils.NullableString(service.Protocol)

	data.ReadTimeout = utils.NullableInt64(service.ReadTimeout)

	data.Retries = utils.NullableInt64(service.Retries)

	data.UpdatedAt = utils.NullableInt64(service.UpdatedAt)

	data.WriteTimeout = utils.NullableInt64(service.WriteTimeout)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ServiceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
