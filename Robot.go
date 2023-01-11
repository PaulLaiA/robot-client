package main

import (
	"log"
	"os"
	"robot-client/robot"
	"strings"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Token  string
	Path   string
	Ignore []string
}

func main() {
	file, err := os.ReadFile(os.Args[1])

	if err != nil {
		return
	}
	config := Config{}
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		return
	}
	log.Println(config.Token)
	log.Println(config.Path)
	robot.Listen(config.Path, func(message robot.XCAutoLog) {
		for _, s := range config.Ignore {
			if strings.Contains(message.Name, s) {
				robot.PushPlusPost(message, config.Token)
				break
			}
		}
	})
	c := make(chan struct{})
	c <- struct{}{}
}
