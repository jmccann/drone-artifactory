package main

import (
	"encoding/json"
	"os"

	"github.com/jmccann/drone-artifactory/artifactory"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var revision string // build number set at compile-time

func main() {
	app := cli.NewApp()
	app.Name = "artifactory plugin"
	app.Usage = "artifactory plugin"
	app.Action = run
	app.Version = revision
	app.Flags = []cli.Flag{

		//
		// plugin args
		//

		cli.StringFlag{
			Name:   "actions",
			Usage:  "Actions to perform against artifactory",
			EnvVar: "PLUGIN_ACTIONS",
		},
		cli.StringFlag{
			Name:   "apikey",
			Usage:  "Artifactory API Key",
			EnvVar: "ARTIFACTORY_APIKEY, PLUGIN_APIKEY",
		},
		cli.BoolFlag{
			Name:   "debug",
			Usage:  "Enable debug logging",
			EnvVar: "PLUGIN_DEBUG",
		},
		cli.BoolFlag{
			Name:   "dry-run",
			Usage:  "Enable dry-run mode",
			EnvVar: "PLUGIN_DRY_RUN",
		},
		cli.StringFlag{
			Name:   "password",
			Usage:  "Artifactory server password",
			EnvVar: "ARTIFACTORY_PASSWORD,PLUGIN_PASSWORD",
		},
		cli.StringFlag{
			Name:   "url",
			Usage:  "Artifactory server URL",
			EnvVar: "PLUGIN_URL",
		},
		cli.StringFlag{
			Name:   "username",
			Usage:  "Artifactory server username",
			EnvVar: "ARTIFACTORY_USERNAME,PLUGIN_USERNAME",
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		logrus.Fatal(err)
	}
}

func run(c *cli.Context) error {
	logrus.WithFields(logrus.Fields{
		"Revision": revision,
	}).Info("Artifactory Drone Plugin Version")

	if c.Bool("debug") {
		logrus.SetLevel(logrus.DebugLevel)
	}

	actions, err := unmarshalActions(c.String("actions"))

	if err != nil {
		return err
	}

	plugin := Plugin{
		Actions: actions,
		Config: artifactory.Config{
			APIKey:   c.String("apikey"),
			Debug:    c.Bool("debug"),
			DryRun:   c.Bool("dry-run"),
			Password: c.String("password"),
			URL:      c.String("url"),
			Username: c.String("username"),
		},
	}

	err = plugin.Exec()

	return err
}

func unmarshalActions(rawJSON string) ([]Action, error) {
	logrus.WithFields(logrus.Fields{
		"raw-actions": rawJSON,
	}).Debug()

	bytes := []byte(rawJSON)
	var actions []Action
	err := json.Unmarshal(bytes, &actions)

	if err != nil {
		return nil, err
	}

	var rawActions []json.RawMessage
	err = json.Unmarshal(bytes, &rawActions)

	if err != nil {
		return nil, err
	}

	for idx := range actions {
		actions[idx].RawArguments = rawActions[idx]
	}

	logrus.WithFields(logrus.Fields{
		"actions": actions,
	}).Debug()

	return actions, nil
}
