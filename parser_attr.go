package main

import (
)

func (hds *CefHeaderData) kv_attr(kv *KeyVal)  (err error) {

    switch {
        case "LOGICAL_FILE_ID" ==  kv.key :
            hds.m_data.DATASETS.FILE.LOGICAL_FILE_ID = kv.val[0]                                // just for testing!!!!
            
        case "VERSION_NUMBER" ==  kv.key :
            hds.m_data.DATASETS.FILE.VERSION_NUMBER = kv.val[0]

        case "FILE_NAME" ==  kv.key :
            hds.m_data.DATASETS.FILE.FILE_NAME = kv.val[0]
            
        default:
            mooi_log("kv_attr::", *kv)

            hds.m_data.DATASETS.UNEXPECTED.ATTR = append(hds.m_data.DATASETS.UNEXPECTED.ATTR, TypeKeyValue{Key: kv.key, Val: kv.val[0] })    
    }
    
    return 
} 
