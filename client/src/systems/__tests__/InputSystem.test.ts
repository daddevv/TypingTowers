// @vitest-environment happy-dom
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
        // Reset player input after each test to avoid cross-test contamination
        stateManager.updatePlayerInput('');
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

    it('works with a mocked gameState', () => {
        // Save original getState
        const originalGetState = stateManager.getState;
        // Mock gameState
        const mockState = {
            gameStatus: 'playing',
            player: { currentInput: '' },
            // ...other properties as needed
        };
        // Mock getState to return mockState
        (stateManager.getState as any) = () => mockState;
        // Mock updatePlayerInput to update mockState
        const originalUpdatePlayerInput = stateManager.updatePlayerInput;
        stateManager.updatePlayerInput = (key: string) => {
            mockState.player.currentInput = key;
        };
        // Simulate input
        simulateKey('z');
        expect(mockState.player.currentInput).toBe('z');
        // Restore originals
        (stateManager.getState as any) = originalGetState;
        stateManager.updatePlayerInput = originalUpdatePlayerInput;
    });
});
