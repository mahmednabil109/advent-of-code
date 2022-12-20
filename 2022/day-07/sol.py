class Node():
    def __init__(self, name, is_dir, size):
        self.name = name
        self.is_dir = is_dir
        self.size = size
        self.children = dict()
    def __str__(self):
        return f"{self.name}<{self.size}>"
    def __repr__(self):
        return f"{self.name}<{self.size}>"

def calc_size(root, dirs):
    if not root.is_dir:
        return root.size
    s = 0
    for file in root.children.values():
        s += calc_size(file, dirs)
    root.size = s
    dirs.append(root)
    return s

with open('input1') as input1:
    actions = input1.read().strip().split('\n')
    
    path_stack = []
    for action in actions:
        match action.split(' '):
            case ['$', 'ls']: continue
            case ['$', cmd, file]:
                if cmd == 'cd' and file == '..':
                    if len(path_stack) > 1: path_stack.pop()
                else:
                    if len(path_stack) == 0:
                        path_stack.append(Node('/', True, 0))
                    else:
                        path_stack.append(
                            path_stack[-1].children[file]
                        )
            case ['dir', file]:
                if file in ['.', '..']: continue
                path_stack[-1].children[file] = Node(file, True, 0)
            case [size, file]:
                path_stack[-1].children[file] = Node(file, False, int(size))
        
    dirs = []
    calc_size(path_stack[0], dirs)
    print(
        "part 1",
        sum(filter(lambda x: x <= 100000, map(lambda x: x.size, dirs)))
    )
    
    needed_space = 30000000 - (70000000 - dirs[-1].size)
    if needed_space <= 0:
        print("<$>")
    else:
        print(
            "part 2",
            next(
                filter(lambda x: x[0] >= 0, sorted(map(lambda x: (x.size - needed_space, x), dirs), key=lambda x: x[1].size))
            )[1].size
        )
                

