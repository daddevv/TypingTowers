// IRenderAdapter.ts
// Defines the minimal contract for a render adapter for the engine.
// The engine must depend only on this interface, never on Phaser or other renderer classes.

import { GameState } from '../state/gameState';

export interface IRenderAdapter {
    /**
     * Initialize the renderer with the given width and height.
     */
    init(width: number, height: number): void;

    /**
     * Render the current game state.
     * @param state The current immutable game state.
     */
    render(state: GameState): void;

    /**
     * Clean up and destroy all renderer resources.
     */
    destroy(): void;
}
