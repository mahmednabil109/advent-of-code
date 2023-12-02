from typing import List, Tuple

class Trie:
    """
    Trie datastructure used for fast string prefix search
    """
    def __init__(self, keys: List[str]= None):
        self._next = [None] * 26
        self.end = False
        self.value = 0
        if  keys:
            for i, key in enumerate(keys):
                self._insert(key, i)
    
    def _insert(self, key: str, value: int):
        curr = self
        for i, ch in enumerate(key):
            idx = ord(ch) - ord('a')
            if curr._next[idx] is None:
                curr._next[idx] = Trie()
            curr = curr._next[idx]
            if i == len(key) - 1:
                curr.end = True
                curr.value = value
    
    def search(self, key: str) -> Tuple[int, bool]:
        curr = self
        for ch in key:
            idx = ord(ch) - ord('a')
            if idx < 0 or idx > 25 or curr._next[idx] is None:
                return 0, False
            curr = curr._next[idx]
            if curr.end:
                return curr.value, True
        return curr.value, curr.end
    
with open('../in.txt') as file:
    trie = Trie([
        "zero",
        "one",
        "two",
        "three",
        "four",
        "five",
        "six",
        "seven",
        "eight",
        "nine"
    ])
    result = 0
    for line in file.readlines():
        if not line: continue;
        fd, ld, d, first, ok = 0, 0, 0, True, False
        for i, ch in enumerate(line):
            if ch.isdigit():
                d, ok = int(ch), True
            else:
                d, ok = trie.search(line[i:])

            if ok:
                if first:
                    fd, first = d, False
                ld = d
            ok = False
        # print(f"[{line.strip()}]: {fd} {ld}")
        result += fd * 10 + ld
    print(result)