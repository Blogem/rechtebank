1. De Architectuur (Docker Compose)

Je kunt dit draaien met drie lichte containers. Omdat Go gecompileerd wordt naar een kleine binary, blijft je footprint op de VM minimaal.

    Frontend (Nginx/SPA): Een simpele React of Svelte app waar de gebruiker de camera opent of een foto uploadt.

    Backend (Go): De API die de foto ontvangt, doorstuurt naar de analyse-service en het "vonnis" teruggeeft.

    Reverse Proxy (Traefik of Nginx): Voor je SSL (HTTPS is nodig om de camera op een smartphone te mogen gebruiken).

2. De Foto-analyse: De "Slimme" Keuze

Voorheen zou ik OpenCV (computer vision) aanraden, maar dat is veel wiskundig werk om perspectief en objecten te herkennen.

Mijn advies: Gebruik de Gemini 1.5 Flash API. => is er niet meer, gebruik gemini-2.5-flash-lite

    Waarom: Het is een "multimodal" model. Je stuurt de foto + een tekstprompt op ("Is deze bank recht?"). Het model begrijpt direct wat een bank is, ziet de hoek van de rugleuning en kan direct in juridisch jargon antwoorden.

    Kosten: Extreem laag. Er is een gratis tier (tot 15 verzoeken per minuut), en daarboven betaal je in 2026 slechts ongeveer $0,50 per 1 miljoen tokens (een paar cent voor duizenden foto's).

Hoe de prompt eruit ziet (System Instruction):

    "Je bent een strenge rechter van de Rechtbank voor Meubilair. Analyseer de geüploade foto. Als het geen bank is, verklaar de zaak dan 'niet-ontvankelijk'. Als het wel een bank is, beoordeel of deze 180 graden recht is. Geef een score van 1-10 en een grappig juridisch vonnis (bijv. 'Veroordeeld tot de brandstapel wegens een rugleuning-afwijking van 5 graden')."

3. De Go Backend (De "Griffier")

In Go kun je de net/http library gebruiken of een framework zoals Gin of Echo. De workflow is simpel:

    Endpoint /v1/judge: Ontvangt de multipart/form-data (de afbeelding).

    API Call: Go stuurt de image bytes naar Google Gemini API (via de officiële google-generative-ai-go SDK).

    JSON Response: Gemini geeft een grappige tekst terug, die Go weer naar je frontend stuurt.

5. Het Interactieve Element (De "Live Waterpas")

Om het echt "af" te maken, kun je in de frontend een CSS-overlay over de camerafeed leggen.

    Gebruik de DeviceOrientationEvent API van de browser.

    Terwijl de gebruiker de foto probeert te maken, ziet hij een virtuele waterpas in beeld.

    Pas als de telefoon kaarsrecht wordt gehouden, wordt de "Upload" knop actief. Dit dwingt de gebruiker om serieus mee te doen met de absurditeit.