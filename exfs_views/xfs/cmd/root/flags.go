package root

// Verbose enables verbose output across all commands.
var Verbose bool

func registerGlobalFlags() {
	Cmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "V", false, "enable verbose output")
}
