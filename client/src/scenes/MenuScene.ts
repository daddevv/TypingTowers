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
        if (!this.worlds || !this.levelManager) {
            this.add.text(400, 300, 'Menu data not loaded. Please restart the game.', { fontSize: '24px', color: '#f00' }).setOrigin(0.5);
            return;
        }
        this.menuItems.forEach(item => item.destroy());
        this.menuItems = [];
        let y = 120;
        this.worlds.forEach((world, wIdx) => {
            const color = wIdx === this.selectedWorld ? '#ff0' : '#fff';
            const txt = this.add.text(400, y, `World ${wIdx + 1}: ${world.name}`, { fontSize: '28px', color, backgroundColor: wIdx === this.selectedWorld ? '#333' : undefined }).setOrigin(0.5).setInteractive({ useHandCursor: true });
            txt.on('pointerdown', () => this.selectWorld(wIdx));
            this.menuItems.push(txt);
            y += 48;
        });
    }

    handleInput(event: KeyboardEvent) {
        if (event.key === 'ArrowDown') {
            if (this.selectedWorld < this.worlds.length - 1) {
                this.selectedWorld++;
                this.renderMenu();
            }
        } else if (event.key === 'ArrowUp') {
            if (this.selectedWorld > 0) {
                this.selectedWorld--;
                this.renderMenu();
            }
        } else if (event.key === 'Enter') {
            this.selectWorld(this.selectedWorld);
        }
    }

    selectWorld(worldIdx: number) {
        this.scene.start('LevelMenuScene', { worldId: this.worlds[worldIdx].id });
    }
}
// Contains AI-generated edits.
