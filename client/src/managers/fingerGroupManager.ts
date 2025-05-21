/**
 * Manages tracking of player progress across finger groups.
 * Provides feedback on which finger should be used for each key.
 */
import {
    FINGER_GROUP_KEYS,
    FingerType,
    getKeyInfo
} from '../curriculum/fingerGroups';

interface FingerStats {
    totalKeyPresses: number;
    correctFingerUses: number;
    mistypedKeys: number;
    accuracy: number;
    averageSpeed: number;
}

export default class FingerGroupManager {
    private fingerStats: Map<FingerType, FingerStats> = new Map();
    private keyPressHistory: Map<string, number[]> = new Map(); // Key -> array of press times

    constructor() {
        // Initialize stats for all finger types
        Object.values(FingerType).forEach(finger => {
            this.fingerStats.set(finger, {
                totalKeyPresses: 0,
                correctFingerUses: 0,
                mistypedKeys: 0,
                accuracy: 0,
                averageSpeed: 0
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

        // Record timestamp for this key
        if (!this.keyPressHistory.has(key)) {
            this.keyPressHistory.set(key, []);
        }
        this.keyPressHistory.get(key)?.push(timestamp);

        // Keep only the last 20 presses for performance
        const history = this.keyPressHistory.get(key);
        if (history && history.length > 20) {
            this.keyPressHistory.set(key, history.slice(history.length - 20));
        }

        // Update finger stats
        const stats = this.fingerStats.get(keyInfo.finger);
        if (stats) {
            stats.totalKeyPresses++;
            if (usedCorrectFinger) {
                stats.correctFingerUses++;
            } else {
                stats.mistypedKeys++;
            }
            stats.accuracy = stats.correctFingerUses / stats.totalKeyPresses;

            // Update average speed if we have enough history
            this.updateAverageSpeed(keyInfo.finger);
        }
    }

    /**
     * Updates the average typing speed for a finger
     */
    private updateAverageSpeed(finger: FingerType): void {
        const fingerKeys = this.getKeysForFinger(finger);
        let totalIntervals = 0;
        let intervalCount = 0;

        fingerKeys.forEach(key => {
            const history = this.keyPressHistory.get(key);
            if (history && history.length > 1) {
                for (let i = 1; i < history.length; i++) {
                    const interval = history[i] - history[i - 1];
                    // Only count reasonable intervals (< 2 seconds)
                    if (interval > 0 && interval < 2000) {
                        totalIntervals += interval;
                        intervalCount++;
                    }
                }
            }
        });

        const stats = this.fingerStats.get(finger);
        if (stats && intervalCount > 0) {
            // Average time between keypresses in milliseconds
            const avgInterval = totalIntervals / intervalCount;
            // Convert to keys per minute
            stats.averageSpeed = 60000 / avgInterval;
        }
    }

    /**
     * Gets all keys associated with a specific finger
     */
    getKeysForFinger(finger: FingerType): string[] {
        const keys: string[] = [];

        Object.values(FINGER_GROUP_KEYS).forEach(group => {
            group.forEach(keyMapping => {
                if (keyMapping.finger === finger) {
                    keys.push(keyMapping.key);
                }
            });
        });

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
        return keyInfo?.finger;
    }

    /**
     * Checks if a key is mastered (high accuracy and good speed)
     */
    isKeyMastered(key: string): boolean {
        const keyInfo = getKeyInfo(key);
        if (!keyInfo) return false;

        const stats = this.fingerStats.get(keyInfo.finger);
        if (!stats) return false;

        // Consider a key mastered if accuracy > 95% and we have enough sample data
        return stats.accuracy > 0.95 && stats.totalKeyPresses > 20;
    }

    /**
     * Gets overall typing proficiency as a percentage
     */
    getOverallProficiency(): number {
        let totalAccuracy = 0;
        let fingerCount = 0;

        this.fingerStats.forEach(stats => {
            if (stats.totalKeyPresses > 0) {
                totalAccuracy += stats.accuracy;
                fingerCount++;
            }
        });

        return fingerCount > 0 ? (totalAccuracy / fingerCount) * 100 : 0;
    }

    /**
     * Saves stats to localStorage
     */
    saveStats(): void {
        const statsData = Array.from(this.fingerStats.entries());
        localStorage.setItem('typeDefense_fingerStats', JSON.stringify(statsData));
    }

    /**
     * Loads stats from localStorage
     */
    loadStats(): boolean {
        const statsData = localStorage.getItem('typeDefense_fingerStats');

        if (statsData) {
            try {
                const parsedData = JSON.parse(statsData) as [FingerType, FingerStats][];
                this.fingerStats = new Map(parsedData);
                return true;
            } catch (e) {
                console.error('Failed to load finger stats data', e);
            }
        }

        return false;
    }
}

//Contains AI - generated edits.
