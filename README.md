# Strip TOML/YAML Front Matter

Little utility to strip TOML/YAML front matter from STDIN.

Created to be used in combination with [Marked 2].

Usage:

```sh
go build strip-front-matter.go
chmod +x strip-front-matter
cat markdown_file.md | ./strip-front-matter
```

[Marked 2]: https://marked2app.com/
