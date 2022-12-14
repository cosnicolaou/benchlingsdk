/*
Benchling API

Testing RegistryApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package benchlingsdk

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_benchlingsdk_RegistryApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test RegistryApiService BulkGetRegisteredEntities", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var registryId string

        resp, httpRes, err := apiClient.RegistryApi.BulkGetRegisteredEntities(context.Background(), registryId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test RegistryApiService GetRegistry", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var registryId string

        resp, httpRes, err := apiClient.RegistryApi.GetRegistry(context.Background(), registryId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test RegistryApiService ListBatchSchemasByRegistry", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var registryId string

        resp, httpRes, err := apiClient.RegistryApi.ListBatchSchemasByRegistry(context.Background(), registryId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test RegistryApiService ListBoxSchemasByRegistry", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var registryId string

        resp, httpRes, err := apiClient.RegistryApi.ListBoxSchemasByRegistry(context.Background(), registryId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test RegistryApiService ListContainerSchemasByRegistry", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var registryId string

        resp, httpRes, err := apiClient.RegistryApi.ListContainerSchemasByRegistry(context.Background(), registryId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test RegistryApiService ListDropdownsByRegistry", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var registryId string

        resp, httpRes, err := apiClient.RegistryApi.ListDropdownsByRegistry(context.Background(), registryId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test RegistryApiService ListEntitySchemasByRegistry", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var registryId string

        resp, httpRes, err := apiClient.RegistryApi.ListEntitySchemasByRegistry(context.Background(), registryId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test RegistryApiService ListLocationSchemasByRegistry", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var registryId string

        resp, httpRes, err := apiClient.RegistryApi.ListLocationSchemasByRegistry(context.Background(), registryId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test RegistryApiService ListPlateSchemasByRegistry", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var registryId string

        resp, httpRes, err := apiClient.RegistryApi.ListPlateSchemasByRegistry(context.Background(), registryId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test RegistryApiService ListRegistries", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.RegistryApi.ListRegistries(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test RegistryApiService RegisterEntities", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var registryId string

        resp, httpRes, err := apiClient.RegistryApi.RegisterEntities(context.Background(), registryId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test RegistryApiService UnregisterEntities", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var registryId string

        resp, httpRes, err := apiClient.RegistryApi.UnregisterEntities(context.Background(), registryId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
