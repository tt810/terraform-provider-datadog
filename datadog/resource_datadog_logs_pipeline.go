package datadog

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/zorkian/go-datadog-api"
	"strings"
)

func resourceDatadogLogsPipeline() *schema.Resource {
	return &schema.Resource{
		Create: resourceDatadogLogsPipelineCreate,
		Update: resourceDatadogLogsPipelineUpdate,
		Read:   resourceDatadogLogsPipelineRead,
		Delete: resourceDatadogLogsPipelineDelete,
		Exists: resourceDatadogLogsPipelineExists,
		Importer: &schema.ResourceImporter{
			State: resourceDatadogLogsPipelineImport,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of this pipeline.",
			},
			"is_enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether this pipeline is enabled.",
			},
			"filter": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"query": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The query of this pipeline filter.",
						},
					},
				},
			},
		},
	}
}

func resourceDatadogLogsPipelineCreate(d *schema.ResourceData, meta interface{}) error {
	fmt.Println("create... pipeline ")
	pipeline := buildDatadogLogsPipeline(d)
	createdPipeline, err := meta.(*datadog.Client).CreateLogsPipeline(pipeline)
	if err != nil {
		return fmt.Errorf("failed to create logs pipeline using Datadog API: %s", err.Error())
	}
	d.SetId(*createdPipeline.Id)
	fmt.Println(d)
	fmt.Println(d.Get("filter"))
	return resourceDatadogLogsPipelineRead(d, meta)
}

func resourceDatadogLogsPipelineRead(d *schema.ResourceData, meta interface{}) error {
	fmt.Printf("reading pipeline %s\n", d.Id())
	pipeline, err := meta.(*datadog.Client).GetLogsPipeline(d.Id())
	fmt.Println(pipeline)
	if err != nil {
		if strings.Contains(err.Error(), "404 Not Found") {
			d.SetId("")
		}
		return err
	}
	if err = d.Set("name", pipeline.Name); err != nil {
		return err
	}
	if err = d.Set("is_enabled", pipeline.IsEnabled); err != nil {
		return err
	}
	filter := make([]interface{}, 1)
	query := make(map[string]string)
	query["query"] = *pipeline.Filter.Query
	filter[0] = query
	if err = d.Set("filter", filter); err != nil {
		return err
	}
	fmt.Println("dddddddd")
	fmt.Println(d.Get("filter"))
	return nil
}

func resourceDatadogLogsPipelineUpdate(d *schema.ResourceData, meta interface{}) error {
	pipeline := buildDatadogLogsPipeline(d)
	client := meta.(*datadog.Client)
	if _, err := client.UpdateLogsPipeline(d.Id(), pipeline); err != nil {
		return fmt.Errorf("error updating logs pipeline: (%s)", err.Error())
	}
	return resourceDatadogLogsPipelineRead(d, meta)
}

func resourceDatadogLogsPipelineDelete(d *schema.ResourceData, meta interface{}) error {
	fmt.Println("deleted.....")
	if err := meta.(*datadog.Client).DeleteLogsPipeline(d.Id()); err != nil {
		fmt.Println(err.Error())
		if strings.Contains(err.Error(), "400 Bad Request") {
			return nil
		}
		return err
	}
	fmt.Println("deleted.....")
	return nil
}

func resourceDatadogLogsPipelineExists(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	fmt.Println("check if existing " + d.Id())
	client := meta.(*datadog.Client)
	if _, err := client.GetLogsPipeline(d.Id()); err != nil {
		if strings.Contains(err.Error(), "404 Not Found") {
			return false, nil
		}
		return false, err
	}
	return false, nil
}

func resourceDatadogLogsPipelineImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	fmt.Println("ddddddddd")
	pipeline, err := meta.(*datadog.Client).GetLogsPipeline(d.Id())

	fmt.Println(pipeline)
	if err != nil {
		return nil, err
	}
	if err = d.Set("name", pipeline.Name); err != nil {
		return nil, err
	}
	if err = d.Set("is_enabled", pipeline.IsEnabled); err != nil {
		return nil, err
	}
	filter := make([]interface{}, 1)
	query := make(map[string]string)
	query["query"] = *pipeline.Filter.Query
	filter[0] = query
	if err = d.Set("filter", filter); err != nil {
		return nil, err
	}
	return []*schema.ResourceData{d}, nil
}

func buildDatadogLogsPipeline(d *schema.ResourceData) *datadog.LogsPipeline {
	var pipeline datadog.LogsPipeline
	if v, ok := d.GetOk("name"); ok {
		name := v.(string)
		pipeline.Name = &name
	}
	if v, ok := d.GetOk("is_enabled"); ok {
		isEnabled := v.(bool)
		pipeline.IsEnabled = &isEnabled
	}
	tfFilters := d.Get("filter").([]interface{})
	ddFilter := &datadog.FilterConfiguration{}
	tfFilter := tfFilters[0]
	query := tfFilter.(map[string]interface{})["query"].(string)
	ddFilter.Query = &query
	pipeline.Filter = ddFilter
	return &pipeline
}
