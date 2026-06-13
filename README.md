# Slow

## An app that cares for your mental health

- Only for Mac at the moment!

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

Call it with trust, as still a new home brew

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

|   APP   | WINDOW | TOTAL TIME | TIME AV | SWITCHES |
|:-------:|:------:|:----------:|:-------:|:--------:|
| Ghostty |        |   2.1 s    |  2.1 s  |    1     |
|   Arc   |        |   0.6 s    |  0.6 s  |    1     |
|  None   |  None  |   0.3 s    |  0.3 s  |    1     |

### Only app tracking

You can also run the tracking loop alone:

```bash
slow track 
```

### To only do Dim loops (may feel more secure)

And the Dim loop alone (same flags can be set up)

```bash
slow dim 
```
