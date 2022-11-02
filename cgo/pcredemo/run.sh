gcc main.c -I pcre/include/ -L pcre/lib/ -lpcre2-8
./a.out  "\d+" "2022 hello 2033 world"