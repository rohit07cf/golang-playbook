"""
gRPC Echo Client -- Minimal Python Equivalent

This script simulates what a gRPC client call looks like in Python.
It uses a mock channel since we cannot import grpcio (stdlib only).

The structure mirrors the real gRPC Python client pattern.

Real client would look like:
    import grpc
    channel = grpc.insecure_channel('localhost:50051')
    stub = echo_pb2_grpc.EchoServiceStub(channel)
    response = stub.Echo(echo_pb2.EchoRequest(message="hello", request_id="001"))
"""

from datetime import datetime, timezone


# --- Mock gRPC types (simulating generated code) ---

class EchoRequest:
    def __init__(self, message="", request_id=""):
        self.message = message
        self.request_id = request_id


class EchoResponse:
    def __init__(self, message="", request_id="", server_time=""):
        self.message = message
        self.request_id = request_id
        self.server_time = server_time

    def __repr__(self):
        return (f"EchoResponse(message={self.message!r}, "
                f"request_id={self.request_id!r}, "
                f"server_time={self.server_time!r})")


# --- Mock stub (simulates server behavior locally) ---

class MockEchoServiceStub:
    """
    In real gRPC:
        stub = echo_pb2_grpc.EchoServiceStub(channel)
        response = stub.Echo(request, timeout=5)
    """

    def Echo(self, request, timeout=None):
        if timeout is not None and timeout < 0.01:
            raise TimeoutError("deadline exceeded")

        return EchoResponse(
            message=f"echo: {request.message}",
            request_id=request.request_id,
            server_time=datetime.now(timezone.utc).isoformat(),
        )


# --- Client demo ---

def main():
    # In real gRPC:
    #   channel = grpc.insecure_channel('localhost:50051')
    #   stub = echo_pb2_grpc.EchoServiceStub(channel)
    stub = MockEchoServiceStub()

    print("=== gRPC Echo Client (Python mock) ===")

    # Request 1: normal echo
    print("\n--- Request 1: normal echo ---")
    resp = stub.Echo(EchoRequest(message="hello gRPC", request_id="req-001"), timeout=5)
    print(f"response: {resp}")

    # Request 2: with tracing ID
    print("\n--- Request 2: with tracing ID ---")
    resp = stub.Echo(EchoRequest(message="interview prep", request_id="trace-abc-123"), timeout=5)
    print(f"response: {resp}")

    # Request 3: deadline exceeded
    print("\n--- Request 3: deadline exceeded (simulated) ---")
    try:
        resp = stub.Echo(
            EchoRequest(message="this should timeout", request_id="req-timeout"),
            timeout=0.001,
        )
        print(f"unexpected success: {resp}")
    except TimeoutError as e:
        print(f"expected error: {e}")

    print("\ndone")


if __name__ == "__main__":
    main()
