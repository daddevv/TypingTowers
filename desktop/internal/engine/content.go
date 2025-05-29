package engine

import (
	"td/internal/content"
)

var (
	LoadedLevels []content.LevelConfig
	LoadedMobs   []content.MobConfig
	LoadedWorlds []content.WorldConfig
)

func LoadAllContent() error {
	var err error
	LoadedLevels, err = LoadJSONConfig[content.LevelConfig]("content/levels.json")
	if err != nil {
		return err
	}
	LoadedMobs, err = LoadJSONConfig[content.MobConfig]("content/mobs.json")
	if err != nil {
		return err
	}
	LoadedWorlds, err = LoadJSONConfig[content.WorldConfig]("content/worlds.json")
	if err != nil {
		return err
	}
	return nil
}
