"use client";

import { useState } from "react";
import Link from "next/link";
import {
  Zap,
  Activity,
  Shield,
  BarChart3,
  ArrowRight,
  Code2,
  Globe,
  Clock,
  Terminal,
  Gauge,
  Lock,
  CheckCircle2,
  Server,
  Layers,
  Eye,
  ChevronRight,
} from "lucide-react";
import { cn } from "@/lib/utils";

/* -- Scroll helper -------------------------------------------------------- */

function scrollTo(id: string) {
  document.getElementById(id)?.scrollIntoView({ behavior: "smooth" });
}

/* -- Terminal window component ------------------------------------------- */

function TerminalWindow({
  title,
  children,
  className,
}: {
  title: string;
  children: React.ReactNode;
  className?: string;
}) {
  return (
    <div
      className={cn(
        "overflow-hidden rounded-xl border border-border bg-card/80 shadow-2xl backdrop-blur-sm",
        className
      )}
    >
      <div className="flex items-center gap-2 border-b border-border bg-muted/20 px-4 py-2.5">
        <div className="flex gap-1.5">
          <div className="h-3 w-3 rounded-full bg-destructive/70" />
          <div className="h-3 w-3 rounded-full bg-chart-2/50" />
          <div className="h-3 w-3 rounded-full bg-emerald-500/60" />
        </div>
        <span className="ml-2 text-xs text-muted-foreground font-mono">
          {title}
        </span>
      </div>
      <div className="p-5 font-mono text-sm leading-relaxed">{children}</div>
    </div>
  );
}

/* -- Feature card -------------------------------------------------------- */

function FeatureCard({
  icon: Icon,
  title,
  description,
  code,
  codeTitle,
  codeAnimKey,
  slideClass,
}: {
  icon: React.ElementType;
  title: string;
  description: string;
  code: string;
  codeTitle: string;
  codeAnimKey?: React.Key;
  slideClass?: string;
}) {
  return (
    <div className="group rounded-2xl border border-border bg-card/40 p-6 backdrop-blur-sm hover:border-border hover:bg-card/80 hover:shadow-lg hover:shadow-foreground/[0.04] transition-all duration-300">
      <div className="flex items-start gap-3 mb-4">
        <div className="flex h-10 w-10 shrink-0 items-center justify-center rounded-xl bg-muted text-foreground ring-1 ring-border">
          <Icon className="h-5 w-5" />
        </div>
        <div>
          <h3 className="font-display font-semibold text-foreground">{title}</h3>
          <p className="text-sm text-muted-foreground mt-0.5 leading-relaxed">
            {description}
          </p>
        </div>
      </div>
      <div key={codeAnimKey} className={slideClass}>
        <TerminalWindow title={codeTitle}>
          <pre className="text-[13px] text-emerald-400/80 whitespace-pre-wrap">
            <code>{code}</code>
          </pre>
        </TerminalWindow>
      </div>
    </div>
  );
}

/* -- Step card ----------------------------------------------------------- */

function StepCard({
  step,
  icon: Icon,
  title,
  description,
  highlight,
}: {
  step: number;
  icon: React.ElementType;
  title: string;
  description: string;
  highlight: string;
}) {
  return (
    <div className="relative flex flex-col items-center text-center group">
      <div className="absolute -top-3 text-xs font-bold text-muted-foreground font-mono">
        0{step}
      </div>
      <div className="flex h-14 w-14 items-center justify-center rounded-2xl bg-muted text-foreground ring-1 ring-border transition-all group-hover:scale-110 group-hover:shadow-lg group-hover:shadow-foreground/5 mb-4 mt-3">
        <Icon className="h-6 w-6" />
      </div>
      <h3 className="font-display font-semibold text-foreground mb-1">{title}</h3>
      <p className="text-sm text-muted-foreground mb-3 max-w-[220px] leading-relaxed">
        {description}
      </p>
      <span className="text-xs font-mono text-muted-foreground bg-muted px-3 py-1.5 rounded-lg border border-border">
        {highlight}
      </span>
    </div>
  );
}

/* -- Page ---------------------------------------------------------------- */

type SdkLang = "go" | "express" | "hono" | "fastify" | "elysia";
const SDK_LANGS: SdkLang[] = ["go", "express", "hono", "fastify", "elysia"];
const SDK_LABELS: Record<SdkLang, string> = {
  go: "Go",
  express: "Express",
  hono: "Hono",
  fastify: "Fastify",
  elysia: "Elysia",
};

export default function LandingPage() {
  const [sdkLang, setSdkLang] = useState<SdkLang>("go");
  const [slideDir, setSlideDir] = useState<"left" | "right">("right");

  function handleLangChange(lang: SdkLang) {
    const newIdx = SDK_LANGS.indexOf(lang);
    const curIdx = SDK_LANGS.indexOf(sdkLang);
    setSlideDir(newIdx >= curIdx ? "right" : "left");
    setSdkLang(lang);
  }

  return (
    <div className="min-h-screen bg-background text-foreground">
      {/* -- Navigation --------------------------------------------------- */}
      <nav className="sticky top-0 z-50 border-b border-border bg-background/80 backdrop-blur-xl">
        <div className="mx-auto flex h-16 max-w-6xl items-center justify-between px-6">
          <Link href="/" className="flex items-center gap-2.5">
            <div className="flex h-8 w-8 items-center justify-center rounded-lg bg-foreground shadow-sm">
              <Zap className="h-4 w-4 text-background" />
            </div>
            <span className="text-lg font-display font-bold tracking-tight">Kivia</span>
          </Link>
          <div className="hidden sm:flex items-center gap-6 text-sm text-muted-foreground">
            <button
              onClick={() => scrollTo("features")}
              className="hover:text-foreground transition-colors"
            >
              Features
            </button>
            <button
              onClick={() => scrollTo("how-it-works")}
              className="hover:text-foreground transition-colors"
            >
              How it Works
            </button>
            <button
              onClick={() => scrollTo("sdk")}
              className="hover:text-foreground transition-colors"
            >
              SDK
            </button>
          </div>
          <div className="flex items-center gap-3">
            <Link
              href="/login"
              className="text-sm font-medium text-muted-foreground hover:text-foreground transition-colors px-3 py-1.5"
            >
              Sign in
            </Link>
            <Link
              href="/register"
              className="inline-flex items-center gap-1.5 rounded-lg bg-foreground px-4 py-2 text-sm font-medium text-background hover:bg-foreground/90 transition-all shadow-lg shadow-foreground/10 hover:shadow-foreground/15"
            >
              Get Started Free
              <ArrowRight className="h-3.5 w-3.5" />
            </Link>
          </div>
        </div>
      </nav>

      {/* -- Hero --------------------------------------------------------- */}
      <section className="relative overflow-hidden dot-grid">
        {/* Gradient orbs */}
        <div className="pointer-events-none absolute inset-0 bg-gradient-to-b from-foreground/[0.03] via-transparent to-transparent" />
        <div className="pointer-events-none absolute top-0 left-1/2 -translate-x-1/2 h-[600px] w-[900px] rounded-full bg-foreground/[0.04] blur-[100px]" />
        <div className="pointer-events-none absolute top-40 -left-32 h-[300px] w-[300px] rounded-full bg-chart-4/[0.04] blur-[80px]" />
        <div className="pointer-events-none absolute top-60 -right-32 h-[300px] w-[300px] rounded-full bg-chart-2/[0.04] blur-[80px]" />

        <div className="relative mx-auto max-w-6xl px-6 pb-20 pt-28 md:pb-36 md:pt-40">
          <div className="mx-auto max-w-3xl text-center">
            <div
              className="mb-8 inline-flex items-center gap-2 rounded-full border border-border bg-muted/30 px-4 py-1.5 text-sm text-muted-foreground backdrop-blur-sm animate-fade-in-up"
            >
              <span className="flex h-2 w-2 rounded-full bg-emerald-400 animate-pulse" />
              Now in public beta
            </div>

            <h1
              className="font-display text-5xl font-bold tracking-tight sm:text-6xl md:text-7xl lg:text-[80px] leading-[1.05] animate-fade-in-up"
              style={{ animationDelay: "80ms" }}
            >
              <span className="text-foreground">Observe your APIs</span>
              <br />
              <span>
                in real time
              </span>
            </h1>

            <p
              className="mx-auto mt-6 max-w-xl text-lg text-muted-foreground leading-relaxed md:text-xl animate-fade-in-up"
              style={{ animationDelay: "160ms" }}
            >
              Kivia is a lightweight API observability platform that captures
              every request, manages your keys, and gives you instant insights.
            </p>

            <div
              className="mt-10 flex flex-col items-center gap-4 sm:flex-row sm:justify-center animate-fade-in-up"
              style={{ animationDelay: "240ms" }}
            >
              <Link
                href="/register"
                className="inline-flex items-center gap-2 rounded-xl bg-foreground px-7 py-3.5 text-base font-medium text-background hover:bg-foreground/90 transition-all shadow-xl shadow-foreground/10 hover:shadow-foreground/15 hover:-translate-y-0.5"
              >
                Get Started Free
                <ArrowRight className="h-4 w-4" />
              </Link>
              <button
                onClick={() => scrollTo("features")}
                className="inline-flex items-center gap-2 rounded-xl border border-border bg-muted/30 px-7 py-3.5 text-base font-medium text-muted-foreground hover:text-foreground hover:bg-muted/50 transition-all backdrop-blur-sm"
              >
                See how it works
                <ChevronRight className="h-4 w-4" />
              </button>
            </div>
          </div>

          {/* Hero Terminal */}
          <div
            className="mx-auto mt-16 max-w-3xl md:mt-20 animate-fade-in-up"
            style={{ animationDelay: "400ms" }}
          >
            {/* Lang tabs */}
            <div className="flex gap-1 mb-3 w-fit flex-wrap">
              {SDK_LANGS.map((lang) => (
                <button
                  key={lang}
                  onClick={() => handleLangChange(lang)}
                  className={cn(
                    "px-3 py-1 rounded-md text-xs font-mono font-medium transition-all",
                    sdkLang === lang
                      ? "bg-foreground text-background"
                      : "text-muted-foreground hover:text-foreground"
                  )}
                >
                  {SDK_LABELS[lang]}
                </button>
              ))}
            </div>
            <TerminalWindow title={`terminal -- kivia (${SDK_LABELS[sdkLang]})`}>
              <div
                key={sdkLang}
                className={cn(
                  "space-y-1",
                  slideDir === "right" ? "animate-slide-in-right" : "animate-slide-in-left"
                )}
              >
                <p>
                  <span className="text-emerald-400">$</span>{" "}
                  <span className="text-foreground/70">
                    {sdkLang === "go"
                      ? "go get github.com/kivia-observe/kivia-sdk-go"
                      : "npm install @kivia/sdk"}
                  </span>
                </p>
                <p className="text-muted-foreground">
                  {sdkLang === "go" ? "Resolving dependencies..." : "added 1 package in 0.8s"}
                </p>
                <p className="text-emerald-400">
                  {sdkLang === "go"
                    ? "\u2713 kivia-sdk-go@v1.0.0 installed"
                    : "\u2713 @kivia/sdk@1.0.0 installed"}
                </p>
                <p className="mt-3">
                  <span className="text-emerald-400">$</span>{" "}
                  <span className="text-foreground/70">
                    {sdkLang === "go" ? "go run main.go" : "node index.js"}
                  </span>
                </p>
                <p className="text-chart-2">
                  {sdkLang === "go"
                    ? "Kivia middleware active on :8080"
                    : "Kivia middleware active on :3000"}
                </p>
                <p className="text-muted-foreground">
                  Listening for requests...
                </p>
                <p className="mt-3 text-muted-foreground/40">
                  &#9484;&#9472;&#9472;&#9472;&#9472;&#9472;&#9472;&#9472;&#9472;&#9472;&#9472;&#9472;&#9472;&#9472;&#9472;&#9472;&#9472;&#9472;&#9472;&#9472;&#9472;&#9472;&#9472;&#9472;&#9472;&#9472;&#9472;&#9472;&#9472;&#9472;&#9472;&#9472;&#9472;&#9472;&#9472;&#9472;&#9472;&#9472;&#9472;&#9472;&#9472;&#9472;&#9472;&#9472;&#9472;&#9472;&#9472;&#9472;&#9472;&#9472;&#9472;&#9472;&#9472;&#9472;&#9472;&#9488;
                </p>
                <p className="text-foreground/60">
                  {"  "}
                  <span className="text-chart-2 font-semibold">POST</span>
                  {" "}/auth/login{" "}
                  <span className="text-muted-foreground/30">
                    ................
                  </span>{" "}
                  <span className="text-emerald-400">200</span>{" "}
                  <span className="text-muted-foreground">45ms</span>
                </p>
                <p className="text-foreground/60">
                  {"  "}
                  <span className="text-emerald-400 font-semibold">GET</span>
                  {"  "}/api/v1/projects/all{" "}
                  <span className="text-muted-foreground/30">
                    .......
                  </span>{" "}
                  <span className="text-emerald-400">200</span>{" "}
                  <span className="text-muted-foreground">12ms</span>
                </p>
                <p className="text-foreground/60">
                  {"  "}
                  <span className="text-chart-2 font-semibold">POST</span>
                  {" "}/api/v1/logs/create{" "}
                  <span className="text-muted-foreground/30">
                    ........
                  </span>{" "}
                  <span className="text-emerald-400">201</span>{" "}
                  <span className="text-muted-foreground">8ms</span>
                </p>
                <p className="text-foreground/60">
                  {"  "}
                  <span className="text-emerald-400 font-semibold">GET</span>
                  {"  "}/api/v1/logs/all/abc123{" "}
                  <span className="text-muted-foreground/30">
                    ..
                  </span>{" "}
                  <span className="text-emerald-400">200</span>{" "}
                  <span className="text-muted-foreground">3ms</span>
                </p>
                <p className="text-muted-foreground/40">
                  &#9492;&#9472;&#9472;&#9472;&#9472;&#9472;&#9472;&#9472;&#9472;&#9472;&#9472;&#9472;&#9472;&#9472;&#9472;&#9472;&#9472;&#9472;&#9472;&#9472;&#9472;&#9472;&#9472;&#9472;&#9472;&#9472;&#9472;&#9472;&#9472;&#9472;&#9472;&#9472;&#9472;&#9472;&#9472;&#9472;&#9472;&#9472;&#9472;&#9472;&#9472;&#9472;&#9472;&#9472;&#9472;&#9472;&#9472;&#9472;&#9472;&#9472;&#9472;&#9472;&#9472;&#9472;&#9472;&#9496;
                </p>
                <p className="mt-1">
                  <span className="text-emerald-400">$</span>{" "}
                  <span className="animate-pulse text-muted-foreground">
                    &#9612;
                  </span>
                </p>
              </div>
            </TerminalWindow>
          </div>
        </div>
      </section>

      {/* -- Developer Experience ----------------------------------------- */}
      <section id="features" className="border-t border-border">
        <div className="mx-auto max-w-6xl px-6 py-24 md:py-32">
          <div className="mb-16 max-w-2xl">
            <p className="text-sm font-semibold text-chart-1 mb-3 tracking-wide uppercase font-display">
              Features
            </p>
            <h2 className="font-display text-3xl font-bold tracking-tight sm:text-4xl md:text-5xl text-foreground">
              First-class
              <br />
              developer experience
            </h2>
            <p className="mt-4 text-muted-foreground text-lg max-w-lg">
              Integrate Kivia in minutes. Native SDKs for Go and JavaScript —
              works with Express, Hono, Fastify, and Elysia. Zero config, zero overhead.
            </p>
          </div>

          <div className="grid gap-6 md:grid-cols-2">
            <FeatureCard
              icon={Code2}
              title="Online in three lines"
              description="Create a client, wrap your handler, and Kivia captures everything automatically."
              codeTitle={sdkLang === "go" ? "main.go" : "index.js"}
              code={
                sdkLang === "go"
                  ? `client := kiviasdk.NewClient("kivia_sk_...")\nhandler := client.NewLog(mux)\nhttp.ListenAndServe(":8080", handler)`
                  : sdkLang === "express"
                  ? `const client = new KiviaClient("kivia_sk_...");\napp.use(client.middleware());\napp.listen(3000);`
                  : sdkLang === "hono"
                  ? `const client = new KiviaClient("kivia_sk_...");\napp.use("*", client.middleware());\nexport default app;`
                  : sdkLang === "fastify"
                  ? `const client = new KiviaClient("kivia_sk_...");\nawait app.register(client.plugin());\nawait app.listen({ port: 3000 });`
                  : `const client = new KiviaClient("kivia_sk_...");\nnew Elysia().use(client.plugin()).listen(3000);`
              }
              codeAnimKey={sdkLang}
              slideClass={slideDir === "right" ? "animate-slide-in-right" : "animate-slide-in-left"}
            />
            <FeatureCard
              icon={Eye}
              title="Instant Observability"
              description="View live traffic, status codes, and latency in the dashboard."
              codeTitle="dashboard"
              code={`POST /auth/login            200  45ms
GET  /api/v1/projects/all   200  12ms
POST /api/v1/logs/create    201   8ms
GET  /api/v1/logs/all/abc   200   3ms`}
            />
            <FeatureCard
              icon={Lock}
              title="API Key Management"
              description="Create and revoke API keys per-project from the dashboard. Keys are hashed and stored securely."
              codeTitle="dashboard — api keys"
              code={`Name          Key              Status
production    kivia_sk_a3f•••    ● Active
staging       kivia_sk_7x2•••    ● Active
deprecated    kivia_sk_old•••    ○ Revoked`}
            />
            <FeatureCard
              icon={Layers}
              title="Multi-project support"
              description="Scope keys and logs per project. Ideal for microservice architectures."
              codeTitle="projects"
              code={`production-api    3 keys   12,847 reqs
staging-api       2 keys    1,203 reqs
internal-tools    1 key       456 reqs`}
            />
          </div>
        </div>
      </section>

      {/* -- How It Works ------------------------------------------------- */}
      <section
        id="how-it-works"
        className="border-t border-border bg-muted/20 dot-grid"
      >
        <div className="mx-auto max-w-6xl px-6 py-24 md:py-32">
          <div className="text-center mb-16">
            <p className="text-sm font-semibold text-chart-1 mb-3 tracking-wide uppercase font-display">
              How it works
            </p>
            <h2 className="font-display text-3xl font-bold tracking-tight sm:text-4xl md:text-5xl text-foreground">
              Up and running in minutes
            </h2>
          </div>

          <div className="grid gap-8 sm:grid-cols-2 lg:grid-cols-4">
            <StepCard
              step={1}
              icon={Terminal}
              title="Install SDK"
              description="Add Kivia to your Go or JavaScript project"
              highlight="go get / npm install"
            />
            <StepCard
              step={2}
              icon={Lock}
              title="Create API Key"
              description="Generate a key in the dashboard for your project"
              highlight="kivia_sk_..."
            />
            <StepCard
              step={3}
              icon={Server}
              title="Add Middleware"
              description="One line captures every request automatically"
              highlight="client.middleware()"
            />
            <StepCard
              step={4}
              icon={Gauge}
              title="View Dashboard"
              description="See live traffic, latency, and status codes instantly"
              highlight="dashboard >"
            />
          </div>
        </div>
      </section>

      {/* -- SDK Integration ----------------------------------------------- */}
      <section id="sdk" className="border-t border-border">
        <div className="mx-auto max-w-6xl px-6 py-24 md:py-32">
          <div className="grid gap-12 lg:grid-cols-2 items-center">
            <div>
              <p className="text-sm font-semibold text-emerald-400 mb-3 tracking-wide uppercase font-display">
                SDK Integration
              </p>
              <h2 className="font-display text-3xl font-bold tracking-tight sm:text-4xl text-foreground">
                Drop-in middleware
                <br />
                for any backend
              </h2>
              <p className="mt-4 text-muted-foreground text-lg leading-relaxed max-w-md">
                Our lightweight SDKs wrap your existing HTTP handlers. No code
                refactoring, no configuration files. Just import and observe.
              </p>
              <ul className="mt-8 space-y-3">
                {[
                  "Zero-config request capture",
                  "Automatic latency measurement",
                  "Async log shipping (no overhead)",
                  "Status code & path extraction",
                ].map((item) => (
                  <li
                    key={item}
                    className="flex items-center gap-3 text-sm text-muted-foreground"
                  >
                    <CheckCircle2 className="h-4 w-4 text-emerald-400 shrink-0" />
                    {item}
                  </li>
                ))}
              </ul>
            </div>

            <div className="lg:ml-auto lg:max-w-lg w-full">
              {/* SDK lang tabs */}
              <div className="flex gap-1 mb-3 flex-wrap">
                {SDK_LANGS.map((lang) => (
                  <button
                    key={lang}
                    onClick={() => handleLangChange(lang)}
                    className={cn(
                      "px-3 py-1 rounded-md text-xs font-mono font-medium transition-all",
                      sdkLang === lang
                        ? "bg-foreground text-background"
                        : "text-muted-foreground hover:text-foreground"
                    )}
                  >
                    {SDK_LABELS[lang]}
                  </button>
                ))}
              </div>

              <div
                key={sdkLang}
                className={slideDir === "right" ? "animate-slide-in-right" : "animate-slide-in-left"}
              >
              {sdkLang === "go" ? (
                <TerminalWindow title="main.go">
                  <pre className="text-[13px] leading-relaxed">
                    <code>
                      <span className="text-chart-1">package</span>{" "}
                      <span className="text-foreground/70">main</span>
                      {"\n\n"}
                      <span className="text-chart-1">import</span>
                      <span className="text-muted-foreground"> (</span>
                      {"\n"}
                      {"    "}
                      <span className="text-emerald-400/80">
                        &quot;net/http&quot;
                      </span>
                      {"\n"}
                      {"    "}
                      <span className="text-emerald-400/80">
                        &quot;github.com/kivia-observe/kivia-sdk-go&quot;
                      </span>
                      {"\n"}
                      <span className="text-muted-foreground">)</span>
                      {"\n\n"}
                      <span className="text-chart-1">func</span>{" "}
                      <span className="text-chart-2">main</span>
                      <span className="text-muted-foreground">()</span>{" "}
                      <span className="text-muted-foreground">{"{"}</span>
                      {"\n"}
                      {"    "}
                      <span className="text-muted-foreground">
                        // Initialize the client
                      </span>
                      {"\n"}
                      {"    "}
                      <span className="text-foreground/70">client</span>{" "}
                      <span className="text-chart-1">:=</span>{" "}
                      <span className="text-chart-2">kiviasdk.NewClient</span>
                      <span className="text-muted-foreground">(</span>
                      {"\n"}
                      {"        "}
                      <span className="text-emerald-400/80">
                        &quot;kivia_sk_your_key&quot;
                      </span>
                      <span className="text-muted-foreground">,</span>
                      {"\n"}
                      {"    "}
                      <span className="text-muted-foreground">)</span>
                      {"\n\n"}
                      {"    "}
                      <span className="text-muted-foreground">
                        // Wrap your handler — that&apos;s it
                      </span>
                      {"\n"}
                      {"    "}
                      <span className="text-foreground/70">handler</span>{" "}
                      <span className="text-chart-1">:=</span>{" "}
                      <span className="text-chart-2">client.NewLog</span>
                      <span className="text-muted-foreground">(</span>
                      <span className="text-foreground/70">mux</span>
                      <span className="text-muted-foreground">)</span>
                      {"\n\n"}
                      {"    "}
                      <span className="text-chart-2">http.ListenAndServe</span>
                      <span className="text-muted-foreground">(</span>
                      <span className="text-emerald-400/80">
                        &quot;:8080&quot;
                      </span>
                      <span className="text-muted-foreground">,</span>{" "}
                      <span className="text-foreground/70">handler</span>
                      <span className="text-muted-foreground">)</span>
                      {"\n"}
                      <span className="text-muted-foreground">{"}"}</span>
                    </code>
                  </pre>
                </TerminalWindow>
              ) : sdkLang === "express" ? (
                <TerminalWindow title="express.js">
                  <pre className="text-[13px] leading-relaxed">
                    <code>
                      <span className="text-chart-1">import</span>{" "}
                      <span className="text-muted-foreground">{"{ "}</span>
                      <span className="text-foreground/70">KiviaClient</span>
                      <span className="text-muted-foreground">{" }"}</span>{" "}
                      <span className="text-chart-1">from</span>{" "}
                      <span className="text-emerald-400/80">
                        &quot;@kivia/sdk&quot;
                      </span>
                      <span className="text-muted-foreground">;</span>
                      {"\n"}
                      <span className="text-chart-1">import</span>{" "}
                      <span className="text-foreground/70">express</span>{" "}
                      <span className="text-chart-1">from</span>{" "}
                      <span className="text-emerald-400/80">
                        &quot;express&quot;
                      </span>
                      <span className="text-muted-foreground">;</span>
                      {"\n\n"}
                      <span className="text-foreground/70">const</span>{" "}
                      <span className="text-foreground/70">app</span>{" "}
                      <span className="text-chart-1">=</span>{" "}
                      <span className="text-chart-2">express</span>
                      <span className="text-muted-foreground">();</span>
                      {"\n\n"}
                      <span className="text-muted-foreground">
                        // Initialize the client
                      </span>
                      {"\n"}
                      <span className="text-foreground/70">const</span>{" "}
                      <span className="text-foreground/70">client</span>{" "}
                      <span className="text-chart-1">=</span>{" "}
                      <span className="text-chart-1">new</span>{" "}
                      <span className="text-chart-2">KiviaClient</span>
                      <span className="text-muted-foreground">(</span>
                      <span className="text-emerald-400/80">
                        &quot;kivia_sk_your_key&quot;
                      </span>
                      <span className="text-muted-foreground">);</span>
                      {"\n\n"}
                      <span className="text-muted-foreground">
                        // Plug in the middleware — that&apos;s it
                      </span>
                      {"\n"}
                      <span className="text-foreground/70">app</span>
                      <span className="text-muted-foreground">.</span>
                      <span className="text-chart-2">use</span>
                      <span className="text-muted-foreground">(</span>
                      <span className="text-foreground/70">client</span>
                      <span className="text-muted-foreground">.</span>
                      <span className="text-chart-2">middleware</span>
                      <span className="text-muted-foreground">());</span>
                      {"\n\n"}
                      <span className="text-foreground/70">app</span>
                      <span className="text-muted-foreground">.</span>
                      <span className="text-chart-2">listen</span>
                      <span className="text-muted-foreground">(</span>
                      <span className="text-emerald-400/80">3000</span>
                      <span className="text-muted-foreground">);</span>
                    </code>
                  </pre>
                </TerminalWindow>
              ) : sdkLang === "hono" ? (
                <TerminalWindow title="index.ts">
                  <pre className="text-[13px] leading-relaxed">
                    <code>
                      <span className="text-chart-1">import</span>{" "}
                      <span className="text-muted-foreground">{"{ "}</span>
                      <span className="text-foreground/70">KiviaClient</span>
                      <span className="text-muted-foreground">{" }"}</span>{" "}
                      <span className="text-chart-1">from</span>{" "}
                      <span className="text-emerald-400/80">
                        &quot;@kivia/sdk&quot;
                      </span>
                      <span className="text-muted-foreground">;</span>
                      {"\n"}
                      <span className="text-chart-1">import</span>{" "}
                      <span className="text-muted-foreground">{"{ "}</span>
                      <span className="text-foreground/70">Hono</span>
                      <span className="text-muted-foreground">{" }"}</span>{" "}
                      <span className="text-chart-1">from</span>{" "}
                      <span className="text-emerald-400/80">
                        &quot;hono&quot;
                      </span>
                      <span className="text-muted-foreground">;</span>
                      {"\n\n"}
                      <span className="text-foreground/70">const</span>{" "}
                      <span className="text-foreground/70">app</span>{" "}
                      <span className="text-chart-1">=</span>{" "}
                      <span className="text-chart-1">new</span>{" "}
                      <span className="text-chart-2">Hono</span>
                      <span className="text-muted-foreground">();</span>
                      {"\n\n"}
                      <span className="text-muted-foreground">
                        // Initialize the client
                      </span>
                      {"\n"}
                      <span className="text-foreground/70">const</span>{" "}
                      <span className="text-foreground/70">client</span>{" "}
                      <span className="text-chart-1">=</span>{" "}
                      <span className="text-chart-1">new</span>{" "}
                      <span className="text-chart-2">KiviaClient</span>
                      <span className="text-muted-foreground">(</span>
                      <span className="text-emerald-400/80">
                        &quot;kivia_sk_your_key&quot;
                      </span>
                      <span className="text-muted-foreground">);</span>
                      {"\n\n"}
                      <span className="text-muted-foreground">
                        // Add middleware — that&apos;s it
                      </span>
                      {"\n"}
                      <span className="text-foreground/70">app</span>
                      <span className="text-muted-foreground">.</span>
                      <span className="text-chart-2">use</span>
                      <span className="text-muted-foreground">(</span>
                      <span className="text-emerald-400/80">&quot;*&quot;</span>
                      <span className="text-muted-foreground">,{" "}</span>
                      <span className="text-foreground/70">client</span>
                      <span className="text-muted-foreground">.</span>
                      <span className="text-chart-2">middleware</span>
                      <span className="text-muted-foreground">());</span>
                      {"\n\n"}
                      <span className="text-chart-1">export default</span>{" "}
                      <span className="text-foreground/70">app</span>
                      <span className="text-muted-foreground">;</span>
                    </code>
                  </pre>
                </TerminalWindow>
              ) : sdkLang === "fastify" ? (
                <TerminalWindow title="index.js">
                  <pre className="text-[13px] leading-relaxed">
                    <code>
                      <span className="text-chart-1">import</span>{" "}
                      <span className="text-muted-foreground">{"{ "}</span>
                      <span className="text-foreground/70">KiviaClient</span>
                      <span className="text-muted-foreground">{" }"}</span>{" "}
                      <span className="text-chart-1">from</span>{" "}
                      <span className="text-emerald-400/80">
                        &quot;@kivia/sdk&quot;
                      </span>
                      <span className="text-muted-foreground">;</span>
                      {"\n"}
                      <span className="text-chart-1">import</span>{" "}
                      <span className="text-foreground/70">Fastify</span>{" "}
                      <span className="text-chart-1">from</span>{" "}
                      <span className="text-emerald-400/80">
                        &quot;fastify&quot;
                      </span>
                      <span className="text-muted-foreground">;</span>
                      {"\n\n"}
                      <span className="text-foreground/70">const</span>{" "}
                      <span className="text-foreground/70">app</span>{" "}
                      <span className="text-chart-1">=</span>{" "}
                      <span className="text-chart-2">Fastify</span>
                      <span className="text-muted-foreground">();</span>
                      {"\n\n"}
                      <span className="text-muted-foreground">
                        // Initialize the client
                      </span>
                      {"\n"}
                      <span className="text-foreground/70">const</span>{" "}
                      <span className="text-foreground/70">client</span>{" "}
                      <span className="text-chart-1">=</span>{" "}
                      <span className="text-chart-1">new</span>{" "}
                      <span className="text-chart-2">KiviaClient</span>
                      <span className="text-muted-foreground">(</span>
                      <span className="text-emerald-400/80">
                        &quot;kivia_sk_your_key&quot;
                      </span>
                      <span className="text-muted-foreground">);</span>
                      {"\n\n"}
                      <span className="text-muted-foreground">
                        // Register plugin — that&apos;s it
                      </span>
                      {"\n"}
                      <span className="text-chart-1">await</span>{" "}
                      <span className="text-foreground/70">app</span>
                      <span className="text-muted-foreground">.</span>
                      <span className="text-chart-2">register</span>
                      <span className="text-muted-foreground">(</span>
                      <span className="text-foreground/70">client</span>
                      <span className="text-muted-foreground">.</span>
                      <span className="text-chart-2">plugin</span>
                      <span className="text-muted-foreground">());</span>
                      {"\n"}
                      <span className="text-chart-1">await</span>{" "}
                      <span className="text-foreground/70">app</span>
                      <span className="text-muted-foreground">.</span>
                      <span className="text-chart-2">listen</span>
                      <span className="text-muted-foreground">({"{ "}</span>
                      <span className="text-foreground/70">port</span>
                      <span className="text-muted-foreground">: </span>
                      <span className="text-emerald-400/80">3000</span>
                      <span className="text-muted-foreground">{' }});'}</span>
                    </code>
                  </pre>
                </TerminalWindow>
              ) : (
                <TerminalWindow title="index.ts">
                  <pre className="text-[13px] leading-relaxed">
                    <code>
                      <span className="text-chart-1">import</span>{" "}
                      <span className="text-muted-foreground">{"{ "}</span>
                      <span className="text-foreground/70">KiviaClient</span>
                      <span className="text-muted-foreground">{" }"}</span>{" "}
                      <span className="text-chart-1">from</span>{" "}
                      <span className="text-emerald-400/80">
                        &quot;@kivia/sdk&quot;
                      </span>
                      <span className="text-muted-foreground">;</span>
                      {"\n"}
                      <span className="text-chart-1">import</span>{" "}
                      <span className="text-muted-foreground">{"{ "}</span>
                      <span className="text-foreground/70">Elysia</span>
                      <span className="text-muted-foreground">{" }"}</span>{" "}
                      <span className="text-chart-1">from</span>{" "}
                      <span className="text-emerald-400/80">
                        &quot;elysia&quot;
                      </span>
                      <span className="text-muted-foreground">;</span>
                      {"\n\n"}
                      <span className="text-muted-foreground">
                        // Initialize the client
                      </span>
                      {"\n"}
                      <span className="text-foreground/70">const</span>{" "}
                      <span className="text-foreground/70">client</span>{" "}
                      <span className="text-chart-1">=</span>{" "}
                      <span className="text-chart-1">new</span>{" "}
                      <span className="text-chart-2">KiviaClient</span>
                      <span className="text-muted-foreground">(</span>
                      <span className="text-emerald-400/80">
                        &quot;kivia_sk_your_key&quot;
                      </span>
                      <span className="text-muted-foreground">);</span>
                      {"\n\n"}
                      <span className="text-muted-foreground">
                        // Plug in — that&apos;s it
                      </span>
                      {"\n"}
                      <span className="text-foreground/70">const</span>{" "}
                      <span className="text-foreground/70">app</span>{" "}
                      <span className="text-chart-1">=</span>{" "}
                      <span className="text-chart-1">new</span>{" "}
                      <span className="text-chart-2">Elysia</span>
                      <span className="text-muted-foreground">()</span>
                      {"\n"}
                      {"  "}
                      <span className="text-muted-foreground">.</span>
                      <span className="text-chart-2">use</span>
                      <span className="text-muted-foreground">(</span>
                      <span className="text-foreground/70">client</span>
                      <span className="text-muted-foreground">.</span>
                      <span className="text-chart-2">plugin</span>
                      <span className="text-muted-foreground">())</span>
                      {"\n"}
                      {"  "}
                      <span className="text-muted-foreground">.</span>
                      <span className="text-chart-2">listen</span>
                      <span className="text-muted-foreground">(</span>
                      <span className="text-emerald-400/80">3000</span>
                      <span className="text-muted-foreground">);</span>
                    </code>
                  </pre>
                </TerminalWindow>
              )}
              </div>
            </div>
          </div>
        </div>
      </section>

      {/* -- Capabilities Grid -------------------------------------------- */}
      <section className="border-t border-border bg-muted/20">
        <div className="mx-auto max-w-6xl px-6 py-24 md:py-32">
          <div className="text-center mb-16">
            <p className="text-sm font-semibold text-chart-2 mb-3 tracking-wide uppercase font-display">
              Capabilities
            </p>
            <h2 className="font-display text-3xl font-bold tracking-tight sm:text-4xl md:text-5xl">
              <span className="text-foreground">Everything you need,</span>
              <br />
              <span className="text-gradient-blue">
                nothing you don&apos;t.
              </span>
            </h2>
          </div>

          <div className="grid gap-px sm:grid-cols-2 lg:grid-cols-3 rounded-2xl border border-border overflow-hidden bg-border">
            {[
              {
                icon: Activity,
                title: "Real-time Monitoring",
                description:
                  "Watch every API request as it happens with live-updating logs and status indicators.",
              },
              {
                icon: Shield,
                title: "API Key Management",
                description:
                  "Generate, rotate, and revoke API keys with fine-grained per-project control.",
              },
              {
                icon: BarChart3,
                title: "Latency Analytics",
                description:
                  "Pinpoint slow endpoints with detailed latency breakdowns and percentile tracking.",
              },
              {
                icon: Code2,
                title: "Drop-in SDK",
                description:
                  "Three lines of code. Works with Express, Hono, Fastify, Elysia, and Go out of the box.",
              },
              {
                icon: Globe,
                title: "Per-Project Scoping",
                description:
                  "Isolated keys and logs per project. Perfect for multi-service architectures.",
              },
              {
                icon: Clock,
                title: "Historical Logs",
                description:
                  "Query past requests by date range, status code, and path with full context.",
              },
            ].map(({ icon: Icon, title, description }) => (
              <div
                key={title}
                className="group bg-background p-7 hover:bg-card transition-colors"
              >
                <div className="mb-4 flex h-10 w-10 items-center justify-center rounded-xl bg-muted text-muted-foreground ring-1 ring-border transition-colors group-hover:text-foreground group-hover:ring-foreground/10 group-hover:bg-foreground/5">
                  <Icon className="h-5 w-5" />
                </div>
                <h3 className="mb-1.5 font-display font-semibold text-foreground/90">
                  {title}
                </h3>
                <p className="text-sm leading-relaxed text-muted-foreground">
                  {description}
                </p>
              </div>
            ))}
          </div>
        </div>
      </section>

      {/* -- Stats -------------------------------------------------------- */}
      <section className="border-t border-border">
        <div className="mx-auto max-w-6xl px-6 py-20 md:py-28">
          <div className="grid gap-8 sm:grid-cols-3 text-center">
            {[
              {
                value: "< 2 min",
                label: "Integration time",
                color: "text-chart-1",
              },
              {
                value: "< 1ms",
                label: "Middleware overhead",
                color: "text-emerald-400",
              },
              {
                value: "100%",
                label: "Request capture rate",
                color: "text-chart-2",
              },
            ].map(({ value, label, color }) => (
              <div key={label}>
                <div
                  className={cn(
                    "font-display text-4xl font-bold tracking-tight md:text-6xl",
                    color
                  )}
                >
                  {value}
                </div>
                <p className="mt-3 text-muted-foreground font-medium">
                  {label}
                </p>
              </div>
            ))}
          </div>
        </div>
      </section>

      {/* -- CTA ---------------------------------------------------------- */}
      <section className="border-t border-border relative overflow-hidden">
        <div className="pointer-events-none absolute inset-0 bg-gradient-to-t from-foreground/[0.02] via-transparent to-transparent" />
        <div className="pointer-events-none absolute bottom-0 left-1/2 -translate-x-1/2 h-[400px] w-[700px] rounded-full bg-foreground/[0.03] blur-[100px]" />
        <div className="relative mx-auto max-w-6xl px-6 py-24 md:py-36">
          <div className="mx-auto max-w-2xl text-center">
            <h2 className="font-display text-3xl font-bold tracking-tight sm:text-4xl md:text-5xl text-foreground">
              Ready to see what
              <br />
              your APIs are doing?
            </h2>
            <p className="mt-5 text-lg text-muted-foreground">
              Set up Kivia in under two minutes. No credit card required.
            </p>
            <div className="mt-10 flex flex-col items-center gap-4 sm:flex-row sm:justify-center">
              <Link
                href="/register"
                className="inline-flex items-center gap-2 rounded-xl bg-foreground px-8 py-3.5 text-base font-medium text-background hover:bg-foreground/90 transition-all shadow-xl shadow-foreground/10 hover:shadow-foreground/15 hover:-translate-y-0.5"
              >
                Create free account
                <ArrowRight className="h-4 w-4" />
              </Link>
              <Link
                href="/login"
                className="inline-flex items-center gap-2 rounded-xl border border-border bg-muted/30 px-8 py-3.5 text-base font-medium text-muted-foreground hover:text-foreground hover:bg-muted/50 transition-all backdrop-blur-sm"
              >
                Sign in
              </Link>
            </div>
          </div>
        </div>
      </section>

      {/* -- Footer ------------------------------------------------------- */}
      <footer className="border-t border-border">
        <div className="mx-auto max-w-6xl px-6 py-12">
          <div className="grid gap-8 sm:grid-cols-2 lg:grid-cols-4">
            <div className="lg:col-span-1">
              <div className="flex items-center gap-2 mb-4">
                <div className="flex h-7 w-7 items-center justify-center rounded-md bg-foreground">
                  <Zap className="h-3.5 w-3.5 text-background" />
                </div>
                <span className="font-display font-bold">Kivia</span>
              </div>
              <p className="text-sm text-muted-foreground leading-relaxed max-w-xs">
                Lightweight API observability for modern backends. Monitor,
                manage, and understand your API traffic.
              </p>
            </div>

            <div>
              <p className="text-xs font-display font-semibold uppercase tracking-wider text-muted-foreground mb-4">
                Product
              </p>
              <ul className="space-y-2.5 text-sm text-muted-foreground">
                <li>
                  <button
                    onClick={() => scrollTo("features")}
                    className="hover:text-foreground transition-colors"
                  >
                    Features
                  </button>
                </li>
                <li>
                  <button
                    onClick={() => scrollTo("sdk")}
                    className="hover:text-foreground transition-colors"
                  >
                    SDK
                  </button>
                </li>
                <li>
                  <button
                    onClick={() => scrollTo("how-it-works")}
                    className="hover:text-foreground transition-colors"
                  >
                    How it Works
                  </button>
                </li>
              </ul>
            </div>

            <div>
              <p className="text-xs font-display font-semibold uppercase tracking-wider text-muted-foreground mb-4">
                Resources
              </p>
              <ul className="space-y-2.5 text-sm text-muted-foreground">
                <li>
                  <Link
                    href="/login"
                    className="hover:text-foreground transition-colors"
                  >
                    Dashboard
                  </Link>
                </li>
                <li>
                  <Link
                    href="/register"
                    className="hover:text-foreground transition-colors"
                  >
                    Get Started
                  </Link>
                </li>
              </ul>
            </div>

            <div>
              <p className="text-xs font-display font-semibold uppercase tracking-wider text-muted-foreground mb-4">
                Legal
              </p>
              <ul className="space-y-2.5 text-sm text-muted-foreground/50">
                <li>Terms</li>
                <li>Privacy</li>
              </ul>
            </div>
          </div>

          <div className="mt-10 pt-6 border-t border-border flex flex-col sm:flex-row items-center justify-between gap-3">
            <p className="text-xs text-muted-foreground/50">
              &copy; {new Date().getFullYear()} Kivia. All rights reserved.
            </p>
            <p className="text-xs text-muted-foreground/30">
              Built for developers who care about their APIs.
            </p>
          </div>
        </div>
      </footer>
    </div>
  );
}
