package podinfo

import (
	"testing"

	openapiclient "github.com/bluecontainer/podinfo-client-go"
	"github.com/stretchr/testify/assert"
)

func Test_podinfo_model_http(t *testing.T) {

	t.Run("Test HttpRuntimeResponse copy", func(t *testing.T) {

		src := new(openapiclient.HttpRuntimeResponse)

		src.SetHostname("host1")

		dest := new(openapiclient.HttpRuntimeResponse)
		src.DeepCopyInto(dest)

		assert.Equal(t, src.GetHostname(), dest.GetHostname())

	})
}
