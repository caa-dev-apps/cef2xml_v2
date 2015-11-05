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
    
    m_data CAA_MetaData
}
 
func (hds *CefHeaderData) start_meta(v *string)  error {
    err := error(nil)
    
    fmt.Println("-x Start Meta", *v)
    
    hds.m_state = META
    return err
} 
 
func (hds *CefHeaderData) start_var(v *string)  error {
    err := error(nil)
    
    hds.m_state = VAR
    return err
} 
 
///////////////////////////////////////////////////////////////////////////////
//

func (hds *CefHeaderData) end_meta(v *string)  error {
    err := error(nil)
    
    fmt.Println("-x End Meta", *v)
    
    hds.m_state = ATTR
    return err
} 
 
func (hds *CefHeaderData) end_var(v *string)  error {
    err := error(nil)
    
    hds.m_state = ATTR
    return err
} 

///////////////////////////////////////////////////////////////////////////////
//

func (hds *CefHeaderData) kv_meta(k, v *string)  error {
    err := error(nil)

    hds.m_state = META
    return err
} 

func (hds *CefHeaderData) kv_var(k, v *string)  error {
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

 
func (hds *CefHeaderData) set_error(message string)  error {

    err := errors.New(message)
    return err
} 
 

///////////////////////////////////////////////////////////////////////////////
//

func (hds *CefHeaderData) add_kv(k, v *string)  error {
    err := error(nil)
    
    fmt.Println("--", *k, *v)

    //x ATTR, META, VAR, ERROR
    switch {
        case strings.EqualFold("START_META", *k) == true :
        
            switch hds.m_state {
                case ATTR:
                    err = hds.start_meta(v)
                default:
                    err = hds.set_error("START_META: invalid State")
            }
            
        case strings.EqualFold("START_VARIABLE", *k) == true :
        
            switch hds.m_state {
                case ATTR:
                    err = hds.start_var(v)
                default:
                    err = hds.set_error("START_VARIABLE: invalid State")
            }
        
        case strings.EqualFold("END_META", *k) == true :

            switch hds.m_state {
                case META:
                    err = hds.end_meta(v)
                default:
                    err = hds.set_error("END_META: invalid State")
            }

        case strings.EqualFold("END_VARIABLE", *k) == true :
                    
            switch hds.m_state {
                case VAR:
                    err = hds.end_var(v)
                default:
                    err = hds.set_error("END_VARIABLE: invalid State")
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
                    err = hds.set_error("DEFAULT K,V: invalid State")
            }
    }   
    
    return err
}
