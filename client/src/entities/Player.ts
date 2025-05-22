// Player.ts
// Represents the player character in the game.
import Phaser from 'phaser';
import stateManager from '../state/stateManager';

export default class Player extends Phaser.GameObjects.Sprite {
    private healthText: Phaser.GameObjects.Text | null = null;

    constructor(scene: Phaser.Scene, x: number, y: number) {
        super(scene, x, y, 'player');
        scene.add.existing(this);
        this.setOrigin(0.5, 0.5);
        // Display health above the player
        this.healthText = scene.add.text(x, y - 40, `Health: 0`, {
            fontSize: '20px',
            color: '#ff5555',
            fontStyle: 'bold',
        }).setOrigin(0.5);
    }

    update(time: number, delta: number) {
        // Player-specific update logic (if any)
        const playerState = stateManager.getState().player;
        this.x = playerState.position.x;
        this.y = playerState.position.y;
        if (this.healthText) {
            this.healthText.setPosition(this.x, this.y - 40);
            this.healthText.setText(`Health: ${playerState.health}`);
        }
    }

    // Remove takeDamage and local health, use StateManager instead
}
