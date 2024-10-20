package install

import (
	"context"
	"fmt"
	"github.com/charmbracelet/huh/spinner"
	ollama "github.com/jmorganca/ollama/api"
	"github.com/yusufcanb/tlm/shell"
	"os"
	"time"
)

func (i *Install) installModelfile(name, modelfile string) error {
	var err error
	err = i.api.Create(context.Background(), &ollama.CreateRequest{Model: name, Modelfile: modelfile}, func(res ollama.ProgressResponse) error {
		return nil
	})
	return err
}

func (i *Install) deployTlm() {
	var err error

	_ = spinner.New().Type(spinner.Line).Title("Getting latest my_codellama_model").Action(func() {
		err = i.api.Pull(context.Background(), &ollama.PullRequest{Model: "my_codellama_model"}, func(res ollama.ProgressResponse) error {
			return nil
		})
		if err != nil {
			fmt.Println("- Installing my_codellama_model. " + shell.Err())
			os.Exit(-1)
		}
	}).Run()
	fmt.Println("- Getting latest my_codellama_model. " + shell.Ok())

	// 6. Install the modelfile (Suggest)
	_ = spinner.New().Type(spinner.Line).Title("Creating Modelfile for suggestions").Action(func() {
		err = i.installModelfile(i.suggest.Tag(), i.explain.Modelfile())
		time.Sleep(1 * time.Second)
		if err != nil {
			fmt.Println("- Creating Modelfile for suggestions. " + shell.Err())
			fmt.Println("\n" + err.Error())
			os.Exit(-1)
		}
	}).Run()
	fmt.Println("- Creating Modelfile for suggestions. " + shell.Ok())

	// 7. Install the modelfile (Explain)
	_ = spinner.New().Type(spinner.Line).Title("Creating Modelfile for explanations").Action(func() {
		err = i.installModelfile(i.explain.Tag(), i.explain.Modelfile())
		time.Sleep(1 * time.Second)
		if err != nil {
			fmt.Println("- Creating Modelfile for explanations. " + shell.Err())
			os.Exit(-1)
		}
	}).Run()
	fmt.Println("- Creating Modelfile for explanations. " + shell.Ok())
}
