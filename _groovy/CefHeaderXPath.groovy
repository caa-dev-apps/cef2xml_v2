package cefpass

import javax.xml.xpath.*
import javax.xml.parsers.DocumentBuilderFactory

import org.w3c.dom.Document;
import org.w3c.dom.Element;

import groovy.xml.DOMBuilder

///////////////////////////////////////////////////////////////////////////////
//
 
public class CefHeaderXPath
{
    def m_xpath = null
    def m_documentElement = null
    
    public CefHeaderXPath(String i_xml)
    {
        def builder         = DocumentBuilderFactory.newInstance().newDocumentBuilder()
        def inputStream     = new ByteArrayInputStream( i_xml.bytes )
        m_documentElement   = builder.parse(inputStream).documentElement
        
        m_xpath             = XPathFactory.newInstance().newXPath()
    }

    // USE ABOVE CONSTRUCTOR - THIS ONE IS NOT WORKING....it should though! TODO = REVIST!!!
    public CefHeaderXPath(Element i_documentElement)
    {
        m_documentElement   = i_documentElement
        m_xpath             = XPathFactory.newInstance().newXPath()
    }

    ///////////////////////////////////////////////////////////////////////////////
    //
    
    def eval_boolean(String i_xpathQuery)           {   m_xpath.evaluate(i_xpathQuery, m_documentElement, XPathConstants.BOOLEAN)   }
    def eval_dom_object_model(String i_xpathQuery)  {   m_xpath.evaluate(i_xpathQuery, m_documentElement, XPathConstants.DOM_OBJECT_MODEL)   }
    def eval_node(String i_xpathQuery)              {   m_xpath.evaluate(i_xpathQuery, m_documentElement, XPathConstants.NODE)   }
    def eval_nodeset(String i_xpathQuery)           {   m_xpath.evaluate(i_xpathQuery, m_documentElement, XPathConstants.NODESET)   }
    def eval_number(String i_xpathQuery)            {   m_xpath.evaluate(i_xpathQuery, m_documentElement, XPathConstants.NUMBER)   }
    def eval_string(String i_xpathQuery)            {   m_xpath.evaluate(i_xpathQuery, m_documentElement, XPathConstants.STRING)   }
    
    def query_nodeset(String i_xpathQuery)
    {
        def ls = []
        
        eval_nodeset(i_xpathQuery).each{
            ls << it
        }
        
        return ls
    }
    
    def getSingleNode(String i_xpathQuery)
    {
        def v = null
        def ls = query_nodeset(i_xpathQuery)
        
        if(ls != null && ls.size() == 1)
        {
            v = ls[0]
        }
        
        return v
    }

    def getSingleNodeValue(String i_xpathQuery)
    {
        def v = null
        def n = getSingleNode(i_xpathQuery)
        
        if(n != null) {
            v = n.getNodeValue()
        }
        
        return v
    }

    def getElementText(String i_xpathQuery)
    {
        def txt = getSingleNodeValue(i_xpathQuery + "/text()")
        txt
    }

    ///////////////////////////////////////////////////////////////////////////////
    //

    def getMeta(i_name)
    {
        getSingleNode('/root/meta[@name="' + i_name +  '"]')
    }

    def getMetaList()
    {
        query_nodeset('/root/meta')
    }
    
    def getMetaEntry(i_name)
    {
        eval_string('/root/meta[@name="' + i_name +  '"]/ENTRY[1]')
    }

    def getMetaEntryPair(i_name)                                                                        
    {                                                                                                   //  e.g.
        [                                                                                               //  <meta name="FILE_TIME_SPAN">
            "ENTRY":      getElementText('/root/meta[@name="' + i_name +  '"]/ENTRY[1]'),               //    <VALUE_TYPE>ISO_TIME_RANGE</VALUE_TYPE>
            "VALUE_TYPE": getElementText('/root/meta[@name="' + i_name +  '"]/VALUE_TYPE[1]')           //    <ENTRY>2011-10-09T00:00:00Z/2011-10-10T00:00:00Z</ENTRY>
        ]                                                                                               //  </meta>
    }

    def getMetaPairNames()
    {
        //  get all meta elements with a VALUE_TYPE
        def rs = []
        String l_xpathQuery = "/root/meta/VALUE_TYPE/../@name"
        
        m_xpath.evaluate(l_xpathQuery, m_documentElement, XPathConstants.NODESET).each{
            def ix = 0
            def l_name = it.getTextContent()
            
            m_xpath.evaluate('/root/meta[@name="' + l_name + '"]/ENTRY', m_documentElement, XPathConstants.NODESET).each {
                ix++ 
            }
            
            if(ix == 1) { rs << l_name}
        }
        
        return rs
    }
    
    ///////////////////////////////////////////////////////////////////////////////
    //

    def getMetaEntryPairs()
    {
        def ps = []
        def ns = getMetaPairNames()
        
        ns.each{ name -> 
            ps << getMetaEntryPair(name)                                                                        
        }
        
        ps
    }
    
    ///////////////////////////////////////////////////////////////////////////////
    //

    def getFilename()
    {
        getElementText("/root/FILE_NAME")
    }

    def getFileFormatVersion()
    {
        getElementText("/root/FILE_FORMAT_VERSION")
    }

    def getEndOfRecordMarker()
    {
        getElementText("/root/END_OF_RECORD_MARKER")
    }

    def getDataUntil()
    {
        getElementText("/root/DATA_UNTIL")
    }
    
    def getLogicalFileId()                                                      //              <meta name="LOGICAL_FILE_ID">
    {                                                                           //                <ENTRY>"C3_CP_EDI_EGD__20111009_V01"</ENTRY>
        getMetaEntry("LOGICAL_FILE_ID")                                         //              </meta>
    }

    def getVersionNumber()                                                      //              <meta name="VERSION_NUMBER">
    {                                                                           //                <ENTRY>"01"</ENTRY>
        getMetaEntry("VERSION_NUMBER")                                          //              </meta>
    }

    def getFileType()                                                           //              <meta name="FILE_TYPE">
    {                                                                           //                <ENTRY>"cef"</ENTRY>
        getMetaEntry("FILE_TYPE")                                               //              </meta>
    }

    def getDatasetVersionNumber()                                               //              <meta name="DATASET_VERSION">
    {                                                                           //                <ENTRY>"2.0"</ENTRY>
        getMetaEntry("DATASET_VERSION")                                         //              </meta>
    }

//x     def getDatasetId()                                                          
//x     {                                                                           
//x         getMetaEntry("DATASET_VERSION")                                         
//x     }
    

    
    ///////////////////////////////////////////////////////////////////////////////
    //

    def hasEmptyMetaObjects()
    {
        eval_boolean('/root/meta[count(*)=0]')                                 // t/f
    }
    
    def countEmptyMetaObjects()
    {
        eval_number('count(/root/meta[count(*)=0])')                           // n >= 0
    }
    
    //i [ "ENTRY": ----. "VALUE_TYPE": ---- ]                 
    def anyMetaEntryTypeMismatches()
    {
        def result = getMetaEntryPairs().any{
            DataTypes.isEntryTypeMismatch(it) == true
        }
        
        result
    }    
    
    def datasetVersionIsValidInt()
    {
        // "\"123\""
        def vers = Utils.getUnQuotedString(getDatasetVersionNumber())

        DataTypes.isValidInt(vers)
    }

    def logicalFileIdMatchesFilename()
    {
        def l1 = Utils.getUnQuotedString(getLogicalFileId())
//x_args        def l2 = CmdLnArgs.getLogicalFileId()
        def l2 = CmdLnArgs_v2.getObject().getLogicalFileId()
        
        return (l1 != null) && (l2 != null) && l1.toUpperCase().equals(l2.toUpperCase())
    }
 
    def versionNumberMatchesFilename()
    {
        def l1 = Utils.getUnQuotedString(getVersionNumber())
//x_args        def l2 = CmdLnArgs.getCefFileVersion()
        def l2 = CmdLnArgs_v2.getObject().getCefFileVersion()
        
        return (l1 != null) && (l2 != null) && l1.toUpperCase().equals(l2.toUpperCase())
    }
 
    def versionNumberIsValidInteger()
    {
        // "\"123\""
        def vers = Utils.getUnQuotedString(getVersionNumber())
        
        DataTypes.isValidInt(vers)
    } 
 
     //x FILE_TIME_SPAN_MUST_BE_ISO_TIME_RANGE... will fail earlier test if VALUE_TYPE is present
    def fileTimeSpanIsValidISOTimeRange()
    {
        def l_value = getMetaEntry("FILE_TIME_SPAN")
        
        DataTypes.isValidISOTimeRange(l_value)
    }
 
    //x FILE_TIME_SPAN_START_TIME_MUST_BE_BEFORE_STOP_TIME
    def fileTimeSpanStartIsBeforeFileTimeSpanStop()
    {
        def l_value = getMetaEntry("FILE_TIME_SPAN")
        
        DataTypes.isValidISOTimeRangeStartStop(l_value)
    }
    
    
    
    
 
}

