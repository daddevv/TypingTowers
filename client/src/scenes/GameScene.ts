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
            // Record each key press in FingerGroupManager
            for (const char of input) {
                const keyInfo = getKeyInfo(char);
                if (keyInfo) {
                    this.fingerGroupManager.recordKeyPress(char, true, time); // 'true' for correct finger (future: detect real finger)
                }
            }
            // Find the closest mob whose word starts with the input
            let closestMob: any = null;
            let minDist = Infinity;
            for (const mob of mobs) {
                if (!mob.isDefeated) {
                    // Only consider mobs whose word starts with the input
                    if (mob.word.toLowerCase().startsWith(input.trim().toLowerCase())) {
                        const dx = mob.x - this.player.x;
                        const dy = mob.y - this.player.y;
                        const dist = Math.sqrt(dx * dx + dy * dy);
                        if (dist < minDist) {
                            minDist = dist;
                            closestMob = mob;
                        }
                    }
                }
            }
            let mobDefeated = false;
            if (closestMob) {
                // If input matches the full word, defeat the mob
                if (input.trim().toLowerCase() === closestMob.word.toLowerCase()) {
                    closestMob.defeat();
                    this.inputHandler.clearInput();
                    mobDefeated = true;
                }
            } else {
                // If no closest mob, check if input matches any mob's word (fallback)
                for (const mob of mobs) {
                    if (!mob.isDefeated && input.trim().toLowerCase() === mob.word.toLowerCase()) {
                        mob.defeat();
                        this.inputHandler.clearInput();
                        mobDefeated = true;
                        break;
                    }
                }
            }
            // If no mob was defeated, reset all mobs (drop combo/progress)
            if (!mobDefeated) {
                this.inputHandler.clearInput();
            // Optionally: add a method to Mob to reset progress if you track per-mob progress
            }
        }
    }
}
// Contains AI-generated edits.
