package client

import (
		"github.com/kong-sdk/pkg/routes"
		"github.com/kong-sdk/pkg/services"
		"github.com/kong-sdk/pkg/apiproducts"
		"github.com/kong-sdk/pkg/apiproductversions"
		"github.com/kong-sdk/pkg/runtimegroups"
)

type Client struct {
		Routes *routes.RoutesService
		Services *services.ServicesService
		ApiProducts *apiproducts.ApiProductsService
		ApiProductVersions *apiproductversions.ApiProductVersionsService
		RuntimeGroups *runtimegroups.RuntimeGroupsService
}

func NewClient(baseUrl string, bearerToken string) *Client {

	return &Client{
			Routes: routes.NewRoutesService(baseUrl, bearerToken),
			Services: services.NewServicesService(baseUrl, bearerToken),
			ApiProducts: apiproducts.NewApiProductsService(baseUrl, bearerToken),
			ApiProductVersions: apiproductversions.NewApiProductVersionsService(baseUrl, bearerToken),
			RuntimeGroups: runtimegroups.NewRuntimeGroupsService(baseUrl, bearerToken),
	}
}
