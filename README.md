# gRPC Zerolog

[![BuildStatus Widget]][BuildStatus Result]
[![CodeCov Widget]][CodeCov Result]
[![GoReport Widget]][GoReport Status]
[![GoDoc Widget]][GoDoc]

[BuildStatus Result]: https://travis-ci.org/philip-bui/grpc-zerolog
[BuildStatus Widget]: https://travis-ci.org/philip-bui/grpc-zerolog.svg?branch=master

[CodeCov Result]: https://codecov.io/gh/philip-bui/grpc-zerolog
[CodeCov Widget]: https://codecov.io/gh/philip-bui/grpc-zerolog/branch/master/graph/badge.svg

[GoReport Status]: https://goreportcard.com/report/github.com/philip-bui/grpc-zerolog
[GoReport Widget]: https://goreportcard.com/badge/github.com/philip-bui/grpc-zerolog

[GoDoc]: https://godoc.org/github.com/philip-bui/grpc-zerolog
[GoDoc Widget]: https://godoc.org/github.com/philip-bui/grpc-zerolog?status.svg

```
import (
	"github.com/philip-bui/grpc-zerolog"
)

func main() {
	grpc.NewServer(
		zerolog.UnaryInterceptor(),
	)
}
```


