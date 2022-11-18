/*
Okta Admin Management

Testing CAPTCHAApiService

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

func Test_okta_CAPTCHAApiService(t *testing.T) {
	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test CAPTCHAApiService CreateCaptchaInstance", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.CAPTCHAApi.CreateCaptchaInstance(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test CAPTCHAApiService DeleteCaptchaInstance", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var captchaId string

		resp, httpRes, err := apiClient.CAPTCHAApi.DeleteCaptchaInstance(context.Background(), captchaId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test CAPTCHAApiService GetCaptchaInstance", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var captchaId string

		resp, httpRes, err := apiClient.CAPTCHAApi.GetCaptchaInstance(context.Background(), captchaId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test CAPTCHAApiService ListCaptchaInstances", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.CAPTCHAApi.ListCaptchaInstances(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test CAPTCHAApiService ReplaceCaptchaInstance", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var captchaId string

		resp, httpRes, err := apiClient.CAPTCHAApi.ReplaceCaptchaInstance(context.Background(), captchaId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test CAPTCHAApiService UpdateCaptchaInstance", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var captchaId string

		resp, httpRes, err := apiClient.CAPTCHAApi.UpdateCaptchaInstance(context.Background(), captchaId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})
}
