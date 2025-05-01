# Feedback
(*Feedback*)

## Overview

Feedback

### Available Operations

* [Get](#get) - Fetch feedback by id.
* [List](#list) - List feedback
* [Upload](#upload) - Upload a piece of feedback

## Get

Fetch feedback by id.

### Example Usage

```go
package main

import(
	"context"
	"github.com/log10-io/log10go"
	"log"
)

func main() {
    ctx := context.Background()

    s := log10go.New(
        log10go.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Feedback.Get(ctx, "<id>", log10go.String("<value>"))
    if err != nil {
        log.Fatal(err)
    }
    if res.Feedback != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `feedbackID`                                             | *string*                                                 | :heavy_check_mark:                                       | The feedback id to fetch.                                |
| `xLog10Organization`                                     | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetResponse](../../models/operations/getresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## List

List feedback

### Example Usage

```go
package main

import(
	"context"
	"github.com/log10-io/log10go"
	"log"
)

func main() {
    ctx := context.Background()

    s := log10go.New(
        log10go.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Feedback.List(ctx, log10go.String("<value>"), nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                 | Type                                                                      | Required                                                                  | Description                                                               |
| ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| `ctx`                                                                     | [context.Context](https://pkg.go.dev/context#Context)                     | :heavy_check_mark:                                                        | The context to use for the request.                                       |
| `xLog10Organization`                                                      | **string*                                                                 | :heavy_minus_sign:                                                        | N/A                                                                       |
| `requestBody`                                                             | [*operations.ListRequestBody](../../models/operations/listrequestbody.md) | :heavy_minus_sign:                                                        | N/A                                                                       |
| `opts`                                                                    | [][operations.Option](../../models/operations/option.md)                  | :heavy_minus_sign:                                                        | The options for this request.                                             |

### Response

**[*operations.ListResponse](../../models/operations/listresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Upload

Upload a piece of feedback

### Example Usage

```go
package main

import(
	"context"
	"github.com/log10-io/log10go"
	"github.com/log10-io/log10go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := log10go.New(
        log10go.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Feedback.Upload(ctx, operations.CreateUploadRequestBodyOne(
        operations.One{
            TaskID: "<id>",
            JSONValues: operations.JSONValues{},
            MatchedCompletionIds: []string{
                "<value>",
                "<value>",
            },
            Comment: "The Apollotech B340 is an affordable wireless mouse with reliable connectivity, 12 months battery life and modern design",
            CompletionTagsSelector: []string{
                "<value>",
            },
        },
    ), log10go.String("<value>"))
    if err != nil {
        log.Fatal(err)
    }
    if res.Feedback != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `requestBody`                                                                | [operations.UploadRequestBody](../../models/operations/uploadrequestbody.md) | :heavy_check_mark:                                                           | N/A                                                                          |
| `xLog10Organization`                                                         | **string*                                                                    | :heavy_minus_sign:                                                           | N/A                                                                          |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*operations.UploadResponse](../../models/operations/uploadresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |