package cefpass

import groovy.xml.XmlUtil

public class Show {

    public static boolean s_show_cef_filename = true;               // main cef file
    public static boolean s_show_search_folders = true;             // list of folders to search for ceh includes

    public static boolean s_show_file_read_line_details = false     // file line as read + line number etc
    public static boolean s_show_header_kv = false                  // header key and value
    public static boolean s_show_header_comment = false             // header comment
    
    public static boolean s_show_cx_files = true;                   // list the main cef and ceh filenames
    public static boolean s_show_nodes = true;                      // show nodebuilder -> nodes dump
    public static boolean s_show_xml_nodes = true;                  // show nodebuilder -> xml dump
    public static boolean s_show_data_line = true;                  // show data as read
    
    ///////////////////////////////////////////////////////////////////////////////
    //

    public static void showCefFilePath(i_path) {
        if(s_show_cef_filename == true) {
            CefLogResults.stage_info  1, "CEF Filename:"    
            CefLogResults.stage_info(1, i_path, Utils.fileExistsS(i_path))
        }
    }
    
    public static void showSearchFolders(i_searchFolders) {
        if(s_show_search_folders == true) {
            CefLogResults.stage_info 1,"Search Folders:"    
            i_searchFolders.each { 
                CefLogResults.stage_info(1, it, Utils.directoryExistsS(it))
            }
        }
    }

    public static void showContextFile(String i_text) 
    {
        if(s_show_cx_files == true) CefLogDev.info i_text
    }
    
    public static void showNodes(i_root) 
    {
        if(s_show_nodes == true) {
            CefLogDev.info "\nNodes:"
            
            def writer = new StringWriter()
            i_root.print(new PrintWriter(writer))
            CefLogDev.info writer.toString()    
        }
    }
    
    public static void showXmlNodes(i_root) 
    {
        if(s_show_xml_nodes == true) {
            CefLogDev.info "\nXML Nodes:"
        
            CefLogDev.info new XmlUtil().serialize(i_root)    
        }
    }
    
    public static void showDataLine(i_line) 
    {
        //x if(s_show_data_line == true) CefLogDev.info "Data: " + i_line
    }
    
    public static void showContexts(i_cefContexts) {
        if(s_show_search_folders == true) {
            CefLogDev.info "\nCef Contexts:"    
            i_cefContexts.each { 
                CefLogDev.info "\t" + it.getString()
            }
        }
    }
 
    public static void showFileReadLineDetails(i_line_details) {
        if(s_show_file_read_line_details == true) CefLogDev.info "H:\t" + i_line_details
    }
    
    public static void showHeaderKV(i_key, i_value) {
        if(s_show_header_kv == true) CefLogDev.info "H:\t" + "k:" + i_key + "v:" + i_value
    }
    
    public static void showHeaderComment(i_comment) {
        if(s_show_header_comment == true) CefLogDev.info "H:\t" + i_comment
    }
 }
