#include <stdio.h>
#include <stdlib.h>
#include <string.h>


typedef struct Trie {
    struct Trie* next[26];
    int value, end;
} Trie;

void trie_insert(Trie*, char*, int);
int trie_search(Trie*, size_t, char*, int*);

Trie* new_trie(char** keys, size_t keys_len) {
    Trie* trie = (Trie*) malloc(sizeof(Trie));
    trie->value = 0;
    trie->end = 0;
    for(int i=0; i < keys_len ; i++) {
        trie_insert(trie, keys[i], i);
    }
    return trie;
}

void trie_insert(Trie* trie, char* key, int value) {
    Trie* curr = trie;
    size_t key_len = strlen(key);
    for(int i=0; i< key_len; i++) {
        size_t idx = key[i] - 'a';
        if(curr->next[idx] == NULL) {
            Trie* t = (Trie *) malloc(sizeof(Trie));
            t->value = 0;
            t->end = 0;
            curr->next[idx] = t;
        }
        curr = curr->next[idx];
        if (i + 1 == key_len) {
            curr->end = 1;
            curr->value = value;
        }
    }
}

int trie_search(Trie *trie, size_t pos, char* key, int* out) {
    Trie *curr = trie;
    size_t key_len = strlen(key);
    for(size_t i=pos; i<key_len; i++) {
        size_t idx = key[i] - 'a';
        if(idx < 0 || idx > 25 || curr->next[idx] == NULL) return 0;
        curr = curr->next[idx];
        if(curr->end == 1){
            *out = curr->value;
            return 1;
        }
    }
    *out = curr->value;
    return curr->end;
}

int main(){
    FILE* file = fopen("../in.txt", "r");
    char line[128];
    int n, sum = 0;
    char *entries[10] = {
        "zero",
        "one",
        "two",
        "three",
        "four",
        "five",
        "six",
        "seven",
        "eight",
        "nine",  
    };
    Trie *trie = new_trie(entries, 10);
    while(n = fscanf(file, "%s", line) != EOF) {
        int fd = 0, ld = 0, d = 0, first = 1, ok = 0;
        for(size_t i=0; i < strlen(line); i++) {
            if(line[i] >= '0' && line[i] <= '9') {
                d = line[i] - '0';
                ok = 1;
            } else {
                ok = trie_search(trie, i, line, &d);
                // printf("%s %d %d \n", (char *)(line+i), ok, d);
            }
            if(ok) {
                fd = first ? d : fd;
                ld = d;
                first = 0;
            }
            ok = 0;
        }
        // printf("[%s] %d %d\n", line, fd, ld);
        sum += fd*10 + ld;
        memset(line, 0, n);
    }
    printf("%d \n", sum);
    fclose(file);
}