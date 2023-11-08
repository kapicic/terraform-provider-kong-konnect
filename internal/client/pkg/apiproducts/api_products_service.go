package apiproducts

import (
	"encoding/json"
	restClient "github.com/kong-sdk/internal/clients/rest"
	"github.com/kong-sdk/internal/clients/rest/httptransport"
	"github.com/kong-sdk/pkg/shared"
)

type ApiProductsService struct {
	client  *restClient.RestClient
	baseUrl string
}

func NewApiProductsService(baseUrl string, bearerToken string) *ApiProductsService {
	return &ApiProductsService{
		client:  restClient.NewRestClient(baseUrl, bearerToken),
		baseUrl: baseUrl,
	}
}

func (api *ApiProductsService) CreateApiProduct(createApiProductDto CreateApiProductDto, opts shared.RequestOptions) (*ApiProduct, error) {
	request := httptransport.NewRequest("POST", api.baseUrl, "/api-products", opts.Headers, opts.QueryParams)
	request.Body = createApiProductDto

	resp, err := api.client.Call(request)
	if err != nil {
		return nil, err.GetError()
	}

	result := &ApiProduct{}
	jsonErr := json.Unmarshal(resp.Body, result)
	if err != nil {
		return nil, jsonErr
	}

	return result, nil

}

func (api *ApiProductsService) GetApiProduct(id string, opts shared.RequestOptions) (*ApiProduct, error) {
	request := httptransport.NewRequest("GET", api.baseUrl, "/api-products/{id}", opts.Headers, opts.QueryParams)

	request.SetPathParam("id", id)

	resp, err := api.client.Call(request)
	if err != nil {
		return nil, err.GetError()
	}

	result := &ApiProduct{}
	jsonErr := json.Unmarshal(resp.Body, result)
	if err != nil {
		return nil, jsonErr
	}

	return result, nil

}

func (api *ApiProductsService) UpdateApiProduct(id string, updateApiProductDto UpdateApiProductDto, opts shared.RequestOptions) (*ApiProduct, error) {
	request := httptransport.NewRequest("PATCH", api.baseUrl, "/api-products/{id}", opts.Headers, opts.QueryParams)
	request.Body = updateApiProductDto

	request.SetPathParam("id", id)

	resp, err := api.client.Call(request)
	if err != nil {
		return nil, err.GetError()
	}

	result := &ApiProduct{}
	jsonErr := json.Unmarshal(resp.Body, result)
	if err != nil {
		return nil, jsonErr
	}

	return result, nil

}

func (api *ApiProductsService) DeleteApiProduct(id string, opts shared.RequestOptions) error {
	request := httptransport.NewRequest("DELETE", api.baseUrl, "/api-products/{id}", opts.Headers, opts.QueryParams)

	request.SetPathParam("id", id)

	_, err := api.client.Call(request)
	if err != nil {
		return err.GetError()
	}

	return nil

}
