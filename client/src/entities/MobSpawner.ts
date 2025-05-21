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
    private mobsPerInterval: number = 1;
    private mobBaseSpeed: number = 90; // Increased default base speed for more challenge
    private progression: number = 0; // 0 to 1, represents game progress
    private minSpawnInterval: number = 600; // ms
    private maxMobSpeed: number = 250;
    private initialSpawnInterval: number;
    private initialMobBaseSpeed: number;

    constructor(scene: Phaser.Scene, words: string[], spawnInterval: number = 2000, mobsPerInterval: number = 1, mobBaseSpeed: number = 90) {
        this.scene = scene;
        this.words = words;
        this.spawnInterval = spawnInterval;
        this.mobsPerInterval = mobsPerInterval;
        this.mobBaseSpeed = mobBaseSpeed;
        this.initialSpawnInterval = spawnInterval;
        this.initialMobBaseSpeed = mobBaseSpeed;
    }

    /**
     * Call this method to update scaling based on progression (0-1)
     */
    public setProgression(progression: number) {
        this.progression = Phaser.Math.Clamp(progression, 0, 1);
        // Linearly interpolate spawn interval and mob speed
        this.spawnInterval = Phaser.Math.Linear(this.initialSpawnInterval, this.minSpawnInterval, this.progression);
        this.mobBaseSpeed = Phaser.Math.Linear(this.initialMobBaseSpeed, this.maxMobSpeed, this.progression);
    }

    update(time: number, delta: number) {
        this.spawnTimer += delta;
        if (this.spawnTimer >= this.spawnInterval) {
            for (let i = 0; i < this.mobsPerInterval; i++) {
                // Choose a random word for the mob
                const word = this.words[Math.floor(Math.random() * this.words.length)];
                // Add y-position variation: spawn mobs at random vertical positions
                const minY = 100;
                const maxY = this.scene.scale.height - 100;
                const y = Math.floor(Math.random() * (maxY - minY + 1)) + minY;
                const mob = new Mob(this.scene, this.scene.scale.width + 50, y, word, this.mobBaseSpeed);
                this.mobs.push(mob);
            }
            this.spawnTimer = 0;
        }
        // Update mobs and let each mob handle its own avoidance
        this.mobs.forEach(mob => mob.update(time, delta, this.mobs));
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

    /**
     * Optionally allow changing mobsPerInterval at runtime
     */
    setMobsPerInterval(count: number) {
        this.mobsPerInterval = count;
    }

    spawnMob() {
        const letter = this.getRandomLetter();
        // Get game width from scene's scale manager
        const gameWidth = this.scene.scale.width;
        // Create a temporary mob to get its width
        const tempMob = new Mob(this.scene, 0, 0, letter, this.mobBaseSpeed);
        const mobWidth = tempMob.displayWidth || 64; // fallback if not loaded
        tempMob.destroy();
        // Spawn X is fully off the right edge
        const x = gameWidth + mobWidth / 2;
        const y = Phaser.Math.Between(100, 500);
        const mob = new Mob(this.scene, x, y, letter, this.mobBaseSpeed);
        this.mobs.push(mob);
    }

    getRandomLetter(): string {
        // Flatten all words into a string, then pick a random character
        const allLetters = this.words.join('').split('');
        return allLetters[Phaser.Math.Between(0, allLetters.length - 1)];
    }
}
