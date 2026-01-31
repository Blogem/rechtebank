import { describe, it, expect } from 'vitest';
import { rotateRight } from './rotation';

describe('rotation utilities', () => {
    describe('rotateRight', () => {
        it('should rotate from 0° to 90°', () => {
            expect(rotateRight(0)).toBe(90);
        });

        it('should rotate from 90° to 180°', () => {
            expect(rotateRight(90)).toBe(180);
        });

        it('should rotate from 180° to 270°', () => {
            expect(rotateRight(180)).toBe(270);
        });

        it('should rotate from 270° to 0°', () => {
            expect(rotateRight(270)).toBe(0);
        });

        it('should cycle correctly through all angles', () => {
            let angle = 0;
            angle = rotateRight(angle); // 90
            expect(angle).toBe(90);
            angle = rotateRight(angle); // 180
            expect(angle).toBe(180);
            angle = rotateRight(angle); // 270
            expect(angle).toBe(270);
            angle = rotateRight(angle); // 0
            expect(angle).toBe(0);
        });
    });
});
