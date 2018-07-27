package cli

import (
	"strings"

	"github.com/coreos/etcd/client"
	"github.com/golang/net/context"
	"github.com/kamilhark/etcdsh/commands"
	"github.com/kamilhark/etcdsh/pathresolver"
)

type Completer struct {
	KeysApi       client.KeysAPI
	CommandsArray []commands.Command
	PathResolver  *pathresolver.PathResolver
}

func (c *Completer) Get(line string) []string {

	tokens := strings.Split(line, " ")

	if len(tokens) == 1 { //user entered only a command (or part of a command) name without arguments
		return c.completeCommand(tokens)
	}

	if len(tokens) == 2 { //user entered full command name and part of argument
		return c.completeArgument(line, tokens)
	}

	return []string{}
}

func (c *Completer) completeCommand(tokens []string) (result []string) {
	for _, commandHandler := range c.CommandsArray {
		if strings.HasPrefix(commandHandler.CommandString(), tokens[0]) {
			result = append(result, commandHandler.CommandString())
		}
	}
	return
}

func (c *Completer) completeArgument(line string, tokens []string) (result []string) {

	commandHandler := c.getCommandHandler(line)
	if commandHandler == nil || !commandHandler.GetAutoCompleteConfig().Available {
		return
	}

	autoCompleteConfig := commandHandler.GetAutoCompleteConfig()

	response, _ := c.KeysApi.Get(context.Background(), c.PathResolver.CurrentPath(), &client.GetOptions{})
	nodes := response.Node.Nodes

	for _, node := range nodes {
		lastIndexOfSlash := strings.LastIndex(node.Key, "/")
		key := node.Key[lastIndexOfSlash+1:]
		if strings.HasPrefix(key, tokens[1]) && (node.Dir || !autoCompleteConfig.OnlyDirs) {
			result = append(result, commandHandler.CommandString()+" "+key)
		}
	}

	return
}

func (c *Completer) getCommandHandler(line string) commands.Command {
	for _, commandHandler := range c.CommandsArray {
		if strings.HasPrefix(line, commandHandler.CommandString()) {
			return commandHandler
		}
	}
	return nil
}
