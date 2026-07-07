package ui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

var (
	checkPassStyle = lipgloss.NewStyle().Bold(true).Foreground(ColorGreen)
	checkFailStyle = lipgloss.NewStyle().Bold(true).Foreground(ColorRed)
	checkWarnStyle = lipgloss.NewStyle().Bold(true).Foreground(ColorYellow)
	checkHintStyle = lipgloss.NewStyle().Foreground(ColorMuted).PaddingLeft(6)
)

// CheckPass renders a passing "doctor"-style check line.
func CheckPass(msg string) string {
	return fmt.Sprintf("  %s  %s", checkPassStyle.Render("✓"), msg)
}

// CheckFail renders a failing "doctor"-style check line, with an optional hint.
func CheckFail(msg, hint string) string {
	s := fmt.Sprintf("  %s  %s", checkFailStyle.Render("✗"), msg)
	if hint != "" {
		s += "\n" + checkHintStyle.Render(hint)
	}
	return s
}

// CheckWarn renders a warning "doctor"-style check line, with an optional hint.
func CheckWarn(msg, hint string) string {
	s := fmt.Sprintf("  %s  %s", checkWarnStyle.Render("⚠"), msg)
	if hint != "" {
		s += "\n" + checkHintStyle.Render(hint)
	}
	return s
}

// CheckSection renders an indented section header for use inside check lists.
func CheckSection(name string) string {
	return fmt.Sprintf("\n  %s\n\n", sectionStyle.Render(name))
}
