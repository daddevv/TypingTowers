import { beforeEach, describe, expect, it, vi } from 'vitest';
import { defaultGameState } from '../gameState';
import { StateManager } from '../stateManager';

describe('StateManager', () => {
    let stateManager: StateManager;

    beforeEach(() => {
        // Clear localStorage and create a fresh StateManager for each test
        if (typeof localStorage !== 'undefined' && localStorage.clear) {
            localStorage.clear();
        }
        stateManager = new StateManager();
    });

    it('initializes with defaultGameState if nothing in localStorage', () => {
        expect(stateManager.getState()).toEqual(defaultGameState);
    });

    it('updates player health and emits event', () => {
        const handler = vi.fn();
        stateManager.on('playerHealthChanged', handler);
        stateManager.updatePlayerHealth(42);
        expect(stateManager.getState().player.health).toBe(42);
        expect(handler).toHaveBeenCalledWith(42);
    });

    it('adds and removes mobs correctly', () => {
        const mob = {
            id: 'mob1',
            word: 'test',
            currentTypedIndex: 0,
            position: { x: 0, y: 0 },
            speed: 100,
            type: 'normal',
            isDefeated: false
        };
        stateManager.addMob(mob);
        expect(stateManager.getState().mobs).toContainEqual(mob);
        stateManager.removeMob('mob1');
        expect(stateManager.getState().mobs).not.toContainEqual(mob);
    });

    it('sets game status and emits event', () => {
        const handler = vi.fn();
        stateManager.on('gameStatusChanged', handler);
        stateManager.setGameStatus('paused');
        expect(stateManager.getState().gameStatus).toBe('paused');
        expect(handler).toHaveBeenCalledWith('paused');
    });

    it('updates player input and emits event', () => {
        const handler = vi.fn();
        stateManager.on('playerInputChanged', handler);
        stateManager.updatePlayerInput('a');
        expect(stateManager.getState().player.currentInput).toBe('a');
        expect(handler).toHaveBeenCalledWith('a');
    });

    it('updates timestamp and delta', () => {
        stateManager.updateTimestampAndDelta(123, 16.6);
        const state = stateManager.getState();
        expect(state.timestamp).toBe(123);
        expect(state.delta).toBe(16.6);
    });

    it('saves and loads state from localStorage', () => {
        stateManager.updatePlayerHealth(99);
        const saved = localStorage.getItem('typeDefenseGameStateV2');
        expect(saved).toBeTruthy();
        // Create a new manager to test loading
        const newManager = new StateManager();
        expect(newManager.getState().player.health).toBe(99);
    });

    it('reset restores defaultGameState', () => {
        stateManager.updatePlayerHealth(1);
        stateManager.reset();
        expect(stateManager.getState()).toEqual(defaultGameState);
    });
});
