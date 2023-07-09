#!/bin/bash

cd server/view/styles
npx tailwindcss -i ./input.css -o ./output.css --watch
