import { describe, it, expect, vi, beforeEach } from 'vitest';
import { ApiAdapter } from './ApiAdapter';
import type { PhotoMetadata } from '../ports/IApiPort';

describe('ApiAdapter', () => {
    let adapter: ApiAdapter;
    let mockFetch: ReturnType<typeof vi.fn>;

    beforeEach(() => {
        adapter = new ApiAdapter('http://localhost:8080', 3, 100);
        mockFetch = vi.fn();
        globalThis.fetch = mockFetch as any;
    });

    const createMockMetadata = (): PhotoMetadata => ({
        userAgent: 'test-agent',
        timestamp: new Date().toISOString(),
        captureMethod: 'camera'
    });

    describe('uploadPhoto', () => {
        it('should upload photo as multipart/form-data', async () => {
            const mockBlob = new Blob(['test'], { type: 'image/jpeg' });
            const mockVerdict = {
                admissible: true,
                score: 5,
                verdict: {
                    crime: 'Scheefhangende Zitting',
                    sentence: 'Test sentence',
                    reasoning: 'Test reasoning',
                    observation: 'Test observation'
                },
                requestId: 'test-id',
                timestamp: new Date().toISOString()
            };

            mockFetch.mockResolvedValue({
                ok: true,
                json: async () => mockVerdict
            });

            const verdict = await adapter.uploadPhoto(mockBlob, createMockMetadata());

            expect(mockFetch).toHaveBeenCalledWith(
                'http://localhost:8080/v1/judge',
                expect.objectContaining({
                    method: 'POST'
                })
            );
            expect(verdict).toEqual(mockVerdict);
        });

        it('should throw error for files larger than 10MB', async () => {
            const largeBlob = new Blob([new ArrayBuffer(11 * 1024 * 1024)], { type: 'image/jpeg' });

            await expect(adapter.uploadPhoto(largeBlob, createMockMetadata())).rejects.toThrow(
                'te groot'
            );
        });

        it('should retry on network errors', async () => {
            const mockBlob = new Blob(['test'], { type: 'image/jpeg' });
            const mockVerdict = {
                type: 'guilty',
                score: 5,
                verdictText: 'Test verdict',
                isFurniture: true
            };

            // Fail twice, then succeed
            mockFetch
                .mockRejectedValueOnce(new TypeError('Failed to fetch'))
                .mockRejectedValueOnce(new TypeError('Failed to fetch'))
                .mockResolvedValue({
                    ok: true,
                    json: async () => mockVerdict
                });

            const verdict = await adapter.uploadPhoto(mockBlob, createMockMetadata());

            expect(mockFetch).toHaveBeenCalledTimes(3);
            expect(verdict).toEqual(mockVerdict);
        });

        it('should throw error after max retries', async () => {
            const mockBlob = new Blob(['test'], { type: 'image/jpeg' });

            mockFetch.mockRejectedValue(new TypeError('Failed to fetch'));

            await expect(adapter.uploadPhoto(mockBlob, createMockMetadata())).rejects.toThrow();
            expect(mockFetch).toHaveBeenCalledTimes(3); // Max retries
        });

        it('should timeout after 30 seconds', async () => {
            vi.useFakeTimers();
            const mockBlob = new Blob(['test'], { type: 'image/jpeg' });

            // Mock fetch that respects abort signal
            mockFetch.mockImplementation(
                (url, options) =>
                    new Promise((resolve, reject) => {
                        const abortHandler = () => {
                            reject(new DOMException('The operation was aborted', 'AbortError'));
                        };
                        options.signal?.addEventListener('abort', abortHandler);
                    })
            );

            const uploadPromise = adapter.uploadPhoto(mockBlob, createMockMetadata());

            // Add catch handler immediately to handle async rejections
            uploadPromise.catch(() => { });

            // Fast-forward time by 30 seconds to trigger timeout
            vi.advanceTimersByTime(30000);

            // Wait a tick for the promise to reject
            await vi.runAllTimersAsync();

            await expect(uploadPromise).rejects.toThrow('te lang beraadslaagd');

            vi.useRealTimers();
        });
    });
});
