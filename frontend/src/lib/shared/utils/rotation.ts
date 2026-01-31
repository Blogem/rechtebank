/**
 * Rotate counter-clockwise by 90°.
 * 
 * @param current - Current rotation angle (0, 90, 180, or 270)
 * @returns New rotation angle after rotating left
 * @example
 * rotateLeft(0)   // returns 270
 * rotateLeft(90)  // returns 0
 */
export function rotateLeft(current: number): number {
    return (current - 90 + 360) % 360;
}

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
