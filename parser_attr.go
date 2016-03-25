package main

import ()

func (hds *CefHeaderData) kv_attr(kv *KeyVal) (err error) {

	switch {
	case "FILE_NAME" == kv.key:
		hds.m_data.DATASETS.FILE.FILE_NAME = val_string(kv)
	case "FILE_FORMAT_VERSION" == kv.key:
		hds.m_data.DATASETS.FILE.FILE_FORMAT_VERSION = val_string(kv)
	case "END_OF_RECORD_MARKER" == kv.key:
		hds.m_data.DATASETS.FILE.END_OF_RECORD_MARKER = val_string(kv)
	case "DATA_UNTIL" == kv.key:
		hds.m_data.DATASETS.FILE.DATA_UNTIL = val_string(kv)
	default:
		mooi_log("kv_attr::", *kv)
		hds.m_data.DATASETS.UNEXPECTED.ATTR = append(hds.m_data.DATASETS.UNEXPECTED.ATTR, TypeKeyValue{Key: kv.key, Val: val_string(kv)})
	}

	return
}
