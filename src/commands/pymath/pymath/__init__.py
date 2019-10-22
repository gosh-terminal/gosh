"""Main pymath module"""
import pymath.algebra
import pymath.physics
import pymath.geometry
import pymath.chemistry
def unitConverter(original, translated, item):
    if original == "c" and translated == "f":
        f = item * 1.8000 + 32
        print(f"{f}° F")
