package config

import (
	"fmt"
	"os"

	props "github.com/magiconair/properties"
)

const envDetectorKey = "configEnvironment"

// GetProps takes in resourceDir which points the location of config files and loads them according to the env.
func GetProps(resourceDir string) *props.Properties {
	p, err := load(resourceDir)
	if err != nil {
		fmt.Printf("configuration loading err - %v. \n Empty Properties returned \n", err)
		p = props.NewProperties()
	}
	return p
}

func load(resourceDir string) (p *props.Properties, err error) {
	filesToLoad := []string{resourceDir + "/app.properties"}
	if v, found := os.LookupEnv(envDetectorKey); found {
		filesToLoad = append(filesToLoad, resourceDir+"/app-"+v+".properties")
	}
	p, err = props.LoadFiles(filesToLoad, props.UTF8, false)
	return
}
