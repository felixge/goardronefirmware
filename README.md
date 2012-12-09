# goardronefirmware

This is my attempt at developing an open source firmware for the Parrot AR Drone
2.0.

I am using [Go](http://golang.org/) for several reasons:

* Small memory footprint
* Easy to cross compile for ARM
* Produces static binaries
* Reasonable performance
* Good concurrency model
* Nice type system
* Built-in http server

It is unclear if garbage collection pauses will become an issue, but I am
hopeful that if such issues manifest, they can be mitigated by reusing objects.

This project draws ideas and inspiration from:

* [Unofficial open firmware for 1.0 Drone in C](https://github.com/ardrone/ardrone): A lot of things have changed with the AR Drone 2, but some of the information and code is still very useful.
* [TU Delft - Search and Rescue with AR Drone 2](http://paparazzi.enac.fr/wiki/TU_Delft_-_Search_and_Rescue_with_AR_Drone_2): Their project also includes reverse engineering the hardware controls.

Besides that I am also using `strace` and the `/proc` folder on the drone
analyze the official firmware

## API Docs

The API is not stable, but if you're curious, here the documentation:

http://go.pkgdoc.org/github.com/felixge/goardronefirmware

## Current status

* Full control over each rotor (set speed between 0 - 511)
* Full control over all LEDs (set off, green, red or orange)
* Partial support for reading sensor data
* Basic http interface for some functionality

## Todo

* Fully parse sensor data
* Kalman filter for translating sensor data into attitude information
* Flight stablization loop
* Access to video cameras
* Edge detection with bottom camera to fight drift
* Full HTTP API
* Full UDP API (probably use the Parrot protocol for compatibility)

If you are interested in helping / exchanging information, please get in touch.
