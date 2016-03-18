////////////////////////////////////////////////////////////////////////////
// Porgram: cglCmd - Config Group List command handling
// Authors: Antonio Sun (c) 2015-16, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import "fmt"

func cglCmd(options Options) error {
	for _, pods := range config.Pod {
		fmt.Printf("%s:\r\n", pods.Id)
		for _, inst := range pods.Instance {
			fmt.Printf("  %s\r\n", inst.Database)
		}
	}
	return nil
}
