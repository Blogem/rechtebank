package gemini

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"rechtebank/backend/internal/core/domain"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

// const systemPrompt = `Je bent de Eerwaarde Rechter van de Meubilair-rechtbank, een humoristisch gerechtshof dat oordeelt over de rechtheid en uitlijning van meubels.

// BELANGRIJK: Wees GENEREUS in wat je als meubilair accepteert. Elk object dat voor zitten, liggen, opbergen of versieren dient (bank, stoel, kast, tafel, bed, nachtkastje, plank, enz.) is meubilair.

// Analyseer de foto en bepaal:
// 1. Of het object meubilair is (admissible: true/false) - wees ruimhartig hierin!
// 2. Hoe recht/uitgelijnd het meubilair staat (score: 1-10, waarbij 10 perfect recht is)
// 3. Een vonnis in formele Nederlandse juridische stijl

// Als er ABSOLUUT GEEN meubilair of meubelachtig object zichtbaar is (bijv. een persoon, dier, voedsel):
// - admissible: false
// - score: 0
// - crime: "Geen meubilair gedetecteerd"
// - sentence: "Zaak niet-ontvankelijk verklaard"
// - reasoning: "Dit Hof oordeelt alleen over meubilair en meubelachtige objecten"

// Als er WEL meubilair (of iets meubelachtigs) zichtbaar is, geef dan ALTIJD:
// - admissible: true
// - Een passende score (1-10) gebaseerd op de rechtheid
// - Een creatieve "misdaad" beschrijving (bijv. "Scheve zitting van 7 graden", "Horizontale ongehoorzaamheid")
// - Een humoristisch vonnis (bijv. "Veroordeeld tot heroriëntatie", "Vrijgesproken wegens voorbeeldige uitlijning")
// - Juridische redenering met verwijzing naar fictieve wetsartikelen (bijv. "Artikel 42 van de Meubilair-wet verbiedt afwijkingen van meer dan 3 graden...")

// Scoreverdeling:
// - 9-10: Uitzonderlijk recht, hoogste lof en vrijspraak
// - 7-8: Goed recht, complimenten
// - 5-6: Lichte afwijking, berisping
// - 3-4: Matige scheefstand, strenge waarschuwing
// - 1-2: Ernstige scheefstand, zware veroordeling

// Wees creatief, humoristisch en overdreven formeel in je juridische taalgebruik!`

const systemPrompt = `Je bent de Eerwaarde Rechter van de Meubilair-rechtbank (Rechtbank.org). Je bent een absurdistische, hyper-formele magistraat die gespecialiseerd is in de 'Wet op de Verticale Integriteit'.

GEBRUIK DE VOLGENDE RICHTLIJNEN:
1. IDENTIFICATIE: 
   - Beschouw elk object met een structuur (poten, vlakken, leuningen) als meubilair. 
   - Zelfs een omgevallen stoel of een rommelige tafel is "bewijsmateriaal".
   - Alleen als er absoluut geen fysiek object herkenbaar is (bijv. alleen een zwart scherm of een selfie), verklaar je de zaak niet-ontvankelijk.

2. ANALYSE (Stap-voor-stap):
   - Stap 1: Benoem het object (bijv. "Een houten zetel met groene ribstof").
   - Stap 2: Meet de hoek ten opzichte van de horizon. 
   - Stap 3: Zoek naar 'strafbare feiten' zoals scheve poten, een doorgezakte zitting of 'ongeoorloofde hellingshoeken'.

3. JURIDISCHE STIJL:
   - Gebruik termen als: "In naam der Koning der Meubelen", "Overwegende dat", "Het Hof gelast", "Wetsartikel 3.14 van het Wetboek van Stoelgang".
   - Wees streng maar rechtvaardig. Een score van 10/10 is zeldzaam; er is altijd wel een splinter die niet deugt.

4. UITSPRAAK:
- observation: Wat de rechter ziet
- admissible: true/false
- score: 1-10
- crime: De juridische naam van de afwijking
- reasoning: Juridische onderbouwing met fictieve wetsartikelen
- sentence: De humoristische straf of de volledige vrijspraak

WEES CREATIEF, HUMORISTISCH EN OVERDREVEN FORMEEL IN JE JURIDISCHE TAALGEBRUIK!`

// VerdictSchema defines the JSON schema for Gemini structured output
type VerdictSchema struct {
	Observation string `json:"observation"`
	Admissible  bool   `json:"admissible"`
	Score       int    `json:"score"`
	Crime       string `json:"crime"`
	Sentence    string `json:"sentence"`
	Reasoning   string `json:"reasoning"`
}

// RealGeminiClient wraps the actual Gemini API client
type RealGeminiClient struct {
	client *genai.Client
	model  *genai.GenerativeModel
}

// NewRealGeminiClient creates a new client connected to the Gemini API
func NewRealGeminiClient(ctx context.Context, apiKey string) (*RealGeminiClient, error) {
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, fmt.Errorf("failed to create Gemini client: %w", err)
	}

	model := client.GenerativeModel("gemini-2.5-flash-lite")
	model.SystemInstruction = genai.NewUserContent(genai.Text(systemPrompt))

	// Configure JSON schema for structured output
	model.ResponseMIMEType = "application/json"
	model.ResponseSchema = &genai.Schema{
		Type: genai.TypeObject,
		Properties: map[string]*genai.Schema{
			"admissible":  {Type: genai.TypeBoolean},
			"score":       {Type: genai.TypeInteger},
			"crime":       {Type: genai.TypeString},
			"sentence":    {Type: genai.TypeString},
			"reasoning":   {Type: genai.TypeString},
			"observation": {Type: genai.TypeString},
		},
		Required: []string{"admissible", "score", "crime", "sentence", "reasoning", "observation"},
	}

	return &RealGeminiClient{
		client: client,
		model:  model,
	}, nil
}

// GenerateContent sends an image to Gemini and returns the verdict
func (c *RealGeminiClient) GenerateContent(ctx context.Context, imageData []byte) (*GeminiResponse, error) {
	// Detect MIME type from magic bytes
	mimeType := detectMIMEType(imageData)
	if mimeType == "" {
		return nil, errors.New("unsupported image format")
	}

	resp, err := c.model.GenerateContent(ctx,
		genai.ImageData(mimeType, imageData),
		genai.Text("Analyseer dit meubelstuk en spreek je vonnis uit."),
	)
	if err != nil {
		return nil, err
	}

	if len(resp.Candidates) == 0 || len(resp.Candidates[0].Content.Parts) == 0 {
		return nil, &InvalidResponseError{Message: "empty response from Gemini"}
	}

	// Extract text from response
	textPart, ok := resp.Candidates[0].Content.Parts[0].(genai.Text)
	if !ok {
		return nil, &InvalidResponseError{Message: "unexpected response format"}
	}

	// Parse JSON response
	var schema VerdictSchema
	if err := json.Unmarshal([]byte(textPart), &schema); err != nil {
		return nil, &InvalidResponseError{Message: fmt.Sprintf("failed to parse response: %v", err)}
	}

	return &GeminiResponse{
		Admissible:  schema.Admissible,
		Score:       schema.Score,
		Crime:       schema.Crime,
		Sentence:    schema.Sentence,
		Reasoning:   schema.Reasoning,
		Observation: schema.Observation,
	}, nil
}

// Close closes the Gemini client
func (c *RealGeminiClient) Close() error {
	return c.client.Close()
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

// NewGeminiAnalyzer creates a new GeminiAnalyzer with the given API key
func NewGeminiAnalyzer(apiKey string, timeout time.Duration) (*GeminiAnalyzer, error) {
	if apiKey == "" {
		return nil, errors.New("GEMINI_API_KEY environment variable is required")
	}

	ctx := context.Background()
	client, err := NewRealGeminiClient(ctx, apiKey)
	if err != nil {
		return nil, err
	}

	return &GeminiAnalyzer{
		realClient: client,
		timeout:    timeout,
		maxRetries: 3,
	}, nil
}

// GeminiAnalyzer implements IPhotoAnalyzer using Google Gemini API
type GeminiAnalyzer struct {
	client     GeminiClientInterface // For testing with mocks
	realClient *RealGeminiClient     // For real API calls
	timeout    time.Duration
	maxRetries int
}

// GeminiResponse represents the parsed response from Gemini API
type GeminiResponse struct {
	Admissible  bool
	Score       int
	Crime       string
	Sentence    string
	Reasoning   string
	Observation string
}

// GeminiClientInterface defines the interface for the Gemini client
type GeminiClientInterface interface {
	GenerateContent(ctx context.Context, imageData []byte) (*GeminiResponse, error)
}

// RateLimitError indicates a rate limit was hit
type RateLimitError struct {
	RetryAfter time.Duration
}

func (e *RateLimitError) Error() string {
	return "rate limit exceeded"
}

// InvalidResponseError indicates an invalid response from the API
type InvalidResponseError struct {
	Message string
}

func (e *InvalidResponseError) Error() string {
	return e.Message
}

// AnalyzePhoto analyzes the image using Gemini API
func (a *GeminiAnalyzer) AnalyzePhoto(ctx context.Context, imageData []byte) (*domain.VerdictResponse, error) {
	var lastErr error
	attempts := a.maxRetries + 1
	if attempts < 1 {
		attempts = 1
	}

	// Use mock client if set (for testing), otherwise use real client
	client := a.getClient()

	for i := 0; i < attempts; i++ {
		ctxWithTimeout, cancel := context.WithTimeout(ctx, a.timeout)
		response, err := client.GenerateContent(ctxWithTimeout, imageData)
		cancel()

		if err == nil {
			return &domain.VerdictResponse{
				Admissible: response.Admissible,
				Score:      response.Score,
				Verdict: domain.VerdictDetails{
					Crime:     response.Crime,
					Sentence:  response.Sentence,
					Reasoning: response.Reasoning,
				},
			}, nil
		}

		lastErr = err

		// Check for specific error types
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, errors.New("AI analysis timeout")
		}

		var rateLimitErr *RateLimitError
		if errors.As(err, &rateLimitErr) {
			if i < a.maxRetries {
				// Exponential backoff
				backoff := time.Duration(1<<uint(i)) * time.Second
				if backoff > 8*time.Second {
					backoff = 8 * time.Second
				}
				time.Sleep(rateLimitErr.RetryAfter)
				continue
			}
			return nil, errors.New("AI analysis service temporarily unavailable")
		}

		var invalidErr *InvalidResponseError
		if errors.As(err, &invalidErr) {
			return nil, errors.New("Invalid AI response format")
		}

		return nil, fmt.Errorf("AI analysis failed: %w", err)
	}

	return nil, lastErr
}

func (a *GeminiAnalyzer) getClient() GeminiClientInterface {
	if a.client != nil {
		return a.client
	}
	return a.realClient
}

// Close closes the analyzer and its resources
func (a *GeminiAnalyzer) Close() error {
	if a.realClient != nil {
		return a.realClient.Close()
	}
	return nil
}
