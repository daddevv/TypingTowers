import Phaser from 'phaser';
import stateManager from '../state/stateManager';

export default class MainMenuScene extends Phaser.Scene {
    private onGameStatusChanged?: (status: string) => void;

    constructor() {
        super('MainMenuScene');
    }

    create() {
        // Remove all children to clear any lingering error messages or UI
        this.children.removeAll();
        const { width, height } = this.scale;
        this.add.text(width / 2, height / 2 - 80, 'Type Defense', {
            fontSize: '48px',
            color: '#fff',
            fontFamily: 'Arial',
        }).setOrigin(0.5);

        // Defensive: Check stateManager
        if (!stateManager || typeof stateManager.setGameStatus !== 'function') {
            this.add.text(width / 2, height / 2 + 80, 'State manager error. Cannot start game.', { fontSize: '24px', color: '#f00' }).setOrigin(0.5);
            return;
        }

        // Add Play button
        const playButton = this.add.text(width / 2, height / 2 + 20, 'Play', {
            fontSize: '36px',
            color: '#0f0',
            backgroundColor: '#222',
            padding: { left: 32, right: 32, top: 12, bottom: 12 },
        }).setOrigin(0.5).setInteractive({ useHandCursor: true });

        // Reset game state on Play
        playButton.on('pointerdown', () => {
            stateManager.resetState();
            // Transition to world select or game start as appropriate
            stateManager.setGameStatus('worldSelect');
        });
        // Remove Back button from title screen
        // Listen for gameStatus changes and transition if needed
        this.onGameStatusChanged = (status: string) => {
            if (status !== 'mainMenu') {
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
}
