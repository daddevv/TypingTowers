// Remove local window/document mocks; use global setupTests.ts for browser globals

import { describe, expect, it } from 'vitest';
import Mob from '../Mob';

// Deep mock for Phaser.GameObjects.Sprite and Phaser.GameObjects.Text for Mob tests
class MockSprite {
    anims: { on: Function };
    constructor(public scene: any, public x: number, public y: number, public texture: string) {
        this.anims = { on: () => { } };
    }
    setOrigin() { return this; }
}
class MockText {
    constructor() { return { setOrigin: () => this }; }
}
(global as any).Phaser = {
    GameObjects: {
        Sprite: MockSprite,
        Text: MockText,
    },
    Math: {
        Between: (min: number, max: number) => min, // Always pick min for deterministic tests
        Clamp: (value: number, min: number, max: number) => Math.max(min, Math.min(max, value)),
    },
};

// Helper to create a mock scene with sys.queueDepthSort
function createMockScene() {
    return {
        scale: { width: 800, height: 600 },
        add: { existing: () => { }, particles: () => { }, text: () => ({ setOrigin: () => ({}) }) },
        time: { delayedCall: () => { } },
        tweens: { add: () => { } },
        sys: {
            queueDepthSort: () => { }, // Mock as function for Phaser internals
            displayList: { add: () => { } },
            updateList: { add: () => { } },
        },
    } as any;
}

describe('Mob', () => {
    it('should initialize with the correct word', () => {
        const mob = new Mob(createMockScene(), 0, 0, 'testword', 100);
        expect(mob.word).toBe('testword');
        expect(mob.isDefeated).toBe(false);
    });

    it('should start with currentLetterIndex at 0', () => {
        const mob = new Mob(createMockScene(), 0, 0, 'abc', 100);
        expect((mob as any).currentLetterIndex).toBe(0);
    });

    it('should not be targeted by default', () => {
        const mob = new Mob(createMockScene(), 0, 0, 'abc', 100);
        expect((mob as any).isTargeted).toBe(false);
    });
});
// Contains AI-generated edits.
