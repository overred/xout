package xformat

import (
	"fmt"
	"strings"
	"time"

	"github.com/overred/xout/xfields"
	"github.com/overred/xout/xlevel"
	"gopkg.in/gookit/color.v1"
)

// LogrusText makes logrus-like text format.
func LogrusText() Formatter {
	start := time.Now()
	return func(level xlevel.Level, input string, fields xfields.Fields) string {
		name := strings.ToUpper(fmt.Sprintf("%.4s", level.Higher().String()))

		clr := map[xlevel.Level]color.Color{
			xlevel.Trace: color.FgGray,
			xlevel.Debug: color.FgGray,
			xlevel.Info:  color.FgBlue,
			xlevel.Warn:  color.FgYellow,
			xlevel.Error: color.FgRed,
			xlevel.Fatal: color.FgRed,
			xlevel.Panic: color.FgRed,
		}[level]
		if clr == 0 {
			clr = color.FgWhite
		}

		values := strings.Builder{}
		for _, field := range fields {
			values.WriteString(fmt.Sprintf("%s=%v ", clr.Render(field.Name), field.String()))
		}

		format := fmt.Sprintf(
			"%-4s[%04d] %-44s %s\n",
			clr.Render(name),
			int(time.Since(start).Seconds()),
			input,
			values.String(),
		)
		return format
	}
}
