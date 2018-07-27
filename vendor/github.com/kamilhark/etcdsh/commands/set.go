package commands

import (
	"fmt"
	"strings"

	"github.com/coreos/etcd/client"
	"github.com/golang/net/context"
	"github.com/kamilhark/etcdsh/common"
	"github.com/kamilhark/etcdsh/pathresolver"
)

type SetCommand struct {
	PathResolver *pathresolver.PathResolver
	KeysApi      client.KeysAPI
}

func (c *SetCommand) Supports(command string) bool {
	return strings.EqualFold(command, "set")
}

func (c *SetCommand) Handle(args []string) {
	key := c.PathResolver.Resolve(args[0])
	value := args[1]
	_, err := c.KeysApi.Set(context.Background(), key, value, &client.SetOptions{Dir: false})
	if err != nil {
		fmt.Println(err)
	}
}

func (c *SetCommand) Verify(args []string) error {
	if len(args) != 2 {
		return common.NewStringError("wrong number of arguments, set command requires two argument")
	}
	return nil
}

func (c *SetCommand) CommandString() string {
	return "set"
}

func (o *SetCommand) GetAutoCompleteConfig() AutoCompleteConfig {
	return AutoCompleteConfig{Available: true, OnlyDirs: true}
}
