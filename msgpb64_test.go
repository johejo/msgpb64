package msgpb64_test

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"os"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/johejo/msgpb64"
)

func Test_msgpb64(t *testing.T) {
	type test struct {
		Foo string
		Bar int
	}
	v := test{Foo: "foo", Bar: 99}

	var buf bytes.Buffer
	if err := msgpb64.NewEncoder(base64.StdEncoding, &buf).Encode(v); err != nil {
		t.Fatal(err)
	}

	s := buf.String()
	const want = "gqNGb2+jZm9vo0JhctMAAAAAAAAAYw=="
	if s != want {
		t.Fatalf("invalid encode result: want=%s, got=%s", want, s)
	}

	var vv test
	if err := msgpb64.NewDecoder(base64.StdEncoding, strings.NewReader(s)).Decode(&vv); err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(v, vv); diff != "" {
		t.Fatal(diff)
	}
}

func Test_msgpb64_large(t *testing.T) {
	f, err := os.Open("testdata/large.json")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	var m []map[string]interface{}
	if err := json.NewDecoder(f).Decode(&m); err != nil {
		t.Fatal(err)
	}

	var buf bytes.Buffer
	if err := msgpb64.NewEncoder(base64.StdEncoding, &buf).Encode(m); err != nil {
		t.Fatal(err)
	}

	var mm []map[string]interface{}
	if err := msgpb64.NewDecoder(base64.StdEncoding, &buf).Decode(&mm); err != nil {
		t.Fatal(err)
	}

	if diff := cmp.Diff(m, mm); diff != "" {
		t.Fatal(diff)
	}
}
