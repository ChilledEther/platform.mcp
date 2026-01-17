# Data Model: Workflow Configuration

Since this feature implements a CI/CD pipeline, the "Data Model" describes the configuration structure of the GitHub Actions workflow file (`deploy.yml`).

## Workflow Entity

### `deploy.yml`

| Field | Value/Type | Description |
|-------|------------|-------------|
| `name` | String | "Build and Publish Docker Image" |
| `on` | Object | Trigger configuration |
| `jobs` | Object | List of jobs to execute |

### Trigger Configuration (`on`)

```yaml
push:
  branches:
    - main
```

### Job: `build-and-publish`

| Step | Action | Configuration |
|------|--------|---------------|
| `checkout` | `actions/checkout@v4` | Fetches code |
| `login` | `docker/login-action@v3` | Registry: `ghcr.io`<br>Username: `${{ github.actor }}`<br>Password: `${{ secrets.GITHUB_TOKEN }}` |
| `meta` | `docker/metadata-action@v5` | Generates tags (`latest`, `sha-...`) |
| `build-push` | `docker/build-push-action@v5` | Context: `.`<br>Push: `true`<br>Tags: `${{ steps.meta.outputs.tags }}`<br>Labels: `${{ steps.meta.outputs.labels }}` |

## Permissions Model

The workflow requires the following permissions token:

```yaml
permissions:
  contents: read
  packages: write
```
