import { describe, expect, it, vi } from 'vitest';
import Mob from '../Mob';

// Deep mock for Phaser.GameObjects.Sprite and its parent classes for MobSpawner tests
class MockGameObject {
    anims: { on: Function };
    constructor(public scene: any, public x: number, public y: number, public texture: string) {
        this.anims = { on: () => { } };
    }
    setOrigin() { return this; }
}
(global as any).Phaser = {
    GameObjects: {
        Sprite: MockGameObject,
        Text: class { constructor() { return { setOrigin: () => this }; } },
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
        add: { existing: vi.fn(), particles: vi.fn(), text: vi.fn(() => ({ setOrigin: vi.fn().mockReturnThis() })) },
        time: { delayedCall: vi.fn() },
        tweens: { add: vi.fn() },
        sys: {
            queueDepthSort: () => { }, // Mock as function for Phaser internals
            displayList: { add: vi.fn() },
            updateList: { add: vi.fn() },
        },
    } as any;
}

describe('Mob', () => {
    it('should create an instance of Mob', () => {
        const scene = createMockScene();
        const mob = new Mob(scene, 100, 100, 'testTexture');
        expect(mob).toBeInstanceOf(Mob);
    });

    it('should have the correct initial properties', () => {
        const scene = createMockScene();
        const mob = new Mob(scene, 100, 100, 'testTexture');
        expect(mob.x).toBe(100);
        expect(mob.y).toBe(100);
        expect(mob.texture.key).toBe('testTexture');
    });

    it('should play the correct animation on spawn', () => {
        const scene = createMockScene();
        const mob = new Mob(scene, 100, 100, 'testTexture');
        mob.anims = { play: vi.fn() }; // Mock the anims property
        mob.spawn(); // Assuming spawn is the method that triggers the animation
        expect(mob.anims.play).toHaveBeenCalledWith('mobSpawnAnimation'); // Replace with the actual animation key
    });
});
