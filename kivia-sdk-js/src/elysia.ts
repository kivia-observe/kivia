import { KiviaClient } from './client';
import { KiviaClientOptions } from './types';

/**
 * Elysia plugin for Kivia observability.
 *
 * Usage:
 *   import { Elysia } from 'elysia';
 *   import { kiviaElysiaPlugin } from '@kivia/sdk';
 *
 *   const app = new Elysia()
 *     .use(kiviaElysiaPlugin({ apiKey: 'YOUR_KIVIA_API_KEY' }));
 */
export function kiviaElysiaPlugin(options: KiviaClientOptions) {
  const client = new KiviaClient(options);
  return client.elysiaPlugin();
}
