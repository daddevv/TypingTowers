// GameJourneyStats.integration.test.ts
// Integration test for simulating a full game journey with detailed stats reporting

import { afterAll, afterEach, beforeAll, beforeEach, describe, expect, it, vi } from 'vitest';
import { WorldConfig, WORLDS } from '../../curriculum/worldConfig'; // Updated imports
import { NullRenderAdapter } from '../../render/NullRenderAdapter';
import stateManager from '../../state/stateManager';
import HeadlessGameEngine, { HeadlessGameEngineOptions } from '../HeadlessGameEngine'; // Updated import

// Track player performance metrics
interface PlayerMetrics {
    totalKeystrokes: number;
    accuracy: number;
    wordsPerMinute: number;
    avgCombo: number;
    highestCombo: number;
    timeSpent: number;
}

// Track level statistics
interface LevelStats {
    levelId: string;
    startTime: number;
    endTime: number;
    timeToComplete: number;
    score: number;
    mobsDefeated: number;
    keystrokes: number;
    correctKeystrokes: number;
    accuracy: number;
    avgWordLength: number;
    comboHighest: number;
    comboAverage: number;
    wordsPerMinute: number;
    difficultyRating: string;
}

// Track world statistics
interface WorldStats {
    worldId: number;
    worldName: string;
    levels: LevelStats[];
    startTime: number;
    endTime: number;
    totalTimeToComplete: number;
    totalScore: number;
    totalMobsDefeated: number;
    averageLevelCompletion: number;
    keystrokes: number;
    accuracy: number;
    wordsPerMinute: number;
    difficultyProgression: string[];
    playerPerformance: PlayerMetrics;
}

// Track journey statistics
interface JourneyStats {
    initialState: Record<string, any>;
    worldProgress: WorldStats[];
    finalStats: Record<string, any>;
    totalTimeToComplete: number;
    totalLevelsCompleted: number;
    totalWordsTyped: number;
    overallAccuracy: number;
    overallWPM: number;
    skillProgression: Record<string, number>;
    difficultyByWorld: Record<string, string>;
}

// Helper function to fast forward time in the engine
function fastForward(engine: HeadlessGameEngine, ms: number, step: number = 16) {
    let elapsed = 0;
    while (elapsed < ms) {
        engine.step(step);
        elapsed += step;
    }
}

// Helper function to complete a level with detailed metrics
async function completeLevel(
    engine: HeadlessGameEngine,
    levelStats: LevelStats,
    mobsToDefeatTarget: number // Use mobsToDefeat from LevelConfig
): Promise<void> {
    // Fast-forward until mobs spawn
    fastForward(engine, 2000);

    const startTime = Date.now();
    let defeatedCount = 0;
    let attempts = 0;
    const maxAttempts = mobsToDefeatTarget * 5 + 20; // Prevent infinite loops, allow more attempts for more mobs
    let totalLettersTyped = 0;
    let totalWords = 0;
    let totalWordLength = 0;
    let comboSum = 0;
    let comboCount = 0;
    let highestCombo = 0;
    let correctKeystrokes = 0;
    let incorrectKeystrokes = 0;

    while (defeatedCount < mobsToDefeatTarget && attempts < maxAttempts) {
        const state = engine.getState();

        // If there are mobs, defeat one
        if (state.mobs.length > 0) {
            const mob = state.mobs[0];
            totalWords++;
            totalWordLength += mob.word.length;

            // Simulate typing each letter in the mob's word
            for (const char of mob.word) {
                // 90% chance of correct keystroke, 10% chance of error then correction
                if (Math.random() < 0.9) {
                    engine.injectInput(char);
                    correctKeystrokes++;
                } else {
                    // Simulate typo and correction
                    engine.injectInput(Math.random().toString(36).substring(2, 3)); // Random character
                    incorrectKeystrokes++;
                    fastForward(engine, 50); // Small delay for "mistake realization"
                    engine.injectInput(char); // Correct character
                    correctKeystrokes++;
                }

                totalLettersTyped++;
                fastForward(engine, 100); // Give time for input processing

                // Track combo
                comboSum += state.player.combo;
                comboCount++;
                highestCombo = Math.max(highestCombo, state.player.combo);
            }

            defeatedCount++;
        }

        // Fast-forward to spawn more mobs if needed
        fastForward(engine, 500);
        attempts++;
    }

    const endTime = Date.now();
    const timeToComplete = endTime - startTime;

    // Calculate metrics
    levelStats.timeToComplete = timeToComplete;
    levelStats.endTime = endTime;
    levelStats.mobsDefeated = defeatedCount;
    levelStats.keystrokes = correctKeystrokes + incorrectKeystrokes;
    levelStats.correctKeystrokes = correctKeystrokes;
    levelStats.accuracy = correctKeystrokes / (correctKeystrokes + incorrectKeystrokes) * 100;
    levelStats.avgWordLength = totalWordLength / totalWords;
    levelStats.comboHighest = highestCombo;
    levelStats.comboAverage = comboSum / comboCount;
    levelStats.wordsPerMinute = (totalWords / (timeToComplete / 1000 / 60));

    // Determine difficulty rating based on metrics
    if (levelStats.wordsPerMinute > 30 && levelStats.accuracy > 95) {
        levelStats.difficultyRating = "Easy";
    } else if (levelStats.wordsPerMinute > 20 && levelStats.accuracy > 90) {
        levelStats.difficultyRating = "Moderate";
    } else {
        levelStats.difficultyRating = "Challenging";
    }

    // Update score from final state
    const finalState = engine.getState();
    levelStats.score = finalState.player.score;

    // Make sure level is completed
    const finalStateCheck = engine.getState();
    // Give the engine a few more steps to process win condition if not yet complete
    let retries = 0;
    while (finalStateCheck.level.levelStatus !== 'complete' && retries < 5) {
        engine.step(16);
        Object.assign(finalStateCheck, engine.getState());
        retries++;
    }
    expect(finalStateCheck.level.levelStatus).toBe('complete');
}

// Helper to get statistics
function getPlayerStats(engine: HeadlessGameEngine): Record<string, any> {
    const state = engine.getState();
    return {
        score: state.player.score,
        health: state.player.health,
        combo: state.player.combo,
        completedLevels: [...state.progression.completedLevels],
        unlockedLevels: [...state.progression.unlockedLevels],
        unlockedWorlds: [...state.progression.unlockedWorlds],
    };
}

describe('Game Journey - Complete Playthrough with Statistics', () => {
    let engine: HeadlessGameEngine;
    const journeyStats: JourneyStats = {
        initialState: {},
        worldProgress: [],
        finalStats: {},
        totalTimeToComplete: 0,
        totalLevelsCompleted: 0,
        totalWordsTyped: 0,
        overallAccuracy: 0,
        overallWPM: 0,
        skillProgression: {},
        difficultyByWorld: {}
    };

    beforeAll(() => {
        if (typeof (globalThis as any).setDeterministicRandomSequence === 'function') {
            (globalThis as any).setDeterministicRandomSequence([0.1, 0.9, 0.5, 0.7, 0.3]);
        }
        // Explicitly restore Math functions in case they were stubbed globally
        (globalThis as any).Math.min = Math.min;
        (globalThis as any).Math.max = Math.max;
        (globalThis as any).Math.abs = Math.abs;
        (globalThis as any).Math.floor = Math.floor;
    });
    afterAll(() => {
        if (typeof vi !== 'undefined' && vi.restoreAllMocks) {
            vi.restoreAllMocks();
        }
        // Explicitly restore Math functions again
        (globalThis as any).Math.min = Math.min;
        (globalThis as any).Math.max = Math.max;
        (globalThis as any).Math.abs = Math.abs;
        (globalThis as any).Math.floor = Math.floor;
    });

    beforeEach(() => {
        // Set up a clean engine instance for each test
        // Initial engine options are minimal, as reset will configure per-level
        engine = new HeadlessGameEngine({
            // winThreshold will be set per level by LevelConfig.mobsToDefeat
            // availableKeys will be set per level
            // wordList can be omitted to use WordGenerator with availableKeys
        }, new NullRenderAdapter());
        // engine.reset() will be called before each level with specific options.
        // Initial stateManager.reset() is good here.
        stateManager.reset();


        // Save initial state
        journeyStats.initialState = getPlayerStats(engine);
        // Clear world progress for fresh test run
        journeyStats.worldProgress = [];
        journeyStats.skillProgression = {};
        journeyStats.difficultyByWorld = {};

    });

    afterEach(() => {
        // Reset state manager after each test
        stateManager.reset();
    });

    it('should progress through all worlds and levels with comprehensive statistics', async () => {
        const journeyStartTime = Date.now();

        // Step 1: Start at main menu
        stateManager.setGameStatus('mainMenu');
        expect(engine.getState().gameStatus).toBe('mainMenu');

        // Step 2: Navigate to world selection
        stateManager.setGameStatus('worldSelect');
        expect(engine.getState().gameStatus).toBe('worldSelect');

        // Get all worlds from config
        const configuredWorlds: WorldConfig[] = WORLDS;

        // Cumulative stats for entire journey
        let totalKeystrokes = 0;
        let totalCorrectKeystrokes = 0;
        let totalWords = 0;
        let totalTimePlayed = 0;

        // For each world
        for (const world of configuredWorlds) {
            const worldStats: WorldStats = {
                worldId: world.id,
                worldName: world.name,
                levels: [],
                startTime: Date.now(),
                endTime: 0,
                totalTimeToComplete: 0,
                totalScore: 0,
                totalMobsDefeated: 0,
                averageLevelCompletion: 0,
                keystrokes: 0,
                accuracy: 0,
                wordsPerMinute: 0,
                difficultyProgression: [],
                playerPerformance: {
                    totalKeystrokes: 0,
                    accuracy: 0,
                    wordsPerMinute: 0,
                    avgCombo: 0,
                    highestCombo: 0,
                    timeSpent: 0
                }
            };

            console.log(`Starting World ${world.id}: ${world.name}`);
            // Select the world in stateManager
            stateManager.updateCurrentLevelContext({ currentWorld: world.id });

            // Add world to unlocked worlds if not already there
            // (Simulating game's progression logic)
            const currentProgression = stateManager.getState().progression;
            if (!currentProgression.unlockedWorlds.includes(String(world.id))) { // Ensure world.id is string if stored as string
                stateManager.updateProgression({
                    unlockedWorlds: [...currentProgression.unlockedWorlds, String(world.id)]
                });
            }

            // Navigate to level selection
            stateManager.setGameStatus('levelSelect');
            expect(engine.getState().gameStatus).toBe('levelSelect');

            // Get levels for this world from config
            const levelsInWorld = world.levels;

            // World level stats
            let worldKeystrokes = 0;
            let worldCorrectKeystrokes = 0;
            let worldTotalTime = 0;
            let worldMobsDefeated = 0;
            let worldTotalWords = 0;
            let worldComboSum = 0;
            let worldComboHighest = 0;

            // For each level in the world
            for (const levelConfig of levelsInWorld) { // Iterate through LevelConfig objects
                const levelId = levelConfig.id;
                const levelStats: LevelStats = {
                    levelId,
                    startTime: Date.now(),
                    endTime: 0,
                    timeToComplete: 0,
                    score: 0,
                    mobsDefeated: 0,
                    keystrokes: 0,
                    correctKeystrokes: 0,
                    accuracy: 0,
                    avgWordLength: 0,
                    comboHighest: 0,
                    comboAverage: 0,
                    wordsPerMinute: 0,
                    difficultyRating: ''
                };

                console.log(`  Starting Level ${levelId} (${levelConfig.name})`);
                // Select the level in stateManager
                stateManager.updateCurrentLevelContext({ currentLevelId: levelId, currentWorld: world.id });

                // Add level to unlocked levels if not already there
                // (Simulating game's progression logic)
                const currentLevelProgression = stateManager.getState().progression;
                if (!currentLevelProgression.unlockedLevels.includes(levelId)) {
                    stateManager.updateProgression({
                        unlockedLevels: [...currentLevelProgression.unlockedLevels, levelId]
                    });
                }

                // Prepare engine options for this specific level
                const engineOptions: HeadlessGameEngineOptions = {
                    availableKeys: levelConfig.availableKeys,
                    winThreshold: levelConfig.mobsToDefeat,
                    mobBaseSpeed: levelConfig.enemySpeed,
                    spawnInterval: levelConfig.enemySpawnRate,
                    wordList: undefined, // Ensure WordGenerator uses availableKeys
                };

                // Reset the engine with new level-specific configuration
                // This also sets gameStatus to 'playing' and levelStatus to 'playing'
                engine.reset(engineOptions);

                // Verify game status after reset
                expect(engine.getState().gameStatus).toBe('playing');
                expect(engine.getState().level.levelStatus).toBe('playing');
                expect(engine.getState().level.currentLevelId).toBe(levelId);


                // Complete the level with detailed metrics
                await completeLevel(engine, levelStats, levelConfig.mobsToDefeat);

                // Check if level is completed (gameStatus should be levelComplete)
                const finalLevelState = engine.getState();
                expect(finalLevelState.level.levelStatus).toBe('complete');
                expect(finalLevelState.gameStatus).toBe('levelComplete');

                // Add to completed levels in stateManager
                const finalProgression = stateManager.getState().progression;
                stateManager.updateProgression({
                    completedLevels: [...finalProgression.completedLevels, levelId]
                });

                // Add level stats to world stats
                worldStats.levels.push(levelStats);
                worldStats.totalScore += levelStats.score;
                worldStats.difficultyProgression.push(levelStats.difficultyRating);

                // Update world aggregate stats
                worldKeystrokes += levelStats.keystrokes;
                worldCorrectKeystrokes += levelStats.correctKeystrokes;
                worldTotalTime += levelStats.timeToComplete;
                worldMobsDefeated += levelStats.mobsDefeated;
                worldTotalWords += levelStats.mobsDefeated; // Assuming one word per mob
                worldComboSum += levelStats.comboAverage;
                worldComboHighest = Math.max(worldComboHighest, levelStats.comboHighest);

                // Update total journey stats
                totalKeystrokes += levelStats.keystrokes;
                totalCorrectKeystrokes += levelStats.correctKeystrokes;
                totalWords += levelStats.mobsDefeated;
                totalTimePlayed += levelStats.timeToComplete;

                console.log(`  Completed Level ${levelId} - ${levelStats.difficultyRating} - ${levelStats.wordsPerMinute.toFixed(2)} WPM - ${levelStats.accuracy.toFixed(2)}% accuracy`);

                // engine.reset() is called at the start of the next level's loop iteration.
                // No need to call it explicitly here unless transitioning to a menu.
            }

            // Complete world stats
            worldStats.endTime = Date.now();
            worldStats.totalTimeToComplete = worldTotalTime;
            worldStats.totalMobsDefeated = worldMobsDefeated;
            worldStats.averageLevelCompletion = worldTotalTime / worldStats.levels.length;
            worldStats.keystrokes = worldKeystrokes;
            worldStats.accuracy = (worldCorrectKeystrokes / worldKeystrokes) * 100;
            worldStats.wordsPerMinute = (worldTotalWords / (worldTotalTime / 1000 / 60));

            // Complete player performance metrics for this world
            worldStats.playerPerformance = {
                totalKeystrokes: worldKeystrokes,
                accuracy: (worldCorrectKeystrokes / worldKeystrokes) * 100,
                wordsPerMinute: (worldTotalWords / (worldTotalTime / 1000 / 60)),
                avgCombo: worldComboSum / worldStats.levels.length,
                highestCombo: worldComboHighest,
                timeSpent: worldTotalTime
            };

            // Add world stats to journey stats
            journeyStats.worldProgress.push(worldStats);

            // Add world difficulty to journey stats
            journeyStats.difficultyByWorld[`World ${world.id}: ${world.name}`] = worldStats.difficultyProgression
                .reduce((acc, curr, idx, arr) => {
                    return acc + (idx === 0 ? '' : ', ') + curr + (idx === arr.length - 1 ? ' (Boss)' : '');
                }, '');

            console.log(`Completed World ${world.id}: ${world.name}`);
            console.log(`  Total Score: ${worldStats.totalScore}`);
            console.log(`  Average WPM: ${worldStats.wordsPerMinute.toFixed(2)}`);
            console.log(`  Accuracy: ${worldStats.accuracy.toFixed(2)}%`);
            console.log(`  Time to Complete: ${(worldStats.totalTimeToComplete / 1000).toFixed(2)} seconds`);

            // Calculate skill progression for this world to measure player improvement
            journeyStats.skillProgression[`World ${world.id}`] =
                (worldStats.wordsPerMinute * (worldStats.accuracy / 100)) /
                (worldStats.totalTimeToComplete / 1000 / 60);
        }

        // Return to main menu
        stateManager.setGameStatus('mainMenu');

        // Complete journey stats
        const journeyEndTime = Date.now();
        journeyStats.totalTimeToComplete = journeyEndTime - journeyStartTime;
        journeyStats.totalLevelsCompleted = stateManager.getState().progression.completedLevels.length;
        journeyStats.totalWordsTyped = totalWords;
        journeyStats.overallAccuracy = (totalCorrectKeystrokes / totalKeystrokes) * 100;
        journeyStats.overallWPM = (totalWords / (totalTimePlayed / 1000 / 60));
        journeyStats.finalStats = getPlayerStats(engine);

        // Output comprehensive journey stats
        console.log('\n============ GAME JOURNEY STATISTICS ============');
        console.log('OVERALL PERFORMANCE:');
        console.log('Total Worlds Completed:', journeyStats.worldProgress.length);
        console.log('Total Levels Completed:', journeyStats.totalLevelsCompleted);
        console.log('Total Words Typed:', journeyStats.totalWordsTyped);
        console.log('Overall Accuracy:', journeyStats.overallAccuracy.toFixed(2) + '%');
        console.log('Average WPM:', journeyStats.overallWPM.toFixed(2));
        console.log('Total Time to Complete:', (journeyStats.totalTimeToComplete / 1000 / 60).toFixed(2), 'minutes');

        console.log('\nPROGRESSION BY WORLD:');
        for (const world of journeyStats.worldProgress) {
            console.log(`World ${world.worldId} (${world.worldName}):`);
            console.log(`  Score: ${world.totalScore}`);
            console.log(`  WPM: ${world.wordsPerMinute.toFixed(2)}`);
            console.log(`  Accuracy: ${world.accuracy.toFixed(2)}%`);
            console.log(`  Skill Rating: ${journeyStats.skillProgression['World ' + world.worldId].toFixed(2)}`);
        }

        console.log('\nDIFFICULTY PROGRESSION:');
        for (const [world, difficulty] of Object.entries(journeyStats.difficultyByWorld)) {
            console.log(`${world}: ${difficulty}`);
        }

        console.log('\nSKILL IMPROVEMENT:');
        const worldIds = Object.keys(journeyStats.skillProgression).sort();
        const firstWorld = worldIds[0];
        const lastWorld = worldIds[worldIds.length - 1];
        const skillImprovement = ((journeyStats.skillProgression[lastWorld] / journeyStats.skillProgression[firstWorld]) - 1) * 100;
        console.log(`Skill Improvement from ${firstWorld} to ${lastWorld}: ${skillImprovement.toFixed(2)}%`);

        // Expectations for final state
        const totalConfiguredLevels = configuredWorlds.reduce((acc, w) => acc + w.levels.length, 0);
        expect(stateManager.getState().progression.completedLevels.length).toBe(totalConfiguredLevels);
        expect(stateManager.getState().progression.unlockedWorlds.length).toBe(configuredWorlds.length);
        expect(journeyStats.overallAccuracy).toBeGreaterThan(80); // Expect decent accuracy

        // Return the journey stats for analysis
        return journeyStats;
    });
});
