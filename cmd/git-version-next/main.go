package main

import (
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/Masterminds/semver/v3"
	"github.com/jessevdk/go-flags"
	v "github.com/kazeburo/git-version-next/internal/version"
)

// version by Makefile
var version string

type cmdOpts struct {
	Patch   bool `long:"patch" description:"update patch version"`
	Minor   bool `long:"minor" description:"update minor version"`
	Major   bool `long:"major" description:"update major version"`
	Version bool `short:"v" long:"version" description:"show version"`
}

func _main(args []string) int {
	opts := cmdOpts{}
	psr := flags.NewParser(&opts, flags.HelpFlag|flags.PassDoubleDash)
	args, err := psr.ParseArgs(args)
	if opts.Version {
		fmt.Fprintf(os.Stderr, "Version: %s\nCompiler: %s %s\n",
			version,
			runtime.Compiler,
			runtime.Version())
		os.Exit(0)
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	pwd := "."
	if len(args) > 0 {
		pwd = args[0]
	}

	cv, err := v.Current(pwd)
	if err != nil {
		log.Printf("failed to fetch tags: %v", err)
		return 1
	}

	var nv semver.Version
	if opts.Major {
		nv = cv.IncMajor()
	} else if opts.Minor {
		nv = cv.IncMinor()
	} else if opts.Patch {
		nv = cv.IncPatch()
	} else {
		n, err := v.Next(cv)
		if err != nil {
			log.Printf("failed to calc next tag: %v", err)
			return 1
		}
		nv = *n
	}

	fmt.Printf("%s\n", nv.String())
	return 0
}

func main() {
	os.Exit(_main(os.Args[1:]))
}
