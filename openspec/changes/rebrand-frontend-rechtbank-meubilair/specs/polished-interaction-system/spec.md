## ADDED Requirements

### Requirement: Visual depth through shadow system
The system SHALL implement a layered shadow system to provide visual depth and hierarchy, avoiding the flat appearance of the previous design.

#### Scenario: Shadow tokens are defined
- **WHEN** the application styles are loaded
- **THEN** CSS variables SHALL be available for shadow levels:
  - `--shadow-sm` for subtle elements
  - `--shadow-base` for resting cards and surfaces
  - `--shadow-md` for elevated or hovered elements
  - `--shadow-lg` for modals and prominent overlays

#### Scenario: Cards use layered shadows
- **WHEN** a card component (intro, case submission, verdict, error) is rendered
- **THEN** it SHALL use `--shadow-base` in its resting state
- **AND** the shadow SHALL be a layered shadow combining ambient and directional light (not a single flat shadow)

#### Scenario: Interactive elements increase shadow on hover
- **WHEN** a user hovers over an interactive card or button
- **THEN** the shadow SHALL smoothly transition to a deeper level (e.g., base → md)
- **AND** the transition SHALL use the defined interactive timing (300ms ease-out)

### Requirement: Button micro-interactions
The system SHALL implement polished button interactions with hover, active, and focus states to provide tactile feedback and avoid the dated static feel.

#### Scenario: Button hover state provides lift feedback
- **WHEN** a user hovers over a primary button (e.g., "Foto bevestigen", "Nieuwe zaak")
- **THEN** the button SHALL translate upward by 2-4px (translateY)
- **AND** the button shadow SHALL increase from base to md level
- **AND** the button background color MAY subtly lighten or shift
- **AND** all transitions SHALL complete within 200-300ms using ease-out timing

#### Scenario: Button active state provides press feedback
- **WHEN** a user presses down on a button (active state)
- **THEN** the button SHALL scale down slightly (e.g., scale(0.98))
- **AND** the shadow SHALL reduce to indicate depression
- **AND** the transition SHALL be immediate (< 100ms)

#### Scenario: Button focus state is visually designed
- **WHEN** a button receives keyboard focus
- **THEN** it SHALL display a custom focus ring using the court accent color
- **AND** the focus ring SHALL use offset shadow styling (not default browser outline)
- **AND** the focus ring SHALL be clearly visible against all background colors

### Requirement: Verdict reveal animation
The system SHALL animate the appearance of verdict content to add gravitas to the judgment moment without being playful or distracting.

#### Scenario: Verdict sections appear in sequence
- **WHEN** a verdict is displayed after upload completion
- **THEN** the verdict sections SHALL fade in sequentially:
  - First: Photo evidence/header
  - Then: Feiten section (100-200ms delay)
  - Then: Overwegingen section (100-200ms delay)
  - Finally: Uitspraak section (100-200ms delay)
- **AND** each section SHALL fade in with a slight upward slide (translateY(8px) → 0)
- **AND** each section transition SHALL use 400ms ease-out timing

#### Scenario: Court seal has subtle presence on verdict load
- **WHEN** a verdict is first displayed
- **THEN** the court seal (if present in the verdict view) MAY have a subtle animation
- **AND** the animation SHALL be restrained (e.g., fade-in or small rotation)
- **AND** the animation SHALL complete within 600ms
- **AND** the animation SHALL not distract from the content reveal

### Requirement: Page transition smoothness
The system SHALL implement smooth content transitions when app state changes to provide continuity and avoid jarring visual jumps.

#### Scenario: Content fades in on state change
- **WHEN** the app transitions between states (e.g., camera → uploading → verdict, or camera → error)
- **THEN** the new content SHALL fade in from slight opacity (0.8 → 1.0)
- **AND** the content SHALL slide up slightly during fade-in (translateY(8px) → 0)
- **AND** the transition SHALL use 300-400ms ease-out timing

#### Scenario: Previous content fades out before new content appears
- **WHEN** transitioning between major app states
- **THEN** the outgoing content MAY fade out before the incoming content appears
- **AND** fade-out SHALL be faster than fade-in (150-200ms)
- **AND** there SHALL be minimal delay between fade-out and fade-in (< 50ms) to maintain responsiveness

### Requirement: Transition timing consistency
The system SHALL use consistent, purposeful timing values for all animations to create a cohesive interaction language.

#### Scenario: Standard timing tokens are defined
- **WHEN** the application styles are loaded
- **THEN** CSS variables or constants SHALL define:
  - Default transition: 200ms ease-out (colors, small changes)
  - Interactive transition: 300ms ease-out (transforms, shadows, hover states)
  - Reveal transition: 400ms ease-out (content appearing, state changes)
  - Quick feedback: 100ms or less (active/press states)

#### Scenario: All interactive elements use standard timings
- **WHEN** any UI element implements a transition or animation
- **THEN** it SHALL use one of the defined standard timing values
- **AND** it SHALL NOT use arbitrary timing values (e.g., 250ms, 350ms)
- **AND** it SHALL prefer ease-out easing for most interactions (entrance, expansion)

### Requirement: Performance and reduced motion
The system SHALL respect user preferences and maintain performance while providing polished interactions.

#### Scenario: Animations respect prefers-reduced-motion
- **WHEN** a user has `prefers-reduced-motion: reduce` set in their browser
- **THEN** all non-essential animations SHALL be disabled or reduced to instant transitions
- **AND** essential feedback (e.g., focus rings, state changes) SHALL remain visible but without motion

#### Scenario: Animations use CSS transforms for performance
- **WHEN** implementing position or scale changes
- **THEN** the implementation SHALL use CSS `transform` properties (translateY, scale)
- **AND** it SHALL NOT animate `top`, `left`, `margin`, or other layout properties
- **AND** it SHALL NOT cause layout thrashing or reflows during animation
