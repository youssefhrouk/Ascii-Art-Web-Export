# ascii-art-web-dockerize

## Objectives
The goal of this project is to enhance the original Ascii-Art-Web by making the site:
- More appealing, interactive, and intuitive
- More user-friendly
- Provide more feedback to the user
- **Responsive**: Ensuring the site works well on various devices and screen sizes

## Instructions
To achieve the objectives, the project must meet the following requirements:
- The project must contain CSS to style the website.
- The code must follow good coding practices.
- The text must be legible regardless of the colors used.
- The website must be consistent, responsive, and interactive.

We recommend researching the principles of good website design to guide your enhancements.

## Features
- **Convert input text to ASCII art**: Users can input text and choose from multiple banner styles (standard, shadow, and thinkertoy) to generate ASCII art.
- **Improved User Interface**: Enhanced UI with CSS for a more appealing and user-friendly experience.
- **Responsive Design**: Ensures the website is usable on various devices and screen sizes.
- **Interactive Elements**: Provides feedback to users through interactive elements.

## Technologies Used
- **Backend**: Go (Golang)
- **Frontend**: HTML, CSS
- **HTTP Server**: Go's `net/http` package
- **Templates**: Go's `html/template` package

<!DOCTYPE html>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<title>Ascii-Art-Web Project Structure</title>
<style>
    pre {
        font-family: monospace;
        font-size: 14px;
    }
</style>
<body>
    <h1>Ascii-Art-Web Project Structure</h1>
    <pre>
ascii-art-stylize/
├── main.go
├── Handlers/
│   ├── AsciiHandler.go
│   │── ErrorHandler.go
│   └── IndexHandler.go
├── ascii/
│   ├── PrintAndSplit.go
│   ├── Converter.go
│   ├── FileReader.go
│   ├── GetBannerPath.go
│   ├── HandleNewLines.go
│   ├── Printer.go
│   └── ValidateInput.go
└── templates/
    ├── index.html
    ├── error.html
    └── ascii-art.html
└── stylize/
    └── index.css
    </pre>
</body>

## Usage
1. Run the application:
2. Open a web browser and visit `http://localhost:8080`
3. Enter text in the input field, select a banner style, and click "Generate" to create ASCII art

## API Endpoints
- GET `/`: Displays the main page with the input form
- POST `/ascii-art`: Generates and displays the ASCII art

## Error Handling
The application includes error handling for various scenarios:
- 400 Bad Request: For invalid form inputs
- 404 Not Found: For non-existent pages
- 405 Method Not Allowed: For incorrect HTTP methods
- 500 Internal Server Error: For server-side issues

## Authors
- @asoudri
- @yhrouk