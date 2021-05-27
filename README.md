# goroutineschannels
Sample to program to illustrate use of channels in Go routines

## Functionality
This program makes use of channels for transmitting and consuming data between Go routines.

Here one go routine prints only odd numbers and another one prints only even number. Goal here is to print the numbers in sequential order only. To achieve this program uses two channels; one for passing odd number from even Go routine to odd Go routine and second for vice-versa.
