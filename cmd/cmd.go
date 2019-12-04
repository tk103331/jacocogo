package cmd

import "fmt"

var cmds map[string]cmd = make(map[string]cmd, 0)

type cmd interface {
	name() string
	desc() string
	parse([]string)
	exec() error
}

func Exec(cmdName string, args []string) error {
	if cmd, ok := cmds[cmdName]; ok {
		cmd.parse(args)
		return cmd.exec()
	} else {
		Usage()
	}
	return nil
}

func Usage() {
	fmt.Println("Usage: jacocogo <command> [arguments]")
	fmt.Println("Supported commands:")
	for _, c := range cmds {
		fmt.Printf("\t%s:  %s\n", c.name(), c.desc())
	}
}
