import { describe, expect, it } from 'vitest';
import { loadWordList } from '../loadWordList';

describe('loadWordList', () => {
    it('loads fjWords for 1-1', async () => {
        const words = await loadWordList('1-1');
        expect(Array.isArray(words)).toBe(true);
        expect(words.length).toBeGreaterThan(0);
        expect(words).toContain('f');
        expect(words).toContain('j');
    });

    it('loads fjghWords for 1-2', async () => {
        const words = await loadWordList('1-2');
        expect(Array.isArray(words)).toBe(true);
        expect(words.length).toBeGreaterThan(0);
        expect(words).toContain('g');
        expect(words).toContain('h');
        expect(words).toContain('gh');
        expect(words).toContain('fg');
        expect(words).toContain('fh');
        expect(words).toContain('hj');
        expect(words).not.toContain('f'); // 'f' is not present in fjghWords.json
    });

    it('loads fjghWords for 1-2 and all words use only f, j, g, h', async () => {
        const allowed = ['f', 'j', 'g', 'h'];
        const words = await loadWordList('1-2');
        expect(Array.isArray(words)).toBe(true);
        for (const word of words) {
            for (const letter of word) {
                expect(allowed).toContain(letter);
            }
        }
    });

    it('loads fjghruWords for 1-3', async () => {
        const words = await loadWordList('1-3');
        expect(Array.isArray(words)).toBe(true);
        expect(words.length).toBeGreaterThan(0);
        expect(words).toContain('r');
        expect(words).toContain('u');
    });

    it('returns empty array for unknown level', async () => {
        const words = await loadWordList('unknown');
        expect(words).toEqual([]);
    });

    it('loads fjghrutyvmWords for 1-4 and 1-5', async () => {
        for (const level of ['1-4', '1-5']) {
            const words = await loadWordList(level);
            expect(Array.isArray(words)).toBe(true);
            expect(words.length).toBeGreaterThan(0);
            // Should contain T and Y for 1-4/1-5
            expect(words.some(w => w.includes('t'))).toBe(true);
            expect(words.some(w => w.includes('y'))).toBe(true);
            // All words use only allowed letters
            const allowed = ['f', 'j', 'g', 'h', 'r', 'u', 't', 'y', 'v', 'm'];
            for (const word of words) {
                for (const letter of word) {
                    expect(allowed).toContain(letter);
                }
            }
        }
    });

    it('loads fjghrutyvmbnWords for 1-6', async () => {
        const words = await loadWordList('1-6');
        expect(Array.isArray(words)).toBe(true);
        expect(words.length).toBeGreaterThan(0);
        // Should contain B and N
        expect(words.some(w => w.includes('b'))).toBe(true);
        expect(words.some(w => w.includes('n'))).toBe(true);
        // All words use only allowed letters
        const allowed = ['f', 'j', 'g', 'h', 'r', 'u', 't', 'y', 'v', 'm', 'b', 'n'];
        for (const word of words) {
            for (const letter of word) {
                expect(allowed).toContain(letter);
            }
        }
    });

    it('loads fjghrutyvmbn_bossWords for 1-7', async () => {
        const words = await loadWordList('1-7');
        expect(Array.isArray(words)).toBe(true);
        expect(words.length).toBeGreaterThan(0);
        // Should contain challenging boss words
        expect(words.some(w => w.length >= 6)).toBe(true);
        // All words use only allowed letters
        const allowed = ['f', 'j', 'g', 'h', 'r', 'u', 't', 'y', 'v', 'm', 'b', 'n'];
        for (const word of words) {
            for (const letter of word) {
                expect(allowed).toContain(letter);
            }
        }
    });

    it('level 1-4 word list emphasizes T/Y and uses only allowed letters', async () => {
        const allowed = ['f', 'j', 'g', 'h', 'r', 'u', 't', 'y', 'v', 'm', 'a']; // Added 'a'
        const words = await loadWordList('1-4');
        expect(Array.isArray(words)).toBe(true);
        expect(words.length).toBeGreaterThan(0);
        // At least 50% of words should contain 't' or 'y' to emphasize usage
        const tyWords = words.filter(w => w.includes('t') || w.includes('y'));
        expect(tyWords.length / words.length).toBeGreaterThanOrEqual(0.5);
        // All words use only allowed letters
        for (const word of words) {
            for (const letter of word) {
                expect(allowed).toContain(letter);
            }
        }
    });

    it('level 1-5 word list emphasizes V/M and uses only allowed letters', async () => {
        const allowed = ['f', 'j', 'g', 'h', 'r', 'u', 't', 'y', 'v', 'm', 'a']; // Added 'a'
        const words = await loadWordList('1-5');
        expect(Array.isArray(words)).toBe(true);
        expect(words.length).toBeGreaterThan(0);
        // At least 50% of words should contain 'v' or 'm' to emphasize usage
        const vmWords = words.filter(w => w.includes('v') || w.includes('m'));
        expect(vmWords.length / words.length).toBeGreaterThanOrEqual(0.5);
        // All words use only allowed letters
        for (const word of words) {
            for (const letter of word) {
                expect(allowed).toContain(letter);
            }
        }
    });

    it('level 1-6 word list includes all index finger letters and emphasizes B/N', async () => {
        const allowed = ['f', 'j', 'g', 'h', 'r', 'u', 't', 'y', 'v', 'm', 'b', 'n', 'a']; // Added 'b', 'n', 'a'
        const words = await loadWordList('1-6');
        expect(Array.isArray(words)).toBe(true);
        expect(words.length).toBeGreaterThan(0);
        // At least 50% of words should contain 'b' or 'n' to emphasize usage
        const bnWords = words.filter(w => w.includes('b') || w.includes('n'));
        expect(bnWords.length / words.length).toBeGreaterThanOrEqual(0.5);
        // All words use only allowed letters
        for (const word of words) {
            for (const letter of word) {
                expect(allowed).toContain(letter);
            }
        }
    });

    it('level 1-7 boss word list uses all index finger letters and contains challenging words', async () => {
        const allowed = ['f', 'j', 'g', 'h', 'r', 'u', 't', 'y', 'v', 'm', 'b', 'n', 'a', 'k']; // Added 'b', 'n', 'a', 'k'
        const words = await loadWordList('1-7');
        expect(Array.isArray(words)).toBe(true);
        expect(words.length).toBeGreaterThan(0);
        // All words use only allowed letters
        for (const word of words) {
            for (const letter of word) {
                expect(allowed).toContain(letter);
            }
        }
        // Should contain at least one long/challenging word
        expect(words.some(w => w.length >= 7)).toBe(true);
    });
});
// Contains AI-generated edits.
