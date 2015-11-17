package main

import (
 	"strings"
//x 	"fmt"
    "errors"
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
        



//x func (hds *CefHeaderData) kv_meta_entry(v *string)  (err error){
//x 
//x     //x fmt.Println("##  #  #  #  #  #  #  #  #  #  #   # # # ##", hds.m_meta, *k, *v)
//x 
//x     switch hds.m_meta {
//x         case LOGICAL_FILE_ID :
//x         case VERSION_NUMBER:
//x         case FILE_TIME_SPAN:
//x         case GENERATION_DATE:
//x         case FILE_CAVEATS:                      
//x         case MISSION:                           hds.m_data.MISSION = *v
//x         case MISSION_TIME_SPAN:                 hds.m_data.MISSION_TIME_SPAN = *v
//x         case MISSION_AGENCY:                    hds.m_data.MISSION_AGENCY = *v
//x         case MISSION_DESCRIPTION:               hds.m_data.MISSION_DESCRIPTION += (` ` + *v)
//x         case MISSION_KEY_PERSONNEL:             hds.m_data.MISSION_KEY_PERSONNEL = append(hds.m_data.MISSION_KEY_PERSONNEL, *v)    
//x         case MISSION_REFERENCES:                hds.m_data.MISSION_REFERENCES = *v
//x         case MISSION_REGION:                    hds.m_data.MISSION_REGION = append(hds.m_data.MISSION_REGION, *v)    
//x         case MISSION_CAVEATS:                   hds.m_data.MISSION_CAVEATS = append(hds.m_data.MISSION_CAVEATS, *v)    
//x         case EXPERIMENT:
//x         case EXPERIMENT_DESCRIPTION:
//x         case INVESTIGATOR_COORDINATES:
//x         case EXPERIMENT_REFERENCES:
//x         case EXPERIMENT_KEY_PERSONNEL:
//x         case EXPERIMENT_CAVEATS:
//x         case OBSERVATORY:
//x         case OBSERVATORY_CAVEATS:
//x         case OBSERVATORY_DESCRIPTION:
//x         case OBSERVATORY_TIME_SPAN:
//x         case OBSERVATORY_REGION:
//x         case INSTRUMENT_NAME:
//x         case INSTRUMENT_DESCRIPTION:
//x         case INSTRUMENT_TYPE:
//x         case MEASUREMENT_TYPE:
//x         case INSTRUMENT_CAVEATS:
//x         case DATASET_ID:
//x         case DATA_TYPE:
//x         case DATASET_TITLE:
//x         case DATASET_DESCRIPTION:
//x         case CONTACT_COORDINATES:
//x         case TIME_RESOLUTION:
//x         case MIN_TIME_RESOLUTION:
//x         case MAX_TIME_RESOLUTION:
//x         case PROCESSING_LEVEL:
//x         case ACKNOWLEDGEMENT:
//x         case DATASET_CAVEATS:
//x         case DATASET_VERSION:
//x         case FILE_TYPE:
//x         case METADATA_TYPE:
//x         case METADATA_VERSION:
//x         default:
//x     }
//x     
//x     return 
//x } 


func (hds *CefHeaderData) v_set(io_data, v *string)  (err error){
    // check value_type here
    *io_data = *v
    return
}

func (hds *CefHeaderData) v_add(io_data, v *string)  (err error){
    // check value_type here
    *io_data += (` ` + *v)
    return
}


func (hds *CefHeaderData) v_append(io_data *[]string, v *string)  (err error){
    // check value_type here
    *io_data  = append(*io_data, *v)    
    return
}


func (hds *CefHeaderData) kv_meta_entry(v *string)  (err error){

    //x fmt.Println("##  #  #  #  #  #  #  #  #  #  #   # # # ##", hds.m_meta, *k, *v)

    switch hds.m_meta {
        
        case MISSION:                           err = hds.v_set(&hds.m_data.MISSION, v)
        case MISSION_TIME_SPAN:                 err = hds.v_set(&hds.m_data.MISSION_TIME_SPAN, v)
        case MISSION_AGENCY:                    err = hds.v_set(&hds.m_data.MISSION_AGENCY, v)
        case MISSION_DESCRIPTION:               err = hds.v_add(&hds.m_data.MISSION_DESCRIPTION, v)
        case MISSION_KEY_PERSONNEL:             err = hds.v_append(&hds.m_data.MISSION_KEY_PERSONNEL, v)    
        case MISSION_REFERENCES:                err = hds.v_set(&hds.m_data.MISSION_REFERENCES, v)
        case MISSION_REGION:                    err = hds.v_append(&hds.m_data.MISSION_REGION, v)    
        case MISSION_CAVEATS:                   err = hds.v_append(&hds.m_data.MISSION_CAVEATS, v)   
        
        case EXPERIMENT:                        err = hds.v_set(&hds.m_data.EXPERIMENTS.EXPERIMENT, v)
        case EXPERIMENT_DESCRIPTION:            err = hds.v_add(&hds.m_data.EXPERIMENTS.EXPERIMENT_DESCRIPTION, v)
        case INVESTIGATOR_COORDINATES:          err = hds.v_append(&hds.m_data.EXPERIMENTS.INVESTIGATOR_COORDINATES, v)
        case EXPERIMENT_REFERENCES:             err = hds.v_append(&hds.m_data.EXPERIMENTS.EXPERIMENT_REFERENCES, v)
        case EXPERIMENT_KEY_PERSONNEL:          err = hds.v_append(&hds.m_data.EXPERIMENTS.EXPERIMENT_KEY_PERSONNEL, v)
        case EXPERIMENT_CAVEATS:                err = hds.v_add(&hds.m_data.EXPERIMENTS.EXPERIMENT_CAVEATS, v)
        
        case OBSERVATORY:                       err = hds.v_append(&hds.m_data.OBSERVATORIES.OBSERVATORY, v)
        case OBSERVATORY_CAVEATS:               err = hds.v_append(&hds.m_data.OBSERVATORIES.OBSERVATORY_CAVEATS, v)
        case OBSERVATORY_DESCRIPTION:           err = hds.v_add(&hds.m_data.OBSERVATORIES.OBSERVATORY_DESCRIPTION, v)
        case OBSERVATORY_TIME_SPAN:             err = hds.v_set(&hds.m_data.OBSERVATORIES.OBSERVATORY_TIME_SPAN, v)
        case OBSERVATORY_REGION:                err = hds.v_append(&hds.m_data.OBSERVATORIES.OBSERVATORY_REGION, v)
        
        case INSTRUMENT_NAME:                   err = hds.v_append(&hds.m_data.EXPERIMENTS.INSTRUMENTS.INSTRUMENT_NAME, v)
        case INSTRUMENT_DESCRIPTION:            err = hds.v_add(&hds.m_data.EXPERIMENTS.INSTRUMENTS.INSTRUMENT_DESCRIPTION, v)
        case INSTRUMENT_TYPE:                   err = hds.v_append(&hds.m_data.EXPERIMENTS.INSTRUMENTS.INSTRUMENT_TYPE, v)
        case MEASUREMENT_TYPE:                  err = hds.v_append(&hds.m_data.EXPERIMENTS.INSTRUMENTS.MEASUREMENT_TYPE, v)
        case INSTRUMENT_CAVEATS:                err = hds.v_append(&hds.m_data.EXPERIMENTS.INSTRUMENTS.INSTRUMENT_CAVEATS, v)
        
        case DATASET_ID:                        err = hds.v_set(&hds.m_data.DATASETS.DATASET_ID, v)
        case DATA_TYPE:                         err = hds.v_set(&hds.m_data.DATASETS.DATA_TYPE, v)
        case DATASET_TITLE:                     err = hds.v_set(&hds.m_data.DATASETS.DATASET_TITLE, v)
        case DATASET_DESCRIPTION:               err = hds.v_add(&hds.m_data.DATASETS.DATASET_DESCRIPTION, v)
        
        case CONTACT_COORDINATES:               err = hds.v_append(&hds.m_data.DATASETS.CONTACT_COORDINATES, v)
        
        case PROCESSING_LEVEL:                  err = hds.v_set(&hds.m_data.DATASETS.PROCESSING_LEVEL, v)
        case TIME_RESOLUTION:                   err = hds.v_set(&hds.m_data.DATASETS.TIME_RESOLUTION, v)
        case MIN_TIME_RESOLUTION:               err = hds.v_set(&hds.m_data.DATASETS.MIN_TIME_RESOLUTION, v)
        case MAX_TIME_RESOLUTION:               err = hds.v_set(&hds.m_data.DATASETS.MAX_TIME_RESOLUTION, v)
        case DATASET_CAVEATS:                   err = hds.v_add(&hds.m_data.DATASETS.DATASET_CAVEATS, v)
        case ACKNOWLEDGEMENT:                   err = hds.v_set(&hds.m_data.DATASETS.ACKNOWLEDGEMENT, v)


        case DATASET_VERSION:                   err = hds.v_set(&hds.m_data.DATASETS.FILE.DATASET_VERSION, v)
        case FILE_TYPE:                         err = hds.v_set(&hds.m_data.DATASETS.FILE.FILE_TYPE, v)
        case METADATA_TYPE:                     err = hds.v_set(&hds.m_data.DATASETS.FILE.METADATA_TYPE, v)
        case METADATA_VERSION:                  err = hds.v_set(&hds.m_data.DATASETS.FILE.METADATA_VERSION, v)

        case LOGICAL_FILE_ID :                  err = hds.v_set(&hds.m_data.DATASETS.FILE.LOGICAL_FILE_ID, v)
        case VERSION_NUMBER:                    err = hds.v_set(&hds.m_data.DATASETS.FILE.VERSION_NUMBER, v)
        case FILE_TIME_SPAN:                    err = hds.v_set(&hds.m_data.DATASETS.FILE.FILE_TIME_SPAN, v)
        case GENERATION_DATE:                   err = hds.v_set(&hds.m_data.DATASETS.FILE.GENERATION_DATE, v)
        case FILE_CAVEATS:                      err = hds.v_set(&hds.m_data.DATASETS.FILE.FILE_CAVEATS, v)


        
        default:
    }
    
    return 
} 

func (hds *CefHeaderData) kv_meta_value_type(v *string)  (err error){
    
    return 
} 


func (hds *CefHeaderData) kv_meta(k, v *string)  (err error){
    
    switch {
        case strings.EqualFold("ENTRY", *k) == true: 
            err = hds.kv_meta_entry(v)
        case strings.EqualFold("VALUE_TYPE", *k) == true:            
            err = hds.kv_meta_value_type(v)
        default:
            return errors.New("META KEY: unknown")
    }
    
    return 
}

