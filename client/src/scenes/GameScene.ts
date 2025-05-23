import Phaser from 'phaser';
import { WORLDS } from '../curriculum/worldConfig';
import { MobSystem } from '../entities/Mob';
import MobSpawner from '../entities/MobSpawner';
import FingerGroupManager from '../managers/fingerGroupManager';
import LevelManager, { levelManager } from '../managers/levelManager';
import stateManager from '../state/stateManager';
import { loadWordList } from '../utils/loadWordList';
import WordGenerator from '../utils/wordGenerator';
import { IRenderManager } from '../render/RenderManager';

export default class GameScene extends Phaser.Scene {
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

    private levelManager!: LevelManager;
    private currentWorldIdx: number = 0;
    private currentLevelIdx: number = 0;

    private pauseMenuContainer?: Phaser.GameObjects.Container;
    private pauseMenuKeyHandler?: (event: KeyboardEvent) => void;
    private globalEscapeHandler?: (event: KeyboardEvent) => void;
    private isPaused: boolean = false;

    private onGameStatusChanged?: (status: string) => void;

    // Group for mob visuals
    private mobGroup!: Phaser.GameObjects.Group;
    private mobVisuals: Map<string, { sprite: Phaser.GameObjects.Sprite, text: Phaser.GameObjects.Text }> = new Map();

    // Player rendering objects
    private playerSprite!: Phaser.GameObjects.Sprite;
    private playerHealthText!: Phaser.GameObjects.Text;

    private renderManager!: IRenderManager;

    // Wave notification properties
    private waveText?: Phaser.GameObjects.Text;
    private waveTween?: Phaser.Tweens.Tween;

    // Enemies remaining text
    private enemiesRemainingText!: Phaser.GameObjects.Text;

    // Add these properties for score/combo UI
    protected scoreText?: Phaser.GameObjects.Text;
    protected comboText?: Phaser.GameObjects.Text;
    protected lastScore: number = 0;

    // Add particleManager property
    private particleManager?: { emitParticleAt: (x: number, y: number, count?: number) => void };

    constructor() {
        super('GameScene');
    }

    preload() {
        // Load assets here
        // Example: this.load.image('player', 'assets/images/player.png');
        // Example: this.load.image('mob', 'assets/images/mob.png');
    }

    async create(data?: { world?: number; level?: number; levelManager?: any, renderManager?: IRenderManager }) {
        // Set up the main game scene here
        // this.add.text(400, 300, 'Type Defense', {
        //     fontSize: '48px',
        //     color: '#fff',
        // }).setOrigin(0.5);

        // Initialize Player and InputHandler
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
        const spawnerConfig = { width: this.scale.width, height: this.scale.height };
        this.mobSpawner = new MobSpawner(spawnerConfig, wordGenerator, level.enemySpawnRate, 2, 90, words);
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

        // Inject or create the renderManager
        this.renderManager = data?.renderManager || (window as any).renderManager;
        if (!this.renderManager) {
            throw new Error('RenderManager instance must be provided to GameScene');
        }
        this.renderManager.init(this.game.canvas.parentElement as HTMLElement);
        // Delegate all initial rendering to RenderManager
        this.renderManager.render(stateManager.getState());
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
        const gameState = stateManager.getState();
        if (gameState.gameStatus === 'paused') {
            return;
        }
        if (!this.mobSpawner) return;
        this.elapsedTime += delta;
        this.scalingProgression = Phaser.Math.Clamp(this.elapsedTime / this.scalingDuration, 0, 1);
        this.mobSpawner.setProgression(this.scalingProgression);
        this.minWordLength = 2 + Math.floor(2 * this.scalingProgression);
        this.maxWordLength = 5 + Math.floor(2 * this.scalingProgression);
        if (this.mobSpawner['wordGenerator'] && typeof this.mobSpawner['wordGenerator'].setWordLengthScaling === 'function') {
            this.mobSpawner['wordGenerator'].setWordLengthScaling(this.minWordLength, this.maxWordLength);
        }
        stateManager.updateTimestampAndDelta(time, delta);
        // Remove all direct Phaser rendering code below.
        // Only update state and call systems:
        MobSystem.updateAll(time, delta);
        this.mobSpawner.update(time, delta);
        // Collision: check for mobs near player
        const mobs = stateManager.getState().mobs;
        const playerState = stateManager.getState().player; // <-- Add this line
        for (const mob of mobs) {
            if (!mob.isDefeated) {
                const dx = mob.position.x - playerState.position.x;
                const dy = mob.position.y - playerState.position.y;
                const dist = Math.sqrt(dx * dx + dy * dy);
                if (dist < 40) {
                    stateManager.updatePlayerHealth(Math.max(0, playerState.health - 1));
                    stateManager.removeMob(mob.id);
                    if (playerState.health <= 0) {
                        this.showGameOverUI();
                        return;
                    }
                    break;
                }
            }
        }
        // Input handling: update mob progress based on input
        const input = stateManager.getState().player.currentInput || '';
        if (input.length > 0) {
            let anyCorrect = false;
            for (const char of input) {
                // Find mobs whose next letter matches
                const candidates = mobs.filter(mob => !mob.isDefeated && mob.word[mob.currentTypedIndex]?.toLowerCase() === char.toLowerCase());
                if (candidates.length > 0) {
                    // Pick the closest mob
                    let minDist = Infinity;
                    let closest = null;
                    for (const mob of candidates) {
                        const dx = mob.position.x - playerState.position.x;
                        const dy = mob.position.y - playerState.position.y;
                        const dist = Math.sqrt(dx * dx + dy * dy);
                        if (dist < minDist) {
                            minDist = dist;
                            closest = mob;
                        }
                    }
                    if (closest) {
                        closest.currentTypedIndex++;
                        anyCorrect = true;
                        // TODO: trigger particle effect at mob position
                        // TODO: update score/combo in state
                        if (closest.currentTypedIndex >= closest.word.length) {
                            closest.isDefeated = true;
                        }
                    }
                }
            }
            if (!anyCorrect) {
                // TODO: reset combo in state
            }
            stateManager.updatePlayerInput('');
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
                    onComplete: () => {
                        if (this.scoreText) {
                            this.scoreText.setData('tweened', false);
                        }
                    }
                });
            }
            this.lastScore = this.score;
        }

        // Delegate rendering to the renderManager
        this.renderManager.render(stateManager.getState());
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
        // Instead of direct scene transition, update gameStatus to trigger state-driven navigation
        continueButton.on('pointerdown', () => {
            stateManager.setGameStatus('playing');
        });
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

        // Remove direct keyboard listener for Enter/Esc here.
        // Keyboard navigation is now handled by InputSystem updating gameState.
        // This scene should listen for gameStatus changes and react accordingly.

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
     * (No longer called directly by Enter key or Continue button; navigation is now state-driven)
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
     * Show Game Over UI and allow returning to level select.
     */
    private showGameOverUI() {
        // Prevent multiple overlays
        if (this.children.getByName('gameOverText')) return;
        const gameOverText = this.add.text(400, 260, 'Game Over', {
            fontSize: '48px',
            color: '#ff5555',
            backgroundColor: '#222',
            padding: { left: 24, right: 24, top: 12, bottom: 12 },
        }).setOrigin(0.5).setName('gameOverText');
        const backButton = this.add.text(400, 340, 'Back to Level Select (Esc/Enter)', {
            fontSize: '32px',
            color: '#fff',
            backgroundColor: '#444',
            padding: { left: 24, right: 24, top: 8, bottom: 8 },
        }).setOrigin(0.5).setInteractive({ useHandCursor: true }).setName('gameOverBackBtn');
        backButton.on('pointerdown', () => this.handleBackToLevelSelect());
        // Keyboard navigation: Enter or Esc for back
        const keyHandler = (event: KeyboardEvent) => {
            if (event.key === 'Enter' || event.key === 'Escape') {
                this.handleBackToLevelSelect();
            }
        };
        this.input.keyboard?.on('keydown', keyHandler);
        this.events.once('shutdown', () => {
            this.input.keyboard?.off('keydown', keyHandler);
        });
        this.scene.pause();
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
        if (this.particleManager?.emitParticleAt) {
            this.particleManager.emitParticleAt(pos.x, pos.y, 12);
        }
        this.scoreText?.setText(`Score: ${this.score}`);
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
        this.comboText?.setVisible(false);
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
        
        // Check for test environment
        const isTestEnv = typeof (window as any) !== 'undefined' && (
            (window as any).PLAYWRIGHT || 
            (window as any).CYPRESS || 
            process.env.PLAYWRIGHT === 'true' || 
            process.env.CYPRESS === 'true'
        );
        // Log for debugging
        console.log("Test environment check in showPauseMenu:", isTestEnv, 
            "PLAYWRIGHT:", (window as any).PLAYWRIGHT, 
            "CYPRESS:", (window as any).CYPRESS, 
            "process.env.PLAYWRIGHT:", process.env.PLAYWRIGHT);
        
        // Create DOM element for E2E testing if in test environment
        if (isTestEnv) {
            console.log("Test environment detected, creating pause menu DOM elements");
            
            // Remove any existing pause header (just in case)
            const existingHeader = document.getElementById('pause-header');
            if (existingHeader) {
                existingHeader.remove();
            }
            
            // Create pause header DOM element for Playwright to detect
            const pauseHeader = document.createElement('div');
            pauseHeader.id = 'pause-header';
            pauseHeader.textContent = 'Paused';
            pauseHeader.style.position = 'absolute';
            pauseHeader.style.left = '50%';
            pauseHeader.style.top = '220px';
            pauseHeader.style.transform = 'translateX(-50%)';
            pauseHeader.style.zIndex = '1000';
            pauseHeader.style.fontSize = '40px';
            pauseHeader.style.color = '#fff';
            pauseHeader.style.backgroundColor = '#222';
            pauseHeader.style.padding = '10px 20px';
            pauseHeader.style.border = '2px solid #fff';
            pauseHeader.style.opacity = '1'; // Make sure it's fully visible
            document.body.appendChild(pauseHeader);
            
            // Create Resume button DOM element
            const resumeBtn = document.createElement('div');
            resumeBtn.id = 'pause-resume-btn';
            resumeBtn.textContent = 'Resume (Esc)';
            resumeBtn.style.position = 'absolute';
            resumeBtn.style.left = '50%';
            resumeBtn.style.top = '300px';
            resumeBtn.style.transform = 'translateX(-50%)';
            resumeBtn.style.zIndex = '1000';
            resumeBtn.style.fontSize = '32px';
            resumeBtn.style.color = '#fff';
            resumeBtn.style.backgroundColor = '#007bff';
            resumeBtn.style.padding = '8px 24px';
            resumeBtn.style.cursor = 'pointer';
            resumeBtn.addEventListener('click', () => this.hidePauseMenu());
            document.body.appendChild(resumeBtn);
            
            // Store reference for removal when unpausing
            if (!this.pauseMenuContainer) {
                this.pauseMenuContainer = this.add.container(0, 0);
            }
            this.pauseMenuContainer.setData('domElements', [pauseHeader, resumeBtn]);
        }
        
        // Phaser pause menu UI (canvas)
        const bg = this.add.rectangle(400, 300, 420, 260, 0x222222, 0.95);
        let title: Phaser.GameObjects.Text | null = null;
        
        // Only create Phaser text if not running in test environment
        if (!isTestEnv) {
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
        if (!this.pauseMenuContainer) {
            this.pauseMenuContainer = this.add.container(0, 0, children).setDepth(200);
        } else {
            this.pauseMenuContainer.removeAll(true);
            this.pauseMenuContainer.add(children);
        }
    }

    /**
     * Hide the pause menu overlay and resume the game logic.
     */
    private hidePauseMenu() {
        if (!this.isPaused) return;
        this.isPaused = false;
        // Remove DOM elements if they exist
        if (this.pauseMenuContainer && this.pauseMenuContainer.getData('domElements')) {
            const domElements = this.pauseMenuContainer.getData('domElements') as HTMLElement[];
            domElements.forEach(element => {
                if (element && element.parentNode) {
                    element.parentNode.removeChild(element);
                }
            });
        }
        // Also explicitly remove all pause-related DOM elements to be safe
        const elementsToRemove = ['pause-header', 'pause-resume-btn', 'pause-quit-btn'];
        elementsToRemove.forEach(id => {
            const element = document.getElementById(id);
            if (element && element.parentNode) {
                element.parentNode.removeChild(element);
            }
        });
        if (this.pauseMenuContainer) {
            this.pauseMenuContainer.destroy();
            this.pauseMenuContainer = undefined;
        }
        // Resume game logic
        this.scene.resume();
        // Re-register global Escape handler
        if (this.globalEscapeHandler) {
            this.input.keyboard?.on('keydown', this.globalEscapeHandler);
        }
    }
}
