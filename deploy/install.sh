#!/bin/bash
set -o nounset
set -o pipefail


echo Usage: $(basename "$0") "[--registry | --node ] "

if [ $# -gt 0 ]; then
  case $1 in
    -y|--tar)
      echo "install registry ......"
      set -e
      pack_tar
      ;;
    -r|--rpm)
      echo "install node......"
      pack_rpm
      ;;
    *)
      echo "Please specify your install mode"
      echo $1
      exit 1
      ;;
  esac
else
  echo "default action: install node"
  pack_tar
fi