resource "konnect_dashboard" "my_dashboard" {
  provider = konnect-beta
  definition = {
    preset_filters = [
      {
        field    = "a2a_method"
        operator = "not_in"
        value    = "{ \"see\": \"documentation\" }"
      }
    ]
    tiles = [
      {
        chart = {
          definition = {
            chart = {
              choropleth_map = {
                chart_title = "...my_chart_title..."
                type        = "choropleth_map"
              }
            }
            query = {
              llm_usage = {
                datasource = "llm_usage"
                dimensions = [
                  "consumer"
                ]
                filters = [
                  {
                    field    = "application"
                    operator = "empty"
                    value    = "{ \"see\": \"documentation\" }"
                  }
                ]
                granularity = "tenMinutely"
                metrics = [
                  "ai_request_count"
                ]
                time_range = {
                  absolute = {
                    end   = "2022-11-26T07:30:44.592Z"
                    start = "2022-01-09T02:25:36.303Z"
                    type  = "absolute"
                    tz    = "Etc/UTC"
                  }
                }
              }
            }
          }
          layout = {
            position = {
              col = 4
              row = 5
            }
            size = {
              cols = 6
              rows = 8
            }
          }
          type = "chart"
        }
      }
    ]
  }
  labels = {
    key = "value"
  }
  name = "...my_name..."
}