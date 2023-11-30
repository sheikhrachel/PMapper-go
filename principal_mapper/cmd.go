package principal_mapper

type Command struct {
}

// Globals - Global flags available to any subcommand
type Globals struct {
	Profile string `help:"The AWS CLI (botocore) profile to use to call the AWS API."`
	Account string `help:"When running offline operations, this parameter determines which account to act against."`
	Debug   bool   `help:"Produces debug-level output of the underlying Principal Mapper library during execution."`
}

type Graph struct{}
type Orgs struct{}
type Query struct{}
type ArgQuery struct{}
type Repl struct{}
type Visualize struct{}
type Analysis struct{}

// CLI includes global flags and registers subcommands
type CLI struct {
	Globals

	Graph     Graph     `cmd:"graph" help:"Pulls information for an AWS account's use of IAM."`
	Orgs      Orgs      `cmd:"orgs" help:"Pulls information for an AWS Organization"`
	Query     Query     `cmd:"query" help:"Displays information corresponding to a query"`
	ArgQuery  ArgQuery  `cmd:"arg_query" help:"Displays information corresponding to an arg-specified query"`
	Repl      Repl      `cmd:"repl" help:"Runs a REPL for querying"`
	Visualize Visualize `cmd:"visualize" help:"Generates an image representing the AWS account"`
	Analysis  Analysis  `cmd:"analysis" help:"Analyzes and reports identified issues"`
}
