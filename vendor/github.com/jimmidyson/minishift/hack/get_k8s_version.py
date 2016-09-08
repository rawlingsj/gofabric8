#!/usr/bin/env python

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

"This package gets the LD flags used to set the version of kubernetes."

import json
import subprocess
import sys

K8S_PACKAGE = 'k8s.io/kubernetes/'
X_ARG_BASE = '-X github.com/jimmidyson/minishift/vendor/k8s.io/kubernetes/pkg/version.'

def get_rev():
  return 'gitCommit=%s' % get_from_godep('Rev')

def get_version():
    return 'gitVersion=%s' % get_from_godep('Comment')

def get_from_godep(key):
  with open('./Godeps/Godeps.json') as f:
    contents = json.load(f)
    for dep in contents['Deps']:
      if dep['ImportPath'].startswith(K8S_PACKAGE):
        return dep[key]

def get_tree_state():
  git_status = subprocess.check_output(['git', 'status', '--porcelain'])
  if git_status:
    result = 'dirty'
  else :
    result = 'clean'
  return 'gitTreeState=%s' % result

def main():
  args = [get_rev(), get_version(), get_tree_state()]
  return ' '.join([X_ARG_BASE + arg for arg in args])

if __name__ == '__main__':
  sys.exit(main())
