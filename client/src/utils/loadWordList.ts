// Utility to load a word list JSON file for a level
export async function loadWordList(levelId: string): Promise<string[]> {
    switch (levelId) {
        case "1-1":
            return import('../wordpacks/fjWords.json').then(m => m.default);
        case "1-2":
            return import('../wordpacks/fjghWords.json').then(m => Array.isArray(m.default) ? m.default : []);
        // Add more cases for other levels as needed
        default:
            return [];
    }
}
// Contains AI-generated edits.
