# Quickstart: Core Foundation

## Installation

```bash
go get github.com/jjr/platform.mcp/pkg/scaffold
```

## Usage

```go
import "github.com/jjr/platform.mcp/pkg/scaffold"

func main() {
    cfg := scaffold.Config{
        ProjectName: "my-app",
        WorkflowType: "go",
    }
    
    files, err := scaffold.Generate(cfg)
    if err != nil {
        panic(err)
    }
    
    for _, f := range files {
        fmt.Printf("Generated %s\n", f.Path)
    }
}
```
