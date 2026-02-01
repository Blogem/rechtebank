// Request for creating a shareable verdict URL
export interface ShareVerdictRequest {
    /** ISO 8601 timestamp of the verdict */
    timestamp: string;
    /** Unique request identifier */
    requestId: string;
}

// Response from share verdict endpoint
export interface ShareVerdictResponse {
    /** Base64url-encoded verdict ID */
    id: string;
}

// Response from verdict retrieval endpoint (GET /v1/verdict/:id)
export interface VerdictWithImageResponse {
    /** The verdict data */
    verdict: import('./Verdict').Verdict;
    /** Base64-encoded image data URL (e.g., "data:image/jpeg;base64,...") */
    image: string;
}
