// BootScene.ts
// Initializes StateManager, loads essential assets, and sets gameStatus to 'mainMenu'
import Phaser from 'phaser';
import stateManager from '../state/stateManager';
import InputSystem from '../systems/InputSystem';
import { IRenderManager } from '../render/RenderManager';

export default class BootScene extends Phaser.Scene {
    constructor() {
        super({ key: 'BootScene' });
    }

    preload() {
        // Load any global assets here (e.g., loading bar, logo, etc.)
        // this.load.image('logo', 'assets/images/logo.png');
    }

    create() {
        console.log('[BootScene] create() called');
        // Defensive: Reset to default state if state is missing/corrupt
        const state = stateManager.getState();
        if (!state || typeof state !== 'object' || !('gameStatus' in state)) {
            console.warn('[BootScene] State missing/corrupt, resetting to default');
            // Reset to default state
            stateManager.reset();
        }

        // Register global input listeners (only once)
        InputSystem.getInstance().registerListeners();

        // Create or get the RenderManager instance (singleton or global)
        const renderManager: IRenderManager = (window as any).renderManager;
        if (!renderManager) {
            throw new Error('RenderManager instance must be available on window');
        }

        // Listen for gameStatus changes and switch scenes accordingly, passing renderManager
        stateManager.on('gameStatusChanged', (status) => {
            console.log('[BootScene] gameStatusChanged:', status);
            if (status === 'mainMenu') {
                this.scene.start('MainMenuScene', { renderManager });
            } else if (status === 'worldSelect') {
                this.scene.start('MenuScene', { renderManager });
            } else if (status === 'levelSelect') {
                this.scene.start('LevelMenuScene', { renderManager });
            } else if (status === 'playing') {
                this.scene.start('GameScene', { renderManager });
            }
            // Add more as needed
        });

        // Always set gameStatus to 'mainMenu' on load for a clean start
        console.log('[BootScene] Setting gameStatus to mainMenu');
        stateManager.setGameStatus('mainMenu');
    }
}
