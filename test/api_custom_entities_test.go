/*
Benchling API

Testing CustomEntitiesApiService

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

func Test_benchlingsdk_CustomEntitiesApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test CustomEntitiesApiService ArchiveCustomEntities", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.CustomEntitiesApi.ArchiveCustomEntities(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test CustomEntitiesApiService BulkCreateCustomEntities", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.CustomEntitiesApi.BulkCreateCustomEntities(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test CustomEntitiesApiService BulkGetCustomEntities", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.CustomEntitiesApi.BulkGetCustomEntities(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test CustomEntitiesApiService BulkUpdateCustomEntities", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.CustomEntitiesApi.BulkUpdateCustomEntities(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test CustomEntitiesApiService CreateCustomEntity", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.CustomEntitiesApi.CreateCustomEntity(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test CustomEntitiesApiService GetCustomEntity", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var customEntityId string

        resp, httpRes, err := apiClient.CustomEntitiesApi.GetCustomEntity(context.Background(), customEntityId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test CustomEntitiesApiService ListCustomEntities", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.CustomEntitiesApi.ListCustomEntities(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test CustomEntitiesApiService UnarchiveCustomEntities", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.CustomEntitiesApi.UnarchiveCustomEntities(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test CustomEntitiesApiService UpdateCustomEntity", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var customEntityId string

        resp, httpRes, err := apiClient.CustomEntitiesApi.UpdateCustomEntity(context.Background(), customEntityId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
