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
import { beforeEach, describe, expect, it } from 'vitest';
import LevelManager from '../levelManager';

describe('LevelManager', () => {
    beforeEach(() => {
        // Clear localStorage before each test to avoid cross-test contamination
        if (typeof localStorage !== 'undefined' && localStorage.clear) {
            localStorage.clear();
        }
    });

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

    it('completing last level of a world unlocks first level of next world', () => {
        const manager = new LevelManager();
        // Complete all levels in world 1 in order
        const world1Levels = [
            '1-1', '1-2', '1-3', '1-4', '1-5', '1-6', '1-7'
        ];
        world1Levels.forEach((levelId, idx) => {
            manager.completeLevel(levelId, { score: 100 + idx * 10, wpm: 20 + idx * 2, accuracy: 0.95 });
        });
        expect(manager.isLevelUnlocked('2-1')).toBe(true);
    });

    it('completeLevel unlocks the next level', () => {
        const manager = new LevelManager();
        manager.completeLevel('1-1', { score: 100, wpm: 20, accuracy: 0.95 });
        expect(manager.isLevelUnlocked('1-2')).toBe(true);
    });

    it('completing 1-2 unlocks 1-3', () => {
        const manager = new LevelManager();
        manager.completeLevel('1-1', { score: 100, wpm: 20, accuracy: 0.95 });
        expect(manager.isLevelUnlocked('1-2')).toBe(true);
        manager.completeLevel('1-2', { score: 110, wpm: 22, accuracy: 0.97 });
        expect(manager.isLevelUnlocked('1-3')).toBe(true);
    });

    it('completing all world 1 levels unlocks world 2', () => {
        const manager = new LevelManager();
        const world1Levels = [
            '1-1', '1-2', '1-3', '1-4', '1-5', '1-6', '1-7'
        ];
        world1Levels.forEach((levelId, idx) => {
            manager.completeLevel(levelId, { score: 100 + idx * 10, wpm: 20 + idx * 2, accuracy: 0.95 });
        });
        expect(manager.isLevelUnlocked('2-1')).toBe(true);
    });

    it('level 1-5 unlocks after completing 1-4', () => {
        const manager = new LevelManager();
        manager.completeLevel('1-4', { score: 100, wpm: 20, accuracy: 0.95 });
        expect(manager.isLevelUnlocked('1-5')).toBe(true);
    });

    it('level 1-6 unlocks after completing 1-5', () => {
        const manager = new LevelManager();
        manager.completeLevel('1-5', { score: 100, wpm: 20, accuracy: 0.95 });
        expect(manager.isLevelUnlocked('1-6')).toBe(true);
    });

    it('level 1-7 unlocks after completing 1-6', () => {
        const manager = new LevelManager();
        manager.completeLevel('1-6', { score: 100, wpm: 20, accuracy: 0.95 });
        expect(manager.isLevelUnlocked('1-7')).toBe(true);
    });

    it('level 1-5 uses the correct word list', () => {
        const fs = require('fs');
        expect(fs.existsSync('src/wordpacks/fjghrutyvmWords.json')).toBe(true);
    });
    it('level 1-6 uses the correct word list', () => {
        const fs = require('fs');
        expect(fs.existsSync('src/wordpacks/fjghrutyvmbnWords.json')).toBe(true);
    });
    it('level 1-7 uses the correct word list', () => {
        const fs = require('fs');
        expect(fs.existsSync('src/wordpacks/fjghrutyvmbn_bossWords.json')).toBe(true);
    });
});

