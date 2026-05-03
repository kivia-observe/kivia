import { KiviaClient } from './client';
import { KiviaClientOptions } from './types';

/**
 * Fastify plugin for Kivia observability.
 * Usage: fastify.register(kiviaFastifyPlugin, { apiKey: '...' })
 */
export async function kiviaFastifyPlugin(fastify: any, options: KiviaClientOptions) {
  const client = new KiviaClient(options);
  fastify.addHook('onResponse', client.logFastifyOnResponse());
}
