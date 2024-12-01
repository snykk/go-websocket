## WebSocket Implementations: Basic vs Advanced

### Overview

These implementations represent **basic examples** designed for **learning purposes**. They demonstrate foundational WebSocket concepts but lack features like authentication, scalability, and advanced error handling.

---

### Basic WebSocket

-   **Purpose**: Simple implementation for one-to-one communication.
-   **Key Features**:
    -   Handles a single client at a time.
    -   Echoes messages back to the sender.
    -   Easy to implement, suitable for learning WebSocket basics.
-   **Usage**:
    -   Best for lightweight applications where broadcasting is unnecessary.

---

### Advanced WebSocket

-   **Purpose**: Introduction to multi-client communication.
-   **Key Features**:
    -   Maintains a list of connected clients.
    -   Broadcasts messages to all clients.
    -   Utilizes `sync.Mutex` for thread-safe access to the client registry.
-   **Usage**:
    -   Ideal for chat systems, notifications, or collaborative applications.

---

### Simulation Steps

1. **Basic**:
    - Run the server and connect with a single client.
    - Send messages and observe the server's direct response.
2. **Advanced**:
    - Run the server and connect multiple clients.
    - Send a message from one client to observe its broadcast to all connected clients.
