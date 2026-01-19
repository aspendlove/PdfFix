#!/bin/bash
# Usage: tmp_path input output
# Ensure bash fails on pipe error
set -euo pipefail
IFS=$'\n\t'

INPUT_FILE=$(readlink -f "$2")
OUTPUT_FILE=$(readlink -f "$3")

# sanitize inputs
if [ ! -f "$INPUT_FILE" ]; then
    echo "Error: Input file does not exist" >&2
    exit 1
fi
MIME_TYPE=$(file -b --mime-type "$INPUT_FILE")
if [ "$MIME_TYPE" != "application/pdf" ]; then
    echo "Error: Not a pdf" >&2
    exit 1
fi

# Generate unique subfolder
uuid="$(uuidgen)"
tmpDir="$1/$uuid"
mkdir -p $tmpDir
clean_up () {
  popd
  rm -rf "$tmpDir"
}
trap clean_up EXIT
pushd $tmpDir

gs -dSAFER -dBATCH -dNOPAUSE -sDEVICE=png16m -r300 -sOutputFile=temp_%03d.png "$INPUT_FILE"
magick -density 300 temp_*.png "$OUTPUT_FILE"
rm temp_*.png
