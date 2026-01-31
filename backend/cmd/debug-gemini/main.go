package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"path/filepath"
	"time"

	"rechtebank/backend/internal/adapters/gemini"
	"rechtebank/backend/internal/core/domain"

	_ "golang.org/x/image/webp"
)

const (
	exitSuccess = 0
	exitError   = 1
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(exitError)
	}
	os.Exit(exitSuccess)
}

func run() error {
	// Parse command-line arguments
	if len(os.Args) != 2 {
		return fmt.Errorf("usage: %s <image-path>", filepath.Base(os.Args[0]))
	}

	imagePath := os.Args[1]

	// Validate file exists
	if _, err := os.Stat(imagePath); os.IsNotExist(err) {
		return fmt.Errorf("image file not found: %s", imagePath)
	}

	// Load image file
	imageData, err := os.ReadFile(imagePath)
	if err != nil {
		return fmt.Errorf("failed to read image file: %w", err)
	}

	// Validate image format
	mimeType := detectMIMEType(imageData)
	if mimeType == "" {
		return fmt.Errorf("unsupported image format (must be JPEG, PNG, or WebP)")
	}

	// Get image dimensions
	dimensions, err := getImageDimensions(imageData)
	if err != nil {
		return fmt.Errorf("failed to get image dimensions: %w", err)
	}

	// Get file size
	fileSize := formatFileSize(len(imageData))

	// Load API key
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		return fmt.Errorf("GEMINI_API_KEY environment variable is required")
	}

	// Print header
	printSection("GEMINI DEBUG TOOL")
	fmt.Printf("Image: %s\n", imagePath)
	fmt.Printf("Size: %s\n", fileSize)
	fmt.Printf("MIME Type: %s\n", mimeType)
	fmt.Printf("Dimensions: %s\n", dimensions)
	fmt.Println()

	// Print system prompt
	printSection("SYSTEM PROMPT")
	fmt.Println(gemini.GetSystemPrompt())
	fmt.Println()

	// Print user prompt
	printSection("USER PROMPT")
	fmt.Println(gemini.GetUserPrompt())
	fmt.Println()

	// Print request metadata
	timeout := 30 * time.Second
	maxRetries := 3
	printSection("REQUEST METADATA")
	fmt.Printf("Model: gemini-2.5-flash-lite\n")
	fmt.Printf("Timeout: %v\n", timeout)
	fmt.Printf("Max Retries: %d\n", maxRetries)
	fmt.Println()

	// Initialize analyzer
	analyzer, err := gemini.NewGeminiAnalyzer(apiKey, timeout)
	if err != nil {
		return fmt.Errorf("failed to initialize analyzer: %w", err)
	}
	defer analyzer.Close()

	// Call API
	ctx := context.Background()
	printSection("CALLING GEMINI API")
	fmt.Println("Sending request...")
	fmt.Println()

	// For debug purposes, we need to capture the raw response
	// We'll create a debug version that exposes internals
	response, rawJSON, err := analyzeWithDebug(ctx, analyzer, imageData, timeout)

	if err != nil {
		printSection("ERROR")
		fmt.Printf("API call failed: %v\n", err)
		fmt.Println()

		// Try to show any available context
		if rawJSON != "" {
			printSection("RAW RESPONSE (ERROR)")
			fmt.Println(rawJSON)
			fmt.Println()
		}

		return err
	}

	// Print raw response
	printSection("RAW RESPONSE")
	fmt.Println(rawJSON)
	fmt.Println()

	// Print parsed verdict
	printSection("PARSED VERDICT")
	fmt.Printf("Observation: %s\n", response.Verdict.Observation)
	fmt.Printf("Admissible: %v\n", response.Admissible)
	fmt.Printf("Score: %d/10\n", response.Score)
	fmt.Printf("Crime: %s\n", response.Verdict.Crime)
	fmt.Printf("Sentence: %s\n", response.Verdict.Sentence)
	fmt.Printf("Reasoning: %s\n", response.Verdict.Reasoning)
	fmt.Println()

	return nil
}

func printSection(title string) {
	fmt.Printf("=== %s ===\n", title)
}

func detectMIMEType(data []byte) string {
	if len(data) < 12 {
		return ""
	}

	// JPEG: FF D8 FF
	if data[0] == 0xFF && data[1] == 0xD8 && data[2] == 0xFF {
		return "jpeg"
	}

	// PNG: 89 50 4E 47 0D 0A 1A 0A
	if data[0] == 0x89 && data[1] == 0x50 && data[2] == 0x4E && data[3] == 0x47 {
		return "png"
	}

	// WebP: RIFF....WEBP
	if data[0] == 0x52 && data[1] == 0x49 && data[2] == 0x46 && data[3] == 0x46 &&
		data[8] == 0x57 && data[9] == 0x45 && data[10] == 0x42 && data[11] == 0x50 {
		return "webp"
	}

	return ""
}

func getImageDimensions(imageData []byte) (string, error) {
	img, _, err := image.DecodeConfig(bytes.NewReader(imageData))
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%dx%d", img.Width, img.Height), nil
}

func formatFileSize(bytes int) string {
	const kb = 1024
	const mb = kb * 1024

	if bytes < kb {
		return fmt.Sprintf("%d bytes", bytes)
	} else if bytes < mb {
		return fmt.Sprintf("%.1f KB", float64(bytes)/float64(kb))
	}
	return fmt.Sprintf("%.1f MB", float64(bytes)/float64(mb))
}

// analyzeWithDebug calls the analyzer and returns the response with raw JSON
func analyzeWithDebug(ctx context.Context, analyzer *gemini.GeminiAnalyzer, imageData []byte, timeout time.Duration) (*domain.VerdictResponse, string, error) {
	response, err := analyzer.AnalyzePhoto(ctx, imageData)
	if err != nil {
		return nil, "", err
	}

	// Pretty-print the raw JSON for readability
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(response.RawJSON), "", "  "); err != nil {
		// If pretty-printing fails, just use the raw JSON as-is
		return response, response.RawJSON, nil
	}

	return response, prettyJSON.String(), nil
}
