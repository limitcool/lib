package lib

import (
	"encoding/json"

	"github.com/golang/protobuf/ptypes/any"
	"github.com/golang/protobuf/ptypes/wrappers"
	"golang.org/x/exp/constraints"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

func IntToBool[T constraints.Signed](a T) bool {
	switch a {
	case 0, -1:
		return false
	}
	return true
}

func ConvertAnyToInterface(anyValue *any.Any) (interface{}, error) {
	var value interface{}
	bytesValue := &wrappers.BytesValue{}
	err := anypb.UnmarshalTo(anyValue, bytesValue, proto.UnmarshalOptions{})
	if err != nil {
		return value, err
	}
	uErr := json.Unmarshal(bytesValue.Value, &value)
	if uErr != nil {
		return value, uErr
	}
	return value, nil
}

func ConvertInterfaceToAny(v interface{}) (*any.Any, error) {
	anyValue := &any.Any{}
	bytes, _ := json.Marshal(v)
	bytesValue := &wrappers.BytesValue{
		Value: bytes,
	}
	err := anypb.MarshalFrom(anyValue, bytesValue, proto.MarshalOptions{})
	return anyValue, err
}

func MustConvertInterfaceToAny(v interface{}) *any.Any {
	anyValue := &any.Any{}
	bytes, _ := json.Marshal(v)
	bytesValue := &wrappers.BytesValue{
		Value: bytes,
	}
	anypb.MarshalFrom(anyValue, bytesValue, proto.MarshalOptions{})
	return anyValue
}

func MustConvertAnyToInterface(anyValue *any.Any) interface{} {
	var value interface{}
	bytesValue := &wrappers.BytesValue{}
	err := anypb.UnmarshalTo(anyValue, bytesValue, proto.UnmarshalOptions{})
	if err != nil {
		return value
	}
	json.Unmarshal(bytesValue.Value, &value)

	return value
}
