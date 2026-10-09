// Vitest test setup file
// The /vitest entry is what augments vitest's `Assertion` with the jest-dom
// matchers. The bare import only augments jest's, so under kit 3's stricter
// tsconfig every `toBeInTheDocument`/`toHaveClass` failed to type-check.
import '@testing-library/jest-dom/vitest';
