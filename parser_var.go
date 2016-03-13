package main

import "fmt"

//x  	"strings"
//x 	"fmt"
//x     "errors"

///////////////////////////////////////////////////////////////////////////////
//
//
//  type CefHeaderData struct {
//      m_state State
//      m_name string
//      m_data CAA_MetaData
//      m_meta eMeta
//      m_var Parameter
//  }

//old   func (hds *CefHeaderData) kv_var(k, v *string)  (err error) {
//old
//old       //x println("x x x x x x x x x x  x x x x x x x x x x x ", *k)
//old
//old       evar, err := getVar(k)
//old       if(err != nil) {
//old   //x         return err
//old       }
//old
//old       switch evar {
//old   //x     case PARAMETER_ID:        see below
//old           case PARAMETER_TYPE:      hds.m_var.PARAMETER_TYPE     = *v
//old           case CATDESC:             hds.m_var.CATDESC            = *v
//old           case UNITS:               hds.m_var.UNITS              = *v
//old           case SI_CONVERSION:       hds.m_var.SI_CONVERSION      = *v
//old           case SIZES:               hds.m_var.SIZES              = *v
//old           case VALUE_TYPE:          hds.m_var.VALUE_TYPE         = *v
//old           case SIGNIFICANT_DIGITS:  hds.m_var.SIGNIFICANT_DIGITS = *v
//old           case FILLVAL:             hds.m_var.FILLVAL            = *v
//old           case FIELDNAM:            hds.m_var.FIELDNAM           = *v
//old           case LABLAXIS:            hds.m_var.LABLAXIS           = *v
//old           case DELTA_PLUS:          hds.m_var.DELTA_PLUS         = *v
//old           case DELTA_MINUS:         hds.m_var.DELTA_MINUS        = *v
//old           case ENTITY:              hds.m_var.ENTITY             = *v
//old           case PROPERTY:            hds.m_var.PROPERTY           = *v
//old           case QUALITY:             hds.m_var.QUALITY            = *v
//old           case DEPEND_0:            hds.m_var.DEPEND_0           = *v
//old
//old           case FRAME:               hds.m_var.FRAME              =  *v
//old           case TENSOR_ORDER:        hds.m_var.TENSOR_ORDER       =  *v
//old           case COORDINATE_SYSTEM:   hds.m_var.COORDINATE_SYSTEM  =  *v
//old           case FRAME_VELOCITY:      hds.m_var.FRAME_VELOCITY     =  *v
//old           case REPRESENTATION_1:    hds.m_var.REPRESENTATION_1   =  *v
//old           case LABEL_1:             hds.m_var.LABEL_1            =  *v
//old
//old           default:
//old               // nothing will get here since getVar filters
//old               println("====================================================================== TODO", *k)
//old       }
//old
//old       return
//old   }
//old
//old   func (hds *CefHeaderData) stx_var(name *string)  (err error) {
//old
//old       hds.m_var = Parameter{}
//old       hds.m_var.PARAMETER_ID = *name
//old
//old       return
//old   }
//old
//old   func (hds *CefHeaderData) etx_var()  (err error) {
//old
//old       hds.m_data.DATASETS.PARAMETERS.PARAMETER = append(hds.m_data.DATASETS.PARAMETERS.PARAMETER, hds.m_var)
//old
//old       return
//old   }

func (hds *CefHeaderData) kv_var(kv *KeyVal) (err error) {

	evar, err := getVar(&kv.key)
	if err != nil {
		//x         return err
	}

	if len(kv.val) > 1 {
		fmt.Println("------------------ TODO -> kv_var Array len > 1 ------------------ ", kv)
	}

	switch evar {
	//x  case PARAMETER_ID:        see below
	case PARAMETER_TYPE:           hds.m_var.PARAMETER_TYPE = kv.val[0]
	case CATDESC:                  hds.m_var.CATDESC = kv.val[0]
	case UNITS:                    hds.m_var.UNITS = kv.val[0]
	case SI_CONVERSION:            hds.m_var.SI_CONVERSION = kv.val[0]
	case SIZES:                    hds.m_var.SIZES = kv.val[0]
	case VALUE_TYPE:               hds.m_var.VALUE_TYPE = kv.val[0]
	case SIGNIFICANT_DIGITS:       hds.m_var.SIGNIFICANT_DIGITS = kv.val[0]
	case FILLVAL:                  hds.m_var.FILLVAL = kv.val[0]
	case FIELDNAM:                 hds.m_var.FIELDNAM = kv.val[0]
	case LABLAXIS:                 hds.m_var.LABLAXIS = kv.val[0]
	case DELTA_PLUS:               hds.m_var.DELTA_PLUS = kv.val[0]
	case DELTA_MINUS:              hds.m_var.DELTA_MINUS = kv.val[0]
	case ENTITY:                   hds.m_var.ENTITY = kv.val[0]
	case PROPERTY:                 hds.m_var.PROPERTY = kv.val[0]
	case QUALITY:                  hds.m_var.QUALITY = kv.val[0]
	case DEPEND_0:                 hds.m_var.DEPEND_0 = kv.val[0]

	case FRAME:                    hds.m_var.FRAME = kv.val[0]
	case TENSOR_ORDER:             hds.m_var.TENSOR_ORDER = kv.val[0]
	case COORDINATE_SYSTEM:        hds.m_var.COORDINATE_SYSTEM = kv.val[0]
	case FRAME_VELOCITY:           hds.m_var.FRAME_VELOCITY = kv.val[0]
	case REPRESENTATION_1:         hds.m_var.REPRESENTATION_1 = kv.val[0]
	case LABEL_1:                  hds.m_var.LABEL_1 = kv.val[0]

	default:
		// nothing will get here since getVar filters
		println("====================================================================== TODO", kv.key[0])
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
