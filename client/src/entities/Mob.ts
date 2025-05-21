// Mob.ts
// Represents an enemy entity in the game.
import Phaser from 'phaser';

export default class Mob extends Phaser.GameObjects.Sprite {
    public word: string;
    public isDefeated: boolean = false;
    private currentLetterIndex: number = 0;
    private isTargeted: boolean = false;
    private letterChars: Phaser.GameObjects.Text[] = [];
    public baseSpeed: number;

    constructor(scene: Phaser.Scene, x: number, y: number, word: string, baseSpeed: number = 60) {
        super(scene, x, y, 'mob');
        this.word = word;
        this.baseSpeed = baseSpeed;
        scene.add.existing(this);
        this.setOrigin(0.5, 0.5);
        // Display each letter as a separate text object for animation
        let offset = -(word.length - 1) * 18 / 2;
        for (let i = 0; i < word.length; i++) {
            const char = word[i];
            const txt = scene.add.text(x + offset + i * 18, y, char, {
                fontSize: '32px',
                color: '#fff',
                fontStyle: 'bold',
                stroke: '#000',
                strokeThickness: 4,
            }).setOrigin(0.5);
            this.letterChars.push(txt);
        }
    }

    update(time: number, delta: number) {
        if (!this.isDefeated) {
            const targetX = 100;
            const targetY = 300;
            const dx = targetX - this.x;
            const dy = targetY - this.y;
            const dist = Math.sqrt(dx * dx + dy * dy);
            if (dist > 1) {
                const speed = this.baseSpeed;
                this.x += (dx / dist) * speed * (delta / 1000);
                this.y += (dy / dist) * speed * (delta / 1000);
            }
            // Update all letter positions and highlight
            let offset = -(this.word.length - 1) * 18 / 2;
            for (let i = 0; i < this.letterChars.length; i++) {
                this.letterChars[i].setPosition(this.x + offset + i * 18, this.y);
                if (i === this.currentLetterIndex && this.isTargeted) {
                    this.letterChars[i].setColor('#ff0');
                    this.letterChars[i].setStyle({ fontStyle: 'bold', backgroundColor: '#333' });
                } else if (i < this.currentLetterIndex) {
                    this.letterChars[i].setColor('#888');
                    this.letterChars[i].setStyle({ fontStyle: 'italic', backgroundColor: undefined });
                } else {
                    this.letterChars[i].setColor('#fff');
                    this.letterChars[i].setStyle({ fontStyle: 'bold', backgroundColor: undefined });
                }
            }
        }
    }

    getNextLetter(): string {
        return this.word[this.currentLetterIndex];
    }

    advanceLetter() {
        if (this.currentLetterIndex < this.word.length - 1) {
            // Animate matched letter
            const txt = this.letterChars[this.currentLetterIndex];
            this.scene.tweens.add({
                targets: txt,
                alpha: 0.5,
                scale: 1.2,
                duration: 120,
                yoyo: true,
                onComplete: () => {
                    txt.setAlpha(1);
                    txt.setScale(1);
                }
            });
            this.currentLetterIndex++;
        } else {
            this.defeat();
        }
    }

    setTargeted(isTargeted: boolean) {
        this.isTargeted = isTargeted;
        for (const txt of this.letterChars) {
            if (isTargeted) {
                txt.setShadow(0, 0, '#ff0', 8, true, true);
            } else {
                txt.setShadow(0, 0, '#000', 0, false, false);
            }
        }
        if (!isTargeted) {
            for (let i = 0; i < this.letterChars.length; i++) {
                this.letterChars[i].setColor(i < this.currentLetterIndex ? '#888' : '#fff');
                this.letterChars[i].setStyle({ fontStyle: i < this.currentLetterIndex ? 'italic' : 'bold', backgroundColor: undefined });
            }
        }
    }

    defeat() {
        this.isDefeated = true;
        for (const txt of this.letterChars) {
            this.scene.tweens.add({
                targets: txt,
                alpha: 0,
                scale: 2,
                duration: 180,
                ease: 'Cubic.easeOut',
                onComplete: () => txt.setVisible(false)
            });
        }
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
            this.scene.time.delayedCall(350, () => {
                particles.destroy();
            });
        }
    }
}
