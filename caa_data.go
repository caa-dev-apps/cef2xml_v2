package main

import (
	"encoding/xml"
	"fmt"
 	"os"
)

///////////////////////////////////////////////////////////////////////////////
//

type CAA_MetaData struct {
    XMLName                     xml.Name            `xml:"CAA_METADATA"`
                                              
    MISSION                     string              `xml:"MISSION_METADATA>MISSION"`
    MISSION_TIME_SPAN           string              `xml:"MISSION_METADATA>MISSION_TIME_SPAN"`
    MISSION_AGENCY              string              `xml:"MISSION_METADATA>MISSION_AGENCY"`
    MISSION_DESCRIPTION         string              `xml:"MISSION_METADATA>MISSION_DESCRIPTION"`
    MISSION_KEY_PERSONNEL       string              `xml:"MISSION_METADATA>MISSION_KEY_PERSONNEL"`
    MISSION_REFERENCES          string              `xml:"MISSION_METADATA>MISSION_REFERENCES"`
    MISSION_REGION              []string            `xml:"MISSION_METADATA>MISSION_REGION"`
    MISSION_CAVEATS             string              `xml:"MISSION_METADATA>MISSION_CAVEATS"`
                                              
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
    EXPERIMENT_CAVEATS          []string            `xml:"EXPERIMENT_METADATA>EXPERIMENT_CAVEATS"`
                                              
    INSTRUMENTS                 Instruments         `xml:"EXPERIMENT_METADATA>INSTRUMENTS"`
}                                             
                                              
type Instruments struct {                     
                                              
    INSTRUMENT_NAME             []string            `xml:"INSTRUMENT_METADATA>INSTRUMENT_NAME"`
    INSTRUMENT_DESCRIPTION      []string            `xml:"INSTRUMENT_METADATA>INSTRUMENT_DESCRIPTION"`
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


func (m *CAA_MetaData) dump()  error {

    fmt.Println("+ = = = == ==  + ")

	output, err := xml.MarshalIndent(m, "", "  ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	os.Stdout.Write(output)
    
    return err
} 
    

///////////////////////////////////////////////////////////////////////////////
//
    
func (m *CAA_MetaData) start_meta(v *string)  error {
    err := error(nil)

    switch {
            
        default:
            fmt.Println("start_meta::", *v)

            m.DATASETS.UNEXPECTED.META                          = append(m.DATASETS.UNEXPECTED.META, TypeKeyValue{Key: "START_META", Val: *v })    
    }
    
    return err
} 

func (m *CAA_MetaData) start_var(v *string)  error {
    err := error(nil)

    switch {
            
        default:
            fmt.Println("start_var::", *v)

            m.DATASETS.UNEXPECTED.VAR                          = append(m.DATASETS.UNEXPECTED.VAR, TypeKeyValue{Key: "START_VAR", Val: *v })    
    }
    
    return err
}     
    

func (m *CAA_MetaData) kv_attr(k, v *string)  error {
    err := error(nil)

    switch {
        case "LOGICAL_FILE_ID" ==  *k :
            m.DATASETS.FILE.LOGICAL_FILE_ID = *v                                // just for testing!!!!
            
        case "VERSION_NUMBER" ==  *k :
            m.DATASETS.FILE.VERSION_NUMBER = *v
            
        default:
            fmt.Println("kv_attr::", *k, *v)

            m.DATASETS.UNEXPECTED.ATTR                          = append(m.DATASETS.UNEXPECTED.ATTR, TypeKeyValue{Key: *k, Val: *v })    
            
    }
    
    return err
} 


