# Quickstart: Testing the Deployment Workflow

## Prerequisites

1. Access to the GitHub repository.
2. A valid `Dockerfile` in the root of the repository.

## Testing the Workflow

1. **Make a Change**: Create a dummy change or a real feature update.
2. **Commit and Push**:
   ```bash
   git add .
   git commit -m "feat: trigger deployment"
   git push origin main
   ```
3. **Verify Execution**:
   - Go to the **Actions** tab in your GitHub repository.
   - Look for the "Build and Publish Docker Image" workflow run.
   - Ensure it completes successfully (Green checkmark).
4. **Verify Artifact**:
   - Go to the main repository page.
   - Look at the right sidebar under **Packages**.
   - Verify a new Docker package version exists with tags `latest` and `sha-XXXXXXX`.

## Local Validation (Dry Run)

You can build the image locally to ensure the `Dockerfile` is valid before pushing:

```bash
docker build . -t test-image:local
```
