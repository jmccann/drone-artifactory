package main

import (
	"os"
	"strings"

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

		cli.BoolTFlag{
			Name:   "dry-run",
			Usage:  "if the plugin should pretend to upload but don't actually upload",
			EnvVar: "PLUGIN_DRY_RUN",
		},
		cli.BoolTFlag{
			Name:   "flat",
			Usage:  "artifacts are uploaded to the exact target path specified and their hierarchy in the source file system is ignored",
			EnvVar: "PLUGIN_FLAT",
		},
		cli.BoolFlag{
			Name:   "include-dirs",
			Usage:  "source path applies to bottom-chain directories and not only to files. Bottom-chain directories are either empty or do not include other directories that match the source path",
			EnvVar: "PLUGIN_INCLUDE_DIRS",
		},
		cli.StringFlag{
			Name:     "path",
			Usage:    "Where to upload artifacts to",
			EnvVar:   "PLUGIN_PATH",
			FilePath: ".artifactory_path",
		},
		cli.StringFlag{
			Name:   "password",
			Usage:  "Artifactory server password",
			EnvVar: "ARTIFACTORY_PASSWORD,PLUGIN_PASSWORD",
		},
		cli.StringFlag{
			Name:   "apikey",
			Usage:  "Artifactory API Key",
			EnvVar: "ARTIFACTORY_APIKEY, PLUGIN_APIKEY",
		},
		cli.BoolTFlag{
			Name:   "recursive",
			Usage:  "artifacts are also collected from sub-folders of the source directory for upload.",
			EnvVar: "PLUGIN_RECURSIVE",
		},
		cli.BoolFlag{
			Name:   "regexp",
			Usage:  "command will interpret the sources as a regular expression.",
			EnvVar: "PLUGIN_REGEXP",
		},
		cli.StringSliceFlag{
			Name: "sources",
			Usage: "local file system path to artifacts which should be uploaded to " +
				"Artifactory. You can specify multiple artifacts by using " +
				"wildcards or a regular expression as designated by the --regexp command option.",
			EnvVar: "PLUGIN_SOURCES",
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

	plugin := Plugin{
		Config: Config{
			DryRun:      c.Bool("dry-run"),
			Flat:        c.Bool("flat"),
			IncludeDirs: c.Bool("include-dirs"),
			Path:        strings.TrimSpace(c.String("path")),
			APIKey:      c.String("apikey"),
			Password:    c.String("password"),
			Recursive:   c.Bool("recursive"),
			Regexp:      c.Bool("regexp"),
			Sources:     c.StringSlice("sources"),
			URL:         c.String("url"),
			Username:    c.String("username"),
		},
	}

	err := plugin.Exec()

	return err
}
