package dochead

import (
    "fmt"
)

type ApiDefinition struct {
    Resources []ApiResource
}

type ApiResource struct {
	Name string
    Verb string
	URI  string
    Html string
}

func (resource ApiResource) String() string {
    return fmt.Sprintf("API Resource: %s %s (Name: %s)\n%s", resource.Verb, resource.URI, resource.Name, resource.Html)
}