// Code generated by protoc-gen-go-binary. DO NOT EDIT.
// source: pbmesh/v2beta1/computed_gateway_routes.proto

package meshv2beta1

import (
	"google.golang.org/protobuf/proto"
)

// MarshalBinary implements encoding.BinaryMarshaler
func (msg *ComputedGatewayRoutes) MarshalBinary() ([]byte, error) {
	return proto.Marshal(msg)
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler
func (msg *ComputedGatewayRoutes) UnmarshalBinary(b []byte) error {
	return proto.Unmarshal(b, msg)
}