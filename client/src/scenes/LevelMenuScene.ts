// LevelMenuScene.ts
// Scene for selecting a level within a selected world
import Phaser from 'phaser';
import { WORLDS, WorldConfig } from '../curriculum/worldConfig';
import { levelManager } from '../managers/levelManager';
import stateManager from '../state/stateManager';

export default class LevelMenuScene extends Phaser.Scene {
    private world!: WorldConfig;
    private menuItems: Phaser.GameObjects.Text[] = [];
    private selectedLevel: number = 0;
    private levelManager = levelManager;
    private onGameStatusChanged?: (status: string) => void;

    constructor() {
        super({ key: 'LevelMenuScene' });
    }

    init(data: { worldId?: number, levelManager?: typeof levelManager } = {}) {
        let effectiveWorldId: number | undefined | null = data.worldId;

        // If worldId is not passed directly (e.g., when BootScene starts this scene based on gameStatus),
        // retrieve it from the global state.
        if (typeof effectiveWorldId !== 'number') {
            const state = stateManager.getState();
            // Ensure state.level and state.level.currentWorld exist and are of the expected type.
            // GameState['level']['currentWorld'] should be number | null.
            effectiveWorldId = state.level?.currentWorld;
        }

        // eslint-disable-next-line no-console
        console.log('[LevelMenuScene] Initializing with worldId:', effectiveWorldId);

        if (typeof effectiveWorldId !== 'number') {
            // eslint-disable-next-line no-console
            console.error('[LevelMenuScene] Invalid or missing worldId. Cannot initialize scene properly.');
            // Setting this.world to undefined will be caught by the check in create()
            this.world = undefined!;
            this.levelManager = data.levelManager || levelManager;
            return;
        }

        this.world = WORLDS.find(w => w.id === effectiveWorldId)!;

        if (!this.world) {
        // eslint-disable-next-line no-console
            console.error(`[LevelMenuScene] World configuration not found for ID: ${effectiveWorldId}.`);
            // this.world will be undefined, and the create() method's check will handle UI.
        }

        this.levelManager = data.levelManager || levelManager;
    }

    create() {
        // Defensive: Check for valid world data
        if (!this.world || !this.world.levels || this.world.levels.length === 0) {
            this.add.text(400, 300, 'No levels available. Please check game data.', { fontSize: '24px', color: '#f00' }).setOrigin(0.5);
            // Add Back button even if no levels
            const backButton = this.add.text(400, 500, 'Back (Esc)', {
                fontSize: '24px', color: '#fff', backgroundColor: '#333', padding: { left: 24, right: 24, top: 8, bottom: 8 }
            }).setOrigin(0.5).setInteractive({ useHandCursor: true });
            backButton.on('pointerdown', () => stateManager.setGameStatus('worldSelect'));
            this.input.keyboard?.on('keydown', (event: KeyboardEvent) => {
                if (event.key === 'Escape') {
                    stateManager.setGameStatus('worldSelect');
                }
            });
            return;
        }
        this.add.text(400, 40, this.world.name, { fontSize: '36px', color: '#fff' }).setOrigin(0.5);
        this.renderMenu();
        // Add Back button
        const backButton = this.add.text(400, 500, 'Back (Esc)', {
            fontSize: '24px', color: '#fff', backgroundColor: '#333', padding: { left: 24, right: 24, top: 8, bottom: 8 }
        }).setOrigin(0.5).setInteractive({ useHandCursor: true });
        backButton.on('pointerdown', () => stateManager.setGameStatus('worldSelect'));
        // Listen for Escape key
        this.input.keyboard?.on('keydown', (event: KeyboardEvent) => {
            if (event.key === 'Escape') {
                stateManager.setGameStatus('worldSelect');
            }
        });
        // Re-render menu when scene is resumed (e.g., after completing a level)
        this.events.on('resume', () => {
            this.renderMenu();
        });
        // Refresh menu when scene is woken (e.g., after returning from GameScene)
        this.events.on('wake', () => {
            this.renderMenu();
        });
        // Listen for gameStatus changes and transition if needed
        this.onGameStatusChanged = (status: string) => {
            if (status !== 'levelSelect') {
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

    renderMenu() {
        this.menuItems.forEach(item => item.destroy());
        this.menuItems = [];
        const levels = this.world.levels;
        levels.forEach((level, idx) => {
            const y = 120 + idx * 48;
            const isUnlocked = this.levelManager.isLevelUnlocked(level.id);
            const color = isUnlocked ? '#fff' : '#888';
            const item = this.add.text(400, y, `${level.name}${isUnlocked ? '' : ' (Locked)'}`, {
                fontSize: '28px',
                color,
                backgroundColor: idx === this.selectedLevel ? '#444' : undefined,
                padding: { left: 12, right: 12, top: 4, bottom: 4 },
            }).setOrigin(0.5).setInteractive({ useHandCursor: isUnlocked });
            if (isUnlocked) {
                item.on('pointerdown', () => this.selectLevel(idx));
            }
            this.menuItems.push(item);
        });
    }

    handleInput(event: KeyboardEvent) {
        const levels = this.world.levels;
        if (event.key === 'ArrowUp') {
            this.selectedLevel = (this.selectedLevel - 1 + levels.length) % levels.length;
            this.renderMenu();
        } else if (event.key === 'ArrowDown') {
            this.selectedLevel = (this.selectedLevel + 1) % levels.length;
            this.renderMenu();
        } else if (event.key === 'Enter') {
            this.selectLevel(this.selectedLevel);
        } else if (event.key === 'Escape') {
            stateManager.setGameStatus('worldSelect');
        }
    }

    selectLevel(idx: number) {
        const level = this.world.levels[idx];
        if (!this.levelManager.isLevelUnlocked(level.id)) return;
        this.levelManager.setCurrentLevel(this.world.id, level.id);
        stateManager.setGameStatus('playing');
        // ...could also update current level in stateManager here...
    }
}
