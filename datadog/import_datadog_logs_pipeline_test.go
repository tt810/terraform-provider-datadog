package datadog

import (
	"github.com/hashicorp/terraform/helper/resource"
	"testing"
)

const pipelineConfigForImport = `
resource "datadog_logs_pipeline" "test_import" {
	name = "imported pipeline"
	is_enabled = false
	filter {
		query = "source:kafka"
	}
}	
`

func TestAccLogsPipeline_importBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPipelineDestroy,
		Steps: []resource.TestStep{
			{
				Config: pipelineConfigForImport,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPipelineExists("datadog_logs_pipeline.test_import"),
				),
			},
			{
				ResourceName:      "datadog_logs_pipeline.test_import",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
