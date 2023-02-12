package util

import (
	"fmt"
	"io"
)

func PrintManual(writer io.Writer) {
	_, _ = fmt.Fprint(writer, "Manual for commit-message-check:\n")

	_, _ = fmt.Fprint(writer, "- Flags:\n")
	_, _ = fmt.Fprint(writer, "\tv\t\tShows the current version of the executable.\n\n")

	_, _ = fmt.Fprint(writer, "- Subcommands:\n")
	_, _ = fmt.Fprint(writer, "\tsetup\t\tInstalls the commit-msg script in every hook directory.\n")
	_, _ = fmt.Fprint(writer, "\tuninstall\tRemoves all commit-msg scripts.\n")
	_, _ = fmt.Fprint(writer, "\tvalidate <PATH>\tValidates the commit message written in <PATH>.\n")
}
