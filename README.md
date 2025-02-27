# Fingerpori RSS Feed Generator

A simple Go application that fetches Fingerpori comics from Helsingin Sanomat's API and generates an RSS feed with the comic images.

## Features

- Fetches the latest Fingerpori comics from HS.fi
- Generates an RSS feed with the comic images as content
- Replaces placeholder values in image URLs with appropriate values (WIDTH → 1440, EXT → jpg)
- Saves the RSS feed to a file and outputs it to stdout

## Usage

```bash
# Run the application
go run main.go

# The RSS feed will be saved to fingerpori.xml
```

## Requirements

- Go 1.16 or later
- Internet connection to fetch the comics from HS.fi

## Dependencies

- [github.com/gorilla/feeds](https://github.com/gorilla/feeds) - For RSS feed generation

## License

This project is for personal use only. Fingerpori comics are copyrighted by Pertti Jarla and published by Helsingin Sanomat.
