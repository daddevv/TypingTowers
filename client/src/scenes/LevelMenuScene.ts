// LevelMenuScene.ts
// Scene for selecting a level within a selected world
import Phaser from 'phaser';
import { WORLDS, WorldConfig } from '../curriculum/worldConfig';
import { levelManager } from '../managers/levelManager';

export default class LevelMenuScene extends Phaser.Scene {
    private world!: WorldConfig;
    private menuItems: Phaser.GameObjects.Text[] = [];
    private selectedLevel: number = 0;

    constructor() {
        super({ key: 'LevelMenuScene' });
    }

    init(data: { worldId: number }) {
        this.world = WORLDS.find(w => w.id === data.worldId)!;
    }

    create() {
        this.add.text(400, 40, this.world.name, { fontSize: '36px', color: '#fff' }).setOrigin(0.5);
        this.renderMenu();
        this.input.keyboard?.on('keydown', this.handleInput, this);
        // Re-render menu when scene is resumed (e.g., after completing a level)
        this.events.on('resume', () => {
            this.renderMenu();
        });
        // Refresh menu when scene is woken (e.g., after returning from GameScene)
        this.events.on('wake', () => {
            this.renderMenu();
        });
    }

    renderMenu() {
        this.menuItems.forEach(item => item.destroy());
        this.menuItems = [];
        const levels = this.world.levels;
        levels.forEach((level, idx) => {
            const y = 120 + idx * 48;
            const isUnlocked = levelManager.isLevelUnlocked(level.id);
            const color = isUnlocked ? '#fff' : '#888';
            const item = this.add.text(400, y, `${level.name}${isUnlocked ? '' : ' (Locked)'}`, {
                fontSize: '28px',
                color,
                backgroundColor: idx === this.selectedLevel ? '#444' : undefined,
                padding: { left: 12, right: 12, top: 4, bottom: 4 },
            }).setOrigin(0.5).setInteractive({ useHandCursor: isUnlocked });
            if (isUnlocked) {
                item.on('pointerdown', () => this.selectLevel(idx));
            }
            this.menuItems.push(item);
        });
    }

    handleInput(event: KeyboardEvent) {
        const levels = this.world.levels;
        if (event.key === 'ArrowUp') {
            this.selectedLevel = (this.selectedLevel - 1 + levels.length) % levels.length;
            this.renderMenu();
        } else if (event.key === 'ArrowDown') {
            this.selectedLevel = (this.selectedLevel + 1) % levels.length;
            this.renderMenu();
        } else if (event.key === 'Enter') {
            this.selectLevel(this.selectedLevel);
        } else if (event.key === 'Escape') {
            this.scene.start('MenuScene');
        }
    }

    selectLevel(idx: number) {
        const level = this.world.levels[idx];
        if (!levelManager.isLevelUnlocked(level.id)) return;
        levelManager.setCurrentLevel(this.world.id, level.id);
        this.scene.start('GameScene', { worldId: this.world.id, levelId: level.id });
    }
}
// Contains AI-generated edits.
