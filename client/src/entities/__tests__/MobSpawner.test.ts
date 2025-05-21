import { beforeEach, describe, expect, it, vi } from 'vitest';
import Mob from '../Mob';
import MobSpawner from '../MobSpawner';

describe('MobSpawner', () => {
    it('spawns mobs at correct interval and y position', () => {
        const scene = {
            scale: { width: 800, height: 600 },
            add: { existing: vi.fn(), particles: vi.fn() },
            time: { delayedCall: vi.fn() },
            tweens: { add: vi.fn() },
        } as any;
        const words = ['f', 'j'];
        const spawner = new MobSpawner(scene, words, 1000, 2, 90);
        spawner.update(0, 1000);
        const mobs = spawner.getMobs();
        expect(mobs.length).toBe(2);
        expect(mobs[0].y).toBeGreaterThanOrEqual(100);
        expect(mobs[0].y).toBeLessThanOrEqual(500);
    });

    it('removes mobs correctly', () => {
        const scene = {
            scale: { width: 800, height: 600 },
            add: { existing: vi.fn(), particles: vi.fn() },
            time: { delayedCall: vi.fn() },
            tweens: { add: vi.fn() },
        } as any;
        const words = ['f', 'j'];
        const spawner = new MobSpawner(scene, words, 1000, 1, 90);
        spawner.update(0, 1000);
        const mob = spawner.getMobs()[0];
        spawner.removeMob(mob);
        expect(spawner.getMobs()).not.toContain(mob);
    });
});

describe('MobSpawner Integration', () => {
    let scene: any;
    beforeEach(() => {
        scene = {
            scale: { width: 800, height: 600 },
            add: { existing: vi.fn(), particles: vi.fn(), text: vi.fn(() => ({ setOrigin: vi.fn().mockReturnThis() })) },
            time: { delayedCall: vi.fn() },
            tweens: { add: vi.fn() },
        };
    });

    it('spawns mobs over multiple intervals and prevents overlap', () => {
        const words = ['foo', 'bar', 'baz', 'qux'];
        const spawner = new MobSpawner(scene, words, 500, 2, 100);
        // Simulate 3 spawn cycles (should spawn 6 mobs)
        for (let i = 0; i < 3; i++) {
            spawner.update(0, 500);
        }
        const mobs = spawner.getMobs();
        expect(mobs.length).toBe(6);
        // Check that mobs have unique y positions (no overlap)
        const yPositions = mobs.map(m => m.y);
        const uniqueY = new Set(yPositions);
        expect(uniqueY.size).toBe(mobs.length);
        // Check that all mobs are instances of Mob
        mobs.forEach(mob => expect(mob).toBeInstanceOf(Mob));
    });

    it('scales spawn interval and mob speed with progression', () => {
        const words = ['foo', 'bar'];
        const spawner = new MobSpawner(scene, words, 1000, 1, 90);
        spawner.setProgression(1); // Max progression
        expect(spawner['spawnInterval']).toBeLessThan(1000);
        expect(spawner['mobBaseSpeed']).toBeGreaterThan(90);
    });

    it('removes mobs and updates mob list', () => {
        const words = ['foo'];
        const spawner = new MobSpawner(scene, words, 100, 1, 90);
        spawner.update(0, 100);
        const mob = spawner.getMobs()[0];
        spawner.removeMob(mob);
        expect(spawner.getMobs()).not.toContain(mob);
    });
});
// Contains AI-generated edits.
