package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/urfave/cli"
)

type app struct {
	Name    string `json:"NAME" xml:"name,attr"`
	Path    string `json:"PATH" xml:"path,attr"`
	Type    string `json:"TYPE" xml:"type,attr"`
	UUID    string `json:"UUID" xml:"uuid"`
	Version string `json:"VERSION" xml:"version"`
}

type application struct {
	Release struct {
		PkgName  string `json:"PKG_NAME"`
		AppID    string `json:"APPID"`
		QuayRepo string `json:"QUAY_REPO"`
		Channel  string `json:"CHANNEL"`
	} `json:"RELEASE"`

	Manifest struct {
		Apps []app `json:"APPS"`
	} `json:"MANIFEST"`
}

type manifest struct {
	Apps []app `xml:"app"`
}

func convertToXML(jsonFile, xmlFile string) error {
	// Read the json file
	data, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		return err
	}

	jsonData := application{}
	if err = json.Unmarshal(data, &jsonData); err != nil {
		return err
	}

	// Now process the json data and convert to XML
	xmlData := manifest{}
	for _, appl := range jsonData.Manifest.Apps {
		if appl.Path == `` {
			appl.Path = fmt.Sprintf("%s/%s.sh", appl.Name, appl.Name)
		}

		if appl.Version == `` {
			appl.Version = fmt.Sprintf("@%s_version@", appl.Name)
		}

		xmlData.Apps = append(xmlData.Apps, appl)
	}

	xmlStr, err := xml.MarshalIndent(xmlData, ``, `  `)
	if err != nil {
		return err
	}

	if xmlFile != `` {
		if !filepath.IsAbs(xmlFile) {
			return fmt.Errorf("invalid out file: %s", xmlFile)
		}
		dir := filepath.Dir(xmlFile)
		_, err := os.Stat(dir)
		if os.IsNotExist(err) {
			return fmt.Errorf("parent dir does not exist: %s", dir)
		}

		// Write to output file
		xmlBytes := []byte(xml.Header + string(xmlStr))
		if err := ioutil.WriteFile(xmlFile, xmlBytes, 0644); err != nil {
			return err
		}
	} else {
		fmt.Println(xml.Header + string(xmlStr))
	}
	return nil
}

func main() {
	// Maintains the arguments for the app
	type cliState struct {
		JsonFile string
		XmlFile  string
	}

	// Create a state instance
	state := cliState{}

	app := cli.NewApp()
	app.Name = "jsonToXml"
	app.Usage = "A simple custom json to xml converter"
	app.Authors = []cli.Author{
		{
			Name:  "Sanshit Sharma",
			Email: "present87@gmail.com",
		},
	}

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "in",
			Usage:       "Absolute path to the file containing the json data",
			Required:    true,
			Destination: &state.JsonFile,
		},
		cli.StringFlag{
			Name:        "out",
			Usage:       "Output XML file. Provide absoulte path. Default: stdout",
			Destination: &state.XmlFile,
		},
	}

	app.Action = func(c *cli.Context) error {
		return convertToXML(state.JsonFile, state.XmlFile)
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Printf("failed to convert. err: '%v'\n", err)
		os.Exit(1)
	}
}
