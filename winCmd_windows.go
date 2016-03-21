////////////////////////////////////////////////////////////////////////////
// Porgram: rdCmd - Result Dump command handling
// Authors: Antonio Sun (c) 2015-16, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"os/exec"
	"strings"
	"time"
)

func rdCmd(options Options) error {
	progname += "::rd"
	PerfCounterExport(options, options.Rd.Id,
		options.Rd.MachineNameFilter, options.Rd.PathOut)
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
				fmt.Printf("%s: Wait 3 minutes for them to come back...\r\n\r\n",
					progname)
				time.Sleep(time.Minute * 3)
				for _, m := range strings.Fields(inst.Servers) {
					ping(m)
				}
			}
		}
	}
	return ret
}

func reboot(machine string) {
	fmt.Printf("%s: Rebooting %s\r\n", progname, machine)

	_cmd := "shutdown"
	args := append([]string{},
		"-t", "0",
		"-r",
		"-f",
		"-m",
	)
	args = append(args, `\\`+machine)

	cmd := exec.Command(_cmd, args...)
	out, err := cmd.Output()
	if err != nil {
		println(err.Error())
	}

	print(string(out))
}

func ping(machine string) {
	fmt.Printf("%s: Pinging %s\r\n", progname, machine)

	_cmd := "ping"
	args := []string{}
	args = append(args, machine)

	cmd := exec.Command(_cmd, args...)
	out, err := cmd.Output()
	if err != nil {
		println(err.Error())
	}

	println(string(out))
}
