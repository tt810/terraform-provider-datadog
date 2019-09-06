package datadog

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/zorkian/go-datadog-api"
	"strings"
	"testing"
)

const pipelineConfigForCreation = `
resource "datadog_logs_pipeline" "my_pipeline_test" {
	name = "my first pipeline"
	is_enabled = true
	filter {
		query = "source:redis"
	}
}
`
const pipelineConfigForUpdate = `
resource "datadog_logs_pipeline" "my_pipeline_test" {
	name = "updated pipeline"
	is_enabled = false
	filter {
		query = "source:kafka"
	}
}
`

func TestAccExampleThing_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPipelineDestroy,
		Steps: []resource.TestStep{
			{
				Config: pipelineConfigForCreation,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPipelineExists("datadog_logs_pipeline.my_pipeline_test"),
					resource.TestCheckResourceAttr(
						"datadog_logs_pipeline.my_pipeline_test", "name", "my first pipeline"),
					resource.TestCheckResourceAttr(
						"datadog_logs_pipeline.my_pipeline_test", "is_enabled", "true"),
					resource.TestCheckResourceAttr(
						"datadog_logs_pipeline.my_pipeline_test", "filter.0.query", "source:redis"),
				),
			}, {
				Config: pipelineConfigForUpdate,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPipelineExists("datadog_logs_pipeline.my_pipeline_test"),
					resource.TestCheckResourceAttr(
						"datadog_logs_pipeline.my_pipeline_test", "name", "updated pipeline"),
					resource.TestCheckResourceAttr(
						"datadog_logs_pipeline.my_pipeline_test", "is_enabled", "false"),
					resource.TestCheckResourceAttr(
						"datadog_logs_pipeline.my_pipeline_test", "filter.0.query", "source:kafka"),
				),
			},
		},
	})
}


func testAccCheckPipelineExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*datadog.Client)
		if err := pipelineExistsChecker(s, client); err != nil {
			return err
		}
		return nil
	}
}

func pipelineExistsChecker(s *terraform.State, client *datadog.Client) error {
	for _, r := range s.RootModule().Resources {
		id := r.Primary.ID
		if _, err := client.GetLogsPipeline(id); err != nil {
			return fmt.Errorf("received an error when retrieving pipeline, (%s)", err)
		}
	}
	return nil
}

func testAccCheckPipelineDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*datadog.Client)
	if err := pipelineDestroyHelper(s, client); err != nil {
		return err
	}
	return nil
}

func pipelineDestroyHelper(s *terraform.State, client *datadog.Client) error {
	for _, r := range s.RootModule().Resources {
		id := r.Primary.ID
		p, err := client.GetLogsPipeline(id)

		if err != nil {
			if strings.Contains(err.Error(), "404 Not Found") {
				continue
			}
			return fmt.Errorf("received an error when retrieving pipeline, (%s)", err)
		}
		if p != nil {
			return fmt.Errorf("pipeline still exists")
		}
	}
	fmt.Println("XXXXXxxxxxxxxxxxxx")
	fmt.Println(s.RootModule().Resources)
	return nil
}
