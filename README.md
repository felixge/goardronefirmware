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

* [Unofficial open firmware for 1.0 Drone in C](https://github.com/ardrone/ardrone)
* [TU Delft - Search and Rescue with AR Drone 2](http://paparazzi.enac.fr/wiki/TU_Delft_-_Search_and_Rescue_with_AR_Drone_2)
