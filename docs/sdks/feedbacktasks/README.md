# FeedbackTasks
(*FeedbackTasks*)

## Overview

FeedbackTasks

### Available Operations

* [List](#list) - List feedback tasks.
* [Create](#create) - Create a new task.
* [Get](#get) - Retrieves feedback task `taskId`.

## List

List feedback tasks.

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

    res, err := s.FeedbackTasks.List(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.Tasks != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ListFeedbackTasksResponse](../../models/operations/listfeedbacktasksresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Create

Create a new task.

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

    res, err := s.FeedbackTasks.Create(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Task != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `request`                                                | [components.Task](../../models/components/task.md)       | :heavy_check_mark:                                       | The request object to use for the request.               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.CreateFeedbackTaskResponse](../../models/operations/createfeedbacktaskresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Get

Retrieves feedback task `taskId`.

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

    res, err := s.FeedbackTasks.Get(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.Task != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `taskID`                                                 | *string*                                                 | :heavy_check_mark:                                       | The task id to fetch.                                    |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetFeedbackTaskResponse](../../models/operations/getfeedbacktaskresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |