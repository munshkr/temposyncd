# Design

*Work in progress*. Expect things to change

## Features

* No need to synchronize system clocks on all machines, only time
  offsets are relevant.
* Intented to work only over a local-area network (LAN).  Uses UDP multicast.
* Decentralized: All processes are identical, and at any time one of them is
  elected as the Tempo Leader.

## Leader and followers

A new process starts as a *follower*, and at the beginning sends a START
message and waits for a HEARTBEAT response from a *leader*.  Only leaders send
the HEARBEAT message.

After some randomized time, if a follower does not receive a HEARTBEAT, it
becomes a leader and begins a new *term*.  Otherwise, the timeout is reset and
waits again.

If a leader `i` receives a HEARTBEAT from another process `j`, and if `j < i`
or it is a leader from a future term, then `j` is the legitimate leader, so
process `i` becomes a follower again.  Otherwise, it ignores those messages.

It is important for a new process to first synchronize with the cluster,
because it might be out of phase in time with respect to the synchronized
clock.

## Time synchronization

Time offsets and round-trip delay are calculated as in NTP:

  Time offset θ is defined by

    θ = ((t1 - t0) + (t2 - t3)) / 2

  and the round-trip delay δ by

    δ = (t3 - t0) - (t2 - t1)

  where

  - t0 is the follower's timestamp of the request packet transmission
  - t1 is the leader's timestamp of the request packet reception
  - t2 is the leader's timestamp of the response packet transmission
  - t3 is the follower's timestamp of the response packet reception

## Changing current tempo

Any process can send the TEMPO message to the leader to set a new tempo value.

## Easing functions for dynamic tempo changing

to do...
