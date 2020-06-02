# msgpb64

## Description

Package msgpb64 provides encoder/decoder that combines msgpack and base64 to serialize/deserialize any structure as a string.

This is useful when we need a string as the payload, like HTTP headers.

## Install

```
go get github.com/johejo/msgpb64
```

## Example

```go
package msgpb64_test

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/johejo/msgpb64"
)

func Example() {
	type test struct {
		Foo string
		Bar int
	}

	v := test{Foo: "foo", Bar: 99}

	var buf bytes.Buffer
	if err := msgpb64.NewEncoder(base64.StdEncoding, &buf).Encode(v); err != nil {
		panic(err)
	}

	s := buf.String()
	fmt.Println(s)

	var vv test
	if err := msgpb64.NewDecoder(base64.StdEncoding, strings.NewReader(s)).Decode(&vv); err != nil {
		panic(err)
	}
	fmt.Printf("%v", vv)

	// Output:
	// gqNGb2+jZm9vo0JhctMAAAAAAAAAYw==
	// {foo 99}
}
```


## Benchmark

Comparison with `encoding/json`

```
Machine: Dell XPS 15 7590
OS: Ubuntu 20.04
CPU: Intel(R) Core(TM) i9-9980HK CPU @ 2.40GHz
Memory: 32GB
```

```
$ go test -bench . -benchmem
goos: linux
goarch: amd64
pkg: github.com/johejo/msgpb64
Benchmark_msgpb64_one-16          471180              4478 ns/op            8013 B/op         20 allocs/op
Benchmark_json_one-16            1000000              1676 ns/op            1080 B/op         11 allocs/op
Benchmark_msgpb64_large-16          9640            107349 ns/op           56837 B/op        170 allocs/op
Benchmark_json_large-16             9469            137980 ns/op           43953 B/op        499 allocs/op
PASS
ok      github.com/johejo/msgpb64       7.141s
```

## License

MIT

## Author

Mitsuo Heijo (@johejo)