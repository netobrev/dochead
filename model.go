package dochead

import (
    "net/http"
)

type ApiDefinition struct {
    Resources []ApiResource
}

type ApiResource struct {
	Name string
    Verb string
	URI  string
    Description string
    
    Parameters []Parameter
    Body Body
    Return Return

    Examples []Example
}

type Parameter struct {
    Name string
    Type string
    Description string
}

type Body struct {
    Accept string
    Schema string
}

type Return struct {
    ContentType string
    Schema string
    Codes map[int]string
}

type Example struct {
    Name string
    Request *http.Request
    Response *http.Response
}