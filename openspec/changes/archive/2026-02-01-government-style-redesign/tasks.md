## 1. Global Styling and Layout

- [x] 1.1 Update `:global(body)` in +page.svelte to replace purple gradient with solid #fafafa background
- [x] 1.2 Update header styling to use solid #2e2e2e background with white text
- [x] 1.3 Add 3px solid #4a4a4a bottom border to header
- [x] 1.4 Update footer styling to use solid #2e2e2e background with white text
- [x] 1.5 Add 3px solid #4a4a4a top border to footer
- [x] 1.6 Update header h1 font-weight to 600 (semibold)

## 2. Court Seal Integration

- [x] 2.1 Move court-seal.png from `lib/assets/` to `src/lib/assets/` (Svelte convention)
- [x] 2.2 Import court seal image in +page.svelte
- [x] 2.3 Add court seal img element to header with 80px dimensions
- [x] 2.4 Apply absolute positioning to seal (left: 2rem, vertical center)
- [x] 2.5 Set seal opacity to 0.95
- [x] 2.6 Add responsive styling to reduce seal size to 60px on mobile (max-width: 768px)

## 3. Card Styling Updates

- [x] 3.1 Update .introduction card: change border-radius from 8px to 2px
- [x] 3.2 Update .introduction card: add 3px solid #4a4a4a top border
- [x] 3.3 Update .introduction card: change box-shadow to crisp style (0 1px 3px rgba(0,0,0,0.1))
- [x] 3.4 Update .camera-section card: change border-radius from 8px to 2px
- [x] 3.5 Update .camera-section card: add 3px solid #4a4a4a top border
- [x] 3.6 Update .camera-section card: change box-shadow to crisp style
- [x] 3.7 Update .welcome-text: increase line-height to 1.7
- [x] 3.8 Add horizontal divider (.section-divider) between introduction and camera sections

## 4. VerdictDisplay Component Updates

- [x] 4.1 Create helper function to generate case number (RVM-{year}-{timestamp})
- [x] 4.2 Create helper function to format Dutch timestamp ("Uitspraak d.d.: {day} {month} {year}, {HH:mm}")
- [x] 4.3 Add case metadata section to verdict display markup (above verdict content)
- [x] 4.4 Display "Zaaknummer: {caseNumber}" in metadata section
- [x] 4.5 Display "Uitspraak d.d.: {formattedDate}" in metadata section
- [x] 4.6 Add horizontal rule separator (2px solid #d1d1d1) below metadata section
- [x] 4.7 Update verdict card border-radius from 8px to 2px
- [x] 4.8 Add 3px solid #4a4a4a top border to verdict card
- [x] 4.9 Update verdict card box-shadow to crisp style
- [x] 4.10 Update heading font-weights to 600 where appropriate

## 5. ErrorDisplay Component Updates

- [x] 5.1 Update error card border-radius from 8px to 2px
- [x] 5.2 Add 3px solid #4a4a4a top border to error card
- [x] 5.3 Update error card box-shadow to crisp style
- [x] 5.4 Update error card background to #ffffff
- [x] 5.5 Update button border-radius to 2px for formal appearance

## 6. PhotoCapture Component Updates

- [x] 6.1 Update capture button border-radius to 2px
- [x] 6.2 Update photo preview container border-radius to 2px
- [x] 6.3 Update confirm/retake button border-radius to 2px
- [x] 6.4 Ensure consistent crisp shadows on interactive elements

## 7. CameraPermission Component Updates

- [x] 7.1 Update permission card border-radius from 8px to 2px
- [x] 7.2 Add 3px solid #4a4a4a top border to permission card
- [x] 7.3 Update permission card box-shadow to crisp style
- [x] 7.4 Update button border-radius to 2px

## 8. UploadProgress Component Updates

- [x] 8.1 Update progress card border-radius from 8px to 2px
- [x] 8.2 Add 3px solid #4a4a4a top border to progress card
- [x] 8.3 Update progress card box-shadow to crisp style
- [x] 8.4 Ensure progress bar maintains government color scheme

## 9. Testing and Verification

- [x] 9.1 Verify no gradient backgrounds remain anywhere in application
- [x] 9.2 Test court seal loads correctly and displays at proper size
- [x] 9.3 Test case number generation produces correct format (RVM-YYYY-timestamp)
- [x] 9.4 Test Dutch date formatting displays correctly
- [x] 9.5 Test responsive design on mobile (seal size reduction, layout integrity)
- [x] 9.6 Verify all cards have 2px border-radius and 3px top borders
- [x] 9.7 Check color contrast ratios meet WCAG AA standards
- [x] 9.8 Test on Chrome, Firefox, and Safari
- [x] 9.9 Visual regression check: compare before/after screenshots
- [x] 9.10 Verify horizontal dividers display correctly between sections
