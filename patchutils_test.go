package patchutils

import (
	"io/ioutil"
	"reflect"
	"testing"
)

func TestCombine(t *testing.T) {
	out, err := Combine("./test/tmp1.patch", "./test/tmp2.patch")
	if err != nil {
		t.Fatal(err)
	}

	b, err := ioutil.ReadAll(out)
	if err != nil {
		t.Fatal(err)
	}

	expected, err := ioutil.ReadFile("./test/combined.patch")

	if !reflect.DeepEqual(expected, b) {
		t.Errorf("expected\n%s\ngot\n%s", string(expected), string(b))
	}
}