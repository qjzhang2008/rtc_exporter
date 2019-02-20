package utils

import (
	"encoding/json"
	"fmt"
	"rtc_exporter/structure"
)

//读取json
func ReadJson(line []byte) (structure.BasicInfo, error) {

	var config structure.BasicInfo
	fmt.Println(string(line))
	err := json.Unmarshal(line, &config)
	if err != nil {
		fmt.Println("unmarshal error!")
		return config, err
	}
	return config, nil
}
