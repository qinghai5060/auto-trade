package main

import (
	"flag"
	"github.com/golang/glog"
	"os"
	"io/ioutil"

	"encoding/json"
)


const  ConfFile = "config.json"


type KeyInfo struct {
	Key string `json:"key"`
	Secret string `json:"secret"`
}


func main() {
	flag.Parse()

	defer glog.Flush()
	// Method call
	glog.Infof("Start to get pairs")
	keyInfo := loadConfigFromFile(ConfFile)
	if keyInfo == nil {
		glog.Fatal("Failed get config file.")
		os.Exit(0)
	}
	gateApi := &GateApi{Key:keyInfo.Key, Secret: keyInfo.Secret}
	// all pairs
	var ret  = gateApi.getPairs()
	glog.Infof("pairs getted is %s", ret)

}


func loadConfigFromFile(f string) *KeyInfo{
	if _, err := os.Stat(f); err != nil {
		glog.Errorf("Failed to load config from file :%s",f)
		return nil
	}

	fileByte,err := ioutil.ReadFile(f)
	if err != nil {
		glog.Errorf("Failed to read file :%s",f)
		return nil
	}

	keyInfo := &KeyInfo{}
	err = json.Unmarshal(fileByte, keyInfo)
	if err != nil {
		glog.Errorf("Failed to unmarshal file (%s), error: %s",f,err.Error())
		return nil
	}

	return keyInfo

}