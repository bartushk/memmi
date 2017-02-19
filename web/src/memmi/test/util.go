package test

import (
	"github.com/golang/protobuf/proto"
	"testing"
)

func AssertProtoEq(t *testing.T, expected proto.Message, actual proto.Message, desc string) {
	if !proto.Equal(expected, actual) {
		t.Error(desc,
			"Expected:", expected,
			"Got:", actual)
	}
}
