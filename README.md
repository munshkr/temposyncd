# TempoSync

Reliable music tempo synchronization over the network.

Please note this is a *work in progress*. See [DESIGN.md](DESIGN.md) to know
more about how this *would* work.

## Install

### Dependencies

To use TempoSync you need the Go compiler. Go to [Downloads](https://golang.org/dl/)
on [golang.org](golang.org) and install it.

### temposyncd Server

```
go get -u github.com/munshkr/temposyncd
```
### Client

Finally, you need to install a client to use it.
Follow instructions on each repository for more information.

#### Current implementations
* [Supercollider](https://github.com/munshkr/temposync-sc)

## Usage

Run as a Leader on (only) one machine:

```
temposyncd --leader
```

Then on the rest of the machines, run as a follower:

```
temposyncd
```

You can see more logging information with `--verbose`

## Development

1. Clone the repository
2. Install dependencies

   ```
   go get github.com/hypebeast/go-osc/osc
   ```
   
3. Build and run

   ```
   go build
   ./temposyncd -verbose -leader
   ```

## Contributing

Bug reports and pull requests are welcome on GitHub at
https://github.com/munshkr/temposyncd. This project is intended to be a safe,
welcoming space for collaboration, and contributors are expected to adhere to
the [Contributor Covenant](http://contributor-covenant.org) code of conduct.

## License

temposyncd  is under the Apache 2.0 license. See the [LICENSE](LICENSE) file
for details.
