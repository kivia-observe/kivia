import { KiviaClientOptions, Log } from './types';

export class KiviaClient {
  private apiKey: string;
  private baseUrl: string;

  constructor(options: KiviaClientOptions) {
    this.apiKey = options.apiKey;
    this.baseUrl = options.baseUrl || 'https://nginx-production-a9aa.up.railway.app/api/v1';
  }

  /**
   * Express/Connect middleware for automatically logging requests.
   */
  public logMiddleware() {
    return (req: any, res: any, next: any) => {
      const start = Date.now();

      // Listen for the response to finish
      res.on('finish', () => {
        const latency = Date.now() - start;

        const logEntry: Log = {
          path: req.originalUrl || req.url,
          status: res.statusCode,
          ip_address: req.ip || req.connection?.remoteAddress,
          timestamp: new Date().toISOString(),
          latency,
        };

        this.sendLog(logEntry);
      });

      next();
    };
  }

  /**
   * Fastify hook variant to be used with the 'onResponse' lifecycle hook.
   * Usage: fastify.addHook('onResponse', kiviaClient.logFastifyOnResponse());
   */
  public logFastifyOnResponse() {
    return async (request: any, reply: any) => {
      // reply.getResponseTime() is available in Fastify to get latency in ms
      const latency = reply.getResponseTime ? reply.getResponseTime() : 0;

      const logEntry: Log = {
        path: request.url || request.raw?.url,
        status: reply.statusCode || reply.raw?.statusCode,
        ip_address: request.ip || request.raw?.connection?.remoteAddress,
        timestamp: new Date().toISOString(),
        latency: Math.round(latency),
      };

      this.sendLog(logEntry);
    };
  }

  /**
   * Hono middleware for automatically logging requests.
   *
   * Usage with Hono:
   *   app.use('*', kiviaClient.logHono());
   */
  public logHono() {
    return async (c: any, next: any) => {
      const start = Date.now();

      await next();

      const latency = Date.now() - start;

      const logEntry: Log = {
        path: c.req.path || c.req.url,
        status: c.res.status,
        ip_address:
          c.req.header?.('x-forwarded-for')?.split(',')[0] ||
          c.req.header?.('x-real-ip') ||
          c.req.raw?.headers?.['x-forwarded-for']?.split(',')[0] ||
          c.req.raw?.socket?.remoteAddress,
        timestamp: new Date().toISOString(),
        latency,
      };

      this.sendLog(logEntry);
    };
  }

  /**
   * Elysia plugin for automatically logging requests.
   *
   * Usage with Elysia:
   *   import { Elysia } from 'elysia';
   *   import { kiviaElysiaPlugin } from '@kivia/sdk';
   *
   *   const app = new Elysia().use(kiviaElysiaPlugin({ apiKey: 'YOUR_KIVIA_API_KEY' }));
   */
  public elysiaPlugin() {
    const client = this;

    return (app: any) =>
      app
        // Capture start time per request
        .derive(() => ({ kiviaStart: Date.now() }))
        // Log after the response has been sent
        .onAfterResponse(({ request, set, kiviaStart, ip }: any) => {
          const end = Date.now();
          const latency = typeof kiviaStart === 'number' ? end - kiviaStart : 0;

          let path = '/';
          try {
            if (request?.url) {
              const url = new URL(request.url);
              path = url.pathname || path;
            }
          } catch {
            if (request?.url) path = request.url;
          }

          const logEntry: Log = {
            path,
            status: typeof set?.status === 'number' ? set.status : 200,
            ip_address: ip,
            timestamp: new Date().toISOString(),
            latency,
          };

          client.sendLog(logEntry);
        });
  }

  private async sendLog(logEntry: Log): Promise<void> {
    try {
      const response = await fetch(`${this.baseUrl}/logs/create`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          'X-kivia-api-key': this.apiKey,
        },
        body: JSON.stringify(logEntry),
      });

      if (!response.ok) {
        const errBody = await response.text();
        console.error(`kiviasdk: log rejected (${response.status}): ${errBody}`);
      }
    } catch (error) {
      console.error('kiviasdk: failed to send log:', error);
    }
  }
}
