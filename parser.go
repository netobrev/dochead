package dochead

import (
	"bytes"
	//"fmt"
    "regexp"
	"io/ioutil"
	"github.com/PuerkitoBio/goquery"
	"github.com/russross/blackfriday"
)

func processError(err error) {
    if err != nil {
		panic(err)
	}
}

func parseURI(httpURI string) (verb string, uri string) {
    matching, err := regexp.Match("[A-Z]+\\s+.*", []byte (httpURI))
    if !matching || err != nil {
        return
    }
    
    whitespacePattern := regexp.MustCompile("\\s+")
    splitted := whitespacePattern.Split(httpURI, 2)
    return splitted[0], splitted[1]
}

// ReadAPIDefinition reads an API definition from a markdown file.
func ReadAPIDefinition(file string) ApiDefinition {

	dat, err := ioutil.ReadFile(file)
	processError(err)
    
	output := blackfriday.MarkdownCommon([]byte(string(dat)))
	//fmt.Printf("--------- My HTML output: ---------\n%s\n", output)
	//fmt.Printf("--------- End of HTML output ------\n")

	reader := bytes.NewBuffer([]byte(string(output)))

	doc, err := goquery.NewDocumentFromReader(reader)
	processError(err)

    var apiResources []ApiResource

	doc.Find("h1").Each(func(i int, resource *goquery.Selection) {
		resourceName := resource.Text()

		definitionSelection := resource.NextUntil("h1").WrapAll("*")

        basicHTML, _ := definitionSelection.NextUntil("h2").WrapAll("*").Html()

        verb, uri := parseURI(definitionSelection.Find("code:first-of-type").First().Text())
        apiResource := ApiResource{resourceName, verb, uri, basicHTML}
        
        apiResources = append(apiResources, apiResource)
	})

	/*doc.Find("h2:contains(\"Examples\")").Each(func(i int, s *goquery.Selection) {
	        fmt.Printf("Review2 %d - Example: %s\n", i, s.Text())
	})*/

    return ApiDefinition { apiResources }
}
