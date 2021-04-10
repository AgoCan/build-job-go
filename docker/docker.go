package docker

import (
	"build-job-go/config"
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"io"
	"os"

	"fmt"
	"log"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/archive"
)

type Options struct {
	ImageName string
	Workspace string
	Client    *client.Client
	Context   context.Context
}

func NewOptions() *Options {
	var ImageName string

	if config.IsOnlyCode == "1" {
		ImageName = fmt.Sprintf("%s/%s/%s-%s-%s-%s:1.0.%s",
			config.HarborRegistry,
			config.EagleRegister,
			config.PreImageName,
			config.UserID,
			config.Project,
			config.CodeDir,
			config.Count,
		)
	} else {
		ImageName = fmt.Sprintf("%s/%s/%s-%s-%s-%s-code:1.0.%s",
			config.HarborRegistry,
			config.EagleRegister,
			config.PreImageName,
			config.UserID,
			config.Project,
			config.CodeDir,
			config.Count,
		)
	}
	workspace := fmt.Sprintf("%s/%s/%s/%s",
		config.Workspace,
		config.Username,
		config.Project,
		config.CodeDir,
	)
	client, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Fatal(err)
	}
	return &Options{
		ImageName: ImageName,
		Workspace: workspace,
		Client:    client,
		Context:   context.Background(),
	}
}

func (o *Options) Build() (err error) {

	tar, err := archive.TarWithOptions(o.Workspace, &archive.TarOptions{})
	if err != nil {
		log.Fatal(err, " :unable to open Dockerfile")
	}

	res, err := o.Client.ImageBuild(
		o.Context,
		tar,
		types.ImageBuildOptions{
			Dockerfile: "Dockerfile",
			Remove:     true,
			Tags:       []string{o.ImageName},
		})
	if err != nil {
		log.Fatal(err, " :unable to build docker image")
	}
	defer res.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(res.Body)
	newStr := buf.String()
	newStr = strings.TrimSpace(newStr)
	newSplice := strings.Split(newStr, "\r\n")
	buildLog := make(map[string]interface{}, 4)

	for _, val := range newSplice {
		json.Unmarshal([]byte(val), &buildLog)
		fmt.Printf("%v", buildLog["stream"])

		if _, ok := buildLog["errorDetail"]; ok {
			fmt.Printf("%v\n", buildLog["errorDetail"])
		}
		if _, ok := buildLog["error"]; ok {
			fmt.Printf("%v\n", buildLog["error"])
		}
	}

	return nil
}

func (o *Options) Push() {

	authConfig := types.AuthConfig{
		Username: config.HarborUsername,
		Password: config.HarborPassword,
	}
	encodedJSON, err := json.Marshal(authConfig)
	if err != nil {
		log.Fatal(err, " :unable to build docker image")
	}
	authStr := base64.URLEncoding.EncodeToString(encodedJSON)
	res, err := o.Client.ImagePush(o.Context, o.ImageName, types.ImagePushOptions{RegistryAuth: authStr})
	io.Copy(os.Stdout, res)
	if err != nil {
		log.Fatal(err, " :unable to build docker image")
	}

}

func (o *Options) Remove() {
	o.Client.ImageRemove(o.Context, o.ImageName, types.ImageRemoveOptions{Force: true, PruneChildren: true})
}

// 指令解析
func (o *Options) ParseInstruction() {

}

// 参数解析
func (o *Options) ParseParam() {

}
