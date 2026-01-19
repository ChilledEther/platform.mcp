# Quickstart: Core Foundation

## Installation

```bash
<<<<<<< HEAD
go get github.com/modelcontextprotocol/platform.mcp/pkg/scaffold
=======
go get github.com/jjr/platform.mcp/pkg/scaffold
>>>>>>> 001-core-foundation
```

## Usage

```go
<<<<<<< HEAD
import "github.com/modelcontextprotocol/platform.mcp/pkg/scaffold"
=======
import "github.com/jjr/platform.mcp/pkg/scaffold"
>>>>>>> 001-core-foundation

func main() {
    cfg := scaffold.Config{
        ProjectName: "my-app",
<<<<<<< HEAD
        WorkflowType: "standard",
=======
        WorkflowType: "go",
>>>>>>> 001-core-foundation
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
