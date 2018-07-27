package cli

import "testing"
import (
	"github.com/kamilhark/etcdsh/mocks"
	"github.com/kamilhark/etcdsh/commands"
	"github.com/kamilhark/etcdsh/pathresolver"
	"strings"
	"github.com/coreos/etcd/client"
)

var keysApiMock = mocks.NewKeysApiMock()
var pathResolver = &pathresolver.PathResolver{}
var commandsArray = []commands.Command{
	&commands.CdCommand{PathResolver: pathResolver, KeysApi: keysApiMock},
	&commands.LsCommand{PathResolver: pathResolver, KeysApi: keysApiMock},
	&commands.GetCommand{PathResolver: pathResolver, KeysApi: keysApiMock},
	&commands.SetCommand{PathResolver: pathResolver, KeysApi: keysApiMock},
	&commands.RmCommand{PathResolver: pathResolver, KeysApi: keysApiMock},
	&commands.ExitCommand{},
}
var completer = (&Completer{keysApiMock, commandsArray, pathResolver}).Get

func TestCompleteCommandsNames(t *testing.T) {
	assertContainHint(t, completer("c"), "cd")
	assertContainHint(t, completer("s"), "set")
	assertContainHint(t, completer("r"), "rm")
	assertContainHint(t, completer("l"), "ls")

	//when there is no input, all commands should be given
	hints := completer("")
	if len(hints) != len(commandsArray) {
		t.Fail()
	}
}

func TestCompleteFirstDirArgumentWhenInRootDir(t *testing.T) {

	node := createDirNode("/aa")
	nodes := client.Nodes{node, createDirNode("/ab"), createDirNode("/bb"), createValueNode("aaa")}

	rootNode := &client.Node{
		Nodes: nodes,
	}

	response := &client.Response{"", rootNode, nil, 1}
	keysApiMock.MockGet(pathResolver.CurrentPath(), response)

	hints := completer("cd a")

	assertLength(t, hints, 2)
	assertContainHint(t, hints, "cd aa", "cd ab")
}

func TestCompleteFirstDirArgumentWhenInChildDir(t *testing.T) {
	pathResolver.Add("child")
	rootNode := &client.Node{}
	rootNode.Nodes = []*client.Node{createDirNode("/child/aa"), createDirNode("/child/ab")}

	response := &client.Response{"", rootNode, nil, 1}
	keysApiMock.MockGet(pathResolver.CurrentPath(), response)

	hints := completer("cd a")

	assertLength(t, hints, 2)
	assertContainHint(t, hints, "cd aa", "cd ab")
}

func TestShouldNotCompleteForExitCommand(t *testing.T) {
	rootNode := &client.Node{}
	rootNode.Nodes = []*client.Node{createDirNode("/aa"), createDirNode("/ab")}

	response := &client.Response{"", rootNode, nil, 1}
	keysApiMock.MockGet(pathResolver.CurrentPath(), response)

	hints := completer("exit ")

	assertLength(t, hints, 0)
}

func TestShouldCompleteValueNodesWhenGetCommand(t *testing.T) {
	rootNode := &client.Node{
		Nodes: client.Nodes{createDirNode("/aa"), createValueNode("/ab")},
	}

	response := &client.Response{"", rootNode, nil, 1}
	keysApiMock.MockGet(pathResolver.CurrentPath(), response)

	hints := completer("get a")

	assertLength(t, hints, 2)
	assertContainHint(t, hints, "get aa", "get ab")
}

func assertContainHint(t *testing.T, actualHints []string, expectedHints ...string) {
	for _, hint := range expectedHints {
		found := false
		for _, a := range actualHints {
			if a == hint {
				found = true
			}
		}
		if !found {
			t.Errorf("actual hints [%s] does not contain %s", strings.Join(actualHints, ","), hint)
		}
	}
}

func assertLength(t *testing.T, slice []string, expectedLength int) {
	if len(slice) != expectedLength {
		t.Errorf("expected size %d but was %d", expectedLength, len(slice))
		t.Fail()
	}
}

func createDirNode(key string) *client.Node {
	return &client.Node{Dir:true, Key:key}
}

func createValueNode(key string) *client.Node {
	return &client.Node{Dir:false, Key:key}
}

