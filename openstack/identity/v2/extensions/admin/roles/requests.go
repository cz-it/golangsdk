package roles

import (
	"github.com/huaweicloudsdk/golangsdk"
	"github.com/huaweicloudsdk/golangsdk/pagination"
)

// List is the operation responsible for listing all available global roles
// that a user can adopt.
func List(client *golangsdk.ServiceClient) pagination.Pager {
	return pagination.NewPager(client, rootURL(client), func(r pagination.PageResult) pagination.Page {
		return RolePage{pagination.SinglePageBase(r)}
	})
}

// AddUser is the operation responsible for assigning a particular role to
// a user. This is confined to the scope of the user's tenant - so the tenant
// ID is a required argument.
func AddUser(client *golangsdk.ServiceClient, tenantID, userID, roleID string) (r UserRoleResult) {
	_, r.Err = client.Put(userRoleURL(client, tenantID, userID, roleID), nil, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201},
	})
	return
}

// DeleteUser is the operation responsible for deleting a particular role
// from a user. This is confined to the scope of the user's tenant - so the
// tenant ID is a required argument.
func DeleteUser(client *golangsdk.ServiceClient, tenantID, userID, roleID string) (r UserRoleResult) {
	_, r.Err = client.Delete(userRoleURL(client, tenantID, userID, roleID), nil)
	return
}
