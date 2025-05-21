import { describe, expect, it } from 'vitest';
import { FingerType } from '../../curriculum/fingerGroups';
import FingerGroupManager from '../fingerGroupManager';

describe('FingerGroupManager', () => {
    it('initializes stats for all finger types', () => {
        const manager = new FingerGroupManager();
        expect(manager.getFingerStats(FingerType.LEFT_INDEX)).toBeDefined();
        expect(manager.getFingerStats(FingerType.RIGHT_INDEX)).toBeDefined();
    });

    it('records key presses and updates stats', () => {
        const manager = new FingerGroupManager();
        manager.recordKeyPress('f', true, 1000);
        manager.recordKeyPress('f', false, 2000);
        const stats = manager.getFingerStats(FingerType.LEFT_INDEX)!;
        expect(stats.totalKeyPresses).toBe(2);
        expect(stats.correctFingerUses).toBe(1);
        expect(stats.mistypedKeys).toBe(1);
        expect(stats.accuracy).toBeCloseTo(0.5);
    });

    it('calculates average speed for a finger', () => {
        const manager = new FingerGroupManager();
        manager.recordKeyPress('f', true, 1000);
        manager.recordKeyPress('f', true, 1600);
        manager.recordKeyPress('f', true, 2200);
        const stats = manager.getFingerStats(FingerType.LEFT_INDEX)!;
        expect(stats.averageSpeed).toBeGreaterThan(0);
    });

    it('returns correct keys for a finger', () => {
        const manager = new FingerGroupManager();
        const keys = manager.getKeysForFinger(FingerType.LEFT_INDEX);
        expect(keys).toContain('f');
        expect(keys).toContain('g');
    });

    it('determines if a key is mastered', () => {
        const manager = new FingerGroupManager();
        // Simulate high accuracy and fast speed
        for (let i = 0; i < 20; i++) {
            manager.recordKeyPress('f', true, 1000 + i * 100);
        }
        expect(manager.isKeyMastered('f')).toBe(true);
    });
});
// Contains AI-generated edits.
