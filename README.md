# go-utils

## Release

1. Update Go version in `go.mod` file.

2. Update packages used in the application.

```bash
go get -u ./util ./monitoring
```

3. Run go mod tidy, which removes any dependencies the module might have accumulated that are no longer necessary.

```bash
go mod tidy
```

4. Run go test ./... a final time to make sure everything is working.

```bash
go test ./...
```

5. Tag the project with a new version number using the git tag command.

```bash
git add .
git commit -m "util: changes for v0.1.0"
git tag v0.1.0
```

6. Push the new tag to the origin repository.

```bash
git push --atomic origin main v0.1.0
```

7. Make the module available by running the go list command to prompt Go to update its index of modules with information about the module you’re publishing.

```bash
GOPROXY=proxy.golang.org go list -m github.com/antonputra/go-utils@v0.1.0
```
