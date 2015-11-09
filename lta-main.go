////////////////////////////////////////////////////////////////////////////
// Porgram: lta-main - load test assistant, main
// Purpose: load test assistant kit
// authors: Antonio Sun (c) 2015, All rights reserved
// Credits: https://github.com/voxelbrain/goptions/tree/master/examples
//
//
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"os"
)

import (
	"github.com/voxelbrain/goptions"
)

//	"os"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Options struct {
	Server string `goptions:"--cs, description='Connection Server, Server of PerfCounter info from\n\t\t\t\tDefault: local machine'"`

	PerfDb string `goptions:"--cd, description='Connection DB, DB that holds the PerfCounter info\n\t\t\t\t'"`

	SqlConnectionString string `goptions:"--conn, description='ConnectionString of Go MSSQL Odbc to MS SQL Server\n\t\t\t\tTo override the above --cs/cd setting. Sample:\n\t\t\t\tdriver=sql server;server=(local);database=LoadTest2010;uid=user;pwd=pass\n'"`

	Verbosity []bool        "goptions:\"-v, --verbose, description='Be verbose'\""
	Help      goptions.Help `goptions:"-h, --help, description='Show this help\n\nSub-commands (Verbs):\n\n\tcgl\t\tConfig Group List\n\t\t\tList machine groups defined in config file\n\n\trd\t\tResult Dump\n\t\t\tDump load test result, all machine counters\n\trdg\t\tResult Dump Group\n\t\t\tDump load test results, for the machine group\n\n\trbg\t\tReBoot Group\n\t\t\tReboot the machine group'"`

	goptions.Verbs

	cgl struct{} `goptions:"cgl"`

	rd  struct{} `goptions:"rd"`
	rdg struct{} `goptions:"rdg"`

	rbg struct{} `goptions:"rbg"`
}

var options = Options{ // Default values goes here
	Server: "(local)",
	PerfDb: "LoadTest2010",
}

type Command func(Options) error

var commands = map[goptions.Verbs]Command{
	"cgl": cmd_cgl,
}

var (
	VERBOSITY = 0
)

func main() {
	goptions.ParseAndFail(&options)
	//fmt.Printf("] %#v\r\n", options)

	if len(options.Verbs) == 0 {
		goptions.PrintHelp()
		os.Exit(2)
	}

	VERBOSITY = len(options.Verbosity)

	if cmd, found := commands[options.Verbs]; found {
		err := cmd(options)
		if err != nil {
			fmt.Println("error:", err)
			os.Exit(1)
		}
	}
}

func cmd_cgl(options Options) error {
	return nil
}
