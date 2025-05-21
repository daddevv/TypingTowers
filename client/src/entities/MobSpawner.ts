// MobSpawner.ts
// Responsible for spawning and managing mobs in the game.
import Phaser from 'phaser';
import Mob from './Mob';

export default class MobSpawner {
    private scene: Phaser.Scene;
    private mobs: Mob[] = [];
    private spawnTimer: number = 0;
    private spawnInterval: number;
    private words: string[];

    constructor(scene: Phaser.Scene, words: string[], spawnInterval: number = 2000) {
        this.scene = scene;
        this.words = words;
        this.spawnInterval = spawnInterval;
    }

    update(time: number, delta: number) {
        this.spawnTimer += delta;
        if (this.spawnTimer >= this.spawnInterval) {
            this.spawnMob();
            this.spawnTimer = 0;
        }
        this.mobs.forEach(mob => mob.update(time, delta));
        this.mobs = this.mobs.filter(mob => !mob.isDefeated);
    }

    /**
     * Returns mobs for collision/proximity checks
     */
    getMobs(): Mob[] {
        return this.mobs;
    }

    /**
     * Removes a mob from the list (e.g., after hitting the player)
     */
    removeMob(mob: Mob) {
        this.mobs = this.mobs.filter(m => m !== mob);
    }

    spawnMob() {
        const letter = this.getRandomLetter();
        const x = Phaser.Math.Between(600, 800); // Spawn off to the right
        const y = Phaser.Math.Between(100, 500);
        const mob = new Mob(this.scene, x, y, letter);
        this.mobs.push(mob);
    }

    getRandomLetter(): string {
        // Flatten all words into a string, then pick a random character
        const allLetters = this.words.join('').split('');
        return allLetters[Phaser.Math.Between(0, allLetters.length - 1)];
    }
}
// Contains AI-generated edits.
