# Design

*Work in progress*. Expect things to change

## Requirements and constraints

* There should be no need to synchronize system clocks on all machines, only
  time offsets are relevant.
* To reliably work over a local-area network (LAN) should be the priority, as
  most live musical events are local performances.
* **Decentralized*: All processes should be identical, and at any time one of
  them could be elected as the Tempo Leader.
* It should allow any machine to enter or exit the *stage* at any given time
  while clock is running, and keep everyone in sync.
* Users should be able to quantize tempo or transport changes (i.e. change tempo
  at the next beat).

## Short description

### Version 0.1

This is a first version, inspired by the MIDI protocol.  Followers only advance
their clocks when receiving TICKs from the leader.

* A specific process is started as a leader. Otherwise, by default all processes
  are followers.
* Leader starts immediately broadcasting TICK messages, with BPM, current beat
  and tick as metadata.
* Followers only listen to TICK messages, and when a follower receives it, she
  forwards it via OSC to localhost:57120 (Supercollider)

### Version 0.2

There is no need to constantly send TICKs to everyone. Each machine has a much
more precise clock than one that ticks via UDP.  More importantly, followers
should set their clocks and start at the same time as the leader.

* When a follower enters the stage, she broadcasts a HELLO message.
* The leader listens to a HELLO message from some new process, and tries to
  connect (via TCP) to that follower.
* At first (and then, periodically), they exchange a number of SYNC messages
  to figure out the time offset between them.
* The follower then maintains connections for all known followers and keeps
  track of their respective time offsets.
* When the leader starts beating, it sends a PULSE message to each acquainted
  follower.
* PULSE messages contains metadata, like BPM, next beat timestamp (local to
  follower) and current beat.
* See below on Time Synchronization for more information on how the time offset
  between the leader and her followers is calculated.

## Misc.

*These notes are from old brainstorming sessions, but they are still here as
they could be useful.*

### Leader and followers

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

### Time synchronization

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

### Changing current tempo

Any process can send the TEMPO message to the leader to set a new tempo value.

### Easing functions for dynamic tempo changing

to do...
