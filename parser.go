package main

import (
//x 	"bufio"
//x 	"compress/gzip"
 	"strings"
	"fmt"
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
}
 
///////////////////////////////////////////////////////////////////////////////
//

func (hds *CefHeaderData) add_kv(k, v *string)  (err error) {
    
//x     fmt.Println("--", *k, *v)

    //x ATTR, META, VAR, ERROR
    switch {
        case strings.EqualFold("START_META", *k) == true :
            fmt.Println("START_META", *v)
    
            switch hds.m_state {
                case ATTR:
                    hds.m_name = *v
                    hds.m_state = META
                    
                    hds.m_meta, err = getMeta(*v) //(m eMeta, err error) 
                    if(err != nil) {
                        return err
                    }
                    
                    //x fmt.Println("################################## ", hds.m_meta, *v)
                    
                default:
                    return errors.New("START_META: invalid State")
            }
            
        case strings.EqualFold("START_VARIABLE", *k) == true :
            fmt.Println("START_VARIABLE", *v)
        
            switch hds.m_state {
                case ATTR:
                    hds.m_name = *v
                    hds.m_state = VAR
                default:
                    return errors.New("START_VARIABLE: invalid State")
            }
        
        case strings.EqualFold("END_META", *k) == true :
            fmt.Println("END_META", *v)

            switch hds.m_state {
                case META:
                    if hds.m_name != *v {
                        return errors.New("END_META: invalid Name")
                    }
                    hds.m_state = ATTR
                default:
                    return errors.New("END_META: invalid State")
            }

        case strings.EqualFold("END_VARIABLE", *k) == true :
            fmt.Println("END_VARIABLE", *v)
                    
            switch hds.m_state {
                case VAR:
                    if hds.m_name != *v {
                        return errors.New("END_VARIABLE: invalid Name")
                    }
                    hds.m_state = ATTR
                default:
                    return errors.New("END_VARIABLE: invalid State")
            }
                    
        default :
            fmt.Println("  ", *k, *v)

            switch hds.m_state {
                case ATTR:
                    err = hds.kv_attr(k,v)
                case META:
                    err = hds.kv_meta(k,v)
                case VAR:
                    err = hds.kv_var(k,v)
                default:
                    return errors.New("DEFAULT K,V: invalid State")

            }
    }   
    
    return err
}
