package commands

import (
    "github.com/kamilhark/etcdsh/pathresolver"
	"github.com/kamilhark/etcdsh/common"
	"github.com/coreos/etcd/client"
	"github.com/golang/net/context"
)

type CdCommand struct {
	PathResolver *pathresolver.PathResolver
	KeysApi      client.KeysAPI
}

func (cdCommand *CdCommand) Supports(command string) bool {
	return command == "cd"
}

func (cdCommand *CdCommand) Handle(args []string) {
	if len(args) == 1 {
		cdCommand.PathResolver.GoTo(args[0])
	} else {
		cdCommand.PathResolver.GoTo("")
	}
}

func (cdCommand *CdCommand) Verify(args []string) error {
	if len(args) > 1 {
		return common.NewStringError("'cd' command supports only one argument")
	}

	if len(args) == 0 {
		return nil
	}

	nextPath := cdCommand.PathResolver.Resolve(args[0])
	response, err := cdCommand.KeysApi.Get(context.Background(), nextPath, &client.GetOptions{})
	if err != nil {
		return err
	}

	if !response.Node.Dir {
		return common.NewStringError("not a directory")
	}

	return nil
}

func (cdCommand *CdCommand) CommandString() string {
	return "cd"
}

func (o *CdCommand) GetAutoCompleteConfig() AutoCompleteConfig {
	return AutoCompleteConfig{Available:true, OnlyDirs:true}
}

