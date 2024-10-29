# Focus Guard

Focus Guard is a web application designed to help users create and manage personalized schedules for blocking distracting websites, enhancing productivity across various platforms.

## Features
- Create and manage website blocking schedules.
- User-friendly interface for easy navigation.
- Cross-platform support with executable binaries for Windows, Linux, and macOS.

## Installation

### Pre-built Binaries
1. Download the appropriate binary executable for your operating system from the [Releases](https://github.com/getaseww/focus-guard/releases) page.
2. Make the binary executable (for Linux and macOS):
   ```bash
   chmod +x focus-guard
   ```
3. Move the binary to your preferred directory (optional):
   ```bash
   mv focus-guard /usr/local/bin/
   ```

### Building from Source
1. Ensure you have Go installed. If not, download and install it from [the official Go website](https://golang.org/dl/).
2. Clone the repository:
   ```bash
   git clone https://github.com/getaseww/focus-guard.git
   cd focus-guard
   ```
3. Install the necessary dependencies:
   ```bash
   go mod tidy
   ```
4. Build the application:
   ```bash
   go build -o focus-guard
   ```

## Usage
1. Run the application:
   - For Windows:
     ```bash
     focus-guard.exe
     ```
   - For Linux/macOS:
     ```bash
     ./focus-guard
     ```
2. Access the web interface by navigating to `http://localhost:8080` in your web browser.
3. Follow the on-screen instructions to create and manage your website blocking schedules.

## Contributing
Contributions are welcome! Please feel free to submit a pull request or open an issue for any enhancements, bug fixes, or suggestions.

## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
```

### Instructions
- Make sure to replace `yourusername` with your actual GitHub username in the clone and download links.
- Update the repository link and any other specifics as your project evolves.
