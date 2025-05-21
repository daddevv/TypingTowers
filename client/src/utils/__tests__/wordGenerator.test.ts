import { describe, expect, it } from 'vitest';
import WordGenerator from '../wordGenerator';

describe('WordGenerator', () => {
    it('generates a word of the correct length', () => {
        const gen = new WordGenerator(['f', 'j']);
        const word = gen.generateWord(4);
        expect(word.length).toBe(4);
        expect(['f', 'j']).toContain(word[0]);
    });

    it('throws if no available letters', () => {
        expect(() => new WordGenerator([]).generateWord(3)).toThrow();
        expect(() => new WordGenerator([]).generatePseudoWord(3)).toThrow();
    });

    it('filters valid words', () => {
        const gen = new WordGenerator(['f', 'j']);
        const valid = gen.filterValidWords(['fj', 'abc', 'jf']);
        expect(valid).toEqual(['fj', 'jf']);
    });

    it('canCreateWord returns true for valid, false for invalid', () => {
        const gen = new WordGenerator(['a', 'b']);
        expect(gen.canCreateWord('ab')).toBe(true);
        expect(gen.canCreateWord('abc')).toBe(false);
    });

    it('generates pronounceable pseudo-words', () => {
        const gen = new WordGenerator(['f', 'j', 'u']);
        const word = gen.generatePseudoWord(4);
        expect(word.length).toBe(4);
        // Should only use available letters
        expect([...word].every(l => ['f', 'j', 'u'].includes(l))).toBe(true);
    });

    it('generates pseudo-words with only vowels', () => {
        const gen = new WordGenerator(['a', 'e']);
        const word = gen.generatePseudoWord(5);
        expect([...word].every(l => ['a', 'e'].includes(l))).toBe(true);
    });

    it('generates pseudo-words with only consonants', () => {
        const gen = new WordGenerator(['b', 'c']);
        const word = gen.generatePseudoWord(5);
        expect([...word].every(l => ['b', 'c'].includes(l))).toBe(true);
    });

    it('getWord uses pseudo-words if enabled', () => {
        const gen = new WordGenerator(['f', 'j'], true);
        const word = gen.getWord(3);
        expect(word.length).toBe(3);
    });

    it('getWord uses random letters if pseudo-words disabled', () => {
        const gen = new WordGenerator(['f', 'j'], false);
        const word = gen.getWord(3);
        expect(word.length).toBe(3);
    });

    it('generateWordSet returns correct number and lengths', () => {
        const gen = new WordGenerator(['f', 'j']);
        const set = gen.generateWordSet(5, 2, 4);
        expect(set.length).toBe(5);
        set.forEach(w => {
            expect(w.length).toBeGreaterThanOrEqual(2);
            expect(w.length).toBeLessThanOrEqual(4);
        });
    });

    it('uses correct letters for level 1-2', () => {
        const allowed = ['f', 'j', 'g', 'h'];
        const gen = new WordGenerator(allowed);
        const word = gen.generateWord(6);
        expect([...word].every(l => allowed.includes(l))).toBe(true);
    });

    it('handles zero/negative word length gracefully', () => {
        const gen = new WordGenerator(['f', 'j']);
        expect(gen.generateWord(0)).toBe('');
        expect(gen.generatePseudoWord(0)).toBe('');
        expect(gen.getWord(0)).toBe('');
        expect(gen.generateWord(-2)).toBe('');
    });
});

// Contains AI-generated edits.
