// Mock localStorage for Node.js test environment
if (typeof global !== 'undefined' && typeof global.localStorage === 'undefined') {
    (global as any).localStorage = {
        store: {} as Record<string, string>,
        getItem(key: string) { return this.store[key] || null; },
        setItem(key: string, value: string) { this.store[key] = value; },
        removeItem(key: string) { delete this.store[key]; },
        clear() { this.store = {}; }
    };
}

import { describe, expect, it } from 'vitest';
import LevelManager from '../levelManager';

describe('LevelManager', () => {
    it('initializes with first level unlocked', () => {
        const manager = new LevelManager();
        expect(manager.isLevelUnlocked('1-1')).toBe(true);
    });

    it('sets and gets level progress', () => {
        const manager = new LevelManager();
        manager.setLevelProgress('1-2', { completed: true, highScore: 100, bestWPM: 20, bestAccuracy: 0.98, attempts: 2 });
        const progress = manager.getLevelProgress('1-2');
        expect(progress).toBeDefined();
        expect(progress!.completed).toBe(true);
        expect(progress!.highScore).toBe(100);
    });

    it('updates level progress and unlocks next level', () => {
        const manager = new LevelManager();
        manager.setLevelProgress('1-1', { completed: true, highScore: 120, bestWPM: 25, bestAccuracy: 0.99, attempts: 1 });
        expect(manager.isLevelUnlocked('1-2')).toBe(true);
    });

    it('saves and loads progress from localStorage', () => {
        const manager = new LevelManager();
        manager.setLevelProgress('1-1', { completed: true, highScore: 100, bestWPM: 20, bestAccuracy: 0.98, attempts: 2 });
        const newManager = new LevelManager();
        newManager.loadProgress();
        expect(newManager.getLevelProgress('1-1')).toBeDefined();
    });
});
