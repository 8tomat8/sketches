# URL Screenshoter

A simple tool to take screenshots of multiple URLs using Playwright.

## Setup

### On Windows:

1. Install dependencies:
```bash
pip install -r requirements.txt
```

2. Install Playwright browsers:
```bash
python -m playwright install
```

## Usage

### Take screenshots from a list of URLs:
```bash
python main.py --urls https://www.google.com https://www.github.com
```

### Take screenshots from a file containing URLs (one per line):
```bash
python main.py --file urls.txt
```

### Specify a custom output directory:
```bash
python main.py --file urls.txt --output my_screenshots
```

## Features

- Takes full-page screenshots
- Uses page titles as filenames
- Processes URLs in parallel using a single browser with multiple tabs
- Handles errors gracefully 