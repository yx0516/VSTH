package cli

import "flag"
import "fmt"
import "strings"
import "log"
import "github.com/kamilhark/etcdsh/commands"
import "github.com/kamilhark/etcdsh/pathresolver"

import (
	"github.com/coreos/etcd/client"
	"github.com/peterh/liner"
	"time"
)

func Start() {
	etcdUrl := getEtcdUrl()
	pathResolver := new(pathresolver.PathResolver)
	cfg := client.Config{
		Endpoints: []string{etcdUrl},
		Transport: client.DefaultTransport,
		// set timeout per request to fail fast when the target endpoint is unavailable
		HeaderTimeoutPerRequest: time.Second,
	}

	c, err := client.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	api := client.NewKeysAPI(c)

	fmt.Println("connected to etcd")

	console := liner.NewLiner()
	console.SetTabCompletionStyle(liner.TabCircular)
	commandsArray := []commands.Command{
		&commands.ExitCommand{State: console},
		&commands.CdCommand{PathResolver: pathResolver, KeysApi: api},
		&commands.LsCommand{PathResolver: pathResolver, KeysApi: api},
		&commands.GetCommand{PathResolver: pathResolver, KeysApi: api},
		&commands.SetCommand{PathResolver: pathResolver, KeysApi: api},
		&commands.RmCommand{PathResolver: pathResolver, KeysApi: api},
	}

	defer console.Close()
	console.SetCtrlCAborts(true)
	completer := (&Completer{api, commandsArray, pathResolver}).Get
	console.SetCompleter(completer)

	for {
		line, err := console.Prompt(pathResolver.CurrentPath() + ">")

		if err != nil && err == liner.ErrPromptAborted {
			return
		}

		if len(line) == 0 {
			continue
		}

		tokens := strings.Split(line, " ")
		if len(tokens) == 0 {
			continue
		}

		console.AppendHistory(line)

		command := tokens[0]
		args := tokens[1:]
		found := false
		for _, commandHandler := range commandsArray {
			if commandHandler.Supports(command) {
				found = true
				err := commandHandler.Verify(args)
				if err != nil {
					fmt.Println(err)
				} else {
					commandHandler.Handle(args)
				}
				break
			}
		}
		if !found {
			fmt.Println("invalid command")
		}
		printPrompt(pathResolver)
	}
}

func getEtcdUrl() string {
	var url = flag.String("url", "http://localhost:2379", "etcd url")
	flag.Parse()
	return *url
}

func printPrompt(pathResolver *pathresolver.PathResolver) {
	fmt.Print()
}
