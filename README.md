# clihq

Shared Charm-based (lipgloss/huh) terminal styling for devops-chris CLIs (lockr, cloudctx, platformr): a common purple/green/red/yellow/muted palette, semantic message helpers, a themed huh form, and a styled table renderer.

```go
import "github.com/devops-chris/clihq/ui"

fmt.Println(ui.Banner("myapp", "does a thing"))
fmt.Println(ui.Success("Done"))
```

See [ui/styles.go](ui/styles.go), [ui/theme.go](ui/theme.go), [ui/table.go](ui/table.go), and [ui/checks.go](ui/checks.go).
