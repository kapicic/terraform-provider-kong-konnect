// +build unit

package apiproductversion

import (
	"context"
	"testing"
	"github.com/kong-sdk/pkg/client"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/stretchr/testify/assert"
)

func TestConfigureResource(t *testing.T) {
	// mock client
	mockClient := &client.Client{}

	// create ApiProductVersionResource instance
	resourceInstance := NewApiProductVersionResource()

	// Type-assert to the concrete type
	r, ok := resourceInstance.(*ApiProductVersionResource)
	if !ok {
		t.Fatalf("Failed to type assert resourceInstance to *ApiProductVersionResource")
	}

	// create mock ConfigureRequest
	req := resource.ConfigureRequest{
		ProviderData: mockClient,
	}

	var resp resource.ConfigureResponse

	r.Configure(context.Background(), req, &resp)

	// assertions
	assert.False(t, resp.Diagnostics.HasError())
	assert.Equal(t, mockClient, r.client, "Expected client to be set correctly in ApiProductVersionResource")
}
