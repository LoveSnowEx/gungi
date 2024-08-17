package config

import (
	"io"
	"path/filepath"

	"github.com/LoveSnowEx/gungi/tool/filetool"
	"github.com/LoveSnowEx/gungi/tool/pathtool"
	"github.com/spf13/viper"
)

type LoggerConfig interface {
	Get(name string) io.Writer
	Set(name string, writer io.Writer)
}

type loggerConfig struct {
	wrtiers map[string]io.Writer
}

type outputInfo struct {
	path   string
	append bool
}

func (c *loggerConfig) Get(name string) io.Writer {
	return c.wrtiers[name]
}

func (c *loggerConfig) Set(name string, writer io.Writer) {
	c.wrtiers[name] = writer
}

func loadLoggerConfig(conf LoggerConfig) {
	if !viper.IsSet("logger") {
		return
	}
	for name := range viper.GetStringMap("logger") {
		outputInfo := parseOutputInfo(name)
		if outputInfo == nil {
			continue
		}
		f, err := filetool.Open(outputInfo.path, outputInfo.append)
		if err != nil {
			continue
		}
		conf.Set(name, f)
	}
}

func parseOutputInfo(name string) (info *outputInfo) {
	if !viper.IsSet("logger." + name) {
		return
	}
	info = &outputInfo{}
	info.path = filepath.Join(
		pathtool.ProjectRoot(),
		viper.GetString("logger."+name+".path"),
	)
	info.append = viper.GetBool("logger." + name + ".append")
	if info.path == "" {
		info = nil
		return
	}
	return
}
