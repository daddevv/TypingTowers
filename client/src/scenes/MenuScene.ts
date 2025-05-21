// MenuScene.ts
// World and Level selection menu with lock/unlock and local storage persistence
import Phaser from 'phaser';
import { WorldConfig } from '../curriculum/worldConfig';
import LevelManager from '../managers/levelManager';

export default class MenuScene extends Phaser.Scene {
    private worlds: WorldConfig[] = [];
    private levelManager!: LevelManager;
    private menuItems: Phaser.GameObjects.Text[] = [];
    private selectedWorld: number = 0;
    private selectedLevel: number = 0;

    constructor() {
        super({ key: 'MenuScene' });
    }

    init(data: { worlds: WorldConfig[], levelManager: LevelManager }) {
        this.worlds = data.worlds;
        this.levelManager = data.levelManager;
    }

    preload() { }

    create() {
        this.add.text(400, 40, 'TypeDefense', { fontSize: '40px', color: '#fff' }).setOrigin(0.5);
        this.renderMenu();
        this.input.keyboard?.on('keydown', this.handleInput, this);
    }

    renderMenu() {
        this.menuItems.forEach(item => item.destroy());
        this.menuItems = [];
        let y = 120;
        this.worlds.forEach((world, wIdx) => {
            this.add.text(400, y, `World ${wIdx + 1}: ${world.name}`, { fontSize: '28px', color: '#ff0' }).setOrigin(0.5);
            y += 36;
            world.levels.forEach((level, lIdx) => {
                const levelKey = `${wIdx + 1}-${lIdx + 1}`;
                const unlocked = this.levelManager.isLevelUnlocked(levelKey);
                const completed = this.levelManager.isLevelCompleted(levelKey);
                let label = `  Level ${lIdx + 1}: ${level.name}`;
                if (!unlocked) label += ' (Locked)';
                if (completed) label += ' ✓';
                const color = unlocked ? (completed ? '#0f0' : '#fff') : '#888';
                const txt = this.add.text(400, y, label, { fontSize: '22px', color }).setOrigin(0.5);
                if (wIdx === this.selectedWorld && lIdx === this.selectedLevel) txt.setStyle({ backgroundColor: '#333' });
                this.menuItems.push(txt);
                y += 28;
            });
            y += 10;
        });
    }

    handleInput(event: KeyboardEvent) {
        const world = this.worlds[this.selectedWorld];
        if (event.key === 'ArrowDown') {
            if (this.selectedLevel < world.levels.length - 1) {
                this.selectedLevel++;
            } else if (this.selectedWorld < this.worlds.length - 1) {
                this.selectedWorld++;
                this.selectedLevel = 0;
            }
            this.renderMenu();
        } else if (event.key === 'ArrowUp') {
            if (this.selectedLevel > 0) {
                this.selectedLevel--;
            } else if (this.selectedWorld > 0) {
                this.selectedWorld--;
                this.selectedLevel = this.worlds[this.selectedWorld].levels.length - 1;
            }
            this.renderMenu();
        } else if (event.key === 'Enter') {
            const levelKey = `${this.selectedWorld + 1}-${this.selectedLevel + 1}`;
            if (this.levelManager.isLevelUnlocked(levelKey)) {
                this.scene.start('GameScene', { world: this.selectedWorld, level: this.selectedLevel });
            }
        }
    }
}
// Contains AI-generated edits.
