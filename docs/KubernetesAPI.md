# \KubernetesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HealthzGet**](KubernetesAPI.md#HealthzGet) | **Get** /healthz | Liveness check
[**MetricsGet**](KubernetesAPI.md#MetricsGet) | **Get** /metrics | Prometheus metrics
[**ReadyzDisablePost**](KubernetesAPI.md#ReadyzDisablePost) | **Post** /readyz/disable | Disable ready state
[**ReadyzEnablePost**](KubernetesAPI.md#ReadyzEnablePost) | **Post** /readyz/enable | Enable ready state
[**ReadyzGet**](KubernetesAPI.md#ReadyzGet) | **Get** /readyz | Readiness check



## HealthzGet

> string HealthzGet(ctx).Execute()

Liveness check



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
	resp, r, err := apiClient.KubernetesAPI.HealthzGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.HealthzGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HealthzGet`: string
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.HealthzGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiHealthzGetRequest struct via the builder pattern


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


## MetricsGet

> string MetricsGet(ctx).Execute()

Prometheus metrics



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
	resp, r, err := apiClient.KubernetesAPI.MetricsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.MetricsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MetricsGet`: string
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.MetricsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiMetricsGetRequest struct via the builder pattern


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


## ReadyzDisablePost

> string ReadyzDisablePost(ctx).Execute()

Disable ready state



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
	resp, r, err := apiClient.KubernetesAPI.ReadyzDisablePost(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.ReadyzDisablePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReadyzDisablePost`: string
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.ReadyzDisablePost`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiReadyzDisablePostRequest struct via the builder pattern


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


## ReadyzEnablePost

> string ReadyzEnablePost(ctx).Execute()

Enable ready state



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
	resp, r, err := apiClient.KubernetesAPI.ReadyzEnablePost(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.ReadyzEnablePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReadyzEnablePost`: string
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.ReadyzEnablePost`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiReadyzEnablePostRequest struct via the builder pattern


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


## ReadyzGet

> string ReadyzGet(ctx).Execute()

Readiness check



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
	resp, r, err := apiClient.KubernetesAPI.ReadyzGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KubernetesAPI.ReadyzGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReadyzGet`: string
	fmt.Fprintf(os.Stdout, "Response from `KubernetesAPI.ReadyzGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiReadyzGetRequest struct via the builder pattern


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

