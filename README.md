# smith
The Agent

This program is personal and experimental, so please don't use it in production.

- smith watches a specified file, and detects changes on the file
- if file is renamed, smith reopens a file with the same name and keep watching
- when smith detects new line in the file, it sends that data to specified server


Huge thanks to
--------------

- https://github.com/go-fsnotify/fsnotify/blob/master/example_test.go
- https://github.com/fujiwara/fluent-agent-hydra
