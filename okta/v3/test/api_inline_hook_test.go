/*
Okta Admin Management

Testing InlineHookApiService

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

func Test_okta_InlineHookApiService(t *testing.T) {
	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test InlineHookApiService ActivateInlineHook", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var inlineHookId string

		resp, httpRes, err := apiClient.InlineHookApi.ActivateInlineHook(context.Background(), inlineHookId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test InlineHookApiService CreateInlineHook", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.InlineHookApi.CreateInlineHook(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test InlineHookApiService DeactivateInlineHook", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var inlineHookId string

		resp, httpRes, err := apiClient.InlineHookApi.DeactivateInlineHook(context.Background(), inlineHookId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test InlineHookApiService DeleteInlineHook", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var inlineHookId string

		resp, httpRes, err := apiClient.InlineHookApi.DeleteInlineHook(context.Background(), inlineHookId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test InlineHookApiService ExecuteInlineHook", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var inlineHookId string

		resp, httpRes, err := apiClient.InlineHookApi.ExecuteInlineHook(context.Background(), inlineHookId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test InlineHookApiService GetInlineHook", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var inlineHookId string

		resp, httpRes, err := apiClient.InlineHookApi.GetInlineHook(context.Background(), inlineHookId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test InlineHookApiService ListInlineHooks", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.InlineHookApi.ListInlineHooks(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test InlineHookApiService ReplaceInlineHook", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var inlineHookId string

		resp, httpRes, err := apiClient.InlineHookApi.ReplaceInlineHook(context.Background(), inlineHookId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})
}
