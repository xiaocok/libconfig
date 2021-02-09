package libconfig

import (
	"log"
	"testing"
)

func TestStartEndString(t *testing.T) {
	f := func(s, expectedResult string) {
		t.Helper()
		result := startEndString(s)
		if result != expectedResult {
			t.Fatalf("unexpected result for startEndString(%q); got %q; want %q", s, result, expectedResult)
		}
	}
	f("", "")
	f("foo", "foo")

	getString := func(n int) string {
		b := make([]byte, 0, n)
		for i := 0; i < n; i++ {
			b = append(b, 'a'+byte(i%26))
		}
		return string(b)
	}
	s := getString(maxStartEndStringLen)
	f(s, s)

	f(getString(maxStartEndStringLen+1), "abcdefghijklmnopqrstuvwxyzabcdefghijklmn...pqrstuvwxyzabcdefghijklmnopqrstuvwxyzabc")
	f(getString(100*maxStartEndStringLen), "abcdefghijklmnopqrstuvwxyzabcdefghijklmn...efghijklmnopqrstuvwxyzabcdefghijklmnopqr")
}

func TestMatchFile(t *testing.T) {
	log.Println(matchFile("test.cfg", "test*.cfg"))  //true
	log.Println(matchFile("test1.cfg", "test*.cfg")) //true
	log.Println(matchFile("test.cfg", "*.cfg"))      //true
	log.Println(matchFile("test.conf", "*.cfg"))     //false
	log.Println(matchFile("test.conf", "test.*"))    //true
	log.Println(matchFile("test.cfg", "test.*"))     //true

	log.Println(matchFile("test1_demo2_example3.cfg", "test*demo*example.cfg"))  //false
	log.Println(matchFile("test1_demo2_example3.cfg", "test*demo*example*.cfg")) //true
}
