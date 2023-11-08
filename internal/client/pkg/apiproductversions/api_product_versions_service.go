package apiproductversions

import (
	"encoding/json"
	restClient "github.com/kong-sdk/internal/clients/rest"
	"github.com/kong-sdk/internal/clients/rest/httptransport"
	"github.com/kong-sdk/pkg/shared"
)

type ApiProductVersionsService struct {
	client  *restClient.RestClient
	baseUrl string
}

func NewApiProductVersionsService(baseUrl string, bearerToken string) *ApiProductVersionsService {
	return &ApiProductVersionsService{
		client:  restClient.NewRestClient(baseUrl, bearerToken),
		baseUrl: baseUrl,
	}
}

func (api *ApiProductVersionsService) CreateApiProductVersion(apiProductId string, createApiProductVersionDto CreateApiProductVersionDto, opts shared.RequestOptions) (*ApiProductVersion, error) {
	request := httptransport.NewRequest("POST", api.baseUrl, "/api-products/{apiProductId}/product-versions", opts.Headers, opts.QueryParams)
	request.Body = createApiProductVersionDto

	request.SetPathParam("apiProductId", apiProductId)

	resp, err := api.client.Call(request)
	if err != nil {
		return nil, err.GetError()
	}

	result := &ApiProductVersion{}
	jsonErr := json.Unmarshal(resp.Body, result)
	if err != nil {
		return nil, jsonErr
	}

	return result, nil

}

func (api *ApiProductVersionsService) GetApiProductVersion(apiProductId string, id string, opts shared.RequestOptions) (*ApiProductVersion, error) {
	request := httptransport.NewRequest("GET", api.baseUrl, "/api-products/{apiProductId}/product-versions/{id}", opts.Headers, opts.QueryParams)

	request.SetPathParam("apiProductId", apiProductId)
	request.SetPathParam("id", id)

	resp, err := api.client.Call(request)
	if err != nil {
		return nil, err.GetError()
	}

	result := &ApiProductVersion{}
	jsonErr := json.Unmarshal(resp.Body, result)
	if err != nil {
		return nil, jsonErr
	}

	return result, nil

}

func (api *ApiProductVersionsService) UpdateApiProductVersion(apiProductId string, id string, updateApiProductVersionDto UpdateApiProductVersionDto, opts shared.RequestOptions) (*ApiProductVersion, error) {
	request := httptransport.NewRequest("PATCH", api.baseUrl, "/api-products/{apiProductId}/product-versions/{id}", opts.Headers, opts.QueryParams)
	request.Body = updateApiProductVersionDto

	request.SetPathParam("apiProductId", apiProductId)
	request.SetPathParam("id", id)

	resp, err := api.client.Call(request)
	if err != nil {
		return nil, err.GetError()
	}

	result := &ApiProductVersion{}
	jsonErr := json.Unmarshal(resp.Body, result)
	if err != nil {
		return nil, jsonErr
	}

	return result, nil

}

func (api *ApiProductVersionsService) DeleteApiProductVersion(apiProductId string, id string, opts shared.RequestOptions) error {
	request := httptransport.NewRequest("DELETE", api.baseUrl, "/api-products/{apiProductId}/product-versions/{id}", opts.Headers, opts.QueryParams)

	request.SetPathParam("apiProductId", apiProductId)
	request.SetPathParam("id", id)

	_, err := api.client.Call(request)
	if err != nil {
		return err.GetError()
	}

	return nil

}
