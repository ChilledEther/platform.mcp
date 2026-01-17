# Unit Tests: Core Foundation

## Package: `pkg/scaffold`

| Test Case | Description | File |
|-----------|-------------|------|
| `TestGenerate_Basic` | Verifies basic file generation for Go projects. | `scaffold_test.go` |
| `TestGenerate_Table` | Table-driven tests for various project types and Docker options. | `scaffold_test.go` |
| `TestValidateConfig` | Verifies Config validation logic and edge cases. | `scaffold_test.go` |
| `TestGenerate_NoSideEffects` | Ensures the generator is a pure function (no disk I/O). | `scaffold_test.go` |
| `TestTemplates_Embedded` | Verifies that embedded templates are correctly loaded and parsed. | `templates_test.go` |
