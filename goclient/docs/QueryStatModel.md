# QueryStatModel

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Color** | **map[string]interface{}** | Map values to a display color NOTE: this interface is under development in the frontend... so simple map for now | [optional] [default to null]
**Custom** | **map[string]interface{}** | Panel Specific Values | [optional] [default to null]
**Decimals** | **int32** |  | [optional] [default to null]
**Description** | **string** | Description is human readable field metadata | [optional] [default to null]
**DisplayName** | **string** | DisplayName overrides Grafana default naming, should not be used from a data source | [optional] [default to null]
**DisplayNameFromDS** | **string** | DisplayNameFromDS overrides Grafana default naming in a better way that allows users to override it easily. | [optional] [default to null]
**Filterable** | **bool** | Filterable indicates if the Field&#39;s data can be filtered by additional calls. | [optional] [default to null]
**Interval** | **float64** | Interval indicates the expected regular step between values in the series. When an interval exists, consumers can identify \&quot;missing\&quot; values when the expected value is not present. The grafana timeseries visualization will render disconnected values when missing values are found it the time field. The interval uses the same units as the values.  For time.Time, this is defined in milliseconds. | [optional] [default to null]
**Links** | [**[]DataLinkModel**](DataLink.md) | The behavior when clicking on a result | [optional] [default to null]
**Mappings** | [***ValueMappingsModel**](ValueMappings.md) |  | [optional] [default to null]
**Max** | [***ConfFloat64Model**](ConfFloat64.md) |  | [optional] [default to null]
**Min** | [***ConfFloat64Model**](ConfFloat64.md) |  | [optional] [default to null]
**NoValue** | **string** | Alternative to empty string | [optional] [default to null]
**Path** | **string** | Path is an explicit path to the field in the datasource. When the frame meta includes a path, this will default to &#x60;${frame.meta.path}/${field.name}  When defined, this value can be used as an identifier within the datasource scope, and may be used as an identifier to update values in a subsequent request | [optional] [default to null]
**Thresholds** | [***ThresholdsConfigModel**](ThresholdsConfig.md) |  | [optional] [default to null]
**Unit** | **string** | Numeric Options | [optional] [default to null]
**Value** | **float64** |  | [optional] [default to null]
**Writeable** | **bool** | Writeable indicates that the datasource knows how to update this value | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


