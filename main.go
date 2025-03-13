package main

import (
	"fmt"
	"io"
	"net/http"
	"encoding/json"
	"github.com/fatih/color"
	"github.com/joho/godotenv"
	"os"
)


func clearScreen() {
    fmt.Print("\033[H\033[2J")
}

func displayWelcomeBanner() {
      clearScreen()

	green := color.New(color.FgGreen).SprintFunc()

	fmt.Println(green("/* :::::::::::::::::::::::::::: */"))
	fmt.Println(green("/* :::::::::::::::::::::::::::: */"))
	fmt.Println(green("/* ::                        :: */"))
	fmt.Println(green("/* ::    █████████           :: */"))
	fmt.Println(green("/* ::   ███░░░░░███          :: */"))
	fmt.Println(green("/* ::  ███     ░░░   ██████  :: */"))
	fmt.Println(green("/* :: ░███          ███░░███ :: */"))
	fmt.Println(green("/* :: ░███    █████░███ ░███ :: */"))
	fmt.Println(green("/* :: ░░███  ░░███ ░███ ░███ :: */"))
	fmt.Println(green("/* ::  ░░█████████ ░░██████  :: */"))
	fmt.Println(green("/* ::   ░░░░░░░░░   ░░░░░░   :: */"))
	fmt.Println(green("/* ::                        :: */"))
	fmt.Println(green("/* :::::::::::::::::::::::::::: */"))
	fmt.Println(green("/* :::::::::::::::::::::::::::: */"))
}

// Prompt user for city input
func showMessage() string {
	displayWelcomeBanner()
	yellow := color.New(color.FgYellow).SprintFunc()

	fmt.Println(yellow("\n🌍 Enter city name:"))
	var cityName string
	fmt.Scanf("%s", &cityName)

	return cityName
}

type Sys struct {
	Country string `json:"country"`
}

type Main struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	Humidity  int     `json:"humidity"`
	Pressure  int     `json:"pressure"`
}

type Weather struct {
	Main        string `json:"main"`
	Description string `json:"description"`
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg   int     `json:"deg"`
}

type Clouds struct {
	All int `json:"all"`
}

type WeatherResponse struct {
	Name    string    `json:"name"`
	Sys     Sys       `json:"sys"`
	Main    Main      `json:"main"`
	Weather []Weather `json:"weather"`
	Wind    Wind      `json:"wind"`
	Clouds  Clouds    `json:"clouds"`
}
func makeRequest(cityName string) (*WeatherResponse, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("❌ Failed to load .env file")
	}

	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("❌ API_KEY is missing in .env file")
	}

	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", cityName, apiKey)
	response, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("❌ Error making request: %v", err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("❌ Error reading response: %v", err)
	}

	var weather WeatherResponse
	err = json.Unmarshal(body, &weather)
	if err != nil {
		return nil, fmt.Errorf("❌ Error parsing JSON: %v", err)
	}

	return &weather, nil
}

// Main function
func main() {
	cityName := showMessage()

	red := color.New(color.FgRed).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()
	blue := color.New(color.FgBlue).SprintFunc()
	cyan := color.New(color.FgCyan).SprintFunc()

	weather, err := makeRequest(cityName)
	if err != nil {
		fmt.Println(red("\n❌ Failed to fetch weather:"), err)
		return
	}

	fmt.Println(cyan("\n================ WEATHER REPORT ================\n"))
	fmt.Printf("📍 Location: %s, %s\n", blue(weather.Name), green(weather.Sys.Country))
	fmt.Printf("🌡️ Temperature: %.2f°C (Feels like %.2f°C)\n", weather.Main.Temp, weather.Main.FeelsLike)
	fmt.Printf("💧 Humidity: %d%%\n", weather.Main.Humidity)
	fmt.Printf("🌬️ Wind: %.2f m/s at %d°\n", weather.Wind.Speed, weather.Wind.Deg)
	fmt.Printf("☁️ Cloud Cover: %d%%\n", weather.Clouds.All)
	fmt.Printf("🌍 Pressure: %d hPa\n", weather.Main.Pressure)
	fmt.Printf("🌤️ Condition: %s - %s\n", red(weather.Weather[0].Main), yellow(weather.Weather[0].Description))
	fmt.Println(cyan("\n================================================\n"))
}