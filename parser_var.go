package main

import (
//x  	"strings"
//x 	"fmt"
//x     "errors"
)

///////////////////////////////////////////////////////////////////////////////
//
// 
//  type CefHeaderData struct {
//      m_state State
//      m_name string 
//      m_data CAA_MetaData
//      m_meta eMeta
//      m_var Parameter
//  }
 
 
 
//  todo    PARAMETER_ID = iota
//  todo    PARAMETER_TYPE         
//  todo    CATDESC                
//  todo    UNITS                  
//  todo    SI_CONVERSION          
//  todo    SIZES                  
//  todo    VALUE_TYPE             
//  todo    SIGNIFICANT_DIGITS     
//  todo    FILLVAL                
//  todo    FIELDNAM               
//  todo    LABLAXIS               
//  todo    DELTA_PLUS             
//  todo    DELTA_MINUS            
//  todo    ENTITY                 
//  todo    PROPERTY               
//  todo    QUALITY                
//  todo    DEPEND_0               
 
 
func (hds *CefHeaderData) kv_var(k, v *string)  (err error) {
//x     err := hds.m_data.kv_var(k, v)

//  todo    switch hds.m_meta {
//  todo        
//  todo    }
//  todo

    hds.m_state = VAR
    return 
} 

func (hds *CefHeaderData) new_var()  (err error) {
    
    hds.m_var = Parameter{}
    hds.m_data.DATASETS.PARAMETERS.PARAMETER = append(hds.m_data.DATASETS.PARAMETERS.PARAMETER, hds.m_var)
    
    return 
} 

