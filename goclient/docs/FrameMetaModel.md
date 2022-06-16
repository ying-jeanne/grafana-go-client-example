# FrameMetaModel

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | **string** | Channel is the path to a stream in grafana live that has real-time updates for this data. | [optional] [default to null]
**Custom** | **interface{}** | Custom datasource specific values. | [optional] [default to null]
**ExecutedQueryString** | **string** | ExecutedQueryString is the raw query sent to the underlying system. All macros and templating have been applied.  When metadata contains this value, it will be shown in the query inspector. | [optional] [default to null]
**Notices** | [**[]NoticeModel**](Notice.md) | Notices provide additional information about the data in the Frame that Grafana can display to the user in the user interface. | [optional] [default to null]
**Path** | **string** | Path is a browsable path on the datasource. | [optional] [default to null]
**PathSeparator** | **string** | PathSeparator defines the separator pattern to decode a hierarchy. The default separator is &#39;/&#39;. | [optional] [default to null]
**PreferredVisualisationType** | [***VisTypeModel**](VisType.md) |  | [optional] [default to null]
**Stats** | [**[]QueryStatModel**](QueryStat.md) | Stats is an array of query result statistics. | [optional] [default to null]
**Type_** | [***FrameTypeModel**](FrameType.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


