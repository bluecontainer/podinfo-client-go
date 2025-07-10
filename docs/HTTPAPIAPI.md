# \HTTPAPIAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiEchoPost**](HTTPAPIAPI.md#ApiEchoPost) | **Post** /api/echo | Echo
[**ApiInfoGet**](HTTPAPIAPI.md#ApiInfoGet) | **Get** /api/info | Runtime information
[**CacheKeyDelete**](HTTPAPIAPI.md#CacheKeyDelete) | **Delete** /cache/{key} | Delete payload from cache
[**CacheKeyGet**](HTTPAPIAPI.md#CacheKeyGet) | **Get** /cache/{key} | Get payload from cache
[**CacheKeyPost**](HTTPAPIAPI.md#CacheKeyPost) | **Post** /cache/{key} | Save payload in cache
[**ChunkedSecondsGet**](HTTPAPIAPI.md#ChunkedSecondsGet) | **Get** /chunked/{seconds} | Chunked transfer encoding
[**DelaySecondsGet**](HTTPAPIAPI.md#DelaySecondsGet) | **Get** /delay/{seconds} | Delay
[**EnvGet**](HTTPAPIAPI.md#EnvGet) | **Get** /env | Environment
[**HeadersGet**](HTTPAPIAPI.md#HeadersGet) | **Get** /headers | Headers
[**PanicGet**](HTTPAPIAPI.md#PanicGet) | **Get** /panic | Panic
[**RootGet**](HTTPAPIAPI.md#RootGet) | **Get** / | Index
[**StatusCodeGet**](HTTPAPIAPI.md#StatusCodeGet) | **Get** /status/{code} | Status code
[**StoreHashGet**](HTTPAPIAPI.md#StoreHashGet) | **Get** /store/{hash} | Download file
[**StorePost**](HTTPAPIAPI.md#StorePost) | **Post** /store | Upload file
[**TokenPost**](HTTPAPIAPI.md#TokenPost) | **Post** /token | Generate JWT token
[**TokenValidatePost**](HTTPAPIAPI.md#TokenValidatePost) | **Post** /token/validate | Validate JWT token
[**VersionGet**](HTTPAPIAPI.md#VersionGet) | **Get** /version | Version
[**WsEchoPost**](HTTPAPIAPI.md#WsEchoPost) | **Post** /ws/echo | Echo over websockets



## ApiEchoPost

> string ApiEchoPost(ctx).Value(value).Execute()

Echo



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/bluecontainer/podinfo-client-go"
)

func main() {
	value := "value_example" // string | The text to echo.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HTTPAPIAPI.ApiEchoPost(context.Background()).Value(value).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HTTPAPIAPI.ApiEchoPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiEchoPost`: string
	fmt.Fprintf(os.Stdout, "Response from `HTTPAPIAPI.ApiEchoPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiEchoPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **value** | **string** | The text to echo. | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiInfoGet

> HttpRuntimeResponse ApiInfoGet(ctx).Execute()

Runtime information



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/bluecontainer/podinfo-client-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HTTPAPIAPI.ApiInfoGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HTTPAPIAPI.ApiInfoGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiInfoGet`: HttpRuntimeResponse
	fmt.Fprintf(os.Stdout, "Response from `HTTPAPIAPI.ApiInfoGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiInfoGetRequest struct via the builder pattern


### Return type

[**HttpRuntimeResponse**](HttpRuntimeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CacheKeyDelete

> CacheKeyDelete(ctx, key).Execute()

Delete payload from cache



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/bluecontainer/podinfo-client-go"
)

func main() {
	key := "key_example" // string | Key to delete

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.HTTPAPIAPI.CacheKeyDelete(context.Background(), key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HTTPAPIAPI.CacheKeyDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | Key to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiCacheKeyDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CacheKeyGet

> string CacheKeyGet(ctx, key).Execute()

Get payload from cache



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/bluecontainer/podinfo-client-go"
)

func main() {
	key := "key_example" // string | Key to load from cache

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HTTPAPIAPI.CacheKeyGet(context.Background(), key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HTTPAPIAPI.CacheKeyGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CacheKeyGet`: string
	fmt.Fprintf(os.Stdout, "Response from `HTTPAPIAPI.CacheKeyGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | Key to load from cache | 

### Other Parameters

Other parameters are passed through a pointer to a apiCacheKeyGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CacheKeyPost

> CacheKeyPost(ctx, key).Value(value).Execute()

Save payload in cache



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/bluecontainer/podinfo-client-go"
)

func main() {
	key := "key_example" // string | Key to save to
	value := "value_example" // string | The value to cache.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.HTTPAPIAPI.CacheKeyPost(context.Background(), key).Value(value).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HTTPAPIAPI.CacheKeyPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | Key to save to | 

### Other Parameters

Other parameters are passed through a pointer to a apiCacheKeyPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **value** | **string** | The value to cache. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChunkedSecondsGet

> map[string]string ChunkedSecondsGet(ctx, seconds).Execute()

Chunked transfer encoding



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/bluecontainer/podinfo-client-go"
)

func main() {
	seconds := int32(56) // int32 | seconds to wait for

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HTTPAPIAPI.ChunkedSecondsGet(context.Background(), seconds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HTTPAPIAPI.ChunkedSecondsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChunkedSecondsGet`: map[string]string
	fmt.Fprintf(os.Stdout, "Response from `HTTPAPIAPI.ChunkedSecondsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**seconds** | **int32** | seconds to wait for | 

### Other Parameters

Other parameters are passed through a pointer to a apiChunkedSecondsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DelaySecondsGet

> map[string]string DelaySecondsGet(ctx, seconds).Execute()

Delay



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/bluecontainer/podinfo-client-go"
)

func main() {
	seconds := int32(56) // int32 | seconds to wait for

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HTTPAPIAPI.DelaySecondsGet(context.Background(), seconds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HTTPAPIAPI.DelaySecondsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DelaySecondsGet`: map[string]string
	fmt.Fprintf(os.Stdout, "Response from `HTTPAPIAPI.DelaySecondsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**seconds** | **int32** | seconds to wait for | 

### Other Parameters

Other parameters are passed through a pointer to a apiDelaySecondsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnvGet

> []string EnvGet(ctx).Execute()

Environment



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/bluecontainer/podinfo-client-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HTTPAPIAPI.EnvGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HTTPAPIAPI.EnvGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EnvGet`: []string
	fmt.Fprintf(os.Stdout, "Response from `HTTPAPIAPI.EnvGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiEnvGetRequest struct via the builder pattern


### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HeadersGet

> []string HeadersGet(ctx).Execute()

Headers



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/bluecontainer/podinfo-client-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HTTPAPIAPI.HeadersGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HTTPAPIAPI.HeadersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HeadersGet`: []string
	fmt.Fprintf(os.Stdout, "Response from `HTTPAPIAPI.HeadersGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiHeadersGetRequest struct via the builder pattern


### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PanicGet

> PanicGet(ctx).Execute()

Panic



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/bluecontainer/podinfo-client-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.HTTPAPIAPI.PanicGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HTTPAPIAPI.PanicGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPanicGetRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RootGet

> string RootGet(ctx).Execute()

Index



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/bluecontainer/podinfo-client-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HTTPAPIAPI.RootGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HTTPAPIAPI.RootGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RootGet`: string
	fmt.Fprintf(os.Stdout, "Response from `HTTPAPIAPI.RootGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRootGetRequest struct via the builder pattern


### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StatusCodeGet

> map[string]string StatusCodeGet(ctx, code).Execute()

Status code



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/bluecontainer/podinfo-client-go"
)

func main() {
	code := int32(56) // int32 | status code to return

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HTTPAPIAPI.StatusCodeGet(context.Background(), code).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HTTPAPIAPI.StatusCodeGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StatusCodeGet`: map[string]string
	fmt.Fprintf(os.Stdout, "Response from `HTTPAPIAPI.StatusCodeGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **int32** | status code to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiStatusCodeGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StoreHashGet

> string StoreHashGet(ctx, hash).Execute()

Download file



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/bluecontainer/podinfo-client-go"
)

func main() {
	hash := "hash_example" // string | hash value

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HTTPAPIAPI.StoreHashGet(context.Background(), hash).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HTTPAPIAPI.StoreHashGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StoreHashGet`: string
	fmt.Fprintf(os.Stdout, "Response from `HTTPAPIAPI.StoreHashGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hash** | **string** | hash value | 

### Other Parameters

Other parameters are passed through a pointer to a apiStoreHashGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorePost

> map[string]string StorePost(ctx).Value(value).Execute()

Upload file



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/bluecontainer/podinfo-client-go"
)

func main() {
	value := "value_example" // string | The content to store.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HTTPAPIAPI.StorePost(context.Background()).Value(value).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HTTPAPIAPI.StorePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StorePost`: map[string]string
	fmt.Fprintf(os.Stdout, "Response from `HTTPAPIAPI.StorePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStorePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **value** | **string** | The content to store. | 

### Return type

**map[string]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TokenPost

> HttpTokenResponse TokenPost(ctx).Execute()

Generate JWT token



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/bluecontainer/podinfo-client-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HTTPAPIAPI.TokenPost(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HTTPAPIAPI.TokenPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TokenPost`: HttpTokenResponse
	fmt.Fprintf(os.Stdout, "Response from `HTTPAPIAPI.TokenPost`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTokenPostRequest struct via the builder pattern


### Return type

[**HttpTokenResponse**](HttpTokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TokenValidatePost

> HttpTokenValidationResponse TokenValidatePost(ctx).Execute()

Validate JWT token



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/bluecontainer/podinfo-client-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HTTPAPIAPI.TokenValidatePost(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HTTPAPIAPI.TokenValidatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TokenValidatePost`: HttpTokenValidationResponse
	fmt.Fprintf(os.Stdout, "Response from `HTTPAPIAPI.TokenValidatePost`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTokenValidatePostRequest struct via the builder pattern


### Return type

[**HttpTokenValidationResponse**](HttpTokenValidationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VersionGet

> map[string]string VersionGet(ctx).Execute()

Version



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/bluecontainer/podinfo-client-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HTTPAPIAPI.VersionGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HTTPAPIAPI.VersionGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VersionGet`: map[string]string
	fmt.Fprintf(os.Stdout, "Response from `HTTPAPIAPI.VersionGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiVersionGetRequest struct via the builder pattern


### Return type

**map[string]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WsEchoPost

> map[string]string WsEchoPost(ctx).Execute()

Echo over websockets



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/bluecontainer/podinfo-client-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HTTPAPIAPI.WsEchoPost(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HTTPAPIAPI.WsEchoPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WsEchoPost`: map[string]string
	fmt.Fprintf(os.Stdout, "Response from `HTTPAPIAPI.WsEchoPost`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiWsEchoPostRequest struct via the builder pattern


### Return type

**map[string]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

