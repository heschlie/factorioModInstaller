package models

type Mod struct {
	Category       string `json:"category"`
	DownloadsCount int    `json:"downloads_count"`
	Name           string `json:"name"`
	Owner          string `json:"owner"`
	Releases       []Release `json:"releases"`
	Score     float64 `json:"score"`
	Summary   string  `json:"summary"`
	Thumbnail string  `json:"thumbnail"`
	Title     string  `json:"title"`
}

type ReleaseNotFound struct {
	err string
}

func (e *ReleaseNotFound) Error() string {
	return "No release found"
}

func (m *Mod) GetLatestReleaseByVersion(ver string) (*Release, error) {
	if len(m.Releases) == 0 {
		return nil, &ReleaseNotFound{}
	}
	var latest Release
	for _, r := range m.Releases {
		if ver == r.InfoJSON.FactorioVersion && r.ReleasedAt.After(latest.ReleasedAt) {
			latest = r
		}
	}

	return &latest, nil
}
