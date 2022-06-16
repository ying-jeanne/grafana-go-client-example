# \DashboardVersionsApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDashboardVersion**](DashboardVersionsApi.md#GetDashboardVersion) | **Get** /dashboards/id/{DashboardID}/versions/{DashboardVersionID} | Get a specific dashboard version.
[**GetDashboardVersionByUID**](DashboardVersionsApi.md#GetDashboardVersionByUID) | **Get** /dashboards/uid/{uid}/versions/{DashboardVersionID} | Get a specific dashboard version using UID.
[**GetDashboardVersions**](DashboardVersionsApi.md#GetDashboardVersions) | **Get** /dashboards/id/{DashboardID}/versions | Gets all existing versions for the dashboard.
[**GetDashboardVersionsByUID**](DashboardVersionsApi.md#GetDashboardVersionsByUID) | **Get** /dashboards/uid/{uid}/versions | Gets all existing versions for the dashboard using UID.
[**RestoreDashboardVersion**](DashboardVersionsApi.md#RestoreDashboardVersion) | **Post** /dashboards/id/{DashboardID}/restore | Restore a dashboard to a given dashboard version.
[**RestoreDashboardVersionByUID**](DashboardVersionsApi.md#RestoreDashboardVersionByUID) | **Post** /dashboards/uid/{uid}/restore | Restore a dashboard to a given dashboard version using UID.


# **GetDashboardVersion**
> DashboardVersionMetaModel GetDashboardVersion(ctx, dashboardID, dashboardVersionID)
Get a specific dashboard version.

Please refer to [updated API](#/dashboard_versions/getDashboardVersionByUID) instead

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dashboardID** | **int64**|  | 
  **dashboardVersionID** | **int64**|  | 

### Return type

[**DashboardVersionMetaModel**](DashboardVersionMeta.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDashboardVersionByUID**
> DashboardVersionMetaModel GetDashboardVersionByUID(ctx, uid, dashboardVersionID)
Get a specific dashboard version using UID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uid** | **string**|  | 
  **dashboardVersionID** | **int64**|  | 

### Return type

[**DashboardVersionMetaModel**](DashboardVersionMeta.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDashboardVersions**
> []DashboardVersionDtoModel GetDashboardVersions(ctx, dashboardID, optional)
Gets all existing versions for the dashboard.

Please refer to [updated API](#/dashboard_versions/getDashboardVersionsByUID) instead

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dashboardID** | **int64**|  | 
 **optional** | ***DashboardVersionsApiGetDashboardVersionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DashboardVersionsApiGetDashboardVersionsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int64**| Maximum number of results to return | [default to 0]
 **start** | **optional.Int64**| Version to start from when returning queries | [default to 0]

### Return type

[**[]DashboardVersionDtoModel**](DashboardVersionDTO.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDashboardVersionsByUID**
> []DashboardVersionDtoModel GetDashboardVersionsByUID(ctx, uid, optional)
Gets all existing versions for the dashboard using UID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uid** | **string**|  | 
 **optional** | ***DashboardVersionsApiGetDashboardVersionsByUIDOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DashboardVersionsApiGetDashboardVersionsByUIDOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int64**| Maximum number of results to return | [default to 0]
 **start** | **optional.Int64**| Version to start from when returning queries | [default to 0]

### Return type

[**[]DashboardVersionDtoModel**](DashboardVersionDTO.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestoreDashboardVersion**
> InlineResponse2004Model RestoreDashboardVersion(ctx, dashboardID, body)
Restore a dashboard to a given dashboard version.

Please refer to [updated API](#/dashboard_versions/restoreDashboardVersionByUID) instead

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dashboardID** | **int64**|  | 
  **body** | [**RestoreDashboardVersionCommandModel**](RestoreDashboardVersionCommandModel.md)|  | 

### Return type

[**InlineResponse2004Model**](inline_response_200_4.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestoreDashboardVersionByUID**
> InlineResponse2004Model RestoreDashboardVersionByUID(ctx, uid, body)
Restore a dashboard to a given dashboard version using UID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uid** | **string**|  | 
  **body** | [**RestoreDashboardVersionCommandModel**](RestoreDashboardVersionCommandModel.md)|  | 

### Return type

[**InlineResponse2004Model**](inline_response_200_4.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

