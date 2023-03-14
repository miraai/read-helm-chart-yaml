package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"

	githubactions "github.com/sethvargo/go-githubactions"
	"gopkg.in/yaml.v2"
)

const ChartFile = "Chart.yaml"

type Chart struct {
	ApiVersion   string   `yaml:"apiVersion"`
	Name         string   `yaml:"name"`
	Version      string   `yaml:"version"`
	KubeVersion  string   `yaml:"kubeVersion"`
	Description  string   `yaml:"description"`
	Type         string   `yaml:"type"`
	Keywords     []string `yaml:"keywords"`
	Home         string   `yaml:"home"`
	Sources      []string `yaml:"sources"`
	Dependencies []*Chart `yaml:"dependencies"`
	Repository   string   `yaml:"repository"`
	Icon         string   `yaml:"icon"`
	AppVersion   string   `yaml:"appVersion"`
	Deprecated   bool     `yaml:"deprecated"`
}

func main() {
	chartPath := path.Join(os.Getenv("INPUT_PATH"), ChartFile)
	fmt.Println(fmt.Sprintf(`Reading values from "%s" file.`, chartPath))

	chartData, readChartError := ioutil.ReadFile(chartPath)
	if readChartError != nil {
		log.Fatalf("Chart Read Error: %v", readChartError)
	}

	chart := Chart{}
	yamlParseError := yaml.Unmarshal([]byte(chartData), &chart)
	if yamlParseError != nil {
		log.Fatalf("YAML Parse Error: %v", yamlParseError)
	}

	githubactions.SetOutput("apiVersion", chart.ApiVersion)
	githubactions.SetOutput("name", chart.Name)
	githubactions.SetOutput("version", chart.Version)
	githubactions.SetOutput("kubeVersion", chart.KubeVersion)
	githubactions.SetOutput("description", chart.Description)
	githubactions.SetOutput("type", chart.Type)
	githubactions.SetOutput("keywords", fmt.Sprintf("%s", strings.Join(chart.Keywords[:], ",")))
	githubactions.SetOutput("home", chart.Home)
	githubactions.SetOutput("sources", fmt.Sprintf("%s", strings.Join(chart.Sources[:], ",")))
	githubactions.SetOutput("repository", chart.Repository)
	githubactions.SetOutput("icon", chart.Icon)
	githubactions.SetOutput("appVersion", chart.AppVersion)
	githubactions.SetOutput("deprecated", fmt.Sprintf("%t", chart.Deprecated))

	for _, dependency := range chart.Dependencies {
		githubactions.SetOutput(fmt.Sprintf(`dependencies_%s_version`, dependency.Name), dependency.Version)
		githubactions.SetOutput(fmt.Sprintf(`dependencies_%s_repository`, dependency.Name), dependency.Repository)
	}
}
