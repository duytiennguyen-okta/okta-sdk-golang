/*
Okta Admin Management

Testing ResourceSetApiService

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

func Test_okta_ResourceSetApiService(t *testing.T) {
	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ResourceSetApiService AddMembersToBinding", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var resourceSetId string
		var roleIdOrLabel string

		resp, httpRes, err := apiClient.ResourceSetApi.AddMembersToBinding(context.Background(), resourceSetId, roleIdOrLabel).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test ResourceSetApiService AddResourceSetResource", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var resourceSetId string

		resp, httpRes, err := apiClient.ResourceSetApi.AddResourceSetResource(context.Background(), resourceSetId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test ResourceSetApiService CreateResourceSet", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ResourceSetApi.CreateResourceSet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test ResourceSetApiService CreateResourceSetBinding", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var resourceSetId string

		resp, httpRes, err := apiClient.ResourceSetApi.CreateResourceSetBinding(context.Background(), resourceSetId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test ResourceSetApiService DeleteBinding", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var resourceSetId string
		var roleIdOrLabel string

		resp, httpRes, err := apiClient.ResourceSetApi.DeleteBinding(context.Background(), resourceSetId, roleIdOrLabel).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test ResourceSetApiService DeleteResourceSet", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var resourceSetId string

		resp, httpRes, err := apiClient.ResourceSetApi.DeleteResourceSet(context.Background(), resourceSetId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test ResourceSetApiService DeleteResourceSetResource", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var resourceSetId string
		var resourceId string

		resp, httpRes, err := apiClient.ResourceSetApi.DeleteResourceSetResource(context.Background(), resourceSetId, resourceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test ResourceSetApiService GetBinding", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var resourceSetId string
		var roleIdOrLabel string

		resp, httpRes, err := apiClient.ResourceSetApi.GetBinding(context.Background(), resourceSetId, roleIdOrLabel).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test ResourceSetApiService GetMemberOfBinding", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var resourceSetId string
		var roleIdOrLabel string
		var memberId string

		resp, httpRes, err := apiClient.ResourceSetApi.GetMemberOfBinding(context.Background(), resourceSetId, roleIdOrLabel, memberId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test ResourceSetApiService GetResourceSet", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var resourceSetId string

		resp, httpRes, err := apiClient.ResourceSetApi.GetResourceSet(context.Background(), resourceSetId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test ResourceSetApiService ListBindings", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var resourceSetId string

		resp, httpRes, err := apiClient.ResourceSetApi.ListBindings(context.Background(), resourceSetId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test ResourceSetApiService ListMembersOfBinding", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var resourceSetId string
		var roleIdOrLabel string

		resp, httpRes, err := apiClient.ResourceSetApi.ListMembersOfBinding(context.Background(), resourceSetId, roleIdOrLabel).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test ResourceSetApiService ListResourceSetResources", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var resourceSetId string

		resp, httpRes, err := apiClient.ResourceSetApi.ListResourceSetResources(context.Background(), resourceSetId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test ResourceSetApiService ListResourceSets", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ResourceSetApi.ListResourceSets(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test ResourceSetApiService ReplaceResourceSet", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var resourceSetId string

		resp, httpRes, err := apiClient.ResourceSetApi.ReplaceResourceSet(context.Background(), resourceSetId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test ResourceSetApiService UnassignMemberFromBinding", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var resourceSetId string
		var roleIdOrLabel string
		var memberId string

		resp, httpRes, err := apiClient.ResourceSetApi.UnassignMemberFromBinding(context.Background(), resourceSetId, roleIdOrLabel, memberId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})
}
