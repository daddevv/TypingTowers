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

import tkinter as tk
from tkinter import ttk
from tkinter import messagebox
import threading

DEFAULT_CONFIG_PATH = Path("v1/config.json")


def load_config(path: Path) -> dict:
    """Load configuration from JSON file."""
    with open(path) as f:
        return json.load(f)


def save_config(cfg: dict, path: Path) -> None:
    """Save configuration back to JSON file."""
    with open(path, "w") as f:
        json.dump(cfg, f, indent=2)
        f.write("\n")


def apply_overrides(cfg: dict, overrides: dict) -> None:
    """Apply key=value overrides to the configuration."""
    for key, value in overrides.items():
        # attempt numeric conversion if possible
        try:
            if "." in value:
                value = float(value)
            else:
                value = int(value)
        except ValueError:
            pass
        cfg[key] = value


def simulate(cfg: dict, waves: int = 100) -> dict:
    """Simulate key balancing metrics across multiple waves, using more parameters."""
    results = {
        "wave": [],
        "enemy_hp": [],
        "enemy_armor": [],
        "enemy_speed": [],
        "ttk": [],
        "tsurvive": [],
        "alive_mobs": [],
        "tower_damage": [],
        "tower_fire_rate": [],
        "tower_range": [],
        "tower_splash": [],
        "tower_crit": [],
    }
    enemy_base_hp = cfg.get("enemy_base_health", 1)
    enemy_growth = cfg.get("enemy_health_multiplier", 0)
    enemy_armor = cfg.get("enemy_base_armor", 0)
    enemy_armor_mult = cfg.get("enemy_armor_multiplier", 0)
    enemy_speed = cfg.get("enemy_base_speed", 1.0)
    enemy_speed_mult = cfg.get("enemy_speed_multiplier", 0)
    dmg = cfg.get("tower_base_damage", 1)
    dmg_mult = cfg.get("tower_damage_multiplier", 1)
    fire_rate = cfg.get("tower_base_fire_rate", 1.0)
    fire_rate_mult = cfg.get("tower_fire_rate_multiplier", 1.0)
    rng = cfg.get("tower_base_range", 500)
    rng_mult = cfg.get("tower_range_multiplier", 1.0)
    splash = cfg.get("tower_base_splash_radius", 0)
    splash_mult = cfg.get("tower_splash_radius_multiplier", 1.0)
    crit = cfg.get("tower_base_crit_chance", 0)
    crit_mult = cfg.get("tower_crit_chance_multiplier", 1.0)
    reload_rate = cfg.get("tower_base_reload_rate", 1.0)
    reload_rate_mult = cfg.get("tower_reload_rate_multiplier", 1.0)
    ammo = cfg.get("tower_base_ammo_capacity", 5)
    spawn_interval = cfg.get("enemy_base_spawn_interval", 1.0)
    spawn_interval_mult = cfg.get("enemy_spawn_interval_multiplier", 1.0)
    mobs_per_wave_base = cfg.get("enemies_per_wave_base", 3)
    mobs_per_wave_growth = cfg.get("enemies_per_wave_growth", 1)

    for w in range(1, waves + 1):
        hp = float(enemy_base_hp) * (1 + (w - 1) * enemy_growth)
        armor = float(enemy_armor) * (1 + (w - 1) * enemy_armor_mult)
        speed = float(enemy_speed) * (1 + (w - 1) * enemy_speed_mult)
        tdmg = float(dmg) * dmg_mult
        tfr = float(fire_rate) * fire_rate_mult
        trng = float(rng) * rng_mult
        tsplash = float(splash) * splash_mult
        tcrit = float(crit) * crit_mult
        treload = float(reload_rate) * reload_rate_mult
        eff_rate = max(tfr, treload)
        shots = max(1, math.ceil((hp - armor) / tdmg))
        ttk = shots * eff_rate
        tsurvive = trng / speed if speed else float("inf")
        mobs_per_wave = mobs_per_wave_base + (w - 1) * mobs_per_wave_growth
        spawn = spawn_interval * spawn_interval_mult
        alive = math.ceil(ttk / spawn) if spawn else mobs_per_wave
        results["wave"].append(w)
        results["enemy_hp"].append(hp)
        results["enemy_armor"].append(armor)
        results["enemy_speed"].append(speed)
        results["ttk"].append(ttk)
        results["tsurvive"].append(tsurvive)
        results["alive_mobs"].append(alive)
        results["tower_damage"].append(tdmg)
        results["tower_fire_rate"].append(tfr)
        results["tower_range"].append(trng)
        results["tower_splash"].append(tsplash)
        results["tower_crit"].append(tcrit)
    return results


def plot_results(results: dict, ax=None, canvas=None) -> None:
    """Display matplotlib plots for the simulated progression. Optionally update an existing canvas."""
    if plt is None:
        print(
            "matplotlib is required for plotting. Install it with `pip install matplotlib`."
        )
        return
    waves = results["wave"]
    if ax is None:
        plt.figure(figsize=(10, 6))
        ax = plt.gca()
    ax.clear()
    ax.plot(waves, results["enemy_hp"], label="Enemy HP")
    ax.plot(waves, results["enemy_armor"], label="Enemy Armor")
    ax.plot(waves, results["ttk"], label="Time to Kill")
    ax.plot(waves, results["tsurvive"], label="Mob Survival Time")
    ax.plot(waves, results["alive_mobs"], label="Alive Mobs")
    ax.plot(waves, results["tower_damage"], label="Tower Damage")
    ax.plot(waves, results["tower_fire_rate"], label="Tower Fire Rate")
    ax.plot(waves, results["tower_range"], label="Tower Range")
    ax.plot(waves, results["tower_splash"], label="Tower Splash")
    ax.plot(waves, results["tower_crit"], label="Tower Crit Chance")
    ax.set_xlabel("Wave")
    ax.set_ylabel("Value")
    ax.set_title("Progression Simulation")
    ax.legend()
    ax.grid(True)
    if canvas:
        canvas.draw()
    else:
        plt.tight_layout()
        plt.show()


def run_gui():
    """Launch a Tkinter GUI for live parameter editing and plotting."""
    import matplotlib

    matplotlib.use("TkAgg")
    from matplotlib.backends.backend_tkagg import FigureCanvasTkAgg
    import matplotlib.pyplot as plt

    param_defs = [
        ("enemy_base_health", 1, 1000, float),
        ("enemy_health_multiplier", 1, 5, float),
        ("enemy_base_armor", 0, 100, float),
        ("enemy_armor_multiplier", 1, 5, float),
        ("enemy_base_speed", 0.01, 1, float),
        ("enemy_speed_multiplier", 1, 5, float),
        ("tower_base_damage", 1, 100, float),
        ("tower_damage_multiplier", 1, 5, float),
        ("tower_base_fire_rate", 0.1, 5, float),
        ("tower_fire_rate_multiplier", 1, 5, float),
        ("tower_base_range", 100, 2000, float),
        ("tower_range_multiplier", 1, 5, float),
        ("tower_base_splash_radius", 0, 500, float),
        ("tower_splash_radius_multiplier", 1, 5, float),
        ("tower_base_crit_chance", 0, 1, float),
        ("tower_crit_chance_multiplier", 1, 5, float),
        ("tower_base_reload_rate", 0.1, 5, float),
        ("tower_reload_rate_multiplier", 1, 5, float),
        ("tower_base_ammo_capacity", 1, 20, int),
        ("enemy_base_spawn_interval", 0.1, 5, float),
        ("enemy_spawn_interval_multiplier", 1, 5, float),
        ("enemies_per_wave_base", 1, 20, int),
        ("enemies_per_wave_growth", 0, 10, int),
    ]

    cfg = load_config(DEFAULT_CONFIG_PATH)
    root = tk.Tk()
    root.title("TypingTowers Balance Editor")

    # --- Filterable metrics/controls ---
    filter_options = {
        "All": [p[0] for p in param_defs],
        "TTK < Tsurvive": [
            "enemy_base_health",
            "enemy_health_multiplier",
            "tower_base_damage",
            "tower_damage_multiplier",
            "tower_base_fire_rate",
            "tower_fire_rate_multiplier",
            "tower_base_reload_rate",
            "tower_reload_rate_multiplier",
            "tower_base_range",
            "tower_range_multiplier",
            "enemy_base_speed",
            "enemy_speed_multiplier",
        ],
        "Alive Mobs <= Ammo": [
            "tower_base_ammo_capacity",
            "enemy_base_spawn_interval",
            "enemy_spawn_interval_multiplier",
            "enemies_per_wave_base",
            "enemies_per_wave_growth",
            "enemy_base_health",
            "enemy_health_multiplier",
            "tower_base_damage",
            "tower_damage_multiplier",
            "tower_base_fire_rate",
            "tower_fire_rate_multiplier",
            "tower_base_reload_rate",
            "tower_reload_rate_multiplier",
        ],
    }
    filter_var = tk.StringVar(value="All")

    def update_filter(*args):
        selected = filter_var.get()
        visible = set(filter_options[selected])
        for i, (key, *_rest) in enumerate(param_defs):
            widgets = param_widgets[key]
            if key in visible:
                for w in widgets:
                    w.grid()
            else:
                for w in widgets:
                    w.grid_remove()
        update_plot()

    # --- Controls ---
    param_vars = {}
    param_widgets = {}
    for i, (key, vmin, vmax, typ) in enumerate(param_defs):
        label = tk.Label(root, text=key)
        var = (
            tk.DoubleVar(value=cfg.get(key, vmin))
            if typ is float
            else tk.IntVar(value=cfg.get(key, vmin))
        )
        param_vars[key] = var
        scale = tk.Scale(
            root,
            variable=var,
            from_=vmin,
            to=vmax,
            resolution=0.01 if typ is float else 1,
            orient="horizontal",
            length=200,
        )
        param_widgets[key] = (label, scale)
        label.grid(row=i, column=0, sticky="w")
        scale.grid(row=i, column=1)

    # --- Filter dropdown ---
    filter_label = tk.Label(root, text="Show Controls For:")
    filter_label.grid(row=0, column=3, sticky="w")
    filter_menu = tk.OptionMenu(
        root, filter_var, *filter_options.keys(), command=update_filter
    )
    filter_menu.grid(row=1, column=3, sticky="w")

    # --- Save/Load Buttons ---
    def save_config_gui():
        for key, var in param_vars.items():
            cfg[key] = var.get()
        save_config(cfg, DEFAULT_CONFIG_PATH)
        messagebox.showinfo("Saved", "Configuration saved.")

    def load_config_gui():
        loaded = load_config(DEFAULT_CONFIG_PATH)
        for key, var in param_vars.items():
            if key in loaded:
                var.set(loaded[key])
        update_plot()

    save_btn = tk.Button(root, text="Save", command=save_config_gui)
    save_btn.grid(row=len(param_defs) + 2, column=0, sticky="w")
    load_btn = tk.Button(root, text="Reload", command=load_config_gui)
    load_btn.grid(row=len(param_defs) + 2, column=1, sticky="w")

    # --- Clean close ---
    def on_close():
        root.quit()
        root.destroy()

    root.protocol("WM_DELETE_WINDOW", on_close)

    fig, ax = plt.subplots(figsize=(8, 5))
    canvas = FigureCanvasTkAgg(fig, master=root)
    canvas.get_tk_widget().grid(row=0, column=4, rowspan=len(param_defs) + 5)

    def update_plot(*args):
        for key, var in param_vars.items():
            cfg[key] = var.get()
        results = simulate(cfg, 100)
        # Check for warnings
        warnings = []
        # TTK > Tsurvive check
        for w, ttk, tsurvive in zip(
            results["wave"], results["ttk"], results["tsurvive"]
        ):
            if ttk > tsurvive:
                warnings.append(f"Wave {w}: TTK > Tsurvive! (Mob may reach base)")
                break
        # Alive mobs > ammo capacity check
        ammo_capacity = cfg.get("tower_base_ammo_capacity", 5)
        for w, alive in zip(results["wave"], results["alive_mobs"]):
            if alive > ammo_capacity:
                warnings.append(
                    f"Wave {w}: Alive mobs ({alive}) > Ammo Capacity ({ammo_capacity})!"
                )
                break
        warnings.append("Note: Actual difficulty depends on player skill and accuracy.")
        if warnings:
            warning_label.config(text="\n".join(warnings), fg="red")
        else:
            warning_label.config(text="No critical issues detected.", fg="green")
        # Only plot relevant metrics if filtered
        selected = filter_var.get()
        plot_keys = {
            "All": [
                "enemy_hp",
                "enemy_armor",
                "ttk",
                "tsurvive",
                "alive_mobs",
                "tower_damage",
                "tower_fire_rate",
                "tower_range",
                "tower_splash",
                "tower_crit",
            ],
            "TTK < Tsurvive": [
                "enemy_hp",
                "ttk",
                "tsurvive",
                "tower_damage",
                "tower_fire_rate",
                "tower_range",
                "enemy_speed",
            ],
            "Alive Mobs <= Ammo": [
                "alive_mobs",
                "tower_base_ammo_capacity",
                "enemy_hp",
                "ttk",
            ],
        }
        keys = plot_keys.get(selected, plot_keys["All"])
        ax.clear()
        waves = results["wave"]
        for k in keys:
            if k in results:
                ax.plot(waves, results[k], label=k)
            elif k in cfg:
                ax.plot(waves, [cfg[k]] * len(waves), label=k)
        ax.set_xlabel("Wave")
        ax.set_ylabel("Value")
        ax.set_title("Progression Simulation")
        ax.legend()
        ax.grid(True)
        if canvas:
            canvas.draw()
        else:
            plt.tight_layout()
            plt.show()

    # Add warning label to GUI
    warning_label = tk.Label(
        root, text="", fg="red", font=("Arial", 10, "bold"), justify="left"
    )
    warning_label.grid(row=len(param_defs), column=0, columnspan=4, sticky="w")

    # Reserve config fields for future upgrades/mechanics
    reserved_fields = [
        "future_upgrade_1",
        "future_upgrade_2",
        "future_mechanic_1",
        "future_mechanic_2",
    ]
    for field in reserved_fields:
        if field not in cfg:
            cfg[field] = 0

    for var in param_vars.values():
        var.trace_add("write", update_plot)

    update_filter()
    root.mainloop()


def main() -> None:
    parser = argparse.ArgumentParser(description="TypingTowers balance editor")
    parser.add_argument(
        "--config", type=Path, default=DEFAULT_CONFIG_PATH, help="Path to config.json"
    )
    parser.add_argument(
        "--set",
        action="append",
        default=[],
        metavar="KEY=VALUE",
        help="Override configuration value (may be repeated)",
    )
    parser.add_argument(
        "--save", action="store_true", help="Save overrides back to the config file"
    )
    parser.add_argument(
        "--waves", type=int, default=100, help="Number of waves to simulate"
    )
    parser.add_argument(
        "--gui", action="store_true", help="Launch GUI for live editing"
    )
    args = parser.parse_args()

    if args.gui:
        run_gui()
        return

    cfg = load_config(args.config)

    overrides = {}
    for item in args.set:
        if "=" not in item:
            parser.error(f"invalid format for --set {item!r}")
        k, v = item.split("=", 1)
        overrides[k] = v
    if overrides:
        apply_overrides(cfg, overrides)
        if args.save:
            save_config(cfg, args.config)

    results = simulate(cfg, args.waves)
    plot_results(results)


if __name__ == "__main__":
    main()
