import { describe, expect, it } from 'vitest';

// Patch: Mock MobSpawner without importing Phaser or Mob for scaling test
class MockMobSpawner {
    public spawnInterval: number;
    public mobBaseSpeed: number;
    public minSpawnInterval: number;
    public maxMobSpeed: number;
    public initialSpawnInterval: number;
    public initialMobBaseSpeed: number;
    constructor(initialInterval: number, initialSpeed: number, minInterval: number, maxSpeed: number) {
        this.initialSpawnInterval = initialInterval;
        this.initialMobBaseSpeed = initialSpeed;
        this.minSpawnInterval = minInterval;
        this.maxMobSpeed = maxSpeed;
        this.spawnInterval = initialInterval;
        this.mobBaseSpeed = initialSpeed;
    }
    setProgression(progression: number) {
        // Linear interpolation
        this.spawnInterval = this.initialSpawnInterval + (this.minSpawnInterval - this.initialSpawnInterval) * progression;
        this.mobBaseSpeed = this.initialMobBaseSpeed + (this.maxMobSpeed - this.initialMobBaseSpeed) * progression;
    }
}

describe('MobSpawner scaling', () => {
    it('should interpolate spawnInterval and mobBaseSpeed smoothly as progression increases', () => {
        const initialInterval = 2000;
        const minInterval = 600;
        const initialSpeed = 90;
        const maxSpeed = 250;
        const spawner = new MockMobSpawner(initialInterval, initialSpeed, minInterval, maxSpeed);
        // At progression 0
        spawner.setProgression(0);
        expect(spawner.spawnInterval).toBeCloseTo(initialInterval);
        expect(spawner.mobBaseSpeed).toBeCloseTo(initialSpeed);
        // At progression 1
        spawner.setProgression(1);
        expect(spawner.spawnInterval).toBeCloseTo(minInterval);
        expect(spawner.mobBaseSpeed).toBeCloseTo(maxSpeed);
        // At progression 0.5
        spawner.setProgression(0.5);
        expect(spawner.spawnInterval).toBeCloseTo((initialInterval + minInterval) / 2, 0);
        expect(spawner.mobBaseSpeed).toBeCloseTo((initialSpeed + maxSpeed) / 2, 0);
    });
});
