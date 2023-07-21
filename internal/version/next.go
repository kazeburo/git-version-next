package version

import (
	"fmt"
	"os"

	"github.com/Masterminds/semver/v3"
	"github.com/manifoldco/promptui"
)

type versionItem struct {
	title string
	sv    semver.Version
}

func (nv *versionItem) String() string {
	return fmt.Sprintf("%s update (%s)", nv.title, nv.sv.String())
}

func Next(cv *semver.Version) (*semver.Version, error) {
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
