Given a root directory it traverses each of the sub directories in parallel and then displays the tree of directories in both text and JSON. It just displays directories, not files because... reasons.

Uses:
- Goroutines
- WaitGroups
- A mutex to update the "children" collection
- JSON marshalling


