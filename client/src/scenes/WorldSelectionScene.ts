// MenuScene.ts
// World and Level selection menu with lock/unlock and local storage persistence
import Phaser from 'phaser';
import { WorldConfig, WORLDS } from '../curriculum/worldConfig';
import { levelManager } from '../managers/levelManager';
import stateManager from '../state/stateManager';

export default class WorldSelectionScene extends Phaser.Scene {
    private worlds: WorldConfig[] = [];
    private levelManager = levelManager;
    private menuItems: Phaser.GameObjects.Text[] = [];
    private selectedWorld: number = 0;
    private onGameStatusChanged?: (status: string) => void;

    constructor() {
        super({ key: 'MenuScene' });
    }

    init(data?: { worlds?: WorldConfig[], levelManager?: typeof levelManager }) {
        // Always load worlds from WORLDS if not provided
        this.worlds = (data && data.worlds) ? data.worlds : WORLDS;
        this.levelManager = (data && data.levelManager) || levelManager;
    }

    preload() { }

    create() {
        // Remove all children to clear any lingering error messages or UI
        this.children.removeAll();
        this.add.text(400, 40, 'TypeDefense', { fontSize: '40px', color: '#fff' }).setOrigin(0.5);
        this.renderMenu();
        // Add Back button
        const backButton = this.add.text(400, 500, 'Back (Esc)', {
            fontSize: '24px', color: '#fff', backgroundColor: '#333', padding: { left: 24, right: 24, top: 8, bottom: 8 }
        }).setOrigin(0.5).setInteractive({ useHandCursor: true });
        backButton.on('pointerdown', () => {
            this.scene.stop();
            // Always go to main menu
            this.scene.start('MainMenuScene');
        });
        // Listen for Escape key
        this.input.keyboard?.on('keydown', (event: KeyboardEvent) => {
            if (event.key === 'Escape') {
                this.scene.stop();
                this.scene.start('MainMenuScene');
            }
        });
        // Input is now handled via InputSystem and stateManager
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
        // Defensive: Check for valid world data
        if (!this.worlds || !Array.isArray(this.worlds) || this.worlds.length === 0) {
            this.add.text(400, 300, 'No worlds available. Please check game data.', { fontSize: '24px', color: '#f00' }).setOrigin(0.5);
            return;
        }
        if (!this.levelManager) {
            this.add.text(400, 340, 'Level manager not loaded.', { fontSize: '20px', color: '#f00' }).setOrigin(0.5);
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
        // Store selected world in stateManager.level
        const state = stateManager.getState();
        state.level.currentWorld = world.id;
        stateManager.updateProgression({}); // trigger save
        // Optionally, store selected world in stateManager if needed
        // For robust navigation, update current world in stateManager (if you have such a method)
        // stateManager.setCurrentWorld(world.id); // Uncomment if implemented
        stateManager.setGameStatus('levelSelect');
    }
}
