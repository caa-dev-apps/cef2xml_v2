package cefpass

import javax.xml.parsers.DocumentBuilder;
import javax.xml.parsers.DocumentBuilderFactory;
 
import org.w3c.dom.Attr;
import org.w3c.dom.Document;
import org.w3c.dom.Element;
 
import groovy.xml.XmlUtil

///////////////////////////////////////////////////////////////////////////////
//

public class CefHeaderXml{
    DocumentBuilderFactory docFactory = DocumentBuilderFactory.newInstance();
    DocumentBuilder docBuilder = docFactory.newDocumentBuilder();
    Document m_doc = docBuilder.newDocument();
    
    Element m_root = m_doc.createElement("root");
    def m_cur = m_root
    
    def error(i_message) { CefLogDev.error i_message; System.exit(-1) }
    def appendDocument(i_from_document) { 
        def l_next = i_from_document.m_root.getFirstChild()
        
        while(l_next != null)
        {
            m_cur.appendChild(m_doc.importNode(l_next, true))
            l_next = l_next.getNextSibling()
        }
    }

    def stxMix(t, n) {  
        m_cur = m_doc.createElement(t); 
        m_cur.setAttribute("name", n)
        m_root.appendChild(m_cur)
    }

    def stxMeta(n)              { stxMix("meta", n) }
    def etxMeta(n)              { m_cur = m_root }
    def stxVar(n)               { stxMix("var", n) } 
    def etxVar(n)               { m_cur = m_root }
    def addAttr(k,v)            { m_cur.appendChild(m_doc.createElement(k)).appendChild(m_doc.createTextNode(v)) }
    def addComment(c)           { m_cur.appendChild(m_doc.createComment(c)) }
    
    def getDocumentElement()    { return m_root }
    def getXmlNodesAsString()   { XmlUtil.serialize(m_root) }
    def dumpX()                 { CefLogDev.diag getXmlNodesAsString() }
    
    def getHeaderXPath()        { return new CefHeaderXPath( getXmlNodesAsString() ) }
    //x !!! ********* - this should work and be faster too - than above(reparsing) - need a break from looking at it ******** !!!
    //x def getHeaderXPath()    { return new CefHeaderXPath( m_root ) }
    
    def removeQuotes(v) {
        def l = v.length()
        def q = (l > 2 && v[0] == '"' && v[l-1] == '"') 
        if (q==true) { v = v.substring(1,l-1) }
        
        [v, q]
    }

    def add_kv(k, v){
        Show.showHeaderKV(k,v)
        
        //x if(CefParser.test_kv(k, v) == CefParser.Error.eNULL)
            
        // throws Exception on fatal error
        CefParser.test_kv(k, v)
    
        if("START_META".compareToIgnoreCase(k) == 0)			stxMeta(v) 
        else if ("END_META".compareToIgnoreCase(k) == 0)		etxMeta(v) 
        else if ("START_VARIABLE".compareToIgnoreCase(k) == 0) 	stxVar(v) 
        else if ("END_VARIABLE".compareToIgnoreCase(k) == 0)	etxVar(v) 
        else 													addAttr(k,v) 
    }
}    
