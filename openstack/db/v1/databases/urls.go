package databases

import "github.com/huaweicloudsdk/golangsdk"

func baseURL(c *golangsdk.ServiceClient, instanceID string) string {
	return c.ServiceURL("instances", instanceID, "databases")
}

func dbURL(c *golangsdk.ServiceClient, instanceID, dbName string) string {
	return c.ServiceURL("instances", instanceID, "databases", dbName)
}
