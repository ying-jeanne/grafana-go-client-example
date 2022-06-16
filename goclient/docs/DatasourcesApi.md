# \DatasourcesApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddDatasource**](DatasourcesApi.md#AddDatasource) | **Post** /datasources | Create a data source.
[**CheckDatasourceHealth**](DatasourcesApi.md#CheckDatasourceHealth) | **Get** /datasources/uid/{uid}/health | Check data source health by Id.
[**CheckDatasourceHealthByID**](DatasourcesApi.md#CheckDatasourceHealthByID) | **Get** /datasources/{id}/health | Check data source health by Id.
[**DatasourceProxyDELETEByUIDcalls**](DatasourcesApi.md#DatasourceProxyDELETEByUIDcalls) | **Delete** /datasources/proxy/uid/{uid}/{datasource_proxy_route} | Data source proxy DELETE calls.
[**DatasourceProxyDELETEcalls**](DatasourcesApi.md#DatasourceProxyDELETEcalls) | **Delete** /datasources/proxy/{id}/{datasource_proxy_route} | Data source proxy DELETE calls.
[**DatasourceProxyGETByUIDcalls**](DatasourcesApi.md#DatasourceProxyGETByUIDcalls) | **Get** /datasources/proxy/uid/{uid}/{datasource_proxy_route} | Data source proxy GET calls.
[**DatasourceProxyGETcalls**](DatasourcesApi.md#DatasourceProxyGETcalls) | **Get** /datasources/proxy/{id}/{datasource_proxy_route} | Data source proxy GET calls.
[**DatasourceProxyPOSTByUIDcalls**](DatasourcesApi.md#DatasourceProxyPOSTByUIDcalls) | **Post** /datasources/proxy/uid/{uid}/{datasource_proxy_route} | Data source proxy POST calls.
[**DatasourceProxyPOSTcalls**](DatasourcesApi.md#DatasourceProxyPOSTcalls) | **Post** /datasources/proxy/{id}/{datasource_proxy_route} | Data source proxy POST calls.
[**DeleteDatasourceByID**](DatasourcesApi.md#DeleteDatasourceByID) | **Delete** /datasources/{id} | Delete an existing data source by id.
[**DeleteDatasourceByName**](DatasourcesApi.md#DeleteDatasourceByName) | **Delete** /datasources/name/{name} | Delete an existing data source by name.
[**DeleteDatasourceByUID**](DatasourcesApi.md#DeleteDatasourceByUID) | **Delete** /datasources/uid/{uid} | Delete an existing data source by UID.
[**FetchDatasourceResources**](DatasourcesApi.md#FetchDatasourceResources) | **Get** /datasources/uid/{uid}/resources/{datasource_proxy_route} | Fetch data source resources.
[**FetchDatasourceResourcesByID**](DatasourcesApi.md#FetchDatasourceResourcesByID) | **Get** /datasources/{id}/resources/{datasource_proxy_route} | Fetch data source resources by Id.
[**GetDatasourceByID**](DatasourcesApi.md#GetDatasourceByID) | **Get** /datasources/{id} | Get a single data source by Id.
[**GetDatasourceByName**](DatasourcesApi.md#GetDatasourceByName) | **Get** /datasources/name/{name} | Get a single data source by Name.
[**GetDatasourceByUID**](DatasourcesApi.md#GetDatasourceByUID) | **Get** /datasources/uid/{uid} | Get a single data source by UID.
[**GetDatasourceIdByName**](DatasourcesApi.md#GetDatasourceIdByName) | **Get** /datasources/id/{name} | Get data source Id by Name.
[**GetDatasources**](DatasourcesApi.md#GetDatasources) | **Get** /datasources | Get all data sources.
[**UpdateDatasourceByID**](DatasourcesApi.md#UpdateDatasourceByID) | **Put** /datasources/{id} | Update an existing data source by its sequential ID.
[**UpdateDatasourceByUID**](DatasourcesApi.md#UpdateDatasourceByUID) | **Put** /datasources/uid/{uid} | Update an existing data source.


# **AddDatasource**
> InlineResponse2006Model AddDatasource(ctx, body)
Create a data source.

By defining `password` and `basicAuthPassword` under secureJsonData property Grafana encrypts them securely as an encrypted blob in the database. The response then lists the encrypted fields under secureJsonFields.  If you are running Grafana Enterprise and have Fine-grained access control enabled you need to have a permission with action: `datasources:create`

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AddDataSourceCommandModel**](AddDataSourceCommandModel.md)|  | 

### Return type

[**InlineResponse2006Model**](inline_response_200_6.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CheckDatasourceHealth**
> SuccessResponseBodyModel CheckDatasourceHealth(ctx, uid)
Check data source health by Id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uid** | **string**|  | 

### Return type

[**SuccessResponseBodyModel**](SuccessResponseBody.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CheckDatasourceHealthByID**
> SuccessResponseBodyModel CheckDatasourceHealthByID(ctx, id)
Check data source health by Id.

Please refer to [updated API](#/datasources/checkDatasourceHealth) instead

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**SuccessResponseBodyModel**](SuccessResponseBody.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DatasourceProxyDELETEByUIDcalls**
> DatasourceProxyDELETEByUIDcalls(ctx, uid, datasourceProxyRoute)
Data source proxy DELETE calls.

Proxies all calls to the actual data source.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uid** | **string**|  | 
  **datasourceProxyRoute** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DatasourceProxyDELETEcalls**
> DatasourceProxyDELETEcalls(ctx, id, datasourceProxyRoute)
Data source proxy DELETE calls.

Proxies all calls to the actual data source.  Please refer to [updated API](#/datasources/datasourceProxyDELETEByUIDcalls) instead

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **datasourceProxyRoute** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DatasourceProxyGETByUIDcalls**
> DatasourceProxyGETByUIDcalls(ctx, datasourceProxyRoute, uid)
Data source proxy GET calls.

Proxies all calls to the actual data source.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **datasourceProxyRoute** | **string**|  | 
  **uid** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DatasourceProxyGETcalls**
> DatasourceProxyGETcalls(ctx, datasourceProxyRoute, id)
Data source proxy GET calls.

Proxies all calls to the actual data source.  Please refer to [updated API](#/datasources/datasourceProxyGETByUIDcalls) instead

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **datasourceProxyRoute** | **string**|  | 
  **id** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DatasourceProxyPOSTByUIDcalls**
> DatasourceProxyPOSTByUIDcalls(ctx, datasourceProxyParam, datasourceProxyRoute, uid)
Data source proxy POST calls.

Proxies all calls to the actual data source. The data source should support POST methods for the specific path and role as defined

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **datasourceProxyParam** | **interface{}**|  | 
  **datasourceProxyRoute** | **string**|  | 
  **uid** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DatasourceProxyPOSTcalls**
> DatasourceProxyPOSTcalls(ctx, datasourceProxyParam, datasourceProxyRoute, id)
Data source proxy POST calls.

Proxies all calls to the actual data source. The data source should support POST methods for the specific path and role as defined  Please refer to [updated API](#/datasources/datasourceProxyPOSTByUIDcalls) instead

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **datasourceProxyParam** | **interface{}**|  | 
  **datasourceProxyRoute** | **string**|  | 
  **id** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDatasourceByID**
> SuccessResponseBodyModel DeleteDatasourceByID(ctx, id)
Delete an existing data source by id.

If you are running Grafana Enterprise and have Fine-grained access control enabled you need to have a permission with action: `datasources:delete` and scopes: `datasources:*`, `datasources:id:*` and `datasources:id:1` (single data source).  Please refer to [updated API](#/datasources/deleteDatasourceByUID) instead

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**SuccessResponseBodyModel**](SuccessResponseBody.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDatasourceByName**
> InlineResponse2008Model DeleteDatasourceByName(ctx, name)
Delete an existing data source by name.

If you are running Grafana Enterprise and have Fine-grained access control enabled you need to have a permission with action: `datasources:delete` and scopes: `datasources:*`, `datasources:name:*` and `datasources:name:test_datasource` (single data source).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 

### Return type

[**InlineResponse2008Model**](inline_response_200_8.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDatasourceByUID**
> SuccessResponseBodyModel DeleteDatasourceByUID(ctx, uid)
Delete an existing data source by UID.

If you are running Grafana Enterprise and have Fine-grained access control enabled you need to have a permission with action: `datasources:delete` and scopes: `datasources:*`, `datasources:uid:*` and `datasources:uid:kLtEtcRGk` (single data source).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uid** | **string**|  | 

### Return type

[**SuccessResponseBodyModel**](SuccessResponseBody.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FetchDatasourceResources**
> SuccessResponseBodyModel FetchDatasourceResources(ctx, uid, datasourceProxyRoute)
Fetch data source resources.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uid** | **string**|  | 
  **datasourceProxyRoute** | **string**|  | 

### Return type

[**SuccessResponseBodyModel**](SuccessResponseBody.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FetchDatasourceResourcesByID**
> SuccessResponseBodyModel FetchDatasourceResourcesByID(ctx, id, datasourceProxyRoute)
Fetch data source resources by Id.

Please refer to [updated API](#/datasources/fetchDatasourceResources) instead

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **datasourceProxyRoute** | **string**|  | 

### Return type

[**SuccessResponseBodyModel**](SuccessResponseBody.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDatasourceByID**
> DataSourceModel GetDatasourceByID(ctx, id)
Get a single data source by Id.

If you are running Grafana Enterprise and have Fine-grained access control enabled you need to have a permission with action: `datasources:read` and scopes: `datasources:*`, `datasources:id:*` and `datasources:id:1` (single data source).  Please refer to [updated API](#/datasources/getDatasourceByUID) instead

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**DataSourceModel**](DataSource.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDatasourceByName**
> DataSourceModel GetDatasourceByName(ctx, name)
Get a single data source by Name.

If you are running Grafana Enterprise and have Fine-grained access control enabled you need to have a permission with action: `datasources:read` and scopes: `datasources:*`, `datasources:name:*` and `datasources:name:test_datasource` (single data source).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 

### Return type

[**DataSourceModel**](DataSource.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDatasourceByUID**
> DataSourceModel GetDatasourceByUID(ctx, uid)
Get a single data source by UID.

If you are running Grafana Enterprise and have Fine-grained access control enabled you need to have a permission with action: `datasources:read` and scopes: `datasources:*`, `datasources:uid:*` and `datasources:uid:kLtEtcRGk` (single data source).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uid** | **string**|  | 

### Return type

[**DataSourceModel**](DataSource.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDatasourceIdByName**
> InlineResponse2007Model GetDatasourceIdByName(ctx, name)
Get data source Id by Name.

If you are running Grafana Enterprise and have Fine-grained access control enabled you need to have a permission with action: `datasources:read` and scopes: `datasources:*`, `datasources:name:*` and `datasources:name:test_datasource` (single data source).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 

### Return type

[**InlineResponse2007Model**](inline_response_200_7.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDatasources**
> DataSourceListModel GetDatasources(ctx, )
Get all data sources.

If you are running Grafana Enterprise and have Fine-grained access control enabled you need to have a permission with action: `datasources:read` and scope: `datasources:*`.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**DataSourceListModel**](DataSourceList.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDatasourceByID**
> InlineResponse2006Model UpdateDatasourceByID(ctx, body, id)
Update an existing data source by its sequential ID.

Similar to creating a data source, `password` and `basicAuthPassword` should be defined under secureJsonData in order to be stored securely as an encrypted blob in the database. Then, the encrypted fields are listed under secureJsonFields section in the response.  If you are running Grafana Enterprise and have Fine-grained access control enabled you need to have a permission with action: `datasources:write` and scopes: `datasources:*`, `datasources:id:*` and `datasources:id:1` (single data source).  Please refer to [updated API](#/datasources/updateDatasourceByUID) instead

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateDataSourceCommandModel**](UpdateDataSourceCommandModel.md)|  | 
  **id** | **string**|  | 

### Return type

[**InlineResponse2006Model**](inline_response_200_6.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDatasourceByUID**
> InlineResponse2006Model UpdateDatasourceByUID(ctx, body, uid)
Update an existing data source.

Similar to creating a data source, `password` and `basicAuthPassword` should be defined under secureJsonData in order to be stored securely as an encrypted blob in the database. Then, the encrypted fields are listed under secureJsonFields section in the response.  If you are running Grafana Enterprise and have Fine-grained access control enabled you need to have a permission with action: `datasources:write` and scopes: `datasources:*`, `datasources:uid:*` and `datasources:uid:1` (single data source).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateDataSourceCommandModel**](UpdateDataSourceCommandModel.md)|  | 
  **uid** | **string**|  | 

### Return type

[**InlineResponse2006Model**](inline_response_200_6.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

