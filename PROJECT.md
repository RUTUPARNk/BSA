# BSA — the story

Every project before this that wired LLMs together — [AMKGC](../AMKGC/PROJECT.md)
with its conflict and hallucination tracking, the [brain](../My_attempt_at_brain/PROJECT.md)
with its reconciliation, [psychometric](../psychometric/PROJECT.md) — kept running
into the same wall. **Multi-agent LLM systems lose the plot.** Agents act on stale
or conflicting views of the world, hallucinate state that was never there, and
quietly corrupt each other's data until the whole thing degrades into noise.

BSA is what you build when you've hit that wall enough times to stop patching it
per-app and go build the missing piece underneath.

**Deterministic State Authority** — a centralized source of truth for agent
world-state, the way git is a source of truth for code. Every agent sees a
**version-locked, immutable snapshot** for a given version (no two agents
disagreeing about "now"). Changes come in as *proposals*, and a concurrent-safe
**reconciliation loop** validates them and commits atomically to the canonical
state. Agents propose; the authority decides; nobody drifts.

Two things make it notable in the arc:
- It's an **infrastructure** move, not another app. The instinct matured — from
  "build the thing" to "build the thing that fixes the recurring problem in all my
  other things."
- He reached for **Go** for the high-throughput core (his first real Go), keeping
  Python only as the thin client layer for easy agent integration. Right tool for a
  concurrent, throughput-sensitive job — a deliberate, grown-up architecture choice.

## What it is

- `bsa-core-go/` — the Go service: a thread-safe `BSA` holding the live index
  (`sync.RWMutex`), a proposal/reconciliation model, an HTTP API
  (`GET /api/v1/state?version=…`, `POST /api/v1/propose`), Dockerized.
- `bsa_client.py` + `models.py` — a Python client with pydantic models so agents
  integrate in a few lines.
- `BSA.pdf` — the design write-up framing it as a systems problem (state drift) and
  its solution.

## Note

Early/prototype stage (some methods are honest stubs — `GetState` returns the live
index with a `// In a real implementation, it would checkout the specific git
version` note, so versioning is designed but not fully wired). The README points to
an `example_usage.ipynb` that isn't in the repo. The idea and the architecture are
the substance here.
