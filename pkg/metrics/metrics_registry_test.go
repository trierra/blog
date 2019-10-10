package metrics

import (
	"bytes"
	"fmt"
	"net/http"
	"testing"
	"time"

	_ "github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/require"
)

func TestErrorInc(t *testing.T) {
	httpAddr := fmt.Sprintf("http://127.0.0.1%s%s", port, endpoint)

	t.Run("server", func(*testing.T) {
		go func() {
			err := Run()
			require.NoError(t, err)
		}()
	})

	// sleep is here because the thread above could need some time to start
	time.Sleep(time.Second * 2)

	expected := "# HELP errors_total Total amount of errors\n" +
		"# TYPE errors_total counter\n" +
		"errors_total{error=\"error\"} 5\n"

	ErrorInc() //1
	ErrorInc() //2
	ErrorInc() //3
	ErrorInc() //4
	ErrorInc() //5

	resp, err := http.Get(httpAddr)
	require.NoError(t, err)

	if resp.StatusCode != 200 {
		t.Fatalf("Expected 200, got: %#v", resp)
	}

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(resp.Body)
	require.NoError(t, err)

	res := buf.String()
	require.Equal(t, expected, res)
}
