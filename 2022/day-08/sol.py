dirs = [(0,1), (1,0), (0,-1), (-1,0)]
lmap = lambda f,i: list(map(f,i))

def dfs(grid, pos, d):
    x, y = pos
    t = grid[x][y]
    nx, ny = x + d[0], y + d[1]
    m, n = len(grid), len(grid[0])
    while nx >= 0 and nx < m  and ny >= 0 and ny < n:
        if grid[nx][ny] >= t:
            return False
        nx, ny = nx + d[0], ny + d[1]
    return True

with open('input1') as input1:
    trees = lmap(
        lambda x: lmap(int, list(x)),
        input1\
            .read()\
            .strip()\
            .split('\n')
    )
    visible_trees = (len(trees) - 1) * (len(trees[0]) - 1)
    print(visible_trees)
    for i in range(1, len(trees)-1):
        for j in range(1, len(trees[0])-1):
            for d in dirs:
                if dfs(trees, (i, j), d):
                    visible_trees += 1
                    break
    print(visible_trees)

