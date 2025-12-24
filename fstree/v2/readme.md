Like the [v1](../v1/readme.md) but this implementation uses a different approach.

Changes:
- Implementation keeps pretty much the same code
- Mutex is removed
- There will be a single "writer" Goroutine and it will update the "children" collection (single writer)
- Communication between goroutines is achieved via a channel

Uses:
- Goroutines
- WaitGroups
- Channels
- JSON marshalling
