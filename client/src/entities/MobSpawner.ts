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

    spawnMob() {
        const word = this.getRandomWord();
        const x = Phaser.Math.Between(600, 800); // Spawn off to the right
        const y = Phaser.Math.Between(100, 500);
        const mob = new Mob(this.scene, x, y, word);
        this.mobs.push(mob);
    }

    getRandomWord(): string {
        return this.words[Phaser.Math.Between(0, this.words.length - 1)];
    }

    getMobs(): Mob[] {
        return this.mobs;
    }
}
// Contains AI-generated edits.
