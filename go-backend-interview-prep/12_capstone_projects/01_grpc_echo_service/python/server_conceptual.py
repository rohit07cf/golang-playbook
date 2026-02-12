"""
gRPC Echo Server -- Conceptual Python Equivalent

This shows what a gRPC server looks like in Python.
In production, you would use the 'grpcio' package.

This file is for LEARNING ONLY -- it does not actually run a gRPC server.
It demonstrates the structure and patterns.

Real setup:
    pip install grpcio grpcio-tools
    python -m grpc_tools.protoc -I. --python_out=. --grpc_python_out=. echo.proto
"""

from datetime import datetime, timezone


# --- What generated code would give you ---
# echo_pb2.EchoRequest      -- message class
# echo_pb2.EchoResponse     -- message class
# echo_pb2_grpc.EchoServiceServicer  -- base class to implement
# echo_pb2_grpc.add_EchoServiceServicer_to_server  -- registration function


# --- Server handler (what you implement) ---

class EchoServiceServicer:
    """
    Implements the EchoService defined in echo.proto.

    In real gRPC Python, this would subclass:
        echo_pb2_grpc.EchoServiceServicer
    """

    def Echo(self, request, context):
        """
        Handle an Echo RPC call.

        Args:
            request: EchoRequest with .message and .request_id
            context: gRPC context (carries deadline, metadata, etc.)

        Returns:
            EchoResponse with echoed message, request_id, and server_time
        """
        print(f"[server] received: message={request.message!r} "
              f"request_id={request.request_id!r}")

        # In real code, this returns an echo_pb2.EchoResponse object
        return {
            "message": f"echo: {request.message}",
            "request_id": request.request_id,
            "server_time": datetime.now(timezone.utc).isoformat(),
        }


# --- Server startup (what main() looks like) ---

def serve():
    """
    In real gRPC Python:

        import grpc
        from concurrent import futures

        server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
        echo_pb2_grpc.add_EchoServiceServicer_to_server(EchoServiceServicer(), server)
        server.add_insecure_port('[::]:50051')
        server.start()
        print("gRPC server listening on :50051")
        server.wait_for_termination()
    """
    print("This is a conceptual server -- see README.md for real gRPC setup.")
    print("In production, use: pip install grpcio grpcio-tools")
    print()

    # Demo: simulate handling a request
    servicer = EchoServiceServicer()

    class FakeRequest:
        def __init__(self, message, request_id):
            self.message = message
            self.request_id = request_id

    response = servicer.Echo(FakeRequest("hello from conceptual server", "req-001"), None)
    print(f"simulated response: {response}")


if __name__ == "__main__":
    serve()
