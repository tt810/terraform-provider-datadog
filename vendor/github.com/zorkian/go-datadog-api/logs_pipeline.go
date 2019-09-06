/*
 * Datadog API for Go
 *
 * Please see the included LICENSE file for licensing information.
 *
 * Copyright 2019 by authors and contributors.
 */

package datadog

import (
	"fmt"
	"log"
)

const (
	LOGS_PIPELINE_PATH = "/v1/logs/config/pipelines"
)

type LogsPipeline struct {
	Id        *string              `json:"id,omitempty"`
	Name      *string              `json:"name"`
	IsEnabled *bool                `json:"is_enabled,omitempty"`
	Filter    *FilterConfiguration `json:"filter"`
	//Processers []Processor `json:"processors"`
}

type FilterConfiguration struct {
	Query *string `json:"query"`
}

func (client *Client) GetLogsPipeline(id string) (*LogsPipeline, error) {
	var pipeline LogsPipeline
	if err := client.doJsonRequest("GET", fmt.Sprintf(LOGS_PIPELINE_PATH+"/%s", id), nil, &pipeline); err != nil {
		return nil, err
	}
	return &pipeline, nil
}

func (client *Client) GetLogsPipelines() ([]*LogsPipeline, error) {
	var pipelines []*LogsPipeline
	if err := client.doJsonRequest("GET", LOGS_PIPELINE_PATH, nil, &pipelines); err != nil {
		return nil, err
	}
	return pipelines, nil
}

func (client *Client) CreateLogsPipeline(pipeline *LogsPipeline) (*LogsPipeline, error) {
	var createdPipeline = &LogsPipeline{}
	log.Println(pipeline)
	if err := client.doJsonRequest("POST", LOGS_PIPELINE_PATH, pipeline, createdPipeline); err != nil {
		return nil, err
	}
	log.Println(createdPipeline)
	return createdPipeline, nil
}

func (client *Client) UpdateLogsPipeline(id string, pipeline *LogsPipeline) (*LogsPipeline, error) {
	var updatedPipeline = &LogsPipeline{}
	if err := client.doJsonRequest("PUT", fmt.Sprintf(LOGS_PIPELINE_PATH+"/%s", id), pipeline, updatedPipeline); err != nil {
		return nil, err
	}
	return updatedPipeline, nil
}

// DeleteLogsPipeline returns 200 OK when operation succeed
func (client *Client) DeleteLogsPipeline(id string) error {
	return client.doJsonRequest("DELETE", fmt.Sprintf(LOGS_PIPELINE_PATH+"/%s", id), nil, nil)
}
