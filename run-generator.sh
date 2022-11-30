
# NOTE that openapi-fixed-dates.yaml avoids using:
#  enum:
#   - date
# and instead defines the type as 'date' directly.
#
# The enum form is rejected by the openapi-generator tool.
#
export GO_POST_PROCESS_FILE="gofmt -w"
openapi-generator generate -g go -o . -i openapi-fixed-dates.yaml --config config.yaml
