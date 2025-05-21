import Phaser from 'phaser';
import { getKeyInfo } from '../curriculum/fingerGroups';
import { WORLDS } from '../curriculum/worldConfig';
import InputHandler from '../entities/InputHandler';
import MobSpawner from '../entities/MobSpawner';
import Player from '../entities/Player';
import FingerGroupManager from '../managers/fingerGroupManager';
import LevelManager from '../managers/levelManager';
import { loadWordList } from '../utils/loadWordList';

export default class GameScene extends Phaser.Scene {
    private player!: Player;
    private inputHandler!: InputHandler;
    private mobSpawner!: MobSpawner;
    private fingerGroupManager!: FingerGroupManager;
    private targetedMob: any = null; // Track the currently targeted mob
    private elapsedTime: number = 0;
    private scalingDuration: number = 120000; // 2 minutes to max difficulty
    private defeatedCount: number = 0; // Track number of defeated enemies
    private winThreshold: number = 50; // Number of enemies to defeat to win
    private levelCompleteText?: Phaser.GameObjects.Text;

    protected score: number = 0;
    protected combo: number = 0;
    protected particleManager!: Phaser.GameObjects.Particles.ParticleEmitter;
    protected scoreText!: Phaser.GameObjects.Text;
    protected comboText!: Phaser.GameObjects.Text;

    private levelManager!: LevelManager;
    private currentWorldIdx: number = 0;
    private currentLevelIdx: number = 0;

    constructor() {
        super('GameScene');
    }

    preload() {
        // Load assets here
        // Example: this.load.image('player', 'assets/images/player.png');
        // Example: this.load.image('mob', 'assets/images/mob.png');
    }

    async create(data?: { world?: number; level?: number; levelManager?: LevelManager }) {
        // Set up the main game scene here
        this.add.text(400, 300, 'Type Defense', {
            fontSize: '48px',
            color: '#fff',
        }).setOrigin(0.5);

        // Initialize Player and InputHandler
        this.player = new Player(this, 100, 300);
        this.inputHandler = new InputHandler(this);
        this.fingerGroupManager = new FingerGroupManager();

        // Support starting at a specific world/level
        if (data && typeof data.world === 'number' && typeof data.level === 'number') {
            this.currentWorldIdx = data.world;
            this.currentLevelIdx = data.level;
        } else {
            this.currentWorldIdx = 0;
            this.currentLevelIdx = 0;
        }
        // Use LevelManager from data or create new
        if (data && data.levelManager) {
            this.levelManager = data.levelManager;
        } else {
            this.levelManager = new LevelManager();
            this.levelManager.loadProgress();
        }
        const world = WORLDS[this.currentWorldIdx];
        const level = world.levels[this.currentLevelIdx];
        const words = await loadWordList(level.id);
        // Spawn 2 mobs per interval for demonstration (can be made configurable)
        this.mobSpawner = new MobSpawner(this, words, level.enemySpawnRate, 2);

        // Score and combo UI
        this.score = 0;
        this.combo = 0;
        this.scoreText = this.add.text(16, 16, 'Score: 0', { fontSize: '28px', color: '#fff', stroke: '#000', strokeThickness: 4 });
        this.comboText = this.add.text(16, 52, '', { fontSize: '22px', color: '#ff0', stroke: '#000', strokeThickness: 3 });
        this.comboText.setVisible(false);

        // Particle manager for bursts (white circle texture)
        if (!this.textures.exists('white')) {
            const g = this.add.graphics();
            g.fillStyle(0xffffff, 1);
            g.fillCircle(8, 8, 8);
            g.generateTexture('white', 16, 16);
            g.destroy();
        }
        // Only create the emitter for bursts, not a persistent emitter at (0,0)
        this.particleManager = this.add.particles(0, 0, 'white', {
            speed: { min: 80, max: 180 },
            angle: { min: 0, max: 360 },
            lifespan: 350,
            quantity: 12,
            scale: { start: 0.5, end: 0 },
            alpha: { start: 1, end: 0 },
            emitting: false // Do not emit constantly
        });
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
            let anyCorrect = false;
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
                        anyCorrect = true;
                        // Particle burst at mob position
                        this.particleManager.emitParticleAt(this.targetedMob.x, this.targetedMob.y, 12);
                        // Score and combo
                        this.combo++;
                        this.score += 10 * this.combo;
                        this.scoreText.setText(`Score: ${this.score}`);
                        this.comboText.setText(`Combo x${this.combo}`);
                        this.comboText.setVisible(this.combo > 1);
                        // If mob is defeated, clear target and increment defeated count
                        if (this.targetedMob.isDefeated) {
                            this.targetedMob.setTargeted(false);
                            this.targetedMob = null;
                            this.defeatedCount++;
                            // Check win condition
                            if (this.defeatedCount >= this.winThreshold && !this.levelCompleteText) {
                                this.handleLevelComplete();
                            }
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
                            // Reset the previous target's progress if switching
                            if (this.targetedMob && this.targetedMob.resetProgress) {
                                this.targetedMob.resetProgress();
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
                            if (this.targetedMob && this.targetedMob.resetProgress) {
                                this.targetedMob.resetProgress();
                            }
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
            // If no correct keypress, reset combo
            if (!anyCorrect) {
                this.combo = 0;
                this.comboText.setVisible(false);
            }
            this.inputHandler.clearInput();
        }
        this.elapsedTime += delta;
        // Progression: 0 at start, 1 at scalingDuration (capped)
        const progression = Phaser.Math.Clamp(this.elapsedTime / this.scalingDuration, 0, 1);
        if (this.mobSpawner && typeof this.mobSpawner.setProgression === 'function') {
            this.mobSpawner.setProgression(progression);
        }
    }

    /**
     * Handles level completion logic: show message and pause game
     */
    private handleLevelComplete() {
        this.levelCompleteText = this.add.text(400, 300, 'Level Complete!', {
            fontSize: '48px',
            color: '#44ff44',
            fontStyle: 'bold',
            backgroundColor: '#222',
            padding: { left: 24, right: 24, top: 12, bottom: 12 },
        }).setOrigin(0.5);
        // Add Continue button
        const continueButton = this.add.text(400, 380, 'Continue', {
            fontSize: '32px',
            color: '#fff',
            backgroundColor: '#007bff',
            padding: { left: 24, right: 24, top: 8, bottom: 8 },
        }).setOrigin(0.5).setInteractive({ useHandCursor: true });
        continueButton.on('pointerdown', () => this.handleContinue());
        // Add Enter key handler
        const enterHandler = (event: KeyboardEvent) => {
            if (event.key === 'Enter') {
                this.handleContinue();
            }
        };
        this.input.keyboard?.on('keydown', enterHandler);
        // Clean up listeners when scene shuts down
        this.events.once('shutdown', () => {
            this.input.keyboard?.off('keydown', enterHandler);
        });
        this.scene.pause();

        // Mark level as completed and unlock next
        const world = WORLDS[this.currentWorldIdx];
        const level = world.levels[this.currentLevelIdx];
        if (this.levelManager) {
            // Use completeLevel to mark as completed and save progress
            this.levelManager.completeLevel(level.id, { score: this.score, wpm: 0, accuracy: 1 });
        }
    }

    /**
     * Handles advancing to the next level or returning to menu
     */
    private handleContinue() {
        // Find current world/level
        const world = WORLDS[this.currentWorldIdx];
        const nextLevelIdx = this.currentLevelIdx + 1;
        if (world && nextLevelIdx < world.levels.length) {
            // Advance to next level in the same world
            this.scene.restart({ world: this.currentWorldIdx, level: nextLevelIdx, levelManager: this.levelManager });
        } else if (this.currentWorldIdx < WORLDS.length - 1) {
            // Go to first level of next world
            this.scene.restart({ world: this.currentWorldIdx + 1, level: 0, levelManager: this.levelManager });
        } else {
            // No more levels, go back to menu
            this.scene.start('MenuScene', { worlds: WORLDS, levelManager: this.levelManager });
        }
    }

    /**
     * For testing: Simulate a correct keystroke at a given position (triggers particle burst, updates score/combo)
     */
    public handleCorrectKeystroke(pos: { x: number; y: number }) {
        this.combo++;
        this.score += 10 * this.combo;
        if (this.particleManager && typeof this.particleManager.emitParticleAt === 'function') {
            this.particleManager.emitParticleAt(pos.x, pos.y, 12);
        }
        if (this.scoreText) this.scoreText.setText(`Score: ${this.score}`);
        if (this.comboText) {
            this.comboText.setText(`Combo x${this.combo}`);
            this.comboText.setVisible(this.combo > 1);
        }
    }

    /**
     * For testing: Simulate an incorrect keystroke (resets combo)
     */
    public handleIncorrectKeystroke() {
        this.combo = 0;
        if (this.comboText) this.comboText.setVisible(false);
    }
}
// Contains AI-generated edits.
