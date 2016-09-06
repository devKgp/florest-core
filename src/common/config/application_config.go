package config

import (
	"github.com/jabong/florest-core/src/common/monitor"
	"github.com/jabong/florest-core/src/common/utils/http"
	"github.com/jabong/florest-core/src/components/cache"
	"github.com/jabong/florest-core/src/core/common/cachestrategy"
)

type AppConfig struct {
	AppName             string
	AppVersion          string
	ServerPort          string
	LogConfFile         string
	MonitorConfig       monitor.MonitorConf
	CacheStrategyConfig cachestrategy.Config
	Performance         PerformanceConfigs
	DynamicConfig       DynamicConfigInfo
	HttpConfig          http.Config
	Profiler            ProfilerConfig
	ResponseHeaders     ResponseHeaderFields
	ApplicationConfig   interface{}
}

type PerformanceConfigs struct {
	UseCorePercentage float64
	GCPercentage      float64
}

type ResponseHeaderFields struct {
	CacheControl CacheControlHeaders
}

type CacheControlHeaders struct {
	ResponseType    string
	NoCache         bool
	NoStore         bool
	MaxAgeInSeconds int
}

type Application struct {
	ResponseHeaders ResponseHeaderFields
}

type DynamicConfigInfo struct {
	Active          bool
	RefreshInterval int
	ConfigKey       string
	Cache           cache.Config
}

type ProfilerConfig struct {
	Enable       bool
	SamplingRate float64
}

//Global ApplicationConfig Singleton
var GlobalAppConfig = new(AppConfig)

var ApplicationConfig interface{}