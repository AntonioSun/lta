////////////////////////////////////////////////////////////////////////////
// Porgram: ltXml - loadtest XML file structs
// Authors: Antonio Sun (c) 2016, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import "encoding/xml"

type Xml struct {
	Xml string `xml:",innerxml"`
}

type XmlBase struct {
	Name           string `xml:"DisplayName,attr"`
	RuleParameters Xml
}

/*
type LoadTest struct {
	Name                    string `xml:" Name,attr"  json:",omitempty"`
	Description             string `xml:" Description,attr"  json:",omitempty"`
	Owner                   string `xml:" Owner,attr"  json:",omitempty"`
	storage                 string `xml:" storage,attr"  json:",omitempty"`
	Priority                string `xml:" Priority,attr"  json:",omitempty"`
	Enabled                 string `xml:" Enabled,attr"  json:",omitempty"`
	CssProjectStructure     string `xml:" CssProjectStructure,attr"  json:",omitempty"`
	CssIteration            string `xml:" CssIteration,attr"  json:",omitempty"`
	DeploymentItemsEditable string `xml:" DeploymentItemsEditable,attr"  json:",omitempty"`
	WorkItemIds             string `xml:" WorkItemIds,attr"  json:",omitempty"`
	TraceLevel              string `xml:" TraceLevel,attr"  json:",omitempty"`
	CurrentRunConfig        string `xml:" CurrentRunConfig,attr"  json:",omitempty"`
	Id                      string `xml:" Id,attr"  json:",omitempty"`
	xmlns                   string `xml:" xmlns,attr"  json:",omitempty"`
	//CounterSets             *CounterSets       `xml:"http://microsoft.com/schemas/VisualStudio/TeamTest/2010 CounterSets,omitempty" json:"CounterSets,omitempty"`
	//LoadTestPlugins *LoadTestPlugins `xml:"http://microsoft.com/schemas/VisualStudio/TeamTest/2010 LoadTestPlugins,omitempty" json:"LoadTestPlugins,omitempty"`
	//RunConfigurations *RunConfigurations `xml:"http://microsoft.com/schemas/VisualStudio/TeamTest/2010 RunConfigurations,omitempty" json:"RunConfigurations,omitempty"`
	//Scenarios         *Scenarios         `xml:"http://microsoft.com/schemas/VisualStudio/TeamTest/2010 Scenarios,omitempty" json:"Scenarios,omitempty"`
	XMLName xml.Name `xml:"http://microsoft.com/schemas/VisualStudio/TeamTest/2010 LoadTest,omitempty" json:"LoadTest,omitempty"`
}
*/

type Scenario struct {
	Name                           string `xml:" Name,attr"  json:",omitempty"`
	DelayBetweenIterations         string `xml:" DelayBetweenIterations,attr"  json:",omitempty"`
	PercentNewUsers                string `xml:" PercentNewUsers,attr"  json:",omitempty"`
	IPSwitching                    string `xml:" IPSwitching,attr"  json:",omitempty"`
	TestMixType                    string `xml:" TestMixType,attr"  json:",omitempty"`
	ApplyDistributionToPacingDelay string `xml:" ApplyDistributionToPacingDelay,attr"  json:",omitempty"`
	MaxTestIterations              string `xml:" MaxTestIterations,attr"  json:",omitempty"`
	DisableDuringWarmup            string `xml:" DisableDuringWarmup,attr"  json:",omitempty"`
	DelayStartTime                 string `xml:" DelayStartTime,attr"  json:",omitempty"`
	AllowedAgents                  string `xml:" AllowedAgents,attr"  json:",omitempty"`
	//BrowserMix                     *BrowserMix   `xml:"http://microsoft.com/schemas/VisualStudio/TeamTest/2010 BrowserMix,omitempty" json:"BrowserMix,omitempty"`
	LoadProfile *LoadProfile `xml:"http://microsoft.com/schemas/VisualStudio/TeamTest/2010 LoadProfile,omitempty" json:"LoadProfile,omitempty"`
	//NetworkMix                     *NetworkMix   `xml:"http://microsoft.com/schemas/VisualStudio/TeamTest/2010 NetworkMix,omitempty" json:"NetworkMix,omitempty"`
	TestMix      *TestMix      `xml:"http://microsoft.com/schemas/VisualStudio/TeamTest/2010 TestMix,omitempty" json:"TestMix,omitempty"`
	ThinkProfile *ThinkProfile `xml:"http://microsoft.com/schemas/VisualStudio/TeamTest/2010 ThinkProfile,omitempty" json:"ThinkProfile,omitempty"`
	XMLName      xml.Name      `xml:"http://microsoft.com/schemas/VisualStudio/TeamTest/2010 Scenario,omitempty" json:"Scenario,omitempty"`
}

type TestMix struct {
	TestProfile []TestProfile `xml:"http://microsoft.com/schemas/VisualStudio/TeamTest/2010 TestProfile,omitempty" json:"TestProfile,omitempty"`
	XMLName     xml.Name      `xml:"http://microsoft.com/schemas/VisualStudio/TeamTest/2010 TestMix,omitempty" json:"TestMix,omitempty"`
}

type TestProfile struct {
	Name       string   `xml:" Name,attr"  json:",omitempty"`
	Path       string   `xml:" Path,attr"  json:",omitempty"`
	Id         string   `xml:" Id,attr"  json:",omitempty"`
	Percentage string   `xml:" Percentage,attr"  json:",omitempty"`
	Type       string   `xml:" Type,attr"  json:",omitempty"`
	XMLName    xml.Name `xml:"http://microsoft.com/schemas/VisualStudio/TeamTest/2010 TestProfile,omitempty" json:"TestProfile,omitempty"`
}

type ThinkProfile struct {
	Value   string   `xml:" Value,attr"  json:",omitempty"`
	Pattern string   `xml:" Pattern,attr"  json:",omitempty"`
	XMLName xml.Name `xml:"http://microsoft.com/schemas/VisualStudio/TeamTest/2010 ThinkProfile,omitempty" json:"ThinkProfile,omitempty"`
}

type LoadProfile struct {
	Pattern      string   `xml:" Pattern,attr"  json:",omitempty"`
	InitialUsers string   `xml:" InitialUsers,attr"  json:",omitempty"`
	MaxUsers     string   `xml:" MaxUsers,attr"  json:",omitempty"`
	StepUsers    string   `xml:" StepUsers,attr"  json:",omitempty"`
	StepDuration string   `xml:" StepDuration,attr"  json:",omitempty"`
	StepRampTime string   `xml:" StepRampTime,attr"  json:",omitempty"`
	XMLName      xml.Name `xml:"http://microsoft.com/schemas/VisualStudio/TeamTest/2010 LoadProfile,omitempty" json:"LoadProfile,omitempty"`
}

type RunConfiguration struct {
	Name                                string              `xml:" Name,attr"  json:",omitempty"`
	Description                         string              `xml:" Description,attr"  json:",omitempty"`
	ResultsStoreType                    string              `xml:" ResultsStoreType,attr"  json:",omitempty"`
	TimingDetailsStorage                string              `xml:" TimingDetailsStorage,attr"  json:",omitempty"`
	SaveTestLogsOnError                 string              `xml:" SaveTestLogsOnError,attr"  json:",omitempty"`
	SaveTestLogsFrequency               string              `xml:" SaveTestLogsFrequency,attr"  json:",omitempty"`
	MaxErrorDetails                     string              `xml:" MaxErrorDetails,attr"  json:",omitempty"`
	MaxErrorsPerType                    string              `xml:" MaxErrorsPerType,attr"  json:",omitempty"`
	MaxThresholdViolations              string              `xml:" MaxThresholdViolations,attr"  json:",omitempty"`
	MaxRequestUrlsReported              string              `xml:" MaxRequestUrlsReported,attr"  json:",omitempty"`
	UseTestIterations                   string              `xml:" UseTestIterations,attr"  json:",omitempty"`
	RunDuration                         string              `xml:" RunDuration,attr"  json:",omitempty"`
	WarmupTime                          string              `xml:" WarmupTime,attr"  json:",omitempty"`
	CoolDownTime                        string              `xml:" CoolDownTime,attr"  json:",omitempty"`
	TestIterations                      string              `xml:" TestIterations,attr"  json:",omitempty"`
	WebTestConnectionModel              string              `xml:" WebTestConnectionModel,attr"  json:",omitempty"`
	WebTestConnectionPoolSize           string              `xml:" WebTestConnectionPoolSize,attr"  json:",omitempty"`
	SampleRate                          string              `xml:" SampleRate,attr"  json:",omitempty"`
	ValidationLevel                     string              `xml:" ValidationLevel,attr"  json:",omitempty"`
	SqlTracingConnectString             string              `xml:" SqlTracingConnectString,attr"  json:",omitempty"`
	SqlTracingConnectStringDisplayValue string              `xml:" SqlTracingConnectStringDisplayValue,attr"  json:",omitempty"`
	SqlTracingDirectory                 string              `xml:" SqlTracingDirectory,attr"  json:",omitempty"`
	SqlTracingEnabled                   string              `xml:" SqlTracingEnabled,attr"  json:",omitempty"`
	SqlTracingFileCount                 string              `xml:" SqlTracingFileCount,attr"  json:",omitempty"`
	SqlTracingRolloverEnabled           string              `xml:" SqlTracingRolloverEnabled,attr"  json:",omitempty"`
	SqlTracingMinimumDuration           string              `xml:" SqlTracingMinimumDuration,attr"  json:",omitempty"`
	RunUnitTestsInAppDomain             string              `xml:" RunUnitTestsInAppDomain,attr"  json:",omitempty"`
	CoreCount                           string              `xml:" CoreCount,attr"  json:",omitempty"`
	ContextParameters                   *ContextParameters  `xml:"http://microsoft.com/schemas/VisualStudio/TeamTest/2010 ContextParameters,omitempty" json:"ContextParameters,omitempty"`
	CounterSetMappings                  *CounterSetMappings `xml:"http://microsoft.com/schemas/VisualStudio/TeamTest/2010 CounterSetMappings,omitempty" json:"CounterSetMappings,omitempty"`
	XMLName                             xml.Name            `xml:"http://microsoft.com/schemas/VisualStudio/TeamTest/2010 RunConfiguration,omitempty" json:"RunConfiguration,omitempty"`
}

type CounterSetMappings struct {
	CounterSetMapping []*CounterSetMapping `xml:"http://microsoft.com/schemas/VisualStudio/TeamTest/2010 CounterSetMapping,omitempty" json:"CounterSetMapping,omitempty"`
	XMLName           xml.Name             `xml:"http://microsoft.com/schemas/VisualStudio/TeamTest/2010 CounterSetMappings,omitempty" json:"CounterSetMappings,omitempty"`
}

type CounterSetMapping struct {
	ComputerName         string                `xml:" ComputerName,attr"  json:",omitempty"`
	CounterSetReferences *CounterSetReferences `xml:"http://microsoft.com/schemas/VisualStudio/TeamTest/2010 CounterSetReferences,omitempty" json:"CounterSetReferences,omitempty"`
	XMLName              xml.Name              `xml:"http://microsoft.com/schemas/VisualStudio/TeamTest/2010 CounterSetMapping,omitempty" json:"CounterSetMapping,omitempty"`
}

type CounterSetReferences struct {
	CounterSetReference []*CounterSetReference `xml:"http://microsoft.com/schemas/VisualStudio/TeamTest/2010 CounterSetReference,omitempty" json:"CounterSetReference,omitempty"`
	XMLName             xml.Name               `xml:"http://microsoft.com/schemas/VisualStudio/TeamTest/2010 CounterSetReferences,omitempty" json:"CounterSetReferences,omitempty"`
}

type CounterSetReference struct {
	CounterSetName string   `xml:" CounterSetName,attr"  json:",omitempty"`
	XMLName        xml.Name `xml:"http://microsoft.com/schemas/VisualStudio/TeamTest/2010 CounterSetReference,omitempty" json:"CounterSetReference,omitempty"`
}

type ContextParameters struct {
	ContextParameter []*ContextParameter `xml:"http://microsoft.com/schemas/VisualStudio/TeamTest/2010 ContextParameter,omitempty" json:"ContextParameter,omitempty"`
	XMLName          xml.Name            `xml:"http://microsoft.com/schemas/VisualStudio/TeamTest/2010 ContextParameters,omitempty" json:"ContextParameters,omitempty"`
}

type ContextParameter struct {
	Name    string   `xml:" Name,attr"  json:",omitempty"`
	Value   string   `xml:" Value,attr"  json:",omitempty"`
	XMLName xml.Name `xml:"http://microsoft.com/schemas/VisualStudio/TeamTest/2010 ContextParameter,omitempty" json:"ContextParameter,omitempty"`
}

type LoadTestPlugin struct {
	XmlBase
	Classname   string   `xml:" Classname,attr"  json:",omitempty"`
	DisplayName string   `xml:" DisplayName,attr"  json:",omitempty"`
	Description string   `xml:" Description,attr"  json:",omitempty"`
	XMLName     xml.Name `xml:"http://microsoft.com/schemas/VisualStudio/TeamTest/2010 LoadTestPlugin,omitempty" json:"LoadTestPlugin,omitempty"`
}
