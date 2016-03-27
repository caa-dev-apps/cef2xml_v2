package main

import (
	"encoding/xml"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

type CAA_MetaData struct {
	XMLName xml.Name `xml:"CAA_METADATA"`

	MISSION               string   `xml:"MISSION_METADATA>MISSION"`
	MISSION_TIME_SPAN     string   `xml:"MISSION_METADATA>MISSION_TIME_SPAN"`
	MISSION_AGENCY        string   `xml:"MISSION_METADATA>MISSION_AGENCY"`
	MISSION_DESCRIPTION   string   `xml:"MISSION_METADATA>MISSION_DESCRIPTION"`
	MISSION_KEY_PERSONNEL []string `xml:"MISSION_METADATA>MISSION_KEY_PERSONNEL"`
	MISSION_REFERENCES    string   `xml:"MISSION_METADATA>MISSION_REFERENCES"`
	MISSION_REGION        []string `xml:"MISSION_METADATA>MISSION_REGION"`
	MISSION_CAVEATS       []string `xml:"MISSION_METADATA>MISSION_CAVEATS"`

	OBSERVATORIES Observatories `xml:"MISSION_METADATA>OBSERVATORIES"`
	EXPERIMENTS   Experiments   `xml:"MISSION_METADATA>EXPERIMENTS"`
	DATASETS      Datasets      `xml:"MISSION_METADATA>DATASETS"`
}

type Observatories struct {
	OBSERVATORY             []string `xml:"OBSERVATORY_METADATA>OBSERVATORY"`
	OBSERVATORY_CAVEATS     []string `xml:"OBSERVATORY_METADATA>OBSERVATORY_CAVEATS"`
	OBSERVATORY_DESCRIPTION string   `xml:"OBSERVATORY_METADATA>OBSERVATORY_DESCRIPTION"`
	OBSERVATORY_TIME_SPAN   string   `xml:"OBSERVATORY_METADATA>OBSERVATORY_TIME_SPAN"`
	OBSERVATORY_REGION      []string `xml:"OBSERVATORY_METADATA>OBSERVATORY_REGION"`
}

type Experiments struct {
	EXPERIMENT             string `xml:"EXPERIMENT_METADATA>EXPERIMENT"`
	EXPERIMENT_DESCRIPTION string `xml:"EXPERIMENT_METADATA>EXPERIMENT_DESCRIPTION"`

	INVESTIGATOR_COORDINATES []string `xml:"EXPERIMENT_METADATA>INVESTIGATOR_COORDINATES"`
	EXPERIMENT_REFERENCES    []string `xml:"EXPERIMENT_METADATA>EXPERIMENT_REFERENCES"`
	EXPERIMENT_KEY_PERSONNEL []string `xml:"EXPERIMENT_METADATA>EXPERIMENT_KEY_PERSONNEL"`
	EXPERIMENT_CAVEATS       string   `xml:"EXPERIMENT_METADATA>EXPERIMENT_CAVEATS"`

	INSTRUMENTS Instruments `xml:"EXPERIMENT_METADATA>INSTRUMENTS"`
}

type Instruments struct {
	INSTRUMENT_NAME        []string `xml:"INSTRUMENT_METADATA>INSTRUMENT_NAME"`
	INSTRUMENT_DESCRIPTION string   `xml:"INSTRUMENT_METADATA>INSTRUMENT_DESCRIPTION"`
	INSTRUMENT_TYPE        []string `xml:"INSTRUMENT_METADATA>INSTRUMENT_TYPE"`
	MEASUREMENT_TYPE       []string `xml:"INSTRUMENT_METADATA>MEASUREMENT_TYPE"`
	INSTRUMENT_CAVEATS     []string `xml:"INSTRUMENT_METADATA>INSTRUMENT_CAVEATS"`
}

type Datasets struct {
	DATASET_ID          string `xml:"DATASET_METADATA>DATASET_ID"`
	DATA_TYPE           string `xml:"DATASET_METADATA>DATA_TYPE"`
	DATASET_TITLE       string `xml:"DATASET_METADATA>DATASET_TITLE"`
	DATASET_DESCRIPTION string `xml:"DATASET_METADATA>DATASET_DESCRIPTION"`

	CONTACT_COORDINATES []string `xml:"DATASET_METADATA>CONTACT_COORDINATES"`

	PROCESSING_LEVEL    string `xml:"DATASET_METADATA>PROCESSING_LEVEL"`
	TIME_RESOLUTION     string `xml:"DATASET_METADATA>TIME_RESOLUTION"`
	MIN_TIME_RESOLUTION string `xml:"DATASET_METADATA>MIN_TIME_RESOLUTION"`
	MAX_TIME_RESOLUTION string `xml:"DATASET_METADATA>MAX_TIME_RESOLUTION"`
	DATASET_CAVEATS     string `xml:"DATASET_METADATA>DATASET_CAVEATS"`
	ACKNOWLEDGEMENT     string `xml:"DATASET_METADATA>ACKNOWLEDGEMENT"`

	PARAMETERS Parameters `xml:"DATASET_METADATA>PARAMETERS"`
	FILE       File       `xml:"DATASET_METADATA>FILE"`

	UNEXPECTED Unexpected `xml:"UNEXPECTED"`
}

type Parameters struct {
	PARAMETER []Parameter `xml:"PARAMETER"`
}

type Parameter struct {
	PARAMETER_ID       string `xml:"PARAMETER_ID,omitempty"`
	PARAMETER_TYPE     string `xml:"PARAMETER_TYPE,omitempty"`
	CATDESC            string `xml:"CATDESC,omitempty"`
	UNITS              string `xml:"UNITS,omitempty"`
	SI_CONVERSION      string `xml:"SI_CONVERSION,omitempty"`
	SIZES              string `xml:"SIZES,omitempty"`
	VALUE_TYPE         string `xml:"VALUE_TYPE,omitempty"`
	SIGNIFICANT_DIGITS string `xml:"SIGNIFICANT_DIGITS,omitempty"`
	FILLVAL            string `xml:"FILLVAL,omitempty"`
	FIELDNAM           string `xml:"FIELDNAM,omitempty"`
	LABLAXIS           string `xml:"LABLAXIS,omitempty"`
	DELTA_PLUS         string `xml:"DELTA_PLUS,omitempty"`
	DELTA_MINUS        string `xml:"DELTA_MINUS,omitempty"`

	ENTITY   string `xml:"ENTITY,omitempty"`
	PROPERTY string `xml:"PROPERTY,omitempty"`
	QUALITY  string `xml:"QUALITY,omitempty"`
	DEPEND_0 string `xml:"DEPEND_0,omitempty"`
	DEPEND_1 string `xml:"DEPEND_1,omitempty"`
	DEPEND_2 string `xml:"DEPEND_2,omitempty"`
	DEPEND_3 string `xml:"DEPEND_3,omitempty"`

	FRAME             string `xml:"FRAME,omitempty"`
	TENSOR_ORDER      string `xml:"TENSOR_ORDER,omitempty"`
	COORDINATE_SYSTEM string `xml:"COORDINATE_SYSTEM,omitempty"`
	FRAME_VELOCITY    string `xml:"FRAME_VELOCITY,omitempty"`
	REPRESENTATION_0  string `xml:"REPRESENTATION_0,omitempty"`
	REPRESENTATION_1  string `xml:"REPRESENTATION_1,omitempty"`
	REPRESENTATION_2  string `xml:"REPRESENTATION_2,omitempty"`
	REPRESENTATION_3  string `xml:"REPRESENTATION_3,omitempty"`
	REPRESENTATION_4  string `xml:"REPRESENTATION_4,omitempty"`

	LABEL_0 string `xml:"LABEL_0,omitempty"`
	LABEL_1 string `xml:"LABEL_1,omitempty"`
	LABEL_2 string `xml:"LABEL_2,omitempty"`
	LABEL_3 string `xml:"LABEL_3,omitempty"`
	LABEL_4 string `xml:"LABEL_4,omitempty"`

	// Guessed that they belong here
	SCALEMIN string `xml:"SCALEMIN,omitempty"`
	SCALEMAX string `xml:"SCALEMAX,omitempty"`
	SCALETYP string `xml:"SCALETYP,omitempty"`

	DISPLAYTYPE string `xml:"DISPLAYTYPE,omitempty"`
	DATA        string `xml:"DATA,omitempty"`
	//
	PARAMETER_CAVEATS string `xml:"PARAMETER_CAVEATS,omitempty"`
	FLUCTUATIONS      string `xml:"FLUCTUATIONS,omitempty"`
	TENSOR_FRAME      string `xml:"TENSOR_FRAME,omitempty"`
    
    ERROR_PLUS          string `xml:"ERROR_PLUS,omitempty"`
    ERROR_MINUS         string `xml:"ERROR_MINUS,omitempty"`
    
}

type File struct {
	FILE_TYPE        string `xml:"FILE_TYPE,omitempty"`
	METADATA_TYPE    string `xml:"METADATA_TYPE,omitempty"`
	METADATA_VERSION string `xml:"METADATA_VERSION,omitempty"`
	LOGICAL_FILE_ID  string `xml:"LOGICAL_FILE_ID,omitempty"`
	VERSION_NUMBER   string `xml:"VERSION_NUMBER,omitempty"`
	DATASET_VERSION  string `xml:"DATASET_VERSION,omitempty"`
	FILE_CAVEATS     string `xml:"FILE_CAVEATS,omitempty"`
	FILE_TIME_SPAN   string `xml:"FILE_TIME_SPAN,omitempty"`
	GENERATION_DATE  string `xml:"GENERATION_DATE,omitempty"`

	// guessed the following go here!
	PARENT_DATASET string `xml:"PARENT_DATASET,omitempty"`
	INGESTION_DATE string `xml:"INGESTION_DATE,omitempty"`
	FILE_SIZE      string `xml:"FILE_SIZE,omitempty"`

	FILE_NAME            string `xml:"FILE_NAME,omitempty"`
	FILE_FORMAT_VERSION  string `xml:"FILE_FORMAT_VERSION,omitempty"`
	END_OF_RECORD_MARKER string `xml:"END_OF_RECORD_MARKER,omitempty"`
	DATA_UNTIL           string `xml:"DATA_UNTIL,omitempty"`
}

type Unexpected struct {
	ATTR  []TypeKeyValue `xml:"ATTR>A"`
	META  []TypeKeyValue `xml:"META>M"`
	VAR   []TypeKeyValue `xml:"VAR>V"`
	ERROR []TypeKeyValue `xml:"ERROR>E"`
}

type TypeKeyValue struct {
	Key string `xml:"key,attr"`
	Val string `xml:"val,attr"`
}

func debug_show_results(output []byte) (err error) {

	s1 := time.Now().Format(time.RFC3339)
	s1 = strings.Replace(s1, ":", "-", -1)

	l_filepath := "C:/Dump/" + s1 + "-Test.xml"

	f, err := os.Create(l_filepath)
	if err != nil {
		panic("file create error\n")
	}
	defer f.Close()

	f.Write(output)
	err = exec.Command("open", l_filepath).Start()

	return
}

func (m *CAA_MetaData) dump() error {

	output, err := xml.MarshalIndent(m, "", "  ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	os.Stdout.Write(output)

	//x debug_show_results(output)

	return err
}
