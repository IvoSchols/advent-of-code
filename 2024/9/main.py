# Define block type
class Block:
    def __init__(self, id, block_size):
        self.id = id
        self.block_size = block_size

    def __str__(self):
        return f"({self.id},{self.block_size})"

def expand_disk_map(disk_map, as_block=True):
    expanded_disk_map = []
    file_block_id = -1

    for i in range(0, len(disk_map), 2):
        block_size = disk_map[i]
        file_block_id += 1

        if as_block:
            expanded_disk_map.append(Block(file_block_id, block_size))

            if i + 1 < len(disk_map):

                free_space_size = disk_map[i + 1]
                expanded_disk_map.append(free_space_size)
        else:
            expanded_disk_map.extend([Block(file_block_id, 1)] * block_size)

            if i + 1 < len(disk_map):
                free_space_size = disk_map[i + 1]
                expanded_disk_map.extend([1] * free_space_size)


    return expanded_disk_map

def sort_disk_map(disk_map):

    for left_free_space_index in range(len(disk_map)):
        if isinstance(disk_map[left_free_space_index], Block):
            continue
        # Reverse search for rightmost block type
        right_block_index = next((len(disk_map) - 1 - index for index, element in enumerate(reversed(disk_map)) if isinstance(element, Block)),None)
 
        if left_free_space_index > right_block_index:
            break

        # Swap the block type with the free space
        disk_map[left_free_space_index], disk_map[right_block_index] = disk_map[right_block_index], disk_map[left_free_space_index]

    return disk_map

def calculate_disk_map_checksum(disk_map):
    checksum = 0
    position = 0
    for i in range(len(disk_map)):
        if isinstance(disk_map[i], int):
            position += disk_map[i]
            continue

        for _ in range(disk_map[i].block_size):
            checksum += disk_map[i].id * position
            position += 1


    return checksum

def part_one(disk_map):
    expanded_disk_map = expand_disk_map(disk_map, as_block=False)
    sorted_disk_map = sort_disk_map(expanded_disk_map)
    checksum = calculate_disk_map_checksum(sorted_disk_map)
    print(checksum)

def sort_disk_map_as_block(disk_map):
    # Remove all 0 free spaces
    disk_map = [element for element in disk_map if element != 0]
    print([str(element) for element in disk_map])

    right_blocks = [right_block for right_block in disk_map if isinstance(right_block, Block)]
    right_blocks.reverse()

    # Search for the leftmost free space (i.e. int type)
    for right_block in right_blocks:
        right_block_index = len(disk_map) - 1 - disk_map[::-1].index(right_block)


        left_free_space_indexes = [index for index, element in enumerate(disk_map) if isinstance(element, int)]



        for left_free_space_index in left_free_space_indexes:
            if left_free_space_index > right_block_index:
                continue
            # Check if the right block fits in the free space
            if disk_map[left_free_space_index] < right_block.block_size:
                    continue
            else:
                remaining_free_space = disk_map[left_free_space_index] - right_block.block_size

                if remaining_free_space == 0:
                    disk_map[left_free_space_index], disk_map[right_block_index] = right_block, disk_map[left_free_space_index]
                    break
                elif remaining_free_space > 0:
                    disk_map[left_free_space_index] = right_block
                    disk_map[right_block_index] = right_block.block_size
                    disk_map.insert(left_free_space_index + 1, remaining_free_space)
                    break
                else:
                    raise ValueError("Invalid disk map")


    return disk_map

def part_two(disk_map):
    expanded_disk_map = expand_disk_map(disk_map, as_block=True)
    sorted_disk_map = sort_disk_map_as_block(expanded_disk_map)
    checksum = calculate_disk_map_checksum(sorted_disk_map)
    print(checksum)

def main():
    # Read disk map
    with open("./9/input.txt", "r") as file:
        disk_map = [int(digit) for digit in file.read().strip()]
    part_one(disk_map)
    part_two(disk_map)

if __name__ == "__main__":
    main()