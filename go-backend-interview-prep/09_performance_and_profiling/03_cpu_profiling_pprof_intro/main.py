"""CPU profiling -- Python equivalent of the Go example.

Uses cProfile (stdlib) to profile workloads and pstats to display results.
"""

import cProfile
import hashlib
import pstats
import io


def cpu_intensive_work(iterations: int) -> bytes:
    """Repeated hashing to create measurable CPU load."""
    data = b"hello world performance profiling"
    for _ in range(iterations):
        data = hashlib.sha256(data).digest()
    return data


def string_work(iterations: int) -> str:
    """String manipulation hotspot."""
    s = ""
    for _ in range(iterations):
        s += "x"
        if len(s) > 1000:
            s = s[:100]
    return s


def workload() -> None:
    """Combined workload to profile."""
    cpu_intensive_work(200_000)
    string_work(100_000)


def main() -> None:
    print("=== CPU Profiling with cProfile ===\n")

    # Profile the workload
    profiler = cProfile.Profile()
    profiler.enable()
    workload()
    profiler.disable()

    # Display results
    print("--- Top 10 functions by cumulative time ---\n")
    stream = io.StringIO()
    stats = pstats.Stats(profiler, stream=stream)
    stats.sort_stats("cumulative")
    stats.print_stats(10)
    print(stream.getvalue())

    print("--- Top 10 by total (self) time ---\n")
    stream2 = io.StringIO()
    stats2 = pstats.Stats(profiler, stream=stream2)
    stats2.sort_stats("tottime")
    stats2.print_stats(10)
    print(stream2.getvalue())

    print("--- How to analyze ---")
    print("  python -m cProfile -s cumulative script.py")
    print("  python -m cProfile -o output.prof script.py")
    print()
    print("Key differences from Go pprof:")
    print("  - cProfile is deterministic (counts every call)")
    print("  - pprof is sampling-based (samples at intervals)")
    print("  - cProfile has higher overhead -- not for production")
    print("  - Go's pprof is production-safe (sampling = low overhead)")


if __name__ == "__main__":
    main()
