#!/bin/bash
find -name "*.sh" -printf "%f\n" | cut -d "/" -f1 | cut -d "." -f1
