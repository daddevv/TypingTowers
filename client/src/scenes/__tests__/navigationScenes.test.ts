import Phaser from 'phaser';
import { beforeEach, describe, expect, it, vi } from 'vitest';
import LevelMenuScene from '../LevelMenuScene';
import MenuScene from '../MenuScene';

// Mock Phaser game objects and methods for headless testing
class MockText {
    constructor(...args: any[]) { }
    setOrigin() { return this; }
    setInteractive() { return this; }
    on() { return this; }
    destroy() { }
}

// Mock Phaser.Scene base
class MockScene extends Phaser.Scene {
    add = { text: vi.fn(() => new MockText()) };
    input = { keyboard: { on: vi.fn(), off: vi.fn() } };
    scene = { start: vi.fn(), restart: vi.fn(), pause: vi.fn() };
    scale = { width: 800, height: 600 };
}

describe('MenuScene navigation', () => {
    let scene: MenuScene;
    beforeEach(() => {
        scene = new MenuScene();
        Object.assign(scene, new MockScene());
        // Setup minimal worlds and levelManager mocks
        (scene as any).worlds = [
            { id: 1, name: 'World 1', levels: [{ id: '1-1', name: 'Level 1-1' }] },
            { id: 2, name: 'World 2', levels: [{ id: '2-1', name: 'Level 2-1' }] }
        ];
        (scene as any).levelManager = { isLevelUnlocked: () => true };
        (scene as any).menuItems = [];
        (scene as any).selectedWorld = 0;
    });

    it('renders menu and responds to keyboard navigation', () => {
        scene.renderMenu();
        scene.handleInput({ key: 'ArrowDown' } as any);
        expect((scene as any).selectedWorld).toBe(1);
        scene.handleInput({ key: 'ArrowUp' } as any);
        expect((scene as any).selectedWorld).toBe(0);
    });

    it('selects world on Enter', () => {
        const startSpy = vi.spyOn((scene as any).scene, 'start');
        scene.handleInput({ key: 'Enter' } as any);
        expect(startSpy).toHaveBeenCalledWith('LevelMenuScene', { worldId: 1 });
    });
});

describe('LevelMenuScene navigation', () => {
    let scene: LevelMenuScene;
    beforeEach(() => {
        scene = new LevelMenuScene();
        Object.assign(scene, new MockScene());
        (scene as any).world = {
            id: 1,
            name: 'World 1',
            levels: [
                { id: '1-1', name: 'Level 1-1' },
                { id: '1-2', name: 'Level 1-2' }
            ]
        };
        (scene as any).selectedLevel = 0;
        (scene as any).menuItems = [];
        (scene as any).levelManager = { isLevelUnlocked: () => true, setCurrentLevel: vi.fn() };
    });

    it('renders level menu and responds to keyboard navigation', () => {
        scene.renderMenu();
        scene.handleInput({ key: 'ArrowDown' } as any);
        expect((scene as any).selectedLevel).toBe(1);
        scene.handleInput({ key: 'ArrowUp' } as any);
        expect((scene as any).selectedLevel).toBe(0);
    });

    it('selects level on Enter', () => {
        const startSpy = vi.spyOn((scene as any).scene, 'start');
        scene.handleInput({ key: 'Enter' } as any);
        expect(startSpy).toHaveBeenCalledWith('GameScene', { worldId: 1, levelId: '1-1' });
    });

    it('returns to MenuScene on Escape', () => {
        const startSpy = vi.spyOn((scene as any).scene, 'start');
        scene.handleInput({ key: 'Escape' } as any);
        expect(startSpy).toHaveBeenCalledWith('MenuScene');
    });
});

// Contains AI-generated edits.
