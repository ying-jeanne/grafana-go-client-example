# \ReportsApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateReport**](ReportsApi.md#CreateReport) | **Post** /reports | Create a report.
[**DeleteReport**](ReportsApi.md#DeleteReport) | **Delete** /reports/{reportID} | Delete a report.
[**GetReport**](ReportsApi.md#GetReport) | **Get** /reports/{reportID} | Get a report.
[**GetReportSettings**](ReportsApi.md#GetReportSettings) | **Get** /reports/settings | Get settings.
[**GetReports**](ReportsApi.md#GetReports) | **Get** /reports | List reports.
[**RenderReportPDF**](ReportsApi.md#RenderReportPDF) | **Get** /reports/render/pdf/{DashboardID} | Render report for dashboard.
[**RenderReportPDFs**](ReportsApi.md#RenderReportPDFs) | **Get** /reports/render/pdfs | Render report for multiple dashboards.
[**SaveReportSettings**](ReportsApi.md#SaveReportSettings) | **Post** /reports/settings | Save settings.
[**SendReport**](ReportsApi.md#SendReport) | **Post** /reports/email | Send a report.
[**SendTestEmail**](ReportsApi.md#SendTestEmail) | **Post** /reports/test-email | Send test report via email.
[**UpdateReport**](ReportsApi.md#UpdateReport) | **Put** /reports/{reportID} | Update a report.


# **CreateReport**
> interface{} CreateReport(ctx, body)
Create a report.

Available to org admins only and with a valid license.  You need to have a permission with action `reports.admin:create`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateOrUpdateConfigCmdModel**](CreateOrUpdateConfigCmdModel.md)|  | 

### Return type

**interface{}**

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteReport**
> SuccessResponseBodyModel DeleteReport(ctx, reportID)
Delete a report.

Available to org admins only and with a valid or expired license  You need to have a permission with action `reports.delete` with scope `reports:id:<report ID>`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reportID** | **int64**|  | 

### Return type

[**SuccessResponseBodyModel**](SuccessResponseBody.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReport**
> ConfigDtoModel GetReport(ctx, reportID)
Get a report.

Available to org admins only and with a valid or expired license  You need to have a permission with action `reports:read` with scope `reports:id:<report ID>`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reportID** | **int64**|  | 

### Return type

[**ConfigDtoModel**](ConfigDTO.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReportSettings**
> SettingsDtoModel GetReportSettings(ctx, )
Get settings.

Available to org admins only and with a valid or expired license  You need to have a permission with action `reports.settings:read`x.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SettingsDtoModel**](SettingsDTO.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReports**
> []ConfigDtoModel GetReports(ctx, )
List reports.

Available to org admins only and with a valid or expired license  You need to have a permission with action `reports:read` with scope `reports:*`.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]ConfigDtoModel**](ConfigDTO.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RenderReportPDF**
> []int32 RenderReportPDF(ctx, dashboardID)
Render report for dashboard.

Please refer to [reports enterprise](#/reports/renderReportPDFs) instead. This will be removed in Grafana 10.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dashboardID** | **int64**|  | 

### Return type

**[]int32**

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/pdf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RenderReportPDFs**
> []int32 RenderReportPDFs(ctx, )
Render report for multiple dashboards.

Available to all users and with a valid license.

### Required Parameters
This endpoint does not need any parameter.

### Return type

**[]int32**

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/pdf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SaveReportSettings**
> SuccessResponseBodyModel SaveReportSettings(ctx, body)
Save settings.

Available to org admins only and with a valid or expired license  You need to have a permission with action `reports.settings:write`xx.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SettingsDtoModel**](SettingsDtoModel.md)|  | 

### Return type

[**SuccessResponseBodyModel**](SuccessResponseBody.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SendReport**
> SuccessResponseBodyModel SendReport(ctx, body)
Send a report.

Generate and send a report. This API waits for the report to be generated before returning. We recommend that you set the client’s timeout to at least 60 seconds. Available to org admins only and with a valid license.  Only available in Grafana Enterprise v7.0+. This API endpoint is experimental and may be deprecated in a future release. On deprecation, a migration strategy will be provided and the endpoint will remain functional until the next major release of Grafana.  You need to have a permission with action `reports:send`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ReportEmailDtoModel**](ReportEmailDtoModel.md)|  | 

### Return type

[**SuccessResponseBodyModel**](SuccessResponseBody.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SendTestEmail**
> SuccessResponseBodyModel SendTestEmail(ctx, body)
Send test report via email.

Available to org admins only and with a valid license.  You need to have a permission with action `reports:send`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateOrUpdateConfigCmdModel**](CreateOrUpdateConfigCmdModel.md)|  | 

### Return type

[**SuccessResponseBodyModel**](SuccessResponseBody.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateReport**
> SuccessResponseBodyModel UpdateReport(ctx, body, reportID)
Update a report.

Available to org admins only and with a valid or expired license  You need to have a permission with action `reports.admin:write` with scope `reports:id:<report ID>`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateOrUpdateConfigCmdModel**](CreateOrUpdateConfigCmdModel.md)|  | 
  **reportID** | **int64**|  | 

### Return type

[**SuccessResponseBodyModel**](SuccessResponseBody.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

