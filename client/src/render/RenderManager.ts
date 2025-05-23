/**
 * RenderManager.ts
 * Abstraction/interface for rendering game state.
 * Concrete implementations (PhaserRenderManager, ThreeJsRenderManager, etc.) should implement this interface.
 */

import { GameState } from '../state/gameState';

export interface IRenderManager {
    /**
     * Initialize the renderer with any required setup.
     */
    init(container: HTMLElement): void;

    /**
     * Render the current game state.
     * @param state The current immutable game state.
     */
    render(state: GameState): void;

    /**
     * Clean up and destroy all renderer resources.
     */
    destroy(): void;

    /**
     * Optionally, handle resize events.
     */
    resize?(width: number, height: number): void;
}

// Example stub for a concrete implementation (Phaser, Three.js, etc.)
// export class PhaserRenderManager implements IRenderManager { ... }
