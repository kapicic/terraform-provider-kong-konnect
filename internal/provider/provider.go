package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/liblab-sdk/pkg/client"

	"github.com/liblab/internal/provider/route"

	"github.com/liblab/internal/provider/service"

	"github.com/liblab/internal/provider/apiproduct"

	"github.com/liblab/internal/provider/apiproductversion"

	"github.com/liblab/internal/provider/runtimegroup"
)

// Ensure Provider satisfies various provider interfaces.
var _ provider.Provider = &Provider{}

type Provider struct {
	version string
}

type liblabProviderModel struct {
	Host types.String `tfsdk:"host"`

	VerifyTls types.Bool `tfsdk:"verifytls"`
}

func (p *Provider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "liblab"
	resp.Version = "1.0.0"
}

func (p *Provider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{

			"host": schema.StringAttribute{
				Required:    true,
				Description: "The API host.",
			},

			"verify_tls": schema.BoolAttribute{
				Optional:    true,
				Description: "Whether or not to verify TLS.",
			},
		},
	}
}

func (p *Provider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data liblabProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	if data.Host.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("host"),
			"Unknown Host",
			"Cannot create API client with unknown host.",
		)
		return
	}

	apiClient := client.NewClient(data.Host.ValueString())

	// Example of setting the client in resp
	resp.DataSourceData = apiClient
	resp.ResourceData = apiClient
}

func (p *Provider) Resources(ctx context.Context) []func() resource.Resource {
	resources := []func() resource.Resource{}

	resources = append(resources, route.NewRouteResource)

	resources = append(resources, service.NewServiceResource)

	resources = append(resources, apiproduct.NewApiproductResource)

	resources = append(resources, apiproductversion.NewApiproductversionResource)

	resources = append(resources, runtimegroup.NewRuntimegroupResource)

	return resources
}

func (p *Provider) DataSources(ctx context.Context) []func() datasource.DataSource {
	dataSources := []func() datasource.DataSource{}

	return dataSources
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &Provider{
			version: version,
		}
	}
}
