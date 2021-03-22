package go_commands

import (
	"fmt"
	execute "github.com/alexellis/go-execute/pkg/v1"
)

func GoModInit(username, projectName string) {
	cmd := execute.ExecTask{
		Command:     "go",
		Args:        []string{"mod", "init", "github.com/" + username + "/" + projectName},
		StreamStdio: false,
		Cwd:         projectName,
	}

	res, err := cmd.Execute()
	if err != nil {
		panic(err)
	}

	fmt.Printf("output: %s", res.Stderr)
}

func GoModTidy(projectName string) {
	cmd := execute.ExecTask{
		Command:     "go",
		Args:        []string{"mod", "tidy"},
		StreamStdio: false,
		Cwd:         projectName,
	}

	res, err := cmd.Execute()
	if err != nil {
		panic(err)
	}

	fmt.Printf("output: %s", res.Stderr)
}

func GoGet(packageName, projectName string) {
	cmd := execute.ExecTask{
		Command:     "go",
		Args:        []string{"get", "-u", packageName},
		StreamStdio: false,
		Cwd:         projectName,
	}

	res, err := cmd.Execute()
	if err != nil {
		panic(err)
	}

	fmt.Printf("output: %s", res.Stderr)
}

func GoModVendor(projectName string) {
	cmd := execute.ExecTask{
		Command:     "go",
		Args:        []string{"mod", "vendor"},
		StreamStdio: false,
		Cwd:         projectName,
	}

	res, err := cmd.Execute()
	if err != nil {
		panic(err)
	}

	fmt.Printf("output: %s", res.Stderr)
}
