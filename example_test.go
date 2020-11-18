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
	// gqNGb2+jZm9vo0JhcmM=
	// {foo 99}
}
