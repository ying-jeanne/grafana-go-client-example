# \LegacyAlertsNotificationChannelsApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAlertNotificationChannel**](LegacyAlertsNotificationChannelsApi.md#CreateAlertNotificationChannel) | **Post** /alert-notifications | Create notification channel.
[**DeleteAlertNotificationChannel**](LegacyAlertsNotificationChannelsApi.md#DeleteAlertNotificationChannel) | **Delete** /alert-notifications/{notification_channel_id} | Delete alert notification by ID.
[**DeleteAlertNotificationChannelByUID**](LegacyAlertsNotificationChannelsApi.md#DeleteAlertNotificationChannelByUID) | **Delete** /alert-notifications/uid/{notification_channel_uid} | Delete alert notification by UID.
[**GetAlertNotificationChannelByID**](LegacyAlertsNotificationChannelsApi.md#GetAlertNotificationChannelByID) | **Get** /alert-notifications/{notification_channel_id} | Get notification channel by ID.
[**GetAlertNotificationChannelByUID**](LegacyAlertsNotificationChannelsApi.md#GetAlertNotificationChannelByUID) | **Get** /alert-notifications/uid/{notification_channel_uid} | Get notification channel by UID
[**GetAlertNotificationChannels**](LegacyAlertsNotificationChannelsApi.md#GetAlertNotificationChannels) | **Get** /alert-notifications | Get all notification channels.
[**LookupAlertNotificationChannels**](LegacyAlertsNotificationChannelsApi.md#LookupAlertNotificationChannels) | **Get** /alert-notifications/lookup | Get all notification channels (lookup)
[**NotificationChannelTest**](LegacyAlertsNotificationChannelsApi.md#NotificationChannelTest) | **Post** /alert-notifications/test | Test notification channel.
[**UpdateAlertNotificationChannel**](LegacyAlertsNotificationChannelsApi.md#UpdateAlertNotificationChannel) | **Put** /alert-notifications/{notification_channel_id} | Update notification channel by ID.
[**UpdateAlertNotificationChannelByUID**](LegacyAlertsNotificationChannelsApi.md#UpdateAlertNotificationChannelByUID) | **Put** /alert-notifications/uid/{notification_channel_uid} | Update notification channel by UID.


# **CreateAlertNotificationChannel**
> AlertNotificationModel CreateAlertNotificationChannel(ctx, body)
Create notification channel.

You can find the full list of [supported notifiers](https://grafana.com/docs/grafana/latest/alerting/old-alerting/notifications/#list-of-supported-notifiers) on the alert notifiers page.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateAlertNotificationCommandModel**](CreateAlertNotificationCommandModel.md)|  | 

### Return type

[**AlertNotificationModel**](AlertNotification.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAlertNotificationChannel**
> SuccessResponseBodyModel DeleteAlertNotificationChannel(ctx, notificationChannelId)
Delete alert notification by ID.

Deletes an existing notification channel identified by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **notificationChannelId** | **int64**|  | 

### Return type

[**SuccessResponseBodyModel**](SuccessResponseBody.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAlertNotificationChannelByUID**
> InlineResponse2001Model DeleteAlertNotificationChannelByUID(ctx, notificationChannelUid)
Delete alert notification by UID.

Deletes an existing notification channel identified by UID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **notificationChannelUid** | **string**|  | 

### Return type

[**InlineResponse2001Model**](inline_response_200_1.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAlertNotificationChannelByID**
> AlertNotificationModel GetAlertNotificationChannelByID(ctx, notificationChannelId)
Get notification channel by ID.

Returns the notification channel given the notification channel ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **notificationChannelId** | **int64**|  | 

### Return type

[**AlertNotificationModel**](AlertNotification.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAlertNotificationChannelByUID**
> AlertNotificationModel GetAlertNotificationChannelByUID(ctx, notificationChannelUid)
Get notification channel by UID

Returns the notification channel given the notification channel UID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **notificationChannelUid** | **string**|  | 

### Return type

[**AlertNotificationModel**](AlertNotification.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAlertNotificationChannels**
> []AlertNotificationModel GetAlertNotificationChannels(ctx, )
Get all notification channels.

Returns all notification channels that the authenticated user has permission to view.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]AlertNotificationModel**](AlertNotification.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LookupAlertNotificationChannels**
> []AlertNotificationLookupModel LookupAlertNotificationChannels(ctx, )
Get all notification channels (lookup)

Returns all notification channels, but with less detailed information. Accessible by any authenticated user and is mainly used by providing alert notification channels in Grafana UI when configuring alert rule.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]AlertNotificationLookupModel**](AlertNotificationLookup.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NotificationChannelTest**
> SuccessResponseBodyModel NotificationChannelTest(ctx, body)
Test notification channel.

Sends a test notification to the channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**NotificationTestCommandModel**](NotificationTestCommandModel.md)|  | 

### Return type

[**SuccessResponseBodyModel**](SuccessResponseBody.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAlertNotificationChannel**
> AlertNotificationModel UpdateAlertNotificationChannel(ctx, body, notificationChannelId)
Update notification channel by ID.

Updates an existing notification channel identified by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateAlertNotificationCommandModel**](UpdateAlertNotificationCommandModel.md)|  | 
  **notificationChannelId** | **int64**|  | 

### Return type

[**AlertNotificationModel**](AlertNotification.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAlertNotificationChannelByUID**
> AlertNotificationModel UpdateAlertNotificationChannelByUID(ctx, body, notificationChannelUid)
Update notification channel by UID.

Updates an existing notification channel identified by uid.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateAlertNotificationWithUidCommandModel**](UpdateAlertNotificationWithUidCommandModel.md)|  | 
  **notificationChannelUid** | **string**|  | 

### Return type

[**AlertNotificationModel**](AlertNotification.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

