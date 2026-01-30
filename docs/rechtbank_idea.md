1. The Architecture (Docker Compose)

You can run this using three lightweight containers. Since Go compiles into a small binary, your footprint on the VM remains minimal.

    Frontend (Nginx/SPA): A simple React or Svelte app where the user opens the camera or uploads a photo.

    Backend (Go): The API that receives the photo, forwards it to the analysis service, and returns the "verdict."

    Reverse Proxy (Traefik or Nginx): For your SSL (HTTPS is required to use the camera on a smartphone).

2. Photo Analysis: The "Smart" Choice

Previously, I would have recommended OpenCV (computer vision), but that requires a lot of mathematical work to recognize perspective and objects.

My advice: Use the Gemini 2.5 Flash Lite API.

    Why: It is a "multimodal" model. You send the photo + a text prompt ("Is this couch straight?"). The model immediately understands what a couch is, sees the angle of the backrest, and can respond immediately in legal jargon.

    Costs: Extremely low. There is a free tier, and above that in 2026, you only pay about $0.10 per 1 million tokens (a few cents for thousands of photos).

System Instruction Prompt:

    "Je bent een strenge rechter van de Rechtbank voor Meubilair. Analyseer de geüploade foto. Als het geen bank is, verklaar de zaak dan 'niet-ontvankelijk'. Als het wel een bank is, beoordeel of deze 180 graden recht is. Geef een score van 1-10 en een grappig juridisch vonnis (bijv. 'Veroordeeld tot de brandstapel wegens een rugleuning-afwijking van 5 graden')."

3. The Go Backend (The "Clerk")

In Go, you can use the net/http library or a framework like Gin or Echo. The workflow is simple:

    Endpoint /v1/judge: Receives the multipart/form-data (the image).

    API Call: Go sends the image bytes to the Google Gemini API (via the official google-generative-ai-go SDK).

    JSON Response: Gemini returns a funny text, which Go sends back to your frontend.

5. The Interactive Element (The "Live Level")

To really "finish" the project, you can place a CSS overlay over the camera feed in the frontend.

    Use the browser's DeviceOrientationEvent API.

    While the user attempts to take the photo, they see a virtual spirit level (waterpas) on the screen.

    The "Upload" button only becomes active when the phone is held perfectly straight. This forces the user to fully engage with the absurdity of the "trial."