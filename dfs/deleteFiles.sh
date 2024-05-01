#!/bin/bash

directory=/bigdata/students/skalse

for file in "$directory"/*_chunk_*;do
    rm "$file"
done