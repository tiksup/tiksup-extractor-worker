# Tiksup-Stractor-Worker

**Tiksup-Extractor-Worker** is a **gRPC-based worker service** designed to extract data from **Redis** and **MongoDB** upon receiving events. It processes incoming gRPC requests, retrieves the required data, and forwards it to another gRPC server.

---

## Features

- **gRPC Server**: Receives events and triggers data extraction.
- **Data Sources**:
  - **Redis**: Fetches key-value data.
  - **MongoDB**: Retrieves structured data from collections.
- **Event-Driven Architecture**: Processes events with a specific flag (`next: true`) to start the extraction process.
- **Data Forwarding**: Sends extracted data to a target gRPC server for further processing.

---

## Workflow

1. **Receive Event**: 
   - The gRPC server receives an event.
   - The event payload must include a `next: true` flag to initiate processing.
   
2. **Extract Data**:
   - Queries Redis for relevant key-value data.
   - Retrieves structured documents from MongoDB.

3. **Forward Data**:
   - Sends the extracted data to another gRPC server.

---

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/tiksup/tiksup-extractor-worker.git
   cd tiksup-extractor-worker
