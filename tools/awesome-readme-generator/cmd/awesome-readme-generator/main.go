package main

import (
	"bytes"
	"flag"
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/ThorstenHans/awesome-tech-conferences/tools/awesome-readme-generator/pkg/confs"
	helpers "github.com/ThorstenHans/awesome-tech-conferences/tools/awesome-readme-generator/pkg/template"
	"gopkg.in/yaml.v2"
)

var (
	version            = "dev"
	templateFile       string
	dataFile           string
	readmeTemplateFile string
	readmeFile         string
	token              string
)

func init() {
	flag.StringVar(&dataFile, "data", "", "Path to the file containing all data that should be processed")
	flag.StringVar(&templateFile, "template", "", "Go template that is used")
	flag.StringVar(&readmeTemplateFile, "readme-template", "", "Path to the README template file")
	flag.StringVar(&readmeFile, "output", "", "Where should the final README be written to")
	flag.StringVar(&token, "token", "$AWESOME$", "The token that should be replaced in the README template")

}

func main() {
	v := flag.Bool("version", false, "Prints the current version")
	flag.Parse()
	if *v {
		log.Printf("awesome-readme-builder version %s", version)
		os.Exit(0)
	}

	tFuncs := template.FuncMap{
		"nice":   helpers.FormatDate,
		"length": helpers.GetMapLength,
		"hasCfp": helpers.HasCallForPaper,
		"keys":   helpers.GetKeys,
		"values": helpers.GetValues,
	}

	f, err := os.ReadFile(dataFile)
	if err != nil {
		log.Fatal("Could not read data file", err)
	}

	t, err := os.ReadFile(readmeTemplateFile)
	if err != nil {
		log.Fatal("Could not read template file", err)
	}

	templateLayout, err := os.ReadFile(templateFile)
	if err != nil {
		log.Fatal("Could not read layout file", err)
	}

	// unmarshal file into new instance interface{}
	var data confs.ConferenceList
	err = yaml.Unmarshal(f, &data)
	if err != nil {
		log.Fatal("Could not process data file. Please ensure that an array is provided as root element", err)
	}

	textTemplate, err := template.New("README").Funcs(tFuncs).Parse(string(templateLayout))
	if err != nil {
		log.Fatal("Could not parse Go template layout", err)
	}
	o, err := os.Create(readmeFile)

	if err != nil {
		log.Fatal("Could not create output file", err)
	}
	defer o.Close()

	model := confs.NewReadmeModel(data)

	var r bytes.Buffer
	err = textTemplate.Execute(&r, model)
	if err != nil {
		log.Fatal("Could not render template", err)
	}

	readmeTemplate := string(t)
	readmeTemplate = strings.Replace(readmeTemplate, token, r.String(), -1)
	o.WriteString(readmeTemplate)

}
