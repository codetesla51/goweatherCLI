s# Weather CLI 🌤️

A simple command-line interface (CLI) application written in Go that provides current weather information for any city around the world.

## Features

- 🌍 Get real-time weather data for any city
- 🌡️ Display temperature in Celsius
- 💧 Show humidity percentage
- 🌬️ Report wind speed and direction
- ☁️ Indicate cloud cover percentage
- 🌍 Display atmospheric pressure
- 🎨 Colorful, easy-to-read output



## Requirements

- Go 1.24 or higher
- OpenWeatherMap API key

## Installation

1. Clone this repository:
   ```
   https://github.com/codetesla51/goweatherCLI.git
   cd goweatherCLI
   ```

2. Create a `.env` file in the project root directory with your OpenWeatherMap API key:
   ```
   API_KEY="your_api_key_here"
   ```

3. Install dependencies:
   ```
   go mod tidy
   ```

4. Build the application:
   ```
   go build
   ```

## Usage

Simply run the application and enter the name of the city when prompted:

```
./weathercli
```

You'll see a welcome banner and be prompted to enter a city name. After entering a valid city name, the application will display detailed weather information including:

- Location (city and country)
- Current temperature and "feels like" temperature
- Humidity level
- Wind speed and direction
- Cloud cover percentage
- Atmospheric pressure
- Current weather condition and description

## Dependencies

This project uses the following external Go packages:

- [github.com/fatih/color](https://github.com/fatih/color) - For colorized terminal output
- [github.com/joho/godotenv](https://github.com/joho/godotenv) - For loading environment variables

## API

This application uses the [OpenWeatherMap API](https://openweathermap.org/api) to fetch weather data. You'll need to sign up for a free API key to use this application.



## Acknowledgments

- Thanks to OpenWeatherMap for providing the weather data API
- Inspiration from various weather applications and CLI tools