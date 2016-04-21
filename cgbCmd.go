////////////////////////////////////////////////////////////////////////////
// Porgram: cgbCmd - Config Group Brief command handling
// Authors: Antonio Sun (c) 2016, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
)

func cgbCmd(options Options) error {
	fmt.Printf("Instance Listing\r\n")
	for _, pods := range config.Pod {
		for _, inst := range pods.Instance {
			fmt.Printf("\r\n%s:\r\n  https://perftest%s.dayforce.com/\r\n",
				pods.Id, pods.Id)
			fmt.Printf("  Client/DB Name: %s\r\n", inst.Database)
			fmt.Printf("  DB Server: %s\r\n", inst.Dbserver)
		}
	}
	return nil
}
