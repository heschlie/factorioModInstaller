package pkg

import (
	"encoding/json"
	"factorioModInstaller/models"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func GetReleaseMeta(name string) *models.Mod {
	url := fmt.Sprintf("https://mods.factorio.com/api/mods/%s", name)

	c := http.Client{Timeout: time.Second * 10}
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	request.Header.Set("User-Agent", "factorioModInstaller")

	resp, err := c.Do(request)
	if err != nil {
		log.WithError(err).WithField("mod", name).Fatal("Failed to get mod")
	}

	defer resp.Body.Close()
	var mod models.Mod
	err = json.NewDecoder(resp.Body).Decode(&mod)
	if err != nil {
		log.WithError(err).WithField("mod", name).Fatal("Failed to parse JSON for mod")
	}

	return &mod
}

func DownloadModRelease(release *models.Release) error {
	user := viper.GetString("username")
	token := viper.GetString("token")
	path := viper.GetString("factorioPath")
	url := fmt.Sprintf("https://mods.factorio.com%s?username=%s&token=%s", release.DownloadURL, user, token)

	out, err := os.Create(filepath.Join(path, "mods", release.FileName))
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
