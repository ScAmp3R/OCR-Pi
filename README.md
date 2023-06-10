# Overview
This project is build for personal use case.
Our boiler can encounter errors in the night which leads to lost heating.
Subzero temperatures outside can lead to frozen pipes if that happens during the winter.

This script grabs a still image of the ESP32-Cam Module in the same network
and runs ssocr lib over it to read the seven segment display of the boiler.

If the recognized text matches an unknown state it triggers the homebridge switch to notify everyone
