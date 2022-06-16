# \ProvisioningApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RouteDeleteAlertRule**](ProvisioningApi.md#RouteDeleteAlertRule) | **Delete** /v1/provisioning/alert-rules/{UID} | Delete a specific alert rule by UID.
[**RouteDeleteContactpoints**](ProvisioningApi.md#RouteDeleteContactpoints) | **Delete** /v1/provisioning/contact-points/{UID} | Delete a contact point.
[**RouteDeleteMuteTiming**](ProvisioningApi.md#RouteDeleteMuteTiming) | **Delete** /v1/provisioning/mute-timings/{name} | Delete a mute timing.
[**RouteDeleteTemplate**](ProvisioningApi.md#RouteDeleteTemplate) | **Delete** /v1/provisioning/templates/{name} | Delete a template.
[**RouteGetAlertRule**](ProvisioningApi.md#RouteGetAlertRule) | **Get** /v1/provisioning/alert-rules/{UID} | Get a specific alert rule by UID.
[**RouteGetContactpoints**](ProvisioningApi.md#RouteGetContactpoints) | **Get** /v1/provisioning/contact-points | Get all the contact points.
[**RouteGetMuteTiming**](ProvisioningApi.md#RouteGetMuteTiming) | **Get** /v1/provisioning/mute-timings/{name} | Get a mute timing.
[**RouteGetMuteTimings**](ProvisioningApi.md#RouteGetMuteTimings) | **Get** /v1/provisioning/mute-timings | Get all the mute timings.
[**RouteGetPolicyTree**](ProvisioningApi.md#RouteGetPolicyTree) | **Get** /v1/provisioning/policies | Get the notification policy tree.
[**RouteGetTemplate**](ProvisioningApi.md#RouteGetTemplate) | **Get** /v1/provisioning/templates/{name} | Get a message template.
[**RouteGetTemplates**](ProvisioningApi.md#RouteGetTemplates) | **Get** /v1/provisioning/templates | Get all message templates.
[**RoutePostAlertRule**](ProvisioningApi.md#RoutePostAlertRule) | **Post** /v1/provisioning/alert-rules | Create a new alert rule.
[**RoutePostContactpoints**](ProvisioningApi.md#RoutePostContactpoints) | **Post** /v1/provisioning/contact-points | Create a contact point.
[**RoutePostMuteTiming**](ProvisioningApi.md#RoutePostMuteTiming) | **Post** /v1/provisioning/mute-timings | Create a new mute timing.
[**RoutePutAlertRule**](ProvisioningApi.md#RoutePutAlertRule) | **Put** /v1/provisioning/alert-rules/{UID} | Update an existing alert rule.
[**RoutePutAlertRuleGroup**](ProvisioningApi.md#RoutePutAlertRuleGroup) | **Put** /v1/provisioning/folder/{FolderUID}/rule-groups/{Group} | Update the interval of a rule group.
[**RoutePutContactpoint**](ProvisioningApi.md#RoutePutContactpoint) | **Put** /v1/provisioning/contact-points/{UID} | Update an existing contact point.
[**RoutePutMuteTiming**](ProvisioningApi.md#RoutePutMuteTiming) | **Put** /v1/provisioning/mute-timings/{name} | Replace an existing mute timing.
[**RoutePutPolicyTree**](ProvisioningApi.md#RoutePutPolicyTree) | **Put** /v1/provisioning/policies | Sets the notification policy tree.
[**RoutePutTemplate**](ProvisioningApi.md#RoutePutTemplate) | **Put** /v1/provisioning/templates/{name} | Updates an existing template.


# **RouteDeleteAlertRule**
> RouteDeleteAlertRule(ctx, uID)
Delete a specific alert rule by UID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uID** | **string**| Alert rule UID | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RouteDeleteContactpoints**
> RouteDeleteContactpoints(ctx, uID)
Delete a contact point.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uID** | **string**| UID is the contact point unique identifier | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RouteDeleteMuteTiming**
> RouteDeleteMuteTiming(ctx, name)
Delete a mute timing.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Mute timing name | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RouteDeleteTemplate**
> RouteDeleteTemplate(ctx, name)
Delete a template.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Template Name | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RouteGetAlertRule**
> AlertRuleModel RouteGetAlertRule(ctx, uID)
Get a specific alert rule by UID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uID** | **string**| Alert rule UID | 

### Return type

[**AlertRuleModel**](AlertRule.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RouteGetContactpoints**
> ContactPointsModel RouteGetContactpoints(ctx, )
Get all the contact points.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ContactPointsModel**](ContactPoints.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RouteGetMuteTiming**
> MuteTimeIntervalModel RouteGetMuteTiming(ctx, name)
Get a mute timing.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Mute timing name | 

### Return type

[**MuteTimeIntervalModel**](MuteTimeInterval.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RouteGetMuteTimings**
> MuteTimingsModel RouteGetMuteTimings(ctx, )
Get all the mute timings.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**MuteTimingsModel**](MuteTimings.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RouteGetPolicyTree**
> RouteModel RouteGetPolicyTree(ctx, )
Get the notification policy tree.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**RouteModel**](Route.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RouteGetTemplate**
> MessageTemplateModel RouteGetTemplate(ctx, name)
Get a message template.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Template Name | 

### Return type

[**MessageTemplateModel**](MessageTemplate.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RouteGetTemplates**
> MessageTemplatesModel RouteGetTemplates(ctx, )
Get all message templates.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**MessageTemplatesModel**](MessageTemplates.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RoutePostAlertRule**
> AlertRuleModel RoutePostAlertRule(ctx, optional)
Create a new alert rule.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProvisioningApiRoutePostAlertRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProvisioningApiRoutePostAlertRuleOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of AlertRuleModel**](AlertRuleModel.md)|  | 

### Return type

[**AlertRuleModel**](AlertRule.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RoutePostContactpoints**
> EmbeddedContactPointModel RoutePostContactpoints(ctx, optional)
Create a contact point.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProvisioningApiRoutePostContactpointsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProvisioningApiRoutePostContactpointsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of EmbeddedContactPointModel**](EmbeddedContactPointModel.md)|  | 

### Return type

[**EmbeddedContactPointModel**](EmbeddedContactPoint.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RoutePostMuteTiming**
> MuteTimeIntervalModel RoutePostMuteTiming(ctx, optional)
Create a new mute timing.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProvisioningApiRoutePostMuteTimingOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProvisioningApiRoutePostMuteTimingOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of MuteTimeIntervalModel**](MuteTimeIntervalModel.md)|  | 

### Return type

[**MuteTimeIntervalModel**](MuteTimeInterval.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RoutePutAlertRule**
> AlertRuleModel RoutePutAlertRule(ctx, uID, optional)
Update an existing alert rule.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uID** | **string**| Alert rule UID | 
 **optional** | ***ProvisioningApiRoutePutAlertRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProvisioningApiRoutePutAlertRuleOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of AlertRuleModel**](AlertRuleModel.md)|  | 

### Return type

[**AlertRuleModel**](AlertRule.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RoutePutAlertRuleGroup**
> AlertRuleGroupModel RoutePutAlertRuleGroup(ctx, folderUID, group, optional)
Update the interval of a rule group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **folderUID** | **string**|  | 
  **group** | **string**|  | 
 **optional** | ***ProvisioningApiRoutePutAlertRuleGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProvisioningApiRoutePutAlertRuleGroupOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of AlertRuleGroupModel**](AlertRuleGroupModel.md)|  | 

### Return type

[**AlertRuleGroupModel**](AlertRuleGroup.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RoutePutContactpoint**
> AckModel RoutePutContactpoint(ctx, uID, optional)
Update an existing contact point.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uID** | **string**| UID is the contact point unique identifier | 
 **optional** | ***ProvisioningApiRoutePutContactpointOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProvisioningApiRoutePutContactpointOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of EmbeddedContactPointModel**](EmbeddedContactPointModel.md)|  | 

### Return type

[**AckModel**](Ack.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RoutePutMuteTiming**
> MuteTimeIntervalModel RoutePutMuteTiming(ctx, name, optional)
Replace an existing mute timing.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Mute timing name | 
 **optional** | ***ProvisioningApiRoutePutMuteTimingOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProvisioningApiRoutePutMuteTimingOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of MuteTimeIntervalModel**](MuteTimeIntervalModel.md)|  | 

### Return type

[**MuteTimeIntervalModel**](MuteTimeInterval.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RoutePutPolicyTree**
> AckModel RoutePutPolicyTree(ctx, optional)
Sets the notification policy tree.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProvisioningApiRoutePutPolicyTreeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProvisioningApiRoutePutPolicyTreeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of RouteModel**](RouteModel.md)| The new notification routing tree to use | 

### Return type

[**AckModel**](Ack.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RoutePutTemplate**
> MessageTemplateModel RoutePutTemplate(ctx, name, optional)
Updates an existing template.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Template Name | 
 **optional** | ***ProvisioningApiRoutePutTemplateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProvisioningApiRoutePutTemplateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of MessageTemplateContentModel**](MessageTemplateContentModel.md)|  | 

### Return type

[**MessageTemplateModel**](MessageTemplate.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

