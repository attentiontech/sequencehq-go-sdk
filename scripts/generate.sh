echo -e

# ::config::
GEN_VERSION="v2.4.1"

# ::steps::
echo "Installing oapi-codegen $GEN_VERSION"
go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@$GEN_VERSION

TEMP_FILE="$(mktemp).json"
echo "Downloading OpenAPI schema to $TEMP_FILE"
wget -q -O $TEMP_FILE https://raw.githubusercontent.com/SequenceHQ/openapi-schema/refs/heads/main/2024-07-30.product.json
echo "Generating Go code"
oapi-codegen --config=scripts/gen.yaml $TEMP_FILE
