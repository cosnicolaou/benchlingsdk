/*
Benchling API

Testing AssayRunsApiService

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

func Test_benchlingsdk_AssayRunsApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test AssayRunsApiService ArchiveAssayRuns", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.AssayRunsApi.ArchiveAssayRuns(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test AssayRunsApiService BulkGetAssayRuns", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.AssayRunsApi.BulkGetAssayRuns(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test AssayRunsApiService CreateAssayRuns", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.AssayRunsApi.CreateAssayRuns(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test AssayRunsApiService GetAssayRun", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var assayRunId string

        resp, httpRes, err := apiClient.AssayRunsApi.GetAssayRun(context.Background(), assayRunId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test AssayRunsApiService ListAssayRuns", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.AssayRunsApi.ListAssayRuns(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test AssayRunsApiService ListAutomationInputGenerators", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var assayRunId string

        resp, httpRes, err := apiClient.AssayRunsApi.ListAutomationInputGenerators(context.Background(), assayRunId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test AssayRunsApiService ListAutomationOutputProcessorsDeprecated", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var assayRunId string

        resp, httpRes, err := apiClient.AssayRunsApi.ListAutomationOutputProcessorsDeprecated(context.Background(), assayRunId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test AssayRunsApiService UnarchiveAssayRuns", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.AssayRunsApi.UnarchiveAssayRuns(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test AssayRunsApiService UpdateAssayRun", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var assayRunId string

        resp, httpRes, err := apiClient.AssayRunsApi.UpdateAssayRun(context.Background(), assayRunId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
