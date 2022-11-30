/*
Benchling API

Testing AASequencesApiService

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

func Test_benchlingsdk_AASequencesApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test AASequencesApiService ArchiveAASequences", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.AASequencesApi.ArchiveAASequences(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test AASequencesApiService AutoAnnotateAaSequences", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.AASequencesApi.AutoAnnotateAaSequences(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test AASequencesApiService BulkCreateAASequences", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.AASequencesApi.BulkCreateAASequences(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test AASequencesApiService BulkGetAASequences", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.AASequencesApi.BulkGetAASequences(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test AASequencesApiService BulkUpdateAASequences", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.AASequencesApi.BulkUpdateAASequences(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test AASequencesApiService CreateAASequence", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.AASequencesApi.CreateAASequence(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test AASequencesApiService GetAASequence", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var aaSequenceId string

        resp, httpRes, err := apiClient.AASequencesApi.GetAASequence(context.Background(), aaSequenceId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test AASequencesApiService ListAASequences", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.AASequencesApi.ListAASequences(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test AASequencesApiService UnarchiveAASequences", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.AASequencesApi.UnarchiveAASequences(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test AASequencesApiService UpdateAASequence", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var aaSequenceId string

        resp, httpRes, err := apiClient.AASequencesApi.UpdateAASequence(context.Background(), aaSequenceId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}