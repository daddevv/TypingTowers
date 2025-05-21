import { describe, expect, it } from 'vitest';
import Mob from '../Mob';

// Mock Phaser.GameObjects.Sprite for Mob inheritance
class MockSprite {
    constructor(public scene: any, public x: number, public y: number, public texture: string) { }
}
(global as any).Phaser = { GameObjects: { Sprite: MockSprite } };

describe('Mob', () => {
    it('should initialize with the correct word', () => {
        const mob = new Mob({} as any, 0, 0, 'testword', 100);
        expect(mob.word).toBe('testword');
        expect(mob.isDefeated).toBe(false);
    });

    it('should start with currentLetterIndex at 0', () => {
        const mob = new Mob({} as any, 0, 0, 'abc', 100);
        expect((mob as any).currentLetterIndex).toBe(0);
    });

    it('should not be targeted by default', () => {
        const mob = new Mob({} as any, 0, 0, 'abc', 100);
        expect((mob as any).isTargeted).toBe(false);
    });
});
// Contains AI-generated edits.
