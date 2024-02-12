package main

import (
 	"strings"
    "errors"
)

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
 
func val_string(kv *KeyVal) (s string) {

    if len(kv.val) == 1 {
        s = strings.Trim(kv.val[0], `" `)
    } else if len(kv.val) > 1 {
        s = strings.Join(kv.val, ", ")
    }

    return
}

func (hds *CefHeaderData) add_kv(kv *KeyVal)  (err error) {
    

    switch {
        case strings.EqualFold("START_META", kv.key) == true :
	kv.val[0] = strings.ToUpper(kv.val[0])

            mooi_log("START_META", kv.val[0])
            switch hds.m_state {
                case ATTR:
                    hds.m_name = kv.val[0]
                    hds.m_state = META
                    
                    hds.m_meta, err = getMeta(kv.val[0]) 
                    if(err != nil) {
                        return err
                    }
                    
                default:
                    return errors.New("START_META: invalid State")
            }
            
        case strings.EqualFold("START_VARIABLE", kv.key) == true :
            mooi_log("START_VARIABLE", kv.val[0])
        
            switch hds.m_state {
                case ATTR:
                    hds.m_name = kv.val[0]
                    hds.m_state = VAR
                    
                    err = hds.stx_var(&kv.val[0])
                default:
                    return errors.New("START_VARIABLE: invalid State")
            }
        
        case strings.EqualFold("END_META", kv.key) == true :
	kv.val[0] = strings.ToUpper(kv.val[0])

            mooi_log("END_META", kv.val[0])

            switch hds.m_state {
                case META:
                    if hds.m_name != kv.val[0] {
                        return errors.New("END_META: invalid Name")
                    }
                    hds.m_state = ATTR
                default:
                    return errors.New("END_META: invalid State")
            }

        case strings.EqualFold("END_VARIABLE", kv.key) == true :
            mooi_log("END_VARIABLE", kv.val[0])
                    
            switch hds.m_state {
                case VAR:
                    if hds.m_name != kv.val[0] {
                        return errors.New("END_VARIABLE: invalid Name")
                    }
                    hds.m_state = ATTR
                    err = hds.etx_var()
                default:
                    return errors.New("END_VARIABLE: invalid State")
            }
                    
        default :
            mooi_log("  ", *kv)

            switch hds.m_state {
                case ATTR:
                    err = hds.kv_attr(kv)
                case META:
                    err = hds.kv_meta(kv)
                case VAR:
                    err = hds.kv_var(kv)
                default:
                    return errors.New("DEFAULT K,V: invalid State")
            }
    }   
    
    return err
}
