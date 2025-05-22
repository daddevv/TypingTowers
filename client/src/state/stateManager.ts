// StateManager for centralized game state management
// Provides immutable access, update functions, event subscription, and save/load

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
        this.state = this.load() || { ...defaultGameState };
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
        this.state = { ...defaultGameState };
        this.emitAndSave('reset');
    }

    // --- Internal ---
    private emitAndSave(event: string, ...args: any[]) {
        this.save();
        this.emitter.emit(event, ...args);
    }
}

// Singleton instance
const stateManager = new StateManager();
export default stateManager;

// For advanced use, export the class
export { StateManager };

