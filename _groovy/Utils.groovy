package cefpass

import java.io.File

///////////////////////////////////////////////////////////////////////////////
//

public class Utils
{
    static def fileExists(i_path)          { File f = new File(i_path); (f != null && f.isFile()) }
    static def directoryExists(i_path)     { File f = new File(i_path); (f != null && f.isDirectory()) }

    static def fileExistsS(i_path)         { fileExists(i_path)         ? "OK" : "NOT FOUND" }
    static def directoryExistsS(i_path)    { directoryExists(i_path)    ? "OK" : "NOT FOUND" }
    
    static def isQuotedString(i_s)         { (i_s != null) && (i_s.size() >= 2) && (i_s[0] =='\"') && (i_s[i_s.size()-1]=='\"') }
    static def getUnQuotedString(i_s)      { 
        def r = null
//x         if(i_s != null)
        if((i_s != null) && isQuotedString(i_s))
        {   
            def l = i_s.size()
            if(l >= 2) r = i_s.substring(1, l-1)
        }
        else {
            r = i_s
        }
            
        r
    }

    // only the filename - not full path
    static def getFilename(i_path)         { File f = new File(i_path); ((f != null) ? f.getName() : null) }

    // just filename + no gz
    static def getCefFilename(i_filepath) { 
        def f = Utils.getFilename(i_filepath); 
        if(f != null) {
            def ix = f.lastIndexOf(".gz")
            f = (ix >= 0) ? f.substring(0, ix) : f
        }
        f
    }

    // filename less the ext    C3_CP_EDI_EGD__20111009_V01
    static def getCefLogicalFileId(i_filename) { 
        def ix = i_filename.indexOf('.')
        (ix >= 0) ? i_filename.substring(0,ix) : null
    }

    // just the number
    static def getCefFileVersion(i_filename) {
        def v = null
        def logical_str = Utils.getCefLogicalFileId(i_filename)
        if(logical_str != null) {
            def ix = logical_str.toUpperCase().lastIndexOf('_V')
            if(ix >= 0) v = logical_str.substring(ix+2)
        }
        
        v
    }
    
//x     static def getIntegerInRange(i_value, i_lo, i_hi, i_default) {
//x         def l_result = i_default
//x 
//x         if(DataTypes.isValidInt(i_value) == true)
//x         {
//x             try {
//x                 def l_value = Integer.parseInt(i_value);
//x                 if((l_value >= i_lo) && (l_value <= i_hi))
//x                     l_result = l_value
//x             } 
//x             catch (Exception e) { 
//x             }
//x         }
//x         
//x         l_result
//x     }
    
    static def getIntegerInRange(i_value, i_lo, i_hi, i_default) {
        def l_result = i_default

        if(i_value instanceof java.lang.String)
        {
            try {
                i_value = Integer.parseInt(i_value);
            } catch (Exception e) { }
        }
        
        if(i_value instanceof java.lang.Integer)
        {
            if((i_value >= i_lo) && (i_value <= i_hi))
                l_result = i_value
        }
        
        l_result
    }
}


