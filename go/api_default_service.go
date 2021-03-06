/*
 * Data Catalog Service - Asset Details
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"context"
	"encoding/json"
	"fmt"

	clientv3 "go.etcd.io/etcd/client/v3"
)

// DefaultApiService is a service that implements the logic for the DefaultApiServicer
// This service should implement the business logic for every endpoint for the DefaultApi API.
// Include any external packages or services that will be required by this service.
type DefaultApiService struct {
	conf clientv3.Config
}

// NewDefaultApiService creates a default api service
func NewDefaultApiService(confMap map[interface{}]interface{}) DefaultApiServicer {
	var conf clientv3.Config

	if confMap["etcd_username"] != nil && confMap["etcd_password"] != nil {
		conf = clientv3.Config{
			Endpoints: []string{fmt.Sprint(confMap["etcd_hostname"]) + ":" + fmt.Sprint(confMap["etcd_port"])},
			Username:  fmt.Sprint(confMap["etcd_username"]),
			Password:  fmt.Sprint(confMap["etcd_password"]),
		}
	} else {
		conf = clientv3.Config{
			Endpoints: []string{fmt.Sprint(confMap["etcd_hostname"]) + ":" + fmt.Sprint(confMap["etcd_port"])},
		}
	}
	uuidStr := fmt.Sprint(confMap["uuid"])
	fmt.Printf("Value: %s\n", uuidStr)

	return &DefaultApiService{conf}
}

func (s *DefaultApiService) getEtcdClient() *clientv3.Client {
	cli, err := clientv3.New(s.conf)

	if err != nil {
		// handle error!
	}

	return cli
}

// CreateAsset - This REST API writes data asset information to the data catalog configured in fybrik
func (s *DefaultApiService) CreateAsset(ctx context.Context, xRequestDatacatalogWriteCred string, createAssetRequest CreateAssetRequest, bodyBytes []byte) (ImplResponse, error) {
	// TODO - update CreateAsset with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(201, CreateAssetResponse{}) or use other options such as http.Ok ...
	//return Response(201, CreateAssetResponse{}), nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	assetID := createAssetRequest.DestinationCatalogID + "/" + createAssetRequest.DestinationAssetID

	cli := s.getEtcdClient()
	_, err := cli.Put(context.TODO(), assetID, string(bodyBytes))
	cli.Close()

	if err != nil {
		fmt.Printf("etcd Put operation failed: %v\n", err)
	}

	return Response(201, CreateAssetResponse{AssetID: assetID}), nil
}

// DeleteAsset - This REST API deletes data asset
func (s *DefaultApiService) DeleteAsset(ctx context.Context, xRequestDatacatalogCred string, deleteAssetRequest DeleteAssetRequest) (ImplResponse, error) {
	// TODO - update DeleteAsset with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, DeleteAssetResponse{}) or use other options such as http.Ok ...
	//return Response(200, DeleteAssetResponse{}), nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	//TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	//return Response(404, nil),nil

	//TODO: Uncomment the next line to return response Response(401, {}) or use other options such as http.Ok ...
	//return Response(401, nil),nil

	assetID := deleteAssetRequest.AssetID
	cli := s.getEtcdClient()
	dresp, err := cli.Delete(context.TODO(), assetID)
	cli.Close()
	if err != nil || dresp.Deleted == 0 {
		return Response(404, DeleteAssetResponse{"Asset ID not found"}), nil
	}
	return Response(200, DeleteAssetResponse{"Deletion Successful"}), nil
}

// GetAssetInfo - This REST API gets data asset information from the data catalog configured in fybrik for the data sets indicated in FybrikApplication yaml
func (s *DefaultApiService) GetAssetInfo(ctx context.Context, xRequestDatacatalogCred string, getAssetRequest GetAssetRequest) (ImplResponse, error) {
	// TODO - update GetAssetInfo with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	//TODO: Uncomment the next line to return response Response(200, GetAssetResponse{}) or use other options such as http.Ok ...
	assetID := getAssetRequest.AssetID

	cli := s.getEtcdClient()
	value, err := cli.Get(context.TODO(), assetID)
	cli.Close()

	if err != nil {
		fmt.Printf("etcd Get operation failed: %v\n", err)
	}

	// Declared an empty map interface
	var result map[string]interface{}
	json.Unmarshal(value.Kvs[0].Value, &result)

	return Response(200, result), nil
}
