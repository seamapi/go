# Seam Go Library

[![fern shield](https://img.shields.io/badge/%F0%9F%8C%BF-SDK%20generated%20by%20Fern-brightgreen)](https://github.com/fern-api/fern)
[![go shield](https://img.shields.io/badge/go-docs-blue)](https://pkg.go.dev/github.com/seamapi/go)

The Seam Go library provides convenient access to the Seam API from Go.

# Requirements

This module requires Go version >= 1.13.

# Installation

Run the following command to use the Seam Go library in your module:
```sh
go get github.com/seamapi/go
```

# Usage

```go
import goclient "github.com/seamapi/go/client"

client := goclient.NewClient(goclient.WithAuthApiKey("<YOUR_AUTH_TOKEN>"))
```

## Access Codes

```go
import goclient "github.com/seamapi/go/client"

client := goclient.NewClient(goclient.WithAuthApiKey("<YOUR_AUTH_TOKEN>"))
accessCode, err := client.AccessCodes.Create(
  context.TODO(),
  &seamgo.AccessCodesCreateRequest{
    DeviceId: someDevice.DeviceId,
    Name:     seamgo.String("Test code"),
    Code:     seamgo.String("4444"),
  },
)
```

## Timeouts

Setting a timeout for each individual request is as simple as using the standard
`context` library. Setting a one second timeout for an individual API call looks
like the following:

```go
import seamgo "github.com/seamapi/go"

ctx, cancel := context.WithTimeout(context.TODO(), time.Second)
defer cancel()

devices, err = client.Devices.List(
  ctx,
  &seamgo.DevicesListRequest{
    DeviceType: seamgo.DeviceTypeAugustLock.Ptr(),
  },
)
```

## Client Options

A variety of client options are included to adapt the behavior of the library, which includes
configuring authorization tokens to be sent on every request, or providing your own instrumented
`*http.Client`. Both of these options are shown below:

```go
client := goclient.NewClient(
  goclient.WithAuthApiKey("<YOUR_AUTH_TOKEN>"),
  goclient.WithHTTPClient(
    &http.Client{
      Timeout: 5 * time.Second,
    },
  ),
)
```

> Providing your own `*http.Client` is recommended. Otherwise, the `http.DefaultClient` will be used,
> and your client will wait indefinitely for a response (unless the per-request, context-based timeout
> is used).

## Errors

Structured error types are returned from API calls that return non-success status codes. For example,
you can check if the error was due to a bad request (i.e. status code 400) with the following:

```go
accessCode, err := client.AccessCodes.Create(
  context.TODO(),
  &seamgo.AccessCodesCreateRequest{
    Name: seamgo.String("Bad Request"),
  },
)
if err != nil {
  if badRequestErr, ok := err.(*seamgo.BadRequestError);
    // Do something with the bad request ...
  }
  return err
}
```

These errors are also compatible with the `errors.Is` and `errors.As` APIs, so you can access the error
like so:

```go
accessCode, err := client.AccessCodes.Create(
  context.TODO(),
  &seamgo.AccessCodesCreateRequest{
    Name: seamgo.String("Bad Request"),
  },
)
if err != nil {
  var badRequestErr *seamgo.BadRequestError
  if errors.As(err, badRequestErr) {
    // Do something with the bad request ...
  }
  return err
}
```

If you'd like to wrap the errors with additional information and still retain the ability to access the type
with `errors.Is` and `errors.As`, you can use the `%w` directive:

```go
accessCode, err := client.AccessCodes.Create(
  context.TODO(),
  &seamgo.AccessCodesCreateRequest{
    Name: seamgo.String("Bad Request"),
  },
)
if err != nil {
  return fmt.Errorf("failed to create access code: %w", err)
}
```

# Beta Status

This SDK is in beta, and there may be breaking changes between versions without a major 
version update. Therefore, we recommend pinning the package version to a specific version. 
This way, you can install the same version each time without breaking changes.

# Contributing

While we value open-source contributions to this SDK, this library is generated programmatically. 
Additions made directly to this library would have to be moved over to our generation code, 
otherwise they would be overwritten upon the next generated release. Feel free to open a PR as
 a proof of concept, but know that we will not be able to merge it as-is. We suggest opening 
an issue first to discuss with us!

On the other hand, contributions to the README are always very welcome!
