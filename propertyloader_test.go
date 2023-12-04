package loadinitms

import (
	"bytes"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadPropertiesFromYAML(t *testing.T) {
	// Preparar un archivo YAML de prueba
	yamlContent := []byte("property1: value1\nproperty2: value2\n")
	yamlFile, err := os.CreateTemp("", "test*.yaml")
	assert.Nil(t, err)
	defer os.Remove(yamlFile.Name())
	os.WriteFile(yamlFile.Name(), yamlContent, 0644)

	// Definir una estructura de prueba
	type TestProperties struct {
		Property1 string `yaml:"property1"`
		Property2 string `yaml:"property2"`
	}

	// Limpiar y configurar las variables globales
	defer func() {
		filePath = "resources/properties.yml"
		propertiesList = make([]interface{}, 0)
	}()

	SetPropertyFilePath(yamlFile.Name())
	AddProperty(&TestProperties{})

	// Capturar la salida estándar para verificar los logs
	var logOutput bytes.Buffer
	log.SetOutput(&logOutput)
	defer func() {
		log.SetOutput(os.Stdout)
	}()

	// Ejecutar la función de prueba
	loadPropertiesFromYAML(yamlFile.Name())

	// Verificar el contenido de la estructura cargada
	assert.Equal(t, "value1", propertiesList[0].(*TestProperties).Property1)
	assert.Equal(t, "value2", propertiesList[0].(*TestProperties).Property2)

	// Verificar los logs (opcional)
	logOutputStr := logOutput.String()
	assert.Contains(t, logOutputStr, "Starting the property loading from the YAML file")
	assert.Contains(t, logOutputStr, "Property loading from the YAML file has been completed")
}

func TestLoadPropertiesFromProperties(t *testing.T) {
	// Preparar un archivo properties de prueba
	propertiesContent := []byte("property1=value1\nproperty2=value2\n")
	propertiesFile, err := os.CreateTemp("", "test*.properties")
	assert.Nil(t, err)
	defer os.Remove(propertiesFile.Name())
	os.WriteFile(propertiesFile.Name(), propertiesContent, 0644)

	// Definir una estructura de prueba
	type TestProperties struct {
		Property1 string `properties:"property1"`
		Property2 string `properties:"property2"`
	}

	// Limpiar y configurar las variables globales
	defer func() {
		filePath = "resources/properties.yml"
		propertiesList = make([]interface{}, 0)
	}()

	SetPropertyFilePath(propertiesFile.Name())
	AddProperty(&TestProperties{})

	// Capturar la salida estándar para verificar los logs
	var logOutput bytes.Buffer
	log.SetOutput(&logOutput)
	defer func() {
		log.SetOutput(os.Stdout)
	}()

	// Ejecutar la función de prueba
	loadPropertiesFromProperties(propertiesFile.Name())

	// Verificar el contenido de la estructura cargada
	assert.Equal(t, "value1", propertiesList[0].(*TestProperties).Property1)
	assert.Equal(t, "value2", propertiesList[0].(*TestProperties).Property2)

	// Verificar los logs (opcional)
	logOutputStr := logOutput.String()
	assert.Contains(t, logOutputStr, "Starting the property loading from the .properties file")
	assert.Contains(t, logOutputStr, "Property loading from the .properties file has been completed")
}
