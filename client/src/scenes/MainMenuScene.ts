import Phaser from 'phaser';
import { WORLDS } from '../curriculum/worldConfig';
import { levelManager } from '../managers/levelManager';

export default class MainMenuScene extends Phaser.Scene {
    constructor() {
        super('MainMenuScene');
    }

    create() {
        const { width, height } = this.scale;
        this.add.text(width / 2, height / 2 - 80, 'Type Defense', {
            fontSize: '48px',
            color: '#fff',
            fontFamily: 'Arial',
        }).setOrigin(0.5);

        // Add Play button
        const playButton = this.add.text(width / 2, height / 2 + 20, 'Play', {
            fontSize: '36px',
            color: '#0f0',
            backgroundColor: '#222',
            padding: { left: 32, right: 32, top: 12, bottom: 12 },
        }).setOrigin(0.5).setInteractive({ useHandCursor: true });
        playButton.on('pointerdown', () => {
            this.scene.start('MenuScene', { worlds: WORLDS, levelManager });
        });
    }
}
// Contains AI-generated edits.
