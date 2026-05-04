# Chunkr.ai Go SDK

Go SDK for [Chunkr.ai](https://chunkr.ai/) — document layout analysis and chunking API.

## Quick Start

```bash
go get github.com/gitslim/chunkr-ai-sdk/v2/chunkrai@latest
```

```go
import (
    client "github.com/gitslim/chunkr-ai-sdk/v2/chunkrai/client"
    option "github.com/gitslim/chunkr-ai-sdk/v2/chunkrai/option"
    chunkrai "github.com/gitslim/chunkr-ai-sdk/v2/chunkrai"
)

c := client.NewClient(
    option.WithToken("<YOUR_API_KEY>"),
)
task, err := c.Task.CreateTaskRoute(ctx, &chunkrai.CreateForm{
    File: "https://example.com/document.pdf",
})
```

See [`chunkrai/README.md`](chunkrai/README.md) for full documentation.

## Development

```bash
# Regenerate SDK
cd fern && npx fern-api generate --local
```

## License

MIT
