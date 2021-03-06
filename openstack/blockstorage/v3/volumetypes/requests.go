package volumetypes

import (
	"github.com/huaweicloudsdk/golangsdk"
	"github.com/huaweicloudsdk/golangsdk/pagination"
)

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToVolumeTypeCreateMap() (map[string]interface{}, error)
}

// CreateOpts contains options for creating a Volume Type. This object is passed to
// the volumetypes.Create function. For more information about these parameters,
// see the Volume Type object.
type CreateOpts struct {
	// The name of the volume type
	Name string `json:"name" required:"true"`
	// The volume type description
	Description string `json:"description,omitempty"`
	// the ID of the existing volume snapshot
	PublicAccess bool `json:"os-volume-type-access:is_public"`
	// Extra spec key-value pairs defined by the user.
	ExtraSpecs map[string]string `json:"extra_specs"`
}

// ToVolumeTypeCreateMap assembles a request body based on the contents of a
// CreateOpts.
func (opts CreateOpts) ToVolumeTypeCreateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "volume_type")
}

// Create will create a new Volume Type based on the values in CreateOpts. To extract
// the Volume Type object from the response, call the Extract method on the
// CreateResult.
func Create(client *golangsdk.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToVolumeTypeCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(createURL(client), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

// Delete will delete the existing Volume Type with the provided ID.
func Delete(client *golangsdk.ServiceClient, id string) (r DeleteResult) {
	_, r.Err = client.Delete(deleteURL(client, id), nil)
	return
}

// Get retrieves the Volume Type with the provided ID. To extract the Volume Type object
// from the response, call the Extract method on the GetResult.
func Get(client *golangsdk.ServiceClient, id string) (r GetResult) {
	_, r.Err = client.Get(getURL(client, id), &r.Body, nil)
	return
}

// ListOptsBuilder allows extensions to add additional parameters to the List
// request.
type ListOptsBuilder interface {
	ToVolumeTypeListQuery() (string, error)
}

// ListOpts holds options for listing Volume Types. It is passed to the volumetypes.List
// function.
type ListOpts struct {
	// Comma-separated list of sort keys and optional sort directions in the
	// form of <key>[:<direction>].
	Sort string `q:"sort"`
	// Requests a page size of items.
	Limit int `q:"limit"`
	// Used in conjunction with limit to return a slice of items.
	Offset int `q:"offset"`
	// The ID of the last-seen item.
	Marker string `q:"marker"`
}

// ToVolumeTypeListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToVolumeTypeListQuery() (string, error) {
	q, err := golangsdk.BuildQueryString(opts)
	return q.String(), err
}

// List returns Volume types.
func List(client *golangsdk.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listURL(client)

	if opts != nil {
		query, err := opts.ToVolumeTypeListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}

	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return VolumeTypePage{pagination.LinkedPageBase{PageResult: r}}
	})
}
