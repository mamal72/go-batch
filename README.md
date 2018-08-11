[![Go Report Card](https://goreportcard.com/badge/github.com/mamal72/go-batch)](https://goreportcard.com/report/github.com/mamal72/go-batch)
[![GoDoc](https://godoc.org/github.com/mamal72/go-batch?status.svg)](https://godoc.org/github.com/mamal72/go-batch)
[![license](https://img.shields.io/github/license/mamal72/go-batch.svg)](https://github.com/mamal72/go-batch/blob/master/LICENSE)


# Go Batch

This is a simple API client to send push notifications using [Batch](https://batch.com).


## Installation

```bash
go get github.com/mamal72/go-batch
# or use dep, glide, vgo or whatever you like
```


## Usage

```go
package main

import "github.com/mamal72/go-batch"

func main() {
    client := batch.NewClient("YOUR_BATCH_REST_API_KEY", "YOUR_BATCH_API_KEY_FOR_YOUR_APP")
    payload := batch.TransactionalPushPayload{
        GroupID: "SOME_GROUP_ID",
        Recipients: batch.TransactionalPushRecipient{
            CustomIDs: []string{"1", "2"},
            // blah blah blah
        },
        Message: batch.TransactionalPushMessage{
            Title: "Greeting",
            Body:  "Hello dear user!",
        },
        // blah blah blah
    }

    response, err := client.TransactionalSend(payload)
    // handle error or get and store token from response
}
```


## Ideas || Issues

Just fill an issue and describe it. I'll check it ASAP!


## Contribution

You can fork the repository, improve or fix some part of it and then send the pull requests back if you want to see them here. I really appreciate that. :heart:


## License
> MIT
