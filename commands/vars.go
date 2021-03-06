package commands

import (
	stks "github.com/itc3-devops/qaz/stacks"

	"github.com/itc3-devops/qaz/logger"
	"github.com/itc3-devops/qaz/repo"
)

var (
	config  stks.Config
	stacks  stks.Map
	region  string
	project string
	gitrepo repo.Repo
	log     = logger.Logger{
		DebugMode: &run.debug,
		Colors:    &run.colors,
	}
)

// config environment variable
const (
	configENV = "QAZ_CONFIG"

	// OutputRegex for printing yaml/json output
	OutputRegex = `(?m)^[ ]*([^\r\n:]+?)\s*:`
)

// run.var used as a central point for command data from flags
var run = struct {
	cfgSource  string
	tplSource  string
	profile    string
	tplSources []string
	stacks     map[string]string
	all        bool
	version    bool
	request    string
	debug      bool
	funcEvent  string
	lambdAsync bool
	changeName string
	stackName  string
	rollback   bool
	colors     bool
	cfgRaw     string
	gituser    string
	gitpass    string
	gitrsa     string
}{}
