package adsc

// Import all Envoy filter types so they are registered and deserialization does not fail
// when using them in the "typed_config" attributes.
import (
	udpa "github.com/cncf/udpa/go/udpa/type/v1"
	_ "istio.io/istio/pkg/config/xds"
	_ "istio.io/istio/pkg/envoy/config/filter/http/alpn/v2alpha1"
)

// Statically link protobuf descriptors from UDPA
var _ = udpa.TypedStruct{}
