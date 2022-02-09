package client

import (
	//mocks "github.com/projectdiscovery/nuclei/v2/pkg/web/api/client/mocks"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTemplates(t *testing.T) {
	setup := NewMockHttpServer(t)
	defer setup()

	client := New(WithBasicAuth("user", "pass"))
	svc := TemplatesService{Client: client}
	t.Run("AddTemplate", func(t *testing.T) {
		resp, err := svc.AddTemplate(AddTemplateRequest{
			Folder: "nuclei-templates",
			Path:   "cves/2000/CVE-2000-0116.yaml",
		})
		if err == nil {
			require.Greater(t, resp, int64(0))
		}
	})
	t.Run("GetTemplates", func(t *testing.T) {
		resp, err := svc.GetTemplates(GetTemplatesRequest{
			Folder: "nuclei-templates",
			Search: "CVE-2000-0114.yaml",
		})
		require.NoError(t, err, "could not get templates")
		require.Greater(t, len(resp), 0)
	})
	t.Run("UpdateTemplate", func(t *testing.T) {
		err := svc.UpdateTemplate(UpdateTemplateRequest{
			Path: "cves/2000/CVE-2000-0116.yaml",
		})
		require.NoError(t, err, "could not update template")
	})
	t.Run("DeleteTemplate", func(t *testing.T) {
		err := svc.DeleteTemplate(DeleteTemplateRequest{
			Path: "cves/2000/CVE-2000-0116.yaml",
		})
		require.NoError(t, err, "could not delete template")
	})
	t.Run("GetTemplateRaw", func(t *testing.T) {
		resp, err := svc.GetTemplateRaw("test.yaml")
		require.NoError(t, err, "could not get raw template")
		require.NotEmpty(t, resp)
	})
	t.Run("ExecuteTemplate", func(t *testing.T) {
		_, err := svc.ExecuteTemplate(ExecuteTemplateRequest{
			Path:   "technologies/ibm/ibm-http-server.yaml",
			Target: "example.com",
		})
		require.NoError(t, err, "could not execute template")
	})
}
