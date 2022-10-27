set -e
pwd=$(pwd)
cd ./pcre2-10.40 && ./configure --prefix="${pwd}"/pcre && make && make install