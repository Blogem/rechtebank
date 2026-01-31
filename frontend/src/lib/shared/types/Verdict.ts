// Verdict response from the backend API
// Matches Go's VerdictResponse structure
export interface Verdict {
    /** Whether the submitted item was identified as furniture */
    admissible: boolean;
    /** Straightness score from 1-10 (10 being perfectly straight) */
    score: number;
    /** Detailed verdict components */
    verdict: VerdictDetails;
    /** Unique request identifier */
    requestId: string;
    /** ISO 8601 timestamp of the verdict */
    timestamp: string;
}

// Detailed verdict components
// Matches Go's VerdictDetails structure
export interface VerdictDetails {
    /** The furniture offense (e.g., "Scheefhangende Zitting") */
    crime: string;
    /** The punishment or acknowledgment */
    sentence: string;
    /** Legal justification with fictitious law articles */
    reasoning: string;
    /** What the judge observed in the photo */
    observation: string;
    /** The verdict classification */
    verdictType: "vrijspraak" | "waarschuwing" | "schuldig";
}
