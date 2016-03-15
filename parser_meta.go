package main

import (
 	"strings"
    "errors"
)


func (hds *CefHeaderData) v_set(io_data *string, kv *KeyVal)  (err error){
    *io_data = val_string(kv)
    return
}

func (hds *CefHeaderData) v_add(io_data *string, kv *KeyVal)  (err error){
    *io_data += (` ` + val_string(kv))
    return
}


func (hds *CefHeaderData) v_append(io_data *[]string, kv *KeyVal)  (err error){
    *io_data  = append(*io_data, val_string(kv))    
    return
}

func (hds *CefHeaderData) kv_meta_entry(kv *KeyVal)  (err error){

    switch hds.m_meta {
        
        case MISSION:                           err = hds.v_set(&hds.m_data.MISSION, kv)
        case MISSION_TIME_SPAN:                 err = hds.v_set(&hds.m_data.MISSION_TIME_SPAN, kv)
        case MISSION_AGENCY:                    err = hds.v_set(&hds.m_data.MISSION_AGENCY, kv)
        case MISSION_DESCRIPTION:               err = hds.v_add(&hds.m_data.MISSION_DESCRIPTION, kv)
        case MISSION_KEY_PERSONNEL:             err = hds.v_append(&hds.m_data.MISSION_KEY_PERSONNEL, kv)    
        case MISSION_REFERENCES:                err = hds.v_set(&hds.m_data.MISSION_REFERENCES, kv)
        case MISSION_REGION:                    err = hds.v_append(&hds.m_data.MISSION_REGION, kv)    
        case MISSION_CAVEATS:                   err = hds.v_append(&hds.m_data.MISSION_CAVEATS, kv)   
        
        case EXPERIMENT:                        err = hds.v_set(&hds.m_data.EXPERIMENTS.EXPERIMENT, kv)
        case EXPERIMENT_DESCRIPTION:            err = hds.v_add(&hds.m_data.EXPERIMENTS.EXPERIMENT_DESCRIPTION, kv)
        case INVESTIGATOR_COORDINATES:          err = hds.v_append(&hds.m_data.EXPERIMENTS.INVESTIGATOR_COORDINATES, kv)
        case EXPERIMENT_REFERENCES:             err = hds.v_append(&hds.m_data.EXPERIMENTS.EXPERIMENT_REFERENCES, kv)
        case EXPERIMENT_KEY_PERSONNEL:          err = hds.v_append(&hds.m_data.EXPERIMENTS.EXPERIMENT_KEY_PERSONNEL, kv)
        case EXPERIMENT_CAVEATS:                err = hds.v_add(&hds.m_data.EXPERIMENTS.EXPERIMENT_CAVEATS, kv)
        
        case OBSERVATORY:                       err = hds.v_append(&hds.m_data.OBSERVATORIES.OBSERVATORY, kv)
        case OBSERVATORY_CAVEATS:               err = hds.v_append(&hds.m_data.OBSERVATORIES.OBSERVATORY_CAVEATS, kv)
        case OBSERVATORY_DESCRIPTION:           err = hds.v_add(&hds.m_data.OBSERVATORIES.OBSERVATORY_DESCRIPTION, kv)
        case OBSERVATORY_TIME_SPAN:             err = hds.v_set(&hds.m_data.OBSERVATORIES.OBSERVATORY_TIME_SPAN, kv)
        case OBSERVATORY_REGION:                err = hds.v_append(&hds.m_data.OBSERVATORIES.OBSERVATORY_REGION, kv)
        
        case INSTRUMENT_NAME:                   err = hds.v_append(&hds.m_data.EXPERIMENTS.INSTRUMENTS.INSTRUMENT_NAME, kv)
        case INSTRUMENT_DESCRIPTION:            err = hds.v_add(&hds.m_data.EXPERIMENTS.INSTRUMENTS.INSTRUMENT_DESCRIPTION, kv)
        case INSTRUMENT_TYPE:                   err = hds.v_append(&hds.m_data.EXPERIMENTS.INSTRUMENTS.INSTRUMENT_TYPE, kv)
        case MEASUREMENT_TYPE:                  err = hds.v_append(&hds.m_data.EXPERIMENTS.INSTRUMENTS.MEASUREMENT_TYPE, kv)
        case INSTRUMENT_CAVEATS:                err = hds.v_append(&hds.m_data.EXPERIMENTS.INSTRUMENTS.INSTRUMENT_CAVEATS, kv)
        
        case DATASET_ID:                        err = hds.v_set(&hds.m_data.DATASETS.DATASET_ID, kv)
        case DATA_TYPE:                         err = hds.v_set(&hds.m_data.DATASETS.DATA_TYPE, kv)
        case DATASET_TITLE:                     err = hds.v_set(&hds.m_data.DATASETS.DATASET_TITLE, kv)
        case DATASET_DESCRIPTION:               err = hds.v_add(&hds.m_data.DATASETS.DATASET_DESCRIPTION, kv)
        
        case CONTACT_COORDINATES:               err = hds.v_append(&hds.m_data.DATASETS.CONTACT_COORDINATES, kv)
        
        case PROCESSING_LEVEL:                  err = hds.v_set(&hds.m_data.DATASETS.PROCESSING_LEVEL, kv)
        case TIME_RESOLUTION:                   err = hds.v_set(&hds.m_data.DATASETS.TIME_RESOLUTION, kv)
        case MIN_TIME_RESOLUTION:               err = hds.v_set(&hds.m_data.DATASETS.MIN_TIME_RESOLUTION, kv)
        case MAX_TIME_RESOLUTION:               err = hds.v_set(&hds.m_data.DATASETS.MAX_TIME_RESOLUTION, kv)
        case DATASET_CAVEATS:                   err = hds.v_add(&hds.m_data.DATASETS.DATASET_CAVEATS, kv)
        case ACKNOWLEDGEMENT:                   err = hds.v_set(&hds.m_data.DATASETS.ACKNOWLEDGEMENT, kv)


        case DATASET_VERSION:                   err = hds.v_set(&hds.m_data.DATASETS.FILE.DATASET_VERSION, kv)
        case FILE_TYPE:                         err = hds.v_set(&hds.m_data.DATASETS.FILE.FILE_TYPE, kv)
        case METADATA_TYPE:                     err = hds.v_set(&hds.m_data.DATASETS.FILE.METADATA_TYPE, kv)
        case METADATA_VERSION:                  err = hds.v_set(&hds.m_data.DATASETS.FILE.METADATA_VERSION, kv)

        case LOGICAL_FILE_ID :                  err = hds.v_set(&hds.m_data.DATASETS.FILE.LOGICAL_FILE_ID, kv)
        case VERSION_NUMBER:                    err = hds.v_set(&hds.m_data.DATASETS.FILE.VERSION_NUMBER, kv)
        case FILE_TIME_SPAN:                    err = hds.v_set(&hds.m_data.DATASETS.FILE.FILE_TIME_SPAN, kv)
        case GENERATION_DATE:                   err = hds.v_set(&hds.m_data.DATASETS.FILE.GENERATION_DATE, kv)
        case FILE_CAVEATS:                      err = hds.v_set(&hds.m_data.DATASETS.FILE.FILE_CAVEATS, kv)
        
        default:
    }
    
    return 
} 

func (hds *CefHeaderData) kv_meta_value_type(kv *KeyVal)  (err error){
    
    return 
} 

func (hds *CefHeaderData) kv_meta(kv *KeyVal)  (err error){
    
    switch {
        case strings.EqualFold("ENTRY", kv.key) == true: 
            err = hds.kv_meta_entry(kv)
        case strings.EqualFold("VALUE_TYPE", kv.key) == true:            
            err = hds.kv_meta_value_type(kv)
        default:
            return errors.New("META KEY: unknown")
    }
    
    return 
}


