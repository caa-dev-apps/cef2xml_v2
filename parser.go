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

///////////////////////////////////////////////////////////////////////////////  ///////////////////////////////////////////////////////////////////////////////
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

///////////////////////////////////////////////////////////////////////////////  ///////////////////////////////////////////////////////////////////////////////
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
}
 
///////////////////////////////////////////////////////////////////////////////
//

func (hds *CefHeaderData) kv_meta(k, v *string)  error {
//x     err := hds.m_data.kv_meta(k, v)
    err := error(nil)

    hds.m_state = META
    return err
} 

func (hds *CefHeaderData) kv_var(k, v *string)  error {
//x     err := hds.m_data.kv_var(k, v)
    err := error(nil)

    hds.m_state = VAR
    return err
} 

func (hds *CefHeaderData) kv_attr(k, v *string)  error {
    
    err := hds.m_data.kv_attr(k, v)
    
    hds.m_state = ATTR
    return err
} 

///////////////////////////////////////////////////////////////////////////////
//

func (hds *CefHeaderData) add_kv(k, v *string)  (err error) {
//x     err := error(nil)
    
    fmt.Println("--", *k, *v)

    //x ATTR, META, VAR, ERROR
    switch {
        case strings.EqualFold("START_META", *k) == true :
        
            switch hds.m_state {
                case ATTR:
                    hds.m_name = *v
                    hds.m_state = META
                default:
                    return errors.New("START_META: invalid State")
            }
            
        case strings.EqualFold("START_VARIABLE", *k) == true :
        
            switch hds.m_state {
                case ATTR:
                    hds.m_name = *v
                    hds.m_state = VAR
                default:
                    return errors.New("START_VARIABLE: invalid State")
            }
        
        case strings.EqualFold("END_META", *k) == true :

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
