package dochead

import (
	"bufio"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/rveen/ogdl"
)

func processError(err error) {
	if err != nil {
		panic(err)
	}
}

func parseURI(httpURI string) (verb string, uri string) {
	matching, err := regexp.Match("[A-Z]+\\s+.*", []byte(httpURI))
	if !matching || err != nil {
		return
	}

	whitespacePattern := regexp.MustCompile("\\s+")
	splitted := whitespacePattern.Split(httpURI, 2)
	return splitted[0], splitted[1]
}

// ReadAPIDefinition reads an API definition from a markdown file.
func ReadAPIDefinition(file string) ApiDefinition {
	document := ogdl.ParseFile(file)

    var apiResources []ApiResource
    for i := 0; i < document.Len(); i++ {
        resource := document.Get(fmt.Sprintf("resource{%d}", i))

        apiResource := parseResource(resource)
        apiResources = append(apiResources, apiResource)
    }
    return ApiDefinition{apiResources}
}

func parseResource(resource *ogdl.Graph) ApiResource {
	name, _ := resource.GetString("name")
	verb, _ := resource.GetString("verb")
	uri, _ := resource.GetString("uri")
	description, _ := resource.GetString("description")

    allParameters, _ := resource.GetSimilar("parameter")
	parameters := parseParameter(allParameters)
	body := parseBody(resource.Node("body"))
	ret := parseReturn(resource.Node("return"))
    
    allExamples, _ := resource.GetSimilar("example")
	examples := parseExamples(allExamples)

	return ApiResource{
		name,
		verb,
		uri,
		description,
		parameters,
		body,
		ret,
		examples,
	}
}

func parseParameter(graph *ogdl.Graph) []Parameter {
	var parameters []Parameter

	for i := 0; i < graph.Len(); i++ {
		parameter := graph.GetAt(i)

		id, _ := parameter.GetString("name")
		valueType, _ := parameter.GetString("type")
		description, _ := parameter.GetString("description")

		parameters = append(parameters, Parameter{id, valueType, description})
	}
	return parameters
}

func parseBody(graph *ogdl.Graph) Body {
	accept, _ := graph.GetString("accept")
	schema, _ := graph.GetString("schema")
	return Body{accept, schema}
}

func parseReturn(graph *ogdl.Graph) Return {
	contentType, _ := graph.GetString("content_type")
	schema, _ := graph.GetString("schema")
    
    allStatusCodes, _ := graph.GetSimilar("status")
	return Return{contentType, schema, parseStatus(allStatusCodes)}
}

func parseStatus(graph *ogdl.Graph) map[int]string {
	codes := make(map[int]string)
	for i := 0; i < graph.Len(); i++ {
		codeGraph := graph.GetAt(i)
		code, _ := codeGraph.GetString("code")
        codeInt, _ := strconv.Atoi(code)
        description, _ := codeGraph.GetString("description")
		codes[codeInt] = description
	}
	return codes
}

func parseExamples(graph *ogdl.Graph) []Example {
	var examples []Example

	for i := 0; i < graph.Len(); i++ {
		example := graph.GetAt(i)

		name, _ := example.GetString("name")

		requestString, _ := example.GetString("request")
		requestString = strings.Replace(requestString, "\n", "\r\n", -1) + "\r\n"
		requestReader := bufio.NewReader(strings.NewReader(requestString))
		httpRequest, _ := http.ReadRequest(requestReader)

		responseString, _ := example.GetString("response")
		responseString = strings.Replace(responseString, "\n", "\r\n", -1) + "\r\n"
		responseReader := bufio.NewReader(strings.NewReader(responseString))
		httpResponse, _ := http.ReadResponse(responseReader, httpRequest)

		examples = append(examples, Example{name, httpRequest, httpResponse})
	}

	return examples
}

func printChildren(graph *ogdl.Graph) {
	fmt.Printf("\n\n")
	fmt.Printf("The node: %s\n", graph)
	for i := 0; i < graph.Len(); i++ {
		fmt.Printf("Child Node: %s\n", graph.GetAt(i).String())
	}
	fmt.Printf("\n\n")
}
