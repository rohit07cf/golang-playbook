"""Cache service with TTL and LRU eviction -- Python equivalent."""

import threading
import time
from collections import OrderedDict


class CacheEntry:
    def __init__(self, value, expires_at):
        self.value = value
        self.expires_at = expires_at


class Cache:
    def __init__(self, capacity):
        self.lock = threading.Lock()
        self.items = OrderedDict()  # key -> CacheEntry (OrderedDict tracks access order)
        self.capacity = capacity

        # Start periodic cleanup
        t = threading.Thread(target=self._cleanup_loop, daemon=True)
        t.start()

    def set(self, key, value, ttl_seconds):
        with self.lock:
            if key in self.items:
                # Update existing: remove and re-insert at end (most recent)
                del self.items[key]

            if len(self.items) >= self.capacity:
                self._evict_lru()

            self.items[key] = CacheEntry(value, time.time() + ttl_seconds)

    def get(self, key):
        with self.lock:
            if key not in self.items:
                return None, False

            entry = self.items[key]

            # Lazy expiration
            if time.time() > entry.expires_at:
                del self.items[key]
                return None, False

            # Move to end (most recently used)
            self.items.move_to_end(key)
            return entry.value, True

    def size(self):
        with self.lock:
            return len(self.items)

    def _evict_lru(self):
        """Remove least recently used entry. Caller must hold lock."""
        if self.items:
            key, _ = self.items.popitem(last=False)  # pop from front (oldest)
            print(f"  [evict] key={key!r} (LRU)")

    def _cleanup_loop(self):
        while True:
            time.sleep(1)
            with self.lock:
                now = time.time()
                expired = [k for k, e in self.items.items() if now > e.expires_at]
                for k in expired:
                    print(f"  [cleanup] expired key={k!r}")
                    del self.items[k]


# --- Demo ---

def main():
    cache = Cache(capacity=3)

    print("=== cache demo (capacity=3) ===\n")

    cache.set("user:1", "Alice", ttl_seconds=5)
    cache.set("user:2", "Bob", ttl_seconds=2)
    cache.set("user:3", "Charlie", ttl_seconds=5)
    print(f"set 3 entries, cache size: {cache.size()}\n")

    # Cache hit
    val, ok = cache.get("user:1")
    if ok:
        print(f"GET user:1 -> {val} (hit)")

    val, ok = cache.get("user:2")
    if ok:
        print(f"GET user:2 -> {val} (hit)")

    # Cache miss
    _, ok = cache.get("user:99")
    if not ok:
        print("GET user:99 -> (miss)")

    # Trigger LRU eviction
    print("\nadding user:4 (should evict LRU)...")
    cache.set("user:4", "Diana", ttl_seconds=5)
    print(f"cache size: {cache.size()}")

    _, ok = cache.get("user:3")
    if not ok:
        print("GET user:3 -> (miss, was evicted as LRU)")

    # Wait for TTL expiration
    print("\nwaiting 3s for user:2 TTL to expire...")
    time.sleep(3)

    _, ok = cache.get("user:2")
    if not ok:
        print("GET user:2 -> (miss, expired)")

    print(f"\nfinal cache size: {cache.size()}")
    print("\ndemo done")


if __name__ == "__main__":
    main()
