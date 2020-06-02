package msgpb64_test

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"os"
	"testing"

	"github.com/johejo/msgpb64"
)

type test struct {
	Foo string
	Bar int
}

var v = test{Foo: "foo", Bar: 99}

func Benchmark_msgpb64_one(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		if err := msgpb64.NewEncoder(base64.StdEncoding, &buf).Encode(v); err != nil {
			b.Fatal(err)
		}
		var vv test
		if err := msgpb64.NewDecoder(base64.StdEncoding, &buf).Decode(&vv); err != nil {
			b.Fatal(vv)
		}
	}
}

func Benchmark_json_one(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		if err := json.NewEncoder(&buf).Encode(v); err != nil {
			b.Fatal(err)
		}
		var vv test
		if err := json.NewDecoder(&buf).Decode(&vv); err != nil {
			b.Fatal(vv)
		}
	}
}

func getMap(b *testing.B) []map[string]interface{} {
	f, err := os.Open("testdata/large.json")
	if err != nil {
		b.Fatal(err)
	}
	defer f.Close()
	var m []map[string]interface{}
	if err := json.NewDecoder(f).Decode(&m); err != nil {
		b.Fatal(err)
	}
	return m
}

func Benchmark_msgpb64_large(b *testing.B) {
	v := getMap(b)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		if err := msgpb64.NewEncoder(base64.StdEncoding, &buf).Encode(v); err != nil {
			b.Fatal(err)
		}
		var vv []test
		if err := msgpb64.NewDecoder(base64.StdEncoding, &buf).Decode(&vv); err != nil {
			b.Fatal(vv)
		}
	}
}

func Benchmark_json_large(b *testing.B) {
	v := getMap(b)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		if err := json.NewEncoder(&buf).Encode(v); err != nil {
			b.Fatal(err)
		}
		var vv []test
		if err := json.NewDecoder(&buf).Decode(&vv); err != nil {
			b.Fatal(vv)
		}
	}
}
