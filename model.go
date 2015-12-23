package dochead

import (
    "fmt"
)

type ApiDefinition struct {
    Resources []ApiResource
}

type ApiResource struct {
	Name string
	URI  string
}

func (resource ApiResource) String() string {
    return fmt.Sprintf("API Resource: %s (Name: %s)", resource.URI, resource.Name)
}