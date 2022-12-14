/*
Benchling API

Testing FeatureLibrariesApiService

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

func Test_benchlingsdk_FeatureLibrariesApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test FeatureLibrariesApiService BulkCreateFeatures", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.FeatureLibrariesApi.BulkCreateFeatures(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test FeatureLibrariesApiService CreateFeature", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.FeatureLibrariesApi.CreateFeature(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test FeatureLibrariesApiService CreateFeatureLibrary", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.FeatureLibrariesApi.CreateFeatureLibrary(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test FeatureLibrariesApiService GetFeature", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var featureId string

        resp, httpRes, err := apiClient.FeatureLibrariesApi.GetFeature(context.Background(), featureId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test FeatureLibrariesApiService GetFeatureLibrary", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var featureLibraryId string

        resp, httpRes, err := apiClient.FeatureLibrariesApi.GetFeatureLibrary(context.Background(), featureLibraryId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test FeatureLibrariesApiService ListFeatureLibraries", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.FeatureLibrariesApi.ListFeatureLibraries(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test FeatureLibrariesApiService ListFeatures", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.FeatureLibrariesApi.ListFeatures(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test FeatureLibrariesApiService UpdateFeature", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var featureId string

        resp, httpRes, err := apiClient.FeatureLibrariesApi.UpdateFeature(context.Background(), featureId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test FeatureLibrariesApiService UpdateFeatureLibrary", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var featureLibraryId string

        resp, httpRes, err := apiClient.FeatureLibrariesApi.UpdateFeatureLibrary(context.Background(), featureLibraryId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
