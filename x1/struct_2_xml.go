//x <?xml version="1.0" encoding="UTF-8"?>
//x <!-- Generated from:
//x /c/services/header_includes//CL_CH_MISSION.ceh
//x /c/services/header_includes/edi//CL_CH_EDI_EXP.ceh
//x /c/services/header_includes//C3_CH_OBS.ceh
//x /c/services/header_includes/edi//C3_CH_EDI3_INST.ceh
//x /c/services/header_includes/edi//C3_CH_EDI_EGD_DATASET.ceh
//x /c/catalog22/products/edi/2011//C3_CP_EDI_EGD__20111009_V01.cef.gz
//x by CEF2XML            CAAtools: 1.3.5-Multi support
//x Cluster Metadata Dictionary(s): 2.0 2.2-pre 2.3  -->
//x <CAA_METADATA>
//x     <MISSION_METADATA>
//x         <MISSION>
//x         <MISSION_TIME_SPAN>
//x         <MISSION_AGENCY>
//x         <MISSION_DESCRIPTION>
//x         <MISSION_KEY_PERSONNEL>
//x         <MISSION_REFERENCES>
//x         <MISSION_REGION>[]...
//x         <MISSION_CAVEATS>
//x         <OBSERVATORIES>
//x             <OBSERVATORY_METADATA>
//x             </OBSERVATORY_METADATA>
//x         </OBSERVATORIES>
//x         <EXPERIMENTS>
//x             <EXPERIMENT_METADATA>
//x             </EXPERIMENT_METADATA>
//x         </EXPERIMENTS>
//x         <DATASETS>
//x             <DATASET_METADATA>
//x                 <DATASET_ID>
//x                 <DATA_TYPE>
//x                 <DATASET_TITLE>
//x                 <DATASET_DESCRIPTION>
//x                 <CONTACT_COORDINATES>
//x                 <PROCESSING_LEVEL>
//x                 <TIME_RESOLUTION>
//x                 <MIN_TIME_RESOLUTION>
//x                 <MAX_TIME_RESOLUTION>
//x                 <DATASET_CAVEATS>
//x                 <ACKNOWLEDGEMENT>
//x                 <PARAMETERS>[]...
//x                 </PARAMETERS>
//x                 <FILE>
//x                     <LOGICAL_FILE_ID>
//x                     <VERSION_NUMBER>
//x                     <FILE_TIME_SPAN>
//x                     <GENERATION_DATE>
//x                     <FILE_CAVEATS>
//x                     <METADATA_TYPE>
//x                     <METADATA_VERSION>
//x                     <FILE_TYPE>
//x                     <DATASET_VERSION>
//x                 </FILE>
//x             </DATASET_METADATA>
//x         </DATASETS>
//x     </MISSION_METADATA>
//x </CAA_METADATA>

//x import (
//x 	"io/ioutil"
//x 	"encoding/xml"
//x 	"fmt"
//x )

package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

func test_write_xml_01() {
    
	type Address struct {
		City, State string
	}
	type Person struct {
		XMLName   xml.Name `xml:"person"`
		Id        int      `xml:"id,attr"`
		FirstName string   `xml:"name>first"`
		LastName  string   `xml:"name>last"`
		Age       int      `xml:"age"`
		Height    float32  `xml:"height,omitempty"`
		Married   bool
		Address
		Comment string `xml:",comment"`
	}

	v := &Person{Id: 13, FirstName: "John", LastName: "Doe", Age: 42}
	v.Comment = " Need more details. "
	v.Address = Address{"Hanga Roa", "Easter Island"}

	output, err := xml.MarshalIndent(v, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	os.Stdout.Write(output)
        
}        
        



//x type File struct {
//x     XMLName                 xml.Name `xml:"FILE"`
//x     
//x     Logical_File_Id         string   `xml:"LOGICAL_FILE_ID"`
//x     VERSION_NUMBER          string   `xml:"VERSION_NUMBER"`
//x     FILE_TIME_SPAN          string   `xml:"FILE_TIME_SPAN"`
//x     GENERATION_DATE         string   `xml:"GENERATION_DATE"`
//x     FILE_CAVEATS            string   `xml:"FILE_CAVEATS"`
//x     METADATA_TYPE           string   `xml:"METADATA_TYPE"`
//x     METADATA_VERSION        string   `xml:"METADATA_VERSION"`
//x     FILE_TYPE               string   `xml:"FILE_TYPE"`
//x     DATASET_VERSION         string   `xml:"DATASET_VERSION"`
//x }


///////////////////////////////////////////////////////////////////////////////
//

type File struct {
    XMLName                 xml.Name `xml:"FILE"`
    
    LOGICAL_FILE_ID         string   // `xml:"LOGICAL_FILE_ID"`
    VERSION_NUMBER          string   
    FILE_TIME_SPAN          string   
    GENERATION_DATE         string   
    FILE_CAVEATS            string   
    METADATA_TYPE           string   
    METADATA_VERSION        string   
    FILE_TYPE               string   
    DATASET_VERSION         string   
}


func test_write_xml_02() {
    
//x 	v := &Person{Id: 13, FirstName: "John", LastName: "Doe", Age: 42}
//x 	v.Comment = " Need more details. "
//x 	v.Address = Address{"Hanga Roa", "Easter Island"}
//x 
    
    v := &File{
        LOGICAL_FILE_ID:    "file_id22",
        VERSION_NUMBER:     "0.1",
        FILE_TIME_SPAN:     "2010-2015",
        GENERATION_DATE:    "2010-01-02",
        FILE_CAVEATS:       "Caveat emptor",
        METADATA_TYPE:      "this is a type",
        METADATA_VERSION:   "1.234",
        FILE_TYPE:          "v01",
        DATASET_VERSION:    "3.2",
    }
    
	output, err := xml.MarshalIndent(v, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	os.Stdout.Write(output)
        
} 

///////////////////////////////////////////////////////////////////////////////
//

//x <CAA_METADATA>
//x     <MISSION_METADATA>
//x         <MISSION>
//x         <MISSION_TIME_SPAN>
//x         <MISSION_AGENCY>
//x         <MISSION_DESCRIPTION>
//x         <MISSION_KEY_PERSONNEL>
//x         <MISSION_REFERENCES>
//x         <MISSION_REGION>[]...
//x         <MISSION_CAVEATS>
//x         <OBSERVATORIES>
//x             <OBSERVATORY_METADATA>
//x             </OBSERVATORY_METADATA>
//x         </OBSERVATORIES>
//x         <EXPERIMENTS>
//x             <EXPERIMENT_METADATA>
//x             </EXPERIMENT_METADATA>
//x         </EXPERIMENTS>
//x         <DATASETS>
//x             <DATASET_METADATA>
//x                 <DATASET_ID>
//x                 <DATA_TYPE>
//x                 <DATASET_TITLE>
//x                 <DATASET_DESCRIPTION>
//x                 <CONTACT_COORDINATES>
//x                 <PROCESSING_LEVEL>
//x                 <TIME_RESOLUTION>
//x                 <MIN_TIME_RESOLUTION>
//x                 <MAX_TIME_RESOLUTION>
//x                 <DATASET_CAVEATS>
//x                 <ACKNOWLEDGEMENT>
//x                 <PARAMETERS>[]...
//x                 </PARAMETERS>
//x                 <FILE>
//x                     <LOGICAL_FILE_ID>
//x                     <VERSION_NUMBER>
//x                     <FILE_TIME_SPAN>
//x                     <GENERATION_DATE>
//x                     <FILE_CAVEATS>
//x                     <METADATA_TYPE>
//x                     <METADATA_VERSION>
//x                     <FILE_TYPE>
//x                     <DATASET_VERSION>
//x                 </FILE>
//x             </DATASET_METADATA>
//x         </DATASETS>
//x     </MISSION_METADATA>
//x </CAA_METADATA>


//x         <OBSERVATORIES>
//x             <OBSERVATORY_METADATA>
//x                 <OBSERVATORY>MULTIPLE</OBSERVATORY>
//x                 <OBSERVATORY>Cluster-1</OBSERVATORY>
//x                 <OBSERVATORY>Cluster-2</OBSERVATORY>
//x                 <OBSERVATORY>Cluster-3</OBSERVATORY>
//x                 <OBSERVATORY>Cluster-4</OBSERVATORY>
//x                 <OBSERVATORY_DESCRIPTION>All Cluster spacecraft</OBSERVATORY_DESCRIPTION>
//x                 <OBSERVATORY_TIME_SPAN>2000-01-01T00:00:00.0Z/2030-12-31T23:59:59.0Z</OBSERVATORY_TIME_SPAN>
//x                 <OBSERVATORY_REGION>Solar_Wind</OBSERVATORY_REGION>
//x                 <OBSERVATORY_REGION>Bow_Shock</OBSERVATORY_REGION>
//x                 <OBSERVATORY_REGION>Magnetosheath</OBSERVATORY_REGION>
//x                 <OBSERVATORY_REGION>Magnetopause</OBSERVATORY_REGION>
//x                 <OBSERVATORY_REGION>Magnetosphere</OBSERVATORY_REGION>
//x                 <OBSERVATORY_REGION>Magnetotail</OBSERVATORY_REGION>
//x                 <OBSERVATORY_REGION>Polar_Cap</OBSERVATORY_REGION>
//x                 <OBSERVATORY_REGION>Auroral_Region</OBSERVATORY_REGION>
//x                 <OBSERVATORY_REGION>Cusp</OBSERVATORY_REGION>
//x                 <OBSERVATORY_REGION>Radiation_Belt</OBSERVATORY_REGION>
//x                 <OBSERVATORY_REGION>Plasmasphere</OBSERVATORY_REGION>
//x             </OBSERVATORY_METADATA>
//x         </OBSERVATORIES>
//x         <EXPERIMENTS>


type Observatories struct {
    OBSERVATORY                 []string            `xml:"Observatories>OBSERVATORY_METADATA>OBSERVATORY"`
    OBSERVATORY_DESCRIPTION     string              `xml:"Observatories>OBSERVATORY_METADATA>OBSERVATORY_DESCRIPTION"`
    OBSERVATORY_TIME_SPAN       string              `xml:"Observatories>OBSERVATORY_METADATA>OBSERVATORY_TIME_SPAN"`
    OBSERVATORY_REGION          []string            `xml:"Observatories>OBSERVATORY_METADATA>OBSERVATORY_REGION"`
}

type Datasets struct {
//x     XMLName                     xml.Name `xml:"DATASETS"`
}

type Experiments struct {
//x     XMLName                     xml.Name `xml:"EXPERIMENTS"`
}

type CAA_MetaData struct {
    XMLName                 xml.Name `xml:"CAA_METADATA"`
    
    MISSION                     string              `xml:"MISSION_METADATA>MISSION"`
    MISSION_TIME_SPAN           string              `xml:"MISSION_METADATA>MISSION_TIME_SPAN"`
    MISSION_AGENCY              string              `xml:"MISSION_METADATA>MISSION_AGENCY"`
    MISSION_DESCRIPTION         string              `xml:"MISSION_METADATA>MISSION_DESCRIPTION"`
    MISSION_KEY_PERSONNEL       string              `xml:"MISSION_METADATA>MISSION_KEY_PERSONNEL"`
    MISSION_REFERENCES          string              `xml:"MISSION_METADATA>MISSION_REFERENCES"`
    MISSION_REGION              []string            `xml:"MISSION_METADATA>MISSION_REGION"`
    MISSION_CAVEATS             string              `xml:"MISSION_METADATA>MISSION_CAVEATS"`
    Observatories               
    Experiments                 
    Datasets                    
}

//x <CAA_METADATA>
//x     <MISSION_METADATA>
//x         <MISSION>
//x         <MISSION_TIME_SPAN>
//x         <MISSION_AGENCY>
//x         <MISSION_DESCRIPTION>
//x         <MISSION_KEY_PERSONNEL>
//x         <MISSION_REFERENCES>
//x         <MISSION_REGION>[]...
//x         <MISSION_CAVEATS>
//x         <OBSERVATORIES>
//x             <OBSERVATORY_METADATA>
//x             </OBSERVATORY_METADATA>
//x         </OBSERVATORIES>
//x         <EXPERIMENTS>
//x             <EXPERIMENT_METADATA>
//x             </EXPERIMENT_METADATA>
//x         </EXPERIMENTS>
//x         <DATASETS>
//x             <DATASET_METADATA>

func test_write_xml_03() {
    
    m := &CAA_MetaData {
        MISSION:                "string",
        MISSION_TIME_SPAN:      "string",
        MISSION_AGENCY:         "string",
        MISSION_DESCRIPTION:    "string",
        MISSION_KEY_PERSONNEL:  "string",
        MISSION_REFERENCES:     "string",
        MISSION_REGION:         []string{"r1", "r2", "r3"},
        MISSION_CAVEATS:        "string",
//x     Observatories
//x     Experiments                 
//x     Datasets                    
    }
        
	output, err := xml.MarshalIndent(m, "", "  ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	os.Stdout.Write(output)
} 


//x type Observatories struct {
//x     OBSERVATORY                 []string            `xml:"OBSERVATORY_METADATA>OBSERVATORY"`
//x     OBSERVATORY_DESCRIPTION     string              `xml:"OBSERVATORY_METADATA>OBSERVATORY_DESCRIPTION"`
//x     OBSERVATORY_TIME_SPAN       string              `xml:"OBSERVATORY_METADATA>OBSERVATORY_TIME_SPAN"`
//x     OBSERVATORY_REGION          []string            `xml:"OBSERVATORY_METADATA>OBSERVATORY_REGION"`
//x }



func test_write_xml_04() {
    
    m := &CAA_MetaData {}
    
    m.MISSION_REGION = append(m.MISSION_REGION, "MISSION_REGION #1")    
    m.MISSION_REGION = append(m.MISSION_REGION, "MISSION_REGION #2")    
    m.MISSION_REGION = append(m.MISSION_REGION, "MISSION_REGION #3")    
    
    m.MISSION_CAVEATS           = "MISSION_CAVEATSMISSION_CAVEATS"

    m.Observatories.OBSERVATORY = append(m.Observatories.OBSERVATORY, "OBSERVATORIES #1")    
    m.Observatories.OBSERVATORY = append(m.Observatories.OBSERVATORY, "OBSERVATORIES #2")    
    m.Observatories.OBSERVATORY = append(m.Observatories.OBSERVATORY, "OBSERVATORIES #3")    
    
    m.Observatories.OBSERVATORY_DESCRIPTION = "observatory_description"
    m.Observatories.OBSERVATORY_TIME_SPAN  = "observatory_time_span"
    
    m.Observatories.OBSERVATORY_REGION     = append(m.Observatories.OBSERVATORY_REGION, "observatory_region #1")    
    m.Observatories.OBSERVATORY_REGION     = append(m.Observatories.OBSERVATORY_REGION, "observatory_region #2")    
    m.Observatories.OBSERVATORY_REGION     = append(m.Observatories.OBSERVATORY_REGION, "observatory_region #3")    
    
    
	output, err := xml.MarshalIndent(m, "", "  ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	os.Stdout.Write(output)
} 

        
func main() {
    //x test_write_xml_01() 
    //x test_write_xml_02() 
    
    fmt.Println("\n\n")
    test_write_xml_03() 
    
    fmt.Println("\n\n")
    test_write_xml_04() 
}
