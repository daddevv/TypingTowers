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

    it('clears input', () => {
        const events = { once: vi.fn() };
        const keyboard = { on: vi.fn(), off: vi.fn() };
        const scene = { input: { keyboard }, events } as any;
        const handler = new InputHandler(scene);
        handler['currentInput'] = 'abc';
        handler.clearInput();
        expect(handler.getInput()).toBe('');
    });
});
// Contains AI-generated edits.
