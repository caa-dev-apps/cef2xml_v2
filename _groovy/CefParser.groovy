package cefpass

import exceptions.RS0_CefParserException

///////////////////////////////////////////////////////////////////////////////
//

public class CefParser
{
    static def s_var = null
    static def s_meta = null
    static def s_k = null
    static def s_v = null
    static def s_includeFile = null
    static def s_readLine = null
    
    static def init() { 
        s_var = null
        s_meta = null
        s_k = null
        s_v = null
        s_includeFile = null
        s_readLine = null
    }

    static def save_kv(k, v)        
    { 
        s_k = k; 
        s_v = v; 
    }
    
    static def fatal_error(i_error)
    {
        throw new RS0_CefParserException(i_error)
    }
    
    static def stxMeta(v)
    {
        if(s_meta != null)                  fatal_error(RS0_CefParserException.Error.R_0_00___START_META___META_UNCLOSED)
        else if(s_var != null)              fatal_error(RS0_CefParserException.Error.R_0_01___START_META___VARIABLE_UNCLOSED)
        else if(v==null || v.length()==0)   fatal_error(RS0_CefParserException.Error.R_0_02___START_META___NAME_ERROR)
        else                                s_meta = v
    }
    
    static def etxMeta(v) 
    {
        if(s_var != null)                   fatal_error(RS0_CefParserException.Error.R_0_10___END_META___VARIABLE_OPENED)         // NEW
        else if(s_meta == null)             fatal_error(RS0_CefParserException.Error.R_0_11___END_META___META_UNOPENED)
        else if(s_meta != v)                fatal_error(RS0_CefParserException.Error.R_0_12___END_META___NAME_ERROR)
        else                                s_meta = null
    }


    static def stxVar(v) 
    {
        if(s_var != null)                   fatal_error(RS0_CefParserException.Error.R_0_20___START_VARIABLE___VARIABLE_UNCLOSED)
        else if(s_meta != null)             fatal_error(RS0_CefParserException.Error.R_0_21___START_VARIABLE___META_UNCLOSED)
        else if(v==null || v.length()==0)   fatal_error(RS0_CefParserException.Error.R_0_22___START_VARIABLE___NAME_ERROR)
        else                                s_var = v
    }
    
    static def etxVar(v) 
    {
        if(s_meta != null)                  fatal_error(RS0_CefParserException.Error.R_0_30___END_VARIABLE___META_OPENED)            // NEW
        else if(s_var == null)              fatal_error(RS0_CefParserException.Error.R_0_31___END_VARIABLE___VARIABLE_UNOPENED)
        else if(s_var != v)                 fatal_error(RS0_CefParserException.Error.R_0_32___END_VARIABLE___NAME_ERROR)
        else                                s_var = null
    }
    
    static def test_string_quotes(k,v)
    {
        def l_isOk = false
        
        if(v != null)
        {
            def l = v.size()
            if(l > 0)
            {
                def c1 = v[0]
                def c2 = v[l-1]
                
                if((((c1 == '\'') || (c1 == '\"') || (c2 == '\'') || (c2 == '\"'))  && 
                    ((c1 != c2) || (l == 1))) == false)
                {
                    l_isOk = true
                }
            }
        }
        
        if(l_isOk == false) {
            fatal_error(RS0_CefParserException.Error.R_0_51___MALFORMED_STRING_QUOTES)
        }
    }
    
    ///////////////////////////////////////////////////////////////////////////////
    //

    static def includeFileDuplicate(f)
    {
        s_includeFile = f
        fatal_error(RS0_CefParserException.Error.R_0_40___INCLUDE_FILE_DUPLICATE)
    }
    
    static def includeFileNotFound(f)
    {
        s_includeFile = f
        fatal_error(RS0_CefParserException.Error.R_0_41___INCLUDE_FILE_UNFOUND)
    }

    static def includeFileLevel8(f)
    {
        s_includeFile = f
        fatal_error(RS0_CefParserException.Error.R_0_42___INCLUDE_FILE_LEVEL_8)
    }
    
    static def MalFormedReadLine(s)
    {
        s_readLine = s
        fatal_error(RS0_CefParserException.Error.R_0_50___MALFORMED_READ_LINE)
    }
    
    ///////////////////////////////////////////////////////////////////////////////
    //

    static def test_kv(k, v)
    {
        save_kv(k,v)
        
        if("START_META".compareToIgnoreCase(k) == 0)            stxMeta(v) 
        else if ("END_META".compareToIgnoreCase(k) == 0)		etxMeta(v) 
        else if ("START_VARIABLE".compareToIgnoreCase(k) == 0) 	stxVar(v) 
        else if ("END_VARIABLE".compareToIgnoreCase(k) == 0)	etxVar(v) 
        else test_string_quotes(k,v)
    }
}


