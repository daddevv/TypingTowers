/**
 * Utility for generating words based on available letters.
 * This is used to create appropriate challenges for each level
 * based on the keys the player has learned.
 */

export default class WordGenerator {
    private availableLetters: string[];
    private usePseudoWords: boolean;

    constructor(availableLetters: string[], usePseudoWords: boolean = true) {
        this.availableLetters = availableLetters.map(letter => letter.toLowerCase());
        this.usePseudoWords = usePseudoWords;
    }

    /**
     * Generates a word of the specified length using only available letters.
     * @param length Desired word length
     * @returns A string of the requested length
     */
    generateWord(length: number = 3): string {
        if (this.availableLetters.length === 0) {
            throw new Error("Cannot generate a word: availableLetters is empty.");
        }
        let word = '';
        for (let i = 0; i < length; i++) {
            const randomIndex = Math.floor(Math.random() * this.availableLetters.length);
            word += this.availableLetters[randomIndex];
        }
        return word;
    }

    /**
     * Generates a set of words for a level challenge.
     * @param count Number of words to generate
     * @param minLength Minimum word length
     * @param maxLength Maximum word length
     * @returns Array of words
     */
    generateWordSet(count: number, minLength: number = 2, maxLength: number = 5): string[] {
        const wordSet: string[] = [];
        for (let i = 0; i < count; i++) {
            const length = Math.floor(Math.random() * (maxLength - minLength + 1)) + minLength;
            wordSet.push(this.generateWord(length));
        }
        return wordSet;
    }

    /**
     * Checks if a given word can be created using the available letters.
     * @param word The word to check
     * @returns True if the word only contains available letters
     */
    canCreateWord(word: string): boolean {
        const lowerWord = word.toLowerCase();
        for (let i = 0; i < lowerWord.length; i++) {
            if (!this.availableLetters.includes(lowerWord[i])) {
                return false;
            }
        }
        return true;
    }

    /**
     * Filters a list of words to only include those that can be created
     * with the available letters.
     * @param words Array of words to filter
     * @returns Array of valid words
     */
    filterValidWords(words: string[]): string[] {
        return words.filter(word => this.canCreateWord(word));
    }

    /**
     * Creates pronounceable pseudo-words that follow basic phonetic rules.
     * Useful when real words can't be formed from limited letter sets.
     * @param length Desired word length
     * @returns A pronounceable pseudo-word
     */
    generatePseudoWord(length: number = 3): string {
        if (length < 2) return this.generateWord(length);

        // Separate letters into vowels and consonants
        const vowels = this.availableLetters.filter(l => ['a', 'e', 'i', 'o', 'u'].includes(l));
        const consonants = this.availableLetters.filter(l => !['a', 'e', 'i', 'o', 'u'].includes(l));

        // If we don't have both vowels and consonants, fall back to random generation
        if (vowels.length === 0 || consonants.length === 0) {
            return this.generateWord(length);
        }

        let word = '';
        let useVowel = Math.random() > 0.5; // Randomly start with vowel or consonant

        for (let i = 0; i < length; i++) {
            if (useVowel) {
                // Add a vowel
                const randomVowelIndex = Math.floor(Math.random() * vowels.length);
                word += vowels[randomVowelIndex];
            } else {
                // Add a consonant
                const randomConsonantIndex = Math.floor(Math.random() * consonants.length);
                word += consonants[randomConsonantIndex];
            }

            // Alternate between vowels and consonants for better pronounceability
            // But occasionally allow double consonants or vowels
            if (Math.random() > 0.2) {
                useVowel = !useVowel;
            }
        }

        return word;
    }

    /**
     * Gets a word appropriate for the current level.
     * Uses pseudo-words if real words aren't possible.
     * @param preferredLength Preferred word length
     * @returns A word using only available letters
     */
    getWord(preferredLength: number = 3): string {
        if (this.usePseudoWords) {
            return this.generatePseudoWord(preferredLength);
        } else {
            return this.generateWord(preferredLength);
        }
    }
}

//Contains AI - generated edits.
