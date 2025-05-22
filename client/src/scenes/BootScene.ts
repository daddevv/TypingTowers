// BootScene.ts
// Initializes StateManager, loads essential assets, and sets gameStatus to 'mainMenu'
import Phaser from 'phaser';
import stateManager from '../state/stateManager';
import InputSystem from '../systems/InputSystem';

export default class BootScene extends Phaser.Scene {
    constructor() {
        super({ key: 'BootScene' });
    }

    preload() {
        // Load any global assets here (e.g., loading bar, logo, etc.)
        // this.load.image('logo', 'assets/images/logo.png');
    }

    create() {
        // Defensive: Reset to default state if state is missing/corrupt
        const state = stateManager.getState();
        if (!state || typeof state !== 'object' || !('gameStatus' in state)) {
            // Reset to default state
            stateManager.reset();
        }

        // Register global input listeners (only once)
        InputSystem.getInstance().registerListeners();

        // Listen for gameStatus changes and switch scenes accordingly
        stateManager.on('gameStatusChanged', (status) => {
            if (status === 'mainMenu') {
                this.scene.start('MainMenuScene');
            } else if (status === 'worldSelect') {
                this.scene.start('MenuScene');
            } else if (status === 'levelSelect') {
                this.scene.start('LevelMenuScene');
            } else if (status === 'playing') {
                this.scene.start('GameScene');
            }
            // Add more as needed
        });

        // Always set gameStatus to 'mainMenu' on load for a clean start
        stateManager.setGameStatus('mainMenu');
    }
}
