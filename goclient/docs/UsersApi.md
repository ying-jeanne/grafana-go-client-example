# \UsersApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUserByID**](UsersApi.md#GetUserByID) | **Get** /users/{user_id} | Get user by id.
[**GetUserByLoginOrEmail**](UsersApi.md#GetUserByLoginOrEmail) | **Get** /users/lookup | Get user by login or email.
[**GetUserOrgList**](UsersApi.md#GetUserOrgList) | **Get** /users/{user_id}/orgs | Get organizations for user.
[**GetUserTeams**](UsersApi.md#GetUserTeams) | **Get** /users/{user_id}/teams | Get teams for user.
[**SearchUsers**](UsersApi.md#SearchUsers) | **Get** /users | Get users.
[**SearchUsersWithPaging**](UsersApi.md#SearchUsersWithPaging) | **Get** /users/search | Get users with paging.
[**UpdateUser**](UsersApi.md#UpdateUser) | **Put** /users/{user_id} | Update user.


# **GetUserByID**
> UserProfileDtoModel GetUserByID(ctx, userId)
Get user by id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **int64**|  | 

### Return type

[**UserProfileDtoModel**](UserProfileDTO.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserByLoginOrEmail**
> UserProfileDtoModel GetUserByLoginOrEmail(ctx, loginOrEmail)
Get user by login or email.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **loginOrEmail** | **string**| loginOrEmail of the user | 

### Return type

[**UserProfileDtoModel**](UserProfileDTO.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserOrgList**
> []UserOrgDtoModel GetUserOrgList(ctx, userId)
Get organizations for user.

Get organizations for user identified by id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **int64**|  | 

### Return type

[**[]UserOrgDtoModel**](UserOrgDTO.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserTeams**
> []TeamDtoModel GetUserTeams(ctx, userId)
Get teams for user.

Get teams for user identified by id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **int64**|  | 

### Return type

[**[]TeamDtoModel**](TeamDTO.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchUsers**
> SearchUserQueryResultModel SearchUsers(ctx, optional)
Get users.

Returns all users that the authenticated user has permission to view, admin permission required.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UsersApiSearchUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiSearchUsersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **perpage** | **optional.Int64**| Limit the maximum number of users to return per page | [default to 1000]
 **page** | **optional.Int64**| Page index for starting fetching users | [default to 1]

### Return type

[**SearchUserQueryResultModel**](SearchUserQueryResult.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchUsersWithPaging**
> SearchUserQueryResultModel SearchUsersWithPaging(ctx, )
Get users with paging.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SearchUserQueryResultModel**](SearchUserQueryResult.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateUser**
> UserProfileDtoModel UpdateUser(ctx, body, userId)
Update user.

Update the user identified by id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateUserCommandModel**](UpdateUserCommandModel.md)| To change the email, name, login, theme, provide another one. | 
  **userId** | **int64**|  | 

### Return type

[**UserProfileDtoModel**](UserProfileDTO.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

