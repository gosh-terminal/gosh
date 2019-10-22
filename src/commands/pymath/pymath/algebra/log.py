"""Module for logs."""


def log(base, awnser):
    """Solves logarithms."""
    if awnser == 1:
        return 0
    if base < 0 and awnser > 0:
        return "Error can not happen"
    result = 0
    exponent = 0
    while result != awnser:
        exponent += 1
        result = base**exponent
        if result == awnser:
            return exponent
