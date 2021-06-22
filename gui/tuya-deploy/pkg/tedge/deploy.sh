#!/bin/bash

tar -zxvf edgex_v1.5.0_Darwin_x86_64.tar.gz
chmod +x tedge
./tedge init -c ./encrypt-config -v 1.4.0
