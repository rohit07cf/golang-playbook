"""Reliability patterns -- Python equivalent: retry, timeout, circuit breaker."""

import random
import threading
import time


# --- Retry with Exponential Backoff ---

def retry(max_attempts, base_delay, fn):
    """Retry fn with exponential backoff + jitter."""
    last_err = None
    for i in range(max_attempts):
        try:
            fn()
            return None
        except Exception as e:
            last_err = e
            if i < max_attempts - 1:
                backoff = base_delay * (2 ** i)
                jitter = random.uniform(0, backoff / 2)
                wait = backoff + jitter
                print(f"  [retry] attempt {i+1} failed: {e}, waiting {wait:.3f}s")
                time.sleep(wait)
    return last_err


# --- Timeout ---

def with_timeout(timeout_sec, fn):
    """Run fn with a timeout. Returns (result, error)."""
    result = [None]
    error = [None]

    def wrapper():
        try:
            fn()
        except Exception as e:
            error[0] = e

    t = threading.Thread(target=wrapper)
    t.start()
    t.join(timeout=timeout_sec)

    if t.is_alive():
        return "timeout"
    return error[0]


# --- Circuit Breaker ---

class CircuitBreaker:
    CLOSED = "CLOSED"
    OPEN = "OPEN"
    HALF_OPEN = "HALF-OPEN"

    def __init__(self, threshold, cooldown_sec):
        self.lock = threading.Lock()
        self.state = self.CLOSED
        self.failures = 0
        self.threshold = threshold
        self.cooldown = cooldown_sec
        self.last_failure_time = 0

    def call(self, fn):
        with self.lock:
            if self.state == self.OPEN:
                if time.time() - self.last_failure_time > self.cooldown:
                    self.state = self.HALF_OPEN
                    print(f"  [circuit] state: OPEN -> HALF-OPEN (probing)")
                else:
                    return "circuit breaker is open"

        # Execute
        try:
            fn()
            err = None
        except Exception as e:
            err = e

        with self.lock:
            if err is not None:
                self.failures += 1
                self.last_failure_time = time.time()
                if self.failures >= self.threshold:
                    prev = self.state
                    self.state = self.OPEN
                    if prev != self.OPEN:
                        print(f"  [circuit] state: {prev} -> OPEN (failures={self.failures})")
                return str(err)

            # Success
            if self.state == self.HALF_OPEN:
                print(f"  [circuit] state: HALF-OPEN -> CLOSED (recovered)")
            self.failures = 0
            self.state = self.CLOSED
            return None


# --- Flaky Service ---

class FlakyService:
    def __init__(self, fail_count):
        self.fail_count = fail_count
        self.call_count = 0

    def call(self):
        self.call_count += 1
        if self.call_count <= self.fail_count:
            raise Exception("service failure")


# --- Demo ---

def main():
    print("=== 1. Retry with Exponential Backoff ===\n")
    svc = FlakyService(fail_count=2)
    err = retry(5, 0.1, svc.call)
    if err:
        print(f"retry result: FAILED -- {err}")
    else:
        print(f"retry result: SUCCESS on attempt {svc.call_count}")

    print("\n=== 2. Timeout ===\n")
    # Fast call
    err = with_timeout(0.5, lambda: time.sleep(0.05))
    print(f"fast call:  {err}")

    # Slow call
    err = with_timeout(0.2, lambda: time.sleep(1.0))
    print(f"slow call:  {err}")

    print("\n=== 3. Circuit Breaker ===\n")
    cb = CircuitBreaker(threshold=3, cooldown_sec=1.0)
    svc2 = FlakyService(fail_count=5)

    for i in range(1, 9):
        err = cb.call(svc2.call)
        print(f"call {i}: err={err}")
        if err == "circuit breaker is open":
            print("  (circuit is open, call was rejected)")
        time.sleep(0.2)

    print("\nwaiting 1.2s for cooldown...")
    time.sleep(1.2)

    err = cb.call(svc2.call)
    print(f"call after cooldown: err={err}")

    print("\ndemo done")


if __name__ == "__main__":
    main()
