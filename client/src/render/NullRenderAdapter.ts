// NullRenderAdapter.ts
// No-op implementation of IRenderAdapter for headless/integration testing.
// All methods are empty; used to decouple engine from any real renderer.

import { GameState } from '../state/gameState';
import { IRenderAdapter } from './IRenderAdapter';

/**
 * NullRenderAdapter: No-op implementation for integration tests and headless engine runs.
 * All methods are empty and do nothing.
 */
export class NullRenderAdapter implements IRenderAdapter {
    init(width: number, height: number): void {
        // No-op
    }
    render(state: GameState): void {
        // No-op
    }
    destroy(): void {
        // No-op
    }
}
