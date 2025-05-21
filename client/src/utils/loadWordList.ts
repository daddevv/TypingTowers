// Utility to load a word list JSON file for a level
export async function loadWordList(levelId: string): Promise<string[]> {
    switch (levelId) {
        case "1-1":
            return import('../wordpacks/fjWords.json').then(m => m.default);
        // Add more cases for other levels as needed
        default:
            return [];
    }
}
// Contains AI-generated edits.
