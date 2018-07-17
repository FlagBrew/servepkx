# servepkx-go Windows by Allen Lydiard

This is the version of servepkx-go that runs on Windows.

# Usage

## Method 1: using CMD/PowerShell

### Pushing a file
`` servepk.exe filename.fileext ``

Example:

``  servepk.exe pikachu.pk7 ``

### Pushing a folder

`` servepk.exe folder ``

Example:

`` servepk.exe pokemon ``

## Method 2: Drag'n'Drop

Simply drag a compatible file/folder onto the executable and 
the pokemon will be uploaded

# Troubleshooting

## The program is stuck on "Looking for 3DS IP, please wait!"

This issue sometimes occurs if the 3DS fails to respond to the port check.
Simply just stop the script by doing ctrl + c and rerun it.

## I get a "DS not found on network!" error

This happens when PKSM isn't listening for incoming connections.
Please check to make sure you have launched the PKSM wireless server on your 3DS
and try again!

## I get an index out of range error

This happens when your pokemon file is corrupted. Please try fetching a clean version of
your pokemon file and try again.

## The program says the pokemon has been uploaded but I don't see it.

This can happen for a few reasons
- Attempting to use a wc7/wc7full/wc6/wc6full file outside of the event injector
- Another device on your network has port 9000 open.

Problem 1 is simple: just use the event injector screen and not the storage or editor screen

Problem 2 is a bit more complicated: I discovered that devices such as
- Google Home
- Google ChromeCast

Are listening for port 9000, which means that they might get picked up as a 3DS IP

I fixed this by checking to see if they also have port 8008 open as well, as these devices do
And if the port is open, then the IP is skipped and the next one is checked.

There might be cases where the wrong IP is chosen and if it does, just post an issue on the repo and
I'll take a look at it.