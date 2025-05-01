# Completions
(*Completions*)

## Overview

Completions

### Available Operations

* [Create](#create) - Create a completion
* [Update](#update) - Update completion by id.
* [ListUngraded](#listungraded) - List ungraded completions i.e. completions that have not been associated with feedback but matches task selector.

## Create

Create a completion

### Example Usage

```go
package main

import(
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
                        Role: components.ChatCompletionRequestToolMessageRoleTool,
                        Content: "<value>",
                        ToolCallID: "<id>",
                    },
                ),
            },
            Model: components.CreateModelStr(
                "gpt-4-turbo",
            ),
            ResponseFormat: &components.ResponseFormat{},
            User: log10go.String("user-1234"),
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

### Parameters

| Parameter                                                      | Type                                                           | Required                                                       | Description                                                    |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `ctx`                                                          | [context.Context](https://pkg.go.dev/context#Context)          | :heavy_check_mark:                                             | The context to use for the request.                            |
| `completion`                                                   | [components.Completion](../../models/components/completion.md) | :heavy_check_mark:                                             | N/A                                                            |
| `xLog10Organization`                                           | **string*                                                      | :heavy_minus_sign:                                             | N/A                                                            |
| `opts`                                                         | [][operations.Option](../../models/operations/option.md)       | :heavy_minus_sign:                                             | The options for this request.                                  |

### Response

**[*operations.CreateResponse](../../models/operations/createresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Update

Update completion by id.

### Example Usage

```go
package main

import(
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

    res, err := s.Completions.Update(ctx, "<id>", components.Completion{
        OrganizationID: "<id>",
        Request: &components.CreateChatCompletionRequest{
            Messages: []components.ChatCompletionRequestMessage{
                components.CreateChatCompletionRequestMessageChatCompletionRequestAssistantMessage(
                    components.ChatCompletionRequestAssistantMessage{
                        Role: components.ChatCompletionRequestAssistantMessageRoleAssistant,
                    },
                ),
                components.CreateChatCompletionRequestMessageChatCompletionRequestUserMessage(
                    components.ChatCompletionRequestUserMessage{
                        Content: components.CreateContentArrayOfChatCompletionRequestMessageContentPart(
                            []components.ChatCompletionRequestMessageContentPart{
                                components.CreateChatCompletionRequestMessageContentPartChatCompletionRequestMessageContentPartText(
                                    components.ChatCompletionRequestMessageContentPartText{
                                        Type: components.TypeText,
                                        Text: "<value>",
                                    },
                                ),
                                components.CreateChatCompletionRequestMessageContentPartChatCompletionRequestMessageContentPartImage(
                                    components.ChatCompletionRequestMessageContentPartImage{
                                        Type: components.ChatCompletionRequestMessageContentPartImageTypeImageURL,
                                        ImageURL: components.ImageURL{
                                            URL: "https://unlucky-hydrolyze.biz/",
                                        },
                                    },
                                ),
                            },
                        ),
                        Role: components.ChatCompletionRequestUserMessageRoleUser,
                    },
                ),
                components.CreateChatCompletionRequestMessageChatCompletionRequestUserMessage(
                    components.ChatCompletionRequestUserMessage{
                        Content: components.CreateContentStr(
                            "<value>",
                        ),
                        Role: components.ChatCompletionRequestUserMessageRoleUser,
                    },
                ),
            },
            Model: components.CreateModelStr(
                "gpt-4-turbo",
            ),
            ResponseFormat: &components.ResponseFormat{},
            User: log10go.String("user-1234"),
        },
    }, log10go.String("<value>"))
    if err != nil {
        log.Fatal(err)
    }
    if res.Completion != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                      | Type                                                           | Required                                                       | Description                                                    |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `ctx`                                                          | [context.Context](https://pkg.go.dev/context#Context)          | :heavy_check_mark:                                             | The context to use for the request.                            |
| `completionID`                                                 | *string*                                                       | :heavy_check_mark:                                             | The completion id to update.                                   |
| `completion`                                                   | [components.Completion](../../models/components/completion.md) | :heavy_check_mark:                                             | N/A                                                            |
| `xLog10Organization`                                           | **string*                                                      | :heavy_minus_sign:                                             | N/A                                                            |
| `opts`                                                         | [][operations.Option](../../models/operations/option.md)       | :heavy_minus_sign:                                             | The options for this request.                                  |

### Response

**[*operations.UpdateResponse](../../models/operations/updateresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListUngraded

List ungraded completions i.e. completions that have not been associated with feedback but matches task selector.

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

    res, err := s.Completions.ListUngraded(ctx, log10go.String("<value>"))
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `xLog10Organization`                                     | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ListUngradedResponse](../../models/operations/listungradedresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |