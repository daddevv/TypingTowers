#!/usr/bin/env python3
"""TypingTowers balance visualization and config editor."""
import argparse
import json
import math
from pathlib import Path

try:
    import matplotlib.pyplot as plt
except ModuleNotFoundError:  # allow running without plotting support
    plt = None

DEFAULT_CONFIG_PATH = Path('v1/config.json')


def load_config(path: Path) -> dict:
    """Load configuration from JSON file."""
    with open(path) as f:
        return json.load(f)


def save_config(cfg: dict, path: Path) -> None:
    """Save configuration back to JSON file."""
    with open(path, 'w') as f:
        json.dump(cfg, f, indent=2)
        f.write('\n')


def apply_overrides(cfg: dict, overrides: dict) -> None:
    """Apply key=value overrides to the configuration."""
    for key, value in overrides.items():
        # attempt numeric conversion if possible
        try:
            if '.' in value:
                value = float(value)
            else:
                value = int(value)
        except ValueError:
            pass
        cfg[key] = value


def simulate(cfg: dict, waves: int = 100) -> dict:
    """Simulate key balancing metrics across multiple waves."""
    results = {
        'wave': [],
        'mob_hp': [],
        'ttk': [],
        'tsurvive': [],
    }
    mob_base_hp = cfg.get('mob_base_health', 1)
    mob_growth = cfg.get('N', 0)
    dmg = cfg.get('tower_damage', 1)
    fire_rate = cfg.get('tower_fire_rate', 1.0)
    reload_rate = cfg.get('tower_reload_rate', 1.0)
    rng = cfg.get('tower_range', 500)
    mob_speed = cfg.get('mob_speed', 1.0)

    eff_rate = max(fire_rate, reload_rate)

    for w in range(1, waves + 1):
        hp = int(float(mob_base_hp) + float(w - 1) * mob_growth)
        shots = max(1, math.ceil(hp / dmg))
        ttk = shots * eff_rate
        tsurvive = rng / mob_speed if mob_speed else float('inf')

        results['wave'].append(w)
        results['mob_hp'].append(hp)
        results['ttk'].append(ttk)
        results['tsurvive'].append(tsurvive)

    return results


def plot_results(results: dict) -> None:
    """Display matplotlib plots for the simulated progression."""
    if plt is None:
        print('matplotlib is required for plotting. Install it with `pip install matplotlib`.')
        return
    waves = results['wave']
    plt.figure(figsize=(10, 6))
    plt.plot(waves, results['mob_hp'], label='Mob HP')
    plt.plot(waves, results['ttk'], label='Time to Kill')
    plt.plot(waves, results['tsurvive'], label='Mob Survival Time')
    plt.xlabel('Wave')
    plt.ylabel('Value')
    plt.title('Progression Simulation')
    plt.legend()
    plt.grid(True)
    plt.tight_layout()
    plt.show()


def main() -> None:
    parser = argparse.ArgumentParser(description='TypingTowers balance editor')
    parser.add_argument('--config', type=Path, default=DEFAULT_CONFIG_PATH, help='Path to config.json')
    parser.add_argument('--set', action='append', default=[], metavar='KEY=VALUE',
                        help='Override configuration value (may be repeated)')
    parser.add_argument('--save', action='store_true', help='Save overrides back to the config file')
    parser.add_argument('--waves', type=int, default=100, help='Number of waves to simulate')
    args = parser.parse_args()

    cfg = load_config(args.config)

    overrides = {}
    for item in args.set:
        if '=' not in item:
            parser.error(f'invalid format for --set {item!r}')
        k, v = item.split('=', 1)
        overrides[k] = v
    if overrides:
        apply_overrides(cfg, overrides)
        if args.save:
            save_config(cfg, args.config)

    results = simulate(cfg, args.waves)
    plot_results(results)


if __name__ == '__main__':
    main()
