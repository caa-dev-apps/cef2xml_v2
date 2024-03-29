package main

import (
    "errors"
)

///////////////////////////////////////////////////////////////////////////////
//

type eMeta int

const (
    LOGICAL_FILE_ID eMeta = iota
    VERSION_NUMBER
    FILE_TIME_SPAN
    GENERATION_DATE
    FILE_NAME
    FILE_CAVEATS
    MISSION
    MISSION_TIME_SPAN
    MISSION_AGENCY
    MISSION_DESCRIPTION
    MISSION_KEY_PERSONNEL
    MISSION_REFERENCES
    MISSION_REGION
    MISSION_CAVEATS
    EXPERIMENT
    EXPERIMENT_DESCRIPTION
    INVESTIGATOR_COORDINATES
    EXPERIMENT_REFERENCES
    EXPERIMENT_KEY_PERSONNEL
    EXPERIMENT_CAVEATS
    OBSERVATORY
    OBSERVATORY_CAVEATS
    OBSERVATORY_DESCRIPTION
    OBSERVATORY_TIME_SPAN
    OBSERVATORY_REGION
    INSTRUMENT_NAME
    INSTRUMENT_DESCRIPTION
    INSTRUMENT_TYPE
    MEASUREMENT_TYPE
    INSTRUMENT_CAVEATS
    DATASET_ID
    DATA_TYPE
    DATASET_TITLE
    DATASET_TYPE
    DATASET_TIME_SPAN
    DATASET_DESCRIPTION
    CONTACT_COORDINATES
    TIME_RESOLUTION
    MIN_TIME_RESOLUTION
    MAX_TIME_RESOLUTION
    PROCESSING_LEVEL
    ACKNOWLEDGEMENT
    DATASET_CAVEATS
    DATASET_VERSION
    FILE_TYPE
    METADATA_TYPE
    METADATA_VERSION
)

var MetaName = []string{
    LOGICAL_FILE_ID:            "LOGICAL_FILE_ID",
    VERSION_NUMBER:             "VERSION_NUMBER",
    FILE_TIME_SPAN:             "FILE_TIME_SPAN",
    GENERATION_DATE:            "GENERATION_DATE",
    FILE_NAME:                  "FILE_NAME",
    FILE_CAVEATS:               "FILE_CAVEATS",
    MISSION:                    "MISSION",
    MISSION_TIME_SPAN:          "MISSION_TIME_SPAN",
    MISSION_AGENCY:             "MISSION_AGENCY",
    MISSION_DESCRIPTION:        "MISSION_DESCRIPTION",
    MISSION_KEY_PERSONNEL:      "MISSION_KEY_PERSONNEL",
    MISSION_REFERENCES:         "MISSION_REFERENCES",
    MISSION_REGION:             "MISSION_REGION",
    MISSION_CAVEATS:            "MISSION_CAVEATS",
    EXPERIMENT:                 "EXPERIMENT",
    EXPERIMENT_DESCRIPTION:     "EXPERIMENT_DESCRIPTION",
    INVESTIGATOR_COORDINATES:   "INVESTIGATOR_COORDINATES",
    EXPERIMENT_REFERENCES:      "EXPERIMENT_REFERENCES",
    EXPERIMENT_KEY_PERSONNEL:   "EXPERIMENT_KEY_PERSONNEL",
    EXPERIMENT_CAVEATS:         "EXPERIMENT_CAVEATS",
    OBSERVATORY:                "OBSERVATORY",
    OBSERVATORY_CAVEATS:        "OBSERVATORY_CAVEATS",
    OBSERVATORY_DESCRIPTION:    "OBSERVATORY_DESCRIPTION",
    OBSERVATORY_TIME_SPAN:      "OBSERVATORY_TIME_SPAN",
    OBSERVATORY_REGION:         "OBSERVATORY_REGION",
    INSTRUMENT_NAME:            "INSTRUMENT_NAME",
    INSTRUMENT_DESCRIPTION:     "INSTRUMENT_DESCRIPTION",
    INSTRUMENT_TYPE:            "INSTRUMENT_TYPE",
    MEASUREMENT_TYPE:           "MEASUREMENT_TYPE",
    INSTRUMENT_CAVEATS:         "INSTRUMENT_CAVEATS",
    DATASET_ID:                 "DATASET_ID",
    DATA_TYPE:                  "DATA_TYPE",
    DATASET_TITLE:              "DATASET_TITLE",
    DATASET_TYPE:               "DATASET_TYPE",
    DATASET_TIME_SPAN:           "DATASET_TIME_SPAN",
    DATASET_DESCRIPTION:        "DATASET_DESCRIPTION",
    CONTACT_COORDINATES:        "CONTACT_COORDINATES",
    TIME_RESOLUTION:            "TIME_RESOLUTION",
    MIN_TIME_RESOLUTION:        "MIN_TIME_RESOLUTION",
    MAX_TIME_RESOLUTION:        "MAX_TIME_RESOLUTION",
    PROCESSING_LEVEL:           "PROCESSING_LEVEL",
    ACKNOWLEDGEMENT:            "ACKNOWLEDGEMENT",
    DATASET_CAVEATS:            "DATASET_CAVEATS",
    DATASET_VERSION:            "DATASET_VERSION",
    FILE_TYPE:                  "FILE_TYPE",
    METADATA_TYPE:              "METADATA_TYPE",
    METADATA_VERSION:           "METADATA_VERSION",
}
     
var MetaNameMap2 = map[string]eMeta {}
    
func init() {

    for i, v := range MetaName {
        MetaNameMap2[v] = eMeta(i)
    }
}
   
func getMeta(s string) (m eMeta, err error) {

    m, ok  := MetaNameMap2[s]
    if(ok != true) {
        return m, errors.New("No Meta Match: " + s)
    }

    return 
}

func getMetaString(m eMeta) (s string, err error) {

    l := len(MetaName)

    if m < 0 || m >= eMeta(l)  {
        return s, errors.New("eMeta Index out of range: " + string(m))
    }
    
    s = MetaName[m]
    
    return 
}


