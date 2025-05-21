import { describe, expect, it, vi } from 'vitest';
import InputHandler from '../InputHandler';

describe('InputHandler', () => {
    it('accumulates input on keydown', () => {
        // Mock Phaser.Scene and input
        const events = { once: vi.fn() };
        const keyboard = { on: vi.fn(), off: vi.fn() };
        const scene = { input: { keyboard }, events } as any;
        const handler = new InputHandler(scene);
        // Simulate keydown event
        const keydownCallback = keyboard.on.mock.calls[0][1];
        keydownCallback({ key: 'a' });
        expect(handler.getInput()).toBe('a');
    });

    it('accumulates multiple keydowns', () => {
        const events = { once: vi.fn() };
        const keyboard = { on: vi.fn(), off: vi.fn() };
        const scene = { input: { keyboard }, events } as any;
        const handler = new InputHandler(scene);
        const keydownCallback = keyboard.on.mock.calls[0][1];
        keydownCallback({ key: 'a' });
        keydownCallback({ key: 'b' });
        keydownCallback({ key: 'c' });
        expect(handler.getInput()).toBe('abc');
    });

    it('clears input', () => {
        const events = { once: vi.fn() };
        const keyboard = { on: vi.fn(), off: vi.fn() };
        const scene = { input: { keyboard }, events } as any;
        const handler = new InputHandler(scene);
        handler['currentInput'] = 'abc';
        handler.clearInput();
        expect(handler.getInput()).toBe('');
    });

    it('does not register events if keyboard is missing', () => {
        const events = { once: vi.fn() };
        const scene = { input: {}, events } as any;
        expect(() => new InputHandler(scene)).not.toThrow();
    });

    it('removes keydown listener on shutdown and destroy', () => {
        const events = { once: vi.fn((event, cb) => cb()) };
        const keyboard = { on: vi.fn(), off: vi.fn() };
        const scene = { input: { keyboard }, events } as any;
        new InputHandler(scene);
        // Should call off twice (shutdown and destroy)
        expect(keyboard.off).toHaveBeenCalledTimes(2);
    });
});
// Contains AI-generated edits.
