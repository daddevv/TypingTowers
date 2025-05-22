// BootScene.ts
// Initializes StateManager, loads essential assets, and sets gameStatus to 'mainMenu'
import Phaser from 'phaser';
import stateManager from '../state/stateManager';

export default class BootScene extends Phaser.Scene {
    constructor() {
        super({ key: 'BootScene' });
    }

    preload() {
        // Load any global assets here (e.g., loading bar, logo, etc.)
        // this.load.image('logo', 'assets/images/logo.png');
    }

    create() {
        // Initialize StateManager (already done on import, but can reset if needed)
        // stateManager.reset(); // Uncomment if you want to always start fresh
        // Set gameStatus to 'mainMenu' to trigger main menu
        stateManager.setGameStatus('mainMenu');
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
    }
}
