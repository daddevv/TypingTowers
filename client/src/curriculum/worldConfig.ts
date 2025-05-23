/**
 * Defines the structure and progression of worlds and levels in the game.
 * Each world focuses on a specific finger group, with progressive levels
 * introducing new keys in a structured manner.
 */

export interface LevelConfig {
    id: string;
    name: string;
    description: string;
    availableKeys: string[];
    targetWPM: number;
    enemySpeed: number;
    enemySpawnRate: number;
    mobsToDefeat: number; // Added: Number of mobs to defeat to win the level
    isBossLevel: boolean;
}

export interface WorldConfig {
    id: number;
    name: string;
    description: string;
    fingerGroup: string;
    levels: LevelConfig[];
}

export const WORLDS: WorldConfig[] = [
    {
        id: 1,
        name: "Index Finger World",
        description: "Master your index fingers, the most versatile digits for typing.",
        fingerGroup: "INDEX_FINGERS",
        levels: [
            {
                id: "1-1",
                name: "Home Base",
                description: "Learn the home row position for your index fingers: F and J.",
                availableKeys: ["f", "j"],
                targetWPM: 10,
                enemySpeed: 50,
                enemySpawnRate: 3000,
                mobsToDefeat: 10, // Added
                isBossLevel: false
            },
            {
                id: "1-2",
                name: "Extending Reach",
                description: "Add G and H to your index finger repertoire.",
                availableKeys: ["f", "j", "g", "h"],
                targetWPM: 15,
                enemySpeed: 60,
                enemySpawnRate: 2800,
                mobsToDefeat: 12, // Added
                isBossLevel: false
            },
            {
                id: "1-3",
                name: "Reaching Up",
                description: "Stretch your index fingers to the top row with R and U.",
                availableKeys: ["f", "j", "g", "h", "r", "u"],
                targetWPM: 18,
                enemySpeed: 70,
                enemySpawnRate: 2600,
                mobsToDefeat: 15, // Added
                isBossLevel: false
            },
            {
                id: "1-4",
                name: "Top Row Mastery",
                description: "Add T and Y (top row) and practice more complex patterns using all index finger keys so far.",
                availableKeys: ["f", "j", "g", "h", "r", "u", "t", "y"],
                targetWPM: 20,
                enemySpeed: 80,
                enemySpawnRate: 2400,
                mobsToDefeat: 18, // Added
                isBossLevel: false
            },
            {
                id: "1-5",
                name: "Bottom Row Drills",
                description: "Add V and M (bottom row) with drills for downward reaches. Emphasizes V/M usage with index finger keys.",
                availableKeys: ["f", "j", "g", "h", "r", "u", "t", "y", "v", "m"],
                targetWPM: 22,
                enemySpeed: 90,
                enemySpawnRate: 2200,
                mobsToDefeat: 20, // Added
                isBossLevel: false
            },
            {
                id: "1-6",
                name: "Complete Control",
                description: "Add the final index finger keys: B and N.",
                availableKeys: ["f", "j", "g", "h", "r", "u", "t", "y", "v", "m", "b", "n"],
                targetWPM: 25,
                enemySpeed: 100,
                enemySpawnRate: 2000,
                mobsToDefeat: 25, // Added
                isBossLevel: false
            },
            {
                id: "1-7",
                name: "Index Overlord",
                description: "Defeat the boss using all your index finger skills!",
                availableKeys: ["f", "j", "g", "h", "r", "u", "t", "y", "v", "m", "b", "n"],
                targetWPM: 30,
                enemySpeed: 120,
                enemySpawnRate: 1500,
                mobsToDefeat: 30, // Added
                isBossLevel: true
            }
        ]
    },
    {
        id: 2,
        name: "Middle Finger World",
        description: "Train your middle fingers to expand your typing reach.",
        fingerGroup: "MIDDLE_FINGERS",
        levels: [
            {
                id: "2-1",
                name: "Middle Position",
                description: "Learn the home row position for your middle fingers: D and K.",
                availableKeys: ["d", "k"],
                targetWPM: 10,
                enemySpeed: 50,
                enemySpawnRate: 3000,
                mobsToDefeat: 10, // Added
                isBossLevel: false
            },
            {
                id: "2-2",
                name: "Expanding Horizons",
                description: "Introduce E and I for a broader reach.",
                availableKeys: ["d", "k", "e", "i"],
                targetWPM: 12,
                enemySpeed: 55,
                enemySpawnRate: 2900,
                mobsToDefeat: 12, // Added
                isBossLevel: false
            },
            {
                id: "2-3",
                name: "Diagonal Drills",
                description: "Practice diagonal movements with C and ,.",
                availableKeys: ["d", "k", "e", "i", "c", ","],
                targetWPM: 15,
                enemySpeed: 60,
                enemySpawnRate: 2800,
                mobsToDefeat: 15, // Added
                isBossLevel: false
            },
            {
                id: "2-4",
                name: "Middle Row Mastery",
                description: "Master the middle row with complex patterns using all learned keys.",
                availableKeys: ["d", "k", "e", "i", "c", ",", "m", "b"],
                targetWPM: 20,
                enemySpeed: 70,
                enemySpawnRate: 2600,
                mobsToDefeat: 18, // Added
                isBossLevel: false
            },
            {
                id: "2-5",
                name: "Middle Master",
                description: "Defeat the middle finger boss!",
                availableKeys: ["d", "e", "c", "k", "i", ","],
                targetWPM: 30,
                enemySpeed: 120,
                enemySpawnRate: 1500,
                mobsToDefeat: 30, // Added
                isBossLevel: true
            }
        ]
    },
    {
        id: 3,
        name: "Ring Finger World",
        description: "Strengthen your ring fingers - often the weakest in typing.",
        fingerGroup: "RING_FINGERS",
        levels: [
            {
                id: "3-1",
                name: "Ring Position",
                description: "Learn the home row position for your ring fingers: S and L.",
                availableKeys: ["s", "l"],
                targetWPM: 10,
                enemySpeed: 50,
                enemySpawnRate: 3000,
                mobsToDefeat: 10, // Added
                isBossLevel: false
            },
            {
                id: "3-2",
                name: "Adding Depth",
                description: "Incorporate D and H for deeper reaches.",
                availableKeys: ["s", "l", "d", "h"],
                targetWPM: 12,
                enemySpeed: 55,
                enemySpawnRate: 2900,
                mobsToDefeat: 12, // Added
                isBossLevel: false
            },
            {
                id: "3-3",
                name: "Complex Patterns",
                description: "Practice complex patterns with all learned keys.",
                availableKeys: ["s", "l", "d", "h", "r", "u", "t", "y"],
                targetWPM: 18,
                enemySpeed: 70,
                enemySpawnRate: 2600,
                mobsToDefeat: 15, // Added
                isBossLevel: false
            },
            {
                id: "3-4",
                name: "Ring Mastery",
                description: "Master ring finger movements with advanced drills.",
                availableKeys: ["s", "l", "d", "h", "r", "u", "t", "y", "v", "m"],
                targetWPM: 22,
                enemySpeed: 90,
                enemySpawnRate: 2200,
                mobsToDefeat: 20, // Added
                isBossLevel: false
            },
            {
                id: "3-5",
                name: "Ring Champion",
                description: "Defeat the ring finger boss!",
                availableKeys: ["s", "w", "x", "l", "o", "."],
                targetWPM: 30,
                enemySpeed: 120,
                enemySpawnRate: 1500,
                mobsToDefeat: 30, // Added
                isBossLevel: true
            }
        ]
    },
    {
        id: 4,
        name: "Pinky Finger World",
        description: "Master your pinky fingers to complete your typing skills.",
        fingerGroup: "PINKY_FINGERS",
        levels: [
            {
                id: "4-1",
                name: "Pinky Position",
                description: "Learn the home row position for your pinky fingers: A and semicolon.",
                availableKeys: ["a", ";"],
                targetWPM: 10,
                enemySpeed: 50,
                enemySpawnRate: 3000,
                mobsToDefeat: 10, // Added
                isBossLevel: false
            },
            {
                id: "4-2",
                name: "Uppercase Challenge",
                description: "Introduce Q and P for uppercase challenges.",
                availableKeys: ["a", ";", "q", "p"],
                targetWPM: 12,
                enemySpeed: 55,
                enemySpawnRate: 2900,
                mobsToDefeat: 12, // Added
                isBossLevel: false
            },
            {
                id: "4-3",
                name: "Symbol Surge",
                description: "Incorporate Z and / for symbol-rich drills.",
                availableKeys: ["a", ";", "q", "p", "z", "/"],
                targetWPM: 15,
                enemySpeed: 60,
                enemySpawnRate: 2800,
                mobsToDefeat: 15, // Added
                isBossLevel: false
            },
            {
                id: "4-4",
                name: "Pinky Mastery",
                description: "Master pinky movements with advanced drills.",
                availableKeys: ["a", ";", "q", "p", "z", "/"],
                targetWPM: 20,
                enemySpeed: 70,
                enemySpawnRate: 2600,
                mobsToDefeat: 18, // Added
                isBossLevel: false
            },
            {
                id: "4-5",
                name: "Typing Conqueror",
                description: "Conquer the typing field using all learned keys!",
                availableKeys: ["a", ";", "q", "p", "z", "/"],
                targetWPM: 30,
                enemySpeed: 120,
                enemySpawnRate: 1500,
                mobsToDefeat: 30, // Added
                isBossLevel: true
            }
        ]
    }
];

// Helper to get a specific level by ID
export function getLevelById(worldId: number, levelId: string): LevelConfig | undefined {
    const world = WORLDS.find(w => w.id === worldId);
    if (!world) return undefined;

    return world.levels.find(l => l.id === levelId);
}

// Get all available keys up to a specific level
export function getAvailableKeysForLevel(worldId: number, levelId: string): string[] {
    const allKeys: string[] = [];

    for (const world of WORLDS) {
        if (world.id > worldId) break;

        for (const level of world.levels) {
            if (world.id === worldId && level.id > levelId) break;

            // Add unique keys
            level.availableKeys.forEach(key => {
                if (!allKeys.includes(key)) {
                    allKeys.push(key);
                }
            });
        }
    }

    return allKeys;
}

// Contains AI-generated edits.
