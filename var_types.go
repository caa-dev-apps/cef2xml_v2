package main

import (
    "errors"
)

type eVar int

const (
    PARAMETER_ID = iota
    PARAMETER_TYPE         
    CATDESC                
    UNITS                  
    SI_CONVERSION          
    SIZES                  
    VALUE_TYPE             
    SIGNIFICANT_DIGITS     
    FILLVAL                
    FIELDNAM               
    LABLAXIS               
    DELTA_PLUS             
    DELTA_MINUS            
    ENTITY                 
    PROPERTY               
    QUALITY                
    DEPEND_0     
    DEPEND_1     
    DEPEND_2     
    DEPEND_3     
    
    FRAME                       
    TENSOR_ORDER                    
    COORDINATE_SYSTEM               
    FRAME_VELOCITY                  
    REPRESENTATION_0                
    REPRESENTATION_1                
    REPRESENTATION_2                
    REPRESENTATION_3                
    LABEL_0                         
    LABEL_1                         
    LABEL_2                         
    LABEL_3    

    SCALEMIN                    
    SCALEMAX                    
    SCALETYP                    
    DISPLAYTYPE                 
    DATA                        
)

//x const (
//x     ACKNOWLEDGEMENT = iota
//x     CATDESC
//x     COMPOUND
//x     COMPOUND_DEF
//x     CONTACT_COORDINATES
//x     COORDINATE_SYSTEM
//x     DATA
//x     DATA_TYPE
//x     DATASET_CAVEATS
//x     DATASET_DESCRIPTION
//x     DATASET_ID
//x     DATASET_TITLE
//x     DELTA_PLUS 
//x     DELTA_MINUS
//x     DEPEND_0
//x     DEPEND_1
//x     DEPEND_2
//x     DEPEND_3
//x     DISPLAYTYPE
//x     ENTITY
//x     ERROR_PLUS and ERROR_MINUS
//x     EXPERIMENT
//x     EXPERIMENT_CAVEATS
//x     EXPERIMENT_DESCRIPTION
//x     EXPERIMENT_KEY_PERSONNEL
//x     EXPERIMENT_REFERENCES
//x     FIELDNAM
//x     FILLVAL
//x     FLUCTUATIONS
//x     FRAME
//x     FRAME_VELOCITY
//x     FRAME_ORIGIN
//x     INSTRUMENT_CAVEATS
//x     INSTRUMENT_DESCRIPTION
//x     INSTRUMENT_NAME
//x     INSTRUMENT_TYPE
//x     INVESTIGATOR_COORDINATES
//x     LABEL_0
//x     LABEL_1
//x     LABEL_2 
//x     LABEL_3 
//x     LABLAXIS
//x     MAX_TIME_RESOLUTION
//x     MEASUREMENT_TYPE
//x     MIN_TIME_RESOLUTION
//x     MISSION
//x     MISSION_AGENCY
//x     MISSION_CAVEATS
//x     MISSION_DESCRIPTION
//x     MISSION_KEY_PERSONNEL
//x     MISSION_REFERENCES
//x     MISSION_REGION
//x     MISSION_TIME_SPAN
//x     OBSERVATORY
//x     OBSERVATORY_CAVEATS
//x     OBSERVATORY_DESCRIPTION
//x     OBSERVATORY_REGION
//x     OBSERVATORY_TIME_SPAN
//x     PARAMETER_CAVEATS
//x     PARAMETER_ID
//x     PARAMETER_TYPE
//x     PARENT_DATASET_ID
//x     PARENT_EXPERIMENT
//x     PARENT_INSTRUMENT
//x     PARENT_MISSION
//x     PARENT_OBSERVATORY
//x     PROCESSING_LEVEL
//x     PROPERTY
//x     QUALITY
//x     REPRESENTATION
//x     REPRESENTATION_1
//x     SCALEMAX
//x     SCALEMIN
//x     SCALETYP
//x     SI_CONVERSION
//x     SIGNIFICANT_DIGITS
//x     SIZES
//x     TARGET_SYSTEM
//x     TENSOR_FRAME
//x     TENSOR_ORDER
//x     TIME_RESOLUTION
//x     UNITS
//x     VALUE_TYPE
//x )

var VarName = []string{
    PARAMETER_ID:              "PARAMETER_ID", 
    PARAMETER_TYPE:            "PARAMETER_TYPE", 
    CATDESC:                   "CATDESC",
    UNITS:                     "UNITS",
    SI_CONVERSION:             "SI_CONVERSION", 
    SIZES:                     "SIZES",
    VALUE_TYPE:                "VALUE_TYPE",
    SIGNIFICANT_DIGITS:        "SIGNIFICANT_DIGITS",     
    FILLVAL:                   "FILLVAL",
    FIELDNAM:                  "FIELDNAM",
    LABLAXIS:                  "LABLAXIS",
    DELTA_PLUS:                "DELTA_PLUS",
    DELTA_MINUS:               "DELTA_MINUS", 
    ENTITY:                    "ENTITY",
    PROPERTY:                  "PROPERTY",
    QUALITY:                   "QUALITY",
    DEPEND_0:                  "DEPEND_0",
    DEPEND_1:                  "DEPEND_1",
    DEPEND_2:                  "DEPEND_2",
    DEPEND_3:                  "DEPEND_3",
    FRAME:                     "FRAME", 
    TENSOR_ORDER:              "TENSOR_ORDER",     
    COORDINATE_SYSTEM:         "COORDINATE_SYSTEM",     
    FRAME_VELOCITY:            "FRAME_VELOCITY",     
    REPRESENTATION_0:          "REPRESENTATION_0",     
    REPRESENTATION_1:          "REPRESENTATION_1",     
    REPRESENTATION_2:          "REPRESENTATION_2",     
    REPRESENTATION_3:          "REPRESENTATION_3",     
    LABEL_0:                   "LABEL_0",     
    LABEL_1:                   "LABEL_1",     
    LABEL_2:                   "LABEL_2",     
    LABEL_3:                   "LABEL_3",    

    SCALEMIN:                  "SCALEMIN",                
    SCALEMAX:                  "SCALEMAX",                    
    SCALETYP:                  "SCALETYP",                    
    DISPLAYTYPE:               "DISPLAYTYPE",                 
    DATA:                      "DATA",                        
}

//x var VarName = []string{
//x     ACKNOWLEDGEMENT:               "ACKNOWLEDGEMENT",
//x     CATDESC:                       "CATDESC",
//x     COMPOUND:                      "COMPOUND",
//x     COMPOUND_DEF:                  "COMPOUND_DEF",
//x     CONTACT_COORDINATES:           "CONTACT_COORDINATES",
//x     COORDINATE_SYSTEM:             "COORDINATE_SYSTEM",
//x     DATA:                          "DATA",
//x     DATA_TYPE:                     "DATA_TYPE",
//x     DATASET_CAVEATS:               "DATASET_CAVEATS",
//x     DATASET_DESCRIPTION:           "DATASET_DESCRIPTION",
//x     DATASET_ID:                    "DATASET_ID",
//x     DATASET_TITLE:                 "DATASET_TITLE",
//x     DELTA_PLUS :                   "DELTA_PLUS",
//x     DELTA_MINUS:                   "DELTA_MINUS",
//x     DEPEND_0:                      "DEPEND_0",
//x     DEPEND_1:                      "DEPEND_1",
//x     DEPEND_2:                      "DEPEND_2",
//x     DEPEND_3:                      "DEPEND_3",
//x     DISPLAYTYPE:                   "DISPLAYTYPE",
//x     ENTITY:                        "ENTITY",
//x     ERROR_PLUS and ERROR_MINUS:    "ERROR_PLUS",
//x     EXPERIMENT:                    "EXPERIMENT",
//x     EXPERIMENT_CAVEATS:            "EXPERIMENT_CAVEATS",
//x     EXPERIMENT_DESCRIPTION:        "EXPERIMENT_DESCRIPTION",
//x     EXPERIMENT_KEY_PERSONNEL:      "EXPERIMENT_KEY_PERSONNEL",
//x     EXPERIMENT_REFERENCES:         "EXPERIMENT_REFERENCES",
//x     FIELDNAM:                      "FIELDNAM",
//x     FILLVAL:                       "FILLVAL",
//x     FLUCTUATIONS:                  "FLUCTUATIONS",
//x     FRAME:                         "FRAME",
//x     FRAME_VELOCITY:                "FRAME_VELOCITY",
//x     FRAME_ORIGIN:                  "FRAME_ORIGIN",
//x     INSTRUMENT_CAVEATS:            "INSTRUMENT_CAVEATS",
//x     INSTRUMENT_DESCRIPTION:        "INSTRUMENT_DESCRIPTION",
//x     INSTRUMENT_NAME:               "INSTRUMENT_NAME",
//x     INSTRUMENT_TYPE:               "INSTRUMENT_TYPE",
//x     INVESTIGATOR_COORDINATES:      "INVESTIGATOR_COORDINATES",
//x     LABEL_0:                       "LABEL_0",
//x     LABEL_1:                       "LABEL_1",
//x     LABEL_2:                       "LABEL_2",
//x     LABEL_3:                       "LABEL_3",
//x     LABLAXIS:                      "LABLAXIS",
//x     MAX_TIME_RESOLUTION:           "MAX_TIME_RESOLUTION",
//x     MEASUREMENT_TYPE:              "MEASUREMENT_TYPE",
//x     MIN_TIME_RESOLUTION:           "MIN_TIME_RESOLUTION",
//x     MISSION:                       "MISSION",
//x     MISSION_AGENCY:                "MISSION_AGENCY",
//x     MISSION_CAVEATS:               "MISSION_CAVEATS",
//x     MISSION_DESCRIPTION:           "MISSION_DESCRIPTION",
//x     MISSION_KEY_PERSONNEL:         "MISSION_KEY_PERSONNEL",
//x     MISSION_REFERENCES:            "MISSION_REFERENCES",
//x     MISSION_REGION:                "MISSION_REGION",
//x     MISSION_TIME_SPAN:             "MISSION_TIME_SPAN",
//x     OBSERVATORY:                   "OBSERVATORY",
//x     OBSERVATORY_CAVEATS:           "OBSERVATORY_CAVEATS",
//x     OBSERVATORY_DESCRIPTION:       "OBSERVATORY_DESCRIPTION",
//x     OBSERVATORY_REGION:            "OBSERVATORY_REGION",
//x     OBSERVATORY_TIME_SPAN:         "OBSERVATORY_TIME_SPAN",
//x     PARAMETER_CAVEATS:             "PARAMETER_CAVEATS",
//x     PARAMETER_ID:                  "PARAMETER_ID",
//x     PARAMETER_TYPE:                "PARAMETER_TYPE",
//x     PARENT_DATASET_ID:             "PARENT_DATASET_ID",
//x     PARENT_EXPERIMENT:             "PARENT_EXPERIMENT",
//x     PARENT_INSTRUMENT:             "PARENT_INSTRUMENT",
//x     PARENT_MISSION:                "PARENT_MISSION",
//x     PARENT_OBSERVATORY:            "PARENT_OBSERVATORY",
//x     PROCESSING_LEVEL:              "PROCESSING_LEVEL",
//x     PROPERTY:                      "PROPERTY",
//x     QUALITY:                       "QUALITY",
//x     REPRESENTATION:                "REPRESENTATION",
//x     REPRESENTATION_1:              "REPRESENTATION_1",
//x     SCALEMAX:                      "SCALEMAX",
//x     SCALEMIN:                      "SCALEMIN",
//x     SCALETYP:                      "SCALETYP",
//x     SI_CONVERSION:                 "SI_CONVERSION",
//x     SIGNIFICANT_DIGITS:            "SIGNIFICANT_DIGITS",
//x     SIZES:                         "SIZES",
//x     TARGET_SYSTEM:                 "TARGET_SYSTEM",
//x     TENSOR_FRAME:                  "TENSOR_FRAME",
//x     TENSOR_ORDER:                  "TENSOR_ORDER",
//x     TIME_RESOLUTION:               "TIME_RESOLUTION",
//x     UNITS:                         "UNITS",
//x     VALUE_TYPE:                    "VALUE_TYPE"
//x }


var VarNameMap2 = map[string]eVar {}
    
func init() {

    for i, v := range VarName {
        VarNameMap2[v] = eVar(i)
    }
}
   
func getVar(s *string) (m eVar, err error) {

    m, ok  := VarNameMap2[*s]
    if(ok != true) {
        return m, errors.New("No Var/Parma Match: " + *s)
    }

    return 
}

func getVarString(m eVar) (s string, err error) {

    l := len(VarName)

    if m < 0 || m >= eVar(l)  {
        return s, errors.New("eVar Index out of range: " + string(m))
    }
    
    s = VarName[m]
    
    return 
}

