// Player.ts
// Represents the player character in the game.
import Phaser from 'phaser';

export default class Player extends Phaser.GameObjects.Sprite {
    public health: number = 5;
    private healthText: Phaser.GameObjects.Text | null = null;

    constructor(scene: Phaser.Scene, x: number, y: number) {
        super(scene, x, y, 'player');
        scene.add.existing(this);
        this.setOrigin(0.5, 0.5);
        // Display health above the player
        this.healthText = scene.add.text(x, y - 40, `Health: ${this.health}`, {
            fontSize: '20px',
            color: '#ff5555',
            fontStyle: 'bold',
        }).setOrigin(0.5);
    }

    update(time: number, delta: number) {
        // Player-specific update logic (if any)
        if (this.healthText) {
            this.healthText.setPosition(this.x, this.y - 40);
            this.healthText.setText(`Health: ${this.health}`);
        }
    }

    takeDamage(amount: number = 1) {
        this.health = Math.max(0, this.health - amount);
    }
}
// Contains AI-generated edits.
