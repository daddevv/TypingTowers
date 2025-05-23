// Global GameState definition for TypeDefense v2
// This interface should be imported wherever state is read or updated.

import FingerGroupManager from '../managers/fingerGroupManager';

// Define types for FingerGroupStats and FingerGroupProgress based on FingerGroupManager's methods
export type FingerGroupStats = ReturnType<FingerGroupManager['getAllFingerGroupStats']>[keyof ReturnType<FingerGroupManager['getAllFingerGroupStats']>];
export type FingerGroupProgress = ReturnType<FingerGroupManager['getAllFingerGroupProgress']>[keyof ReturnType<FingerGroupManager['getAllFingerGroupProgress']>];

export type GameStatus =
    | 'booting'
    | 'mainMenu'
    | 'worldSelect'
    | 'levelSelect'
    | 'playing'
    | 'paused'
    | 'levelComplete'
    | 'gameOver';

export interface PlayerState {
    health: number;
    maxHealth: number;
    score: number;
    combo: number;
    currentInput: string;
    position: { x: number; y: number };
    // Add more as needed (e.g., powerups)
}

export interface LevelState {
    currentWorld?: number | null;
    currentLevelId?: string | null;
    levelStatus: 'notStarted' | 'playing' | 'complete' | 'failed';
    // Add more as needed (e.g., wave, win condition)
}

export interface MobState {
    id: string;
    word: string;
    currentTypedIndex: number;
    position: { x: number; y: number };
    speed: number;
    type: string;
    isDefeated: boolean;
}

export interface MobSpawnerState {
    nextSpawnTime: number;
    currentWave: number;
    mobsRemainingInWave: number;
    // Add more as needed
}

export interface UIState {
    activeModals: string[];
    notifications: string[];
    showPauseMenu: boolean;
    showLevelComplete: boolean;
    // Add more as needed
}

export interface SettingsState {
    volume: number;
    difficulty: 'normal' | 'hard';
    // Add more as needed
}

export interface LevelProgress {
    completed: boolean;
    highScore: number;
    bestWPM: number;
    bestAccuracy: number;
    attempts: number;
}

export interface ProgressionState {
    unlockedWorlds: string[];
    unlockedLevels: string[];
    completedLevels: string[];
    fingerGroupStats: Record<string, FingerGroupStats>; // Deprecated: will move to curriculum
    fingerGroupProgress: Record<string, FingerGroupProgress>; // Deprecated: will move to curriculum
    levelProgress: Record<string, LevelProgress>; // New: stores progress for each level
}

export interface CurriculumState {
    worldConfig: any; // Replace with actual type if available
    fingerGroupStats: Record<string, FingerGroupStats>;
    fingerGroupProgress: Record<string, FingerGroupProgress>;
}

export interface GameState {
    player: PlayerState;
    level: LevelState;
    gameStatus: GameStatus;
    mobs: MobState[];
    mobSpawner: MobSpawnerState;
    ui: UIState;
    settings: SettingsState;
    progression: ProgressionState;
    curriculum: CurriculumState;
    timestamp: number;
    delta: number;
}

// Optionally, you can export a default empty state for initialization
export const defaultGameState: GameState = {
    player: {
        health: 3,
        maxHealth: 3,
        score: 0,
        combo: 0,
        currentInput: '',
        position: { x: 100, y: 300 },
    },
    level: {
        currentWorld: 1,
        currentLevelId: '1-1',
        levelStatus: 'notStarted',
    },
    gameStatus: 'booting',
    mobs: [],
    mobSpawner: {
        nextSpawnTime: 0,
        currentWave: 1,
        mobsRemainingInWave: 0,
    },
    ui: {
        activeModals: [],
        notifications: [],
        showPauseMenu: false,
        showLevelComplete: false,
    },
    settings: {
        volume: 1,
        difficulty: 'normal',
    },
    progression: {
        unlockedWorlds: [],
        unlockedLevels: [],
        completedLevels: [],
        fingerGroupStats: {},
        fingerGroupProgress: {},
        levelProgress: {},
    },
    curriculum: {
        worldConfig: {},
        fingerGroupStats: {},
        fingerGroupProgress: {},
    },
    timestamp: 0,
    delta: 0,
};
