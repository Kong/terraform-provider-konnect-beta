resource "konnect_dashboard" "my_dashboard" {
  provider = konnect-beta
  name     = "Example Dashboard"
  labels = {
    hello = "world"
  }
  definition = {
    preset_filters = null
    tiles = [
      {
        chart = {
          definition = {
            chart = {
              donut          = null
              horizontal_bar = null
              single_value   = null
              timeseries_line = {
                chart_title = "Total traffic over time"
                stacked     = null
                type        = "timeseries_line"
              }
            }
            query = {
              api_usage = {
                datasource = "api_usage"
                dimensions = ["time"]
                filters = [
                ]
                granularity = null
                metrics     = ["request_count"]
                time_range  = null
              }
              llm_usage = null
            }
          }
          layout = {
            position = {
              col = 0
              row = 0
            }
            size = {
              cols = 6
              rows = 2
            }
          }
          type = "chart"
        }
      },
      {
        chart = {
          definition = {
            chart = {
              donut = null
              horizontal_bar = {
                chart_title = "Top gateway services by requests"
                stacked     = true
                type        = "horizontal_bar"
              }
              single_value    = null
              timeseries_line = null
            }
            query = {
              api_usage = {
                datasource = "api_usage"
                dimensions = ["gateway_service"]
                filters = [
                  {
                    field    = "gateway_service"
                    operator = "not_empty"
                    value    = null
                  },
                ]
                granularity = null
                metrics     = ["request_count"]
                time_range  = null
              }
              llm_usage = null
            }
          }
          layout = {
            position = {
              col = 0
              row = 2
            }
            size = {
              cols = 2
              rows = 2
            }
          }
          type = "chart"
        }
      },
      {
        chart = {
          definition = {
            chart = {
              donut = null
              horizontal_bar = {
                chart_title = "Top routes by requests"
                stacked     = true
                type        = "horizontal_bar"
              }
              single_value    = null
              timeseries_line = null
            }
            query = {
              api_usage = {
                datasource = "api_usage"
                dimensions = ["route"]
                filters = [
                  {
                    field    = "route"
                    operator = "not_empty"
                    value    = null
                  },
                ]
                granularity = null
                metrics     = ["request_count"]
                time_range  = null
              }
              llm_usage = null
            }
          }
          layout = {
            position = {
              col = 2
              row = 2
            }
            size = {
              cols = 2
              rows = 2
            }
          }
          type = "chart"
        }
      },
      {
        chart = {
          definition = {
            chart = {
              donut = null
              horizontal_bar = {
                chart_title = "Top consumers by requests"
                stacked     = true
                type        = "horizontal_bar"
              }
              single_value    = null
              timeseries_line = null
            }
            query = {
              api_usage = {
                datasource = "api_usage"
                dimensions = ["consumer"]
                filters = [
                  {
                    field    = "consumer"
                    operator = "not_empty"
                    value    = null
                  },
                ]
                granularity = null
                metrics     = ["request_count"]
                time_range  = null
              }
              llm_usage = null
            }
          }
          layout = {
            position = {
              col = 4
              row = 2
            }
            size = {
              cols = 2
              rows = 2
            }
          }
          type = "chart"
        }
      },
      {
        chart = {
          definition = {
            chart = {
              donut          = null
              horizontal_bar = null
              single_value   = null
              timeseries_line = {
                chart_title = "Latency breakdown over time"
                stacked     = null
                type        = "timeseries_line"
              }
            }
            query = {
              api_usage = {
                datasource = "api_usage"
                dimensions = ["time"]
                filters = [
                ]
                granularity = null
                metrics     = ["response_latency_p99", "response_latency_p95", "response_latency_p50"]
                time_range  = null
              }
              llm_usage = null
            }
          }
          layout = {
            position = {
              col = 0
              row = 4
            }
            size = {
              cols = 3
              rows = 2
            }
          }
          type = "chart"
        }
      },
      {
        chart = {
          definition = {
            chart = {
              donut          = null
              horizontal_bar = null
              single_value   = null
              timeseries_line = {
                chart_title = "Kong vs upstream latency over time"
                stacked     = null
                type        = "timeseries_line"
              }
            }
            query = {
              api_usage = {
                datasource = "api_usage"
                dimensions = ["time"]
                filters = [
                ]
                granularity = null
                metrics     = ["upstream_latency_p99", "kong_latency_p99"]
                time_range  = null
              }
              llm_usage = null
            }
          }
          layout = {
            position = {
              col = 3
              row = 4
            }
            size = {
              cols = 3
              rows = 2
            }
          }
          type = "chart"
        }
      },
      {
        chart = {
          definition = {
            chart = {
              donut = null
              horizontal_bar = {
                chart_title = "Slowest gateway services (average)"
                stacked     = true
                type        = "horizontal_bar"
              }
              single_value    = null
              timeseries_line = null
            }
            query = {
              api_usage = {
                datasource = "api_usage"
                dimensions = ["gateway_service"]
                filters = [
                  {
                    field    = "gateway_service"
                    operator = "not_empty"
                    value    = null
                  },
                ]
                granularity = null
                metrics     = ["response_latency_average"]
                time_range  = null
              }
              llm_usage = null
            }
          }
          layout = {
            position = {
              col = 0
              row = 6
            }
            size = {
              cols = 2
              rows = 2
            }
          }
          type = "chart"
        }
      },
      {
        chart = {
          definition = {
            chart = {
              donut = null
              horizontal_bar = {
                chart_title = "Slowest routes (average)"
                stacked     = true
                type        = "horizontal_bar"
              }
              single_value    = null
              timeseries_line = null
            }
            query = {
              api_usage = {
                datasource = "api_usage"
                dimensions = ["route"]
                filters = [
                  {
                    field    = "route"
                    operator = "not_empty"
                    value    = null
                  },
                ]
                granularity = null
                metrics     = ["response_latency_average"]
                time_range  = null
              }
              llm_usage = null
            }
          }
          layout = {
            position = {
              col = 2
              row = 6
            }
            size = {
              cols = 2
              rows = 2
            }
          }
          type = "chart"
        }
      },
      {
        chart = {
          definition = {
            chart = {
              donut = null
              horizontal_bar = {
                chart_title = "Slowest consumers (average)"
                stacked     = true
                type        = "horizontal_bar"
              }
              single_value    = null
              timeseries_line = null
            }
            query = {
              api_usage = {
                datasource = "api_usage"
                dimensions = ["consumer"]
                filters = [
                  {
                    field    = "consumer"
                    operator = "not_empty"
                    value    = null
                  },
                ]
                granularity = null
                metrics     = ["response_latency_average"]
                time_range  = null
              }
              llm_usage = null
            }
          }
          layout = {
            position = {
              col = 4
              row = 6
            }
            size = {
              cols = 2
              rows = 2
            }
          }
          type = "chart"
        }
      },
    ]
  }
}
