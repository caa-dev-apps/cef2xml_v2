package cefpass

import rules2015.RuleSets
import rules2015.RS0_CefParser

///////////////////////////////////////////////////////////////////////////////
//

public class App {

    CefHeaderXml m_headerXml = null
    String[] m_args
    
    public App()
    {
    }
    
    ///////////////////////////////////////////////////////////////////////////////
    //
    
    class AppResult extends CefResult
    {
        def stage_results = [:]
        
        AppResult() { 
            tag="AppResult" 
        }
        
        def pushResult(i_key, i_result) {
            stage_results[i_key] = i_result
            //x i_result.info()
            i_result.stage_result()
            i_result.isError
        }
    }
    
    ///////////////////////////////////////////////////////////////////////////////
    //
    
    boolean args(String[] i_args) {
        true
    }
    
    def stage_1__cmdln_args = {
        def l_CefLogDev_diag_str = "\nStage 1: "
        
        def l_result = new CefResult("stage_1__cmdln_args")
        
        try{
            CmdLnArgs_v2.getObject().init(m_args)
            if(CmdLnArgs_v2.getObject().getIsOk() == false) { throw new Exception("Error: CmdLnArgs_v2") } 
            CefLogDev.diag l_CefLogDev_diag_str
            
            FileLogs.init()
            
            l_result.setPass()
        }
        catch (Exception e) {
            CefLogDev.diag l_CefLogDev_diag_str
            
            l_result.setException(e)
        }
            
        l_result
    }
    
    def stage_2__parser = {
        CefLogDev.diag "\nStage 2: "
        
        def l_result = new CefResult("stage_2__parser")
        
        try{
            CefReader l_reader = new CefReader()
            //x CefHeaderXml 
            m_headerXml = l_reader.getHeaderXml()
            FileLogs.writeTextFile("nodes.xml", m_headerXml.getXmlNodesAsString())

            RS0_CefParser.showAll()
            
            l_result.setPass()
        }
        catch (Exception e) {
           println e
           l_result.setException(e)
        }
            
        l_result
    }
    
    def stage_3__xsd_schema = {
        CefLogDev.diag "\nStage 3: "
        
        def l_result = new CefResult("stage_3__xsd_schema")
        
        try{
            //  String l_xmlPath = FileLogs.getFilePath("nodes.xml")
            //  String l_xsdPath = "C:/work.dev/2014.09.27.github.cef.pass.v2/cef-pass/xsd/a1.xsd"
            //  if(CefHeaderXsd.validateXMLSchema(l_xsdPath, l_xmlPath) == false) return
            CefLogResults.stage3_info("validateXMLSchema", "Skipped")
            
//x         l_result.setPass()
            l_result.setSkipped()
        }
        catch (Exception e) {
           l_result.setException(e)
        }
            
        l_result
    }
    
    def stage_4__rules = {
        CefLogDev.diag "\nStage 4: "
        
        def l_result = new CefResult("stage_4__rules")
        
        try{
        
            l_result = RuleSets.run(m_headerXml.getHeaderXPath())
            l_result.tag = "stage_4__rules"
        }
        catch (Exception e) {
            
           l_result.setException(e)
        }
            
        l_result
    }

    def stage_5__data = {
        CefLogDev.diag "\nStage 5: "
        
        def l_result = new CefResult("stage_5__data")
        
        try{
            l_result.setPass()
        }
        catch (Exception e) {
           l_result.setException(e)
        }
            
        l_result
    }
    
    ///////////////////////////////////////////////////////////////////////////////
    //
    
    public def stages(String[] i_args) {
        m_args = i_args
        
        def l_result = new AppResult()
        
        def l_stages = [
            STAGE_1: stage_1__cmdln_args,
            STAGE_2: stage_2__parser,
            STAGE_3: stage_3__xsd_schema ,
            STAGE_4: stage_4__rules,
            STAGE_5: stage_5__data
        ] as TreeMap // retain order

        def r = l_stages.any { it ->
            //x CefLog.info it.key
            def res = it.value()
            l_result.pushResult(it.key, res)
        }

        if(r == false) l_result.setPass()   // did not stop
        else           l_result.setFail()
    
        l_result        
    }    
    
    public static void main(String[] i_args) {
        //x CefLogDev.diag "Hello, World!"
        //x CefLogDev.diag "workingDir: "  + System.getProperty("user.dir")
        
        App a = new App();

        try
        {
            def l_result = a.stages(i_args)
            l_result.full_result()
        }
        catch(Exception e) {
            e.printStackTrace()        
        }
        
        FileLogs.close()
    }
}

