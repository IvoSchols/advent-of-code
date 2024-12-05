import random
# Find number in rules
def find_number(rules, number):
    for parent, children in rules.items():
        if number in children:
            return parent
    return None


# Returns the inverted rules tree
def invert_rules(rules):
    inverted_rules = {}

    for parent, children in rules.items():
        for child in children:
            if child not in inverted_rules:
                inverted_rules[child] = []

            if parent not in inverted_rules:
                inverted_rules[parent] = []
                               
            inverted_rules[child].append(parent)
    
    return inverted_rules


def find_root(rules):
    for parent, children in rules.items():
        if len(children) == 0:
            return parent
    return None

def number_is_valid(remaining_page, number, inverted_rules):
    # Check if following numbers can be found in the previous numbers
    for following_number in remaining_page:
        if number not in inverted_rules[following_number]:
            return False
    return True


#
def part_one(rules, pages):
    inverted_rules = invert_rules(rules)

    score = 0

    for page in pages:
        is_valid = True

        for id, number in enumerate(page):
            # Check if following numbers can be found in the previous numbers
            remaining_page = page[id+1:]
            if number_is_valid(remaining_page, number, inverted_rules) == False:
                is_valid = False
                break
        
        if is_valid:
            middle_number = page[len(page) // 2]
            score += middle_number
  
    print(score)



def part_two(rules, pages):
    inverted_rules = invert_rules(rules)

    score = 0

    for page in pages:
        was_valid = True

        for id, number in enumerate(page):
            # Check if following numbers can be found in the previous numbers
            remaining_page = page[id+1:]
            if number_is_valid(remaining_page, number, inverted_rules) == False:
                was_valid = False
                # We will now swap till valid
                while number_is_valid(remaining_page, number, inverted_rules) == False:
                    random_index = random.randint(0, len(remaining_page) - 1)
                    random_number = remaining_page[random_index]
                    page[id], page[id+1+random_index] = random_number, number
                    remaining_page = page[id+1:]
                    number = random_number
        
        if was_valid == False:
            middle_number = page[len(page) // 2]
            score += middle_number
  
    print(score)


def parse_input(input):
    rules = parse_rules(input)
    pages = parse_pages(input)
    return rules, pages

def parse_rules(input):
    # A rule is a mapping of a number to a list of numbers
    rules = {}
    line = input.readline()

    while (line != None) and (line != "\n"):
        rule = line.split("|")

        left = int(rule[0])
        right = int(rule[1])

        if left not in rules:
            rules[left] = []

        if right not in rules:
            rules[right] = []
        
        rules[left].append(right)
        line = input.readline()

    return rules

def parse_pages(input):
    # Pages are a list of numbers
    pages = []
    line = input.readline()

    while line != '':
        numbers = line.split(',')
        pages.append([int(number) for number in numbers])
        line = input.readline()
    return pages

def main():
    file = open("5/input.txt", "r")
    
    # We build a mapping of the rules
    rules, pages = parse_input(file)
    file.close()
    
    part_one(rules, pages)
    part_two(rules, pages)

if __name__ == "__main__":
    main()