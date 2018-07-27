package commands

import (
	"fmt"
	"strings"

	"github.com/coreos/etcd/client"
	"github.com/golang/net/context"
	"github.com/kamilhark/etcdsh/common"
	"github.com/kamilhark/etcdsh/pathresolver"
)

type GetCommand struct {
	PathResolver *pathresolver.PathResolver
	KeysApi      client.KeysAPI
}

func (c *GetCommand) Supports(command string) bool {
	return strings.EqualFold(command, "get")
}

func (c *GetCommand) Handle(args []string) {
	key := c.PathResolver.Resolve(args[0])
	response, err := c.KeysApi.Get(context.Background(), key, &client.GetOptions{})
	if err != nil {
		fmt.Println(err)
	} else {
		if response.Node.Dir {
			fmt.Println("dir provided, no value")
		} else {
			fmt.Println(response.Node.Value)
		}
	}
}

func (c *GetCommand) Verify(args []string) error {
	if len(args) != 1 {
		return common.NewStringError("wrong number of arguments, get command requires one argument")
	}
	return nil
}

func (c *GetCommand) CommandString() string {
	return "get"
}

func (o *GetCommand) GetAutoCompleteConfig() AutoCompleteConfig {
	return AutoCompleteConfig{Available: true, OnlyDirs: false}
}
