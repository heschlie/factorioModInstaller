package pkg

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
)

var pattern = `^Version: (\d+)\.(\d+)\.(\d+) \(.+\)`

type FactorioVersion struct {
	Major int
	Minor int
	Point int
}

func (v *FactorioVersion) String() string{
	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Point)
}

func GetFactorioVersion() (FactorioVersion, error) {
	r := regexp.MustCompile(pattern)

	path := viper.GetString("factorioPath")
	cmd := exec.Command(filepath.Join(path, "bin", "x64", "factorio"), "--version")
	out, err := cmd.Output()
	if err != nil {
		log.WithError(err).Error("Failed to find Factorio version")
		return FactorioVersion{}, err
	}

	result := r.FindSubmatch(out)

	major, _ :=strconv.Atoi(string(result[1]))
	minor, _ :=strconv.Atoi(string(result[2]))
	point, _ :=strconv.Atoi(string(result[3]))
	f := FactorioVersion{
		Major: major,
		Minor: minor,
		Point: point,
	}

	return f, nil
}
