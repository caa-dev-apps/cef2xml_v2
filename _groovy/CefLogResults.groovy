package cefpass

///////////////////////////////////////////////////////////////////////////////
// '(Optional) Output Results Level(0,1,2,3): 0:Debug 1:Info 2:Stages 3:Result Only')


public class CefLogResults
{
    enum Type {
        diag(0),
        stage_info(1),
        stage_result(2),
        full_result(3),
        
        final int id
        public Type(int id) { this.id = id }
        static Type byId(int id) { values().find { it.id == id } } 
    }

    static Type s_type = Type.diag
    static def setOutputType(int id) { 
        s_type = Type.byId(id) ?: Type.diag 

        // switch dev-debug off when arg/config for output > 0
//x         if(id == 0) CefLogDev.setIsEnabled(true)
//x         else        CefLogDev.setIsEnabled(false)
    
        CefLogDev.setIsEnabled(id == 0)
    }
    
    static def p =                                          { i_str -> println i_str}
//x     static def print_ln =                                   { i_type, i_str -> 
//x                                                                 if(i_type >= s_type) CefLogResults.p i_str  
//x                                                                 else CefLogResults.p "x: " + i_str  
//x                                                             }
    static def print_ln =                                   { i_type, i_str -> 
                                                                if(i_type >= s_type) CefLogResults.p i_str  
//x                                                                 else CefLogResults.p "x: " + i_str  
                                                            }
    static def diag =                                       { i_str -> CefLogResults.print_ln(Type.diag, i_str) }

    static def strFmt(i_str1, i_str2)                       { ((i_str2 != null) ? ("    " + i_str1).padRight(100) + i_str2 : ("  " + i_str1))} 
    static def stage_info(i_stage, i_str1, i_str2=null)     { CefLogResults.print_ln(Type.stage_info,  strFmt(i_str1, i_str2)) }

    static def stage1_info(i_str1, i_str2=null)             { stage_info(1, i_str1, i_str2) }   
    static def stage2_info(i_str1, i_str2=null)             { stage_info(2, i_str1, i_str2) }   
    static def stage3_info(i_str1, i_str2=null)             { stage_info(3, i_str1, i_str2) }   
    static def stage4_info(i_str1, i_str2=null)             { stage_info(4, i_str1, i_str2) }   
    static def stage5_info(i_str1, i_str2=null)             { stage_info(5, i_str1, i_str2) }   

    static def stage_result(i_str)                          { CefLogResults.print_ln(Type.stage_result, i_str) }
    static def full_result(i_str)                           { CefLogResults.print_ln(Type.full_result, i_str) }
}