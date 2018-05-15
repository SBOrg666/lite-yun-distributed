package utils

import "github.com/tidwall/gjson"

var (
	Upload_data   []uint64
	Download_data []uint64
	InitUpload    uint64
	InitDownload  uint64
	Current_Month int
	ServersMap    map[string]gjson.Result
	ServersString string
)
