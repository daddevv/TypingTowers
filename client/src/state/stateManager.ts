// StateManager for centralized game state management
// Provides immutable access, update functions, event subscription, and save/load

import { LevelConfig, WORLDS } from '../curriculum/worldConfig'; // Added WORLDS and LevelConfig
import { GameState, defaultGameState } from './gameState';

// Simple event emitter for state change notifications
class EventEmitter {
    private listeners: Record<string, Array<(...args: any[]) => void>> = {};

    on(event: string, handler: (...args: any[]) => void) {
        if (!this.listeners[event]) this.listeners[event] = [];
        this.listeners[event].push(handler);
    }

    off(event: string, handler: (...args: any[]) => void) {
        if (!this.listeners[event]) return;
        this.listeners[event] = this.listeners[event].filter(h => h !== handler);
    }

    emit(event: string, ...args: any[]) {
        if (!this.listeners[event]) return;
        for (const handler of this.listeners[event]) handler(...args);
    }
}

class StateManager {
    private state: GameState;
    private emitter: EventEmitter;
    private static STORAGE_KEY = 'typeDefenseGameStateV2';

    constructor() {
        this.state = this.load() || JSON.parse(JSON.stringify(defaultGameState));
        this.emitter = new EventEmitter();
        // Expose for debugging
        (window as any).gameState = this.state;
    }

    // Get a deep copy of the current state (to prevent direct mutation)
    getState(): GameState {
        return JSON.parse(JSON.stringify(this.state));
    }

    // Subscribe to state changes
    on(event: string, handler: (...args: any[]) => void) {
        this.emitter.on(event, handler);
    }

    off(event: string, handler: (...args: any[]) => void) {
        this.emitter.off(event, handler);
    }

    // --- State update functions ---
    updatePlayerHealth(newHealth: number) {
        this.state.player.health = newHealth;
        this.emitAndSave('playerHealthChanged', newHealth);
    }

    addMob(mob: GameState['mobs'][number]) {
        this.state.mobs.push(mob);
        this.emitAndSave('mobAdded', mob);
    }

    removeMob(mobId: string) {
        this.state.mobs = this.state.mobs.filter(m => m.id !== mobId);
        this.emitAndSave('mobRemoved', mobId);
    }

    setGameStatus(status: GameState['gameStatus']) {
        this.state.gameStatus = status;
        this.emitAndSave('gameStatusChanged', status);
    }

    updatePlayerInput(input: string) {
        this.state.player.currentInput = input;
        this.emitAndSave('playerInputChanged', input);
    }

    // --- Add update methods for delta and timestamp ---
    updateTimestampAndDelta(timestamp: number, delta: number) {
        this.state.timestamp = timestamp;
        this.state.delta = delta;
        this.emitAndSave('timestampDeltaChanged', { timestamp, delta });
    }

    updateMobSpawnerState(newState: Partial<GameState['mobSpawner']>) {
        this.state.mobSpawner = { ...this.state.mobSpawner, ...newState };
        this.emitAndSave('mobSpawnerStateChanged', this.state.mobSpawner);
    }

    updateMobs(newMobs: GameState['mobs']) {
        this.state.mobs = newMobs;
        this.emitAndSave('mobsUpdated', newMobs);
    }

    // --- Level Context (Current World/Level) ---
    updateCurrentLevelContext(context: Partial<GameState['level']>) {
        // Ensure this.state.level is initialized if it's not already
        if (!this.state.level) {
            this.state.level = { currentWorld: null, currentLevelId: null, levelStatus: 'notStarted', ...context };
        } else {
            this.state.level = { ...this.state.level, ...context };
        }
        this.emitAndSave('levelContextChanged', this.state.level);

        if (
            this.state.level.levelStatus === 'complete' &&
            this.state.level.currentWorld != null &&
            this.state.level.currentLevelId != null
        ) {
            this.unlockNextLevel(this.state.level.currentWorld, this.state.level.currentLevelId);
        }
    }

    // --- Level Progression ---
    updateProgression(newProgression: Partial<GameState['progression']>) {
        this.state.progression = { ...this.state.progression, ...newProgression };
        this.emitAndSave('progressionChanged', this.state.progression);
    }

    // --- Curriculum/Finger Group Stats ---
    updateCurriculum(newCurriculum: Partial<GameState['curriculum']>) {
        this.state.curriculum = { ...this.state.curriculum, ...newCurriculum };
        this.emitAndSave('curriculumChanged', this.state.curriculum);
    }

    // Add more update methods as needed for other state parts...

    // --- Save/load ---
    save() {
        try {
            localStorage.setItem(StateManager.STORAGE_KEY, JSON.stringify(this.state));
        } catch (e) {
            // eslint-disable-next-line no-console
            console.warn('Failed to save game state:', e);
        }
    }

    load(): GameState | null {
        try {
            const raw = localStorage.getItem(StateManager.STORAGE_KEY);
            if (!raw) return null;
            return JSON.parse(raw) as GameState;
        } catch (e) {
            // eslint-disable-next-line no-console
            console.warn('Failed to load game state:', e);
            return null;
        }
    }

    reset() {
        this.state = JSON.parse(JSON.stringify(defaultGameState));
        this.emitAndSave('reset');
    }

    // Reset the entire game state to default values
    resetState() {
        this.state = JSON.parse(JSON.stringify(defaultGameState));
        this.emitAndSave('stateReset', this.state);
    }

    // --- Internal ---
    private emitAndSave(event: string, ...args: any[]) {
        this.save();
        this.emitter.emit(event, ...args);
    }

    private unlockNextLevel(worldId: number, completedLevelId: string) {
        const worldConfig = WORLDS.find(w => w.id === worldId);
        if (!worldConfig) {
            console.warn(`[StateManager] unlockNextLevel: World config not found for worldId: ${worldId}`);
            return;
        }

        const completedLevelIndex = worldConfig.levels.findIndex(l => l.id === completedLevelId);
        if (completedLevelIndex === -1) {
            console.warn(`[StateManager] unlockNextLevel: Level config not found for levelId: ${completedLevelId} in worldId: ${worldId}`);
            return;
        }

        let nextLevelToUnlock: LevelConfig | undefined = undefined;

        // Try to find next level in the same world
        if (completedLevelIndex < worldConfig.levels.length - 1) {
            nextLevelToUnlock = worldConfig.levels[completedLevelIndex + 1];
        } else {
            // Try to find first level of the next world
            const nextWorldConfig = WORLDS.find(w => w.id === worldId + 1);
            if (nextWorldConfig && nextWorldConfig.levels.length > 0) {
                nextLevelToUnlock = nextWorldConfig.levels[0];
                // Unlock the next world as well
                const nextWorldIdStr = String(nextWorldConfig.id);
                if (!this.state.progression.unlockedWorlds.includes(nextWorldIdStr)) {
                    this.updateProgression({
                        unlockedWorlds: [...this.state.progression.unlockedWorlds, nextWorldIdStr]
                    });
                }
            }
        }

        if (nextLevelToUnlock) {
            if (!this.state.progression.unlockedLevels.includes(nextLevelToUnlock.id)) {
                this.updateProgression({
                    unlockedLevels: [...this.state.progression.unlockedLevels, nextLevelToUnlock.id]
                });
                console.log(`[StateManager] Unlocked level: ${nextLevelToUnlock.id}`);
            }
        } else {
            console.log(`[StateManager] No next level to unlock after ${completedLevelId}. End of curriculum?`);
        }
    }
}

// Singleton instance
const stateManager = new StateManager();
export default stateManager;

// For advanced use, export the class
export { StateManager };

