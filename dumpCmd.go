////////////////////////////////////////////////////////////////////////////
// Porgram: dumpCmd - loadtest file dump handling
// Authors: Antonio Sun (c) 2016, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

////////////////////////////////////////////////////////////////////////////
// Function definitions

//::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
// Main dispatch function

func dumpCmd(options Options) error {
	fileo := options.Dump.Fileo
	if fileo == nil {
		var err error
		fileo, err = os.Create(
			strings.Replace(options.Dump.Filei.Name(), ".loadtest", ".loadtext", 1))
		check(err)
	}
	//println("] ", fileo.Name())
	defer fileo.Close()

	return treatLtXml(fileo, getDecoder(options.Dump.Filei))
}

//::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
// Script-wide processing

func treatLtXml(w io.Writer, decoder *xml.Decoder) error {

	for {
		// Read tokens from the XML document in a stream.
		token, _ := decoder.Token()
		if token == nil {
			break
		}

		// Inspect the type of the token just read.
		switch t := token.(type) {
		case xml.StartElement:
			// NB, even we're just dealing with a StartElement token
			// the whole XML has already been read into the token
			switch inElement := t.Name.Local; inElement {
			case "Scenario":
				{
					var r Scenario
					// decode a whole chunk of following XML into the
					// variable c which is a Comment (t above)
					decoder.DecodeElement(&r, &t)
					fmt.Fprintf(w,
						"\r\n\r\nSC: %s, %s mix, %s delay, %s gap, %s max, %s%% new\r\n",
						r.Name, r.TestMixType, r.DelayStartTime, r.DelayBetweenIterations,
						r.MaxTestIterations, r.PercentNewUsers)
					for _, av := range r.TestMix.TestProfile {
						if options.Dump.ID {
							fmt.Fprintf(w, "  TP: %s at %s%%.\tId=%s\r\n",
								av.Name, av.Percentage, av.Id)
						} else {
							fmt.Fprintf(w, "  TP: %s\tat %s%%\r\n",
								av.Name, av.Percentage)
						}
					}
					r.LoadProfile.XMLName = xml.Name{}
					fmt.Fprintf(w, "  LP: %+v\r\n", r.LoadProfile)
					fmt.Fprintf(w, "  TK: %s\r\n", r.ThinkProfile.Pattern)
				}
			case "RunConfiguration":
				{
					var r RunConfiguration
					decoder.DecodeElement(&r, &t)
					fmt.Fprintf(w, "\r\nRC: %s, %s:%s:%s dur, %s up, %s srate\r\n",
						r.Name, r.TestIterations, r.UseTestIterations, r.RunDuration,
						r.WarmupTime, r.SampleRate)
					for _, av := range r.CounterSetMappings.CounterSetMapping {
						if av.ComputerName[0] != '[' {
							fmt.Fprintf(w, "  SV: %s\r\n", av.ComputerName)
						}
					}
					fmt.Fprintf(w, "  --:\r\n")
					for _, av := range r.ContextParameters.ContextParameter {
						fmt.Fprintf(w, "  CP: %s=%s\r\n", av.Name, av.Value)
					}
				}
			case "LoadTestPlugin":
				{
					var r LoadTestPlugin
					decoder.DecodeElement(&r, &t)
					fmt.Fprintf(w, "\r\nPL: (%s)\r\n  RP: %s\r\n",
						r.DisplayName, minify(r.RuleParameters.Xml))
				}
			default:
			}
		}
	}
	return nil
}

//::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
// Support functions

func getDecoder(Script *os.File) *xml.Decoder {
	defer Script.Close()

	content, err := ioutil.ReadFile(Script.Name())
	check(err)
	return xml.NewDecoder(bytes.NewBuffer(content))
}

func minify(xs string) string {
	re := regexp.MustCompile("\r*\n *")
	return re.ReplaceAllString(xs, "")
}
