# Progression Planning

This file contains the parameters used to balance the game, including mobs, towers, and other gameplay elements. The values are defined in `config.json` and can be adjusted to modify the game's difficulty and progression.

Note: Each parameter has a base value and a multiplier that can be applied to scale the difficulty as the game progresses. All parameters listed here are either currently supported or are candidates for future upgrades/configuration.

## Parameters

### Mobs

- Health: `MobBaseHealth` * `MobHealthMultiplier`
- Speed: `MobBaseSpeed` * `MobSpeedMultiplier`
- Damage: `MobBaseDamage` * `MobDamageMultiplier`
- Armor: `MobBaseArmor` * `MobArmorMultiplier`
- Magic Resist: `MobBaseMagicResist` * `MobMagicResistMultiplier`
- Healing: `MobBaseHealing` * `MobHealingMultiplier`
- Regeneration: `MobBaseRegen` * `MobRegenMultiplier`
- Evasion: `MobBaseEvasion` * `MobEvasionMultiplier`
- Special Abilities: (e.g., shield, split, spawn, flying, etc.)
- Spawn Interval: `MobBaseSpawnInterval` * `MobSpawnIntervalMultiplier`
- Mobs per Wave: `MobsPerWaveBase` + (`WaveNumber` - 1) * `MobsPerWaveGrowth`
- Wave Growth: `MobsPerWaveGrowth`
- Reward: `MobBaseReward` * `MobRewardMultiplier`
- Size: `MobBaseSize` * `MobSizeMultiplier`
- Status Resistances: (e.g., slow, stun, burn resistances)

### Towers

- Damage: `TowerBaseDamage` * `TowerDamageMultiplier`
- Fire Rate: `TowerBaseFireRate` * `TowerFireRateMultiplier`
- Ammo Capacity: `TowerBaseAmmoCapacity` * `TowerAmmoCapacityMultiplier`
- Projectile Speed: `TowerBaseProjectileSpeed` * `TowerProjectileSpeedMultiplier`
- Range: `TowerBaseRange` * `TowerRangeMultiplier`
- Projectiles per Shot: `TowerBaseProjectilesPerShot` * `TowerProjectilesPerShotMultiplier`
- Bounce Count: `TowerBaseBounceCount` * `TowerBounceCountMultiplier`
- Splash Radius: `TowerBaseSplashRadius` * `TowerSplashRadiusMultiplier`
- Splash Damage: `TowerBaseSplashDamage` * `TowerSplashDamageMultiplier`
- Critical Hit Chance: `TowerBaseCritChance` * `TowerCritChanceMultiplier`
- Critical Hit Multiplier: `TowerBaseCritMultiplier` * `TowerCritMultiplierMultiplier`
- Status Effects: (e.g., slow, stun, burn, poison chance/duration/strength)
- Reload Rate: `TowerBaseReloadRate` * `TowerReloadRateMultiplier`
- Reload Prompt Complexity: (e.g., letter pool size, word length)
- Accuracy Bonus: `TowerBaseAccuracyBonus` * `TowerAccuracyBonusMultiplier`
- Armor Penetration: `TowerBaseArmorPen` * `TowerArmorPenMultiplier`
- Magic Penetration: `TowerBaseMagicPen` * `TowerMagicPenMultiplier`
- Targeting Priority: (e.g., first, last, strongest, weakest, random)
- Upgrade Slots: `TowerBaseUpgradeSlots` * `TowerUpgradeSlotsMultiplier`
- Special Abilities: (e.g., chain lightning, freeze, multi-shot, etc.)

### Base

- Health: `BaseHealth`
- Regeneration: `BaseRegen`
- Armor: `BaseArmor`
- Magic Resist: `BaseMagicResist`
- Shield: `BaseShield`
- Passive Gold Gain: `BaseGoldPerWave`
- Special Abilities: (e.g., retaliate, heal towers, etc.)

### Economy

- Currency Gain: `CurrencyGainBase` * `CurrencyGainMultiplier`
- Upgrade Cost: `UpgradeCostBase` * `UpgradeCostMultiplier`
- Shop Discount: `ShopDiscount`
- Wave Bonus: `WaveBonusGold`
- Interest Rate: `GoldInterestRate`
- Sell Refund Rate: `SellRefundRate`

### Game/Progression

- Wave Timer: `WaveTimerBase` * `WaveTimerMultiplier`
- Endless Scaling Rate: `EndlessScalingMultiplier`
- Unlockable Letters/Words: `LetterUnlockRate`
- Difficulty Scaling: `DifficultyScalingMultiplier`
- Achievement Bonuses: (e.g., per X waves, per Y kills)
- Random Event Frequency: `RandomEventRate`
- Powerup Frequency: `PowerupSpawnRate`
- Powerup Effectiveness: `PowerupEffectMultiplier`

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
- Plan for future upgrades and mechanics by reserving config fields and balancing formulas for new features.
