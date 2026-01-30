// Verdict response from the backend API
export interface Verdict {
    /** Type of verdict rendered by the court */
    type: 'niet-ontvankelijk' | 'guilty' | 'acquittal';
    /** Straightness score from 1-10 (10 being perfectly straight) */
    score: number;
    /** Full verdict text in Dutch legal language */
    verdictText: string;
    /** Sentence or acknowledgment based on verdict type */
    sentence?: string;
    /** Angle deviation in degrees (if applicable) */
    angleDeviation?: number;
    /** Whether the submitted item was identified as furniture */
    isFurniture: boolean;
}
