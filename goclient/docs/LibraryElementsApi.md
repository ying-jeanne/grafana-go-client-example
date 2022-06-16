# \LibraryElementsApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLibraryElement**](LibraryElementsApi.md#CreateLibraryElement) | **Post** /library-elements | Create library element.
[**DeleteLibraryElementByUID**](LibraryElementsApi.md#DeleteLibraryElementByUID) | **Delete** /library-elements/{library_element_uid} | Delete library element.
[**GetLibraryElementByName**](LibraryElementsApi.md#GetLibraryElementByName) | **Get** /library-elements/name/{library_element_name} | Get library element by name.
[**GetLibraryElementByUID**](LibraryElementsApi.md#GetLibraryElementByUID) | **Get** /library-elements/{library_element_uid} | Get library element by UID.
[**GetLibraryElementConnections**](LibraryElementsApi.md#GetLibraryElementConnections) | **Get** /library-elements/{library_element_uid}/connections/ | Get library element connections.
[**GetLibraryElements**](LibraryElementsApi.md#GetLibraryElements) | **Get** /library-elements | Get all library elements.
[**UpdateLibraryElement**](LibraryElementsApi.md#UpdateLibraryElement) | **Patch** /library-elements/{library_element_uid} | Update library element.


# **CreateLibraryElement**
> LibraryElementResponseModel CreateLibraryElement(ctx, body)
Create library element.

Creates a new library element.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateLibraryElementCommandModel**](CreateLibraryElementCommandModel.md)|  | 

### Return type

[**LibraryElementResponseModel**](LibraryElementResponse.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteLibraryElementByUID**
> SuccessResponseBodyModel DeleteLibraryElementByUID(ctx, libraryElementUid)
Delete library element.

Deletes an existing library element as specified by the UID. This operation cannot be reverted. You cannot delete a library element that is connected. This operation cannot be reverted.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **libraryElementUid** | **string**|  | 

### Return type

[**SuccessResponseBodyModel**](SuccessResponseBody.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLibraryElementByName**
> LibraryElementResponseModel GetLibraryElementByName(ctx, libraryElementName)
Get library element by name.

Returns a library element with the given name.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **libraryElementName** | **string**|  | 

### Return type

[**LibraryElementResponseModel**](LibraryElementResponse.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLibraryElementByUID**
> LibraryElementResponseModel GetLibraryElementByUID(ctx, libraryElementUid)
Get library element by UID.

Returns a library element with the given UID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **libraryElementUid** | **string**|  | 

### Return type

[**LibraryElementResponseModel**](LibraryElementResponse.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLibraryElementConnections**
> LibraryElementConnectionsResponseModel GetLibraryElementConnections(ctx, libraryElementUid)
Get library element connections.

Returns a list of connections for a library element based on the UID specified.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **libraryElementUid** | **string**|  | 

### Return type

[**LibraryElementConnectionsResponseModel**](LibraryElementConnectionsResponse.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLibraryElements**
> LibraryElementSearchResponseModel GetLibraryElements(ctx, optional)
Get all library elements.

Returns a list of all library elements the authenticated user has permission to view. Use the `perPage` query parameter to control the maximum number of library elements returned; the default limit is `100`. You can also use the `page` query parameter to fetch library elements from any page other than the first one.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LibraryElementsApiGetLibraryElementsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LibraryElementsApiGetLibraryElementsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchString** | **optional.String**| Part of the name or description searched for. | 
 **kind** | **optional.Int64**| Kind of element to search for. | 
 **sortDirection** | **optional.String**| Sort order of elements. | 
 **typeFilter** | **optional.String**| A comma separated list of types to filter the elements by | 
 **excludeUid** | **optional.String**| Element UID to exclude from search results. | 
 **folderFilter** | **optional.String**| A comma separated list of folder ID(s) to filter the elements by. | 
 **perPage** | **optional.Int64**| The number of results per page. | [default to 100]
 **page** | **optional.Int64**| The page for a set of records, given that only perPage records are returned at a time. Numbering starts at 1. | [default to 1]

### Return type

[**LibraryElementSearchResponseModel**](LibraryElementSearchResponse.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateLibraryElement**
> LibraryElementResponseModel UpdateLibraryElement(ctx, body, libraryElementUid)
Update library element.

Updates an existing library element identified by uid.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PatchLibraryElementCommandModel**](PatchLibraryElementCommandModel.md)|  | 
  **libraryElementUid** | **string**|  | 

### Return type

[**LibraryElementResponseModel**](LibraryElementResponse.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

