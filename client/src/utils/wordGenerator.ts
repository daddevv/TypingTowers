/**
 * Utility for generating words based on available letters.
 * This is used to create appropriate challenges for each level
 * based on the keys the player has learned.
 */

export default class WordGenerator {
    private availableLetters: string[];
    private usePseudoWords: boolean;

    private minWordLength: number = 2;
    private maxWordLength: number = 5;

    constructor(availableLetters: string[], usePseudoWords: boolean = true) {
        this.availableLetters = availableLetters.map(letter => letter.toLowerCase());
        this.usePseudoWords = usePseudoWords;
    }

    /**
     * Generates a word of the specified length using only available letters.
     * If no length is provided, uses the current min/maxWordLength for scaling.
     */
    generateWord(length?: number): string {
        if (this.availableLetters.length === 0) {
            throw new Error("Cannot generate a word: availableLetters is empty.");
        }
        let word = '';
        const len = typeof length === 'number' ? length : (this.minWordLength + Math.floor(Math.random() * (this.maxWordLength - this.minWordLength + 1)));
        for (let i = 0; i < len; i++) {
            const idx = Math.floor(Math.random() * this.availableLetters.length);
            word += this.availableLetters[idx];
        }
        return word;
    }

    /**
     * Set word length scaling for dynamic difficulty.
     */
    setWordLengthScaling(min: number, max: number) {
        this.minWordLength = min;
        this.maxWordLength = max;
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
            const len = Math.floor(Math.random() * (maxLength - minLength + 1)) + minLength;
            wordSet.push(this.getWord(len));
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
     * Filters a list of words, returning only those that can be created with available letters.
     * @param words Array of words to filter
     * @returns Array of valid words
     */
    filterValidWords(words: string[]): string[] {
        return words.filter(word => this.canCreateWord(word));
    }

    /**
     * Generates a pronounceable pseudo-word using available letters (simple CVC pattern).
     * @param length Desired word length
     * @returns A pseudo-word string
     */
    generatePseudoWord(length: number = 4): string {
        // Simple implementation: alternate consonant/vowel if possible
        const vowels = ['a', 'e', 'i', 'o', 'u'].filter(v => this.availableLetters.includes(v));
        const consonants = this.availableLetters.filter(l => !vowels.includes(l));
        if (this.availableLetters.length === 0) {
            throw new Error("Cannot generate a pseudo-word: availableLetters is empty.");
        }
        let word = '';
        let useVowel = Math.random() > 0.5;
        for (let i = 0; i < length; i++) {
            if (useVowel && vowels.length > 0) {
                word += vowels[Math.floor(Math.random() * vowels.length)];
            } else if (consonants.length > 0) {
                word += consonants[Math.floor(Math.random() * consonants.length)];
            } else {
                // fallback if only vowels or only consonants
                word += this.availableLetters[Math.floor(Math.random() * this.availableLetters.length)];
            }
            useVowel = !useVowel;
        }
        return word;
    }

    /**
     * Returns a word, using pseudo-words if enabled, otherwise random letters.
     * @param length Desired word length
     * @returns A word string
     */
    getWord(length: number = 3): string {
        if (this.usePseudoWords) {
            return this.generatePseudoWord(length);
        }
        return this.generateWord(length);
    }
}

// Contains AI-generated edits.
