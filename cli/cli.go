package cli

func CLI(argv []string) (Commands, error) {
	return parseCommands(argv)
}
