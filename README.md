# umimic 🤖

A Go-based API for creating personalized chatbots using OpenRouter, with support for message history and custom personalities.

## Features

- Support for multiple personalities via system messages
- Message history to maintain conversation context
- Integration with OpenRouter (access to various AI models)
- CORS configured for frontend compatibility
- Request validation using go-playground/validator

## Tech Stack

- Go 1.24
- Chi Router (HTTP framework)
- OpenRouter API
- Chi/CORS (middleware)

## Local Development

### Prerequisites

- Go 1.24 or higher
- OpenRouter account (to obtain API key)

### Setup

1. Clone the repository
```bash
git clone https://github.com/xyztavo/umimic.git
cd umimic
```

2. Copy .env.example and fill in your credentials
```bash
cp .env.example .env
```

3. Add your environment variables to .env
```env
OPENROUTER_API_KEY=your_key_here
PORT=3000
ALLOWED_ORIGINS=http://localhost:3000,http://localhost:5173
```

4. Install dependencies
```bash
go mod download
```

5. Run the server
```bash
go run main.go
```

The server will start at `http://localhost:3000`

## Endpoints

### POST /api/message

Sends a message to the bot with optional history.

**Request Body:**
```json
{
  "message": "your message here",
  "history": [
    {
      "role": "system",
      "content": "bot personality here"
    },
    {
      "role": "user",
      "content": "previous message"
    },
    {
      "role": "assistant",
      "content": "previous response"
    }
  ]
}
```

**Response:**
```json
{
  "reply": "bot response"
}
```

### GET /api/healthz

Checks if the API is running.

**Response:**
```json
{
  "status": "ok"
}
```

## Deployment

### Railway (Free Tier)

Railway provides 500 hours per month on the free plan, which is more than sufficient.

1. Create an account at [railway.app](https://railway.app)

2. Connect your GitHub repository

3. Click "New Project" > "Deploy from GitHub repo"

4. Select the `umimic` repository

5. Add environment variables:
   - `OPENROUTER_API_KEY`
   - `PORT` (Railway will set this automatically, but you can force 3000)
   - `ALLOWED_ORIGINS` (add your frontend URL)

6. Railway will automatically detect the Go project and build it

7. After deployment, you will receive a generated URL like `umimic-production.up.railway.app`

### DisCloud (Free Tier)

DisCloud offers a free plan with 256MB RAM and 1GB storage.

1. Create an account at [discloud.app](https://discloud.app)

2. Create a `discloud.config` file in the project root:
```
TYPE=site
MAIN=main.go
RAM=256
AUTORESTART=true
VERSION=recommended
APT=tools
```

3. Upload the compressed project (zip) through the DisCloud panel

4. Add environment variables in the panel:
   - `OPENROUTER_API_KEY`
   - `PORT=8080` (DisCloud uses port 8080)
   - `ALLOWED_ORIGINS`

5. DisCloud will automatically build and deploy the application

6. Access via the provided URL like `your-app.discloud.app`

## Project Structure

```
.
├── config/          # Configuration and environment variables
├── mimic/           # Chatbot logic
├── models/          # Request/response structs
├── openrouter/      # OpenRouter API client
├── utils/           # Utilities (validation, etc.)
├── main.go          # API entrypoint
└── go.mod           # Dependencies
```

## How It Works

1. Frontend sends message + history (including personality as system message)
2. Backend appends the current message to history
3. Makes request to OpenRouter with full context
4. Returns the model's response

Personalities are managed on the frontend and sent as the first message in history with role "system".

## Troubleshooting

**CORS error:** Add your frontend origin to `ALLOWED_ORIGINS`

**API timeout:** OpenRouter can sometimes be slow; consider increasing timeout if necessary

**Rate limit:** OpenRouter has per-minute limits; monitor your usage

## Contributing

Feel free to open issues or pull requests.

## License

Made with ❤️ by ustav