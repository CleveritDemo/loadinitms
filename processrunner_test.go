package loadinitms

import (
	"bytes"
	"log"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// MockPrimaryProcess es una implementación de PrimaryProcess para pruebas.
type MockPrimaryProcess struct {
	started bool
}

func (m *MockPrimaryProcess) Start() {
	m.started = true
}

func TestAddPrimaryAndRunPrimaries(t *testing.T) {
	// Limpiar la lista de primarios
	primaries = make([]func() PrimaryProcess, 0)

	// Mock la salida de log para verificar los logs generados
	var logOutput bytes.Buffer
	log.SetOutput(&logOutput)
	defer func() {
		log.SetOutput(os.Stdout)
	}()

	// Crear un mock de PrimaryProcess y agregarlo a la lista de primarios
	mockPrimary := &MockPrimaryProcess{}
	AddPrimary(func() PrimaryProcess {
		return mockPrimary
	})

	// Ejecutar runPrimaries
	runPrimaries()

	// Esperar un tiempo para permitir que las goroutines se ejecuten
	time.Sleep(time.Millisecond * 10)

	// Verificar que el proceso primario se haya iniciado
	assert.True(t, mockPrimary.started)

	// Verificar los logs (opcional)
	logOutputStr := logOutput.String()
	assert.Contains(t, logOutputStr, "Running primary process")

	// Limpiar la lista de primarios después de las pruebas
	primaries = make([]func() PrimaryProcess, 0)
}
