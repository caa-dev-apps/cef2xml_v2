package main

import (
)


//          GLOBAL (FILE) ATTRIBUTES
//              Keyword
//              DATASET_VERSION
//              FILE_CAVEATS
//              FILE_SIZE x
//              FILE_TIME_SPAN
//              FILE_TYPE
//              GENERATION_DATE
//              INGESTION_DATE x
//              LOGICAL_FILE_ID
//              METADATA_TYPE
//              METADATA_VERSION
//              PARENT_DATASET (not required for CEF) x
//              VERSION_NUMBER


//x           <A key="FILE_NAME"                 val="&#34;C1_CP_PEA_PITCH_3DXLARH_DPFlux__20111031_V01.cef&#34;"></A>
//x           <A key="FILE_FORMAT_VERSION"               val="&#34;CEF-2.0&#34;"></A>
//x           <A key="END_OF_RECORD_MARKER"              val="&#34;$&#34;"></A>
//x           <A key="DATA_UNTIL"                val="&#34;End_of_Data&#34;"></A>
//x 
//x PARENT_DATASET
//x INGESTION_DATE
//x FILE_SIZE
//x 
//x FILE_NAME
//x FILE_FORMAT_VERSION
//x END_OF_RECORD_MARKER
//x DATA_UNTIL
//x 
//x 
//x 
//x 
//x         <FILE>
//x           <FILE_TYPE>cef</FILE_TYPE>
//x           <METADATA_TYPE>CAA</METADATA_TYPE>
//x           <METADATA_VERSION>2.0</METADATA_VERSION>
//x           <LOGICAL_FILE_ID>C1_CP_PEA_PITCH_3DXLARH_DPFlux__20111031_V01</LOGICAL_FILE_ID>
//x           <VERSION_NUMBER>01</VERSION_NUMBER>
//x           <DATASET_VERSION>Version: 3.3.5 (2011/01/31)</DATASET_VERSION>
//x           <FILE_CAVEATS>and Rob Wilson.</FILE_CAVEATS>
//x           <FILE_TIME_SPAN>2011-10-31T00:00:00.000Z/2011-10-31T23:59:59.999Z</FILE_TIME_SPAN>
//x           <GENERATION_DATE>2015-09-25T11:44:03.000Z</GENERATION_DATE>
//x         </FILE>
//x       </DATASET_METADATA>
//x       <UNEXPECTED>
//x         <ATTR>
//x           <A key="FILE_NAME" val="&#34;C1_CP_PEA_PITCH_3DXLARH_DPFlux__20111031_V01.cef&#34;"></A>
//x           <A key="FILE_FORMAT_VERSION" val="&#34;CEF-2.0&#34;"></A>
//x           <A key="END_OF_RECORD_MARKER" val="&#34;$&#34;"></A>
//x           <A key="DATA_UNTIL" val="&#34;End_of_Data&#34;"></A>
//x         </ATTR>
//x         <META></META>
//x 

//x Keyword
//x ACKNOWLEDGEMENT
//x CATDESC
//x COMPOUND
//x COMPOUND_DEF
//x CONTACT_COORDINATES
//x COORDINATE_SYSTEM
//x DATA
//x DATA_TYPE
//x DATASET_CAVEATS
//x DATASET_DESCRIPTION
//x DATASET_ID
//x DATASET_TITLE
//x DELTA_PLUS and DELTA_MINUS
//x //x DEPEND_i
//x DEPEND_0
//x DEPEND_1
//x DEPEND_2
//x DEPEND_3
//x DISPLAYTYPE
//x ENTITY
//x ERROR_PLUS and ERROR_MINUS
//x EXPERIMENT
//x EXPERIMENT_CAVEATS
//x EXPERIMENT_DESCRIPTION
//x EXPERIMENT_KEY_PERSONNEL
//x EXPERIMENT_REFERENCES
//x FIELDNAM
//x FILLVAL
//x FLUCTUATIONS
//x FRAME
//x FRAME VELOCITY
//x FRAME_ORIGIN
//x INSTRUMENT_CAVEATS
//x INSTRUMENT_DESCRIPTION
//x INSTRUMENT_NAME
//x INSTRUMENT_TYPE
//x INVESTIGATOR_COORDINATES
//x //x LABEL_i 
//x LABEL_0
//x LABEL_1
//x LABEL_2 
//x LABEL_3 
//x LABLAXIS
//x MAX_TIME_RESOLUTION
//x MEASUREMENT_TYPE
//x MIN_TIME_RESOLUTION
//x MISSION
//x MISSION_AGENCY
//x MISSION_CAVEATS
//x MISSION_DESCRIPTION
//x MISSION_KEY_PERSONNEL
//x MISSION_REFERENCES
//x MISSION_REGION
//x MISSION_TIME_SPAN
//x OBSERVATORY
//x OBSERVATORY_CAVEATS
//x OBSERVATORY_DESCRIPTION
//x OBSERVATORY_REGION
//x OBSERVATORY_TIME_SPAN
//x PARAMETER_CAVEATS
//x PARAMETER_ID
//x PARAMETER_TYPE
//x PARENT_DATASET_ID
//x PARENT_EXPERIMENT
//x PARENT_INSTRUMENT
//x PARENT_MISSION
//x PARENT_OBSERVATORY
//x PROCESSING_LEVEL
//x PROPERTY
//x QUALITY
//x REPRESENTATION
//x SCALEMAX
//x SCALEMIN
//x SCALETYP
//x SI_CONVERSION
//x SIGNIFICANT_DIGITS
//x SIZES
//x TARGET_SYSTEM
//x TENSOR_FRAME
//x TENSOR_ORDER
//x TIME_RESOLUTION
//x UNITS
//x VALUE_TYPE
//x 

func (hds *CefHeaderData) kv_var(kv *KeyVal) (err error) {

	evar, err := getVar(&kv.key)
	if err != nil {
	}

    val_str := val_string(kv)
    
	switch evar {
	//x  case PARAMETER_ID:        see below
	case PARAMETER_TYPE:           hds.m_var.PARAMETER_TYPE = val_str
	case CATDESC:                  hds.m_var.CATDESC = val_str
	case UNITS:                    hds.m_var.UNITS = val_str
	case SI_CONVERSION:            hds.m_var.SI_CONVERSION = val_str
	case SIZES:                    hds.m_var.SIZES = val_str
	case VALUE_TYPE:               hds.m_var.VALUE_TYPE = val_str
	case SIGNIFICANT_DIGITS:       hds.m_var.SIGNIFICANT_DIGITS = val_str
	case FILLVAL:                  hds.m_var.FILLVAL = val_str
	case FIELDNAM:                 hds.m_var.FIELDNAM = val_str
	case LABLAXIS:                 hds.m_var.LABLAXIS = val_str
	case DELTA_PLUS:               hds.m_var.DELTA_PLUS = val_str
	case DELTA_MINUS:              hds.m_var.DELTA_MINUS = val_str
	case ENTITY:                   hds.m_var.ENTITY = val_str
	case PROPERTY:                 hds.m_var.PROPERTY = val_str
	case QUALITY:                  hds.m_var.QUALITY = val_str
	case DEPEND_0:                 hds.m_var.DEPEND_0 = val_str
	case DEPEND_1:                 hds.m_var.DEPEND_1 = val_str
	case DEPEND_2:                 hds.m_var.DEPEND_2 = val_str
	case DEPEND_3:                 hds.m_var.DEPEND_3 = val_str

	case FRAME:                    hds.m_var.FRAME = val_str
	case TENSOR_ORDER:             hds.m_var.TENSOR_ORDER = val_str
	case COORDINATE_SYSTEM:        hds.m_var.COORDINATE_SYSTEM = val_str
	case FRAME_VELOCITY:           hds.m_var.FRAME_VELOCITY = val_str
	case REPRESENTATION_0:         hds.m_var.REPRESENTATION_0 = val_str
	case REPRESENTATION_1:         hds.m_var.REPRESENTATION_1 = val_str
	case REPRESENTATION_2:         hds.m_var.REPRESENTATION_2 = val_str
	case REPRESENTATION_3:         hds.m_var.REPRESENTATION_3 = val_str
	case LABEL_0:                  hds.m_var.LABEL_0 = val_str
	case LABEL_1:                  hds.m_var.LABEL_1 = val_str
	case LABEL_2:                  hds.m_var.LABEL_2 = val_str
	case LABEL_3:                  hds.m_var.LABEL_3 = val_str

    case SCALEMIN:                 hds.m_var.SCALEMIN = val_str                    
    case SCALEMAX:                 hds.m_var.SCALEMAX = val_str                    
    case SCALETYP:                 hds.m_var.SCALETYP = val_str                    
    case DISPLAYTYPE:              hds.m_var.DISPLAYTYPE = val_str                 
    case DATA:                     hds.m_var.DATA = val_str                        
    
	default:
		println("TODO: ", kv.key, val_str)
        hds.m_data.DATASETS.UNEXPECTED.VAR = append(hds.m_data.DATASETS.UNEXPECTED.VAR, TypeKeyValue{Key: kv.key, Val: val_str })         
	}

	return
}

//x func (hds *CefHeaderData) kv_var(kv *KeyVal) (err error) {
//x 
//x 	evar, err := getVar(&kv.key)
//x 	if err != nil {
//x 	}
//x 
//x     val_str := val_string(kv)
//x     
//x 	switch evar {
//x 
//x     case ACKNOWLEDGEMENT:                  hds.m_var.ACKNOWLEDGEMENT = val_str
//x     case CATDESC:                          hds.m_var.CATDESC = val_str
//x     case COMPOUND:                         hds.m_var.COMPOUND = val_str
//x     case COMPOUND_DEF:                     hds.m_var.COMPOUND_DEF = val_str
//x     case CONTACT_COORDINATES:              hds.m_var.CONTACT_COORDINATES = val_str
//x     case COORDINATE_SYSTEM:                hds.m_var.COORDINATE_SYSTEM = val_str
//x     case DATA:                             hds.m_var.DATA = val_str
//x     case DATA_TYPE:                        hds.m_var.DATA_TYPE = val_str
//x     case DATASET_CAVEATS:                  hds.m_var.DATASET_CAVEATS = val_str
//x     case DATASET_DESCRIPTION:              hds.m_var.DATASET_DESCRIPTION = val_str
//x     case DATASET_ID:                       hds.m_var.DATASET_ID = val_str
//x     case DATASET_TITLE:                    hds.m_var.DATASET_TITLE = val_str
//x     case DELTA_PLUS:                       hds.m_var.DELTA_PLUS = val_str 
//x     case DELTA_MINUS:                      hds.m_var.DELTA_MINUS = val_str
//x     case DEPEND_0:                         hds.m_var.DEPEND_0 = val_str
//x     case DEPEND_1:                         hds.m_var.DEPEND_1 = val_str
//x     case DEPEND_2:                         hds.m_var.DEPEND_2 = val_str
//x     case DEPEND_3:                         hds.m_var.DEPEND_3 = val_str
//x     case DISPLAYTYPE:                      hds.m_var.DISPLAYTYPE = val_str
//x     case ENTITY:                           hds.m_var.ENTITY = val_str
//x     case ERROR_PLUS:                       hds.m_var.ERROR_PLUS = val_str and ERROR_MINUS
//x     case EXPERIMENT:                       hds.m_var.EXPERIMENT = val_str
//x     case EXPERIMENT_CAVEATS:               hds.m_var.EXPERIMENT_CAVEATS = val_str
//x     case EXPERIMENT_DESCRIPTION:           hds.m_var.EXPERIMENT_DESCRIPTION = val_str
//x     case EXPERIMENT_KEY_PERSONNEL:         hds.m_var.EXPERIMENT_KEY_PERSONNEL = val_str
//x     case EXPERIMENT_REFERENCES:            hds.m_var.EXPERIMENT_REFERENCES = val_str
//x     case FIELDNAM:                         hds.m_var.FIELDNAM = val_str
//x     case FILLVAL:                          hds.m_var.FILLVAL = val_str
//x     case FLUCTUATIONS:                     hds.m_var.FLUCTUATIONS = val_str
//x     case FRAME:                            hds.m_var.FRAME = val_str
//x     case FRAME_VELOCITY:                   hds.m_var.FRAME_VELOCITY = val_str
//x     case FRAME_ORIGIN:                     hds.m_var.FRAME_ORIGIN = val_str
//x     case INSTRUMENT_CAVEATS:               hds.m_var.INSTRUMENT_CAVEATS = val_str
//x     case INSTRUMENT_DESCRIPTION:           hds.m_var.INSTRUMENT_DESCRIPTION = val_str
//x     case INSTRUMENT_NAME:                  hds.m_var.INSTRUMENT_NAME = val_str
//x     case INSTRUMENT_TYPE:                  hds.m_var.INSTRUMENT_TYPE = val_str
//x     case INVESTIGATOR_COORDINATES:         hds.m_var.INVESTIGATOR_COORDINATES = val_str
//x     case LABEL_0:                          hds.m_var.LABEL_0 = val_str
//x     case LABEL_1:                          hds.m_var.LABEL_1 = val_str
//x     case LABEL_2:                          hds.m_var.LABEL_2 = val_str 
//x     case LABEL_3:                          hds.m_var.LABEL_3 = val_str 
//x     case LABLAXIS:                         hds.m_var.LABLAXIS = val_str
//x     case MAX_TIME_RESOLUTION:              hds.m_var.MAX_TIME_RESOLUTION = val_str
//x     case MEASUREMENT_TYPE:                 hds.m_var.MEASUREMENT_TYPE = val_str
//x     case MIN_TIME_RESOLUTION:              hds.m_var.MIN_TIME_RESOLUTION = val_str
//x     case MISSION:                          hds.m_var.MISSION = val_str
//x     case MISSION_AGENCY:                   hds.m_var.MISSION_AGENCY = val_str
//x     case MISSION_CAVEATS:                  hds.m_var.MISSION_CAVEATS = val_str
//x     case MISSION_DESCRIPTION:              hds.m_var.MISSION_DESCRIPTION = val_str
//x     case MISSION_KEY_PERSONNEL:            hds.m_var.MISSION_KEY_PERSONNEL = val_str
//x     case MISSION_REFERENCES:               hds.m_var.MISSION_REFERENCES = val_str
//x     case MISSION_REGION:                   hds.m_var.MISSION_REGION = val_str
//x     case MISSION_TIME_SPAN:                hds.m_var.MISSION_TIME_SPAN = val_str
//x     case OBSERVATORY:                      hds.m_var.OBSERVATORY = val_str
//x     case OBSERVATORY_CAVEATS:              hds.m_var.OBSERVATORY_CAVEATS = val_str
//x     case OBSERVATORY_DESCRIPTION:          hds.m_var.OBSERVATORY_DESCRIPTION = val_str
//x     case OBSERVATORY_REGION:               hds.m_var.OBSERVATORY_REGION = val_str
//x     case OBSERVATORY_TIME_SPAN:            hds.m_var.OBSERVATORY_TIME_SPAN = val_str
//x     case PARAMETER_CAVEATS:                hds.m_var.PARAMETER_CAVEATS = val_str
//x     // PARAMETER_ID see belo
//x     case PARAMETER_TYPE:                   hds.m_var.PARAMETER_TYPE = val_str
//x     case PARENT_DATASET_ID:                hds.m_var.PARENT_DATASET_ID = val_str
//x     case PARENT_EXPERIMENT:                hds.m_var.PARENT_EXPERIMENT = val_str
//x     case PARENT_INSTRUMENT:                hds.m_var.PARENT_INSTRUMENT = val_str
//x     case PARENT_MISSION:                   hds.m_var.PARENT_MISSION = val_str
//x     case PARENT_OBSERVATORY:               hds.m_var.PARENT_OBSERVATORY = val_str
//x     case PROCESSING_LEVEL:                 hds.m_var.PROCESSING_LEVEL = val_str
//x     case PROPERTY:                         hds.m_var.PROPERTY = val_str
//x     case QUALITY:                          hds.m_var.QUALITY = val_str
//x     case REPRESENTATION:                   hds.m_var.REPRESENTATION = val_str
//x     case REPRESENTATION_1:                 hds.m_var.REPRESENTATION_1 = val_str
//x     case SCALEMAX:                         hds.m_var.SCALEMAX = val_str
//x     case SCALEMIN:                         hds.m_var.SCALEMIN = val_str
//x     case SCALETYP:                         hds.m_var.SCALETYP = val_str
//x     case SI_CONVERSION:                    hds.m_var.SI_CONVERSION = val_str
//x     case SIGNIFICANT_DIGITS:               hds.m_var.SIGNIFICANT_DIGITS = val_str
//x     case SIZES:                            hds.m_var.SIZES = val_str
//x     case TARGET_SYSTEM:                    hds.m_var.TARGET_SYSTEM = val_str
//x     case TENSOR_FRAME:                     hds.m_var.TENSOR_FRAME = val_str
//x     case TENSOR_ORDER:                     hds.m_var.TENSOR_ORDER = val_str
//x     case TIME_RESOLUTION:                  hds.m_var.TIME_RESOLUTION = val_str
//x     case UNITS:                            hds.m_var.UNITS = val_str
//x     case VALUE_TYPE:                       hds.m_var.VALUE_TYPE = val_str
//x     
//x 	default:
//x 		println("TODO: ", kv.key, val_str)
//x         
//x         hds.m_data.DATASETS.UNEXPECTED.VAR = append(hds.m_data.DATASETS.UNEXPECTED.VAR, TypeKeyValue{Key: kv.key, Val: val_str })         
//x 	}
//x 
//x 	return
//x }

func (hds *CefHeaderData) stx_var(name *string) (err error) {

	hds.m_var = Parameter{}
	hds.m_var.PARAMETER_ID = *name

	return
}

func (hds *CefHeaderData) etx_var() (err error) {

	hds.m_data.DATASETS.PARAMETERS.PARAMETER = append(hds.m_data.DATASETS.PARAMETERS.PARAMETER, hds.m_var)

	return
}
