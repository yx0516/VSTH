package commands

import (
	"fmt"

	"github.com/coreos/etcd/client"
	"github.com/golang/net/context"
	"github.com/kamilhark/etcdsh/common"
	"github.com/kamilhark/etcdsh/pathresolver"
)

type LsCommand struct {
	PathResolver *pathresolver.PathResolver
	KeysApi      client.KeysAPI
}

func (c *LsCommand) Supports(command string) bool {
	return command == "ls"
}

func (c *LsCommand) Handle(args []string) {
	var lsArg = ""
	if len(args) == 1 {
		lsArg = args[0]
	}
	lsPath := c.PathResolver.Resolve(lsArg)
	resp, err := c.KeysApi.Get(context.Background(), lsPath, &client.GetOptions{})

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, node := range resp.Node.Nodes {
		fmt.Println(node.Key)
	}
}

func (c *LsCommand) Verify(args []string) error {
	if len(args) > 2 {
		return common.NewStringError("to many arguments")
	}
	return nil
}

func (c *LsCommand) CommandString() string {
	return "ls"
}

func (o *LsCommand) GetAutoCompleteConfig() AutoCompleteConfig {
	return AutoCompleteConfig{Available: true, OnlyDirs: true}
}
