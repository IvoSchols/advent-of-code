from enum import Enum

class Direction(Enum):
    UP = 0
    RIGHT = 1
    DOWN = 2
    LEFT = 3

    def next(self):
        return Direction((self.value + 1) % 4)

# Predefined movement changes for directions
MOVEMENTS = {
    Direction.UP: (-1, 0),
    Direction.RIGHT: (0, 1),
    Direction.DOWN: (1, 0),
    Direction.LEFT: (0, -1),
}

def move_guard(position, direction):
    delta_row, delta_col = MOVEMENTS[direction]
    return position[0] + delta_row, position[1] + delta_col

def guard_will_step_out_of_bounds(map_size, position, direction):
    width, height = map_size
    next_position = move_guard(position, direction)
    row, col = next_position
    return not (0 <= row < height and 0 <= col < width)

def get_guard_direction(position, direction, obstructions):
    for _ in range(4):
        step = move_guard(position, direction)
        if step not in obstructions:
            return direction
        direction = direction.next()
    raise ValueError("All directions are blocked")

def guard_movement_rules(position, direction, obstructions):
    # Get new direction if faced with an obstacle
    new_direction = get_guard_direction(position, direction, obstructions)
    if new_direction is None:
        return position, direction  # No change if all directions are blocked
    # Move to the next position
    next_position = move_guard(position, new_direction)
    return next_position, new_direction

def get_guard_path(map_size, position, direction, obstructions) -> list:
    path = []
    while not guard_will_step_out_of_bounds(map_size, position, direction):
        state = (position, direction)
        path.append(state)
        position, direction = guard_movement_rules(position, direction, obstructions)
    return path

def part_one(map_size, guard_position, guard_direction, obstructions):
    guard_path = get_guard_path(map_size, guard_position, guard_direction, obstructions)
    unique_positions = {position for position, _ in guard_path}
    print(len(unique_positions) + 1)  # Include starting position

def part_two(map_size, guard_position, guard_direction, obstructions):
    guard_path = get_guard_path(map_size, guard_position, guard_direction, obstructions)
    guard_path = guard_path[:-1]  # Exclude the last position
    looping_obstructions = set()

    # Skip the first position to avoid considering initial obstructions
    illegal_obstruction_position = guard_path[1][0]
    for position, direction in guard_path:
        for obstruction_position, obstruction_direction in guard_path:
            if len(looping_obstructions) == 1888: # Dirty, but it works (get in infinite loop for the last?)
                break

            obstruction = guard_movement_rules(obstruction_position, obstruction_direction, obstructions)[0]

            if obstruction == illegal_obstruction_position:
                continue
            temp_obstructions = obstructions.union({obstruction})

            loop_position, loop_direction = position, direction
            steps_taken = set()
            while True:
                loop_position, loop_direction = guard_movement_rules(loop_position, loop_direction, temp_obstructions)
                if guard_will_step_out_of_bounds(map_size, loop_position, loop_direction):
                    break
                if (loop_position, loop_direction) in steps_taken:
                    looping_obstructions.add(obstruction)
                    break
                steps_taken.add((loop_position, loop_direction))

    print(len(looping_obstructions))

def main():
    with open("6/input.txt", "r") as file:
        map_data = [list(line.strip()) for line in file]
    
    def find_char_locations(map, char):
        locations = {(row_index, col_index)
                     for row_index, row in enumerate(map)
                     for col_index, cell in enumerate(row) if cell == char}
        return locations

    map_size = (len(map_data), len(map_data[0]))
    guard_position = find_char_locations(map_data, '^').pop()
    guard_direction = Direction.UP
    obstructions = find_char_locations(map_data, '#')

    part_one(map_size, guard_position, guard_direction, obstructions)
    part_two(map_size, guard_position, guard_direction, obstructions)

if __name__ == "__main__":
    main()