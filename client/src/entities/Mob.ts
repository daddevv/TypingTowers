// Mob.ts
// Helper class to update all MobStates in the game
import stateManager from '../state/stateManager';

export class MobSystem {
    /**
     * Update all mobs in the game state.
     * @param time The current time
     * @param delta The time since the last update
     */
    static updateAll(time: number, delta: number) {
        const state = stateManager.getState();
        const mobs = state.mobs;
        const playerPos = state.player.position;
        for (const mob of mobs) {
            if (mob.isDefeated) continue;
            // Move toward player
            const dx = playerPos.x - mob.position.x;
            const dy = playerPos.y - mob.position.y;
            const dist = Math.sqrt(dx * dx + dy * dy);
            if (dist > 1) {
                const speed = mob.speed;
                mob.position.x += (dx / dist) * speed * (delta / 1000);
                mob.position.y += (dy / dist) * speed * (delta / 1000);
            }
            // TODO: Add avoidance logic if needed
            // If mob is defeated, mark in state
            if (mob.currentTypedIndex >= mob.word.length) {
                mob.isDefeated = true;
            }
        }
        // Write back updated mobs
        // (In a real ECS, this would be more granular)
        stateManager.updateMobs(mobs);
    }
}

//Contains AI - generated edits.
