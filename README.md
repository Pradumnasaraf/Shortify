# Shortify

Shortify is a URL shortener written in Go. It uses Redis as a database and is built with the Fiber.

## 📦 Installation

- Prerequisite: Docker and Docker Compose installed

To run to project. after cloning the repository. First create a `.env` file inside the `api` folder by copying the `.env.example` file. Also you can run the following command to create the `.env` file (from the root).

```bash
cp api/.env.example api/.env
```

Then run the following command to start the project.

```bash
docker compose up
```

## 📝 Usage

To create a short URL send a `POST` request to the `/api/v1` endpoint with the following body.

```json
{
  "url": "https://pradumnasaraf.dev",
  "short": "pradumna" // this is optional. If not provided a random string will be generated
}
```

In the response you will receive the recive the following body. If we don't provide the `short` field in the request body then the `short` field in the response will be a random string. Also, we can modify the expiry time and rate limit in the `.env` file.

```json
{
  "url": "https://pradumnasaraf.dev",
  "short": "localhost:8080/pradumna",
  "expiry": 24,
  "rate_limit": 9,
  "rate_limit_reset": 30
}
```

Now to check if the short URL received in the response is working, head over to broswer and enter URL. For example, if the short URL is `http://localhost:3000/pradumna` then enter `http://localhost:3000/pradumna` in the browser. You will be redirected to the original URL.

## 📜 License

This project is licensed under the Apache-2.0 license - see the [LICENSE](LICENSE) file for details.

## 🛡 Security

If you discover a security vulnerability within this project, please check the [SECURITY](SECURITY.md) for more information.
