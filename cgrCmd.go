////////////////////////////////////////////////////////////////////////////
// Porgram: cgrCmd - Config Group Report command handling
// Authors: Antonio Sun (c) 2015-16, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"strings"
	"time"
)

func cgrCmd(options Options) error {
	t := time.Now()
	fmt.Printf("h2. Server Listing as of %v\r\n\r\n", t.Format("2006-01-02"))
	fmt.Printf("||Ver||Instance||Servers||\r\n")
	for _, pods := range config.Pod {
		fmt.Printf("||%s|| | |\r\n", pods.Id)
		for _, inst := range pods.Instance {
			fmt.Printf("|| || %s | ", inst.Database)
			fmt.Printf(" * %s |\r\n",
				strings.Join(strings.Fields(inst.Dbserver+" "+inst.Servers), "\r\n * "))
		}
	}
	return nil
}
