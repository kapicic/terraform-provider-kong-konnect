package client

import (
	"github.com/liblab-sdk/pkg/apiproducts"
	"github.com/liblab-sdk/pkg/apiproductversions"
	"github.com/liblab-sdk/pkg/routes"
	"github.com/liblab-sdk/pkg/runtimegroups"
	"github.com/liblab-sdk/pkg/services"
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
