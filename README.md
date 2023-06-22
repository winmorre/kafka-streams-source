## Source Repo
Simple service connecting to [finnhub.io](https://finnhub.io) trade data. The data is emitted into kafka where consumers can pull and process.

### How to run this code
- Register for an API Key at https://finnhub.io
- Add the API key to the [basic.go](services/basic.go) 
- Download packages using ```go mod download ```
- Run the code using ```go run main.go```