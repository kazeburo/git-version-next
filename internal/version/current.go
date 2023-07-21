package version

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"sort"

	"github.com/Masterminds/semver/v3"
)

func Current(pwd string) (*semver.Version, error) {
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
