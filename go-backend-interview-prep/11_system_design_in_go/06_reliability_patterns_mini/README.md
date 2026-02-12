# 06 -- Reliability Patterns (Mini)

## What We Are Building

- Three core reliability patterns in one demo: **retry with backoff**, **timeout**, and **circuit breaker**
- **ELI10:** Retry is knocking again. Timeout is giving up after 10 seconds. Circuit breaker is putting a "closed" sign on the door.
- Shows how failures propagate and how each pattern mitigates them

## Requirements

**Functional:**
- Retry: call a flaky service with exponential backoff
- Timeout: cancel slow calls after a deadline
- Circuit breaker: stop calling a dead service, fail fast, auto-recover

**Non-functional:**
- Composable -- patterns can wrap each other
- Configurable thresholds (max retries, timeout duration, failure threshold)
- Clear logging of state transitions

## High-Level Design

```
Caller --> [Retry] --> [Timeout] --> [CircuitBreaker] --> Flaky Service
              |             |               |
         backoff +      context.         open/closed/
         jitter       WithTimeout       half-open
```

```
+--------+     +-------+     +---------+     +---------+     +-------+
| Caller | --> | Retry | --> | Timeout | --> | Circuit | --> | Svc   |
+--------+     +-------+     +---------+     | Breaker |     | (flaky|
                                             +---------+     +-------+
```

- **Retry**: re-attempts on transient failures with exponential backoff
- **Timeout**: wraps call with `context.WithTimeout`, cancels if too slow
- **Circuit breaker**: tracks failures; opens circuit after threshold; half-open to probe recovery

## Key Go Building Blocks Used

- `context.WithTimeout` -- enforce deadlines on downstream calls
- `time.Sleep` + exponential backoff -- retry delay
- `sync.Mutex` -- protect circuit breaker state
- `time.After` -- circuit breaker cooldown timer
- Error wrapping -- distinguish transient vs permanent failures

## Trade-Offs

- **Retry amplifies load** -- failing service gets hammered; add jitter to spread retries
- **Timeout too short** -- false failures on slow but healthy calls
- **Timeout too long** -- resources held, cascading slowness
- **Circuit breaker threshold** -- too low = flaps open/closed; too high = slow detection
- **In-process only** -- not shared across instances; production uses service mesh (Istio)
- **No bulkhead** -- one slow dependency can exhaust all goroutines

## Common Interview Traps

- Retrying non-idempotent operations (double payment!)
- Not adding jitter to backoff (thundering herd on retry)
- Confusing circuit breaker with rate limiter (different problems)
- Forgetting the half-open state (circuit breaker must probe for recovery)
- Not mentioning that retries + timeout interact (total time = retries * timeout)

## Run It

```bash
go run ./11_system_design_in_go/06_reliability_patterns_mini
python3 ./11_system_design_in_go/06_reliability_patterns_mini/main.py
```

## TL;DR

- **Retry**: exponential backoff + jitter; only for idempotent + transient failures
- **Timeout**: `context.WithTimeout` prevents hanging on slow services
- **Circuit breaker**: fail fast when downstream is dead; half-open probes for recovery
- These patterns compose: retry wraps timeout wraps circuit breaker
- Always mention jitter in interviews -- it prevents thundering herd
- In production, use a service mesh or library (e.g., resilience4j pattern)
