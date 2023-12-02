#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int main(){
    FILE* file = fopen("../in.txt", "r");
    char line[128];
    int n, sum = 0;
    while(n = fscanf(file, "%s", line) != EOF) {
        int fd = 0, ld = 0, first = 1;
        for(int i=0; i < strlen(line); i++)
            if(line[i] >= '0' && line[i] <= '9') {
                if(first) {
                    fd = line[i] - '0';
                    first = 0;
                }
                ld = line[i] - '0';
            }
        sum += fd*10 + ld;
        memset(line, 0, n);
    }
    printf("%d \n", sum);
    fclose(file);
}