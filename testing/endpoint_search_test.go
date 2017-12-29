package testing

import (
	"testing"

	"github.com/huaweicloudsdk/golangsdk"
	th "github.com/huaweicloudsdk/golangsdk/testhelper"
)

func TestApplyDefaultsToEndpointOpts(t *testing.T) {
	eo := golangsdk.EndpointOpts{Availability: golangsdk.AvailabilityPublic}
	eo.ApplyDefaults("compute")
	expected := golangsdk.EndpointOpts{Availability: golangsdk.AvailabilityPublic, Type: "compute"}
	th.CheckDeepEquals(t, expected, eo)

	eo = golangsdk.EndpointOpts{Type: "compute"}
	eo.ApplyDefaults("object-store")
	expected = golangsdk.EndpointOpts{Availability: golangsdk.AvailabilityPublic, Type: "compute"}
	th.CheckDeepEquals(t, expected, eo)
}
