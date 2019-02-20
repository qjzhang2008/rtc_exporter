package main

import (
	"flag"
	"fmt"
	"os"
	"rtc_exporter/common"
	"rtc_exporter/structure"
	"rtc_exporter/utils"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
)

var (
	confFile = flag.String("c", "./config/config.json", "configuration file,json format")

	rtcData = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name:      "rtc_server",
		Namespace: "rtc_server",
		Help:      "rtc server report running manually~",
	}, []string{"server", "key"})
)

//读取程序配置文件
func ReadConfigFile() (string, string, string, float64) {

	common.ProcessOptions()
	if err := common.LoadConfigFromFile(*confFile); err != nil {
		fmt.Println("Load Config File fail,", err)
		return "", "", "", -1
	}
	common.DumpConfigContent()

	// 获取日志配置
	expo_dest, err := common.GetConfigByKey("exporter.dest")
	if err != nil {
		fmt.Println("can not get dest config:", err)
		return "", "", "", -1
	}
	jobname, err := common.GetConfigByKey("exporter.jobname")
	if err != nil {
		fmt.Println("can not get jobname config:", err)
		return "", "", "", -1
	}
	filename, err := common.GetConfigByKey("exporter.filename")
	if err != nil {
		fmt.Println("can not get filename config:", err)
		return "", "", "", -1
	}
	interval, err := common.GetConfigByKey("exporter.interval")
	if err != nil {
		fmt.Println("can not get interval config:", err)
		return "", "", "", -1
	}

	return expo_dest.(string), jobname.(string), filename.(string), interval.(float64)
}

//main函数
func main() {

	var f func()
	var t *time.Timer

	expo_dest, jobname, filename, interval := ReadConfigFile()

	f = func() {
		var info structure.BasicInfo
		info, err := utils.ReadLineJson(filename, utils.ProcessLine)
		if err != nil {
			fmt.Println("Info error!")
		} else {
			registry := prometheus.NewRegistry()
			registry.MustRegister(rtcData)
			pusher := push.New(expo_dest, jobname).Gatherer(registry)

			rtcData.WithLabelValues(info.Server, "user").Set((float64)(info.Data.User))
			rtcData.WithLabelValues(info.Server, "qps").Set((float64)(info.Data.Qps))

			if err := pusher.Add(); err != nil {
				fmt.Println("Could not push to Pushgateway:", err)
				os.Exit(-1)
			}

			fmt.Println("success")
		}
		t = time.AfterFunc(time.Duration(interval)*time.Second, f)
	}

	t = time.AfterFunc(time.Duration(interval)*time.Second, f)

	defer t.Stop()

	//simulate doing stuff
	time.Sleep(time.Minute)

}
