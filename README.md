# Shortify

Shortify is a fast and reliable URL shortener written in Go, using Redis as a database and built with the Fiber web framework.

![Shortify](https://github.com/Pradumnasaraf/Shortify/assets/51878265/2c2686e8-82aa-481c-9d4a-07c2d3711398)

## 📦 Installation

Before you begin, make sure you have Docker and Docker Compose installed on your machine. Then, follow these steps to install and run Shortify:

1. Clone this repository to your local machine.
2. Create a `.env` file inside the `api` folder by copying the `.env.example` file. You can do this by running the following command from the root directory of the project:
   ```bash
   cp api/.env.example api/.env
   ```
3. Start the project by running the following command from the root directory of the project:
   ```bash
   docker compose up
   ```

## 📝 Usage

To create a short URL, send a `POST` request to the `/api/v1` endpoint with the following JSON body:

```json
{
  "url": "https://pradumnasaraf.dev",
  "short": "pradumna"
}
```

The `short` field is optional. If you don't provide it, a random string will be generated. In response, you will receive the following JSON body:

```json
{
  "url": "https://pradumnasaraf.dev",
  "short": "localhost:8080/pradumna",
  "expiry": 24,
  "rate_limit": 9,
  "rate_limit_reset": 30
}
```

You can modify the rate limit in the `.env` file. To test if the short URL works, enter it into your browser's address bar. For example, if the short URL is `http://localhost:8080/pradumna`, you will be redirected to the original URL.

## 📜 License

This project is licensed under the Apache-2.0 license - see the [LICENSE](LICENSE) file for details.

## 🛡 Security

If you discover a security vulnerability within this project, please check the [SECURITY](SECURITY.md) file for more information.
