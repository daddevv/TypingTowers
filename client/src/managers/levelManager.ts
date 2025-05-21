/**
 * Manages level progression and tracking player progress through worlds.
 */
import { getAvailableKeysForLevel, getLevelById, LevelConfig, WorldConfig, WORLDS } from '../curriculum/worldConfig';

export interface LevelProgress {
    completed: boolean;
    highScore: number;
    bestWPM: number;
    bestAccuracy: number;
    attempts: number;
}

export default class LevelManager {
    private currentWorldId: number = 1;
    private currentLevelId: string = '1-1';
    private progress: Map<string, LevelProgress> = new Map();

    constructor() {
        // Initialize with first level unlocked
        this.setLevelProgress('1-1', {
            completed: false,
            highScore: 0,
            bestWPM: 0,
            bestAccuracy: 0,
            attempts: 0
        });
    }

    /**
     * Gets the current world configuration
     */
    getCurrentWorld(): WorldConfig | undefined {
        return WORLDS.find(w => w.id === this.currentWorldId);
    }

    /**
     * Gets the current level configuration
     */
    getCurrentLevel(): LevelConfig | undefined {
        return getLevelById(this.currentWorldId, this.currentLevelId);
    }

    /**
     * Gets the available letters for the current level
     */
    getCurrentLevelLetters(): string[] {
        return getAvailableKeysForLevel(this.currentWorldId, this.currentLevelId);
    }

    /**
     * Sets the current level
     */
    setCurrentLevel(worldId: number, levelId: string): boolean {
        const level = getLevelById(worldId, levelId);
        if (!level) return false;

        this.currentWorldId = worldId;
        this.currentLevelId = levelId;
        return true;
    }

    /**
     * Gets progress for a specific level
     */
    getLevelProgress(levelId: string): LevelProgress | undefined {
        return this.progress.get(levelId);
    }

    /**
     * Sets progress for a specific level
     */
    setLevelProgress(levelId: string, progress: LevelProgress): void {
        this.progress.set(levelId, progress);
    }

    /**
     * Updates progress after completing a level
     */
    updateLevelProgress(
        levelId: string,
        completed: boolean,
        score: number,
        wpm: number,
        accuracy: number
    ): void {
        const currentProgress = this.getLevelProgress(levelId) || {
            completed: false,
            highScore: 0,
            bestWPM: 0,
            bestAccuracy: 0,
            attempts: 0
        };

        const newProgress: LevelProgress = {
            completed: completed || currentProgress.completed,
            highScore: Math.max(score, currentProgress.highScore),
            bestWPM: Math.max(wpm, currentProgress.bestWPM),
            bestAccuracy: Math.max(accuracy, currentProgress.bestAccuracy),
            attempts: currentProgress.attempts + 1
        };

        this.setLevelProgress(levelId, newProgress);

        // If completed, unlock the next level
        if (completed && !currentProgress.completed) {
            this.unlockNextLevel();
        }
    }

    /**
     * Unlocks the next level based on current progress
     */
    private unlockNextLevel(): void {
        const currentWorld = this.getCurrentWorld();
        if (!currentWorld) return;

        const currentLevelIndex = currentWorld.levels.findIndex(l => l.id === this.currentLevelId);
        if (currentLevelIndex === -1) return;

        // Check if there's another level in this world
        if (currentLevelIndex < currentWorld.levels.length - 1) {
            const nextLevel = currentWorld.levels[currentLevelIndex + 1];
            this.setLevelProgress(nextLevel.id, {
                completed: false,
                highScore: 0,
                bestWPM: 0,
                bestAccuracy: 0,
                attempts: 0
            });
        }
        // If this was the last level in the world, unlock the first level of the next world
        else if (currentWorld.id < WORLDS.length) {
            const nextWorld = WORLDS.find(w => w.id === currentWorld.id + 1);
            if (nextWorld && nextWorld.levels.length > 0) {
                const nextLevelId = nextWorld.levels[0].id;
                this.setLevelProgress(nextLevelId, {
                    completed: false,
                    highScore: 0,
                    bestWPM: 0,
                    bestAccuracy: 0,
                    attempts: 0
                });
            }
        }
    }

    /**
     * Checks if a specific level is unlocked
     */
    isLevelUnlocked(worldId: number, levelId: string): boolean {
        return this.progress.has(levelId);
    }

    /**
     * Gets all unlocked levels
     */
    getUnlockedLevels(): string[] {
        return Array.from(this.progress.keys());
    }

    /**
     * Saves progress to localStorage
     */
    saveProgress(): void {
        const progressData = Array.from(this.progress.entries());
        localStorage.setItem('typeDefense_levelProgress', JSON.stringify(progressData));
        localStorage.setItem('typeDefense_currentWorld', this.currentWorldId.toString());
        localStorage.setItem('typeDefense_currentLevel', this.currentLevelId);
    }

    /**
     * Loads progress from localStorage
     */
    loadProgress(): boolean {
        const progressData = localStorage.getItem('typeDefense_levelProgress');
        const worldId = localStorage.getItem('typeDefense_currentWorld');
        const levelId = localStorage.getItem('typeDefense_currentLevel');

        if (progressData) {
            try {
                const parsedData = JSON.parse(progressData) as [string, LevelProgress][];
                this.progress = new Map(parsedData);

                if (worldId && levelId) {
                    this.currentWorldId = parseInt(worldId, 10);
                    this.currentLevelId = levelId;
                }

                return true;
            } catch (e) {
                console.error('Failed to load progress data', e);
            }
        }

        return false;
    }
}

//Contains AI - generated edits.
