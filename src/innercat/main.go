package main;

import (
	"os";
	"io";
	"bufio";
	"fmt";
);

/* Filename passed as as arguments. */
var filenames []string;

/* Constructor routines. */
func init_filenames(args[]string) (filenames []string) {
	/* Get number of program arguments. */
	arglen := len(args);
	if (arglen == 1) {
		fmt.Printf("Usage: %s filepath1 filepath2 ...\n", args[0]);
		os.Exit(0);
	}
	/* Get all filenames. */
	filenames = make([]string, arglen - 1);
	for i := 1; i < arglen; i += 1 {
		filenames[i - 1] = args[i];
	}
	return filenames;
}

/* Destructor routines. */

/* Global constructor. */
func init() {
	filenames = init_filenames(os.Args);
}

/* Global destructor. */
func fini() {
	// Global cleanup
}

func main() {
	defer fini();
	submain(filenames);
}

/* Concatenate lines from each file. */
func innerConcatLines(fds []*os.File) {
	var done bool;
	reader := make([]*bufio.Reader, len(fds));
	for i, fd := range(fds) {
		reader[i] = bufio.NewReader(fd);
	}
	for {
		cat := "";
		for _, r := range(reader) {
			line, err := r.ReadString(10);
			if err == io.EOF {
				done = true;
				break;
			} else if err != nil {
				panic(err);
			}
			cat += line;
		}
		if (done) {
			break;
		}
		fmt.Printf("%s", cat);
	}
}

func openfiles(fn []string) (filedescr []*os.File) {
	var err error;
	/* Open files. */
	l := len(fn);
	filedescr = make([]*os.File, l);
	for i := 0; i < l; i += 1 {
		filedescr[i], err = os.Open(fn[i]);
		if err != nil {
			panic(err);
		}
	}
	return filedescr;
}

func closefiles(fds []*os.File) {
	/* Close all files. */
	for _, fd := range(fds) {
		if err := fd.Close(); err != nil {
			panic(err);
		}
	}
}

func submain(args[]string) {
	files := openfiles(args);
	defer closefiles(files);
	innerConcatLines(files);
}

