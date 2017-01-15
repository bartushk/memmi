package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"pbuf"
)

func main() {
	test_msg := pbuf.Test{
		Label: proto.String("asdf"),
		Type:  proto.Int32(2),
		Reps:  []int64{1, 2, 3},
		Optionalgroup: &pbuf.Test_OptionalGroup{
			RequiredField: proto.String("good bye"),
		},
	}
	fmt.Println(test_msg.GetOptionalgroup().GetRequiredField())

}
