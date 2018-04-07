package main

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
)

func main() {
	exit := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}

	app := cli.NewApp()
	app.Author = "marvin + konsorten GmbH"
	app.Version = versionString
	app.HideHelp = true
	app.Usage = `Consul Versioning Tool

		This tool updates a Consul configuration based on a git repository
		by handling the committed files.

		The following file structure is supported:

		<git repository root>
		  +- rootkey.subkey   (key directory)
		     +- index.yml     (properties file)
		     +- example.json  (value file)
		     +- example2.jpg  (value file)
		
		Every directory is interpreted as a folder in Consul. The directory
		name is allowed to contain dots to match several subfolders.

		Every file in a directory is interpreted as key-value where
		the key is the filename and the value is the contents of the file.
		The filename is not allowed to contain dots.

		The properties file is a special file named "index.yml" and is
		parsed to contain several properties at once. It has the following
		format:

			properties:
			  mykey1: "simple value"
			  mykey2.subkey: "folder with value"
			  mykey3: |
			    multi-line value
			    (including all newlines)
			  mykey4: >
			    multi-line value
			    (removing all newlines)
		`

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "directory, C",
			Value: ".",
			Usage: "The Git repository directory to use. Using the current as default.",
		},
		cli.BoolFlag{
			Name:  "debug, d",
			Usage: "Enable debug output.",
		},
	}

	app.Action = func(c *cli.Context) {
		exit(run(c))
	}

	exit(app.Run(os.Args))
}

func run(c *cli.Context) error {
	// enable debug output
	initTerminal()

	if c.Bool("debug") {
		log.SetLevel(log.DebugLevel)
	}

	// run in remote mode?
	remote := c.String("repository")

	if remote != "" {
		log.Debugf("Git repository URL: %v", remote)

		return runRemoteRepository(remote)
	}

	// run in directory mode?
	dir := c.String("directory")

	log.Debugf("Git repository directory: %v", dir)

	// start handling the directory
	return runLocalRepository(dir)
}
