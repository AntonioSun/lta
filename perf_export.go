////////////////////////////////////////////////////////////////////////////
// Porgram: perf_export.go - Perf Counter Export
// Authors: Antonio Sun (c) 2015, All rights reserved
// Purpose: Export performance counters collected from MS load test to .csv
//			    files for perfmon to view
////////////////////////////////////////////////////////////////////////////

// Translated to GO from C#, http://blogs.msdn.com/b/geoffgr/archive/2013/09/09/

// +build windows

package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"bitbucket.org/kardianos/table"
	_ "github.com/alexbrainman/odbc"
)

/*
PerfCounterExport will wxport performance counters collected from MS load
test to .csv files for perfmon to view
*/
func PerfCounterExport(options Options, ltRunId int, machineNameFilter string, pathOut string) {
	conn := getConnection(getConnectionString(options))
	defer conn.Close()
	log.Printf("[%s] Program started\n", progname)

	// get TraceName according to LoadTestRunId
	r, err := table.Get(conn,
		"SELECT TraceName from LoadTestRun WHERE LoadTestRunId = ?", ltRunId)
	check(err)
	ltTraceName := U8ToGoString(r.MustGetScaler(0, "TraceName").([]uint8))

	resultFilePre := pathOut + string(os.PathSeparator)
	resultFilePre = filepath.Dir(resultFilePre) +
		string(os.PathSeparator) + ltTraceName
	os.Mkdir(resultFilePre, os.ModePerm)
	// so far path only, now append folder name as file prefix
	resultFilePre += string(os.PathSeparator) + ltTraceName

	log.Printf("[%s] Exporting LoadTest %d\n  to %s-...\n  with step of %d\n",
		progname, ltRunId, resultFilePre, options.Step)

	if machineNameFilter != "" {
		fmt.Printf("  limiting to only export machine %s\n\n", machineNameFilter)
		savePerfmonAsCsv(options.NoClobber, conn, machineNameFilter, ltRunId, resultFilePre)
		os.Exit(0)
	}

	/*
			Get all machine names

		    SELECT  category.MachineName
		      FROM  LoadTestPerformanceCounterCategory AS category
		      JOIN  LoadTestPerformanceCounterInstance AS instance
		        ON  category.LoadTestRunId = instance.LoadTestRunId
		       AND  instance.LoadTestRunId = (
		            SELECT MAX(LoadTestRunId) from LoadTestRun )
		     GROUP  BY MachineName

	*/

	machines, err := table.Get(conn,
		"SELECT  category.MachineName"+
			"  FROM  LoadTestPerformanceCounterCategory AS category"+
			"  JOIN  LoadTestPerformanceCounterInstance AS instance"+
			"    ON  category.LoadTestRunId = instance.LoadTestRunId"+
			"   AND  instance.LoadTestRunId = ?"+
			" GROUP  BY MachineName", ltRunId)
	if err != nil {
		log.Fatal(err)
	}

	for _, machine := range machines.Rows {
		// machine.MustGet("MachineName").(string)
		machineName := U8ToGoString(machine.MustGet("MachineName").([]uint8))
		savePerfmonAsCsv(options.NoClobber, conn, machineName, ltRunId, resultFilePre)
	}

	log.Printf("[%s] Exporting finished correctly.\n", progname)
	return
}

func getConnection(connectionString string) *sql.DB {
	conn, err := sql.Open("odbc", connectionString)
	check(err)
	return conn
}

func getConnectionString(options Options) string {
	// Construct the Go MSSQL odbc SqlConnectionString
	// https://code.google.com/p/odbc/source/browse/mssql_test.go
	var c string
	if options.SqlConnectionString == "" {
		var params map[string]string
		params = map[string]string{
			"driver":             "sql server",
			"server":             options.Server,
			"database":           options.PerfDb,
			"trusted_connection": "yes",
		}

		for n, v := range params {
			c += n + "=" + v + ";"
		}
	} else {
		c = options.SqlConnectionString
	}
	//log.Println("Connection string: " + c)
	return c
}

func savePerfmonAsCsv(fNoClobber bool, conn *sql.DB, machineName string, _runId int, resultFilePre string) {
	// Only use right(5)
	const keep = 5
	if len(machineName) > keep {
		machineName = machineName[len(machineName)-keep:]
	}

	log.Printf("[%s]   Collecting data for %s...\n", progname, machineName)

	// if no clobber and the destination file exists, skip
	if options.NoClobber {
		if _, err := os.Stat(resultFilePre + "-" + machineName + ".csv"); err == nil {
			log.Printf("[%s]   (Host %s skipped for no clobbering)\n",
				progname, machineName)
			return
		}
	}

	sql := fmt.Sprintf("exec TSL_prc_PerfCounterCollectionInCsvFormat"+
		" @RunId = %d, @InstanceName=N'\\\\%%%s\\%%'", _runId, machineName)
	//log.Println("] sql string: " + sql)
	table, err := table.Get(conn, sql)
	if err != nil {
		log.Printf("[%s]   Skipping it for the fatal error:\n\t\t    %v\n",
			progname, err.Error())
		return
	}

	log.Printf("[%s]   Exporting %s data...\n", progname, machineName)

	// open the output file
	file, err := os.Create(resultFilePre + "-" + machineName + ".csv")
	if err != nil {
		panic(err)
	}
	// close file on exit and check for its returned error
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	// output header
	for i, element := range table.ColumnName {
		if i != 0 {
			fmt.Fprintf(file, ",")
		}
		fmt.Fprintf(file, "\"%s\"", element)
	}
	fmt.Fprintf(file, "\n")

	// output body
	const layout = "01/02/2006 15:04:05.999"
	for j, row := range table.Rows {
		for i, colname := range table.ColumnName {
			if i != 0 {
				fmt.Fprintf(file, ",")
			}
			switch x := row.MustGet(colname).(type) {
			case string: // x is a string
				fmt.Fprintf(file, "\"%s\"", x)
			case int: // now x is an int
				fmt.Fprintf(file, "\"%d\"", x)
			case int32: // now x is an int32
				fmt.Fprintf(file, "\"%d\"", x)
			case int64: // now x is an int64
				fmt.Fprintf(file, "\"%d\"", x)
			case float32: // now x is an float32
				fmt.Fprintf(file, "\"%f\"", x)
			case float64: // now x is an float64
				fmt.Fprintf(file, "\"%f\"", x)
			case time.Time: // now x is a time.Time
				fmt.Fprintf(file, "\"%s\"", x.Format(layout))
			default:
				fmt.Fprintf(file, "\"%s\"", x)
			}
		}
		fmt.Fprintf(file, "\n")
		if j%options.Step == 0 {
			fmt.Fprintf(os.Stderr, ".")
		}
	}
	fmt.Fprintf(os.Stderr, "\n")

}

func U8ToGoString(c []uint8) string {
	n := -1
	for i, b := range c {
		if b == 0 {
			break
		}
		n = i
	}
	return string(c[:n+1])
}
