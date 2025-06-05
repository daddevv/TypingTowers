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

## Time to Kill (TTK)

A tower can fire one projectile every `tower_fire_rate` seconds. If each projectile deals `tower_damage` and travels at `projectile_speed`, the time to kill a mob with health `H` is approximately:

```txt
shots_needed = ceil(H / tower_damage)
TTK = shots_needed * max(tower_fire_rate, tower_reload_rate)
```

**Note:** The actual time between shots is limited by the slower of the fire rate and reload rate, since the player must type to reload each shot.

### Including Projectile Travel Time

Projectiles take time to reach the mob. If the distance to the mob is `D` and projectile speed is `projectile_speed`, the travel time per shot is:

```txt
travel_time = D / projectile_speed
```

The total time to kill a mob is then:

```txt
TTK_total = TTK + travel_time
```

If the mob is moving toward the base, `D` decreases over time, but for balancing, use the average or initial distance.

## Mob Survival Time (Tsurvive)

Mobs spawn on the right side and move toward the base with speed `mob_speed`. The time for a mob to reach the base is:

```txt
Tsurvive = D / mob_speed
```

To ensure a player can defeat mobs by typing accurately, we require `TTK < Tsurvive` for the first mob of each wave. Since later mobs spawn after `spawn_interval` seconds, we also want the tower to defeat mobs quickly enough that only a manageable number are alive simultaneously.

## Reload Throughput

With an ammo capacity of `C`, a tower can fire `C` shots before running out. Each empty slot generates a reload letter. Assuming perfect accuracy, the fastest time to refill one slot is `tower_reload_rate` seconds. The sustained firing rate considering reload is therefore:

```txt
EffectiveFireRate = max(tower_fire_rate, tower_reload_rate)
```

**Downtime:** After firing `C` shots, the tower cannot fire again until at least one reload letter is typed. If the player is slower than `tower_reload_rate`, the effective fire rate drops further.

## Spawn Rate and Overlapping Mobs

Let `Mwave` be the number of mobs in a wave:

```txt
Mwave = mobs_per_wave_base + (current_wave - 1) * mobs_per_wave_growth
```

Mobs spawn every `spawn_interval` frames. The number of mobs alive at once depends on how quickly mobs are killed and how quickly new ones spawn.

The maximum number of overlapping mobs (`Alive_max`) is:

```txt
Alive_max = ceil(TTK_total / spawn_interval)
```

This assumes the tower is always firing at the closest mob and that the player reloads perfectly.

To keep early waves manageable, we target:

```txt
Alive_max ≤ tower_ammo_capacity
```

so that ammo capacity is sufficient to handle bursts. If this value is too high, increase `spawn_interval` or reduce `Mwave` growth.

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

After all letters are unlocked, waves can grow more challenging by reducing `spawn_interval`, increasing `mob_base_health`, or increasing `mob_speed`. The guiding principle is that `TTK_total` should remain less than `Tsurvive` while letter prompts remain manageable. A gentle exponential increase to health or spawn rate every few waves can provide a long-term challenge without overwhelming the player immediately.

## Summary

Balancing revolves around keeping the ratio of tower damage and firing speed ahead of mob health and spawn rate, while also accounting for projectile travel time, reload throughput, and overlapping spawns. By monitoring `TTK_total`, `Tsurvive`, and reload throughput, we can tune parameters so that any player who types each letter correctly can survive the wave. Difficulty can ramp up slowly through additional mobs per wave and moderate increases to health or speed, ensuring the game remains fun while serving as effective typing practice.

All new mob types (armored, fast, boss) and their special abilities are implemented and covered by unit tests to ensure correct behaviors and stats.
