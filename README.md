# loadinitms

loadinitms is a lightweight microservice framework for managing properties, running primary processes, and handling signals. It simplifies the setup and configuration of microservices in Go.

## Features

- Load properties from YAML or .properties files.
- Run primary processes concurrently.
- Gracefully handle exit signals (SIGHUP, SIGINT, SIGQUIT, SIGABRT, SIGTERM).

## Installation

To use loadinitms in your Go project, you can install it using `go get`:

```bash
go get github.com/gregperez/loadinitms
```

## Usage

code example:

```go
package main

import "github.com/gregperez/loadinitms"

func main() {
	// Create a new instance of LoadInitMS
	lms := loadinitms.NewLoadInitMS()
	
	// Add properties to the list
	lms.AddProperty(&yourPropertyStruct{})

	// Add primary processes to the list
	lms.AddPrimaryProcess(&yourPrimaryProcessStruct{})

	// Set the property file path
	lms.SetPropertyFilePath("path/to/your/properties.yml")

	// Run loadinitms
	lms.Run()
}
```