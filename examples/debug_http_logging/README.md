# Debug HTTP Logging Example

This example demonstrates how to enable debug logging in the OpenAI client to see the full HTTP request details including headers and request body.

## Features

The enhanced debug logging provides:

- **Full HTTP Request Details**: Method, URL, headers, and complete request body
- **Masked API Key**: API key is safely masked in logs (shows first 4 and last 4 characters)
- **Request Body**: Complete JSON payload sent to OpenAI API
- **Operation Context**: Which method triggered the request (Generate, Chat, GenerateWithTools, etc.)

## Usage

1. Set your OpenAI API key:
```bash
export OPENAI_API_KEY="your-api-key-here"
```

2. Run the example:
```bash
go run main.go
```

## Debug Output Example

When debug logging is enabled, you'll see output like:

```
DEBUG: Full HTTP Request Details
{
  "operation": "Generate",
  "method": "POST", 
  "url": "https://api.openai.com/v1/chat/completions",
  "headers": {
    "Content-Type": "application/json",
    "Authorization": "Bearer sk-1***xyz",
    "User-Agent": "agent-sdk-go"
  },
  "body": {
    "model": "gpt-4o-mini",
    "messages": [
      {
        "role": "system",
        "content": "You are a helpful geography assistant."
      },
      {
        "role": "user", 
        "content": "What is the capital of France?"
      }
    ],
    "temperature": 0.7
  }
}
```

## Security Notes

- The API key is automatically masked in debug logs for security
- Only the first 4 and last 4 characters of the API key are shown
- All other request details are logged in full for debugging purposes

## Use Cases

This debug logging is useful for:

- **API Debugging**: See exactly what's being sent to OpenAI
- **Request Optimization**: Analyze request structure and parameters
- **Integration Testing**: Verify request formatting and headers
- **Troubleshooting**: Diagnose API communication issues
- **Performance Analysis**: Track request complexity and size
