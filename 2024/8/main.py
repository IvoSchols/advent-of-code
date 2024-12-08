from itertools import combinations
from typing import Tuple, Dict, List, Set

def delta(a, b):
    return (a[0] - b[0], a[1] - b[1])

def out_of_bounds(map_size, position):
    width, height = map_size
    return not (0 <= position[0] < height and 0 <= position[1] < width)

def part_one(map_size, antennas):
    antinodes = set()

    for antenna_type, positions in antennas.items():
        for (a, b) in combinations(positions, 2):
            delta_a_b = delta(a, b)
            
            potential_antinodes = [
                (a[0] + delta_a_b[0], a[1] + delta_a_b[1]),
                (b[0] - delta_a_b[0], b[1] - delta_a_b[1])
            ]
            for antinode in potential_antinodes:
                if not out_of_bounds(map_size, antinode):
                    antinodes.add(antinode)

    print(len(antinodes))

def part_two(map_size, antennas) -> None:
    antinodes = set()

    for antenna_type, positions in antennas.items():
        for (a, b) in combinations(positions, 2):
            delta_a_b = delta(a, b)

            while not out_of_bounds(map_size, a):
                antinodes.add(a)
                a = (a[0] + delta_a_b[0], a[1] + delta_a_b[1])

            while not out_of_bounds(map_size, b):
                antinodes.add(b)
                b = (b[0] - delta_a_b[0], b[1] - delta_a_b[1])

    
    print(len(antinodes))

def main() -> None:
    with open("./8/input.txt", "r") as file:
        lines = file.readlines()

    antennas = {}
    width, height = len(lines[0]) - 1, len(lines) # -1 to remove newline character
    map_size = (width, height)
    
    for row, line in enumerate(lines):
        for column, antenna_type in enumerate(line.strip()):
            if antenna_type == ".":
                continue

            if antenna_type not in antennas:
                antennas[antenna_type] = []

            antenna_type = antenna_type
            antennas[antenna_type].append((row, column))
        
    part_one(map_size, antennas)
    part_two(map_size, antennas)

if __name__ == "__main__":
    main()