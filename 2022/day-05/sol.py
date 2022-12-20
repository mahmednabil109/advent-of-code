import re

with open('input1') as input1:
    stacks, moves = input1.\
            read().\
            split('\n\n')
    st = list(
        map(
            lambda _: [],
            filter(lambda x: len(x.strip()) != 0, stacks.split('\n')[-1])
        )
    )
    for line in stacks.split('\n')[:-1][::-1]:
        for i in range(len(st)):
            top = line[i*4: i*4+3]
            if top.strip() != "":
                st[i].append(top[1])

    for move in moves.split('\n'):
        if not move: continue
        num, s1, s2 = map(
            int,
            re.match(r'move (\d+) from (\d+) to (\d+)', move).groups()
        )
        # first part
        #for _ in range(num):
        #   st[s2-1].append(st[s1-1].pop())
        # second part
        st[s2-1].extend(st[s1-1][-1*num:])
        st[s1-1] = st[s1-1][:-1*num]
    print(''.join(map(lambda x: '' if len(x) ==0 else x[-1], st)))

