// Utility to load a word list JSON file for a level
export async function loadWordList(levelId: string): Promise<string[]> {
    switch (levelId) {
        case "1-1":
            return import('../wordpacks/fjWords.json').then(m => m.default);
        case "1-2":
            return import('../wordpacks/fjghWords.json').then(m => Array.isArray(m.default) ? m.default : []);
        case "1-3":
            return import('../wordpacks/fjghruWords.json').then(m => Array.isArray(m.default) ? m.default : []);
        case "1-4":
        case "1-5":
            return import('../wordpacks/fjghrutyvmWords.json').then(m => Array.isArray(m.default) ? m.default as string[] : []);
        case "1-6":
            return import('../wordpacks/fjghrutyvmbnWords.json').then(m => Array.isArray(m.default) ? m.default as string[] : []);
        case "1-7":
            return import('../wordpacks/fjghrutyvmbn_bossWords.json').then(m => Array.isArray(m.default) ? m.default as string[] : []);
        // Add more cases for other levels as needed
        default:
            return [];
    }
}
// Contains AI-generated edits.
