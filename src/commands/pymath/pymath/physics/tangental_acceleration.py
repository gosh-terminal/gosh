from math import pi


def tangental_acceleration(radius, time, cycles):
    awns = (4 * pi**2 * radius) / (time / cycles)**2
    return awns
