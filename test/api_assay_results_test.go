/*
Benchling API

Testing AssayResultsApiService

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

func Test_benchlingsdk_AssayResultsApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test AssayResultsApiService AbortAssayResultsTransaction", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var transactionId string

        resp, httpRes, err := apiClient.AssayResultsApi.AbortAssayResultsTransaction(context.Background(), transactionId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test AssayResultsApiService ArchiveAssayResults", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.AssayResultsApi.ArchiveAssayResults(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test AssayResultsApiService BulkGetAssayResults", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.AssayResultsApi.BulkGetAssayResults(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test AssayResultsApiService CommitAssayResultsTransaction", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var transactionId string

        resp, httpRes, err := apiClient.AssayResultsApi.CommitAssayResultsTransaction(context.Background(), transactionId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test AssayResultsApiService CreateAssayResults", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.AssayResultsApi.CreateAssayResults(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test AssayResultsApiService CreateAssayResultsInTransaction", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var transactionId string

        resp, httpRes, err := apiClient.AssayResultsApi.CreateAssayResultsInTransaction(context.Background(), transactionId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test AssayResultsApiService CreateAssayResultsTransaction", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.AssayResultsApi.CreateAssayResultsTransaction(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test AssayResultsApiService GetAssayResult", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var assayResultId string

        resp, httpRes, err := apiClient.AssayResultsApi.GetAssayResult(context.Background(), assayResultId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test AssayResultsApiService ListAssayResults", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.AssayResultsApi.ListAssayResults(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test AssayResultsApiService UnarchiveAssayResults", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.AssayResultsApi.UnarchiveAssayResults(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
