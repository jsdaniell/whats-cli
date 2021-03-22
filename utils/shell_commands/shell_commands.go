package shell_commands

import (
	"fmt"
	execute "github.com/alexellis/go-execute/pkg/v1"
)

func ExecuteSh(file, projectName string) {
	cmd := execute.ExecTask{
		Command:     "sh",
		Args:        []string{file},
		StreamStdio: false,
		Cwd:         projectName,
	}

	res, err := cmd.Execute()
	if err != nil {
		panic(err)
	}

	fmt.Printf("output: %s", res.Stderr)
}

func ExecuteShellCommand(command string, projectName string, args ...string) {
	cmd := execute.ExecTask{
		Command:     command,
		Args:        args,
		StreamStdio: false,
		Cwd:         projectName,
	}

	res, err := cmd.Execute()
	if err != nil {
		panic(err)
	}

	fmt.Printf("output: %s", res.Stderr)
}
