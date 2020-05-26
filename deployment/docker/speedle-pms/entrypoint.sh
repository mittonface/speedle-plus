#!/bin/sh

if [ ! -e /var/lib/speedle/policies.json ]; then
    echo "{}" > /var/lib/speedle/policies.json
fi


speedle-pms --filestore-loc /var/lib/speedle/policies.json --endpoint 0.0.0.0:6733 $*

