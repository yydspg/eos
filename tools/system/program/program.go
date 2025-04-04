package system
import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func ExitWithError(err error) {
	programName := filepath.Base(os.Args[0])
	fmt.Fprintf(os.Stderr,"%s exit -1 : %+v\n",programName,err)
	os.Exit(-1)
}

func SIGTEMExit() {
	programName := filepath.Base(os.Args[0])
	fmt.Fprintf(os.Stderr,"Warning %s receive process terminal SIGTERM exit 0\n",programName)
}

func getProcessName() string {
	args := os.Args
	if len(args) > 0 {
		segments := strings.Split(args[0], "/")
		if len(segments) > 0 {
			return segments[len(segments)-1]
		}
	}
	return ""
}