import { beforeEach, describe, expect, it, vi } from 'vitest';
import GameScene from '../GameScene';

// Mock Phaser dependencies
class MockScene {
    add = {
        particles: vi.fn(),
        rectangle: vi.fn(() => ({ destroy: vi.fn() })),
        text: vi.fn(() => ({ setOrigin: vi.fn(() => ({ setInteractive: vi.fn(() => ({ on: vi.fn() })) })), setInteractive: vi.fn(() => ({ on: vi.fn() })), destroy: vi.fn() })),
        container: vi.fn(() => ({ destroy: vi.fn() }))
    };
    input = { keyboard: { on: vi.fn(), off: vi.fn() } };
    scene = { pause: vi.fn(), resume: vi.fn(), stop: vi.fn(), start: vi.fn() };
}

describe('GameScene Pause Menu', () => {
    let scene: GameScene & MockScene;
    beforeEach(() => {
        scene = new GameScene() as any;
        Object.assign(scene, new MockScene());
        scene.isPaused = false;
        scene.levelCompleteText = undefined;
    });

    it('shows pause menu and pauses game on showPauseMenu()', () => {
        scene.showPauseMenu();
        expect(scene.isPaused).toBe(true);
        expect(scene.scene.pause).toHaveBeenCalled();
        expect(scene.pauseMenuContainer).toBeDefined();
    });

    it('hides pause menu and resumes game on hidePauseMenu()', () => {
        scene.showPauseMenu();
        scene.hidePauseMenu();
        expect(scene.isPaused).toBe(false);
        expect(scene.scene.resume).toHaveBeenCalled();
        expect(scene.pauseMenuContainer).toBeUndefined();
    });

    it('does not show pause menu if already paused', () => {
        scene.isPaused = true;
        scene.showPauseMenu();
        // Should not call pause again
        expect(scene.scene.pause).not.toHaveBeenCalled();
    });

    it('does not hide pause menu if not paused', () => {
        scene.isPaused = false;
        scene.hidePauseMenu();
        // Should not call resume
        expect(scene.scene.resume).not.toHaveBeenCalled();
    });
});
// Contains AI-generated edits.
