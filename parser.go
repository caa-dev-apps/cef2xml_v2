package main

import (
//x 	"bufio"
//x 	"compress/gzip"
 	"strings"
//x 	"fmt"
    "errors"
//x 	"os"
//x 	"regexp"
)

///////////////////////////////////////////////////////////////////////////////  
//

//          Current State	                Key In	            Val In	            Checks	                            Output function	            Next State
//          ATTR                            K	                V	                Well known k	                    Data.File.k = v	            A
//                                          K	                V	                Unexpected		                                                Error
//                                          META START	        NAME	            First Occurrence + Name             Expected Data.M[n] = v	    M
//                                                                                  Unexpected		                                                Error
//                                          VAR START	        NAME	            First Occurrence + Name Expected	Data.V[n] = v	            V
//                                                                                  Unexpected		                                                Error
//--------------------------------------------------------------------------------------------------------------------------------------------------------------
//          META                            K	                V	                Key Expected	                    Data.M[k] = v	            M
//                                          K	                V	                Key Unexpected		                                            Error
//                                          META START				                                                                                Error
//                                          VAR START				                                                                                Error
//                                          VAR END				                                                                                    Error
//                                          META_END	        V	                Name Matches start		                                        A
//                                          META_END	        V	                Name mismatch		                                            Error
//--------------------------------------------------------------------------------------------------------------------------------------------------------------
//          VAR                             K	                V	                Key Expected	                    Data.M[k] = v	            V
//                                          K	                V	                Key Unexpected		                                            Error
//                                          VAR START				                                                                                Error
//                                          META START				                                                                                Error
//                                          META END				                                                                                Error
//                                          VAR_END	            V	                Name Matches start		                                        A
//                                          VAR_END	            V	                Name mismatch		                                            Error

///////////////////////////////////////////////////////////////////////////////  
//

type State int

const (
 ATTR State = iota
 META
 VAR
 ERROR
 )

type CefHeaderData struct {
    m_state State
    m_name string 
    m_data CAA_MetaData
    m_meta eMeta
    m_var Parameter
}
 
///////////////////////////////////////////////////////////////////////////////
//

//old   func (hds *CefHeaderData) add_kv(k, v *string)  (err error) {
//old       
//old   //x     fmt.Println("--", *k, *v)
//old   
//old       //x ATTR, META, VAR, ERROR
//old       switch {
//old           case strings.EqualFold("START_META", *k) == true :
//old               mooi_log("START_META", *v)
//old       
//old               switch hds.m_state {
//old                   case ATTR:
//old                       hds.m_name = *v
//old                       hds.m_state = META
//old                       
//old                       hds.m_meta, err = getMeta(*v) //(m eMeta, err error) 
//old                       if(err != nil) {
//old                           return err
//old                       }
//old                       
//old                   default:
//old                       return errors.New("START_META: invalid State")
//old               }
//old               
//old           case strings.EqualFold("START_VARIABLE", *k) == true :
//old               mooi_log("START_VARIABLE", *v)
//old           
//old               switch hds.m_state {
//old                   case ATTR:
//old                       hds.m_name = *v
//old                       hds.m_state = VAR
//old                       
//old                       err = hds.stx_var(v)
//old                   default:
//old                       return errors.New("START_VARIABLE: invalid State")
//old               }
//old           
//old           case strings.EqualFold("END_META", *k) == true :
//old               mooi_log("END_META", *v)
//old   
//old               switch hds.m_state {
//old                   case META:
//old                       if hds.m_name != *v {
//old                           return errors.New("END_META: invalid Name")
//old                       }
//old                       hds.m_state = ATTR
//old                   default:
//old                       return errors.New("END_META: invalid State")
//old               }
//old   
//old           case strings.EqualFold("END_VARIABLE", *k) == true :
//old               mooi_log("END_VARIABLE", *v)
//old                       
//old               switch hds.m_state {
//old                   case VAR:
//old                       if hds.m_name != *v {
//old                           return errors.New("END_VARIABLE: invalid Name")
//old                       }
//old                       hds.m_state = ATTR
//old                       err = hds.etx_var()
//old                   default:
//old                       return errors.New("END_VARIABLE: invalid State")
//old               }
//old                       
//old           default :
//old               mooi_log("  ", *k, *v)
//old   
//old               switch hds.m_state {
//old                   case ATTR:
//old                       err = hds.kv_attr(k,v)
//old                   case META:
//old                       err = hds.kv_meta(k,v)
//old                   case VAR:
//old                       err = hds.kv_var(k,v)
//old                   default:
//old                       return errors.New("DEFAULT K,V: invalid State")
//old   
//old               }
//old       }   
//old       
//old       return err
//old   }


func (hds *CefHeaderData) add_kv(kv *KeyVal)  (err error) {
    
//x     fmt.Println("--", *k, *v)

    //x ATTR, META, VAR, ERROR
    switch {
//x         case strings.EqualFold("START_META", *k) == true :
//x             mooi_log("START_META", *v)
//x         case strings.EqualFold("START_META", *kv.key) == true :
        case strings.EqualFold("START_META", kv.key) == true :
            mooi_log("START_META", kv.val[0])
    
            switch hds.m_state {
                case ATTR:
//x                     hds.m_name = *v
                    hds.m_name = kv.val[0]
                    hds.m_state = META
                    
//x                     hds.m_meta, err = getMeta(*v) //(m eMeta, err error) 
                    hds.m_meta, err = getMeta(kv.val[0]) //(m eMeta, err error) 
                    if(err != nil) {
                        return err
                    }
                    
                default:
                    return errors.New("START_META: invalid State")
            }
            
//x         case strings.EqualFold("START_VARIABLE", *k) == true :
//x             mooi_log("START_VARIABLE", *v)
        case strings.EqualFold("START_VARIABLE", kv.key) == true :
            mooi_log("START_VARIABLE", kv.val[0])
        
            switch hds.m_state {
                case ATTR:
//x                     hds.m_name = *v
                    hds.m_name = kv.val[0]
                    hds.m_state = VAR
                    
//x                     err = hds.stx_var(v)
                    err = hds.stx_var(&kv.val[0])
                default:
                    return errors.New("START_VARIABLE: invalid State")
            }
        
//x         case strings.EqualFold("END_META", *k) == true :
//x             mooi_log("END_META", *v)
        case strings.EqualFold("END_META", kv.key) == true :
            mooi_log("END_META", kv.val[0])

            switch hds.m_state {
                case META:
//x                     if hds.m_name != *v {
                    if hds.m_name != kv.val[0] {
                        return errors.New("END_META: invalid Name")
                    }
                    hds.m_state = ATTR
                default:
                    return errors.New("END_META: invalid State")
            }

//x         case strings.EqualFold("END_VARIABLE", *k) == true :
//x             mooi_log("END_VARIABLE", *v)
        case strings.EqualFold("END_VARIABLE", kv.key) == true :
            mooi_log("END_VARIABLE", kv.val[0])
                    
            switch hds.m_state {
                case VAR:
//x                     if hds.m_name != *v {
                    if hds.m_name != kv.val[0] {
                        return errors.New("END_VARIABLE: invalid Name")
                    }
                    hds.m_state = ATTR
                    err = hds.etx_var()
                default:
                    return errors.New("END_VARIABLE: invalid State")
            }
                    
        default :
//x             mooi_log("  ", *k, *v)
            mooi_log("  ", *kv)

            switch hds.m_state {
                case ATTR:
//x                     err = hds.kv_attr(k,v)
                    err = hds.kv_attr(kv)
                case META:
//x                     err = hds.kv_meta(k,v)
                    err = hds.kv_meta(kv)
                case VAR:
//x                     err = hds.kv_var(k,v)
                    err = hds.kv_var(kv)
                default:
                    return errors.New("DEFAULT K,V: invalid State")
            }
    }   
    
    return err
}
