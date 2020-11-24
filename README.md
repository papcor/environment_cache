# environment_cache
is a small go module to cache supplied environment variables and convert them to methods which return those values as strings.

For example,

with

<code>
const (Port = "Port")
</code>

and

<code>
var port func() string
</code>,

setting

<code>
port = environment_cache.CacheRequiredEnvironmentVariable(Port)
</code>

will ensure that the "Port" environment variable exists and that port() will return a string containing its value. Running checks like this at program initialization time enables a cleaner way to handle fatal errors.

You can also use:
<code>
port = environment_cache.CacheEnvironmentVariable(Port)
</code>

for a non-mandatory/non-logging version which will also set the port() method to the cached value or a method that returns "" if the variable does not exist.