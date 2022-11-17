/*
Okta Admin Management

Testing ThreatInsightApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package okta

import (
	"context"
	"testing"

	openapiclient "./openapi"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_okta_ThreatInsightApiService(t *testing.T) {
	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ThreatInsightApiService GetCurrentConfiguration", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ThreatInsightApi.GetCurrentConfiguration(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test ThreatInsightApiService UpdateConfiguration", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ThreatInsightApi.UpdateConfiguration(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})
}
