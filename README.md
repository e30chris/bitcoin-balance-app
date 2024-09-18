# Bitcoin Balance App

This is a Go web application that prompts the user for a Bitcoin address and displays the balance of that Bitcoin address using the mempool.space API.

## Setup

1. Clone the repository:

   ```bash
   git clone <repository-url>
   ```

2. Install the project dependencies:

   ```bash
   go mod download
   ```

3. Build the application:

   ```bash
   go build ./cmd
   ```

## Usage

1. Start the application:

   ```bash
   ./cmd
   ```

2. Open your web browser and navigate to `http://localhost:8080`.

3. Enter a Bitcoin address in the provided input field and click the "Get Balance" button.

4. The balance of the Bitcoin address will be displayed on the page.

## Dependencies

- [mempool.space API](https://mempool.space/api-docs/)
- [Go](https://golang.org/)

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, please open an issue or submit a pull request.

## License

This project is licensed under the [MIT License](LICENSE).
```

Please note that you need to replace `<repository-url>` with the actual URL of your repository.