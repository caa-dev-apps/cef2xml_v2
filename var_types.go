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

