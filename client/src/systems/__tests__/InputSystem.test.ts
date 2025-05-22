// InputSystem.test.ts
// Unit tests for InputSystem: ensures input updates gameState as expected
import { afterEach, beforeEach, describe, expect, it } from 'vitest';
import stateManager from '../../state/stateManager';
import InputSystem from '../InputSystem';

function simulateKey(key: string) {
    const event = new KeyboardEvent('keydown', { key });
    window.dispatchEvent(event);
}

describe('InputSystem', () => {
    beforeEach(() => {
        // Reset state before each test
        stateManager.setGameStatus('playing');
        stateManager.updatePlayerInput('');
        InputSystem.getInstance().unregisterListeners();
        InputSystem.getInstance().registerListeners();
    });
    afterEach(() => {
        InputSystem.getInstance().unregisterListeners();
    });

    it('updates gameState.player.currentInput on keydown', () => {
        simulateKey('a');
        expect(stateManager.getState().player.currentInput).toBe('a');
    });

    it('toggles pause on Escape', () => {
        stateManager.setGameStatus('playing');
        simulateKey('Escape');
        expect(stateManager.getState().gameStatus).toBe('paused');
        simulateKey('Escape');
        expect(stateManager.getState().gameStatus).toBe('playing');
    });

    it('ignores input except Escape when paused', () => {
        stateManager.setGameStatus('paused');
        simulateKey('a');
        expect(stateManager.getState().player.currentInput).toBe('');
        simulateKey('Escape');
        expect(stateManager.getState().gameStatus).toBe('playing');
    });
});
