package commands

import (
	"fmt"
	"strings"

	"github.com/coreos/etcd/client"
	"github.com/golang/net/context"
	"github.com/kamilhark/etcdsh/common"
	"github.com/kamilhark/etcdsh/pathresolver"
)

type RmCommand struct {
	PathResolver *pathresolver.PathResolver
	KeysApi      client.KeysAPI
}

func (c *RmCommand) Supports(command string) bool {
	return strings.EqualFold(command, "rm")
}

func (c *RmCommand) Handle(args []string) {
	for i := 0; i < len(args); i++ {
		key := c.PathResolver.Resolve(args[i])
		_, err := c.KeysApi.Delete(context.Background(), key, &client.DeleteOptions{})
		if err != nil {
			fmt.Println(err)
		}
	}
}

func (c *RmCommand) Verify(args []string) error {
	if len(args) < 1 {
		return common.NewStringError("wrong number of arguments, rm command requires at least one argument")
	}
	return nil
}

func (c *RmCommand) CommandString() string {
	return "rm"
}

func (o *RmCommand) GetAutoCompleteConfig() AutoCompleteConfig {
	return AutoCompleteConfig{Available: true, OnlyDirs: false}
}
