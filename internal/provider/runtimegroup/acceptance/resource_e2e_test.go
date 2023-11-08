// +build acceptance

package acceptance

import (
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"testing"
	"github.com/kong/internal/provider"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const (
	providerConfig = `
		provider "kong" {
        host = "host"
        auth_token = "auth_token"
}

	`
)

var (
	testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
		"kong": providerserver.NewProtocol6WithError(provider.New("test")()),
	}
)

func TestAcckongRunTimeGroupResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: providerConfig +
					`
resource "kong_run_time_group" "example" {
    name = "name"

    description = "description"

    labels = {
                name = "name"
}


    cluster_type = "cluster_type"

    auth_type = "auth_type"

}

`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Extend this based on the model attributes
					resource.TestCheckResourceAttr("kong_run_time_group.example", "name", "name"),
					resource.TestCheckResourceAttr("kong_run_time_group.example", "description", "description"),
					resource.TestCheckResourceAttr("kong_run_time_group.example", "labels.name", "name"),
					resource.TestCheckResourceAttr("kong_run_time_group.example", "cluster_type", "cluster_type"),
					resource.TestCheckResourceAttr("kong_run_time_group.example", "auth_type", "auth_type"),
				),
			},
		},
	})
}
