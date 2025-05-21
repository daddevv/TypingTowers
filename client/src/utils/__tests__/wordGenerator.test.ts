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
    });

    it('filters valid words', () => {
        const gen = new WordGenerator(['f', 'j']);
        const valid = gen.filterValidWords(['fj', 'abc', 'jf']);
        expect(valid).toEqual(['fj', 'jf']);
    });

    it('generates pronounceable pseudo-words', () => {
        const gen = new WordGenerator(['f', 'j', 'u']);
        const word = gen.generatePseudoWord(4);
        expect(word.length).toBe(4);
    });

    it('gets a word using getWord', () => {
        const gen = new WordGenerator(['f', 'j']);
        const word = gen.getWord(3);
        expect(word.length).toBe(3);
    });
});
// Contains AI-generated edits.
