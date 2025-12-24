Like the [v1](../v1/readme.md) but this implementation uses a different approach.

Changes:
- Implementation keeps pretty much the same
- Mutex is removed
- There will be a single "writer" Goroutine and will update the "children" collection
- Communication between goroutines is achieved via channels

Uses:
- Goroutines
- WaitGroups
- Channels
- JSON marshalling
