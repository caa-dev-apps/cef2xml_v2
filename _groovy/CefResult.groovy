package cefpass

import exceptions.CefException

///////////////////////////////////////////////////////////////////////////////
//

class CefResult
{
    def tag        = "x"
    def isError     = true
    def status      = "error"
    def lastError   = null      // Generic - for Exceptions and Failing Inner Result
    def exception   = null      // Generic - for Exceptions and Failing Inner Result

    CefResult(){
    }
    
    CefResult(i_tag){
        tag = i_tag
    }

    def setSkipped()                    { isError=false; status="Skipped";  this }
    
    def setOk(i_status="Ok")            { isError=false; status=i_status;  this }
    def setError(i_status="Error")      { isError=true;  status=i_status;  this }
    
    def setPass(i_status="Pass")        { isError=false; status=i_status;  this }
    def setFail(i_status="Fail")        { isError=true;  status=i_status;  this }
    
    def setLastError(i_lastError)       { setFail(); lastError = i_lastError; this }
    def setException(i_exception)       { setFail(); exception = i_exception; this }

    def info()                          { CefLogDev.info "TAG: ${tag}  isError: ${isError}  status:${status}  lastError:${lastError}  exception:${exception}" }

    def stage_result()                  { CefLogResults.stage_result "${tag} : ${status} " }
    def full_result()                   { CefLogResults.full_result  "${tag} : ${status} " }
}
