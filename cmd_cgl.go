////////////////////////////////////////////////////////////////////////////
// Porgram: cmd_cgl - Config Group List sub-command handling
// Authors: Antonio Sun (c) 2015, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import "fmt"

func cmd_cgl(options Options) error {
	for _, pods := range config.Pod {
		fmt.Printf("ID: %s\r\n", pods.Id)
		for _, inst := range pods.Instance {
			fmt.Printf("  DB: %s\r\n", inst.Database)
		}
	}
	return nil
}
