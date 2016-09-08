#!/bin/bash

# Copyright (C) 2016 Red Hat, Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# Setup a persistent directory for /var/lib/openshift
BOOT2DOCKER_DATA=`blkid -o device -l -t LABEL=boot2docker-data`
PARTNAME=`echo "$BOOT2DOCKER_DATA" | sed 's/.*\///'`

mkdir -p /mnt/$PARTNAME/var/lib/minishift
ln -s /mnt/$PARTNAME/var/lib/minishift /var/lib/minishift
mkdir -p /mnt/$PARTNAME/data
ln -s /mnt/$PARTNAME/data /data
