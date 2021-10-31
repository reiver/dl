package help

import (
	"flag"
	"fmt"
	"io"
	"strings"
	"unicode/utf8"
)

func WriteTo(w io.Writer) (int64, error) {

	var msg strings.Builder

	fmt.Fprint(&msg,
		"dl — a program for downloading files and streams"                                          +"\n"+
		""                                                                                          +"\n"+
		"usage:"                                                                                    +"\n",
		"\tdl [options] [target]"                                                                   +"\n"+
		""                                                                                          +"\n"+
		"examples:"                                                                                 +"\n"+
		"\tdl gemini://example.com/apple/banana/cherry"                                             +"\n"+
		"\tdl http://example.com/path/to/file.txt"                                                  +"\n"+
		"\tdl mercury://example.com/path/to/file.txt"                                               +"\n"+
		""                                                                                          +"\n"+
		"\tdl --target='gemini://example.com/apple/banana/cherry'"                                  +"\n"+
		"\tdl --target='http://example.com/path/to/file.txt'"                                       +"\n"+
		"\tdl --target='mercury://example.com/path/to/file.txt'"                                    +"\n"+
		""                                                                                          +"\n"+
		"\tdl --address='cache.something.com' gemini://example.com/apple/banana/cherry"             +"\n"+
		"\tdl --address='cache.something.com' mercury://example.com/apple/banana/cherry"            +"\n"+
		""                                                                                          +"\n"+
		"\tdl --address='cache.something.com' --target='gemini://example.com/apple/banana/cherry'"  +"\n"+
		"\tdl --address='cache.something.com' --target='mercury://example.com/apple/banana/cherry'" +"\n"+
		""                                                                                          +"\n"+
		"flags:"                                                                                    +"\n",
	)
	flag.VisitAll(func(fl *flag.Flag){
		switch fl.Name {
		case "v", "vv", "vvv", "vvvv", "vvvvv", "vvvvvv":
			fmt.Fprint(&msg, " -")
		default:
			switch {
			case 2 > utf8.RuneCountInString(fl.Name):
				fmt.Fprint(&msg, " -")
			default:
				fmt.Fprint(&msg, " --")
			}
		}
		fmt.Fprintf(&msg, "%s\n\t%s\n\n", fl.Name, fl.Usage)

	})


	{
		n, err := io.WriteString(w, msg.String())

		var n64 int64 = int64(n)

		return n64, err
	}
}
