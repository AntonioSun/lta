////////////////////////////////////////////////////////////////////////////
// Porgram: winCmd - Windows specific command handling in dummy linux version
// Authors: Antonio Sun (c) 2015-16, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"strings"
)

func rdCmd(options Options) error {
	return nil
}

func rebootCmd(options Options) error {
	var ret error
	progname += "::reboot"
	for _, pods := range config.Pod {
		for _, inst := range pods.Instance {
			if inst.Database == options.Reboot.Instance {
				fmt.Printf("%s: Rebooting machines for %s\r\n", progname, inst.Database)
				if !options.Reboot.NoDb {
					reboot(inst.Dbserver)
				}
				for _, m := range strings.Fields(inst.Servers) {
					reboot(m)
				}
			}
		}
	}
	return ret
}

func reboot(machine string) {
	fmt.Printf("%s: Rebooting %s\r\n", progname, machine)
}
