# Progression Planning

This file contains the parameters used to balance the game, including mobs, towers, and other gameplay elements. The values are defined in `config.json` and can be adjusted to modify the game's difficulty and progression.

Note: Each parameter has a base value and a multiplier that can be applied to scale the difficulty as the game progresses.

## Parameters

### Mobs

- Health: `MobBaseHealth` * `MobHealthMultiplier`
- Speed: `MobBaseSpeed` * `MobSpeedMultiplier`
- Damage: `MobBaseDamage` * `MobDamageMultiplier`
- Spawn Interval: `MobBaseSpawnInterval` * `MobSpawnIntervalMultiplier`
- Mobs per Wave: `MobsPerWaveBase` + (`WaveNumber` - 1) * `MobsPerWaveGrowth`
- Wave Growth: `MobsPerWaveGrowth`
- Healing: `MobBaseHealing` * `MobHealingMultiplier`

### Towers

- Damage: `TowerBaseDamage` * `TowerDamageMultiplier`
- Fire Rate: `TowerBaseFireRate` * `TowerFireRateMultiplier`
- Ammo Capacity: `TowerBaseAmmoCapacity` * `TowerAmmoCapacityMultiplier`
- Projectile Speed: `TowerBaseProjectileSpeed` * `TowerProjectileSpeedMultiplier`
- Range: `TowerBaseRange` * `TowerRangeMultiplier`

### Economy

- Currency Gain: `CurrencyGainBase` * `CurrencyGainMultiplier`
- Upgrade Cost: `UpgradeCostBase` * `UpgradeCostMultiplier`

## Important Metrics

- Shots Needed = `ceil(MobHealth / TowerDamage)`
- Time to Kill (TTK) = Shots Needed * `max(TowerFireRate, UserReloadRate)`
- Mob Survival Time (Tsurvive) = `TowerRange / MobSpeed`
- Effective Mob Wave Size = `MobsPerWaveBase` + (`WaveNumber` - 1) * `MobsPerWaveGrowth`
- Alive Mobs = `ceil(Tsurvive / MobSpawnInterval)`
- Reload Throughput = `TowerAmmoCapacity / UserReloadRate`
- Effective Reload Throughput = `Reload Throughput` * `TowerFireRate`
- Effective TTK = `TTK + (TowerRange / TowerProjectileSpeed)`
- Effective Tsurvive = `MobBaseHealth / TowerDamage` * `max(TowerFireRate, UserReloadRate)`

## Balancing Considerations

- Ensure `TTK < Tsurvive` for the first mob of each wave.
- Ensure the tower can handle the number of mobs alive at once, factoring in the spawn interval and mob speed.
- Adjust parameters to maintain a balance where players can progress through waves without overwhelming difficulty.
- Consider the impact of player skill and accuracy on the game's difficulty.
