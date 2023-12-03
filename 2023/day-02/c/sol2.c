#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/param.h>

typedef struct darray {
	char** internal;
	size_t len, cap;
} darray;

darray* darray_init(int len, int cap) {
	darray* arr = (darray*) malloc(sizeof(darray));
	arr->internal = (char**) malloc(sizeof(char*) * cap);
	arr->cap = cap;
	arr->len = 0;
	return arr;
}

int darray_append(darray* arr, char* entry) {
	if(arr->len == arr->cap) {
		size_t old_cap = arr->cap;
		char** old = arr->internal;
		arr->cap = 2 * arr->cap;
		arr->cap = arr->cap == 0 ? 2 : arr->cap;
		arr->internal = (char**) malloc(sizeof(char*) * (arr->cap));
		if(arr->internal == NULL) {
			return -1;
		}
		memcpy(arr->internal, old, sizeof(char*)*(old_cap));
		free(old);
	}

	arr->internal[arr->len++] = entry;
	return 0;
}

int darray_index(darray* arr, size_t idx, char** out) {
	if(idx > arr->len) return 0;
	*out = arr->internal[idx];
	return 1;
}

int darray_print(darray* arr) {
	printf("[");
	for(size_t i=0; i<arr->len; i++) {
		printf("%s", arr->internal[i]);
		if(i+1 != arr->len)
			printf("|");
	}
	printf("]\n");
}


int index_of(char* str, char ch) {
	for(size_t i=0; i<strlen(str); i++)
		if (str[i] == ch)
			return i;
	return -1;
}

void str_trim(char* str) {
	int st = 0, i = 0, len = strlen(str);
	for (; st < len && (str[st] == ' ' || str[st] == '\n'); st++);
	for(i=0; st < len && i < len && st != 0; str[i++] = str[st++]);
	str[i] = st == 0 ? str[i] : '\0';
	for(i = len-1; i >= 0 && ( str[i] == ' ' || str[i] == '\n' || str[i] == '\0'); i--);
	str[i+1] = '\0';
}
// with no memory allocs for strs,
// we assume all strs are immutable.
// mmmh, acutally, forgot that we don't have fat_ptrs :(
// so we will have memory allocs for strs, jjjust for now
darray* str_split(char* str, char dilmeter) {
	darray* arr = darray_init(0, 0);
	for(size_t i=0; i<strlen(str); i++) {
		int offset = index_of((char *) (str+i), dilmeter);
		if (offset == -1) {
			offset = strlen(str) - i;
		}
		char* new_str = (char *) malloc(offset+1);
		new_str[offset] = '\0';
		strncpy(new_str, (char *) (str+i), offset);
		str_trim(new_str);
		darray_append(arr, new_str);
		i += offset;
	}
	return arr;
}

int main(){
	FILE *file = fopen("../in.txt", "r");

	char line[1024];
	int i=0, sum = 0, n = 0;
	while(n = fgets(line, sizeof line, file) != NULL) {
		int counts[3] = {0};
		int colon_idx = index_of(line, ':');
		darray* subs = str_split((char *) (line + colon_idx + 1), ';');
		for(int j = 0; j < subs->len; j++) {
			darray* cubes = str_split(subs->internal[j], ',');
			for(int k=0; k < cubes->len; k++) {
				darray* data = str_split(cubes->internal[k], ' ');
				int count = atoi(data->internal[0]);
				char* color = data->internal[1];
				switch (color[0]){
				case 'r':
					counts[0] = MAX(counts[0], count);
					break;
				case 'b':
					counts[1] = MAX(counts[1], count);
					break;
				case 'g':
					counts[2] = MAX(counts[2], count);
				}
			}
		}
		// printf("%s %d %d %d \n", line, counts[0], counts[1], counts[2]);
		sum += counts[0] * counts[1] * counts[2];
	}

	printf("%d\n", sum);
	fclose(file);	
	return 0;
}