package config

import (
	"flag"
	"fmt"
	"log"
	"log/slog"
	"os"
)

var (
	DIR        string
	PORT       int
	PROTECTED  bool
	HELP       bool
	DIR_LOGGER string = "./data/app.log"
	LOGGER     *slog.Logger
	LOGFILE    *os.File
	helpTxt    string = `
	Doodocs days-2 backend project
	
	Usage:
		server [--port <N>] [--dir [S]]
		server --help	
	
	Options:
		--help		Show this screen.
		--porn N	Port number.
		--dir S		Path to the data directory.	
`
)

var (
	ErrMimeNotSupported   error = fmt.Errorf("this mimetype is not supported. Only: application/vnd.openxmlformats-officedocument.wordprocessingml.document | application/xml | image/jpeg | image/png")
	ErrFormatNotSupported error = fmt.Errorf("provided file has unsupported format")
	ErrInvalidZipFile     error = fmt.Errorf("provided file is not ZIP archive")
	ErrWrongArraySize     error = fmt.Errorf("unexpected or wrong size of array")
	ErrEmptyFile          error = fmt.Errorf("file data is empty or file is missing")
	ErrCorruptedFileData  error = fmt.Errorf("file data is empty or corrupted")
)

func init() {
	flag.Usage = func() {
		fmt.Print(helpTxt)
		os.Exit(0)
	}

	flag.StringVar(&DIR, "dir", "./data", "directory to serve")
	flag.IntVar(&PORT, "port", 8080, "port to serve")
	flag.BoolVar(&PROTECTED, "ip", false, "log request ip")
	flag.BoolVar(&HELP, "help", false, "show help")

	flag.Parse()

	if HELP {
		fmt.Print(helpTxt)
		os.Exit(0)
	}

	checkFlags()

	if _, err := os.Stat(DIR_LOGGER); os.IsNotExist(err) {
		os.WriteFile(DIR_LOGGER, []byte(""), 0o644)
	}
	LOGFILE, err := os.OpenFile(DIR_LOGGER, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o666)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to open log file:", err)
	}

	LOGGER = slog.New(slog.NewTextHandler(LOGFILE, nil))
}

func checkFlags() {
	if len(flag.Args()) > 0 {
		log.Fatalf("Unexpected argument: %v\n%s", flag.Args(), helpTxt)
	}

	if PORT < 1024 || PORT > 49151 {
		log.Fatalf("Port number must be in the range [1024, 49151]\n%s", helpTxt)
	}

	if DIR == "" {
		log.Fatalf("Directory path is empty\n%s", helpTxt)
	}

	os.MkdirAll(DIR, os.ModePerm)
}
