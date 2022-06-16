# \OrgInvitesApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddInvite**](OrgInvitesApi.md#AddInvite) | **Post** /org/invites | Add invite.
[**GetInvites**](OrgInvitesApi.md#GetInvites) | **Get** /org/invites | Get pending invites.
[**RevokeInvite**](OrgInvitesApi.md#RevokeInvite) | **Delete** /org/{invitation_code}/invites | Revoke invite.


# **AddInvite**
> InlineResponse20010Model AddInvite(ctx, body)
Add invite.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AddInviteFormModel**](AddInviteFormModel.md)|  | 

### Return type

[**InlineResponse20010Model**](inline_response_200_10.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetInvites**
> []TempUserDtoModel GetInvites(ctx, )
Get pending invites.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]TempUserDtoModel**](TempUserDTO.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RevokeInvite**
> SuccessResponseBodyModel RevokeInvite(ctx, invitationCode)
Revoke invite.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **invitationCode** | **string**|  | 

### Return type

[**SuccessResponseBodyModel**](SuccessResponseBody.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

