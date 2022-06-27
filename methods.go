package xout

import (
	"fmt"

	"github.com/overred/xout/xlevel"
)

// Print acts as a standard fmt.Print.
// Sends data to all targets with xlevel.Text.
func (x Logger) Print(a ...interface{}) (int, error) {
	return fmt.Fprint(x.Writer(xlevel.Text), a...)
}

// Print acts as a standard fmt.Print.
// Sends data to all targets with xlevel.Text.
func Print(a ...interface{}) (int, error) {
	return fmt.Fprint(Default.Writer(xlevel.Text), a...)
}

// Println acts as a standard fmt.Println.
// Sends data to all targets with xlevel.Text.
func (x Logger) Println(a ...interface{}) (int, error) {
	return fmt.Fprintln(x.Writer(xlevel.Text), a...)
}

// Println acts as a standard fmt.Println.
// Sends data to all targets with xlevel.Text.
func Println(a ...interface{}) (int, error) {
	return fmt.Fprintln(Default.Writer(xlevel.Text), a...)
}

// Printf acts as a standard fmt.Printf.
// Sends data to all targets with xlevel.Text.
func (x Logger) Printf(format string, a ...interface{}) (int, error) {
	return fmt.Fprintf(x.Writer(xlevel.Text), format, a...)
}

// Printf acts as a standard fmt.Printf.
// Sends data to all targets with xlevel.Text.
func Printf(format string, a ...interface{}) (int, error) {
	return fmt.Fprintf(Default.Writer(xlevel.Text), format, a...)
}

// Trace acts as a standard fmt.Print.
// Sends data to all targets with xlevel.Trace.
// Line wrapping is controlled by the targets formatter.
func (x Logger) Trace(a ...interface{}) (int, error) {
	return fmt.Fprint(x.Writer(xlevel.Trace), a...)
}

// Trace acts as a standard fmt.Print.
// Sends data to all targets with xlevel.Trace.
// Line wrapping is controlled by the targets formatter.
func Trace(a ...interface{}) (int, error) {
	return fmt.Fprint(Default.Writer(xlevel.Trace), a...)
}

// Tracef acts as a standard fmt.Printf.
// Sends data to all targets with xlevel.Trace.
// Line wrapping is controlled by the targets formatter.
func (x Logger) Tracef(format string, a ...interface{}) (int, error) {
	return fmt.Fprintf(x.Writer(xlevel.Trace), format, a...)
}

// Tracef acts as a standard fmt.Printf.
// Sends data to all targets with xlevel.Trace.
// Line wrapping is controlled by the targets formatter.
func Tracef(format string, a ...interface{}) (int, error) {
	return fmt.Fprintf(Default.Writer(xlevel.Trace), format, a...)
}

// Debug acts as a standard fmt.Print.
// Sends data to all targets with xlevel.Debug.
// Line wrapping is controlled by the targets formatter.
func (x Logger) Debug(a ...interface{}) (int, error) {
	return fmt.Fprint(x.Writer(xlevel.Debug), a...)
}

// Debug acts as a standard fmt.Print.
// Sends data to all targets with xlevel.Debug.
// Line wrapping is controlled by the targets formatter.
func Debug(a ...interface{}) (int, error) {
	return fmt.Fprint(Default.Writer(xlevel.Debug), a...)
}

// Debugf acts as a standard fmt.Printf.
// Sends data to all targets with xlevel.Debug.
// Line wrapping is controlled by the targets formatter.
func (x Logger) Debugf(format string, a ...interface{}) (int, error) {
	return fmt.Fprintf(x.Writer(xlevel.Debug), format, a...)
}

// Debugf acts as a standard fmt.Printf.
// Sends data to all targets with xlevel.Debug.
// Line wrapping is controlled by the targets formatter.
func Debugf(format string, a ...interface{}) (int, error) {
	return fmt.Fprintf(Default.Writer(xlevel.Debug), format, a...)
}

// Info acts as a standard fmt.Print.
// Sends data to all targets with xlevel.Info.
// Line wrapping is controlled by the targets formatter.
func (x Logger) Info(a ...interface{}) (int, error) {
	return fmt.Fprint(x.Writer(xlevel.Info), a...)
}

// Info acts as a standard fmt.Print.
// Sends data to all targets with xlevel.Info.
// Line wrapping is controlled by the targets formatter.
func Info(a ...interface{}) (int, error) {
	return fmt.Fprint(Default.Writer(xlevel.Info), a...)
}

// Infof acts as a standard fmt.Printf.
// Sends data to all targets with xlevel.Info.
// Line wrapping is controlled by the targets formatter.
func (x Logger) Infof(format string, a ...interface{}) (int, error) {
	return fmt.Fprintf(x.Writer(xlevel.Info), format, a...)
}

// Infof acts as a standard fmt.Printf.
// Sends data to all targets with xlevel.Info.
// Line wrapping is controlled by the targets formatter.
func Infof(format string, a ...interface{}) (int, error) {
	return fmt.Fprintf(Default.Writer(xlevel.Info), format, a...)
}

// Warn acts as a standard fmt.Print.
// Sends data to all targets with xlevel.Warn.
// Line wrapping is controlled by the targets formatter.
func (x Logger) Warn(a ...interface{}) (int, error) {
	return fmt.Fprint(x.Writer(xlevel.Warn), a...)
}

// Warn acts as a standard fmt.Print.
// Sends data to all targets with xlevel.Warn.
// Line wrapping is controlled by the targets formatter.
func Warn(a ...interface{}) (int, error) {
	return fmt.Fprint(Default.Writer(xlevel.Warn), a...)
}

// Warnf acts as a standard fmt.Printf.
// Sends data to all targets with xlevel.Warn.
// Line wrapping is controlled by the targets formatter.
func (x Logger) Warnf(format string, a ...interface{}) (int, error) {
	return fmt.Fprintf(x.Writer(xlevel.Warn), format, a...)
}

// Warnf acts as a standard fmt.Printf.
// Sends data to all targets with xlevel.Warn.
// Line wrapping is controlled by the targets formatter.
func Warnf(format string, a ...interface{}) (int, error) {
	return fmt.Fprintf(Default.Writer(xlevel.Warn), format, a...)
}

// Error acts as a standard fmt.Print.
// Sends data to all targets with xlevel.Error.
// Line wrapping is controlled by the targets formatter.
func (x Logger) Error(a ...interface{}) (int, error) {
	return fmt.Fprint(x.Writer(xlevel.Error), a...)
}

// Error acts as a standard fmt.Print.
// Sends data to all targets with xlevel.Error.
// Line wrapping is controlled by the targets formatter.
func Error(a ...interface{}) (int, error) {
	return fmt.Fprint(Default.Writer(xlevel.Error), a...)
}

// Errorf acts as a standard fmt.Printf.
// Sends data to all targets with xlevel.Error.
// Line wrapping is controlled by the targets formatter.
func (x Logger) Errorf(format string, a ...interface{}) (int, error) {
	return fmt.Fprintf(x.Writer(xlevel.Error), format, a...)
}

// Errorf acts as a standard fmt.Printf.
// Sends data to all targets with xlevel.Error.
// Line wrapping is controlled by the targets formatter.
func Errorf(format string, a ...interface{}) (int, error) {
	return fmt.Fprintf(Default.Writer(xlevel.Error), format, a...)
}

// Fatal acts as a standard fmt.Print.
// Sends data to all targets with xlevel.Fatal.
// Calls os.Exit() or redefined exit function.
// Line wrapping is controlled by the targets formatter.
func (x Logger) Fatal(a ...interface{}) (int, error) {
	defer x.exit(1)
	return fmt.Fprint(x.Writer(xlevel.Fatal), a...)
}

// Fatal acts as a standard fmt.Print.
// Sends data to all targets with xlevel.Fatal.
// Calls os.Exit() or redefined exit function.
// Line wrapping is controlled by the targets formatter.
func Fatal(a ...interface{}) (int, error) {
	defer Default.exit(1)
	return fmt.Fprint(Default.Writer(xlevel.Fatal), a...)
}

// Fatalf acts as a standard fmt.Printf.
// Sends data to all targets with xlevel.Fatal.
// Calls os.Exit() or redefined exit function.
// Line wrapping is controlled by the targets formatter.
func (x Logger) Fatalf(format string, a ...interface{}) (int, error) {
	defer x.exit(1)
	return fmt.Fprintf(x.Writer(xlevel.Fatal), format, a...)
}

// Fatalf acts as a standard fmt.Printf.
// Sends data to all targets with xlevel.Fatal.
// Calls os.Exit() or redefined exit function.
// Line wrapping is controlled by the targets formatter.
func Fatalf(format string, a ...interface{}) (int, error) {
	defer Default.exit(1)
	return fmt.Fprintf(Default.Writer(xlevel.Fatal), format, a...)
}

// Panic acts as a standard fmt.Print.
// Sends data to all targets with xlevel.Panic.
// Calls panic().
// Line wrapping is controlled by the targets formatter.
func (x Logger) Panic(a ...interface{}) (int, error) {
	fmt.Fprint(x.Writer(xlevel.Panic), a...)
	panic(fmt.Sprint(a...))
}

// Panic acts as a standard fmt.Print.
// Sends data to all targets with xlevel.Panic.
// Calls panic().
// Line wrapping is controlled by the targets formatter.
func Panic(a ...interface{}) (int, error) {
	fmt.Fprint(Default.Writer(xlevel.Panic), a...)
	panic(fmt.Sprint(a...))
}

// Panicf acts as a standard fmt.Printf.
// Sends data to all targets with xlevel.Panic.
// Calls panic().
// Line wrapping is controlled by the targets formatter.
func (x Logger) Panicf(format string, a ...interface{}) (int, error) {
	fmt.Fprintf(x.Writer(xlevel.Panic), format, a...)
	panic(fmt.Sprintf(format, a...))
}

// Panicf acts as a standard fmt.Printf.
// Sends data to all targets with xlevel.Panic.
// Calls panic().
// Line wrapping is controlled by the targets formatter.
func Panicf(format string, a ...interface{}) (int, error) {
	fmt.Fprintf(Default.Writer(xlevel.Panic), format, a...)
	panic(fmt.Sprintf(format, a...))
}
