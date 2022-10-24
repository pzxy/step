#include <stdio.h>
#include <stdlib.h>

#define DIM(x) (sizeof(x)/sizeof((x)[0]))

static int cmp(const void* a, const void* b) {
    const int* pa = (int*)a;
    const int* pb = (int*)b;
    return *pa - *pb;
}

int main() {
    int values[] = { 42, 8, 109, 97, 23, 25};
    int i;

    qsort(values, DIM(values), sizeof(values[0]), cmp);

    for(i = 0; i < DIM(values); i++) {
        printf (" %d",values[i]);
    }
    return 0;
}
/**
void qsort(
    void* base, size_t num, size_t size,
    int (*cmp)(const void*, const void*)
);
 其中 base 参数是要排序数组的首个元素的地址，num 是数组中元素的个数，size 是数组中每个元素的大小。最关键是 cmp 比较函数，用于对数组中任意两个元素进行排序。cmp 排序函数的两个指针参数分别是要比较的两个元素的地址，如果第一个参数对应元素大于第二个参数对应的元素将返回结果大于 0，如果两个元素相等则返回 0，如果第一个元素小于第二个元素则返回结果小于 0。
*/