package fix

import (
	_ "embed"
	"fmt"
	ollama "github.com/jmorganca/ollama/api"
)

//go:embed Modelfile.fix
var fixModelfile string

type Fix struct {
	api *ollama.Client

	tag       string
	modelfile string
}

func (f *Fix) Tag() string {
	return f.tag
}

func (f *Fix) Modelfile() string {
	return f.modelfile
}

func New(api *ollama.Client, version string) *Fix {
	tag := fmt.Sprintf("tlm:%s-f", version)
	return &Fix{api: api, tag: tag, modelfile: fixModelfile}
}
