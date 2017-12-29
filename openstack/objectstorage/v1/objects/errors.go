package objects

import "github.com/huaweicloudsdk/golangsdk"

// ErrWrongChecksum is the error when the checksum generated for an object
// doesn't match the ETAG header.
type ErrWrongChecksum struct {
	golangsdk.BaseError
}

func (e ErrWrongChecksum) Error() string {
	return "Local checksum does not match API ETag header"
}
