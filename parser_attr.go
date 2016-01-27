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
 
func (hds *CefHeaderData) kv_attr(k, v *string)  (err error) {

    switch {
        case "LOGICAL_FILE_ID" ==  *k :
            hds.m_data.DATASETS.FILE.LOGICAL_FILE_ID = *v                                // just for testing!!!!
            
        case "VERSION_NUMBER" ==  *k :
            hds.m_data.DATASETS.FILE.VERSION_NUMBER = *v
            
        default:
            mooi_log("kv_attr::", *k, *v)

            hds.m_data.DATASETS.UNEXPECTED.ATTR                          = append(hds.m_data.DATASETS.UNEXPECTED.ATTR, TypeKeyValue{Key: *k, Val: *v })    
    }
    
    return 
} 
