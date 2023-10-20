package services

import (
	"encoding/json"

	restClient "github.com/kong-sdk/internal/clients/rest"
	"github.com/kong-sdk/internal/clients/rest/httptransport"
	"github.com/kong-sdk/pkg/shared"
)

type ApiService struct {
	client  *restClient.RestClient
	baseUrl string
}

func NewApiService(baseUrl string) *ApiService {
	return &ApiService{
		client:  restClient.NewRestClient(baseUrl),
		baseUrl: baseUrl,
	}
}

func (api *ApiService) CreateService(runtimeGroupId string, service Service, opts shared.RequestOptions) (*CreateService_201Response, error) {
	request := httptransport.NewRequest("POST", api.baseUrl, "/runtime-groups/{runtimeGroupId}/core-entities/services", opts.Headers, opts.QueryParams)
	request.Body = service

	request.SetPathParam("runtimeGroupId", runtimeGroupId)

	resp, err := api.client.Call(request)
	if err != nil {
		return nil, err.GetError()
	}

	result := &CreateService_201Response{}
	jsonErr := json.Unmarshal(resp.Body, result)
	if err != nil {
		return nil, jsonErr
	}

	return result, nil

}

func (api *ApiService) GetService(runtimeGroupId string, serviceId string, opts shared.RequestOptions) (*GetService_200Response, error) {
	request := httptransport.NewRequest("GET", api.baseUrl, "/runtime-groups/{runtimeGroupId}/core-entities/services/{service_id}", opts.Headers, opts.QueryParams)

	request.SetPathParam("runtimeGroupId", runtimeGroupId)
	request.SetPathParam("service_id", serviceId)

	resp, err := api.client.Call(request)
	if err != nil {
		return nil, err.GetError()
	}

	result := &GetService_200Response{}
	jsonErr := json.Unmarshal(resp.Body, result)
	if err != nil {
		return nil, jsonErr
	}

	return result, nil

}

func (api *ApiService) UpsertService(runtimeGroupId string, serviceId string, service Service, opts shared.RequestOptions) (*UpsertService_200Response, error) {
	request := httptransport.NewRequest("PUT", api.baseUrl, "/runtime-groups/{runtimeGroupId}/core-entities/services/{service_id}", opts.Headers, opts.QueryParams)
	request.Body = service

	request.SetPathParam("runtimeGroupId", runtimeGroupId)
	request.SetPathParam("service_id", serviceId)

	resp, err := api.client.Call(request)
	if err != nil {
		return nil, err.GetError()
	}

	result := &UpsertService_200Response{}
	jsonErr := json.Unmarshal(resp.Body, result)
	if err != nil {
		return nil, jsonErr
	}

	return result, nil

}

func (api *ApiService) DeleteService(runtimeGroupId string, serviceId string, opts shared.RequestOptions) error {
	request := httptransport.NewRequest("DELETE", api.baseUrl, "/runtime-groups/{runtimeGroupId}/core-entities/services/{service_id}", opts.Headers, opts.QueryParams)

	request.SetPathParam("runtimeGroupId", runtimeGroupId)
	request.SetPathParam("service_id", serviceId)

	_, err := api.client.Call(request)
	if err != nil {
		return err.GetError()
	}

	return nil

}
