package Utils

import (
	"Silver/Silver/Helpers"
	"net/http"
	"os"
	"regexp"
)

type Response struct{}

func (response Response) SendString(data string) {
	Helpers.CurrentResponseWriter.Write([]byte(data))
}

type ViewParams map[string]string

func (response Response) SendView(view string) {
	http.ServeFile(Helpers.CurrentResponseWriter, Helpers.CurrentRequest, "./Views/"+view)
}

// SendTemplateView sending view that contains data which should be replaced {{ variable }} -> hello world
func (response Response) SendTemplateView(view string, variables ViewParams) {
	htmlContent, _ := os.ReadFile("./Views/" + view)
	regex := regexp.MustCompile(`\{\{\s*([a-zA-Z_][a-zA-Z0-9_]*)\s*\}\}`)

	updatedHtml := regex.ReplaceAllStringFunc(string(htmlContent), func(match string) string {
		variableName := regex.FindStringSubmatch(match)[1]
		replacement, exists := variables[variableName]

		if exists {
			return replacement
		}

		return match
	})

	Helpers.CurrentResponseWriter.Write([]byte(updatedHtml))
}
