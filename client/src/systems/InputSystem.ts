// InputSystem.ts
// Centralized input handling for TypeDefense v2
// All keyboard/mouse event listeners are registered here and update gameState via StateManager.

import stateManager from '../state/stateManager';

export default class InputSystem {
    private static instance: InputSystem;
    private listenersRegistered = false;

    private constructor() { }

    public static getInstance(): InputSystem {
        if (!InputSystem.instance) {
            InputSystem.instance = new InputSystem();
        }
        return InputSystem.instance;
    }

    public registerListeners(): void {
        if (this.listenersRegistered) return;
        if (typeof window !== 'undefined' && window.addEventListener) {
            window.addEventListener('keydown', this.handleKeyDown);
            window.addEventListener('keyup', this.handleKeyUp);
            this.listenersRegistered = true;
        }
    }

    public unregisterListeners(): void {
        if (!this.listenersRegistered) return;
        if (typeof window !== 'undefined' && window.removeEventListener) {
            window.removeEventListener('keydown', this.handleKeyDown);
            window.removeEventListener('keyup', this.handleKeyUp);
            this.listenersRegistered = false;
        }
    }

    private handleKeyDown = (event: KeyboardEvent) => {
        const gameState = stateManager.getState();
        if (gameState.gameStatus === 'paused' && event.key !== 'Escape') {
            // Ignore all input except Escape when paused
            return;
        }
        if (event.key === 'Escape') {
            // Toggle pause/play or handle back navigation
            if (gameState.gameStatus === 'playing') {
                stateManager.setGameStatus('paused');
            } else if (gameState.gameStatus === 'paused') {
                stateManager.setGameStatus('playing');
            } else if (gameState.gameStatus === 'levelComplete') {
                // Back to level select from level complete
                stateManager.setGameStatus('levelSelect');
            }
            return;
        }
        // Handle Enter for continue on level complete
        if (event.key === 'Enter' && gameState.gameStatus === 'levelComplete') {
            // Advance to next level or world, or back to menu
            // This logic should be handled by the scene listening to gameStatus change
            // Here, just set a flag or update gameStatus to trigger scene transition
            stateManager.setGameStatus('playing');
            return;
        }
        // Example: update currentInput for typing
        if (gameState.gameStatus === 'playing') {
            stateManager.updatePlayerInput(event.key);
        }
    };

    private handleKeyUp = (event: KeyboardEvent) => {
        // Optionally handle keyup events if needed
    };
}

// InputSystem is now the only place for registering input listeners.
// Call InputSystem.getInstance().registerListeners() in BootScene or on game start.
// Optionally, call unregisterListeners() on shutdown if needed.

// Usage: InputSystem.getInstance().registerListeners() in BootScene or GameScene
