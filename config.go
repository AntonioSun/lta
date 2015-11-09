////////////////////////////////////////////////////////////////////////////
// Porgram: config - Config handling
// authors: Antonio Sun (c) 2015, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"io/ioutil"
)

import (
	"gopkg.in/yaml.v2"
)

////////////////////////////////////////////////////////////////////////////
// Config

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
var config map[interface{}]interface{}

func configGet(configFile string) {

	cfgStr, err := ioutil.ReadFile(options.ConfigFile)
	err = yaml.Unmarshal(cfgStr, &config)
	check(err)
	fmt.Printf("] %#v\r\n", config)
}
