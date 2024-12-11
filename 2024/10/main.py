from enum import Enum

class Direction(Enum):
    UP = 1
    DOWN = 2
    LEFT = 3
    RIGHT = 4

MOVEMENTS = {
    Direction.UP: (-1, 0),
    Direction.RIGHT: (0, 1),
    Direction.DOWN: (1, 0),
    Direction.LEFT: (0, -1),
}

def out_of_bounds(x, y, topological_map):
    return x < 0 or y < 0 or x >= len(topological_map) or y >= len(topological_map[0])

def gradient(current, target):
    return target - current

def find_neighbors(x, y, topological_map):
    neighbors = []
    for direction in Direction:
        dx, dy = MOVEMENTS[direction]
        nx, ny = x + dx, y + dy
        if not out_of_bounds(nx, ny, topological_map) and gradient(topological_map[x][y], topological_map[nx][ny]) == 1:
            neighbors.append((nx, ny))
    return neighbors


# Part one: Find the number of trails that can be finished from a given trailhead
def part_one(topological_map, trailhead_locations):
    total = 0

    # We perform a breadth-first search from each trailhead, and count the number of trails that can be finished
    def bfs(x, y, topological_map):
        visited = set()
        queue = [(x, y)]
        count = 0
        while queue:
            current = queue.pop(0)
            if current in visited:
                continue
            visited.add(current)
            if topological_map[current[0]][current[1]] == 9:
                count += 1
                continue
            neighbors = find_neighbors(current[0], current[1], topological_map)
            queue.extend(neighbors)
        return count

    for trailhead in trailhead_locations:
        x, y = trailhead
        count = bfs(x, y, topological_map)
        total += count
    print(total)


def part_two(topological_map, trailhead_locations):
    total = 0

    # We perform a breadth-first search from each trailhead, and count the number of trails that can be finished
    def bfs(x, y, topological_map):
        queue = [(x, y)]
        count = 0
        while queue:
            current = queue.pop(0)
            if topological_map[current[0]][current[1]] == 9:
                count += 1
                continue
            neighbors = find_neighbors(current[0], current[1], topological_map)
            queue.extend(neighbors)
        return count

    for trailhead in trailhead_locations:
        x, y = trailhead
        count = bfs(x, y, topological_map)
        total += count
    print(total)

def main():
    topological_map = []
    trailhead_locations = []
    with open("./10/test_input.txt") as file:
        for i, line in enumerate(file):
            row = [int(digit) for digit in line.strip()]
            topological_map.append(row)
            for j, value in enumerate(row):
                if value == 0:
                    trailhead_locations.append((i, j))

    part_one(topological_map, trailhead_locations)
    part_two(topological_map, trailhead_locations)

if __name__ == "__main__":
    main()