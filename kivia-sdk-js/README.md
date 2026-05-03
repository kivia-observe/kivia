# Kivia TypeScript SDK

The Kivia TypeScript SDK is the official Node.js client for the Kivia observability platform. It allows Node.js developers using Express, Fastify, Hono, or Elysia to instantly track API request metrics, response times, and paths.

## Installation

```bash
npm install @kivia/sdk
```

## Quick Start (with Express)

```typescript
import express from 'express';
import { KiviaClient } from '@kivia/sdk';

const app = express();

const kiviaClient = new KiviaClient({
  apiKey: 'YOUR_KIVIA_API_KEY',
});

// Let Kivia track all your network traffic by setting this global middleware
app.use(kiviaClient.logMiddleware());

app.get('/hello', (req, res) => {
  res.send('Hello from Kivia TS SDK!');
});

app.listen(3000, () => console.log('Server running on port 3000'));
```

## Quick Start (with NestJS)

Because NestJS runs on Express (or Fastify) under the hood, you can simply inject the SDK as a global middleware right in your `main.ts` bootstrap function!

```typescript
// main.ts
import { NestFactory } from '@nestjs/core';
import { AppModule } from './app.module';
import { KiviaClient } from '@kivia/sdk';

async function bootstrap() {
  const app = await NestFactory.create(AppModule);

  const kiviaClient = new KiviaClient({
    apiKey: 'YOUR_KIVIA_API_KEY'
  });

  // Just apply it as global middleware!
  app.use(kiviaClient.logMiddleware());

  await app.listen(3000);
}
bootstrap();
```

## Quick Start (with Fastify Plugin)

We export a native Fastify plugin so you can easily encapsulate options and registration.

```typescript
import Fastify from 'fastify';
import { kiviaFastifyPlugin } from '@kivia/sdk';

const fastify = Fastify({ logger: false });

// Register as a native Fastify plugin
fastify.register(kiviaFastifyPlugin, {
  apiKey: 'YOUR_KIVIA_API_KEY'
});

fastify.get('/hello', async (request, reply) => {
  return { message: 'Hello from Kivia TS SDK using Fastify Plugin!' };
});

fastify.listen({ port: 3000 });
```

## Quick Start (with Hono)

```typescript
import { Hono } from 'hono';
import { kiviaHonoMiddleware } from '@kivia/sdk';

const app = new Hono();

app.use('*', kiviaHonoMiddleware({
  apiKey: 'YOUR_KIVIA_API_KEY',
}));

app.get('/hello', (c) => {
  return c.json({ message: 'Hello from Kivia TS SDK using Hono!' });
});

export default app;
```

## Quick Start (with Elysia)

```typescript
import { Elysia } from 'elysia';
import { kiviaElysiaPlugin } from '@kivia/sdk';

const app = new Elysia()
  .use(
    kiviaElysiaPlugin({
      apiKey: 'YOUR_KIVIA_API_KEY',
    }),
  );

app.get('/hello', () => ({ message: 'Hello from Kivia TS SDK using Elysia!' }));

app.listen(3000);
```
