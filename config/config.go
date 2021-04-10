package config

import (
	"flag"
	"fmt"
	"log"
	"testing"

	"github.com/spf13/viper"
)

// Config 配置结构体
var (
	GitUsername string
	GitEmail    string
	GitPassword string
	GitOrigin   string
	GitURL      string

	HarborRegistry string
	HarborIp       string
	HarborUsername string
	HarborPassword string
	EagleRegister  string
	PreImageName   string

	Username    string
	UserID      string
	CallbackURL string
	CallbackID  string
	Project     string
	Count       string
	CodeDir     string
	LaunchesNum string
	IsPush      string
	Remainder   string
	IsOnlyCode  string

	DownloadDomain       string
	DownloadIp           string
	BaseFrameImageName   string
	DockerfileLogPath    string
	DockerfileLogUrlPath string
	CommandLogPath       string
	Workspace            string
)

// Init init config
func init() {
	var configFile = flag.String("c", "../config/config.yaml", "config fime path.")
	testing.Init()
	fmt.Println(configFile)
	flag.Parse()
	config := viper.New()
	config.AutomaticEnv()
	config.SetConfigFile(*configFile)
	err := config.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	GitUsername = config.GetString("GIT_USERNAME")
	GitEmail = config.GetString("GIT_EMAIL")
	GitPassword = config.GetString("GIT_PASSWORD")
	GitOrigin = config.GetString("GIT_ORIGIN")
	GitURL = config.GetString("GIT_URL")

	HarborRegistry = config.GetString("HARBOR_REGISTRY")
	HarborIp = config.GetString("HARBOR_IP")
	HarborUsername = config.GetString("HARBOR_USERNAME")
	HarborPassword = config.GetString("HARBOR_PASSWORD")
	EagleRegister = config.GetString("EAGLE_REGISTER")
	PreImageName = config.GetString("PRE_IMAGE_NAME")

	Username = config.GetString("USERNAME")
	UserID = config.GetString("USER_ID")
	CallbackURL = config.GetString("CALLBACK_URL")
	CallbackID = config.GetString("CALLBACK_ID")
	Project = config.GetString("PROJECT")
	Count = config.GetString("COUNT")
	CodeDir = config.GetString("CODE_DIR")
	LaunchesNum = config.GetString("LAUNCHES_NUM")
	IsPush = config.GetString("IS_PUSH")
	Remainder = config.GetString("REMAINDER")
	IsOnlyCode = config.GetString("IS_ONLY_CODE")

	DownloadDomain = config.GetString("DOWNLOAD_DOMAIN")
	DownloadIp = config.GetString("DOWNLOAD_IP")
	BaseFrameImageName = config.GetString("BASE_FRAME_IMAGE_NAME")
	DockerfileLogPath = config.GetString("DOCKERFILE_LOG_PATH")
	DockerfileLogUrlPath = config.GetString("DOCKERFILE_LOG_URL_PATH")
	CommandLogPath = config.GetString("COMMAND_LOG_PATH")
	Workspace = config.GetString("WORKSPACE")
}
