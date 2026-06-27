# timetrack

![Timetrack Dashboard](https://media.discordapp.net/attachments/869362798034059314/1520256613175267458/image.png?ex=6a4088cd&is=6a3f374d&hm=a60afedcb61c6b6ed36111bce4e5a65eb914c84877f842e44011c0be4abfc234&=&format=webp&quality=lossless&width=1872&height=719)

A minimalist time-tracking tool built as a personal hands-on experiment to explore Go backend development, concurrent routines, and native HTTP routing.

*Note: This repository is a proof-of-concept and a learning playground, not a production-ready utility.*

## What it does

- **CLI Tracking:** Start and stop time blocks for different projects directly from the terminal.
- **Live Clock:** A real-time status counter powered by Go goroutines and tickers.
- **JSON Persistence:** Saves all session history into a local, human-readable JSON file.
- **Brutalist Dashboard:** A lightweight, native HTTP server that serves a basic JSON API and renders a rustic, monospace HTML bar chart.

## Commands

```bash
# Start a tracking session
go run main.go start <project-name>

# View live elapsed time (Ctrl+C to exit view)
go run main.go status

# Stop the current session
go run main.go stop

# View a text summary report
go run main.go report

# Launch the web dashboard at http://localhost:8080
go run main.go serve