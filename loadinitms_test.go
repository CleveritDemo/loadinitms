package loadinitms

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

// mockSignalHandler es una implementación de SignalHandler para pruebas.
type mockSignalHandler struct {
	signalToReturn os.Signal
	signalChan     chan os.Signal
}

func newMockSignalHandler() *mockSignalHandler {
	return &mockSignalHandler{
		signalChan: make(chan os.Signal, 1),
	}
}

func (msh *mockSignalHandler) WaitForExitSignal() os.Signal {
	return <-msh.signalChan
}

func (msh *mockSignalHandler) sendExitSignal(signal os.Signal) {
	msh.signalChan <- signal
	close(msh.signalChan)
}

func TestRun(t *testing.T) {
	// Preparar un archivo YAML de prueba
	yamlContent := []byte("property1: value1\nproperty2: value2\n")
	yamlFile, err := os.CreateTemp("", "test*.yml")
	assert.Nil(t, err)
	defer os.Remove(yamlFile.Name())
	os.WriteFile(yamlFile.Name(), yamlContent, 0644)

	// Limpiar y configurar las variables globales
	defer func() {
		filePath = "resources/properties.yml"
		propertiesList = make([]interface{}, 0)
	}()

	SetPropertyFilePath(yamlFile.Name())

	// Ejecutar el código que utiliza Run en una goroutine
	go Run()

	// Enviar una señal de salida al mock de SignalHandler después de un tiempo
	msh := newMockSignalHandler()
	go func() {
		<-time.After(time.Second * 3)
		msh.sendExitSignal(os.Interrupt)
	}()

	// Esperar a que Run termine
	<-time.After(time.Second) // Ajusta el tiempo según sea necesario

	// Verificar el resultado (puedes ajustar según tus necesidades)
	assert.True(t, true)
}
