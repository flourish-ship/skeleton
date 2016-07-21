package debuger

import (
	"fmt"
	"runtime"
)

func Display(info string, o interface{}) string {
	var pc, file, line, ok = runtime.Caller(2)
	if !ok {
		return ""
	}

	var buf = new(bytes.Buffter)
	fmt.Fprintf(buf, "[Debug] at %s() [%s:%d]\n", function(pc), file, line)
	fmt.Fprintf(buf, "\n[Variables]\n")

	for i := 0; i < len(data); i += 2 {
		var output = fomateinfo(len(data[i].(string))+3, data[i+1])
		fmt.Fprintf(buf, "%s = %s", data[i], output)
	}

	if displayed {
		log.Print(buf)
	}
	return buf.String()
}

// return data dump and format bytes
func fomateinfo(headlen int, data ...interface{}) []byte {
	var buf = new(bytes.Buffer)

	if len(data) > 1 {
		fmt.Fprint(buf, "    ")

		fmt.Fprint(buf, "[")

		fmt.Fprintln(buf)
	}

	for k, v := range data {
		var buf2 = new(bytes.Buffer)
		var pointers *pointerInfo
		var interfaces = make([]reflect.Value, 0, 10)

		printKeyValue(buf2, reflect.ValueOf(v), &pointers, &interfaces, nil, true, "    ", 1)

		if k < len(data)-1 {
			fmt.Fprint(buf2, ", ")
		}

		fmt.Fprintln(buf2)

		buf.Write(buf2.Bytes())
	}

	if len(data) > 1 {
		fmt.Fprintln(buf)

		fmt.Fprint(buf, "    ")

		fmt.Fprint(buf, "]")
	}

	return buf.Bytes()
}

// return the name of the function containing the PC if possible,
func function(pc uintptr) []byte {
	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return dunno
	}
	name := []byte(fn.Name())
	// The name includes the path name to the package, which is unnecessary
	// since the file name is already included.  Plus, it has center dots.
	// That is, we see
	//	runtime/debug.*T·ptrmethod
	// and want
	//	*T.ptrmethod
	if period := bytes.Index(name, dot); period >= 0 {
		name = name[period+1:]
	}
	name = bytes.Replace(name, centerDot, dot, -1)
	return name
}
