package vips

import "github.com/huaweicloudsdk/golangsdk"

const (
	rootPath     = "lb"
	resourcePath = "vips"
)

func rootURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL(rootPath, resourcePath)
}

func resourceURL(c *golangsdk.ServiceClient, id string) string {
	return c.ServiceURL(rootPath, resourcePath, id)
}
