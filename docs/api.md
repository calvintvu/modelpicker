# HTTP API

The picker returns configured IDs. It never calls providers or generates completions.

## GET /healthz

Returns **200**, empty body, without authentication. The process still needs API_KEY to start.
The listener opens after configuration validation, native runtime/tokenizer/model loading,
and classifier-fit checks. During initialization, no HTTP response is available, including 503.
Startup failures exit before listening.

Health checks HTTP responsiveness only. It does not run inference or repeat startup validation.
Inference can fail while health still returns 200.

## POST /v1/pick

Replaces `/v1/route`; update existing clients to use `/v1/pick`.

```sh
curl --fail http://localhost:8080/v1/pick \
  -H "Authorization: Bearer $API_KEY" -H 'Content-Type: application/json' \
  --data '{"prompt":"Explain a Go function."}'
```

Bearer scheme matching is case-insensitive; the token must match exactly.
Authentication precedes body validation. Content-Type is recommended, not enforced.
The body must be one JSON object with a nonempty string prompt. Unknown fields
are ignored; whitespace-only prompts are accepted. Null/empty/missing prompts,
wrong types, and trailing JSON values are rejected.

Success: **200**, `Content-Type: application/json`, exactly two fields and a final newline:

```json
{"tier":"code","model":"your-code-model"}
```

### Selection

- Highest classifier score wins; ties select the earliest YAML tier.
- One model list entry is then selected at random within that tier.
- Temperature rescales logits before softmax; it does not sample tiers or models.
  Zero selects tier-count calibration. Scores are not exposed as validated confidence.
- `default_tier` handles inference errors only, not low confidence. Fallback returns
  normal 200 JSON without a fallback flag. Logs record its use.
- One tier skips classification per request. Startup still loads the model.
  Initialization failures cannot use fallback.

### Limits

| Setting | Behavior |
| --- | --- |
| Body | 1 MiB, including JSON/whitespace; overflow returns 400. |
| Input | 512 tokenizer IDs total, including question, options, prompt, and special tokens. |
| Prompt | Leading tokens that fit are kept; truncation is silent. One-tier selection skips tokenization. |
| Configuration | 1-10 tiers. Multi-tier question/options must leave at least one prompt-token slot or startup fails. |
| HTTP deadlines | Header read 5s, request read 10s, response write WRITE_TIMEOUT (default 60s), idle 120s. |
| Headers | 32 KiB configured budget; Go can reject oversized headers with 431. |

There is no separate character limit. See [configuration](configuration.md)
for all fields. Timeouts or process failures can close a connection without a full response.

### Errors

Errors are `text/plain; charset=utf-8`, with a final newline, not JSON.

| Status | Body | Cause |
| --- | --- | --- |
| 400 | `invalid JSON body` | Bad JSON/types, trailing values, or body overflow. |
| 400 | `prompt is required` | Missing, null, or empty prompt. |
| 401 | `unauthorized request` | Missing/invalid token; includes `WWW-Authenticate: Bearer`. |
| 500 | `internal server error` | Selection failed without a result; details stay in logs. |
| 404 | `404 page not found` | Unknown path. |
| 405 | `Method Not Allowed` | Wrong method; includes `Allow`. |

The mux generates 404/405. Header rejection occurs before the route handler.
Implementation: [main.go](../main.go), [selection](../inference/pick.go),
[validation](../config/config.go). Checks: [testing](testing.md).
