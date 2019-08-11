package kube

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/cyberbliss/kube-prompt/internal/debug"
)

var LivePrefixState struct {
	LivePrefix string
	IsEnable   bool
}

func ChangeLivePrefix() (string, bool) {
	return LivePrefixState.LivePrefix, LivePrefixState.IsEnable
}

func Executor(s string) {
	s = strings.TrimSpace(s)
	debug.Log(s)
	if s == "" {
		return
	} else if s == "quit" || s == "exit" {
		fmt.Println("Bye!")
		os.Exit(0)
		return
	}

	elems := strings.Split(s, " ")
	namespaceChange := false
	executable := "kubectl "

	switch elems[0] {
	case "ns", "namespace":
		if len(elems) < 2 {
			s = "get namespaces"
		} else {
			currCtx := GetKubeContext()
			s = fmt.Sprintf("config set-context %s --namespace %s", currCtx, elems[1])
			namespaceChange = true
		}
	case "sh":
		if len(elems) < 2 {
			return
		}
		s = strings.Join(elems[1:], " ")
		executable = ""
	}

	cmd := exec.Command("/bin/sh", "-c", executable+s)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Got error: %s\n", err.Error())
	}

	if namespaceChange && err == nil {
		removeLastFetchedAt(activeNamespaceKey)
		LivePrefixState.LivePrefix = elems[1] + ": "
		LivePrefixState.IsEnable = true
	}
	return
}

func ExecuteAndGetResult(s string) string {
	s = strings.TrimSpace(s)
	if s == "" {
		debug.Log("you need to pass the something arguments")
		return ""
	}

	out := &bytes.Buffer{}
	cmd := exec.Command("/bin/sh", "-c", "kubectl "+s)
	cmd.Stdin = os.Stdin
	cmd.Stdout = out
	if err := cmd.Run(); err != nil {
		debug.Log(err.Error())
		return ""
	}
	r := string(out.Bytes())
	return r
}
