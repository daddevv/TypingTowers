package game

// Settings holds user configurable options.
type Settings struct {
	Mute bool `json:"mute"`
}

// DefaultSettings returns a Settings struct with defaults.
func DefaultSettings() Settings {
	return Settings{Mute: false}
}
