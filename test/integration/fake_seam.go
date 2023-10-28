package integration

import (
	"fmt"
	"os"
	"os/exec"
	"sync/atomic"
	"testing"
	"time"

	seamclient "github.com/seamapi/go/client"
)

const (
	fakeSeamToken          = "seam_apikey1_token"
	fakeSeamConnectCommand = "fake-seam-connect"
)

var fakeSeamConnectPort atomic.Uint64

func newFakeSeam(t *testing.T) (*seamclient.Client, func()) {
	port := fakeSeamConnectPort.Add(1)
	cmd := exec.Command(fakeSeamConnectCommand, "--seed")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = append(os.Environ(), fmt.Sprintf("PORT=%d", port))

	go func() {
		if err := cmd.Run(); err != nil {
			_, _ = os.Stderr.WriteString(fmt.Sprintf("%s failed to shutdown gracefully: %v", fakeSeamConnectCommand, err))
		}
	}()

	time.Sleep(5 * time.Second)

	client := seamclient.NewClient(
		seamclient.WithBaseURL(fmt.Sprintf("http://localhost:%d", port)),
		seamclient.WithAuthApiKey(fakeSeamToken),
	)

	cleanup := func() {
		if err := cmd.Process.Kill(); err != nil {
			_, _ = os.Stderr.WriteString(fmt.Sprintf("failed to stop %s: %v", fakeSeamConnectCommand, err))
		}
	}

	return client, cleanup
}

func init() {
	fakeSeamConnectPort.Store(8999)
}
