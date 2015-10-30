package cefpass

import java.io.FileInputStream
import java.io.File
import java.util.zip.GZIPInputStream

//-----------------------------------------------------------------------------
//

public class CefContext {
    def m_filepath
    def m_level = 0
    def m_prefix = []       // e.g. [0.1.0.2]   indicates tree layout taking into account nested include files
    def m_lineCount = 0
    def m_filename_only
    
    def CefContext(i_filepath, i_level, i_prefix) {
        m_filepath = i_filepath
        m_level = i_level
        m_prefix = i_prefix
        
        m_filename_only = new File(i_filepath).getName()
    }
    
    def getString() { 
        return "Cx: " + m_level + "\t" + m_filepath  + "\t" + m_prefix 
    }
    
    def incLineCount= { m_lineCount++ }
    def diag = { it ->
        Show.showFileReadLineDetails("" + m_lineCount + "\t" + m_level + "\t" + m_filename_only + "\t\t\t\t" + it);
    }
}


