// Mob.ts
// Represents an enemy entity in the game.
import Phaser from 'phaser';

export default class Mob extends Phaser.GameObjects.Sprite {
    public word: string;
    public isDefeated: boolean = false;
    private letterText: Phaser.GameObjects.Text;

    constructor(scene: Phaser.Scene, x: number, y: number, word: string) {
        super(scene, x, y, 'mob');
        this.word = word;
        scene.add.existing(this);
        this.setOrigin(0.5, 0.5);
        // Display the letter on the mob
        this.letterText = scene.add.text(x, y, word, {
            fontSize: '32px',
            color: '#fff',
            fontStyle: 'bold',
            stroke: '#000',
            strokeThickness: 4,
        }).setOrigin(0.5);
    }

    update(time: number, delta: number) {
        // Move mob toward the player (assume player is at x=100, y=300)
        if (!this.isDefeated) {
            const targetX = 100;
            const targetY = 300;
            const dx = targetX - this.x;
            const dy = targetY - this.y;
            const dist = Math.sqrt(dx * dx + dy * dy);
            if (dist > 1) {
                const speed = 60; // pixels per second
                this.x += (dx / dist) * speed * (delta / 1000);
                this.y += (dy / dist) * speed * (delta / 1000);
            }
            // Keep the letter text centered on the mob
            this.letterText.setPosition(this.x, this.y);
        }
    }

    defeat() {
        this.isDefeated = true;
        this.setVisible(false);
        this.letterText.setVisible(false);
        // TODO: Add defeat animation, sound, and effects
    }
}
