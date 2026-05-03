import { KiviaClient } from './client';
import { KiviaClientOptions } from './types';

/**
 * Hono middleware factory for Kivia observability.
 *
 * Usage:
 *   import { Hono } from 'hono';
 *   import { kiviaHonoMiddleware } from '@kivia/sdk';
 *
 *   const app = new Hono();
 *   app.use('*', kiviaHonoMiddleware({ apiKey: 'YOUR_KIVIA_API_KEY' }));
 */
export function kiviaHonoMiddleware(options: KiviaClientOptions) {
  const client = new KiviaClient(options);
  return client.logHono();
}
