package tools

import "time"

// SecretId 监控密钥
const SecretID = "****"
const SecretKey = "****"

// TEST

// ServerAddr 日志汇上报地址
const ServerAddr = "log-report-sz.cloud.net:11001"

// Topic zhiyan-log接入信息
const Topic = "sdk-1ca831a4c2b3b66d"

// Limit 1次分页查询数量
const Limit = 100

// DefaultOffset 分页查询的起始偏移量
const DefaultOffset = 0

const Goroutines = 1

type InstanceInfo struct {
	InstanceId   string            `json:"InstanceId"`
	InstanceName string            `json:"InstanceName"`
	DbEngine     string            `json:"DbEngine"`
	Region       string            `json:"Region"`
	Tags         map[string]string `json:"tags"`
	Metric       map[string]string `json:"metric"`
	Time         time.Time         `json:"time"`
}
