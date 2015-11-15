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
 
///////////////////////////////////////////////////////////////////////////////
//

//x         
//x             hds.m_data.DATASETS.FILE.VERSION_NUMBER = *v
//x             
//x         default:
//x             fmt.Println("kv_attr::", *k, *v)
//x 
//x             hds.m_data.DATASETS.UNEXPECTED.ATTR                          = append(hds.m_data.DATASETS.UNEXPECTED.ATTR, TypeKeyValue{Key: *k, Val: *v })    
//x         
        



func (hds *CefHeaderData) kv_meta(k, v *string)  (err error){

    //x fmt.Println("##  #  #  #  #  #  #  #  #  #  #   # # # ##", hds.m_meta, *k, *v)

    switch hds.m_meta {
        case LOGICAL_FILE_ID :
        case VERSION_NUMBER:
        case FILE_TIME_SPAN:
        case GENERATION_DATE:
        case FILE_CAVEATS:                      
        case MISSION:                           hds.m_data.MISSION = *v
        case MISSION_TIME_SPAN:                 hds.m_data.MISSION_TIME_SPAN = *v
        case MISSION_AGENCY:                    hds.m_data.MISSION_AGENCY = *v
        case MISSION_DESCRIPTION:               hds.m_data.MISSION_DESCRIPTION += (` ` + *v)
        case MISSION_KEY_PERSONNEL:             hds.m_data.MISSION_KEY_PERSONNEL = append(hds.m_data.MISSION_KEY_PERSONNEL, *v)    
        case MISSION_REFERENCES:                hds.m_data.MISSION_REFERENCES = *v
        case MISSION_REGION:                    hds.m_data.MISSION_REGION = append(hds.m_data.MISSION_REGION, *v)    
        case MISSION_CAVEATS:                   hds.m_data.MISSION_CAVEATS = append(hds.m_data.MISSION_CAVEATS, *v)    
        case EXPERIMENT:
        case EXPERIMENT_DESCRIPTION:
        case INVESTIGATOR_COORDINATES:
        case EXPERIMENT_REFERENCES:
        case EXPERIMENT_KEY_PERSONNEL:
        case EXPERIMENT_CAVEATS:
        case OBSERVATORY:
        case OBSERVATORY_CAVEATS:
        case OBSERVATORY_DESCRIPTION:
        case OBSERVATORY_TIME_SPAN:
        case OBSERVATORY_REGION:
        case INSTRUMENT_NAME:
        case INSTRUMENT_DESCRIPTION:
        case INSTRUMENT_TYPE:
        case MEASUREMENT_TYPE:
        case INSTRUMENT_CAVEATS:
        case DATASET_ID:
        case DATA_TYPE:
        case DATASET_TITLE:
        case DATASET_DESCRIPTION:
        case CONTACT_COORDINATES:
        case TIME_RESOLUTION:
        case MIN_TIME_RESOLUTION:
        case MAX_TIME_RESOLUTION:
        case PROCESSING_LEVEL:
        case ACKNOWLEDGEMENT:
        case DATASET_CAVEATS:
        case DATASET_VERSION:
        case FILE_TYPE:
        case METADATA_TYPE:
        case METADATA_VERSION:
        default:
    }
    
    return 
} 

