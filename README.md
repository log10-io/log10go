# github.com/log10-io/log10go

<div align="left">
    <a href="https://speakeasyapi.dev/"><img src="https://custom-icon-badges.demolab.com/badge/-Built%20By%20Speakeasy-212015?style=for-the-badge&logoColor=FBE331&logo=speakeasy&labelColor=545454" /></a>
    <a href="https://opensource.org/licenses/MIT">
        <img src="https://img.shields.io/badge/License-MIT-blue.svg" style="width: 100px; height: 28px;" />
    </a>
</div>


## 🏗 **Welcome to your new SDK!** 🏗

It has been generated successfully based on your OpenAPI spec. However, it is not yet ready for production use. Here are some next steps:
- [ ] 🛠 Make your SDK feel handcrafted by [customizing it](https://www.speakeasyapi.dev/docs/customize-sdks)
- [ ] ♻️ Refine your SDK quickly by iterating locally with the [Speakeasy CLI](https://github.com/speakeasy-api/speakeasy)
- [ ] 🎁 Publish your SDK to package managers by [configuring automatic publishing](https://www.speakeasyapi.dev/docs/advanced-setup/publish-sdks)
- [ ] ✨ When ready to productionize, delete this section from the README

<!-- Start Summary [summary] -->
## Summary

Log10 Feedback API Spec: Log10 Feedback API Spec
<!-- End Summary [summary] -->

<!-- Start Table of Contents [toc] -->
## Table of Contents
<!-- $toc-max-depth=2 -->
* [github.com/log10-io/log10go](#githubcomlog10-iolog10go)
  * [🏗 **Welcome to your new SDK!** 🏗](#welcome-to-your-new-sdk)
  * [SDK Installation](#sdk-installation)
  * [SDK Example Usage](#sdk-example-usage)
  * [Available Resources and Operations](#available-resources-and-operations)
  * [Error Handling](#error-handling)
  * [Server Selection](#server-selection)
  * [Custom HTTP Client](#custom-http-client)
  * [Authentication](#authentication)
  * [Retries](#retries)
* [Development](#development)
  * [Maturity](#maturity)
  * [Contributions](#contributions)

<!-- End Table of Contents [toc] -->

<!-- Start SDK Installation [installation] -->
## SDK Installation

To add the SDK as a dependency to your project:
```bash
go get github.com/log10-io/log10go
```
<!-- End SDK Installation [installation] -->

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Example

```go
package main

import (
	"context"
	"github.com/log10-io/log10go"
	"log"
)

func main() {
	ctx := context.Background()

	s := log10go.New(
		log10go.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	res, err := s.Sessions.Create(ctx, log10go.String("<value>"))
	if err != nil {
		log.Fatal(err)
	}
	if res.Object != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

<details open>
<summary>Available methods</summary>

### [Completions](docs/sdks/completions/README.md)

* [Create](docs/sdks/completions/README.md#create) - Create a completion
* [Update](docs/sdks/completions/README.md#update) - Update completion by id.
* [ListUngraded](docs/sdks/completions/README.md#listungraded) - List ungraded completions i.e. completions that have not been associated with feedback but matches task selector.

### [Feedback](docs/sdks/feedback/README.md)

* [Get](docs/sdks/feedback/README.md#get) - Fetch feedback by id.
* [List](docs/sdks/feedback/README.md#list) - List feedback
* [Upload](docs/sdks/feedback/README.md#upload) - Upload a piece of feedback

### [FeedbackTasks](docs/sdks/feedbacktasks/README.md)

* [List](docs/sdks/feedbacktasks/README.md#list) - List feedback tasks.
* [Create](docs/sdks/feedbacktasks/README.md#create) - Create a new task.
* [Get](docs/sdks/feedbacktasks/README.md#get) - Retrieves feedback task `taskId`.


### [Sessions](docs/sdks/sessions/README.md)

* [Create](docs/sdks/sessions/README.md#create) - Create a session

</details>
<!-- End Available Resources and Operations [operations] -->

<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations. All operations return a response object or an error, they will never return both.

By Default, an API error will return `sdkerrors.SDKError`. When custom error responses are specified for an operation, the SDK may also return their associated error. You can refer to respective *Errors* tables in SDK docs for more details on possible error types for each operation.

For example, the `Create` function may return the following errors:

| Error Type         | Status Code | Content Type |
| ------------------ | ----------- | ------------ |
| sdkerrors.SDKError | 4XX, 5XX    | \*/\*        |

### Example

```go
package main

import (
	"context"
	"errors"
	"github.com/log10-io/log10go"
	"github.com/log10-io/log10go/models/components"
	"github.com/log10-io/log10go/models/sdkerrors"
	"log"
)

func main() {
	ctx := context.Background()

	s := log10go.New(
		log10go.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	res, err := s.Completions.Create(ctx, components.Completion{
		OrganizationID: "<id>",
		Request: &components.CreateChatCompletionRequest{
			Messages: []components.ChatCompletionRequestMessage{
				components.CreateChatCompletionRequestMessageChatCompletionRequestToolMessage(
					components.ChatCompletionRequestToolMessage{
						Role:       components.ChatCompletionRequestToolMessageRoleTool,
						Content:    "<value>",
						ToolCallID: "<id>",
					},
				),
			},
			Model: components.CreateModelStr(
				"gpt-4-turbo",
			),
			ResponseFormat: &components.ResponseFormat{},
			User:           log10go.String("user-1234"),
		},
	}, log10go.String("<value>"))
	if err != nil {

		var e *sdkerrors.SDKError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling [errors] -->

<!-- Start Server Selection [server] -->
## Server Selection

### Override Server URL Per-Client

The default server can be overridden globally using the `WithServerURL(serverURL string)` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	"github.com/log10-io/log10go"
	"github.com/log10-io/log10go/models/components"
	"log"
)

func main() {
	ctx := context.Background()

	s := log10go.New(
		log10go.WithServerURL("https://log10.io"),
		log10go.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	res, err := s.Completions.Create(ctx, components.Completion{
		OrganizationID: "<id>",
		Request: &components.CreateChatCompletionRequest{
			Messages: []components.ChatCompletionRequestMessage{
				components.CreateChatCompletionRequestMessageChatCompletionRequestToolMessage(
					components.ChatCompletionRequestToolMessage{
						Role:       components.ChatCompletionRequestToolMessageRoleTool,
						Content:    "<value>",
						ToolCallID: "<id>",
					},
				),
			},
			Model: components.CreateModelStr(
				"gpt-4-turbo",
			),
			ResponseFormat: &components.ResponseFormat{},
			User:           log10go.String("user-1234"),
		},
	}, log10go.String("<value>"))
	if err != nil {
		log.Fatal(err)
	}
	if res.Any != nil {
		// handle response
	}
}

```
<!-- End Server Selection [server] -->

<!-- Start Custom HTTP Client [http-client] -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"
	"github.com/myorg/your-go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = sdk.New(sdk.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client [http-client] -->

<!-- Start Authentication [security] -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security scheme globally:

| Name         | Type   | Scheme  |
| ------------ | ------ | ------- |
| `Log10Token` | apiKey | API key |

You can configure it using the `WithSecurity` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	"github.com/log10-io/log10go"
	"github.com/log10-io/log10go/models/components"
	"log"
)

func main() {
	ctx := context.Background()

	s := log10go.New(
		log10go.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	res, err := s.Completions.Create(ctx, components.Completion{
		OrganizationID: "<id>",
		Request: &components.CreateChatCompletionRequest{
			Messages: []components.ChatCompletionRequestMessage{
				components.CreateChatCompletionRequestMessageChatCompletionRequestToolMessage(
					components.ChatCompletionRequestToolMessage{
						Role:       components.ChatCompletionRequestToolMessageRoleTool,
						Content:    "<value>",
						ToolCallID: "<id>",
					},
				),
			},
			Model: components.CreateModelStr(
				"gpt-4-turbo",
			),
			ResponseFormat: &components.ResponseFormat{},
			User:           log10go.String("user-1234"),
		},
	}, log10go.String("<value>"))
	if err != nil {
		log.Fatal(err)
	}
	if res.Any != nil {
		// handle response
	}
}

```
<!-- End Authentication [security] -->

<!-- Start Retries [retries] -->
## Retries

Some of the endpoints in this SDK support retries. If you use the SDK without any configuration, it will fall back to the default retry strategy provided by the API. However, the default retry strategy can be overridden on a per-operation basis, or across the entire SDK.

To change the default retry strategy for a single API call, simply provide a `retry.Config` object to the call by using the `WithRetries` option:
```go
package main

import (
	"context"
	"github.com/log10-io/log10go"
	"github.com/log10-io/log10go/models/components"
	"github.com/log10-io/log10go/retry"
	"log"
	"models/operations"
)

func main() {
	ctx := context.Background()

	s := log10go.New(
		log10go.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	res, err := s.Completions.Create(ctx, components.Completion{
		OrganizationID: "<id>",
		Request: &components.CreateChatCompletionRequest{
			Messages: []components.ChatCompletionRequestMessage{
				components.CreateChatCompletionRequestMessageChatCompletionRequestToolMessage(
					components.ChatCompletionRequestToolMessage{
						Role:       components.ChatCompletionRequestToolMessageRoleTool,
						Content:    "<value>",
						ToolCallID: "<id>",
					},
				),
			},
			Model: components.CreateModelStr(
				"gpt-4-turbo",
			),
			ResponseFormat: &components.ResponseFormat{},
			User:           log10go.String("user-1234"),
		},
	}, log10go.String("<value>"), operations.WithRetries(
		retry.Config{
			Strategy: "backoff",
			Backoff: &retry.BackoffStrategy{
				InitialInterval: 1,
				MaxInterval:     50,
				Exponent:        1.1,
				MaxElapsedTime:  100,
			},
			RetryConnectionErrors: false,
		}))
	if err != nil {
		log.Fatal(err)
	}
	if res.Any != nil {
		// handle response
	}
}

```

If you'd like to override the default retry strategy for all operations that support retries, you can use the `WithRetryConfig` option at SDK initialization:
```go
package main

import (
	"context"
	"github.com/log10-io/log10go"
	"github.com/log10-io/log10go/models/components"
	"github.com/log10-io/log10go/retry"
	"log"
)

func main() {
	ctx := context.Background()

	s := log10go.New(
		log10go.WithRetryConfig(
			retry.Config{
				Strategy: "backoff",
				Backoff: &retry.BackoffStrategy{
					InitialInterval: 1,
					MaxInterval:     50,
					Exponent:        1.1,
					MaxElapsedTime:  100,
				},
				RetryConnectionErrors: false,
			}),
		log10go.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	res, err := s.Completions.Create(ctx, components.Completion{
		OrganizationID: "<id>",
		Request: &components.CreateChatCompletionRequest{
			Messages: []components.ChatCompletionRequestMessage{
				components.CreateChatCompletionRequestMessageChatCompletionRequestToolMessage(
					components.ChatCompletionRequestToolMessage{
						Role:       components.ChatCompletionRequestToolMessageRoleTool,
						Content:    "<value>",
						ToolCallID: "<id>",
					},
				),
			},
			Model: components.CreateModelStr(
				"gpt-4-turbo",
			),
			ResponseFormat: &components.ResponseFormat{},
			User:           log10go.String("user-1234"),
		},
	}, log10go.String("<value>"))
	if err != nil {
		log.Fatal(err)
	}
	if res.Any != nil {
		// handle response
	}
}

```
<!-- End Retries [retries] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->

# Development

## Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

## Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release!

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
