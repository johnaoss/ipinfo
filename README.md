# ipinfo

ipinfo is a client library for [ipinfo.io](https://ipinfo.io).

## Installation

To install, run the following:

```bash
go get github.com/johnaoss/ipinfo
```

## Usage

Here's an example on how to obtain the City an IP is located within.

```go
package main

import (
    "fmt"

    "github.com/johnaoss/ipinfo"
)

func main() {
    resp, _ := ipinfo.Info("8.8.8.8")
    fmt.Println(resp.City)
}
```

## Testing

To run tests, ensure you have the env var `$IPINFO_TOKEN` set to be your bearer
token obtained from your dashboard.

## License

See the `LICENSE.md` file for more details.