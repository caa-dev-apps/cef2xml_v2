package main

import (
)


func (hds *CefHeaderData) kv_attr(kv *KeyVal)  (err error) {

    switch {
        case "FILE_NAME" ==  kv.key :
            hds.m_data.DATASETS.FILE.FILE_NAME = kv.val[0]
        case "FILE_FORMAT_VERSION" ==  kv.key :
            hds.m_data.DATASETS.FILE.FILE_FORMAT_VERSION = kv.val[0]
        case "END_OF_RECORD_MARKER" ==  kv.key :
            hds.m_data.DATASETS.FILE.END_OF_RECORD_MARKER = kv.val[0]
        case "DATA_UNTIL" ==  kv.key :
            hds.m_data.DATASETS.FILE.DATA_UNTIL = kv.val[0]
        default:
            mooi_log("kv_attr::", *kv)
            hds.m_data.DATASETS.UNEXPECTED.ATTR = append(hds.m_data.DATASETS.UNEXPECTED.ATTR, TypeKeyValue{Key: kv.key, Val: kv.val[0] })    
    }
    
    return 
} 
