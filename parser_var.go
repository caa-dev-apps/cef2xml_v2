package main

import (
)

func (hds *CefHeaderData) kv_var(kv *KeyVal) (err error) {

	evar, err := getVar(&kv.key)
	if err != nil {
	}

    val_str := val_string(kv)
    
	switch evar {
	//-  case PARAMETER_ID:        see below
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

func (hds *CefHeaderData) stx_var(name *string) (err error) {
	hds.m_var = Parameter{}
	hds.m_var.PARAMETER_ID = *name

	return
}

func (hds *CefHeaderData) etx_var() (err error) {
	hds.m_data.DATASETS.PARAMETERS.PARAMETER = append(hds.m_data.DATASETS.PARAMETERS.PARAMETER, hds.m_var)

	return
}
