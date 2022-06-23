package xout

import (
	"fmt"
	"os"

	"github.com/overred/xout/xlevel"
)

// Print acts like fmt.Print.
func (x Logger) Print(a ...interface{}) (int, error) {
	return fmt.Fprint(x.Writer(xlevel.Text), a...)
}

// Println acts like fmt.Println.
func (x Logger) Println(a ...interface{}) (int, error) {
	return fmt.Fprintln(x.Writer(xlevel.Text), a...)
}

// Printf acts like fmt.Printf.
func (x Logger) Printf(format string, a ...interface{}) (int, error) {
	return fmt.Fprintf(x.Writer(xlevel.Text), format, a...)
}

// Trace acts like fmt.Print.
func (x Logger) Trace(a ...interface{}) (int, error) {
	return fmt.Fprint(x.Writer(xlevel.Trace), a...)
}

// Traceln acts like fmt.Println.
func (x Logger) Traceln(a ...interface{}) (int, error) {
	return fmt.Fprintln(x.Writer(xlevel.Trace), a...)
}

// Tracef acts like fmt.Printf.
func (x Logger) Tracef(format string, a ...interface{}) (int, error) {
	return fmt.Fprintf(x.Writer(xlevel.Trace), format, a...)
}

// Debug acts like fmt.Print.
func (x Logger) Debug(a ...interface{}) (int, error) {
	return fmt.Fprint(x.Writer(xlevel.Debug), a...)
}

// Debugln acts like fmt.Println.
func (x Logger) Debugln(a ...interface{}) (int, error) {
	return fmt.Fprintln(x.Writer(xlevel.Debug), a...)
}

// Debugf acts like fmt.Printf.
func (x Logger) Debugf(format string, a ...interface{}) (int, error) {
	return fmt.Fprintf(x.Writer(xlevel.Debug), format, a...)
}

// Info acts like fmt.Print.
func (x Logger) Info(a ...interface{}) (int, error) {
	return fmt.Fprint(x.Writer(xlevel.Info), a...)
}

// Infoln acts like fmt.Println.
func (x Logger) Infoln(a ...interface{}) (int, error) {
	return fmt.Fprintln(x.Writer(xlevel.Info), a...)
}

// Infof acts like fmt.Printf.
func (x Logger) Infof(format string, a ...interface{}) (int, error) {
	return fmt.Fprintf(x.Writer(xlevel.Info), format, a...)
}

// Warn acts like fmt.Print.
func (x Logger) Warn(a ...interface{}) (int, error) {
	return fmt.Fprint(x.Writer(xlevel.Warn), a...)
}

// Warnln acts like fmt.Println.
func (x Logger) Warnln(a ...interface{}) (int, error) {
	return fmt.Fprintln(x.Writer(xlevel.Warn), a...)
}

// Warnf acts like fmt.Printf.
func (x Logger) Warnf(format string, a ...interface{}) (int, error) {
	return fmt.Fprintf(x.Writer(xlevel.Warn), format, a...)
}

// Error acts like fmt.Print.
func (x Logger) Error(a ...interface{}) (int, error) {
	return fmt.Fprint(x.Writer(xlevel.Error), a...)
}

// Errorln acts like fmt.Println.
func (x Logger) Errorln(a ...interface{}) (int, error) {
	return fmt.Fprintln(x.Writer(xlevel.Error), a...)
}

// Errorf acts like fmt.Printf.
func (x Logger) Errorf(format string, a ...interface{}) (int, error) {
	return fmt.Fprintf(x.Writer(xlevel.Error), format, a...)
}

// Fatal acts like fmt.Print.
// BE AWARE: This function calls os.Exit().
func (x Logger) Fatal(a ...interface{}) (int, error) {
	fmt.Fprint(x.Writer(xlevel.Fatal), a...)
	os.Exit(1)
	return 0, nil
}

// Fatalln acts like fmt.Println.
// BE AWARE: This function calls os.Exit().
func (x Logger) Fatalln(a ...interface{}) (int, error) {
	fmt.Fprintln(x.Writer(xlevel.Fatal), a...)
	os.Exit(1)
	return 0, nil
}

// Fatalf acts like fmt.Printf.
// BE AWARE: This function calls os.Exit().
func (x Logger) Fatalf(format string, a ...interface{}) (int, error) {
	fmt.Fprintf(x.Writer(xlevel.Fatal), format, a...)
	os.Exit(1)
	return 0, nil
}

// Panic acts like fmt.Print.
// BE AWARE: This function calls panic().
func (x Logger) Panic(a ...interface{}) (int, error) {
	fmt.Fprint(x.Writer(xlevel.Panic), a...)
	panic(fmt.Sprint(a...))
}

// Panicln acts like fmt.Println.
// BE AWARE: This function calls panic().
func (x Logger) Panicln(a ...interface{}) (int, error) {
	fmt.Fprintln(x.Writer(xlevel.Panic), a...)
	panic(fmt.Sprintln(a...))
}

// Panicf acts like fmt.Printf.
// BE AWARE: This function calls panic().
func (x Logger) Panicf(format string, a ...interface{}) (int, error) {
	fmt.Fprintf(x.Writer(xlevel.Panic), format, a...)
	panic(fmt.Sprintf(format, a...))
}
