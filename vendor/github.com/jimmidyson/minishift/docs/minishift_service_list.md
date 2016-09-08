## minishift service list

Lists the URLs for the services in your local cluster

### Synopsis


Lists the URLs for the services in your local cluster

```
minishift service list [flags]
```

### Options

```
  -n, --namespace="": The services namespace
```

### Options inherited from parent commands

```
      --alsologtostderr[=false]: log to standard error as well as files
      --format="http://{{.IP}}:{{.Port}}": Format to output service URL in
      --log-flush-frequency=5s: Maximum number of seconds between log flushes
      --log_backtrace_at=:0: when logging hits line file:N, emit a stack trace
      --log_dir="": If non-empty, write log files in this directory
      --logtostderr[=false]: log to standard error instead of files
      --show-libmachine-logs[=false]: Whether or not to show logs from libmachine.
      --stderrthreshold=2: logs at or above this threshold go to stderr
      --v=0: log level for V logs
      --vmodule=: comma-separated list of pattern=N settings for file-filtered logging
```

### SEE ALSO
* [minishift service](minishift_service.md)	 - Gets the URL for the specified service in your local cluster

