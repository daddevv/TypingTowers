import Phaser from 'phaser';
import { beforeEach, describe, expect, it, vi } from 'vitest';
import GameScene from '../../scenes/GameScene';

// Mock Phaser dependencies for headless test
globalThis.Phaser = Phaser;

describe('GameScene Combo Multiplier Integration', () => {
    let scene: GameScene;

    beforeEach(() => {
        scene = new GameScene();
        // Minimal mock for required Phaser methods
        (scene as any).add = { text: () => ({ setScrollFactor: () => { }, setDepth: () => { }, setStyle: () => { }, setText: () => { }, setVisible: () => { } }) };
        (scene as any).make = { particles: () => ({ createEmitter: () => ({ explode: () => { } }) }) };
        (scene as any).inputHandler = { on: () => { } };
        (scene as any).mobSpawner = { getMobs: () => [] };
        (scene as any).fingerGroupManager = { getCurrentFingerGroup: () => ({}) };
        (scene as any).levelManager = { getCurrentLevel: () => ({}) };
        (scene as any).score = 0;
        (scene as any).combo = 0;
        (scene as any).scoreText = { setText: () => { } };
        (scene as any).comboText = { setText: () => { }, setVisible: () => { } };
    });

    it('increments combo and score on correct keystroke', () => {
        (scene as any).combo = 0;
        (scene as any).score = 0;
        scene.handleCorrectKeystroke({ x: 100, y: 100 });
        expect((scene as any).combo).toBe(1);
        expect((scene as any).score).toBeGreaterThan(0);
    });

    it('resets combo on incorrect keystroke', () => {
        (scene as any).combo = 3;
        scene.handleIncorrectKeystroke();
        expect((scene as any).combo).toBe(0);
    });

    it('score calculation uses combo multiplier', () => {
        (scene as any).combo = 2;
        (scene as any).score = 0;
        scene.handleCorrectKeystroke({ x: 100, y: 100 });
        // Should add more than base points
        expect((scene as any).score).toBeGreaterThan(10);
    });

    it('triggers particle burst on correct keystroke', () => {
        (scene as any).combo = 0;
        (scene as any).score = 0;
        // Mock particleManager with spy
        const emitSpy = vi.fn();
        (scene as any).particleManager = { emitParticleAt: emitSpy };
        (scene as any).scoreText = { setText: () => { } };
        (scene as any).comboText = { setText: () => { }, setVisible: () => { } };
        // Simulate correct keystroke
        scene.handleCorrectKeystroke({ x: 123, y: 456 });
        expect(emitSpy).toHaveBeenCalledWith(123, 456, 12);
    });
});

// Contains AI-generated edits.
