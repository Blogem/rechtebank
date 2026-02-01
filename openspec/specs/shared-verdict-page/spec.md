## ADDED Requirements

### Requirement: Display shared verdict page
The system SHALL provide a public page that displays a verdict and its associated photo using a shareable URL.

#### Scenario: Access shared verdict via URL
- **WHEN** user navigates to `/verdict/:id` with a valid shareable ID
- **THEN** system displays a page with both the verdict text and the original photo

#### Scenario: Shared page matches original verdict display
- **WHEN** shared verdict page is displayed
- **THEN** page uses the same visual styling and layout as the original verdict display shown after upload

#### Scenario: Photo display in shared verdict
- **WHEN** shared verdict page loads
- **THEN** system displays the original uploaded photo in a bordered frame above or beside the verdict text

#### Scenario: Responsive photo sizing
- **WHEN** shared verdict page is viewed on different screen sizes
- **THEN** photo scales responsively while maintaining aspect ratio

#### Scenario: Shared page with invalid ID
- **WHEN** user navigates to `/verdict/:id` with invalid or malformed ID
- **THEN** system displays an error page with message "Invalid verdict link"

#### Scenario: Shared page for missing verdict
- **WHEN** user navigates to `/verdict/:id` for verdict files that don't exist
- **THEN** system displays an error page with message "Verdict not found"

#### Scenario: Shared page for expired verdict
- **WHEN** user navigates to `/verdict/:id` for verdict cleaned up by retention policy
- **THEN** system displays an error page with message "This verdict is no longer available"

### Requirement: Server-side rendering for social sharing
The system SHALL render the shared verdict page on the server to enable social media previews.

#### Scenario: Open Graph meta tags
- **WHEN** shared verdict page is requested
- **THEN** server includes Open Graph meta tags (og:title, og:description, og:image) in the HTML response

#### Scenario: Social media preview image
- **WHEN** shared link is posted on social media platforms
- **THEN** platform displays a preview with the verdict photo and title

#### Scenario: Twitter card meta tags
- **WHEN** shared verdict page is requested
- **THEN** server includes Twitter card meta tags for enhanced Twitter previews

### Requirement: Share verdict button
The system SHALL provide a button to generate and share verdict URLs.

#### Scenario: Share button on verdict display
- **WHEN** verdict is displayed after upload or on shared page
- **THEN** system displays a "Deel Vonnis" (Share Verdict) button

#### Scenario: Generate shareable link
- **WHEN** user clicks "Deel Vonnis" button
- **THEN** system calls backend to generate shareable ID and constructs full URL

#### Scenario: Native share on mobile
- **WHEN** user clicks share button on mobile device with Web Share API support
- **THEN** system opens native share dialog with verdict title, description, and URL

#### Scenario: Clipboard fallback on desktop
- **WHEN** user clicks share button on device without Web Share API support
- **THEN** system copies the shareable URL to clipboard and shows confirmation message

#### Scenario: Share fails gracefully
- **WHEN** both native share and clipboard copy fail
- **THEN** system displays the URL in a dialog for manual copy

#### Scenario: Share includes verdict preview
- **WHEN** native share dialog is shown
- **THEN** share data includes verdict text preview and score in the description
