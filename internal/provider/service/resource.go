package service

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/liblab-sdk/pkg/client"
	"github.com/liblab-sdk/pkg/services"
	"github.com/liblab-sdk/pkg/shared"
	"github.com/liblab/internal/provider/service/models/client_certificate"
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
	ConnectTimeout    types.Int64                          `tfsdk:"connect_timeout"`
	CreatedAt         types.Int64                          `tfsdk:"created_at"`
	Enabled           types.Bool                           `tfsdk:"enabled"`
	Host              types.String                         `tfsdk:"host"`
	Id                types.String                         `tfsdk:"id"`
	Name              types.String                         `tfsdk:"name"`
	Path              types.String                         `tfsdk:"path"`
	Port              types.Int64                          `tfsdk:"port"`
	Protocol          types.String                         `tfsdk:"protocol"`
	ReadTimeout       types.Int64                          `tfsdk:"read_timeout"`
	Retries           types.Int64                          `tfsdk:"retries"`
	UpdatedAt         types.Int64                          `tfsdk:"updated_at"`
	WriteTimeout      types.Int64                          `tfsdk:"write_timeout"`
	CaCertificates    types.List                           `tfsdk:"ca_certificates"`
	ClientCertificate client_certificate.ClientCertificate `tfsdk:"client_certificate"`
	Tags              types.List                           `tfsdk:"tags"`
	TlsVerify         types.Bool                           `tfsdk:"tls_verify"`
	TlsVerifyDepth    types.Int64                          `tfsdk:"tls_verify_depth"`
	Url               types.String                         `tfsdk:"url"`
	RuntimeGroupId    types.String                         `tfsdk:"runtime_group_id"`
	ServiceId         types.String                         `tfsdk:"service_id"`
}

func (r *ServiceResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_service"
}

func (r *ServiceResource) Schema(_ context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"connect_timeout": schema.Int64Attribute{
				Optional: true,
			},

			"created_at": schema.Int64Attribute{
				Description: "Unix epoch when the resource was last created.",
				Optional:    true,
			},

			"enabled": schema.BoolAttribute{
				Description: "Service enabled boolean",
				Optional:    true,
			},

			"host": schema.StringAttribute{
				Optional: true,
			},

			"id": schema.StringAttribute{
				Optional: true,
			},

			"name": schema.StringAttribute{
				Optional: true,
			},

			"path": schema.StringAttribute{
				Optional: true,
			},

			"port": schema.Int64Attribute{
				Optional: true,
			},

			"protocol": schema.StringAttribute{
				Optional: true,
			},

			"read_timeout": schema.Int64Attribute{
				Optional: true,
			},

			"retries": schema.Int64Attribute{
				Optional: true,
			},

			"updated_at": schema.Int64Attribute{
				Optional: true,
			},

			"write_timeout": schema.Int64Attribute{
				Optional: true,
			},

			"ca_certificates": schema.ListAttribute{
				Description: "Array of `CA Certificate` object UUIDs that are used to build the trust store while verifying upstream server's TLS certificate. If set to `null` when Nginx default is respected. If default CA list in Nginx are not specified and TLS verification is enabled, then handshake with upstream server will always fail (because no CA are trusted).",
				ElementType: types.StringType,
				Optional:    true,
			},

			"client_certificate": schema.SingleNestedAttribute{
				Description: "Certificate to be used as client certificate while TLS handshaking to the upstream server.",
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Optional: true,
					},
				},
				Optional: true,
			},

			"tags": schema.ListAttribute{
				Description: "An optional set of strings associated with the service for grouping and filtering.",
				ElementType: types.StringType,
				Optional:    true,
			},

			"tls_verify": schema.BoolAttribute{
				Description: "Whether to enable verification of upstream server TLS certificate. If set to `null`, then the Nginx default is respected.",
				Optional:    true,
			},

			"tls_verify_depth": schema.Int64Attribute{
				Description: "Maximum depth of chain while verifying Upstream server's TLS certificate. If set to `null`, then the Nginx default is respected.",
				Optional:    true,
			},

			"url": schema.StringAttribute{
				Description: "Helper field to set `protocol`, `host`, `port` and `path` using a URL. This field is write-only and is not returned in responses.",
				Optional:    true,
			},

			"runtime_group_id": schema.StringAttribute{
				Description: "The ID of your runtime group. This variable is available in the Konnect manager",
				Required:    true,
			},

			"service_id": schema.StringAttribute{
				Description: "ID **or** name of the service to lookup",
				Required:    true,
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

	Service, err := r.client.Services.GetService(RuntimeGroupId, ServiceId, shared.RequestOptions{})

	if err != nil {
		resp.Diagnostics.AddError(
			"Unexpected error calling Services.GetService",
			err.Error(),
		)

		return
	}

	data.ConnectTimeout = types.Int64Value(*Service.ConnectTimeout)

	data.CreatedAt = types.Int64Value(*Service.CreatedAt)

	data.Enabled = types.BoolValue(*Service.Enabled)

	data.Host = types.StringValue(*Service.Host)

	data.Id = types.StringValue(*Service.Id)

	data.Name = types.StringValue(*Service.Name)

	data.Path = types.StringValue(*Service.Path)

	data.Port = types.Int64Value(*Service.Port)

	data.Protocol = types.StringValue(*Service.Protocol)

	data.ReadTimeout = types.Int64Value(*Service.ReadTimeout)

	data.Retries = types.Int64Value(*Service.Retries)

	data.UpdatedAt = types.Int64Value(*Service.UpdatedAt)

	data.WriteTimeout = types.Int64Value(*Service.WriteTimeout)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ServiceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ServiceResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	RuntimeGroupId := data.RuntimeGroupId.ValueString()

	// TODO: figure out struct name of createRequest
	createRequest := services.Service{
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
		TlsVerify:      data.TlsVerify.ValueBoolPointer(),
		TlsVerifyDepth: data.TlsVerifyDepth.ValueInt64Pointer(),
		UpdatedAt:      data.UpdatedAt.ValueInt64Pointer(),
		Url:            data.Url.ValueStringPointer(),
		WriteTimeout:   data.WriteTimeout.ValueInt64Pointer(),
	}

	// make request
	Service, err := r.client.Services.CreateService(RuntimeGroupId, createRequest, shared.RequestOptions{})

	if err != nil {
		resp.Diagnostics.AddError(
			"Error Creating Service",
			err.Error(),
		)

		return
	}

	// TODO: this can probably be a function using reflection
	data.ConnectTimeout = types.Int64Value(*Service.ConnectTimeout)

	data.CreatedAt = types.Int64Value(*Service.CreatedAt)

	data.Enabled = types.BoolValue(*Service.Enabled)

	data.Host = types.StringValue(*Service.Host)

	data.Id = types.StringValue(*Service.Id)

	data.Name = types.StringValue(*Service.Name)

	data.Path = types.StringValue(*Service.Path)

	data.Port = types.Int64Value(*Service.Port)

	data.Protocol = types.StringValue(*Service.Protocol)

	data.ReadTimeout = types.Int64Value(*Service.ReadTimeout)

	data.Retries = types.Int64Value(*Service.Retries)

	data.UpdatedAt = types.Int64Value(*Service.UpdatedAt)

	data.WriteTimeout = types.Int64Value(*Service.WriteTimeout)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ServiceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data = &ServiceResourceModel{}

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	RuntimeGroupId := data.RuntimeGroupId.ValueString()
	ServiceId := data.ServiceId.ValueString()

	err := r.client.Services.DeleteService(RuntimeGroupId, ServiceId, shared.RequestOptions{})

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

	// TODO: add query params
	RuntimeGroupId := data.RuntimeGroupId.ValueString()
	ServiceId := data.ServiceId.ValueString()

	updateRequest := services.Service{
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
		TlsVerify:      data.TlsVerify.ValueBoolPointer(),
		TlsVerifyDepth: data.TlsVerifyDepth.ValueInt64Pointer(),
		UpdatedAt:      data.UpdatedAt.ValueInt64Pointer(),
		Url:            data.Url.ValueStringPointer(),
		WriteTimeout:   data.WriteTimeout.ValueInt64Pointer(),
	}

	Service, err := r.client.Services.UpsertService(RuntimeGroupId, ServiceId, updateRequest, shared.RequestOptions{})

	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating Service",
			err.Error(),
		)

		return
	}

	// TODO: this can probably be a function using reflection
	data.ConnectTimeout = types.Int64Value(*Service.ConnectTimeout)

	data.CreatedAt = types.Int64Value(*Service.CreatedAt)

	data.Enabled = types.BoolValue(*Service.Enabled)

	data.Host = types.StringValue(*Service.Host)

	data.Id = types.StringValue(*Service.Id)

	data.Name = types.StringValue(*Service.Name)

	data.Path = types.StringValue(*Service.Path)

	data.Port = types.Int64Value(*Service.Port)

	data.Protocol = types.StringValue(*Service.Protocol)

	data.ReadTimeout = types.Int64Value(*Service.ReadTimeout)

	data.Retries = types.Int64Value(*Service.Retries)

	data.UpdatedAt = types.Int64Value(*Service.UpdatedAt)

	data.WriteTimeout = types.Int64Value(*Service.WriteTimeout)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ServiceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
