// MenuScene.ts
// World and Level selection menu with lock/unlock and local storage persistence
import Phaser from 'phaser';
import { WorldConfig, WORLDS } from '../curriculum/worldConfig';
import { levelManager } from '../managers/levelManager';
import stateManager from '../state/stateManager';
import { IRenderManager } from '../render/RenderManager';

export default class WorldSelectionScene extends Phaser.Scene {
    private worlds: WorldConfig[] = [];
    private levelManager = levelManager;
    private selectedWorld: number = 0;
    private onGameStatusChanged?: (status: string) => void;
    private unlockedWorlds: number[] = []; // Track unlocked worlds
    private renderManager!: IRenderManager;

    constructor() {
        super({ key: 'MenuScene' });
    }

    init(data?: { worlds?: WorldConfig[], levelManager?: typeof levelManager, renderManager?: IRenderManager }) {
        // Always load worlds from WORLDS if not provided
        this.worlds = (data && data.worlds) ? data.worlds : WORLDS;
        this.levelManager = (data && data.levelManager) || levelManager;
        this.renderManager = data?.renderManager || (window as any).renderManager;
        if (!this.renderManager) {
            throw new Error('RenderManager instance must be provided to WorldSelectionScene');
        }
    }

    preload() { }

    create() {
        // Remove all children to clear any lingering error messages or UI
        this.children.removeAll();
        // Always show heading for E2E selectors
        this.add.text(400, 40, 'TypeDefense', { fontSize: '40px', color: '#fff' }).setOrigin(0.5);
        // Load worlds and unlocked worlds from state
        const gameState = stateManager.getState();
        // Use curriculum.worldConfig if present, else fallback to WORLDS
        this.worlds = (gameState.curriculum && Array.isArray(gameState.curriculum.worldConfig) && gameState.curriculum.worldConfig.length > 0)
            ? gameState.curriculum.worldConfig
            : WORLDS;
        // Track unlocked worlds from progression, ensure number[]
        let unlocked = Array.isArray(gameState.progression.unlockedWorlds) ? gameState.progression.unlockedWorlds : [1];
        let unlockedWorlds: number[];
        if (unlocked.length > 0 && typeof unlocked[0] === 'string') {
            unlockedWorlds = (unlocked as string[]).map((id) => parseInt(id, 10));
        } else {
            unlockedWorlds = unlocked as number[];
        }
        this.unlockedWorlds = unlockedWorlds;

        // Remove all direct Phaser rendering code.
        // Instead, delegate to renderManager:
        this.renderManager.init(this.game.canvas.parentElement as HTMLElement);
        this.renderManager.render(stateManager.getState());

        // Add Back button
        const backButton = this.add.text(400, 500, 'Back (Esc)', {
            fontSize: '24px', color: '#fff', backgroundColor: '#333', padding: { left: 24, right: 24, top: 8, bottom: 8 }
        }).setOrigin(0.5).setInteractive({ useHandCursor: true });
        backButton.on('pointerdown', () => {
            stateManager.setGameStatus('mainMenu');
        });
        // Listen for Escape key
        this.input.keyboard?.on('keydown', (event: KeyboardEvent) => {
            if (event.key === 'Escape') {
                stateManager.setGameStatus('mainMenu');
            }
        });
        // Listen for gameStatus changes and transition if needed
        this.onGameStatusChanged = (status: string) => {
            if (status !== 'worldSelect') {
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

    handleInput(event: KeyboardEvent) {
        if (event.key === 'ArrowDown') {
            if (this.selectedWorld < this.worlds.length - 1) {
                this.selectedWorld++;
                this.renderManager.render(stateManager.getState());
            }
        } else if (event.key === 'ArrowUp') {
            if (this.selectedWorld > 0) {
                this.selectedWorld--;
                this.renderManager.render(stateManager.getState());
            }
        } else if (event.key === 'Enter') {
            this.selectWorld(this.selectedWorld);
        }
    }

    // When starting LevelMenuScene, always pass the singleton levelManager
    selectWorld(idx: number) {
        const world = this.worlds[idx];
        if (!world) {
            // eslint-disable-next-line no-console
            console.error(`[WorldSelectionScene] Selected world at index ${idx} is undefined.`);
            // Optionally, provide user feedback here
            return;
        }

        // world.id is already a number as per WorldConfig interface in worldConfig.ts
        // Update the currentWorld in the global state using the new StateManager method.
        // This ensures LevelMenuScene can retrieve it.
        stateManager.updateCurrentLevelContext({ currentWorld: world.id });

        // Optionally, clear currentLevelId if that's desired when changing worlds
        // stateManager.updateCurrentLevelContext({ currentWorld: world.id, currentLevelId: null });

        // Trigger the transition to the level selection scene for the chosen world.
        stateManager.setGameStatus('levelSelect');
    }
}
