# Fingerpori RSS Feed Generator

A simple Go application that fetches Fingerpori comics from Helsingin Sanomat's API and generates an RSS feed with the comic images.

## Features

- Fetches the latest Fingerpori comics from HS.fi
- Generates an RSS feed with the comic images as content
- Replaces placeholder values in image URLs with appropriate values (WIDTH → 1440, EXT → jpg)
- Saves the RSS feed to a file and outputs it to stdout

## Usage

```bash
# Run the application with default settings (saves to fingerpori.xml in current directory)
go run main.go

# Specify an output directory
go run main.go -outdir /path/to/output/directory

# Build and run the binary
go build -o fingerpori-rss
./fingerpori-rss -outdir /path/to/output/directory
```

## Command-line Options

- `-outdir`: Directory where the RSS feed XML file will be created (default: current directory)

## License

This project is for personal use only. Fingerpori comics are copyrighted by Pertti Jarla and published by Helsingin Sanomat.
