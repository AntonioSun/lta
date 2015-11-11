////////////////////////////////////////////////////////////////////////////
// Porgram: cmd_rd - Result Dump handling
// Authors: Antonio Sun (c) 2015, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

func cmd_rd(options Options) error {
	PerfCounterExport(options, options.Rd.Id,
		options.Rd.MachineNameFilter, options.Rd.PathOut)
	return nil
}
