////////////////////////////////////////////////////////////////////////////
// Porgram: lta-main - load test assistant, main
// Purpose: load test assistant kit
// authors: Antonio Sun (c) 2015-16, All rights reserved
// Credits: https://github.com/voxelbrain/goptions/tree/master/examples
//
//
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"os"
)

import "github.com/voxelbrain/goptions"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

////////////////////////////////////////////////////////////////////////////
// Configuration variables definitions

var progname = "lta"
var progdesc = " - load test assistant program"
var buildTime = "2016-03-18"

////////////////////////////////////////////////////////////////////////////
// Commandline option definitions

type Options struct {
	ConfigExt  string `goptions:"--cfe, description='Config File Extension, extension for the config file\n\t\t\t\tDefault config file will be Args[0] plus this extension\n\t\t\t\t'"`
	ConfigFile string `goptions:"--cfn, description='Config File Name, the alternative config file to use\n\t\t\t\tinstead of the default one, with extension\n'"`

	Server string `goptions:"--cs, description='Connection Server, Server of PerfCounter info from\n\t\t\t\tDefault: local machine'"`

	PerfDb string `goptions:"--cd, description='Connection DB, DB that holds the PerfCounter info\n\t\t\t\t'"`

	SqlConnectionString string `goptions:"--conn, description='ConnectionString of Go MSSQL Odbc to MS SQL Server\n\t\t\t\tTo override the above --cs/cd setting. Sample:\n\t\t\t\tdriver=sql server;server=(local);database=LoadTest2010;uid=user;pwd=pass\n'"`

	Step      int  `goptions:"-s, --step, description='Number of record outputed after which indicator is shown\n\t\t\t\t'"`
	NoClobber bool `goptions:"--nc, description='No clobber, do not overwrite existing files\n\t\t\t\tDefault: overwrite them\n'"`

	Verbosity []bool        `goptions:"-v, --verbose, description='Be verbose'"`
	Help      goptions.Help `goptions:"-h, --help, description='Show this help\n\nSub-commands (Verbs):\n\n\tcgl\t\tConfig Group List\n\t\t\tList machine groups defined in config file\n\n\tcgr\t\tConfig Group Report\n\t\t\tReport machines by groups from config file as Confluence Wiki\n\n\trd\t\tResult Dump\n\t\t\tDump load test result\n\n\tdump\t\tDump\n\t\t\tDump loadtest file\n\n\treboot\t\tReboot instance\n\t\t\tReboot the machines defined as the instance in config file'"`

	goptions.Verbs

	Cgl struct{} `goptions:"cgl"`

	Cgr struct{} `goptions:"cgr"`

	Rd struct {
		Id                int    `goptions:"-n, --id, obligatory, description='Loadtest RunId'"`
		PathOut           string `goptions:"-p, --path, obligatory, description='Path to where dumps are saved'"`
		MachineNameFilter string `goptions:"-m, --mfilter, description='machine name filter for exporting the counters\n\t\t\t\tDefault: export all machines\n'"`
	} `goptions:"rd"`

	Dump struct {
		Filei *os.File `goptions:"-i, --input, obligatory, description='The loadtest file to dump', rdonly"`
		Fileo *os.File `goptions:"-o, --output, description='The loadtest dump output (default: .loadtext file of input)', wronly"`
		ID    bool     `goptions:"--id, description='\tDump loadtest ID as well'"`
	} `goptions:"dump"`

	Reboot struct {
		NoDb bool `goptions:"--nodb, description='\tReboot all machines except the DB server'"`
	} `goptions:"reboot"`
}

var options = Options{ // Default values goes here
	ConfigExt: ".conf",
	Server:    "(local)",
	PerfDb:    "LoadTest2010",
	Step:      50,
}

type Command func(Options) error

var commands = map[goptions.Verbs]Command{
	"cgl":  cglCmd,
	"cgr":  cgrCmd,
	"rd":   rdCmd,
	"dump": dumpCmd,
}

var (
	VERBOSITY = 0
)

////////////////////////////////////////////////////////////////////////////
// Main

func main() {
	goptions.ParseAndFail(&options)
	//fmt.Printf("] %#v\r\n", options)

	if len(options.Verbs) == 0 {
		fmt.Printf("%s%s \n      built on %s\n\n", progname, progdesc, buildTime)
		goptions.PrintHelp()
		os.Exit(2)
	}

	VERBOSITY = len(options.Verbosity)

	configGet(os.Args[0])

	if cmd, found := commands[options.Verbs]; found {
		err := cmd(options)
		check(err)
	}
}
