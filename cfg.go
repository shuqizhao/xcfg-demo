package main

import (
	"encoding/xml"
	"xcfg"
)

type RabbitMqConfig struct {
	XMLName      xml.Name      `xml:"amqp-settings"`
	Servers      []ServerGroup `xml:"servers"`
	MajorVersion int           `xml:"majorVersion,attr"`
	MinorVersion int           `xml:"minorVersion,attr"`
}
type ServerGroup struct {
	Server []ServerItem `xml:"server"`
}
type ServerItem struct {
	Name string `xml:"name,attr"`
	Uri  string `xml:"uri,attr"`
}

func NewRabbitMqConfig() RabbitMqConfig {
	rcfg := RabbitMqConfig{}
	xcfg.LoadCfg(&rcfg)

	return rcfg
}

type RecruitWechatNotice struct {
	XMLName              xml.Name                   `xml:"RecruitWechatNoticeConfiguration"`
	SenderConfigurations []SenderConfigurationGroup `xml:"SenderConfigurations"`
	MajorVersion         int                        `xml:"majorVersion,attr"`
	MinorVersion         int                        `xml:"minorVersion,attr"`
}
type SenderConfigurationGroup struct {
	SenderConfiguration []SenderConfiguration `xml:"SenderConfiguration"`
}

type SenderConfiguration struct {
	GroupName string `xml:"GroupName"`
}

func NewRecruitWechatNotice() RecruitWechatNotice {
	cfg := RecruitWechatNotice{}
	xcfg.LoadCfg(&cfg)
	return cfg
}
