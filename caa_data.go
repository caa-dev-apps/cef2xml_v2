package main

import (
	"encoding/xml"
	"fmt"
 	"os"
    "time"
//x     "math/rand"
    "os/exec"
    "strings"
)

///////////////////////////////////////////////////////////////////////////////
//

type CAA_MetaData struct {
    XMLName                     xml.Name            `xml:"CAA_METADATA"`
                                              
    MISSION                     string              `xml:"MISSION_METADATA>MISSION"`
    MISSION_TIME_SPAN           string              `xml:"MISSION_METADATA>MISSION_TIME_SPAN"`
    MISSION_AGENCY              string              `xml:"MISSION_METADATA>MISSION_AGENCY"`
    MISSION_DESCRIPTION         string              `xml:"MISSION_METADATA>MISSION_DESCRIPTION"`
    MISSION_KEY_PERSONNEL       []string            `xml:"MISSION_METADATA>MISSION_KEY_PERSONNEL"`
    MISSION_REFERENCES          string              `xml:"MISSION_METADATA>MISSION_REFERENCES"`
    MISSION_REGION              []string            `xml:"MISSION_METADATA>MISSION_REGION"`
    MISSION_CAVEATS             []string            `xml:"MISSION_METADATA>MISSION_CAVEATS"`
                                              
    OBSERVATORIES               Observatories       `xml:"MISSION_METADATA>OBSERVATORIES"`
    EXPERIMENTS                 Experiments         `xml:"MISSION_METADATA>EXPERIMENTS"`
    DATASETS                    Datasets            `xml:"MISSION_METADATA>DATASETS"`
    
}                                             
                                              
type Observatories struct {                   
                                              
    OBSERVATORY                 []string            `xml:"OBSERVATORY_METADATA>OBSERVATORY"`
    OBSERVATORY_CAVEATS         []string            `xml:"OBSERVATORY_METADATA>OBSERVATORY_CAVEATS"`
    OBSERVATORY_DESCRIPTION     string              `xml:"OBSERVATORY_METADATA>OBSERVATORY_DESCRIPTION"`
    OBSERVATORY_TIME_SPAN       string              `xml:"OBSERVATORY_METADATA>OBSERVATORY_TIME_SPAN"`
    OBSERVATORY_REGION          []string            `xml:"OBSERVATORY_METADATA>OBSERVATORY_REGION"`
}                                             
                                              
type Experiments struct {                     
                                              
    EXPERIMENT                  string              `xml:"EXPERIMENT_METADATA>EXPERIMENT"`
    EXPERIMENT_DESCRIPTION      string              `xml:"EXPERIMENT_METADATA>EXPERIMENT_DESCRIPTION"`
                                              
    INVESTIGATOR_COORDINATES    []string            `xml:"EXPERIMENT_METADATA>INVESTIGATOR_COORDINATES"`
    EXPERIMENT_REFERENCES       []string            `xml:"EXPERIMENT_METADATA>EXPERIMENT_REFERENCES"`
    EXPERIMENT_KEY_PERSONNEL    []string            `xml:"EXPERIMENT_METADATA>EXPERIMENT_KEY_PERSONNEL"`
    EXPERIMENT_CAVEATS          string              `xml:"EXPERIMENT_METADATA>EXPERIMENT_CAVEATS"`
                                              
    INSTRUMENTS                 Instruments         `xml:"EXPERIMENT_METADATA>INSTRUMENTS"`
}                                             
                                              
type Instruments struct {                     
                                              
    INSTRUMENT_NAME             []string            `xml:"INSTRUMENT_METADATA>INSTRUMENT_NAME"`
//x INSTRUMENT_DESCRIPTION      []string            `xml:"INSTRUMENT_METADATA>INSTRUMENT_DESCRIPTION"`
    INSTRUMENT_DESCRIPTION      string              `xml:"INSTRUMENT_METADATA>INSTRUMENT_DESCRIPTION"`
    INSTRUMENT_TYPE             []string            `xml:"INSTRUMENT_METADATA>INSTRUMENT_TYPE"`
    MEASUREMENT_TYPE            []string            `xml:"INSTRUMENT_METADATA>MEASUREMENT_TYPE"`
    INSTRUMENT_CAVEATS          []string            `xml:"INSTRUMENT_METADATA>INSTRUMENT_CAVEATS"`

}

type Datasets struct {
    
    DATASET_ID                  string              `xml:"DATASET_METADATA>DATASET_ID"`
    DATA_TYPE                   string              `xml:"DATASET_METADATA>DATA_TYPE"`
    DATASET_TITLE               string              `xml:"DATASET_METADATA>DATASET_TITLE"`
    DATASET_DESCRIPTION         string              `xml:"DATASET_METADATA>DATASET_DESCRIPTION"`
                                           
    CONTACT_COORDINATES         []string            `xml:"DATASET_METADATA>CONTACT_COORDINATES"`
                                           
    PROCESSING_LEVEL            string              `xml:"DATASET_METADATA>PROCESSING_LEVEL"`
    TIME_RESOLUTION             string              `xml:"DATASET_METADATA>TIME_RESOLUTION"`
    MIN_TIME_RESOLUTION         string              `xml:"DATASET_METADATA>MIN_TIME_RESOLUTION"`
    MAX_TIME_RESOLUTION         string              `xml:"DATASET_METADATA>MAX_TIME_RESOLUTION"`
    DATASET_CAVEATS             string              `xml:"DATASET_METADATA>DATASET_CAVEATS"`
    ACKNOWLEDGEMENT             string              `xml:"DATASET_METADATA>ACKNOWLEDGEMENT"`
                                           
    PARAMETERS                  Parameters          `xml:"DATASET_METADATA>PARAMETERS"`
    FILE                        File                `xml:"DATASET_METADATA>FILE"`
    
    
    UNEXPECTED                  Unexpected           `xml:"UNEXPECTED"`
    
}                                          

type Parameters struct {                    
    PARAMETER                   []Parameter         `xml:"PARAMETER"`
}
                                           
type Parameter struct {                    

    PARAMETER_ID                string              `xml:"PARAMETER_ID,omitempty"`
    PARAMETER_TYPE              string              `xml:"PARAMETER_TYPE,omitempty"`
    CATDESC                     string              `xml:"CATDESC,omitempty"`
    UNITS                       string              `xml:"UNITS,omitempty"`
    SI_CONVERSION               string              `xml:"SI_CONVERSION,omitempty"`
    SIZES                       string              `xml:"SIZES,omitempty"`
    VALUE_TYPE                  string              `xml:"VALUE_TYPE,omitempty"`
    SIGNIFICANT_DIGITS          string              `xml:"SIGNIFICANT_DIGITS,omitempty"`
    FILLVAL                     string              `xml:"FILLVAL,omitempty"`
    FIELDNAM                    string              `xml:"FIELDNAM,omitempty"`
    LABLAXIS                    string              `xml:"LABLAXIS,omitempty"`
    DELTA_PLUS                  string              `xml:"DELTA_PLUS,omitempty"`
    DELTA_MINUS                 string              `xml:"DELTA_MINUS,omitempty"`
                                           
    ENTITY                      string              `xml:"ENTITY,omitempty"`
    PROPERTY                    string              `xml:"PROPERTY,omitempty"`
    QUALITY                     string              `xml:"QUALITY,omitempty"`
    DEPEND_0                    string              `xml:"DEPEND_0,omitempty"`
    
    FRAME                       string              `xml:"FRAME,omitempty"`
    TENSOR_ORDER                string              `xml:"TENSOR_ORDER,omitempty"`
    COORDINATE_SYSTEM           string              `xml:"COORDINATE_SYSTEM,omitempty"`
    FRAME_VELOCITY              string              `xml:"FRAME_VELOCITY,omitempty"`
    REPRESENTATION_1            string              `xml:"REPRESENTATION_1,omitempty"`
    LABEL_1                     string              `xml:"LABEL_1,omitempty"`

}   
                                         
type File struct {                         

    FILE_TYPE                   string              `xml:"FILE_TYPE"`
    METADATA_TYPE               string              `xml:"METADATA_TYPE"`
    METADATA_VERSION            string              `xml:"METADATA_VERSION"`
    LOGICAL_FILE_ID             string              `xml:"LOGICAL_FILE_ID"`
    VERSION_NUMBER              string              `xml:"VERSION_NUMBER"`
    DATASET_VERSION             string              `xml:"DATASET_VERSION"`
    FILE_CAVEATS                string              `xml:"FILE_CAVEATS"`
    FILE_TIME_SPAN              string              `xml:"FILE_TIME_SPAN"`
    GENERATION_DATE             string              `xml:"GENERATION_DATE"`
}


type Unexpected struct {                         

    ATTR                        []TypeKeyValue      `xml:"ATTR>A"`
    META                        []TypeKeyValue      `xml:"META>M"`
    VAR                         []TypeKeyValue      `xml:"VAR>V"`
    ERROR                       []TypeKeyValue      `xml:"ERROR>E"`
}


type TypeKeyValue struct {                         

    Key                         string              `xml:"key,attr"`
    Val                         string              `xml:"val,attr"`
}


//x     for ix, p := range m.DATASETS.PARAMETERS.PARAMETER {
//x         
//x         fmt.Println("")
//x         fmt.Println(ix, "PARAMETER_ID             ", p.PARAMETER_ID)
//x         fmt.Println(ix, "PARAMETER_TYPE           ", p.PARAMETER_TYPE)
//x         fmt.Println(ix, "CATDESC                  ", p.CATDESC)
//x         fmt.Println(ix, "UNITS                    ", p.UNITS)
//x         fmt.Println(ix, "SI_CONVERSION            ", p.SI_CONVERSION)
//x         fmt.Println(ix, "SIZES                    ", p.SIZES)
//x         fmt.Println(ix, "VALUE_TYPE               ", p.VALUE_TYPE)
//x         fmt.Println(ix, "SIGNIFICANT_DIGITS       ", p.SIGNIFICANT_DIGITS)
//x         fmt.Println(ix, "FILLVAL                  ", p.FILLVAL)
//x         fmt.Println(ix, "FIELDNAM                 ", p.FIELDNAM)
//x         fmt.Println(ix, "LABLAXIS                 ", p.LABLAXIS)
//x         fmt.Println(ix, "DELTA_PLUS               ", p.DELTA_PLUS)
//x         fmt.Println(ix, "DELTA_MINUS              ", p.DELTA_MINUS)
//x 
//x         fmt.Println(ix, "ENTITY                   ", p.ENTITY)
//x         fmt.Println(ix, "PROPERTY                 ", p.PROPERTY)
//x         fmt.Println(ix, "QUALITY                  ", p.QUALITY)
//x         fmt.Println(ix, "DEPEND_0                 ", p.DEPEND_0)
//x 
//x         fmt.Println(ix, "FRAME                    ", p.FRAME)
//x         fmt.Println(ix, "TENSOR_ORDER             ", p.TENSOR_ORDER)
//x         fmt.Println(ix, "COORDINATE_SYSTEM        ", p.COORDINATE_SYSTEM)
//x         fmt.Println(ix, "FRAME_VELOCITY           ", p.FRAME_VELOCITY)
//x         fmt.Println(ix, "REPRESENTATION_1         ", p.REPRESENTATION_1)
//x         fmt.Println(ix, "LABEL_1                  ", p.LABEL_1)
//x     }
//x         

func debug_show_results(output []byte) (err error){
    
    s1 := time.Now().Format(time.RFC3339)
    s1 = strings.Replace(s1, ":", "-", -1)
    
    l_filepath := "C:/Dump/" +  s1 + "-Test.xml"
    
    f, err := os.Create(l_filepath)
    if err != nil {
        panic("file create error\n")
    }
    defer f.Close()
    
    f.Write(output)
    err = exec.Command("open", l_filepath).Start()
    
    return
}

func (m *CAA_MetaData) dump()  error {

	output, err := xml.MarshalIndent(m, "", "  ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	os.Stdout.Write(output)
    
    debug_show_results(output)
    
    return err
} 
    
