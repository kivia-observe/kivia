import { useEffect, useRef, useState, useCallback } from "react";
import { getAccessToken } from "@/lib/auth";
import { BASE_URL, type Log } from "@/lib/api";

export function useLogStream(projectId: string, enabled: boolean) {
  const [logs, setLogs] = useState<Log[]>([]);
  const [connected, setConnected] = useState(false);
  const abortRef = useRef<AbortController | null>(null);

  const clear = useCallback(() => setLogs([]), []);

  useEffect(() => {
    if (!enabled) {
      abortRef.current?.abort();
      setConnected(false);
      return;
    }

    const controller = new AbortController();
    abortRef.current = controller;

    async function connect() {
      try {
        const token = getAccessToken();
        const res = await fetch(
          `${BASE_URL}/api/v1/logs/stream/${projectId}`,
          {
            headers: token ? { Authorization: `Bearer ${token}` } : {},
            signal: controller.signal,
          }
        );

        if (!res.ok || !res.body) {
          setConnected(false);
          return;
        }

        setConnected(true);
        const reader = res.body.getReader();
        const decoder = new TextDecoder();
        let buffer = "";

        // eslint-disable-next-line no-constant-condition
        while (true) {
          const { done, value } = await reader.read();
          if (done) break;

          buffer += decoder.decode(value, { stream: true });
          const lines = buffer.split("\n");
          buffer = lines.pop() || "";

          for (const line of lines) {
            if (line.startsWith("data: ")) {
              try {
                const log = JSON.parse(line.slice(6)) as Log;
                setLogs((prev) => [log, ...prev]);
              } catch {
                // ignore malformed SSE data
              }
            }
          }
        }
      } catch (err) {
        if (err instanceof DOMException && err.name === "AbortError") return;
        console.error("SSE connection error:", err);
      } finally {
        setConnected(false);
      }
    }

    connect();

    return () => {
      controller.abort();
    };
  }, [projectId, enabled]);

  return { logs, connected, clear };
}
