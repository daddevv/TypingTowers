// Mob.ts
// Represents an enemy entity in the game.
import Phaser from 'phaser';

export default class Mob extends Phaser.GameObjects.Sprite {
    public word: string;
    public isDefeated: boolean = false;

    constructor(scene: Phaser.Scene, x: number, y: number, word: string) {
        super(scene, x, y, 'mob');
        this.word = word;
        scene.add.existing(this);
        this.setOrigin(0.5, 0.5);
        // TODO: Add animations, health, and other properties as needed
    }

    update(time: number, delta: number) {
        // TODO: Add movement or behavior logic
    }

    defeat() {
        this.isDefeated = true;
        this.setVisible(false);
        // TODO: Add defeat animation, sound, and effects
    }
}
// Contains AI-generated edits.
