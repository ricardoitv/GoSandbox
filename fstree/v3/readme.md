Like [v2](../v2/readme.md) but:
- Will aim to use a bounded worker pool/buffered channel semaphore pattern when traversing the tree of dirs
- Remove the use of `log.Fatal`
- Have `Run()` return errors
- Better separation of concerns
- Add metrics (Add a channel for “nodes processed” - can print progress)
- Support CLI args (like `--path /some/dir --output json --max-workers 10`)
