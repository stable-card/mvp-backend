# Flex-Perks: Programmable Rewards Card MVP

Flex-Perks is an MVP project for a programmable credit/check card benefits card that individuals can design themselves.

## Core Idea

The core value is to allow users to describe the benefits they want in natural language, and an LLM compiles the combination of benefits into a policy that is enforced on-chain. When paying, benefits that match the policy (base cashback + brand top-up) are paid immediately. Merchants/brands can make precise targeting by putting a bonus budget for each time, geo, and category, and registration is operated decentrally (based on a deposit).

## Architecture

The project follows a standard Go backend layered architecture.

```mermaid
graph TD
    subgraph "API Layer"
        A[Router] --> B[Card Handler]
        A --> C[Policy Handler]
    end

    subgraph "Service Layer"
        B --> D[Card Service]
        C --> E[Policy Service]
    end

    subgraph "Repository Layer"
        D --> F[Card Repository]
        D --> G[Policy Repository]
    end

    subgraph "Domain"
        H[Card]
        I[Policy]
    end

    D -- uses --> H
    E -- uses --> I
    F -- manages --> H
    G -- manages --> I

    style A fill:#f9f,stroke:#333,stroke-width:2px
    style B fill:#ccf,stroke:#333,stroke-width:2px
    style C fill:#ccf,stroke:#333,stroke-width:2px
    style D fill:#9cf,stroke:#333,stroke-width:2px
    style E fill:#9cf,stroke:#333,stroke-width:2px
    style F fill:#c9f,stroke:#333,stroke-width:2px
    style G fill:#c9f,stroke:#333,stroke-width:2px
    style H fill:#f99,stroke:#333,stroke-width:2px
    style I fill:#f99,stroke:#333,stroke-width:2px
```

## Payment Flow Diagram

Here is a diagram illustrating the payment flow based on the project proposal.

```mermaid
sequenceDiagram
    participant UserApp as User App
    participant Attestor
    participant SmartAccount as Smart Account (AA)
    participant PolicyGuard
    participant RewardVault
    participant TopupPool

    UserApp->>Attestor: Request payment context signature (time, geo, etc.)
    activate Attestor
    Attestor-->>UserApp: Return EIP-712 signature
    deactivate Attestor

    UserApp->>SmartAccount: Call authorizeAndPay(PaymentIntent, Attestation)
    activate SmartAccount

    SmartAccount->>PolicyGuard: authorizeAndPay(...)
    activate PolicyGuard
    
    note right of PolicyGuard: 1. Verify policy hash<br/>2. Check rules (time, geo, category)<br/>3. Check caps (monthly, per-transaction)<br/>4. Check for brand top-ups

    alt Payment is valid
        PolicyGuard->>RewardVault: Execute payment & payout rewards
        activate RewardVault
        RewardVault-->>PolicyGuard: Success
        deactivate RewardVault
        PolicyGuard-->>SmartAccount: Return receipt ID
    else Payment is invalid
        PolicyGuard-->>SmartAccount: Revert with error code
    end
    deactivate PolicyGuard

    SmartAccount-->>UserApp: Transaction result
    deactivate SmartAccount
```

## API Endpoints

The MVP backend exposes the following endpoints:

- `POST /api/v1/policies/compile`: Compiles a natural language user request into a policy. In the current MVP, this endpoint simulates the compilation of a policy.
- `POST /api/v1/cards/issue`: Issues a new card for a user with a specific policy.

## Getting Started

### Prerequisites

- Go 1.18 or higher

### Installation & Running

1.  **Clone the repository:**
    ```bash
    git clone https://github.com/rrabit42/mvp-backend.git
    cd mvp-backend
    ```

2.  **Create a configuration file:**
    Create a `config.toml` file in the root of the project with the server port.
    ```toml
    [server]
    port = "8080"
    ```

3.  **Install dependencies:**
    ```bash
    go mod tidy
    ```

4.  **Run the server:**
    ```bash
    go run cmd/server/main.go
    ```
    The server will start on the port specified in your `config.toml`.

## Project Structure

```
mvp-backend/
├── cmd/server/main.go      # Application entry point
├── internal/
│   ├── api/                # API handlers and router setup
│   ├── config/             # Configuration loading
│   ├── domain/             # Core domain models (Card, Policy)
│   ├── repository/         # Data access layer (in-memory for MVP)
│   └── service/            # Business logic layer
├── go.mod
├── go.sum
└── README.md
```

## Future Work (Based on Proposal)

- **LLM Integration**: Implement the LLM policy compiler to translate natural language to a policy DSL.
- **Full Database Integration**: Replace in-memory repositories with a persistent database.
- **Brand/Merchant Features**: Implement the Brand Console and the decentralized merchant registration system.
- **Authentication & Security**: Add robust authentication and security measures.
