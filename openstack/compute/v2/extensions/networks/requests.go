package networks

import (
	"github.com/huaweicloudsdk/golangsdk"
	"github.com/huaweicloudsdk/golangsdk/pagination"
)

// List returns a Pager that allows you to iterate over a collection of Network.
func List(client *golangsdk.ServiceClient) pagination.Pager {
	return pagination.NewPager(client, listURL(client), func(r pagination.PageResult) pagination.Page {
		return NetworkPage{pagination.SinglePageBase(r)}
	})
}

// Get returns data about a previously created Network.
func Get(client *golangsdk.ServiceClient, id string) (r GetResult) {
	_, r.Err = client.Get(getURL(client, id), &r.Body, nil)
	return
}
