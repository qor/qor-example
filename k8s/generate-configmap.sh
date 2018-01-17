#!/bin/bash

cat <<-EOF
apiVersion: v1
kind: ConfigMap
metadata:
  name: configs
data:
EOF

for f in config/*yml
do
  echo "  $(basename $f): |+"
  cat $f | sed "s/^/    /g"
done
