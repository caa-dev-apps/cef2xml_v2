package main

import (
//x  	"fmt"
    "errors"
)

///////////////////////////////////////////////////////////////////////////////
//

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
}

var VarNameMap2 = map[string]eVar {}
    
func init() {

    for i, v := range VarName {
        VarNameMap2[v] = eVar(i)
    }
}
   
func getVar(s string) (m eVar, err error) {

    m, ok  := VarNameMap2[s]
    if(ok != true) {
        return m, errors.New("No Var/Parma Match: " + s)
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

