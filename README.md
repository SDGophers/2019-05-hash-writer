# Hash writer

For this meetup we'll be writing a file copying application. The command line
interface will be simple, but the functionality will be powerful. Based on a
config file, the program will recursively copy all files from a root directory
and output to multiple locations. This could be useful for generating backups.


## Usage

```bash
$ hw config_file config_option root
```

## The config file

The config file will contain groups of config options. The first line will be
the name of the config option, the second line will be the capuring expression
and the following lines until an empty line will be the output location
templates. For instructions on how to use go's regex templates refer to the go
documentation on
[`func (\* regexp.Regexp) Expand`](https://golang.org/pkg/regexp/#Regexp.Expand)

Here's an example config file:

```text
name
regexp_capture
regexp_dest1
regexp_dest2
regexp_dest3

gofiles
(.*\.go)
/backup/gofiles/$1
/backup2/lib/go/$1
```

## Architecture

Everything pretty much revolves around the `Config` struct. The map
is populated with `ConfigOption`'s by the `ParseConfig` function. Note
that the parser only populates the `capture` and `tmpls` strings, and does
not compile the regular expressions. In order to compile them,
`func (* ConfigOption) Compile` must be called.

The `func (*Config) Write` method, will execute the recursive walk and call
the `func (*ConfigOption) Write` method at each file. This mechanism
uses [`filepath.Walk`](https://golang.org/pkg/path/filepath/#Walk), with
the `func (*ConfigOption) Write`  method as the `WalkFunc`. This method will
be doing the actual copying of files.

```go
type Config struct {
    m Map
}

type Map interface {
    Get(string) (interface{}, bool)
    Set(string, interface{})
    Del(string)
}

type ConfigOption struct {
    capture string
    tmpls []string

    captureExp *regexp.Regexp
}
```

## The Challange

There are two parts to the challenge:
* the map
* `func (*ConfigOption) Write`

Some things have already been written to get you up and running
and instructions are provided but they don't have to be followed.

### Map

This will be a hashmap implementation. The
[wikipedia page](https://golang.org/pkg/path/filepath/#Walk)
has detailed information on variations and theory of hashmaps.
Essentially, a hash map is a data structure for associating keys
to values (for example, a name to a phone number). In our case, it will
be configuration names to the configuration itself.

Go provides it's own hashmap implemention and you could use it for reference
(or to skip this part of the challange altogether). Go's implementaion is
generic, which means almost any type can be used as a key and any time could
be used as a value. In our case the keys will be strings, but the values
will be `interface{}` which all types can be reduced to.

The simplest way to write a map is separate chaining, it's also not very
efficient. In this algorith, the map has an array (static size) of buckets,
and each bucket is a list (dynamic size) of key-value pairs.

To store a key-value pair (kv pair), the key is passed to a hash function
(provided) which returns a number, the hash. One property of this function
is that for the same input it will always return the same output. Also, the
distribution of these values (should be) uniform. The hash is then
modded by the number of buckets to get the bucket the kv pair belongs to.
The kv pair is then appended to the list at it's bucket.

To retireve or delete a value from the map, its associated key is passed and a
similar proces to storing is executed. The key is hashed, then modded by the
number of buckets to find the bucket. Then, we search through the linked list
and try to find a kv pair with the same key.


### `func (*ConfigOption) Write`

This function is much less abstract than the hash map. It is an implementaion
of `filepath.WalkFunc`, it has the following signature:

```go
func (conf *ConfigOption) Write (path string, info os.FileInfo, err error) error
```

First we have to see if the `path` matches the `conf.capture` regular
expression. If it doesn't we don't have to deal with this file and can return.
If it does, we then have to loop through `conf.tmpls` and generate new file
names for the destinations using
[`func (*Regexp) Expand`](https://golang.org/pkg/regexp/#Regexp.Expand)
Remeber to `defer` closing these files!

## Testing

...

# Have fun!

