package parse

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// Parse parses the file at rthe given path
func Parse(sourceFile string) (*Root, error) {

	b, err := ioutil.ReadFile(sourceFile)
	if err != nil {
		log.Fatalf("error: %v", err)
		return nil, err
	}

	return _parseYaml(b)
}

// extracted method for better testing
func _parseYaml(content []byte) (*Root, error) {

	var root Root
	err := yaml.Unmarshal(content, &root)
	if err != nil {
		log.Fatalf("error: %v", err)
		return nil, err
	}

	return &root, nil
}
