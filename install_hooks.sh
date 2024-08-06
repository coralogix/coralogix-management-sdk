#!/bin/sh
sudo chmod +x ./hooks/pre-commit
cp ./hooks/pre-commit .git/hooks/pre-commit

