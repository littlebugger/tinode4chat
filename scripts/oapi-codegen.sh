#!/bin/bash

go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest

# Define the base directory for OpenAPI specs and output
OPENAPI_DIR="./api/openapi"
OUTPUT_BASE_DIR="./pkg/server"

# Remove previous generated server code directories
rm -rf $OUTPUT_BASE_DIR

# Iterate over all YAML files in the OPENAPI_DIR
for OPENAPI_SPEC in $OPENAPI_DIR/*.yaml; do

  # Extract the spec filename (without path and extension)
  SPEC_NAME=$(basename "$OPENAPI_SPEC" .yaml)

  # Define the output directory for this specific spec
  OUTPUT_DIR="$OUTPUT_BASE_DIR/$SPEC_NAME"

  mkdir -p $OUTPUT_DIR

  echo "Generating Go server code for spec: $SPEC_NAME"

  # Generate Go server code using oapi-codegen
  oapi-codegen --config=./configs/oapi-codegen.yaml -o "$OUTPUT_DIR/$SPEC_NAME.go" "$OPENAPI_SPEC"

  # Check if the generation was successful
  if [ $? -eq 0 ]; then
    echo "Server code generated successfully for $SPEC_NAME at $OUTPUT_DIR"
  else
    echo "Error: Failed to generate server code for $SPEC_NAME"
  fi

done