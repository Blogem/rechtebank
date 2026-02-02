## ADDED Requirements

### Requirement: Frontend constructs shareable URLs with /vonnis/ path
The system SHALL construct shareable URLs using the `/vonnis/[id]` path format when sharing verdicts.

#### Scenario: Generate shareable URL in frontend
- **WHEN** user triggers the share functionality from the verdict display
- **THEN** system constructs the shareable URL as `${origin}/vonnis/${id}` where id is the shareable verdict ID from backend

#### Scenario: Share URL uses vonnis path regardless of current route
- **WHEN** user shares a verdict while viewing it on `/verdict/[id]` route
- **THEN** generated shareable URL uses `/vonnis/[id]` format, not `/verdict/[id]`

#### Scenario: Web Share API receives vonnis URL
- **WHEN** user shares via native Web Share API
- **THEN** the URL passed to navigator.share() contains `/vonnis/[id]` path

#### Scenario: Copy to clipboard contains vonnis URL
- **WHEN** user copies the shareable link
- **THEN** clipboard contains URL with `/vonnis/[id]` path format
