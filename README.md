# Slow

## An app that cares for your mental health

Slow is an app to remind you how much you are using your computer.
I made it thinking on the mental health problems
arising with the new AI addictions and psychosis.

It does something simple:

- Blinks light every cycle (default 60 min)
- Blinks dark after half the total 21:49
- After total time (default 8 cycles) makes your screen brightness go dark.

Additionally:

- It tracks which app and window name your are using
- it keeps it in an SQLite locally
- You can ask for usage

It has two loops to achieve that using go, with minimum impact to
your machine (8M of RAM)

## Install

```bash

brew tap gabesoler/tap
```

```bash
brew trust --cask gabesoler/tap/slow
```

```bash

brew install --cask slow
```

## Use

### For it to run both loops

When using slow or command dim, you can pass flags
'duration' and 'cycles' to customise the blinking behaviour.

```bash
slow
```

Imagine you want to work for four hours you can pass:

```bash
slow -d 30 -c 4
```

With this configuration your computer will blink
every 30 minutes for 4 cycles.
Mid way it will increase 20% darkness.

This aims to give you feedback of how much time is passing while working.
After the desired time, the screen will go darker and you can run the usage
command to see where you have been.

### To check the records

To check your use you just type 'slow usage'. You can pass
the flag '--back' or '-b' to say how many days back to search in the database.
I added not just time,
but also the amount of times you have switched from the same window,
as it shows attention
fragmentation and concentration.

```bash
slow usage 
```

Example
┌──────────────────┬──────────────────────┬────────────┬─────────┬──────────┐
│       APP        │        WINDOW        │ TOTAL TIME │ TIME AV │ SWITCHES │
├──────────────────┼──────────────────────┼────────────┼─────────┼──────────┤
│ Arc              │                      │ 11m        │ 5m      │ 23       │
│ Ghostty          │                      │ 7m         │ 3m      │ 21       │
│ Activity Monitor │                      │ 4m         │ 4m      │ 3        │
│ ghostty          │ ./slow --duration... │ 16s        │ 16s     │ 1        │
│ ghostty          │ ./slow               │ 14s        │ 14s     │ 2        │
│ Arc              │ vi - Find and rep... │ 8s         │ 8s      │ 3        │
│ TextEdit         │                      │ 5s         │ 5s      │ 1        │
│ Prime Video      │ Prime Video          │ 5s         │ 5s      │ 1        │
│ ghostty          │ go run .             │ 3s         │ 3s      │ 1        │
│ Finder           │                      │ 3s         │ 3s      │ 1        │
│ Notion           │ Apple                │ 2s         │ 2s      │ 1        │
└──────────────────┴──────────────────────┴────────────┴─────────┴──────────┘

### To only do app tracking

You can also run the tracking loop alone:

```bash
slow track 
```

### To only do Dim loops (may feel more secure)

And the Dim loop alone (same flags can be set up)

```bash
slow dim 
```
