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

//old   func (hds *CefHeaderData) v_set(io_data, v *string)  (err error){
//old       // check value_type here
//old       *io_data = *v
//old       return
//old   }
//old   
//old   func (hds *CefHeaderData) v_add(io_data, v *string)  (err error){
//old       // check value_type here
//old       *io_data += (` ` + *v)
//old       return
//old   }
//old   
//old   
//old   func (hds *CefHeaderData) v_append(io_data *[]string, v *string)  (err error){
//old       // check value_type here
//old       *io_data  = append(*io_data, *v)    
//old       return
//old   }
//old   
//old   
//old   func (hds *CefHeaderData) kv_meta_entry(v *string)  (err error){
//old   
//old       //x fmt.Println("##  #  #  #  #  #  #  #  #  #  #   # # # ##", hds.m_meta, *k, *v)
//old   
//old       switch hds.m_meta {
//old           
//old           case MISSION:                           err = hds.v_set(&hds.m_data.MISSION, v)
//old           case MISSION_TIME_SPAN:                 err = hds.v_set(&hds.m_data.MISSION_TIME_SPAN, v)
//old           case MISSION_AGENCY:                    err = hds.v_set(&hds.m_data.MISSION_AGENCY, v)
//old           case MISSION_DESCRIPTION:               err = hds.v_add(&hds.m_data.MISSION_DESCRIPTION, v)
//old           case MISSION_KEY_PERSONNEL:             err = hds.v_append(&hds.m_data.MISSION_KEY_PERSONNEL, v)    
//old           case MISSION_REFERENCES:                err = hds.v_set(&hds.m_data.MISSION_REFERENCES, v)
//old           case MISSION_REGION:                    err = hds.v_append(&hds.m_data.MISSION_REGION, v)    
//old           case MISSION_CAVEATS:                   err = hds.v_append(&hds.m_data.MISSION_CAVEATS, v)   
//old           
//old           case EXPERIMENT:                        err = hds.v_set(&hds.m_data.EXPERIMENTS.EXPERIMENT, v)
//old           case EXPERIMENT_DESCRIPTION:            err = hds.v_add(&hds.m_data.EXPERIMENTS.EXPERIMENT_DESCRIPTION, v)
//old           case INVESTIGATOR_COORDINATES:          err = hds.v_append(&hds.m_data.EXPERIMENTS.INVESTIGATOR_COORDINATES, v)
//old           case EXPERIMENT_REFERENCES:             err = hds.v_append(&hds.m_data.EXPERIMENTS.EXPERIMENT_REFERENCES, v)
//old           case EXPERIMENT_KEY_PERSONNEL:          err = hds.v_append(&hds.m_data.EXPERIMENTS.EXPERIMENT_KEY_PERSONNEL, v)
//old           case EXPERIMENT_CAVEATS:                err = hds.v_add(&hds.m_data.EXPERIMENTS.EXPERIMENT_CAVEATS, v)
//old           
//old           case OBSERVATORY:                       err = hds.v_append(&hds.m_data.OBSERVATORIES.OBSERVATORY, v)
//old           case OBSERVATORY_CAVEATS:               err = hds.v_append(&hds.m_data.OBSERVATORIES.OBSERVATORY_CAVEATS, v)
//old           case OBSERVATORY_DESCRIPTION:           err = hds.v_add(&hds.m_data.OBSERVATORIES.OBSERVATORY_DESCRIPTION, v)
//old           case OBSERVATORY_TIME_SPAN:             err = hds.v_set(&hds.m_data.OBSERVATORIES.OBSERVATORY_TIME_SPAN, v)
//old           case OBSERVATORY_REGION:                err = hds.v_append(&hds.m_data.OBSERVATORIES.OBSERVATORY_REGION, v)
//old           
//old           case INSTRUMENT_NAME:                   err = hds.v_append(&hds.m_data.EXPERIMENTS.INSTRUMENTS.INSTRUMENT_NAME, v)
//old           case INSTRUMENT_DESCRIPTION:            err = hds.v_add(&hds.m_data.EXPERIMENTS.INSTRUMENTS.INSTRUMENT_DESCRIPTION, v)
//old           case INSTRUMENT_TYPE:                   err = hds.v_append(&hds.m_data.EXPERIMENTS.INSTRUMENTS.INSTRUMENT_TYPE, v)
//old           case MEASUREMENT_TYPE:                  err = hds.v_append(&hds.m_data.EXPERIMENTS.INSTRUMENTS.MEASUREMENT_TYPE, v)
//old           case INSTRUMENT_CAVEATS:                err = hds.v_append(&hds.m_data.EXPERIMENTS.INSTRUMENTS.INSTRUMENT_CAVEATS, v)
//old           
//old           case DATASET_ID:                        err = hds.v_set(&hds.m_data.DATASETS.DATASET_ID, v)
//old           case DATA_TYPE:                         err = hds.v_set(&hds.m_data.DATASETS.DATA_TYPE, v)
//old           case DATASET_TITLE:                     err = hds.v_set(&hds.m_data.DATASETS.DATASET_TITLE, v)
//old           case DATASET_DESCRIPTION:               err = hds.v_add(&hds.m_data.DATASETS.DATASET_DESCRIPTION, v)
//old           
//old           case CONTACT_COORDINATES:               err = hds.v_append(&hds.m_data.DATASETS.CONTACT_COORDINATES, v)
//old           
//old           case PROCESSING_LEVEL:                  err = hds.v_set(&hds.m_data.DATASETS.PROCESSING_LEVEL, v)
//old           case TIME_RESOLUTION:                   err = hds.v_set(&hds.m_data.DATASETS.TIME_RESOLUTION, v)
//old           case MIN_TIME_RESOLUTION:               err = hds.v_set(&hds.m_data.DATASETS.MIN_TIME_RESOLUTION, v)
//old           case MAX_TIME_RESOLUTION:               err = hds.v_set(&hds.m_data.DATASETS.MAX_TIME_RESOLUTION, v)
//old           case DATASET_CAVEATS:                   err = hds.v_add(&hds.m_data.DATASETS.DATASET_CAVEATS, v)
//old           case ACKNOWLEDGEMENT:                   err = hds.v_set(&hds.m_data.DATASETS.ACKNOWLEDGEMENT, v)
//old   
//old   
//old           case DATASET_VERSION:                   err = hds.v_set(&hds.m_data.DATASETS.FILE.DATASET_VERSION, v)
//old           case FILE_TYPE:                         err = hds.v_set(&hds.m_data.DATASETS.FILE.FILE_TYPE, v)
//old           case METADATA_TYPE:                     err = hds.v_set(&hds.m_data.DATASETS.FILE.METADATA_TYPE, v)
//old           case METADATA_VERSION:                  err = hds.v_set(&hds.m_data.DATASETS.FILE.METADATA_VERSION, v)
//old   
//old           case LOGICAL_FILE_ID :                  err = hds.v_set(&hds.m_data.DATASETS.FILE.LOGICAL_FILE_ID, v)
//old           case VERSION_NUMBER:                    err = hds.v_set(&hds.m_data.DATASETS.FILE.VERSION_NUMBER, v)
//old           case FILE_TIME_SPAN:                    err = hds.v_set(&hds.m_data.DATASETS.FILE.FILE_TIME_SPAN, v)
//old           case GENERATION_DATE:                   err = hds.v_set(&hds.m_data.DATASETS.FILE.GENERATION_DATE, v)
//old           case FILE_CAVEATS:                      err = hds.v_set(&hds.m_data.DATASETS.FILE.FILE_CAVEATS, v)
//old   
//old   
//old           
//old           default:
//old       }
//old       
//old       return 
//old   } 
//old   
//old   func (hds *CefHeaderData) kv_meta_value_type(v *string)  (err error){
//old       
//old       return 
//old   } 
//old   
//old   
//old   func (hds *CefHeaderData) kv_meta(k, v *string)  (err error){
//old       
//old       switch {
//old           case strings.EqualFold("ENTRY", *k) == true: 
//old               err = hds.kv_meta_entry(v)
//old           case strings.EqualFold("VALUE_TYPE", *k) == true:            
//old               err = hds.kv_meta_value_type(v)
//old           default:
//old               return errors.New("META KEY: unknown")
//old       }
//old       
//old       return 
//old   }
//old   

//old   func (hds *CefHeaderData) v_set(io_data, v *string)  (err error){
//old       // check value_type here
//old       *io_data = *v
//old       return
//old   }
//old   
//old   func (hds *CefHeaderData) v_add(io_data, v *string)  (err error){
//old       // check value_type here
//old       *io_data += (` ` + *v)
//old       return
//old   }
//old   
//old   
//old   func (hds *CefHeaderData) v_append(io_data *[]string, v *string)  (err error){
//old       // check value_type here
//old       *io_data  = append(*io_data, *v)    
//old       return
//old   }

func (hds *CefHeaderData) v_set(io_data *string, kv *KeyVal)  (err error){
    // check value_type here
//x     *io_data = kv.val[0]
    
    *io_data = val_string(kv)
    return
}

func (hds *CefHeaderData) v_add(io_data *string, kv *KeyVal)  (err error){
    // check value_type here
//x     *io_data += (` ` + kv.val[0])
    *io_data += (` ` + val_string(kv))
    
    return
}


func (hds *CefHeaderData) v_append(io_data *[]string, kv *KeyVal)  (err error){
    // check value_type here
//x     *io_data  = append(*io_data, *v)    
//x     *io_data  = append(*io_data, kv.val...)    
    *io_data  = append(*io_data, val_string(kv))    
    return
}



//x func (hds *CefHeaderData) kv_meta_entry(v *string)  (err error){
func (hds *CefHeaderData) kv_meta_entry(kv *KeyVal)  (err error){

    //x fmt.Println("##  #  #  #  #  #  #  #  #  #  #   # # # ##", hds.m_meta, *k, *v)

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

//x func (hds *CefHeaderData) kv_meta_value_type(v *string)  (err error){
func (hds *CefHeaderData) kv_meta_value_type(kv *KeyVal)  (err error){
    
    return 
} 

func (hds *CefHeaderData) kv_meta(kv *KeyVal)  (err error){
    
    switch {
//x     case strings.EqualFold("ENTRY", *k) == true: 
        case strings.EqualFold("ENTRY", kv.key) == true: 
//x             err = hds.kv_meta_entry(v)
            err = hds.kv_meta_entry(kv)
//x         case strings.EqualFold("VALUE_TYPE", *k) == true:            
        case strings.EqualFold("VALUE_TYPE", kv.key) == true:            
//x             err = hds.kv_meta_value_type(v)
            err = hds.kv_meta_value_type(kv)
        default:
            return errors.New("META KEY: unknown")
    }
    
    return 
}


