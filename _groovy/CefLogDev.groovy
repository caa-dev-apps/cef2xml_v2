package cefpass

///////////////////////////////////////////////////////////////////////////////
// '(Optional) Output Results Level(0,1,2,3): 0:Debug 1:Info 2:Stages 3:Result Only')

public class CefLogDev
{
    enum Type {
        diag,
        info,
        exception,
        warn,
        error,
        none_today_thanks,
    }

    static Type s_type = Type.diag
    static boolean isEnabled = true
    
    static def pp =                                         { i_str -> println i_str}
//x     static def p(i_type, i_str)                             { if(isEnabled && i_type >= s_type) pp i_str 
//x                                                               else                              pp "x: " + i_str 
//x                                                             }
                                                            
    static def p(i_type, i_str)                             { if(isEnabled && i_type >= s_type) pp i_str      }
    static def diag(i_str)                                  { CefLogDev.p(CefLogDev.Type.diag,         i_str) }
    static def info(i_str)                                  { CefLogDev.p(CefLogDev.Type.info,         i_str) }
    static def exception(i_str)                             { CefLogDev.p(CefLogDev.Type.exception,    i_str) }
    static def warn(i_str)                                  { CefLogDev.p(CefLogDev.Type.warn,         i_str) }
    static def error(i_str)                                 { CefLogDev.p(CefLogDev.Type.error,        i_str) }
}
