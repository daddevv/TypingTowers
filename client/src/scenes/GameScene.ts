import Phaser from 'phaser';
import { getKeyInfo } from '../curriculum/fingerGroups';
import { WORLDS } from '../curriculum/worldConfig';
import InputHandler from '../entities/InputHandler';
import MobSpawner from '../entities/MobSpawner';
import Player from '../entities/Player';
import FingerGroupManager from '../managers/fingerGroupManager';
import LevelManager, { levelManager } from '../managers/levelManager';
import stateManager from '../state/stateManager';
import { loadWordList } from '../utils/loadWordList';
import WordGenerator from '../utils/wordGenerator';

export default class GameScene extends Phaser.Scene {
    private player!: Player;
    private inputHandler!: InputHandler;
    private mobSpawner!: MobSpawner;
    private fingerGroupManager!: FingerGroupManager;
    private targetedMob: any = null; // Track the currently targeted mob
    private elapsedTime: number = 0;
    private scalingDuration: number = 120000; // 2 minutes to max difficulty
    private minWordLength: number = 2;
    private maxWordLength: number = 5;
    private minSpawnInterval: number = 600;
    private maxMobSpeed: number = 250;
    private scalingProgression: number = 0;
    private defeatedCount: number = 0; // Track number of defeated enemies
    private winThreshold: number = 50; // Number of enemies to defeat to win
    private levelCompleteText?: Phaser.GameObjects.Text;

    protected score: number = 0;
    protected combo: number = 0;
    protected particleManager!: Phaser.GameObjects.Particles.ParticleEmitter;
    protected scoreText!: Phaser.GameObjects.Text;
    protected comboText!: Phaser.GameObjects.Text;
    protected enemiesRemainingText!: Phaser.GameObjects.Text; // Enemies remaining UI element
    private waveText?: Phaser.GameObjects.Text;
    private waveTween?: Phaser.Tweens.Tween;
    private lastScore: number = 0; // Store last score for pulsing logic

    private levelManager!: LevelManager;
    private currentWorldIdx: number = 0;
    private currentLevelIdx: number = 0;

    private pauseMenuContainer?: Phaser.GameObjects.Container;
    private pauseMenuKeyHandler?: (event: KeyboardEvent) => void;
    private globalEscapeHandler?: (event: KeyboardEvent) => void;
    private isPaused: boolean = false;

    private onGameStatusChanged?: (status: string) => void;

    constructor() {
        super('GameScene');
    }

    preload() {
        // Load assets here
        // Example: this.load.image('player', 'assets/images/player.png');
        // Example: this.load.image('mob', 'assets/images/mob.png');
    }

    async create(data?: { world?: number; level?: number; levelManager?: any }) {
        // Set up the main game scene here
        // this.add.text(400, 300, 'Type Defense', {
        //     fontSize: '48px',
        //     color: '#fff',
        // }).setOrigin(0.5);

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
        // Use singleton levelManager for progress tracking
        this.levelManager = data?.levelManager || levelManager;
        this.levelManager.loadProgress();
        const world = WORLDS[this.currentWorldIdx];
        const level = world.levels[this.currentLevelIdx];
        const words = await loadWordList(level.id);
        // Create WordGenerator using allowed keys for this level
        const wordGenerator = new WordGenerator(level.availableKeys, true);
        // Pass word list to MobSpawner; MobSpawner will use these words if available
        this.mobSpawner = new MobSpawner(this, wordGenerator, level.enemySpawnRate, 2, 90, words);
        // --- Wave system integration ---
        this.mobSpawner.onWaveStart((wave) => {
            // Increase word length as waves increase
            const minLen = 2 + Math.floor(wave / 2); // e.g. every 2 waves, min increases
            const maxLen = 5 + Math.floor(wave / 2);
            if (this.mobSpawner['wordGenerator'] && typeof this.mobSpawner['wordGenerator'].setWordLengthScaling === 'function') {
                this.mobSpawner['wordGenerator'].setWordLengthScaling(minLen, maxLen);
            }
            this.showWaveNotification(wave);
        });
        this.mobSpawner.onWaveEnd(() => {
            // Start next wave after delay
            this.time.delayedCall(1500, () => this.mobSpawner.startNextWave());
        });
        this.mobSpawner.startNextWave();

        // Score and combo UI
        this.score = 0;
        this.combo = 0;
        this.scoreText = this.add.text(16, 16, 'Score: 0', { fontSize: '28px', color: '#fff', stroke: '#000', strokeThickness: 4 });
        this.comboText = this.add.text(16, 52, '', { fontSize: '22px', color: '#ff0', stroke: '#000', strokeThickness: 3 });
        this.comboText.setVisible(false);

        // Enemies remaining UI
        this.enemiesRemainingText = this.add.text(20, 80, '', {
            fontSize: '20px',
            color: '#fff',
            fontStyle: 'bold',
            stroke: '#000',
            strokeThickness: 3,
        }).setScrollFactor(0).setDepth(100);
        this.updateEnemiesRemainingUI();

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

        // Register global Escape handler
        this.globalEscapeHandler = (event: KeyboardEvent) => {
            if (event.key === 'Escape' && !this.levelCompleteText) {
                if (!this.isPaused) {
                    this.showPauseMenu();
                } else {
                    // Always call hidePauseMenu to ensure pause header is removed and game resumes
                    this.hidePauseMenu();
                }
            }
        };
        this.input.keyboard?.on('keydown', this.globalEscapeHandler);

        this.onGameStatusChanged = (status: string) => {
            if (status !== 'playing') {
                this.scene.stop();
            }
        };
        stateManager.on('gameStatusChanged', this.onGameStatusChanged);
        this.events.once('shutdown', () => {
            if (this.onGameStatusChanged) {
                stateManager.off('gameStatusChanged', this.onGameStatusChanged);
            }
        });
    }

    private showWaveNotification(wave: number) {
        if (this.waveText) this.waveText.destroy();
        this.waveText = this.add.text(400, 200, `Wave ${wave}`, {
            fontSize: '48px', color: '#fff', stroke: '#000', strokeThickness: 6, fontStyle: 'bold',
        }).setOrigin(0.5).setAlpha(0).setScale(0.7);
        this.waveTween = this.tweens.add({
            targets: this.waveText,
            alpha: 1,
            scale: 1.1,
            duration: 400,
            yoyo: true,
            hold: 1600, // Increased hold duration for longer display
            onComplete: () => {
                if (this.waveText) this.waveText.destroy();
            }
        });
    }

    updateEnemiesRemainingUI() {
        if (this.mobSpawner) {
            const remaining = this.mobSpawner.getMobs().length;
            if (remaining > 0) {
                this.enemiesRemainingText.setText(`Enemies Remaining: ${remaining}`);
                this.enemiesRemainingText.setVisible(true);
            } else {
                this.enemiesRemainingText.setVisible(false);
            }
        }
    }

    update(time: number, delta: number) {
        // Guard: wait for mobSpawner to be initialized
        if (!this.mobSpawner) return;
        // --- Difficulty scaling system ---
        this.elapsedTime += delta;
        // Progression: 0 (start) to 1 (max difficulty)
        this.scalingProgression = Phaser.Math.Clamp(this.elapsedTime / this.scalingDuration, 0, 1);
        // Adjust MobSpawner spawn rate and speed
        this.mobSpawner.setProgression(this.scalingProgression);
        // Adjust word complexity (length) as difficulty increases
        this.minWordLength = 2 + Math.floor(2 * this.scalingProgression); // 2 to 4
        this.maxWordLength = 5 + Math.floor(2 * this.scalingProgression); // 5 to 7
        if (this.mobSpawner['wordGenerator'] && typeof this.mobSpawner['wordGenerator'].setWordLengthScaling === 'function') {
            this.mobSpawner['wordGenerator'].setWordLengthScaling(this.minWordLength, this.maxWordLength);
        }
        // --- Update global gameState with delta and timestamp ---
        stateManager.updateTimestampAndDelta(time, delta);
        // Core game loop logic
        if (this.player) {
            this.player.update(time, delta);
        }
        // Update MobSpawner and its mobs
        this.mobSpawner.update(time, delta);
        // Update enemies remaining UI dynamically
        this.updateEnemiesRemainingUI();
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
        // Animate score pop-up when score changes
        if (this.scoreText && this.scoreText.visible) {
            if (this.score > this.lastScore && !this.scoreText.getData('tweened')) {
                this.scoreText.setData('tweened', true);
                this.tweens.add({
                    targets: this.scoreText,
                    scale: 1.2,
                    duration: 120,
                    yoyo: true,
                    onComplete: () => this.scoreText.setData('tweened', false)
                });
            }
            this.lastScore = this.score;
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
        const continueButton = this.add.text(400, 380, 'Continue (Enter)', {
            fontSize: '32px',
            color: '#fff',
            backgroundColor: '#007bff',
            padding: { left: 24, right: 24, top: 8, bottom: 8 },
        }).setOrigin(0.5).setInteractive({ useHandCursor: true });
        continueButton.on('pointerdown', () => this.handleContinue());
        continueButton.setInteractive({ useHandCursor: true });
        continueButton.on('pointerover', () => continueButton.setStyle({ backgroundColor: '#0056b3' }));
        continueButton.on('pointerout', () => continueButton.setStyle({ backgroundColor: '#007bff' }));
        // Add Back button
        const backButton = this.add.text(400, 440, 'Back to Level Select (Esc)', {
            fontSize: '28px',
            color: '#fff',
            backgroundColor: '#444',
            padding: { left: 24, right: 24, top: 8, bottom: 8 },
        }).setOrigin(0.5).setInteractive({ useHandCursor: true });
        backButton.on('pointerdown', () => this.handleBackToLevelSelect());

        // Keyboard navigation: Enter for continue, Esc for back
        const keyHandler = (event: KeyboardEvent) => {
            if (event.key === 'Enter') {
                this.handleContinue();
            } else if (event.key === 'Escape') {
                this.handleBackToLevelSelect();
            }
        };
        this.input.keyboard?.on('keydown', keyHandler);
        // Clean up listeners when scene shuts down
        this.events.once('shutdown', () => {
            this.input.keyboard?.off('keydown', keyHandler);
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
            this.levelManager.setCurrentLevel(world.id, world.levels[nextLevelIdx].id);
            stateManager.setGameStatus('playing');
        } else if (this.currentWorldIdx < WORLDS.length - 1) {
            // Go to first level of next world
            const nextWorld = WORLDS[this.currentWorldIdx + 1];
            this.levelManager.setCurrentLevel(nextWorld.id, nextWorld.levels[0].id);
            stateManager.setGameStatus('worldSelect');
        } else {
            // No more levels, go back to menu
            stateManager.setGameStatus('worldSelect');
        }
    }

    /**
     * Handles returning to the level selection screen for the current world
     */
    private handleBackToLevelSelect() {
        stateManager.setGameStatus('levelSelect');
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

    /**
     * Show the pause menu overlay and pause the game logic.
     */
    private showPauseMenu() {
        if (this.isPaused) return;
        this.isPaused = true;
        if (this.globalEscapeHandler) {
            this.input.keyboard?.off('keydown', this.globalEscapeHandler);
        }
        // Phaser pause menu UI (canvas)
        const bg = this.add.rectangle(400, 300, 420, 260, 0x222222, 0.95);
        let title: Phaser.GameObjects.Text | null = null;
        // Only create Phaser text if not running in Playwright (for e2e test, use DOM)
        // Use a global flag for test env detection
        const isPlaywright = typeof (window as any) !== 'undefined' && ((window as any).PLAYWRIGHT || (window as any).CYPRESS);
        if (!isPlaywright) {
            title = this.add.text(400, 220, 'Paused', { fontSize: '40px', color: '#fff' }).setOrigin(0.5);
        }
        const resumeBtn = this.add.text(400, 300, 'Resume (Esc)', {
            fontSize: '32px', color: '#fff', backgroundColor: '#007bff', padding: { left: 24, right: 24, top: 8, bottom: 8 }
        }).setOrigin(0.5).setInteractive({ useHandCursor: true });
        const quitBtn = this.add.text(400, 360, 'Quit to Menu', {
            fontSize: '28px', color: '#fff', backgroundColor: '#444', padding: { left: 24, right: 24, top: 8, bottom: 8 }
        }).setOrigin(0.5).setInteractive({ useHandCursor: true });
        resumeBtn.on('pointerdown', () => this.hidePauseMenu());
        quitBtn.on('pointerdown', () => {
            this.scene.stop();
            this.scene.start('MenuScene', { worlds: WORLDS, levelManager: this.levelManager });
        });
        const children = [bg, resumeBtn, quitBtn];
        if (title) children.push(title);
        this.pauseMenuContainer = this.add.container(0, 0, children).setDepth(200);
        // Pause game logic
        this.scene.pause();
    }

    /**
     * Hide the pause menu overlay and resume the game logic.
     */
    private hidePauseMenu() {
        if (!this.isPaused) return;
        this.isPaused = false;
        if (this.pauseMenuContainer) {
            this.pauseMenuContainer.destroy();
            this.pauseMenuContainer = undefined;
        }
        // Resume game logic
        this.scene.resume();
        // Re-register global Escape handler
        this.globalEscapeHandler = (event: KeyboardEvent) => {
            if (event.key === 'Escape' && !this.levelCompleteText) {
                if (!this.isPaused) {
                    this.showPauseMenu();
                } else {
                    // Always call hidePauseMenu to ensure pause header is removed and game resumes
                    this.hidePauseMenu();
                }
            }
        };
        this.input.keyboard?.on('keydown', this.globalEscapeHandler);
    }
}
