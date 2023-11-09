package chat

import (
	"github.com/go-ini/ini"
	"log"
	"os"
	"path/filepath"
)

const (
	confName    = ".chat.ini"     // The name of the program's configuration file.
	contextName = ".chat.context" // Name of the file that stores the program context.
)

var (
	cfg            *ini.File
	DefaultSetting = &Default{}
)

type Default struct {
	BaseURL         string
	AuthToken       string
	Model           string
	PS1             string
	UseProxy        bool
	ProxyUrl        string
	UseTranslator   bool
	YouDaoAppKey    string
	YouDaoAppSecret string
	NativeLanguage  string
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}

func InitFromConf() {
	var err error
	chatConfig := filepath.Join(homeDir(), confName)
	cfg, err = ini.Load(chatConfig)
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse '%s': %v", chatConfig, err)
	}
	mapTo("default", DefaultSetting)
}

// mapTo map section
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
