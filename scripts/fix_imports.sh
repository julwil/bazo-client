#!/bin/bash

# Ask the user for their github username
echo This script will fix all imports related to bazo-client package.
echo Plese enter your github username.

read USERNAME

echo Fixing imports...


# Replace all bazo-related imports with the provided username in the entire scope of the project
for file in $(find ../ -name '*.go' -o -name '*.mod' -o -name '*.sum' -o -name '.travis.yml'); do
    sed -i -e 's@"github.com/.*/bazo-@"github.com/'${USERNAME}'/bazo-@g' $file
    echo $file
done

echo
echo Done