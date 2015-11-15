package main

import (
//x  	"strings"
//x 	"fmt"
//x     "errors"
)

///////////////////////////////////////////////////////////////////////////////
//
// 
// type CefHeaderData struct {
//     m_state State
//     m_name string 
//     m_data CAA_MetaData
//     m_meta eMeta
// }
 
func (hds *CefHeaderData) kv_var(k, v *string)  (err error) {
//x     err := hds.m_data.kv_var(k, v)

    hds.m_state = VAR
    return 
} 

