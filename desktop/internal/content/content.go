package content

var (
	LoadedLevels []LevelConfig
	LoadedMobs   []MobConfig
	LoadedWorlds []WorldConfig
)

func LoadContentConfigs() error {
	var err error
	LoadedLevels, err = LoadJSONConfig[LevelConfig]("content/levels.json")
	if err != nil {
		return err
	}
	LoadedMobs, err = LoadJSONConfig[MobConfig]("content/mobs.json")
	if err != nil {
		return err
	}
	LoadedWorlds, err = LoadJSONConfig[WorldConfig]("content/worlds.json")
	if err != nil {
		return err
	}
	return nil
}
