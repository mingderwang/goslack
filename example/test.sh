#!/bin/bash
for i in `seq 1 10000`; 
do
    curl -X POST -d @test.json --header "content-type: application/json" http://localhost:8091/githubUser
done
