# Completion


## Fields

| Field                                                                                               | Type                                                                                                | Required                                                                                            | Description                                                                                         |
| --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- |
| `ID`                                                                                                | **string*                                                                                           | :heavy_minus_sign:                                                                                  | The unique identifier for this task.                                                                |
| `OrganizationID`                                                                                    | *string*                                                                                            | :heavy_check_mark:                                                                                  | The unique identifier for the organization.                                                         |
| `Kind`                                                                                              | [*components.Kind](../../models/components/kind.md)                                                 | :heavy_minus_sign:                                                                                  | The kind of completion i.e. chat messages or prompt                                                 |
| `Status`                                                                                            | [*components.Status](../../models/components/status.md)                                             | :heavy_minus_sign:                                                                                  | The status of this completion.                                                                      |
| `Tags`                                                                                              | []*string*                                                                                          | :heavy_minus_sign:                                                                                  | The tags for this completion.                                                                       |
| `Request`                                                                                           | [*components.CreateChatCompletionRequest](../../models/components/createchatcompletionrequest.md)   | :heavy_minus_sign:                                                                                  | N/A                                                                                                 |
| `Response`                                                                                          | [*components.CreateChatCompletionResponse](../../models/components/createchatcompletionresponse.md) | :heavy_minus_sign:                                                                                  | Represents a chat completion response returned by model, based on the provided input.               |
| `Stacktrace`                                                                                        | [][components.Stacktrace](../../models/components/stacktrace.md)                                    | :heavy_minus_sign:                                                                                  | The stacktrace for this completion.                                                                 |
| `SessionID`                                                                                         | **string*                                                                                           | :heavy_minus_sign:                                                                                  | The session id for this completion.                                                                 |
| `Duration`                                                                                          | **float64*                                                                                          | :heavy_minus_sign:                                                                                  | The duration of this completion in seconds.                                                         |
| `FailureKind`                                                                                       | **string*                                                                                           | :heavy_minus_sign:                                                                                  | The failure kind of this completion.                                                                |
| `FailureReason`                                                                                     | **string*                                                                                           | :heavy_minus_sign:                                                                                  | The failure reason of this completion.                                                              |