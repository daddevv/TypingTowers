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
        // Get game width from scene's scale manager
        const gameWidth = this.scene.scale.width;
        // Create a temporary mob to get its width
        const tempMob = new Mob(this.scene, 0, 0, letter);
        const mobWidth = tempMob.displayWidth || 64; // fallback if not loaded
        tempMob.destroy();
        // Spawn X is fully off the right edge
        const x = gameWidth + mobWidth / 2;
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
