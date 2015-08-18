/*
	MIT License
	Copyright Ming-der Wang<ming@log4analytics.com>
*/
package main

import (
	"errors"
	"io/ioutil"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/mingderwang/goslack/service"
	"gopkg.in/yaml.v2"
)

func getConfig(c *cli.Context) (service.Config, error) {
	yamlPath := c.GlobalString("config")
	config := service.Config{}

	if _, err := os.Stat(yamlPath); err != nil {
		return config, errors.New("config path not valid")
	}

	ymlData, err := ioutil.ReadFile(yamlPath)
	if err != nil {
		return config, err
	}

	err = yaml.Unmarshal([]byte(ymlData), &config)
	return config, err
}

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.TextFormatter{})

	// Output to stderr instead of stdout, could also be a file.
	log.SetOutput(os.Stderr)

	// Only log the warning severity or above.
	log.SetLevel(log.WarnLevel)
}

func main() {
	app := cli.NewApp()
	app.Name = "goslack"
	app.Usage = "start a micro service"
	app.Version = "0.0.3"

	app.Flags = []cli.Flag{
		//	cli.StringFlag{"host", "http://localhost:8080", "use sevice host", "APP_HOST"},
		cli.StringFlag{"config, c", "config/config.yml", "config file to use", "APP_CONFIG"},
	}

	app.Commands = []cli.Command{
		{
			Name:  "serve",
			Usage: "start micro service",
			Action: func(c *cli.Context) {
				//host := c.GlobalString("host")
				cfg, err := getConfig(c)
				if err != nil {
					log.WithFields(log.Fields{
						"file:": "main.go",
						"func:": "main",
						"line:": 74,
					}).Fatal("getConfig return error")
					return
				}
				svc := service.Slack_messageService{}

				if err = svc.Run(cfg); err != nil {
					log.Fatal(err)
				}

			},
		},
		{
			Name:  "migratedb",
			Usage: "Perform database migrations",
			Action: func(c *cli.Context) {
				cfg, err := getConfig(c)
				if err != nil {
					log.Fatal(err)
					return
				}

				svc := service.Slack_messageService{}

				if err = svc.Migrate(cfg); err != nil {
					log.Fatal(err)
				}
			},
		},
	}

	app.Run(os.Args)
}
