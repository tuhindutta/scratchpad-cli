# Scratchpad CLI

<img width="911" height="485" alt="image" src="https://github.com/user-attachments/assets/c9c7fe00-247f-457f-8f53-c66a31436e62" />

A lightweight, terminal-first AI assistant client written in Go.

Scratchpad CLI provides a clean command-line interface for interacting with a locally hosted AI backend. It abstracts backend lifecycle management, session persistence, document ingestion, and streaming AI interactions behind a simple developer-friendly interface.

The project is designed around a thin-client architecture where the CLI is responsible for user interaction and backend orchestration, while all AI-specific functionality remains encapsulated within the Python backend.

[GitHub Repository](https://github.com/tuhindutta/scratchpad-cli)

---

## Features

- 🚀 Terminal-first AI assistant
- 💬 Persistent conversation threads
- 📄 Knowledge ingestion for PDF and TXT documents
- ⚡ Streaming assistant responses
- 🔄 Automatic backend installation and setup
- 🐍 Automatic Python virtual environment creation
- 🔧 Backend lifecycle management
- 🖥 Interactive command shell
- 🧩 Modular and extensible architecture
- 📝 Lightweight Go implementation

---

# Architecture

```
                               +-----------------------+
                               |    Scratchpad CLI     |
                               |      Go + Cobra       |
                               +-----------+-----------+
                                           |
                                   HTTP REST Interface
                                           |
                               +-----------v-----------+
                               |    Python Backend     |
                               |   AI Agent Platform   |
                               +-----------+-----------+
                                           |
             +-----------------------------+-----------------------------+
             |                             |                             |
     Conversation Memory             Knowledge Base                 AI Agents
          (Threads)                      (RAG)                     & Tool Calling
```

The CLI intentionally contains **no AI logic**.

Its responsibilities are limited to:

- User interaction
- Command parsing
- Session persistence
- Backend startup and shutdown
- HTTP communication
- Streaming server responses

The backend is responsible for:

- Agent execution
- LLM orchestration
- Tool calling
- Knowledge retrieval
- Conversation memory
- AI workflows

This separation allows both components to evolve independently while keeping the CLI lightweight.

---

# Repository Structure

```
cmd/
└── scratchpad/
    └── Application entry point

internals/

├── apiRequests/
│   HTTP client layer
│
├── backend/
│   Backend lifecycle management
│
├── cli/
│   Cobra commands and interactive shell
│
└── envSetup/
    Environment bootstrap
        ├── Python verification
        ├── Backend download
        ├── Virtual environment creation
        └── Dependency installation

README.md
LICENSE
go.mod
```

---

# Installation

## Requirements

- Go 1.25+
- Python 3.12.12 or later
- Git

---

## Clone

```bash
git clone https://github.com/tuhindutta/scratchpad-cli.git

cd scratchpad-cli
```

---

## Build

```bash
go build -o scratchpad
```

or

```bash
go install
```

---

# First Run

Start the backend.

```bash
scratchpad start
```

On first execution Scratchpad automatically performs the following:

```
Verify Python installation
        │
        ▼
Download backend repository
        │
        ▼
Create Python virtual environment
        │
        ▼
Install backend dependencies
        │
        ▼
Launch backend server
        │
        ▼
Ready for requests
```

No manual backend installation is required.

---

# Interactive Shell

Launch the built-in shell.

```bash
scratchpad shell
```

Example session

```text
scratchpad>

start

chat Explain Retrieval Augmented Generation.

thread list

knowledge

exit
```

---

# Available Commands

| Command | Description |
|----------|-------------|
| `start` | Download (first run) and start backend |
| `stop` | Shutdown backend |
| `chat` | Send prompt to assistant |
| `knowledge` | Import PDF/TXT documents |
| `thread list` | List conversation threads |
| `thread delete` | Delete a conversation thread |
| `thread clear` | Delete all conversations |
| `credentials get` | Show current credentials |
| `credentials set` | Set user/thread IDs |
| `shell` | Launch interactive shell |
| `version` | Show application version |

---

# Session Management

Scratchpad stores the following metadata locally:

- User ID
- Thread ID
- Backend Port

This enables persistent conversations across CLI sessions.

If credentials are unavailable, Scratchpad automatically creates a temporary session.

---

# Streaming Responses

Assistant responses are streamed incrementally as execution progresses.

Example

```text
===== ASSISTANT =====

[Planner]

Node: Planner

Type: assistant

Content:
Planning task...

[Retriever]

Node: Retriever

Type: tool

Tool Name: search_documents

Content:
Searching knowledge base...

[Assistant]

Node: Assistant

Type: assistant

Content:
Final response...
```

Streaming provides visibility into intermediate execution instead of waiting for the entire workflow to complete.

---

# Knowledge Ingestion

Scratchpad supports importing external documents into the backend knowledge base.

Supported formats

- PDF
- TXT

Once ingested, documents become available to the backend retrieval pipeline for future conversations.

---

# Design Philosophy

Scratchpad was built around a few core principles.

- Terminal-first workflow
- Minimal setup
- Lightweight client
- Persistent AI sessions
- Streaming interactions
- Clean separation of responsibilities
- Extensible architecture
- Simple developer experience

---

# Why Go?

Go was chosen for the CLI because it provides:

- Small self-contained binaries
- Excellent concurrency primitives
- Fast startup time
- Strong standard library
- Easy cross-platform compilation

---

# Why Python?

Python powers the backend because of its rich AI ecosystem.

It provides seamless integration with:

- LLM frameworks
- Agent orchestration
- Retrieval pipelines
- Machine learning libraries
- Tool integrations

Keeping AI functionality inside the backend allows the CLI to remain lightweight and language-agnostic.

---

# Why HTTP?

Communication between the CLI and backend uses HTTP.

Advantages include:

- Simplicity
- Easy debugging
- Backend language independence
- Streaming support
- Future compatibility with remote deployments

---

# Future Work

Planned improvements include

- Multiple model providers
- MCP integration
- Plugin architecture
- Configuration management
- Cross-platform release binaries
- Docker deployment
- Enhanced logging
- Authentication improvements
- Shell autocompletion

---

# License

This project is licensed under the MIT License.

See the LICENSE file for details.

---

# Author

**Tuhin Kumar Dutta**

Portfolio Website: https://www.tuhindutta.com/
