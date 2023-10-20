package client

import (
	"github.com/kong-sdk/pkg/apiproducts"
	"github.com/kong-sdk/pkg/apiproductversions"
	"github.com/kong-sdk/pkg/routes"
	"github.com/kong-sdk/pkg/runtimegroups"
	"github.com/kong-sdk/pkg/services"
)

type Client struct {
	Routes             *routes.ApiService
	Services           *services.ApiService
	ApiProducts        *apiproducts.ApiService
	ApiProductVersions *apiproductversions.ApiService
	RuntimeGroups      *runtimegroups.ApiService
}

func NewClient(baseUrl string) *Client {

	return &Client{
		Routes:             routes.NewApiService(baseUrl),
		Services:           services.NewApiService(baseUrl),
		ApiProducts:        apiproducts.NewApiService(baseUrl),
		ApiProductVersions: apiproductversions.NewApiService(baseUrl),
		RuntimeGroups:      runtimegroups.NewApiService(baseUrl),
	}
}
