// MobSpawner.ts
// Responsible for spawning and managing mobs in the game.
import { v4 as uuidv4 } from 'uuid'; // For unique mob IDs
import stateManager from '../state/stateManager';
import WordGenerator from '../utils/wordGenerator';

function clamp(val: number, min: number, max: number) {
    return Math.max(min, Math.min(max, val));
}
function lerp(a: number, b: number, t: number) {
    return a + (b - a) * t;
}
function randomBetween(min: number, max: number) {
    return Math.floor(Math.random() * (max - min + 1)) + min;
}

export interface MobSpawnerConfig {
    width: number;
    height: number;
}

export default class MobSpawner {
    private width: number;
    private height: number;
    private spawnTimer: number = 0;
    private spawnInterval: number;
    private wordGenerator: WordGenerator;
    private mobsPerInterval: number = 1;
    private mobBaseSpeed: number = 90;
    private progression: number = 0;
    private minSpawnInterval: number = 600;
    private maxMobSpeed: number = 250;
    private initialSpawnInterval: number;
    private initialMobBaseSpeed: number;
    private currentWave: number = 0;
    private waveInProgress: boolean = false;
    private waveDelay: number = 2000;
    private waveTimer: number = 0;
    private mobsPerWave: number = 5;
    private mobsSpawnedThisWave: number = 0;
    private onWaveStartCallback?: (wave: number) => void;
    private onWaveEndCallback?: (wave: number) => void;
    private wordList?: string[];
    private wordListIndex: number = 0;

    constructor(config: MobSpawnerConfig, wordGenerator: WordGenerator, spawnInterval: number = 2000, mobsPerInterval: number = 1, mobBaseSpeed: number = 90, wordList?: string[], mobsPerWave?: number) {
        this.width = config.width;
        this.height = config.height;
        this.wordGenerator = wordGenerator;
        this.spawnInterval = spawnInterval;
        this.mobsPerInterval = mobsPerInterval;
        this.mobBaseSpeed = mobBaseSpeed;
        this.initialSpawnInterval = spawnInterval;
        this.initialMobBaseSpeed = mobBaseSpeed;
        // Always set wordList if provided, even if empty, so empty list disables spawning
        if (typeof wordList !== 'undefined') {
            this.wordList = wordList;
        }
        this.wordListIndex = 0;
        if (typeof mobsPerWave === 'number' && mobsPerWave > 0) {
            this.mobsPerWave = mobsPerWave;
        }
    }

    /**
     * Call this method to update scaling based on progression (0-1)
     */
    public setProgression(progression: number) {
        this.progression = clamp(progression, 0, 1);
        // Linearly interpolate spawn interval and mob speed
        this.spawnInterval = lerp(this.initialSpawnInterval, this.minSpawnInterval, this.progression);
        this.mobBaseSpeed = lerp(this.initialMobBaseSpeed, this.maxMobSpeed, this.progression);
    }

    /**
     * Start the next wave. Call from GameScene to trigger a new wave.
     */
    public startNextWave() {
        // Prevent starting a wave if wordList is defined and empty
        if (this.wordList && this.wordList.length === 0) {
            this.waveInProgress = false;
            if (this.onWaveEndCallback) this.onWaveEndCallback(this.currentWave + 1);
            // Update mobSpawner state in gameState
            stateManager.updateMobSpawnerState({
                nextSpawnTime: this.spawnTimer,
                currentWave: this.currentWave + 1,
                mobsRemainingInWave: 0,
            });
            return;
        }
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
        // Prevent spawning if wordList is defined but empty
        if (this.wordList && this.wordList.length === 0) {
            this.waveInProgress = false;
            if (this.onWaveEndCallback) this.onWaveEndCallback(this.currentWave);
            // Also update mobSpawner state in gameState
            stateManager.updateMobSpawnerState({
                nextSpawnTime: this.spawnTimer,
                currentWave: this.currentWave,
                mobsRemainingInWave: 0,
            });
            return;
        }
        this.spawnTimer += delta;
        // Spawn mobs for this wave
        if (this.mobsSpawnedThisWave < this.mobsPerWave && this.spawnTimer >= this.spawnInterval) {
            for (let i = 0; i < this.mobsPerInterval && this.mobsSpawnedThisWave < this.mobsPerWave; i++) {
                let word: string;
                if (this.wordList && this.wordList.length > 0) {
                    word = this.wordList[this.wordListIndex % this.wordList.length];
                    this.wordListIndex++;
                } else {
                    word = this.wordGenerator.generateWord(randomBetween(2, 5));
                }
                const minY = 100;
                const maxY = this.height - 100;
                const y = randomBetween(minY, maxY);
                const mobState = {
                    id: uuidv4(),
                    word,
                    currentTypedIndex: 0,
                    position: { x: this.width + 50, y },
                    speed: this.mobBaseSpeed,
                    type: 'normal',
                    isDefeated: false,
                };
                stateManager.addMob(mobState);
                this.mobsSpawnedThisWave++;
            }
            this.spawnTimer = 0;
        }
        // Remove defeated mobs from gameState
        const mobs = stateManager.getState().mobs;
        for (const mob of mobs) {
            if (mob.isDefeated) {
                stateManager.removeMob(mob.id);
            }
        }
        // If all mobs for this wave are defeated and all have spawned, end wave
        const remainingMobs = stateManager.getState().mobs.filter(m => !m.isDefeated);
        if (this.mobsSpawnedThisWave >= this.mobsPerWave && remainingMobs.length === 0) {
            this.waveInProgress = false;
            if (this.onWaveEndCallback) this.onWaveEndCallback(this.currentWave);
        }
        // Update mobSpawner state in gameState using a new method
        stateManager.updateMobSpawnerState({
            nextSpawnTime: this.spawnTimer,
            currentWave: this.currentWave,
            mobsRemainingInWave: this.mobsPerWave - this.mobsSpawnedThisWave,
        });
    }

    /**
     * Returns mobs for collision/proximity checks
     */
    getMobs(): any[] {
        return stateManager.getState().mobs;
    }

    /**
     * Removes a mob from the list (e.g., after hitting the player)
     */
    removeMob(mob: any) {
        stateManager.removeMob(mob.id);
    }

    /**
     * Optionally allow changing mobsPerInterval at runtime
     */
    setMobsPerInterval(count: number) {
        this.mobsPerInterval = count;
    }
    /**
     * Optionally allow changing mobsPerWave at runtime
     */
    setMobsPerWave(count: number) {
        if (typeof count === 'number' && count > 0) {
            this.mobsPerWave = count;
        }
    }
}
// Contains AI-generated edits.
