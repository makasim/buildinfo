# buildinfo

```golang
package main

import (
	"fmt"

	"github.com/makasim/buildinfo"
)

func main() {
	fmt.Println(buildinfo.String())
}
```

```shell
COMMIT=$(git rev-list -1 HEAD | cut -c1-8)
BUILD_DATE=${BUILD_DATE:-$(date +"%Y-%m-%d_%T%z")}
VERSION="1.2.3"

go build \
  -X github.com/makasim/buildinfo.Commit=${COMMIT} \
  -X github.com/makasim/buildinfo.Version=${VERSION} \
  -X github.com/makasim/buildinfo.Date=${BUILD_DATE} \
  main.go
```