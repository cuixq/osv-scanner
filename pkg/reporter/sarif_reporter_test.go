package reporter_test

import (
	"bytes"
	"io"
	"testing"

	"github.com/google/osv-scanner/v2/pkg/reporter"
)

func TestSarifReporter_Errorf(t *testing.T) {
	t.Parallel()

	writer := &bytes.Buffer{}
	r := reporter.NewSarifReporter(io.Discard, writer, reporter.ErrorLevel)
	text := "hello world!"

	r.Errorf("%s", text)

	if writer.String() != text {
		t.Error("Error level message should have been printed")
	}
	if !r.HasErrored() {
		t.Error("HasErrored() should have returned true")
	}
}

func TestSarifReporter_Warnf(t *testing.T) {
	t.Parallel()

	text := "hello world!"
	tests := []struct {
		lvl              reporter.VerbosityLevel
		expectedPrintout string
	}{
		{lvl: reporter.WarnLevel, expectedPrintout: text},
		{lvl: reporter.ErrorLevel, expectedPrintout: ""},
	}

	for _, tt := range tests {
		writer := &bytes.Buffer{}
		r := reporter.NewSarifReporter(io.Discard, writer, tt.lvl)

		r.Warnf("%s", text)

		if writer.String() != tt.expectedPrintout {
			t.Errorf("expected \"%s\", got \"%s\"", tt.expectedPrintout, writer.String())
		}
	}
}

func TestSarifReporter_Infof(t *testing.T) {
	t.Parallel()

	text := "hello world!"
	tests := []struct {
		lvl              reporter.VerbosityLevel
		expectedPrintout string
	}{
		{lvl: reporter.InfoLevel, expectedPrintout: text},
		{lvl: reporter.WarnLevel, expectedPrintout: ""},
	}

	for _, tt := range tests {
		writer := &bytes.Buffer{}
		r := reporter.NewSarifReporter(io.Discard, writer, tt.lvl)

		r.Infof("%s", text)

		if writer.String() != tt.expectedPrintout {
			t.Errorf("expected \"%s\", got \"%s\"", tt.expectedPrintout, writer.String())
		}
	}
}
