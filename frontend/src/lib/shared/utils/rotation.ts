/**
 * Rotate clockwise by 90°.
 * 
 * @param current - Current rotation angle (0, 90, 180, or 270)
 * @returns New rotation angle after rotating right
 * @example
 * rotateRight(0)    // returns 90
 * rotateRight(270)  // returns 0
 */
export function rotateRight(current: number): number {
    return (current + 90) % 360;
}
