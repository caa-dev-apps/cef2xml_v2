package cefpass

import groovy.util.CliBuilder
import org.apache.commons.cli.Option
import java.io.File

///////////////////////////////////////////////////////////////////////////////
//

public class CmdLnArgs_v2
{
    static CmdLnArgs_v2 s_object = new CmdLnArgs_v2()
    private CmdLnArgs_v2() {}
    static CmdLnArgs_v2 getObject() { return s_object}

    def m_args
    def m_cli
    
    def isCommentsOn = false
    
    def filePath 
    def searchFolders
    def xmlSchemas
    def logsFolder

    def isOutputHeaderXML
    def isQuickValidation
    
    def testRuleId
    def isNoStopOnFail
    def outputResultsLevel
    
    def configFilePath
    def configEnv
    
    def filename
    def logicalFileId
    def cefFileVersion

    def isDebugInfo4CmdLnArgs
    
    def isOk = false;
    
    // ConfigSlurper returns [:] when missing from config file
    def slurper_val(i_val) {
        i_val != [:] ? i_val : false
    }
    
    
    def init(String[] i_args) {
        m_args = i_args
        m_cli = new CliBuilder(usage:'java -jar cefpass-0.1.0.jar -f <path-to-cef> -i <search include dirs> -x <xml-schema files>')
                
        m_cli.h(longOpt: 'help',                                                                                                       'usage information')
        m_cli.c(longOpt: 'comments-on',                                                                                                '(Optional) include comments in the output data(XML). Default=False')
        
        m_cli.f(longOpt: 'cef',                        args: 1,                            required: true,                             '(Required) path to cef file')
        m_cli.i(longOpt: 'include',                    args: Option.UNLIMITED_VALUES,      required: false,    valueSeparator: ',',    '(Optional) list of include folders to search for ceh files')
        m_cli.l(longOpt: 'logs',                       args: 1,                            required: false,                            '(Optional) path to logs folder')
        m_cli.x(longOpt: 'xsd',                        args: Option.UNLIMITED_VALUES,      required: false,    valueSeparator: ',',    '(Optional) list of xml schema files to validate header data against')
        
        m_cli.u(longOpt: 'xo',                                                                                                         '(Optional) output header meta data in xml format. Default=False')
        m_cli.q(longOpt: 'qv',                                                                                                         '(Optional) quick validation (only checks 1st data row). Default=False ')
              
        m_cli.r(longOpt: 'rule-id',                    args: 1,                            required: false,                            '(Optional) Test specific Rule Id. Default=All e.g. "1.02"')
        m_cli.s(longOpt: 'stop',                                                           required: false,                            '(Optional) No-Stop on Fail. Default=False')
        m_cli.o(longOpt: 'output',                     args: 1,                            required: false,                            '(Optional) Output Results Level(0,1,2,3): 0:Debug 1:Info 2:Stages 3:Result Only')
              
        m_cli.g(longOpt: 'config',                     args: 1,                            required: false,                            '(Optional) A groovy dsl config file')
        m_cli.e(longOpt: 'config-env',                 args: 1,                            required: false,                            '(Optional) A config environment tag (e.g. dev, test, default=none)')

        m_cli.z(longOpt: 'debug-info',                                                                                                 '(Optional) Outputs the cmdline args data and exits immediately. For debug Only!')
        
        m_cli.width = 120
        //x m_cli.usage()

        
        //or ['-h', '--help'].intersect(args?.toList())
        if('-h' in i_args || '--help' in i_args) {
            m_cli.usage() 
            System.exit(-1)
        }
        
        def options = m_cli.parse(i_args)

        // values are false if missing i.e. user does not set them
        isCommentsOn              = options.c                         // -c presence = true, abscence = false
        
        filePath                  = options.f
        
        searchFolders             = options.is ?: []
        xmlSchemas                = options.xs ?: []
        logsFolder                = options.l  

        isOutputHeaderXML         = options.u
        isQuickValidation         = options.q
        
        testRuleId                = options.r
        isNoStopOnFail            = options.s
        outputResultsLevel        = options.o
        
        configFilePath            = options.g
        configEnv                 = options.e

        isDebugInfo4CmdLnArgs     = options.z                           // used to dump internal state of CmdLnArgs_v2 via show() method below.
        
        if(configFilePath)
        {
            def l_env = configEnv ?: ""
            def l_config = new ConfigSlurper(l_env).parse(new File(configFilePath).toURL())                
                                                                                            // slurper changes [:] into false
            isCommentsOn          =  isCommentsOn                                   ?: Boolean.valueOf(slurper_val(l_config.settings.isCommentsOn))
            
            searchFolders         += (slurper_val(l_config.settings.searchFolders)  ?: [])
            xmlSchemas            += (slurper_val(l_config.settings.xmlSchemas)     ?: [])
            logsFolder            = logsFolder                                      ?: slurper_val(l_config.settings.logsFolder)
            
            isOutputHeaderXML     = isOutputHeaderXML                               ?: slurper_val(l_config.settings.outputHeaderXML)
            isQuickValidation     = isQuickValidation                               ?: slurper_val(l_config.settings.quickValidation)
            
            if(testRuleId == false) {
                if(l_config.settings.testRuleId instanceof java.lang.Number) {
                    testRuleId = l_config.settings.testRuleId.toString()
                }
            }

            isNoStopOnFail          = isNoStopOnFail                                ?: slurper_val(l_config.settings.noStopOnFail)
            
            if(outputResultsLevel == false) {
                if(l_config.settings.outputResultsLevel instanceof java.lang.Integer) {
                    outputResultsLevel = l_config.settings.outputResultsLevel
                }
            }
        }
        
        outputResultsLevel        = Utils.getIntegerInRange(outputResultsLevel, 0, 3, 1)
        CefLogResults.setOutputType(outputResultsLevel)
        
        filename              = Utils.getCefFilename(filePath)
        logicalFileId         = Utils.getCefLogicalFileId(filename)
        cefFileVersion        = Utils.getCefFileVersion(filename)
        
        isOk = true
        
        if(isDebugInfo4CmdLnArgs) {
            CefLogDev.setIsEnabled(true);
            
            println "-z Debug Info CmdLnArgs"
            try{
                showAll()
                CefLogDev.diag("-z : Exiting now..")
            }
            catch(Exception e) {
                e.printStackTrace()
            }
            finally {
                System.exit(-1) // used for debug only!
            }
        }
    }
    
    ///////////////////////////////////////////////////////////////////////////////
    //
    
    def showLine(i_key, i_value) {
        CefLogDev.info i_key.padRight(50) + i_value            
    }

    def show() {
        showLine "isCommentsOn",           getIsCommentsOn()

        showLine "filePath",               getFilePath()
        showLine "searchFolders",          getSearchFolders()
        showLine "xmlSchemas",             getXmlSchemas()
        showLine "logsFolder",             getLogsFolder()

        showLine "isOutputHeaderXML",      getIsOutputHeaderXML()
        showLine "isQuickValidation",      getIsQuickValidation()

        showLine "testRuleId",             getTestRuleId()
        showLine "isNoStopOnFail",         getIsNoStopOnFail()
        showLine "outputResultsLevel",     getOutputResultsLevel()

        showLine "filename",               getFilename()
        showLine "logicalFileId",          getLogicalFileId()
        showLine "cefFileVersion",         getCefFileVersion()

        showLine "isDebugInfo4CmdLnArgs",  getIsDebugInfo4CmdLnArgs()
        
        showLine "isOk",                   getIsOk()
    }
    
    def showAll() {
        CefLogDev.diag "Command Line Args:"
        m_args.each {
            CefLogDev.diag "\t" + it
        }
        
        CefLogDev.diag("\n")
        
        m_cli.usage()
        
        CefLogDev.diag("\n")
        
        show()
    }
    
}
