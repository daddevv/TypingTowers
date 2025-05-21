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
});
// Contains AI-generated edits.
