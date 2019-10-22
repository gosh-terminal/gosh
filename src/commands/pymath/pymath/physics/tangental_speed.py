from math import pi


def tangental_speed(radius, time, cycles):
    awns = (2*pi*radius) / (time / cycles)
    return awns
