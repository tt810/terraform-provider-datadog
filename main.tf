resource "datadog_logs_pipeline" "my" {
	name = "my pipeline"
	is_enabled = true
	filter {
		query = "source:kafka"
	}
	processor {
		arithmetic_processor {
			name = "updated arithmetic processor"
			is_enabled = true
			expression = "(time1 - time2)*1000"
			target = "my_arithmetic"
			is_replace_missing = true
		}
	}
	processor {
        attribute_remapper {
            name = "test attribute processor"
            is_enabled = true
            sources = ["db.instance"]
            source_type = "tag"
            target = "db"
            target_type = "tag"
        }
    }
    processor {
        category_processor {
            name = "test category processor"
            target = "redis.severity"
            category {
                name = "debug"
                filter {
                    query = "@severity: \".\""
                }
            }
            category {
                name = "verbose"
                filter {
                    query = "@severity: \"-\""
                }
            }
        }
    }
    processor {
        date_remapper {
            is_enabled = true
            sources = ["date"]
        }
    }
    processor {
        message_remapper {
            name = "test message remapper"
            is_enabled = false
            sources = ["message"]
        }
    }
    processor {
        status_remapper {
            name = "test status remapper"
            sources = ["status", "extra"]
        }
    }
    processor {
        trace_id_remapper {
            name = "test trace id remapper"
            sources = ["dd.trace_id"]
        }
    }
    processor {
        service_remapper {
            name = "test service remapper"
            sources = ["service"]
        }
    }
    processor {
        grok_parser {
            name = "test grok parser"
            source = "message"
            grok {
                support_rules = ""
                match_rules = "Rule %%{word:my_word2} %%{number:my_float2}"
            }
        }
    }
    processor {
        pipeline {
            name = "nested pipeline"
            filter {
                query = "source:redis"
            }
            processor {
                grok_parser {
                    name = "test grok parser"
                    source = "message"
                    grok {
                        support_rules = ""
                        match_rules = "Rule %%{word:my_word2} %%{number:my_float2}"
                    }
                }
            }
            processor {
                url_parser {
                    name = "test url parser"
                    sources = ["url", "extra"]
                    target = "http_url"
                }
            }
        }
    }
    processor {
        user_agent_parser {
            name = "test user agent parser"
            sources = ["user", "agent"]
            target = "http_agent"
        }
    }
}

resource "datadog_logs_pipelineorder" "my_order" {
    name = "my_order"
    depends_on = [
        "datadog_logs_pipeline.my"
    ]
    pipelines = [
        "7mKmvI1DTVuGKPT4hk1aNQ",
        "x4NYdjW1RjyredO5JsTdZA",
        "iuhVw_ZvT4m8RAQ73i0Bdg",
        "${datadog_logs_pipeline.my.id}"
    ]
}