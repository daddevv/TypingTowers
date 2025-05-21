// MobSpawner.ts
// Responsible for spawning and managing mobs in the game.
import Phaser from 'phaser';
import WordGenerator from '../utils/wordGenerator';
import Mob from './Mob';

export default class MobSpawner {
    private scene: Phaser.Scene;
    private mobs: Mob[] = [];
    private spawnTimer: number = 0;
    private spawnInterval: number;
    private wordGenerator: WordGenerator;
    private mobsPerInterval: number = 1;
    private mobBaseSpeed: number = 90; // Increased default base speed for more challenge
    private progression: number = 0; // 0 to 1, represents game progress
    private minSpawnInterval: number = 600; // ms
    private maxMobSpeed: number = 250;
    private initialSpawnInterval: number;
    private initialMobBaseSpeed: number;
    private currentWave: number = 0;
    private waveInProgress: boolean = false;
    private waveDelay: number = 2000; // ms between waves
    private waveTimer: number = 0;
    private mobsPerWave: number = 5;
    private mobsSpawnedThisWave: number = 0;
    private onWaveStartCallback?: (wave: number) => void;
    private onWaveEndCallback?: (wave: number) => void;
    private wordList?: string[];
    private wordListIndex: number = 0;

    constructor(scene: Phaser.Scene, wordGenerator: WordGenerator, spawnInterval: number = 2000, mobsPerInterval: number = 1, mobBaseSpeed: number = 90, wordList?: string[]) {
        this.scene = scene;
        this.wordGenerator = wordGenerator;
        this.spawnInterval = spawnInterval;
        this.mobsPerInterval = mobsPerInterval;
        this.mobBaseSpeed = mobBaseSpeed;
        this.initialSpawnInterval = spawnInterval;
        this.initialMobBaseSpeed = mobBaseSpeed;
        this.wordList = wordList && wordList.length > 0 ? wordList : undefined;
        this.wordListIndex = 0;
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

    /**
     * Start the next wave. Call from GameScene to trigger a new wave.
     */
    public startNextWave() {
        this.currentWave++;
        this.waveInProgress = true;
        this.waveTimer = 0;
        this.mobsSpawnedThisWave = 0;
        if (this.onWaveStartCallback) this.onWaveStartCallback(this.currentWave);
    }

    /**
     * Register a callback for when a wave starts.
     */
    public onWaveStart(cb: (wave: number) => void) {
        this.onWaveStartCallback = cb;
    }
    /**
     * Register a callback for when a wave ends.
     */
    public onWaveEnd(cb: (wave: number) => void) {
        this.onWaveEndCallback = cb;
    }

    /**
     * Returns whether a wave is in progress.
     */
    public isWaveInProgress() {
        return this.waveInProgress;
    }
    /**
     * Returns the current wave number.
     */
    public getCurrentWave() {
        return this.currentWave;
    }

    update(time: number, delta: number) {
        if (!this.waveInProgress) return;
        this.spawnTimer += delta;
        // Spawn mobs for this wave
        if (this.mobsSpawnedThisWave < this.mobsPerWave && this.spawnTimer >= this.spawnInterval) {
            for (let i = 0; i < this.mobsPerInterval && this.mobsSpawnedThisWave < this.mobsPerWave; i++) {
                // Use wordList if available, else WordGenerator
                let word: string;
                if (this.wordList && this.wordList.length > 0) {
                    word = this.wordList[this.wordListIndex % this.wordList.length];
                    this.wordListIndex++;
                } else {
                    word = this.wordGenerator.getWord(Phaser.Math.Between(2, 5));
                }
                const minY = 100;
                const maxY = this.scene.scale.height - 100;
                const y = Math.floor(Math.random() * (maxY - minY + 1)) + minY;
                const mob = new Mob(this.scene, this.scene.scale.width + 50, y, word, this.mobBaseSpeed);
                this.mobs.push(mob);
                this.mobsSpawnedThisWave++;
            }
            this.spawnTimer = 0;
        }
        // Update mobs and let each mob handle its own avoidance
        this.mobs.forEach(mob => mob.update(time, delta, this.mobs));
        this.mobs = this.mobs.filter(mob => !mob.isDefeated);
        // If all mobs for this wave are defeated and all have spawned, end wave
        if (this.mobsSpawnedThisWave >= this.mobsPerWave && this.mobs.length === 0) {
            this.waveInProgress = false;
            if (this.onWaveEndCallback) this.onWaveEndCallback(this.currentWave);
        }
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
        // Use wordList if available, else WordGenerator
        let word: string;
        if (this.wordList && this.wordList.length > 0) {
            word = this.wordList[this.wordListIndex % this.wordList.length];
            this.wordListIndex++;
        } else {
            word = this.wordGenerator.getWord(Phaser.Math.Between(2, 5));
        }
        const gameWidth = this.scene.scale.width;
        const tempMob = new Mob(this.scene, 0, 0, word, this.mobBaseSpeed);
        const mobWidth = tempMob.displayWidth || 64;
        tempMob.destroy();
        const x = gameWidth + mobWidth / 2;
        const y = Phaser.Math.Between(100, 500);
        const mob = new Mob(this.scene, x, y, word, this.mobBaseSpeed);
        this.mobs.push(mob);
    }
}
// Contains AI-generated edits.
