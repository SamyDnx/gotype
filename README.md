# gotype

> A terminal-based typing test written in Go.

`gotype` is a clean and minimal command-line interface (CLI) application designed to test and improve your typing speed and accuracy directly from your terminal.

---

## Compatibility

- **OS:** Linux only (utilizes Linux terminal/raw mode syscalls)
- **Terminals:** Fully tested on terminal emulators such as Foot, Kitty, Alacritty, and standard Linux consoles.

---

## Features

- **Blazing Fast & Lightweight:** Compiled directly to a single binary with zero heavy runtime overhead.
- **Terminal Native:** Built to fit right into your terminal workflow.
- **Real-Time Statistics:** Track your Words Per Minute (WPM) and accuracy on the fly.
- **Clean UI:** Minimalist interface focused purely on the text and your pacing.

---

## Installation

Ensure you have [Go](https://golang.org/) installed (version 1.18+ recommended), then install `gotype` directly via `go install`:

```bash
go install github.com/SamyDnx/gotype@latest
```

Or clone and build from source:

```bash
# Clone the repository
git clone https://github.com/SamyDnx/gotype.git
cd gotype

# Build binary
go build -o gotype main.go
```

---

## Usage

Run the tool simply by executing the binary:

```bash
./gotype

```

*(Optional: Move the binary to your PATH directory, e.g., `sudo mv gotype /usr/local/bin/`, to run it from anywhere.)*

---

## Tech Stack

* **Language:** [Go](https://golang.org/) — for clean structure, and fast execution.
* **Interface:** Terminal UI / ANSI escape sequences.
