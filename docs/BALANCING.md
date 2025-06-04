# Balancing Approach

This document outlines a mathematical approach for balancing TypingTowers so players can reliably progress through waves by keeping up with the reload prompts. The goal is for players who maintain accurate typing to survive indefinitely during the main progression phase. Once all letters are unlocked, the difficulty may scale faster for an optional endless challenge.

## Parameters

The following configuration parameters impact combat balance:

- `tower_damage` – damage dealt per projectile
- `tower_fire_rate` – seconds between shots
- `tower_reload_rate` – delay after a correct letter before another letter can be typed
- `tower_ammo_capacity` – number of shots before reloading is required
- `projectile_speed` – speed of projectiles
- `mob_base_health` – starting health of each mob
- `mob_speed` – movement speed of mobs
- `mobs_per_wave_base` – mobs spawned on wave 1
- `mobs_per_wave_growth` – additional mobs per subsequent wave
- `spawn_interval` – seconds between spawns within a wave

All values are defined in `config.json` and can be reloaded at runtime.

## Time to Kill

A tower can fire one projectile every `tower_fire_rate` seconds. If each projectile deals `tower_damage` and travels at `projectile_speed`, the time to kill a mob with health `H` is approximately:

```
TTK = ceil(H / tower_damage) * tower_fire_rate
```

The travel time from tower to mob also matters, but because projectiles are fast relative to mob speed, it is typically small. For early balancing, this factor can be ignored or approximated as `distance / projectile_speed`.

## Mob Survival Time

Mobs spawn on the right side and move toward the base with speed `mob_speed`. Let `D` be the distance from spawn to base. The time for a mob to reach the base is approximately:

```
Tsurvive = D / mob_speed
```

To ensure a player can defeat mobs by typing accurately, we require `TTK < Tsurvive` for the first mob of each wave. Since later mobs spawn after `spawn_interval` seconds, we also want the tower to defeat mobs quickly enough that only a manageable number are alive simultaneously.

## Reload Throughput

With an ammo capacity of `C`, a tower can fire `C` shots before running out. Each empty slot generates a reload letter. Assuming perfect accuracy, the fastest time to refill one slot is `tower_reload_rate` seconds. The sustained firing rate considering reload is therefore:

```
EffectiveFireRate = max(tower_fire_rate, tower_reload_rate)
```

For balance, the player must be able to maintain this rate by typing each letter correctly when it appears.

## Spawn Rate Limits

Let `Mwave` be the number of mobs in a wave:

```
Mwave = mobs_per_wave_base + (current_wave - 1) * mobs_per_wave_growth
```

We choose `spawn_interval` such that the tower can kill mobs faster than they appear when the player keeps up with reloads. Given `TTK` and `spawn_interval`, the number of mobs alive at once is roughly:

```
Alive ≈ TTK / spawn_interval
```

To keep early waves manageable, we target `Alive ≤ C` so that ammo capacity is sufficient to handle bursts. If this value is too high, increase `spawn_interval` or reduce `Mwave` growth.

## Example Baseline

Using the default configuration:

- `tower_damage = 1`
- `tower_fire_rate = 100`
- `tower_reload_rate = 100`
- `tower_ammo_capacity = 5`
- `projectile_speed = 5`
- `mob_base_health = 1`
- `mob_speed = 1`
- `mobs_per_wave_base = 3`
- `mobs_per_wave_growth = 3`
- `spawn_interval = 1`

If the distance to the base is ~800 pixels, `Tsurvive ≈ 13s`. A mob with `H=1` takes about `TTK = 1.6s`, so `TTK < Tsurvive` is satisfied. The tower can kill about one mob every `1.6s`. With the chosen spawn interval of one second, at most two mobs are alive before the first is destroyed, which is less than the ammo capacity of five. This keeps the game winnable when reloading correctly.

## Endless Scaling

After all letters are unlocked, waves can grow more challenging by reducing `spawn_interval`, increasing `mob_base_health`, or increasing `mob_speed`. The guiding principle is that `TTK` should remain less than `Tsurvive` while letter prompts remain manageable. A gentle exponential increase to health or spawn rate every few waves can provide a long-term challenge without overwhelming the player immediately.

## Summary

Balancing revolves around keeping the ratio of tower damage and firing speed ahead of mob health and spawn rate. By monitoring `TTK`, `Tsurvive`, and reload throughput, we can tune parameters so that any player who types each letter correctly can survive the wave. Difficulty can ramp up slowly through additional mobs per wave and moderate increases to health or speed, ensuring the game remains fun while serving as effective typing practice.
