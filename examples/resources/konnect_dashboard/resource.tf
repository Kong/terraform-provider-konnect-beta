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
            table_chart_tile_definition = {
              chart = {
                chart_title = "...my_chart_title..."
                type        = "table"
              }
              query = {
                columns = [
                  "..."
                ]
                cursor     = "...my_cursor..."
                datasource = "platform_usage"
                entity     = "...my_entity..."
                filters = [
                  {
                    field    = "hostname"
                    operator = "not_in"
                    value = [
                      "..."
                    ]
                  }
                ]
                page_size = 4
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