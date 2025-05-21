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

    it('returns empty array for unknown level', async () => {
        const words = await loadWordList('unknown');
        expect(words).toEqual([]);
    });
});
// Contains AI-generated edits.
