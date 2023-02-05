package util

import (
	"flag"
	"fmt"
	"os"
)

func PrintManual() {
	_, err := fmt.Fprint(os.Stderr, "Manual for commit-message-check:\n")
	_, err = fmt.Fprint(os.Stderr, "- Flags:\n")

	flag.VisitAll(func(f *flag.Flag) {
		_, err = fmt.Fprintf(os.Stderr, "\t%s\t\t%v\n", f.Name, f.Usage)
	})

	_, err = fmt.Fprint(os.Stderr, "- Subcommands:\n")
	_, err = fmt.Fprint(os.Stderr, "\tsetup\t\tInstalls the commit-msg script in every hook directory.\n")
	_, err = fmt.Fprint(os.Stderr, "\tuninstall\tRemoves all commit-msg scripts.\n")
	_, err = fmt.Fprint(os.Stderr, "\tvalidate <PATH>\tValidates the commit message written in <PATH>.\n")

	if err != nil {
		return
	}
}
