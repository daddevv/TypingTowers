/**
 * Manages tracking of player progress across finger groups.
 * Responsibilities:
 * - Record key presses and associate them with finger groups
 * - Track stats: total presses, correct/incorrect usage, accuracy, speed per finger group
 * - Provide progress data for curriculum advancement
 * - Expose methods to reset, update, and retrieve stats
 * - (Future) Integrate with UI for feedback and achievements
 */
import { FINGER_GROUP_KEYS, FingerType, getKeyInfo } from '../curriculum/fingerGroups';

interface FingerStats {
    totalKeyPresses: number;
    correctFingerUses: number;
    mistypedKeys: number;
    accuracy: number;
    averageSpeed: number;
}

export default class FingerGroupManager {
    private fingerStats: Map<FingerType, FingerStats> = new Map();
    private keyPressHistory: Map<string, number[]> = new Map();
    private lastKeyTimestamps: Map<FingerType, number> = new Map();
    constructor() {
        Object.values(FingerType).forEach(finger => {
            this.fingerStats.set(finger, {
                totalKeyPresses: 0,
                correctFingerUses: 0,
                mistypedKeys: 0,
                accuracy: 1,
                averageSpeed: 0,
            });
        });
    }

    /**
     * Records a key press with timing information
     * @param key The key that was pressed
     * @param usedCorrectFinger Whether the correct finger was used (future feature with camera)
     * @param timestamp Time when the key was pressed
     */
    recordKeyPress(key: string, usedCorrectFinger: boolean = true, timestamp: number = Date.now()): void {
        const keyInfo = getKeyInfo(key);
        if (!keyInfo) return;
        const finger = keyInfo.finger;
        const stats = this.fingerStats.get(finger)!;
        stats.totalKeyPresses++;
        if (usedCorrectFinger) {
            stats.correctFingerUses++;
        } else {
            stats.mistypedKeys++;
        }
        stats.accuracy = stats.correctFingerUses / stats.totalKeyPresses;
        // Speed calculation
        const last = this.lastKeyTimestamps.get(finger);
        if (last) {
            const interval = timestamp - last;
            if (!this.keyPressHistory.has(key)) this.keyPressHistory.set(key, []);
            this.keyPressHistory.get(key)!.push(interval);
            this.updateAverageSpeed(finger);
        }
        this.lastKeyTimestamps.set(finger, timestamp);
    }

    /**
     * Updates the average typing speed for a finger
     */
    private updateAverageSpeed(finger: FingerType): void {
        let total = 0;
        let count = 0;
        for (const [key, intervals] of this.keyPressHistory.entries()) {
            const keyInfo = getKeyInfo(key);
            if (keyInfo && keyInfo.finger === finger) {
                total += intervals.reduce((a, b) => a + b, 0);
                count += intervals.length;
            }
        }
        const stats = this.fingerStats.get(finger)!;
        stats.averageSpeed = count > 0 ? total / count : 0;
    }

    /**
     * Gets all keys associated with a specific finger
     */
    getKeysForFinger(finger: FingerType): string[] {
        const keys: string[] = [];
        for (const group of Object.values(FINGER_GROUP_KEYS)) {
            for (const mapping of group) {
                if (mapping.finger === finger) {
                    keys.push(mapping.key);
                }
            }
        }
        return keys;
    }

    /**
     * Gets stats for a specific finger
     */
    getFingerStats(finger: FingerType): FingerStats | undefined {
        return this.fingerStats.get(finger);
    }

    /**
     * Gets the finger that should be used for a specific key
     */
    getFingerForKey(key: string): FingerType | undefined {
        const keyInfo = getKeyInfo(key);
        return keyInfo ? keyInfo.finger : undefined;
    }

    /**
     * Checks if a key is mastered (high accuracy and good speed)
     */
    isKeyMastered(key: string): boolean {
        const keyInfo = getKeyInfo(key);
        if (!keyInfo) return false;
        const stats = this.fingerStats.get(keyInfo.finger);
        if (!stats) return false;
        return stats.accuracy > 0.95 && stats.averageSpeed < 350; // Example mastery criteria
    }

    /**
     * Gets overall typing proficiency as a percentage
     */
    getOverallProficiency(): number {
        let total = 0;
        let mastered = 0;
        for (const finger of Object.values(FingerType)) {
            const stats = this.fingerStats.get(finger);
            if (stats) {
                total++;
                if (stats.accuracy > 0.95 && stats.averageSpeed < 350) mastered++;
            }
        }
        return total > 0 ? (mastered / total) * 100 : 0;
    }

    /**
     * Saves stats to localStorage
     */
    saveStats(): void {
        const obj: Record<string, FingerStats> = {};
        for (const [finger, stats] of this.fingerStats.entries()) {
            obj[finger] = stats;
        }
        localStorage.setItem('fingerStats', JSON.stringify(obj));
    }
}

// Contains AI-generated edits.
