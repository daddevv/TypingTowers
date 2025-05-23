import { describe, it, expect, vi } from 'vitest';
import { GameState } from '../../state/gameState';
import { IRenderManager } from '../RenderManager';

class MockRenderManager implements IRenderManager {
    public init = vi.fn();
    public render = vi.fn();
    public destroy = vi.fn();
    public resize = vi.fn();
}

describe('RenderManager abstraction', () => {
    it('calls init, render, and destroy as expected', () => {
        const mock = new MockRenderManager();
        const container = document.createElement('div');
        mock.init(container);
        expect(mock.init).toHaveBeenCalledWith(container);

        const fakeState = { player: {}, mobs: [], gameStatus: 'playing' } as unknown as GameState;
        mock.render(fakeState);
        expect(mock.render).toHaveBeenCalledWith(fakeState);

        mock.destroy();
        expect(mock.destroy).toHaveBeenCalled();
    });

    it('calls resize if implemented', () => {
        const mock = new MockRenderManager();
        mock.resize?.(800, 600);
        expect(mock.resize).toHaveBeenCalledWith(800, 600);
    });

    it('can be used as a drop-in for scenes', () => {
        // Simulate a scene using the RenderManager
        const mock = new MockRenderManager();
        const state = { player: {}, mobs: [], gameStatus: 'mainMenu' } as unknown as GameState;
        // Typical usage in a scene:
        mock.init(document.body);
        mock.render(state);
        mock.destroy();
        expect(mock.init).toHaveBeenCalled();
        expect(mock.render).toHaveBeenCalledWith(state);
        expect(mock.destroy).toHaveBeenCalled();
    });
});
