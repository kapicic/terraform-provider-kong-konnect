package runtimegroups

import (
	"encoding/json"
	restClient "github.com/kong-sdk/internal/clients/rest"
	"github.com/kong-sdk/internal/clients/rest/httptransport"
	"github.com/kong-sdk/pkg/shared"
)

type RuntimeGroupsService struct {
	client  *restClient.RestClient
	baseUrl string
}

func NewRuntimeGroupsService(baseUrl string, bearerToken string) *RuntimeGroupsService {
	return &RuntimeGroupsService{
		client:  restClient.NewRestClient(baseUrl, bearerToken),
		baseUrl: baseUrl,
	}
}

func (api *RuntimeGroupsService) CreateRuntimeGroup(createRuntimeGroupRequest CreateRuntimeGroupRequest, opts shared.RequestOptions) (*RuntimeGroup, error) {
	request := httptransport.NewRequest("POST", api.baseUrl, "/runtime-groups", opts.Headers, opts.QueryParams)
	request.Body = createRuntimeGroupRequest

	resp, err := api.client.Call(request)
	if err != nil {
		return nil, err.GetError()
	}

	result := &RuntimeGroup{}
	jsonErr := json.Unmarshal(resp.Body, result)
	if err != nil {
		return nil, jsonErr
	}

	return result, nil

}

func (api *RuntimeGroupsService) GetRuntimeGroup(id string, opts shared.RequestOptions) (*RuntimeGroup, error) {
	request := httptransport.NewRequest("GET", api.baseUrl, "/runtime-groups/{id}", opts.Headers, opts.QueryParams)

	request.SetPathParam("id", id)

	resp, err := api.client.Call(request)
	if err != nil {
		return nil, err.GetError()
	}

	result := &RuntimeGroup{}
	jsonErr := json.Unmarshal(resp.Body, result)
	if err != nil {
		return nil, jsonErr
	}

	return result, nil

}

func (api *RuntimeGroupsService) UpdateRuntimeGroup(id string, updateRuntimeGroupRequest UpdateRuntimeGroupRequest, opts shared.RequestOptions) (*RuntimeGroup, error) {
	request := httptransport.NewRequest("PATCH", api.baseUrl, "/runtime-groups/{id}", opts.Headers, opts.QueryParams)
	request.Body = updateRuntimeGroupRequest

	request.SetPathParam("id", id)

	resp, err := api.client.Call(request)
	if err != nil {
		return nil, err.GetError()
	}

	result := &RuntimeGroup{}
	jsonErr := json.Unmarshal(resp.Body, result)
	if err != nil {
		return nil, jsonErr
	}

	return result, nil

}

func (api *RuntimeGroupsService) DeleteRuntimeGroup(id string, opts shared.RequestOptions) error {
	request := httptransport.NewRequest("DELETE", api.baseUrl, "/runtime-groups/{id}", opts.Headers, opts.QueryParams)

	request.SetPathParam("id", id)

	_, err := api.client.Call(request)
	if err != nil {
		return err.GetError()
	}

	return nil

}
