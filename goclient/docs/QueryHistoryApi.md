# \QueryHistoryApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateQuery**](QueryHistoryApi.md#CreateQuery) | **Post** /query-history | Add query to query history.
[**DeleteQuery**](QueryHistoryApi.md#DeleteQuery) | **Delete** /query-history/{query_history_uid} | Delete query in query history.
[**MigrateQueries**](QueryHistoryApi.md#MigrateQueries) | **Post** /query-history/migrate | Migrate queries to query history.
[**PatchQueryComment**](QueryHistoryApi.md#PatchQueryComment) | **Patch** /query-history/{query_history_uid} | Update comment for query in query history.
[**SearchQueries**](QueryHistoryApi.md#SearchQueries) | **Get** /query-history | Query history search.
[**StarQuery**](QueryHistoryApi.md#StarQuery) | **Post** /query-history/star/{query_history_uid} | Add star to query in query history.
[**UnstarQuery**](QueryHistoryApi.md#UnstarQuery) | **Delete** /query-history/star/{query_history_uid} | Remove star to query in query history.


# **CreateQuery**
> QueryHistoryResponseModel CreateQuery(ctx, body)
Add query to query history.

Adds new query to query history.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateQueryInQueryHistoryCommandModel**](CreateQueryInQueryHistoryCommandModel.md)|  | 

### Return type

[**QueryHistoryResponseModel**](QueryHistoryResponse.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteQuery**
> QueryHistoryDeleteQueryResponseModel DeleteQuery(ctx, queryHistoryUid)
Delete query in query history.

Deletes an existing query in query history as specified by the UID. This operation cannot be reverted.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **queryHistoryUid** | **string**|  | 

### Return type

[**QueryHistoryDeleteQueryResponseModel**](QueryHistoryDeleteQueryResponse.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MigrateQueries**
> QueryHistoryMigrationResponseModel MigrateQueries(ctx, body)
Migrate queries to query history.

Adds multiple queries to query history.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MigrateQueriesToQueryHistoryCommandModel**](MigrateQueriesToQueryHistoryCommandModel.md)|  | 

### Return type

[**QueryHistoryMigrationResponseModel**](QueryHistoryMigrationResponse.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchQueryComment**
> QueryHistoryResponseModel PatchQueryComment(ctx, queryHistoryUid, body)
Update comment for query in query history.

Updates comment for query in query history as specified by the UID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **queryHistoryUid** | **string**|  | 
  **body** | [**PatchQueryCommentInQueryHistoryCommandModel**](PatchQueryCommentInQueryHistoryCommandModel.md)|  | 

### Return type

[**QueryHistoryResponseModel**](QueryHistoryResponse.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchQueries**
> QueryHistorySearchResponseModel SearchQueries(ctx, optional)
Query history search.

Returns a list of queries in the query history that matches the search criteria. Query history search supports pagination. Use the `limit` parameter to control the maximum number of queries returned; the default limit is 100. You can also use the `page` query parameter to fetch queries from any page other than the first one.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***QueryHistoryApiSearchQueriesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a QueryHistoryApiSearchQueriesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **datasourceUid** | [**optional.Interface of []string**](string.md)| List of data source UIDs to search for | 
 **searchString** | **optional.String**| Text inside query or comments that is searched for | 
 **onlyStarred** | **optional.Bool**| Flag indicating if only starred queries should be returned | 
 **sort** | **optional.String**| Sort method | [default to time-desc]
 **page** | **optional.Int64**| Use this parameter to access hits beyond limit. Numbering starts at 1. limit param acts as page size. | 
 **limit** | **optional.Int64**| Limit the number of returned results | 
 **from** | **optional.Int64**| From range for the query history search | 
 **to** | **optional.Int64**| To range for the query history search | 

### Return type

[**QueryHistorySearchResponseModel**](QueryHistorySearchResponse.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StarQuery**
> QueryHistoryResponseModel StarQuery(ctx, queryHistoryUid)
Add star to query in query history.

Adds star to query in query history as specified by the UID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **queryHistoryUid** | **string**|  | 

### Return type

[**QueryHistoryResponseModel**](QueryHistoryResponse.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UnstarQuery**
> QueryHistoryResponseModel UnstarQuery(ctx, queryHistoryUid)
Remove star to query in query history.

Removes star from query in query history as specified by the UID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **queryHistoryUid** | **string**|  | 

### Return type

[**QueryHistoryResponseModel**](QueryHistoryResponse.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

