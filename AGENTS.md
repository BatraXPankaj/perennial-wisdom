# Perennial Wisdom — Agent Instructions

## Vision

A living bridge between ancient contemplative wisdom and modern neuroscience.
Stoic, Buddhist, Sufi, Vedantic, Taoist teachings — validated by empirical research.
Epictetus is the apex; other traditions reveal shared truths.

## Domain Model

```
Quote ──────► Philosopher ──────► Tradition
  │                                    │
  ├──► Theme ◄─────────────────────────┘
  │
  └──► Evidence (scientific validation)
```

- **Quote**: Central entity. Wisdom teaching with original text, scholarly context, expositions, practices.
- **Philosopher**: Teacher across traditions. Has era, bio, key teachings.
- **Tradition**: School of thought (Stoicism, Buddhism, Vedanta, Sufism, Taoism, Krishnamurti).
- **Theme**: Universal thread across traditions (impermanence, control, virtue, ego-dissolution).
- **Evidence**: Neuroscience/psychology research validating the insight. Must have citation.

## Architecture

```
Go + Gin + PostgreSQL + HTMX + Tailwind
```

- Server-rendered HTML with HTMX for interactivity
- No SPA, no JavaScript framework
- Explicit SQL via sqlc — no ORM magic
- Content lives in database, not code files
- JSONB for rich metadata (practices, science_notes, thought_experiments)

## Database Schema

Core fields are columns (queryable). Rich fields in JSONB `meta` column.

Key tables:
- `traditions` — schools of thought
- `philosophers` — wisdom teachers  
- `quotes` — central entity with expositions
- `themes` — universal threads
- `evidence` — scientific validation
- `quote_themes`, `quote_evidence` — join tables

## Code Organization

```
perennial-wisdom/
├── main.go          # Entry point, wiring
├── config.go        # Environment config
├── db/              # Database connection, migrations, queries
├── handlers/        # HTTP handlers
├── models/          # Domain types
├── router/          # Route setup
├── store/           # Legacy in-memory store (deprecate)
└── templates/       # HTML templates
```

## Principles

1. **Minimal code, rich content** — Keep codebase under 2000 lines
2. **Every quote traceable to source** — Source work, location, editions
3. **Science notes must have citations** — No vague claims
4. **Server-rendered simplicity** — HTMX, not React
5. **Explicit over magic** — Raw SQL via sqlc, not ORM

## Style

- Dark theme, amber accents, Crimson Pro serif for headings
- Contemplative, not flashy
- Mobile-responsive

## Never Do

- Don't add GraphQL, React, Vue, or complex frontend frameworks
- Don't use ORMs with magic methods (no GORM)
- Don't store content in code files or large JSON
- Don't create features without updating this document
- Don't add authentication/user accounts yet

## Current Focus

PostgreSQL migration + importing 73 Stoic quotes from observe-intelligence-ai.

## Content Source

Rich quote data lives in: `/home/panks/observe-intelligence-ai/data/v2/quotes-stoicism.json`
Fields to import: text, text_scholarly, expositions, source info, themes, evidence links, practices.
