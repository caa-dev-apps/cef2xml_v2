package main

import (
	"encoding/xml"
	"fmt"
	"os"
    "testing"

)

///////////////////////////////////////////////////////////////////////////////
//

func Test_write_xml_04(t *testing.T) {
    
    m := &CAA_MetaData {}
    
    m.MISSION                                           = "-"
    m.MISSION_TIME_SPAN                                 = "-"
    m.MISSION_AGENCY                                    = "-"
    m.MISSION_DESCRIPTION                               = "-"
    m.MISSION_KEY_PERSONNEL                             = "-"
    m.MISSION_REFERENCES                                = "-"
    m.MISSION_REGION                                    = append(m.MISSION_REGION, "[]")    
    m.MISSION_CAVEATS                                   = "-"
                                            
    m.OBSERVATORIES.OBSERVATORY                         = append(m.OBSERVATORIES.OBSERVATORY, "[]")    
    m.OBSERVATORIES.OBSERVATORY_CAVEATS                 = append(m.OBSERVATORIES.OBSERVATORY_CAVEATS, "[]")    
    m.OBSERVATORIES.OBSERVATORY_DESCRIPTION             = "-"
    m.OBSERVATORIES.OBSERVATORY_TIME_SPAN               = "-"
    m.OBSERVATORIES.OBSERVATORY_REGION                  = append(m.OBSERVATORIES.OBSERVATORY_REGION, "[]")    
                                            
    m.EXPERIMENTS.EXPERIMENT                            = "-"
    m.EXPERIMENTS.EXPERIMENT_DESCRIPTION                = "-"
    m.EXPERIMENTS.INVESTIGATOR_COORDINATES              = append(m.EXPERIMENTS.INVESTIGATOR_COORDINATES, "[]")    
    m.EXPERIMENTS.EXPERIMENT_REFERENCES                 = append(m.EXPERIMENTS.EXPERIMENT_REFERENCES, "[]")    
    m.EXPERIMENTS.EXPERIMENT_KEY_PERSONNEL              = append(m.EXPERIMENTS.EXPERIMENT_KEY_PERSONNEL, "[]")    
    m.EXPERIMENTS.EXPERIMENT_CAVEATS                    = append(m.EXPERIMENTS.EXPERIMENT_CAVEATS, "[]")    
    
    m.EXPERIMENTS.INSTRUMENTS.INSTRUMENT_NAME           = append(m.EXPERIMENTS.INSTRUMENTS.INSTRUMENT_NAME, "[]")    
    m.EXPERIMENTS.INSTRUMENTS.INSTRUMENT_DESCRIPTION    = append(m.EXPERIMENTS.INSTRUMENTS.INSTRUMENT_DESCRIPTION, "[]")    
    m.EXPERIMENTS.INSTRUMENTS.INSTRUMENT_TYPE           = append(m.EXPERIMENTS.INSTRUMENTS.INSTRUMENT_TYPE, "[]")    
    m.EXPERIMENTS.INSTRUMENTS.MEASUREMENT_TYPE          = append(m.EXPERIMENTS.INSTRUMENTS.MEASUREMENT_TYPE, "[]")    
    m.EXPERIMENTS.INSTRUMENTS.INSTRUMENT_CAVEATS        = append(m.EXPERIMENTS.INSTRUMENTS.INSTRUMENT_CAVEATS, "[]")    

    
    m.DATASETS.DATASET_ID                               = "-"
    m.DATASETS.DATA_TYPE                                = "-"
    m.DATASETS.DATASET_TITLE                            = "-"
    m.DATASETS.DATASET_DESCRIPTION                      = "-"
                                           
    m.DATASETS.CONTACT_COORDINATES                      = append(m.DATASETS.CONTACT_COORDINATES, "[]")    
                                           
    m.DATASETS.PROCESSING_LEVEL                         = "-"
    m.DATASETS.TIME_RESOLUTION                          = "-"
    m.DATASETS.MIN_TIME_RESOLUTION                      = "-"
    m.DATASETS.MAX_TIME_RESOLUTION                      = "-"
    m.DATASETS.DATASET_CAVEATS                          = "-"
    m.DATASETS.ACKNOWLEDGEMENT                          = "-"
    
    p1 := &Parameter { PARAMETER_ID: "-1",  PARAMETER_TYPE : "-", CATDESC : "-",   UNITS : "-" }
    p2 := &Parameter { PARAMETER_ID: "-2",  PARAMETER_TYPE : "-", CATDESC : "-",   UNITS : "-" }
    
    m.DATASETS.PARAMETERS.PARAMETER                     = append(m.DATASETS.PARAMETERS.PARAMETER, *p1)    
    m.DATASETS.PARAMETERS.PARAMETER                     = append(m.DATASETS.PARAMETERS.PARAMETER, *p2)    
    
    m.DATASETS.FILE.FILE_TYPE                           = "-"
    m.DATASETS.FILE.METADATA_TYPE                       = "-"
    m.DATASETS.FILE.METADATA_VERSION                    = "-"
    m.DATASETS.FILE.LOGICAL_FILE_ID                     = "-"
    m.DATASETS.FILE.VERSION_NUMBER                      = "-"
    m.DATASETS.FILE.DATASET_VERSION                     = "-"
    m.DATASETS.FILE.FILE_CAVEATS                        = "-"
    m.DATASETS.FILE.FILE_TIME_SPAN                      = "-"
    m.DATASETS.FILE.GENERATION_DATE                     = "-"

    
	output, err := xml.MarshalIndent(m, "", "  ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
        t.Errorf("error: %v\n", err)
	}

	os.Stdout.Write(output)
} 
