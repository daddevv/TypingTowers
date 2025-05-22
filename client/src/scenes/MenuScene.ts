// MenuScene.ts
// World and Level selection menu with lock/unlock and local storage persistence
import Phaser from 'phaser';
import { WorldConfig } from '../curriculum/worldConfig';
import { levelManager } from '../managers/levelManager';
import stateManager from '../state/stateManager';

export default class MenuScene extends Phaser.Scene {
    private worlds: WorldConfig[] = [];
    private levelManager = levelManager;
    private menuItems: Phaser.GameObjects.Text[] = [];
    private selectedWorld: number = 0;
    private onGameStatusChanged?: (status: string) => void;

    constructor() {
        super({ key: 'MenuScene' });
    }

    init(data: { worlds: WorldConfig[], levelManager?: typeof levelManager }) {
        this.worlds = data.worlds;
        this.levelManager = data.levelManager || levelManager;
    }

    preload() { }

    create() {
        this.add.text(400, 40, 'TypeDefense', { fontSize: '40px', color: '#fff' }).setOrigin(0.5);
        this.renderMenu();
        this.input.keyboard?.on('keydown', this.handleInput, this);
        // Listen for gameStatus changes and transition if needed
        this.onGameStatusChanged = (status: string) => {
            if (status !== 'worldSelect') {
                this.scene.stop();
            }
        };
        stateManager.on('gameStatusChanged', this.onGameStatusChanged);
        this.events.once('shutdown', () => {
            if (this.onGameStatusChanged) {
                stateManager.off('gameStatusChanged', this.onGameStatusChanged);
            }
        });
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
            // Only unlock world 1 by default; others require previous world completion
            let isUnlocked = true;
            if (wIdx > 0) {
                // Get last level of previous world
                const prevWorld = this.worlds[wIdx - 1];
                const lastLevel = prevWorld.levels[prevWorld.levels.length - 1];
                const progress = this.levelManager.getLevelProgress(lastLevel.id);
                isUnlocked = !!(progress && progress.completed);
            }
            const color = wIdx === this.selectedWorld ? (isUnlocked ? '#ff0' : '#888') : (isUnlocked ? '#fff' : '#888');
            const label = `World ${wIdx + 1}: ${world.name}${isUnlocked ? '' : ' (Locked)'}`;
            const txt = this.add.text(400, y, label, { fontSize: '28px', color, backgroundColor: wIdx === this.selectedWorld ? '#333' : undefined }).setOrigin(0.5).setInteractive({ useHandCursor: isUnlocked });
            if (isUnlocked) {
                txt.on('pointerdown', () => this.selectWorld(wIdx));
            }
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

    // When starting LevelMenuScene, always pass the singleton levelManager
    selectWorld(idx: number) {
        const world = this.worlds[idx];
        // Instead of scene.start, update state
        // Optionally, store selected world in stateManager if needed
        stateManager.setGameStatus('levelSelect');
        // ...could also update current world in stateManager here...
    }
}
