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
 
func (hds *CefHeaderData) kv_var(k, v *string)  (err error) {

    println("x x x x x x x x x x  x x x x x x x x x x x ", *k)

    evar, err := getVar(k) 
    if(err != nil) {
//x         return err
    }

    switch evar {
//x     case PARAMETER_ID:        see below
        case PARAMETER_TYPE:      hds.m_var.PARAMETER_TYPE     = *v    
        case CATDESC:             hds.m_var.CATDESC            = *v    
        case UNITS:               hds.m_var.UNITS              = *v    
        case SI_CONVERSION:       hds.m_var.SI_CONVERSION      = *v    
        case SIZES:               hds.m_var.SIZES              = *v    
        case VALUE_TYPE:          hds.m_var.VALUE_TYPE         = *v    
        case SIGNIFICANT_DIGITS:  hds.m_var.SIGNIFICANT_DIGITS = *v    
        case FILLVAL:             hds.m_var.FILLVAL            = *v    
        case FIELDNAM:            hds.m_var.FIELDNAM           = *v    
        case LABLAXIS:            hds.m_var.LABLAXIS           = *v    
        case DELTA_PLUS:          hds.m_var.DELTA_PLUS         = *v    
        case DELTA_MINUS:         hds.m_var.DELTA_MINUS        = *v    
        case ENTITY:              hds.m_var.ENTITY             = *v    
        case PROPERTY:            hds.m_var.PROPERTY           = *v    
        case QUALITY:             hds.m_var.QUALITY            = *v    
        case DEPEND_0:            hds.m_var.DEPEND_0           = *v    
                
        case FRAME:               hds.m_var.FRAME              =  *v   
        case TENSOR_ORDER:        hds.m_var.TENSOR_ORDER       =  *v   
        case COORDINATE_SYSTEM:   hds.m_var.COORDINATE_SYSTEM  =  *v   
        case FRAME_VELOCITY:      hds.m_var.FRAME_VELOCITY     =  *v   
        case REPRESENTATION_1:    hds.m_var.REPRESENTATION_1   =  *v   
        case LABEL_1:             hds.m_var.LABEL_1            =  *v   
                                        
        default:
            // nothing will get here since getVar filters 
            println("====================================================================== TODO", *k)
    }
    
    return 
} 

func (hds *CefHeaderData) stx_var(name *string)  (err error) {
    
    hds.m_var = Parameter{}
    hds.m_var.PARAMETER_ID = *name
    
    return 
} 

func (hds *CefHeaderData) etx_var()  (err error) {
    
    hds.m_data.DATASETS.PARAMETERS.PARAMETER = append(hds.m_data.DATASETS.PARAMETERS.PARAMETER, hds.m_var)
    
    return 
} 

