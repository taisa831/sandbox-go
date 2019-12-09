package testing

import (
	"bytes"
	"fmt"
	"html/template"
	"math"
	"math/rand"
	"testing"
	"time"
)

func TestAbs(t *testing.T) {
	got := math.Abs(-1)
	if got != 1 {
		t.Errorf("Abs(-1) = %f; want 1", got)
	}
}

func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("hello")
	}
}

func BenchmarkTemplateParallel(b *testing.B) {
	templ := template.Must(template.New("test").Parse("Hello, {{.}}!"))
	b.RunParallel(func(pb *testing.PB) {
		var buf bytes.Buffer
		for pb.Next() {
			buf.Reset()
			templ.Execute(&buf, "World")
		}
	})
}

func ExampleHello() {
	fmt.Println("hello")
	// Output: hello
}

func ExampleSalutations() {
	fmt.Println("hello, and")
	fmt.Println("goodbye")
	// Output:
	// hello, and
	// goodbye
}

func ExamplePerm() {
	for _, value := range rand.Perm(5) {
		fmt.Println(value)
	}
	// Unordered output: 4
	// 2
	// 1
	// 3
	// 0
}

func ExampleF() {
	for _, value := range rand.Perm(4) {
		fmt.Println(value)
	}
}

func TestTimeConsuming(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
}

func TestFoo(t *testing.T) {
	t.Run("A=1", func(t *testing.T) {
		got := math.Abs(-1)
		if got != 1 {
			t.Errorf("Abs(-1) = %f; want 1", got)
		}
	})
	t.Run("A=2", func(t *testing.T) {
		got := math.Abs(-1)
		if got != 1 {
			t.Errorf("Abs(-1) = %f; want 1", got)
		}
	})
	t.Run("B=1", func(t *testing.T) {
		got := math.Abs(-1)
		if got != 1 {
			t.Errorf("Abs(-1) = %f; want 1", got)
		}
	})
}

//     go test -run ''      # Run all tests.
//     go test -run Foo     # Run top-level tests matching "Foo", such as "TestFooBar".
//     go test -run Foo/A=  # For top-level tests matching "Foo", run subtests matching "A=".
//     go test -run /A=1    # For all top-level tests, run subtests matching "A=1".

func TestGroupedParallel(t *testing.T) {
	tests := []struct {
		Name    string
		Want    bool
		WantErr bool
	}{
		{
			Name: "test1",
			Want: true,
			WantErr: false,
		},
		//{
		//	Name: "test2",
		//	Want: true,
		//	WantErr: false,
		//},
		//{
		//	Name: "test3",
		//	Want: true,
		//	WantErr: false,
		//},
		//{
		//	Name: "test4",
		//	Want: true,
		//	WantErr: false,
		//},
	}

	for _, tc := range tests {
		tc := tc // capture range variable
		t.Run(tc.Name, func(t *testing.T) {
			t.Run("A=1", func(t *testing.T) {
				t.Parallel()
				got := math.Abs(-1)
				time.Sleep(2 * time.Second)
				if got != 1 {
					t.Errorf("Abs(-1) = %f; want 1", got)
				}
				t.Logf("Len=2: %s", time.Now())
			})
			t.Run("A=2", func(t *testing.T) {
				t.Parallel()
				got := math.Abs(-1)
				if got != 1 {
					t.Errorf("Abs(-1) = %f; want 1", got)
				}
				t.Logf("Len=2: %s", time.Now())
			})
			t.Run("B=1", func(t *testing.T) {
				t.Parallel()
				got := math.Abs(-1)
				if got != 1 {
					t.Errorf("Abs(-1) = %f; want 1", got)
				}
				t.Logf("Len=2: %s", time.Now())
			})
		})
	}
	t.Logf("%s", "finished")
}

func TestTeardownParallel(t *testing.T) {
	// This Run will not return until the parallel tests finish.
	t.Run("group", func(t *testing.T) {
		t.Run("Test1", func(t *testing.T) {
			t.Parallel()
			got := math.Abs(-1)
			time.Sleep(2 * time.Second)
			if got != 1 {
				t.Errorf("Abs(-1) = %f; want 1", got)
			}
			t.Logf("Len=2: %s", time.Now())
		})
		t.Run("Test2", func(t *testing.T) {
			t.Parallel()
			got := math.Abs(-1)
			time.Sleep(2 * time.Second)
			if got != 1 {
				t.Errorf("Abs(-1) = %f; want 1", got)
			}
			t.Logf("Len=2: %s", time.Now())
		})
		t.Run("Test3", func(t *testing.T) {
			t.Parallel()
			got := math.Abs(-1)
			time.Sleep(2 * time.Second)
			if got != 1 {
				t.Errorf("Abs(-1) = %f; want 1", got)
			}
			t.Logf("Len=2: %s", time.Now())
		})
	})
	// <tear-down code>
	t.Logf("%s", "finished")
}

//func TestMain(m *testing.M) {
//	println("前処理")
//
//	m.Run()
//
//	println("後処理")
//}
