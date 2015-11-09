////////////////////////////////////////////////////////////////////////////
// Porgram: config - Config handling
// Authors: Antonio Sun (c) 2015, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import "io/ioutil"

import (
	"gopkg.in/yaml.v2"
)

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

/*
dbuser: uu
dbpassword: "pp"

pod:

 - id: v746b

   instance:

    - database: perfwhit746
      dbserver: TorsvPerfDb07
      servers: TorsvPerfBje05 TorsvPerfBje06 TorsvPerfApp03 TorsvPerfApp06

    - database: perfwhit746b
      dbserver: TorsvPerfDb07
      servers: TorsvPerfBje05 TorsvPerfBje06 TorsvPerfApp03 TorsvPerfApp06

*/

type instance struct {
	Database string
	Dbserver string
	Servers  string
}

type pod struct {
	Id       string
	Instance []instance
}

var config struct {
	DbUser     string
	DbPassword string
	Pod        []pod
}

////////////////////////////////////////////////////////////////////////////
// Function definitions

func configGet(configFile string) {
	cfgStr, err := ioutil.ReadFile(options.ConfigFile)
	err = yaml.Unmarshal(cfgStr, &config)
	check(err)
	//fmt.Printf("] %#v\r\n", config)
}
