import { describe, expect, it, vi } from 'vitest';
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
// Contains AI-generated edits.
