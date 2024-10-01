#!/bin/sh
sudo chmod +x ./hooks/pre-commit
sudo chmod +x ./hooks/post-commit
cp ./hooks/pre-commit .git/hooks/pre-commit
cp ./hooks/post-commit .git/hooks/post-commit

