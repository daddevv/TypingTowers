import Phaser from 'phaser';
import { getKeyInfo } from '../curriculum/fingerGroups';
import { WORLDS } from '../curriculum/worldConfig';
import InputHandler from '../entities/InputHandler';
import MobSpawner from '../entities/MobSpawner';
import Player from '../entities/Player';
import FingerGroupManager from '../managers/fingerGroupManager';
import { loadWordList } from '../utils/loadWordList';

export default class GameScene extends Phaser.Scene {
    private player!: Player;
    private inputHandler!: InputHandler;
    private mobSpawner!: MobSpawner;
    private fingerGroupManager!: FingerGroupManager;
    private targetedMob: any = null; // Track the currently targeted mob

    constructor() {
        super('GameScene');
    }

    preload() {
        // Load assets here
        // Example: this.load.image('player', 'assets/images/player.png');
        // Example: this.load.image('mob', 'assets/images/mob.png');
    }

    async create() {
        // Set up the main game scene here
        this.add.text(400, 300, 'Type Defense', {
            fontSize: '48px',
            color: '#fff',
        }).setOrigin(0.5);

        // Initialize Player and InputHandler
        this.player = new Player(this, 100, 300);
        this.inputHandler = new InputHandler(this);
        this.fingerGroupManager = new FingerGroupManager();

        // Load Level 1-1 config and word list
        const world = WORLDS[0];
        const level = world.levels[0]; // Level 1-1
        const words = await loadWordList(level.id);
        // Spawn 2 mobs per interval for demonstration (can be made configurable)
        this.mobSpawner = new MobSpawner(this, words, level.enemySpawnRate, 2);
    }

    update(time: number, delta: number) {
        // Guard: wait for mobSpawner to be initialized
        if (!this.mobSpawner) return;
        // Core game loop logic
        if (this.player) {
            this.player.update(time, delta);
        }
        // Update MobSpawner and its mobs
        this.mobSpawner.update(time, delta);
        // Check for mobs reaching the player (collision/proximity)
        const mobs = this.mobSpawner.getMobs();
        for (const mob of mobs) {
            if (!mob.isDefeated) {
                const dx = mob.x - this.player.x;
                const dy = mob.y - this.player.y;
                const dist = Math.sqrt(dx * dx + dy * dy);
                if (dist < 40) { // collision radius
                    this.player.takeDamage(1);
                    this.mobSpawner.removeMob(mob);
                    // Optionally: add hit feedback here
                    if (this.player.health <= 0) {
                        this.add.text(400, 300, 'Game Over', { fontSize: '48px', color: '#ff5555' }).setOrigin(0.5);
                        this.scene.pause();
                    }
                    break;
                }
            }
        }
        // Improved mob input handling for combos and multiple mobs
        const input = this.inputHandler.getInput();
        if (input.length > 0) {
            for (const char of input) {
                const keyInfo = getKeyInfo(char);
                if (keyInfo) {
                    this.fingerGroupManager.recordKeyPress(char, true, time);
                }
                let handled = false;
                // 1. If no mob is targeted, find all mobs whose next letter matches the key
                if (!this.targetedMob) {
                    const candidates = mobs.filter(mob => !mob.isDefeated && mob.getNextLetter && mob.getNextLetter().toLowerCase() === char.toLowerCase());
                    if (candidates.length > 0) {
                        // Pick the closest mob
                        let minDist = Infinity;
                        let closest = null;
                        for (const mob of candidates) {
                            const dx = mob.x - this.player.x;
                            const dy = mob.y - this.player.y;
                            const dist = Math.sqrt(dx * dx + dy * dy);
                            if (dist < minDist) {
                                minDist = dist;
                                closest = mob;
                            }
                        }
                        this.targetedMob = closest;
                        if (this.targetedMob) {
                            this.targetedMob.setTargeted(true);
                        }
                    }
                }
                // 2. If a mob is targeted, check if key matches its next letter
                if (this.targetedMob && this.targetedMob.getNextLetter && !this.targetedMob.isDefeated) {
                    if (this.targetedMob.getNextLetter().toLowerCase() === char.toLowerCase()) {
                        this.targetedMob.advanceLetter();
                        handled = true;
                        // If mob is defeated, clear target
                        if (this.targetedMob.isDefeated) {
                            this.targetedMob.setTargeted(false);
                            this.targetedMob = null;
                        }
                    } else {
                        // 3. If not, check if key matches any other mob's next letter
                        const otherCandidates = mobs.filter(mob => mob !== this.targetedMob && !mob.isDefeated && mob.getNextLetter && mob.getNextLetter().toLowerCase() === char.toLowerCase());
                        if (otherCandidates.length > 0) {
                            // Retarget to the closest matching mob
                            let minDist = Infinity;
                            let closest = null;
                            for (const mob of otherCandidates) {
                                const dx = mob.x - this.player.x;
                                const dy = mob.y - this.player.y;
                                const dist = Math.sqrt(dx * dx + dy * dy);
                                if (dist < minDist) {
                                    minDist = dist;
                                    closest = mob;
                                }
                            }
                            if (this.targetedMob) this.targetedMob.setTargeted(false);
                            this.targetedMob = closest;
                            if (this.targetedMob) {
                                this.targetedMob.setTargeted(true);
                                this.targetedMob.advanceLetter();
                                handled = true;
                                if (this.targetedMob.isDefeated) {
                                    this.targetedMob.setTargeted(false);
                                    this.targetedMob = null;
                                }
                            }
                        } else {
                            // No match, clear target
                            if (this.targetedMob) this.targetedMob.setTargeted(false);
                            this.targetedMob = null;
                        }
                    }
                }
                // Ensure only one mob is targeted at a time
                for (const mob of mobs) {
                    if (mob !== this.targetedMob && mob.setTargeted) {
                        mob.setTargeted(false);
                    }
                }
            }
            this.inputHandler.clearInput();
        }
    }
}
// Contains AI-generated edits.
