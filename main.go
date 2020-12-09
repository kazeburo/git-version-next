package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"sort"

	"github.com/Masterminds/semver/v3"
	"github.com/jessevdk/go-flags"
	"github.com/manifoldco/promptui"
)

// version by Makefile
var version string

type cmdOpts struct {
	Patch   bool `long:"patch" description:"update patch version"`
	Minor   bool `long:"minor" description:"update minor version"`
	Major   bool `long:"major" description:"update major version"`
	Version bool `short:"v" long:"version" description:"show version"`
}

func currentVersion(pwd string) (*semver.Version, error) {
	if pwd[len(pwd):] != "/" {
		pwd = pwd + "/"
	}

	args := []string{}
	args = append(args, fmt.Sprintf("--work-tree=%s", pwd))
	args = append(args, fmt.Sprintf("--git-dir=%s.git/", pwd))
	args = append(args, "tag")
	cmd := exec.Command("git", args...)
	cmd.Stderr = os.Stderr
	outout, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(bytes.NewReader(outout))
	vs := make([]*semver.Version, 0)
	for scanner.Scan() {
		t := scanner.Text()
		v, err := semver.NewVersion(t)
		if err != nil {
			continue
		}
		vs = append(vs, v)
	}
	if len(vs) == 0 {
		initial, _ := semver.NewVersion("0.0.0")
		return initial, nil
	}

	sort.Sort(semver.Collection(vs))

	return vs[len(vs)-1], nil
}

type versionItem struct {
	title string
	sv    semver.Version
}

func (nv *versionItem) String() string {
	return fmt.Sprintf("%s update (%s)", nv.title, nv.sv.String())
}

func nextVersion(cv *semver.Version) (*semver.Version, error) {
	label := fmt.Sprintf("Current tag is %s. Next is", cv.String())
	items := []*versionItem{
		{"patch", cv.IncPatch()},
		{"minor", cv.IncMinor()},
		{"major", cv.IncMajor()},
	}
	prompt := promptui.Select{
		Label:  label,
		Items:  items,
		Stdout: os.Stderr,
	}
	i, _, err := prompt.Run()
	if err != nil {
		return nil, err
	}
	return &items[i].sv, err
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

	cv, err := currentVersion(pwd)
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
		n, err := nextVersion(cv)
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
