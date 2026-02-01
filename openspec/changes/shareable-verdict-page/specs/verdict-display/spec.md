## ADDED Requirements

### Requirement: Display photo with verdict
The system SHALL display the uploaded photo alongside the verdict text when photo data is provided.

#### Scenario: Photo display when available
- **WHEN** verdict display component receives image data
- **THEN** system displays the photo in a bordered frame above or beside the verdict content

#### Scenario: Verdict display without photo
- **WHEN** verdict display component does not receive image data
- **THEN** system displays only the verdict content without photo section (backward compatible)

#### Scenario: Photo responsive sizing
- **WHEN** photo is displayed with verdict
- **THEN** system ensures photo scales appropriately for different screen sizes while maintaining aspect ratio

#### Scenario: Photo loading state
- **WHEN** photo data is being loaded
- **THEN** system displays a placeholder or loading indicator in the photo area

## MODIFIED Requirements

### Requirement: Share verdict option
The system SHALL provide options to share the verdict with a shareable URL.

#### Scenario: Share verdict button
- **WHEN** verdict is displayed
- **THEN** system provides a "Deel Vonnis" (Share Verdict) button that generates a shareable URL

#### Scenario: Share via native dialog
- **WHEN** user clicks share button on mobile device
- **THEN** system opens native share dialog with verdict URL and preview text

#### Scenario: Copy link to clipboard
- **WHEN** user clicks share button on desktop or native share is unavailable
- **THEN** system copies the shareable verdict URL to clipboard and shows confirmation

#### Scenario: Share includes verdict preview
- **WHEN** share dialog is opened
- **THEN** system includes verdict title, score, and a text preview in the share data
