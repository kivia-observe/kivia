export interface Log {
  id?: string;
  path: string;
  status: number;
  ip_address?: string;
  timestamp: string; // ISO 8601
  latency: number; // milliseconds
}

export interface KiviaClientOptions {
  apiKey: string;
  /**
   * Override the default production base URL.
   */
  baseUrl?: string;
}
