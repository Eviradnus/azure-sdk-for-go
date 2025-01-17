//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armservicefabric

// ApplicationTypeVersionsClientCreateOrUpdateResponse contains the response from method ApplicationTypeVersionsClient.CreateOrUpdate.
type ApplicationTypeVersionsClientCreateOrUpdateResponse struct {
	ApplicationTypeVersionResource
}

// ApplicationTypeVersionsClientDeleteResponse contains the response from method ApplicationTypeVersionsClient.Delete.
type ApplicationTypeVersionsClientDeleteResponse struct {
	// placeholder for future response values
}

// ApplicationTypeVersionsClientGetResponse contains the response from method ApplicationTypeVersionsClient.Get.
type ApplicationTypeVersionsClientGetResponse struct {
	ApplicationTypeVersionResource
}

// ApplicationTypeVersionsClientListResponse contains the response from method ApplicationTypeVersionsClient.List.
type ApplicationTypeVersionsClientListResponse struct {
	ApplicationTypeVersionResourceList
}

// ApplicationTypesClientCreateOrUpdateResponse contains the response from method ApplicationTypesClient.CreateOrUpdate.
type ApplicationTypesClientCreateOrUpdateResponse struct {
	ApplicationTypeResource
}

// ApplicationTypesClientDeleteResponse contains the response from method ApplicationTypesClient.Delete.
type ApplicationTypesClientDeleteResponse struct {
	// placeholder for future response values
}

// ApplicationTypesClientGetResponse contains the response from method ApplicationTypesClient.Get.
type ApplicationTypesClientGetResponse struct {
	ApplicationTypeResource
}

// ApplicationTypesClientListResponse contains the response from method ApplicationTypesClient.List.
type ApplicationTypesClientListResponse struct {
	ApplicationTypeResourceList
}

// ApplicationsClientCreateOrUpdateResponse contains the response from method ApplicationsClient.CreateOrUpdate.
type ApplicationsClientCreateOrUpdateResponse struct {
	ApplicationResource
}

// ApplicationsClientDeleteResponse contains the response from method ApplicationsClient.Delete.
type ApplicationsClientDeleteResponse struct {
	// placeholder for future response values
}

// ApplicationsClientGetResponse contains the response from method ApplicationsClient.Get.
type ApplicationsClientGetResponse struct {
	ApplicationResource
}

// ApplicationsClientListResponse contains the response from method ApplicationsClient.List.
type ApplicationsClientListResponse struct {
	ApplicationResourceList
}

// ApplicationsClientUpdateResponse contains the response from method ApplicationsClient.Update.
type ApplicationsClientUpdateResponse struct {
	ApplicationResource
}

// ClusterVersionsClientGetByEnvironmentResponse contains the response from method ClusterVersionsClient.GetByEnvironment.
type ClusterVersionsClientGetByEnvironmentResponse struct {
	ClusterCodeVersionsListResult
}

// ClusterVersionsClientGetResponse contains the response from method ClusterVersionsClient.Get.
type ClusterVersionsClientGetResponse struct {
	ClusterCodeVersionsListResult
}

// ClusterVersionsClientListByEnvironmentResponse contains the response from method ClusterVersionsClient.ListByEnvironment.
type ClusterVersionsClientListByEnvironmentResponse struct {
	ClusterCodeVersionsListResult
}

// ClusterVersionsClientListResponse contains the response from method ClusterVersionsClient.List.
type ClusterVersionsClientListResponse struct {
	ClusterCodeVersionsListResult
}

// ClustersClientCreateOrUpdateResponse contains the response from method ClustersClient.CreateOrUpdate.
type ClustersClientCreateOrUpdateResponse struct {
	Cluster
}

// ClustersClientDeleteResponse contains the response from method ClustersClient.Delete.
type ClustersClientDeleteResponse struct {
	// placeholder for future response values
}

// ClustersClientGetResponse contains the response from method ClustersClient.Get.
type ClustersClientGetResponse struct {
	Cluster
}

// ClustersClientListByResourceGroupResponse contains the response from method ClustersClient.ListByResourceGroup.
type ClustersClientListByResourceGroupResponse struct {
	ClusterListResult
}

// ClustersClientListResponse contains the response from method ClustersClient.List.
type ClustersClientListResponse struct {
	ClusterListResult
}

// ClustersClientListUpgradableVersionsResponse contains the response from method ClustersClient.ListUpgradableVersions.
type ClustersClientListUpgradableVersionsResponse struct {
	UpgradableVersionPathResult
}

// ClustersClientUpdateResponse contains the response from method ClustersClient.Update.
type ClustersClientUpdateResponse struct {
	Cluster
}

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	OperationListResult
}

// ServicesClientCreateOrUpdateResponse contains the response from method ServicesClient.CreateOrUpdate.
type ServicesClientCreateOrUpdateResponse struct {
	ServiceResource
}

// ServicesClientDeleteResponse contains the response from method ServicesClient.Delete.
type ServicesClientDeleteResponse struct {
	// placeholder for future response values
}

// ServicesClientGetResponse contains the response from method ServicesClient.Get.
type ServicesClientGetResponse struct {
	ServiceResource
}

// ServicesClientListResponse contains the response from method ServicesClient.List.
type ServicesClientListResponse struct {
	ServiceResourceList
}

// ServicesClientUpdateResponse contains the response from method ServicesClient.Update.
type ServicesClientUpdateResponse struct {
	ServiceResource
}
