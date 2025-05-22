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
        window.addEventListener('keydown', this.handleKeyDown);
        window.addEventListener('keyup', this.handleKeyUp);
        this.listenersRegistered = true;
    }

    public unregisterListeners(): void {
        if (!this.listenersRegistered) return;
        window.removeEventListener('keydown', this.handleKeyDown);
        window.removeEventListener('keyup', this.handleKeyUp);
        this.listenersRegistered = false;
    }

    private handleKeyDown = (event: KeyboardEvent) => {
        const gameState = stateManager.getState();
        if (gameState.gameStatus === 'paused' && event.key !== 'Escape') {
            // Ignore all input except Escape when paused
            return;
        }
        if (event.key === 'Escape') {
            // Toggle pause/play
            if (gameState.gameStatus === 'playing') {
                stateManager.setGameStatus('paused');
            } else if (gameState.gameStatus === 'paused') {
                stateManager.setGameStatus('playing');
            }
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
