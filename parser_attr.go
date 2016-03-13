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
 
//old   func (hds *CefHeaderData) kv_attr(k, v *string)  (err error) {
//old   
//old       switch {
//old           case "LOGICAL_FILE_ID" ==  *k :
//old               hds.m_data.DATASETS.FILE.LOGICAL_FILE_ID = *v                                // just for testing!!!!
//old               
//old           case "VERSION_NUMBER" ==  *k :
//old               hds.m_data.DATASETS.FILE.VERSION_NUMBER = *v
//old               
//old           default:
//old               mooi_log("kv_attr::", *k, *v)
//old   
//old               hds.m_data.DATASETS.UNEXPECTED.ATTR                          = append(hds.m_data.DATASETS.UNEXPECTED.ATTR, TypeKeyValue{Key: *k, Val: *v })    
//old       }
//old       
//old       return 
//old   } 

func (hds *CefHeaderData) kv_attr(kv *KeyVal)  (err error) {

    switch {
        case "LOGICAL_FILE_ID" ==  kv.key :
            hds.m_data.DATASETS.FILE.LOGICAL_FILE_ID = kv.val[0]                                // just for testing!!!!
            
        case "VERSION_NUMBER" ==  kv.key :
            hds.m_data.DATASETS.FILE.VERSION_NUMBER = kv.val[0]
            
        default:
            mooi_log("kv_attr::", *kv)

//x             hds.m_data.DATASETS.UNEXPECTED.ATTR                          = append(hds.m_data.DATASETS.UNEXPECTED.ATTR, TypeKeyValue{Key: *k, Val: *v })    
            hds.m_data.DATASETS.UNEXPECTED.ATTR                          = append(hds.m_data.DATASETS.UNEXPECTED.ATTR, TypeKeyValue{Key: kv.key, Val: kv.val[0] })    
    }
    
    return 
} 
