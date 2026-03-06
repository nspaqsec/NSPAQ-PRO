# NSPAQ-PRO v1.0
**Advanced Multi-Layer (L4/L7) Network Stress Testing Tool**

NSPAQ-PRO is a high-performance network stress testing utility developed in Go. It is designed to evaluate network infrastructure and firewall resilience by simulating various traffic patterns. The tool supports modern protocols and provides a modular architecture for scalability.

---

## Core Features
* **Dual-Layer Support:** Capability to perform Layer 7 (HTTP GET, POST, HEAD) and Layer 4 (UDP, TCP Flood) testing.
* **HTTP/2 Integration:** Optimized HTTP/2 stack to bypass traditional connection limits and simulate modern browser traffic.
* **Automated Proxy Management:** Integrated proxy scraper with dynamic rotation logic to distribute requests across multiple origins.
* **Dynamic Header Generation:** Randomized header sets including User-Agent rotation for realistic client simulation.
* **Real-time Statistics:** Interactive terminal dashboard displaying Requests Per Second (RPS), total successes, and failure counts.
* **Cross-Platform Compatibility:** Native support for Windows, Linux, and macOS.

---

## Installation and Setup:

### 1. Windows Systems
1. Ensure the Go programming language (1.20+) is installed on your system.
2. Navigate to the project directory and install dependencies:
   ```bash
   go mod tidy
3. Run it:
    go run .

### 2. Kali Linux / Ubuntu Systems
sudo apt update && sudo apt install golang git -y
git clone https://github.com/nspaqsec/NSPAQ-PRO/
go mod tidy
go build -o nspaq
./nspaq

## Operational Guidelines
* **Layer 7 Testing:** Use full URL formats (e.g., http://example.com) for application-level tests.

* **Layer 4 Testing:** Use IP:Port format (e.g., 1.1.1.1:80) for transport-layer flood tests.

* **Concurrency:** Adjust the thread count based on your system resources to optimize the RPS output.

## Legal Disclaimer:

* **This tool is developed strictly for educational purposes and authorized security testing. The use of this software against unauthorized systems is illegal and may result in criminal charges. The developer assumes no liability for any misuse or damage caused by this program. Users are solely responsible for complying with local and international laws.**

Maintained by nspaqsec 