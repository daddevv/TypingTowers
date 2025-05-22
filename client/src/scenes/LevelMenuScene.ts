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
    private stateUnsubscribe?: () => void;

    constructor() {
        super({ key: 'LevelMenuScene' });
    }

    init(data: { worldId?: number, levelManager?: typeof levelManager } = {}) {
        // Always fetch the latest worldId from stateManager if not provided
        let effectiveWorldId: number | undefined | null = data.worldId;
        if (typeof effectiveWorldId !== 'number') {
            const state = stateManager.getState();
            effectiveWorldId = state?.level?.currentWorld || 1;
        }
        // Defensive: fallback to 1 if invalid
        if (!WORLDS.some(w => w.id === effectiveWorldId)) {
            effectiveWorldId = 1;
        }
        this.world = WORLDS.find(w => w.id === effectiveWorldId)!;
        this.selectedLevel = 0;
        this.levelManager = (data && data.levelManager) || levelManager;
        // Defensive: clear menu items
        this.menuItems = [];
        // eslint-disable-next-line no-console
        console.log('[LevelMenuScene] Initializing with worldId:', effectiveWorldId);
    }

    create() {
        // Defensive: Remove all children and menu items
        this.children.removeAll();
        this.menuItems = [];
        // Always re-fetch world from state in case of navigation
        const state = stateManager.getState();
        const worldId = state?.level?.currentWorld || 1;
        this.world = WORLDS.find(w => w.id === worldId) || WORLDS[0];
        this.selectedLevel = 0;
        this.renderMenu();
        // Listen for gameStatus changes to stop scene if needed
        this.onGameStatusChanged = (status: string) => {
            if (status !== 'levelSelect') {
                this.scene.stop();
            }
        };
        stateManager.on('gameStatusChanged', this.onGameStatusChanged);
        // Clean up on shutdown
        this.events.once('shutdown', () => {
            if (this.onGameStatusChanged) {
                stateManager.off('gameStatusChanged', this.onGameStatusChanged);
            }
        });
        // Defensive: Listen for keyboard input
        window.addEventListener('keydown', this.handleInput.bind(this));
        this.events.once('shutdown', () => {
            window.removeEventListener('keydown', this.handleInput.bind(this));
        });
    }

    renderMenu() {
        this.menuItems.forEach(item => item.destroy());
        this.menuItems = [];
        const levels = this.world.levels;
        levels.forEach((level, idx) => {
            const y = 120 + idx * 48;
            // Always unlock 1-1, regardless of progress
            const isUnlocked = (level.id === '1-1') ? true : this.levelManager.isLevelUnlocked(level.id);
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
