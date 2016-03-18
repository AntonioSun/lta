////////////////////////////////////////////////////////////////////////////
// Porgram: rdCmd - Result Dump command handling
// Authors: Antonio Sun (c) 2015-16, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

func rdCmd(options Options) error {
	PerfCounterExport(options, options.Rd.Id,
		options.Rd.MachineNameFilter, options.Rd.PathOut)
	return nil
}
