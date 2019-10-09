package main

import (
	"fmt"
	s "strings"
)

var p = fmt.Println

func main() {
	p("Contains: ", s.Contains("test", "es"))
	p("Count: ", s.Count("test", "t"))
	p("HasPrefix: ", s.HasPrefix("test", "t"))
	p("HasSuffix: ", s.HasSuffix("test", "est"))
	p("Index: ", s.Index("test", "s"))
	p("Join: ", s.Join([]string{"a", "b"}, "-"))
	p("Repeat: ", s.Repeat("a", 5))
	p("Replace: ", s.Replace("foo", "o", "0", -1))
	p("Replace: ", s.Replace("fooo", "o", "0", 2))
	p("Split: ", s.Split("t----e-s-t", "-"))
	p("ToLower: ", s.ToLower("tESt"))
	p("ToUpper: ", s.ToUpper("test"))
	p("Len: ", len("sinsombat"))
}
