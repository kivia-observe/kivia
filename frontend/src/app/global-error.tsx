"use client";

import { RotateCcw, Home, Zap } from "lucide-react";
import { Button } from "@/components/ui/button";
import { Separator } from "@/components/ui/separator";
import "./globals.css";

export default function GlobalError({
  error,
  reset,
}: {
  error: Error & { digest?: string };
  reset: () => void;
}) {
  return (
    <html lang="en" className="dark">
      <body className="antialiased bg-background text-foreground">
        <div className="flex min-h-screen flex-col">
          {/* Header */}
          <header className="flex items-center gap-2.5 px-6 h-16 shrink-0 border-b">
            <a href="/" className="flex items-center gap-2.5">
              <div className="flex h-8 w-8 items-center justify-center rounded-lg bg-primary text-primary-foreground shadow-sm shadow-primary/20">
                <Zap className="h-4 w-4" />
              </div>
              <span className="text-base font-bold tracking-tight">Kivia</span>
            </a>
          </header>

          {/* Content */}
          <div className="flex flex-1 flex-col items-center justify-center px-6">
            <div className="relative mb-8">
              <p className="text-[10rem] leading-none font-black tracking-tighter text-destructive/10 select-none">
                500
              </p>
              <p className="absolute inset-0 flex items-center justify-center text-6xl font-bold tracking-tighter text-destructive">
                500
              </p>
            </div>

            <h1 className="text-2xl font-bold tracking-tight text-foreground">
              Something went wrong
            </h1>
            <p className="mt-3 text-base text-muted-foreground max-w-md text-center leading-relaxed">
              An unexpected error occurred on our end. Please try again or
              contact support if the problem persists.
            </p>

            {error.digest && (
              <p className="mt-3 font-mono text-xs text-muted-foreground/60">
                Error ID: {error.digest}
              </p>
            )}

            <Separator className="my-8 max-w-xs" />

            <div className="flex gap-3">
              <Button size="lg" onClick={reset}>
                <RotateCcw className="h-4 w-4 mr-1.5" />
                Try again
              </Button>
              <Button variant="outline" size="lg" nativeButton={false} render={<a href="/" />}>
                <Home className="h-4 w-4 mr-1.5" />
                Home
              </Button>
            </div>
          </div>

          {/* Footer */}
          <footer className="flex items-center justify-center h-14 text-xs text-muted-foreground border-t">
            Kivia &mdash; API Observability
          </footer>
        </div>
      </body>
    </html>
  );
}
