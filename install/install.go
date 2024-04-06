package install

import (
	ollama "github.com/jmorganca/ollama/api"
	"github.com/yusufcanb/tlm/explain"
	"github.com/yusufcanb/tlm/fix"
	"github.com/yusufcanb/tlm/suggest"
)

var repositoryOwner = "yusufcanb"
var repositoryName = "tlm"

type Install struct {
	api *ollama.Client

	suggest *suggest.Suggest
	explain *explain.Explain
	fix     *fix.Fix

	ReleaseManager *ReleaseManager
}

func New(api *ollama.Client, suggest *suggest.Suggest, explain *explain.Explain, fix *fix.Fix) *Install {
	return &Install{
		api:            api,
		suggest:        suggest,
		explain:        explain,
		fix:            fix,
		ReleaseManager: NewReleaseManager(repositoryOwner, repositoryName),
	}
}
