package loadinitms

import (
	"log"
	"os"
	"path/filepath"

	"github.com/magiconair/properties"
	"gopkg.in/yaml.v3"
)

var (
	filePath       = "resources/properties.yml"
	propertiesList = make([]interface{}, 0)
)

// SetPropertyFilePath sets the property file path. Default is "resources/properties.yml"
func SetPropertyFilePath(path string) {
	filePath = path
}

// AddProperty adds a property to the list of properties.
func AddProperty(p interface{}) {
	propertiesList = append(propertiesList, p)
}

// loadPropertiesFromYAML loads the properties from a YAML file.
func loadPropertiesFromYAML(filename string) {
	log.Println("Starting the property loading from the YAML file: ", filePath)

	yamlFile, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error while reading '%s' file: %s\n", filePath, err)
	}

	yamlEnv := os.ExpandEnv(string(yamlFile))

	for _, p := range propertiesList {
		if err := yaml.Unmarshal([]byte(yamlEnv), p); err != nil {
			log.Fatalf("Error deserializing the YAML file '%s': %s", filePath, err)
		}
	}

	log.Println("Property loading from the YAML file has been completed: ", filePath)
}

// loadPropertiesFromProperties loads the properties from a .properties file.
func loadPropertiesFromProperties(filename string) {
	log.Println("Starting the property loading from the .properties file: ", filePath)

	props, err := properties.LoadFile(filename, properties.UTF8)
	if err != nil {
		log.Fatalf("Error loading the .properties file '%s': %s", filePath, err)
	}

	for _, p := range propertiesList {
		if err := props.Decode(p); err != nil {
			log.Fatalf("Error decoding properties from the .properties file '%s': %s", filePath, err)
		}
	}

	log.Println("Property loading from the .properties file has been completed: ", filePath)
}

// LoadProperties loads the properties from the YAML or .properties file.
func LoadProperties() {
	ext := filepath.Ext(filePath)

	filename, err := filepath.Abs(filePath)
	if err != nil {
		log.Fatalf("Error obtaining the file path '%s': %s", filePath, err)
	}

	if ext == ".yml" {
		loadPropertiesFromYAML(filename)
	} else if ext == ".properties" {
		loadPropertiesFromProperties(filename)
	} else {
		log.Fatalf("The file format is not compatible: %s", ext)
	}
}
