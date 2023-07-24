#!/bin/bash
tee data.yml << EOF
Namespace: <namespace>
EOF
./sct -d data.yml $*
rm data.yml