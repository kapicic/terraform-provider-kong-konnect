//go:build unit
// +build unit

package apiproduct

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/kong-sdk/pkg/client"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfigureResource(t *testing.T) {
	// mock client
	mockClient := &client.Client{}

	// create ApiProductResource instance
	resourceInstance := NewApiProductResource()

	// Type-assert to the concrete type
	r, ok := resourceInstance.(*ApiProductResource)
	if !ok {
		t.Fatalf("Failed to type assert resourceInstance to *ApiProductResource")
	}

	// create mock ConfigureRequest
	req := resource.ConfigureRequest{
		ProviderData: mockClient,
	}

	var resp resource.ConfigureResponse

	r.Configure(context.Background(), req, &resp)

	// assertions
	assert.False(t, resp.Diagnostics.HasError())
	assert.Equal(t, mockClient, r.client, "Expected client to be set correctly in ApiProductResource")
}
