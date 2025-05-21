// Mob.ts
// Represents an enemy entity in the game.
import Phaser from 'phaser';

export default class Mob extends Phaser.GameObjects.Sprite {
    public word: string;
    public isDefeated: boolean = false;
    private letterText: Phaser.GameObjects.Text;
    public baseSpeed: number;

    constructor(scene: Phaser.Scene, x: number, y: number, word: string, baseSpeed: number = 60) {
        super(scene, x, y, 'mob');
        this.word = word;
        this.baseSpeed = baseSpeed;
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
                const speed = this.baseSpeed; // Use baseSpeed property
                this.x += (dx / dist) * speed * (delta / 1000);
                this.y += (dy / dist) * speed * (delta / 1000);
            }
            // Keep the letter text centered on the mob
            this.letterText.setPosition(this.x, this.y);
        }
    }

    defeat() {
        this.isDefeated = true;
        // Play a quick flash effect before hiding
        this.scene.tweens.add({
            targets: [this, this.letterText],
            alpha: 0,
            scale: 2,
            duration: 180,
            ease: 'Cubic.easeOut',
            onComplete: () => {
                this.setVisible(false);
                this.letterText.setVisible(false);
                this.setAlpha(1);
                this.setScale(1);
                this.letterText.setAlpha(1);
                this.letterText.setScale(1);
            }
        });
        // Particle burst: short, fades out, and destroys itself
        if (this.scene.add.particles) {
            const particles = this.scene.add.particles(0, 0, 'white', {
                x: this.x,
                y: this.y,
                speed: { min: 60, max: 180 },
                angle: { min: 0, max: 360 },
                lifespan: 300,
                quantity: 12,
                scale: { start: 0.5, end: 0 },
                alpha: { start: 1, end: 0 }
            });
            particles.emitParticleAt(this.x, this.y, 12);
            // Destroy the particle emitter after the burst
            this.scene.time.delayedCall(350, () => {
                particles.destroy();
            });
        }
    }
}
// Contains AI-generated edits.
