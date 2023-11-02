package routes

import (
	"encoding/json"
	restClient "github.com/kong-sdk/internal/clients/rest"
	"github.com/kong-sdk/internal/clients/rest/httptransport"
	"github.com/kong-sdk/pkg/shared"
)

type RoutesService struct {
	client  *restClient.RestClient
	baseUrl string
}

func NewRoutesService(baseUrl string, bearerToken string) *RoutesService {
	return &RoutesService{
		client:  restClient.NewRestClient(baseUrl, bearerToken),
		baseUrl: baseUrl,
	}
}

func (api *RoutesService) CreateRoute(runtimeGroupId string, route Route, opts shared.RequestOptions) (*Route, error) {
	request := httptransport.NewRequest("POST", api.baseUrl, "/runtime-groups/{runtimeGroupId}/core-entities/routes", opts.Headers, opts.QueryParams)
	request.Body = route

	request.SetPathParam("runtimeGroupId", runtimeGroupId)

	resp, err := api.client.Call(request)
	if err != nil {
		return nil, err.GetError()
	}

	result := &Route{}
	jsonErr := json.Unmarshal(resp.Body, result)
	if err != nil {
		return nil, jsonErr
	}

	return result, nil

}

func (api *RoutesService) GetRoute(runtimeGroupId string, routeId string, opts shared.RequestOptions) (*Route, error) {
	request := httptransport.NewRequest("GET", api.baseUrl, "/runtime-groups/{runtimeGroupId}/core-entities/routes/{route_id}", opts.Headers, opts.QueryParams)

	request.SetPathParam("runtimeGroupId", runtimeGroupId)
	request.SetPathParam("route_id", routeId)

	resp, err := api.client.Call(request)
	if err != nil {
		return nil, err.GetError()
	}

	result := &Route{}
	jsonErr := json.Unmarshal(resp.Body, result)
	if err != nil {
		return nil, jsonErr
	}

	return result, nil

}

func (api *RoutesService) UpsertRoute(runtimeGroupId string, routeId string, route Route, opts shared.RequestOptions) (*Route, error) {
	request := httptransport.NewRequest("PUT", api.baseUrl, "/runtime-groups/{runtimeGroupId}/core-entities/routes/{route_id}", opts.Headers, opts.QueryParams)
	request.Body = route

	request.SetPathParam("runtimeGroupId", runtimeGroupId)
	request.SetPathParam("route_id", routeId)

	resp, err := api.client.Call(request)
	if err != nil {
		return nil, err.GetError()
	}

	result := &Route{}
	jsonErr := json.Unmarshal(resp.Body, result)
	if err != nil {
		return nil, jsonErr
	}

	return result, nil

}

func (api *RoutesService) DeleteRoute(runtimeGroupId string, routeId string, opts shared.RequestOptions) error {
	request := httptransport.NewRequest("DELETE", api.baseUrl, "/runtime-groups/{runtimeGroupId}/core-entities/routes/{route_id}", opts.Headers, opts.QueryParams)

	request.SetPathParam("runtimeGroupId", runtimeGroupId)
	request.SetPathParam("route_id", routeId)

	_, err := api.client.Call(request)
	if err != nil {
		return err.GetError()
	}

	return nil

}
