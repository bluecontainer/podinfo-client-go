# HttpTokenValidationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiresAt** | Pointer to **string** |  | [optional] 
**TokenName** | Pointer to **string** |  | [optional] 

## Methods

### NewHttpTokenValidationResponse

`func NewHttpTokenValidationResponse() *HttpTokenValidationResponse`

NewHttpTokenValidationResponse instantiates a new HttpTokenValidationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHttpTokenValidationResponseWithDefaults

`func NewHttpTokenValidationResponseWithDefaults() *HttpTokenValidationResponse`

NewHttpTokenValidationResponseWithDefaults instantiates a new HttpTokenValidationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiresAt

`func (o *HttpTokenValidationResponse) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *HttpTokenValidationResponse) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *HttpTokenValidationResponse) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *HttpTokenValidationResponse) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetTokenName

`func (o *HttpTokenValidationResponse) GetTokenName() string`

GetTokenName returns the TokenName field if non-nil, zero value otherwise.

### GetTokenNameOk

`func (o *HttpTokenValidationResponse) GetTokenNameOk() (*string, bool)`

GetTokenNameOk returns a tuple with the TokenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenName

`func (o *HttpTokenValidationResponse) SetTokenName(v string)`

SetTokenName sets TokenName field to given value.

### HasTokenName

`func (o *HttpTokenValidationResponse) HasTokenName() bool`

HasTokenName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


