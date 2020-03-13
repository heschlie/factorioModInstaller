package models

import "time"

type Release struct {
	DownloadURL string `json:"download_url"`
	FileName    string `json:"file_name"`
	InfoJSON    struct {
		FactorioVersion string `json:"factorio_version"`
	} `json:"info_json"`
	ReleasedAt time.Time `json:"released_at"`
	Sha1       string    `json:"sha1"`
	Version    string    `json:"version"`
}
