from itertools import product
from enum import Enum

# Using the Enum class to represent directions
class Direction(Enum):
    UP = (-1, 0)
    DOWN = (1, 0)
    LEFT = (0, -1)
    RIGHT = (0, 1)
    UP_LEFT = (-1, -1)
    UP_RIGHT = (-1, 1)
    DOWN_LEFT = (1, -1)
    DOWN_RIGHT = (1, 1)

# Simplified tuple operations using standard tuples
class CustomTuple(tuple):
    def __add__(self, other):
        if isinstance(other, Direction):
            return CustomTuple((self[0] + other.value[0], self[1] + other.value[1]))
        return CustomTuple((self[0] + other[0], self[1] + other[1]))

    def __sub__(self, other):
        if isinstance(other, Direction):
            return CustomTuple((self[0] - other.value[0], self[1] - other.value[1]))
        return CustomTuple((self[0] - other[0], self[1] - other[1]))

    def __eq__(self, other):
        return self[0] == other[0] and self[1] == other[1]

    def __hash__(self):
        return hash((self[0], self[1]))


# Simplifying direction movements initialization
def setup_movements():
    manhattan_movements = {d: CustomTuple(d.value) for d in list(Direction)[:4]}
    all_movements = {d: CustomTuple(d.value) for d in Direction}
    return manhattan_movements, all_movements


MANHATTAN_MOVEMENTS, MOVEMENTS = setup_movements()
INV_MOVEMENTS = {v: k for k, v in MOVEMENTS.items()}

def out_of_bounds(i, j, garden_map):
    return i < 0 or j < 0 or i >= len(garden_map) or j >= len(garden_map[0])

def get_neighbors(i, j, garden_map, movement=MANHATTAN_MOVEMENTS):
    current_char = garden_map[i][j]
    neighbors = []

    for delta in movement.values():
        new_pos = CustomTuple((i, j)) + delta
        if not out_of_bounds(new_pos[0], new_pos[1], garden_map) and garden_map[new_pos[0]][new_pos[1]] == current_char:
            neighbors.append(new_pos)
    return neighbors

def get_segments(map):
    not_visited = {(i, j) for i, j in product(range(len(map)), range(len(map[0])))}
    segments = []

    while not_visited:
        current_position = not_visited.pop()
        queue = [current_position]
        segment = [current_position]

        while queue:
            i, j = queue.pop(0)
            for neighbor in get_neighbors(i, j, map):
                if neighbor in not_visited:
                    not_visited.remove(neighbor)
                    queue.append(neighbor)
                    segment.append(neighbor)

        segments.append(segment)
    return segments

def part_one(map):
    total = 0

    for segment in get_segments(map):
        area = len(segment)
        perimeter = 4 * len(segment) - sum(len(get_neighbors(i, j, map)) for i, j in segment)
        total += area * perimeter
    
    print(total)

def get_segment_neighbors(i, j, segment):
    return [new_pos for delta in MANHATTAN_MOVEMENTS.values() if (new_pos := CustomTuple((i, j)) + delta) in segment]

def count_corner_one_neighbor(position, neighbor, segment):
    return 2

def count_corner_two_neighbors(position, neighbors, segment):
    direction_neighbor_one = neighbors[0] - position
    direction_neighbor_two = neighbors[1] - position
    if direction_neighbor_one + direction_neighbor_two == CustomTuple((0, 0)): # Check if part of a straight line
        return 0

    # Check if inner corner is filled
    inner_corner_position = position + direction_neighbor_one + direction_neighbor_two
    if inner_corner_position in segment:
        return 1 # Still have outside corner
    return 2

def count_corner_three_neighbors(position, neighbors, segment):
    direction_neighbor_one = neighbors[0] - position
    direction_neighbor_two = neighbors[1] - position
    direction_neighbor_three = neighbors[2] - position

    direction = direction_neighbor_one + direction_neighbor_two + direction_neighbor_three
    potential_corners = 2

    if INV_MOVEMENTS[direction] == Direction.UP:
        corner_directions = [MOVEMENTS[Direction.UP_LEFT], MOVEMENTS[Direction.UP_RIGHT]]
        if (position + corner_directions[0]) in segment:
            potential_corners -= 1
        if (position + corner_directions[1]) in segment:
            potential_corners -= 1
    if INV_MOVEMENTS[direction] == Direction.DOWN:
        corner_directions = [MOVEMENTS[Direction.DOWN_LEFT], MOVEMENTS[Direction.DOWN_RIGHT]]
        if (position + corner_directions[0]) in segment:
            potential_corners -= 1
        if (position + corner_directions[1]) in segment:
            potential_corners -= 1
    if INV_MOVEMENTS[direction] == Direction.LEFT:
        corner_directions = [MOVEMENTS[Direction.UP_LEFT], MOVEMENTS[Direction.DOWN_LEFT]]
        if (position + corner_directions[0]) in segment:
            potential_corners -= 1
        if (position + corner_directions[1]) in segment:
            potential_corners -= 1
    if INV_MOVEMENTS[direction] == Direction.RIGHT:
        corner_directions = [MOVEMENTS[Direction.UP_RIGHT], MOVEMENTS[Direction.DOWN_RIGHT]]
        if (position + corner_directions[0]) in segment:
            potential_corners -= 1
        if (position + corner_directions[1]) in segment:
            potential_corners -= 1

    return potential_corners

def count_corner_four_neighbors(position, neighbors, segment):
    potential_corners = 4

    if position + MOVEMENTS[Direction.UP_LEFT] in segment:
        potential_corners -= 1
    if position + MOVEMENTS[Direction.UP_RIGHT] in segment:
        potential_corners -= 1
    if position + MOVEMENTS[Direction.DOWN_LEFT] in segment:
        potential_corners -= 1
    if position + MOVEMENTS[Direction.DOWN_RIGHT] in segment:
        potential_corners -= 1
    return potential_corners


#1122281 too high
#799390 too low
def part_two(map):
    total = 0
    for segment in get_segments(map):
        area = len(segment)
        sides = 0

        for position in segment:
            i, j = position
            position = CustomTuple((i, j))
            neighbor_pos = get_segment_neighbors(i, j, segment)
            neighbor_count = len(neighbor_pos)

            match neighbor_count:
                case 0:
                    sides += 4
                case 1:
                    sides += count_corner_one_neighbor(position, neighbor_pos[0], segment)
                case 2:
                    sides += count_corner_two_neighbors(position, neighbor_pos, segment)
                case 3:
                    sides += count_corner_three_neighbors(position, neighbor_pos, segment)
                case 4:
                    sides += count_corner_four_neighbors(position, neighbor_pos, segment)
                case _:
                    raise ValueError("Invalid neighbor count")

        total += area * sides     


    print(total)

def main():
    garden_map = []

    with open("./12/input.txt", "r") as file:
        for line in file:
            garden_map.append(list(line.strip()))

    part_one(garden_map)
    part_two(garden_map)

if __name__ == "__main__":
    main()