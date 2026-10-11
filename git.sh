#!/bin/bash

echo " Git Helper Script"
echo "1) Create repo and push"
echo "2) Push code"
echo "3) Git status"
echo "4) Pull latest changes"
echo "5) Create new branch"
echo "6) Switch branch"
echo "7) Show branches"
echo "8) Show commit log"
echo "9) Clone repository"
echo "10) Stash changes"
echo "11) Apply stashed changes"
echo "12) Reset last commit"
echo "13) Fetch remote changes"
echo "14) Show remotes"
echo "15) Show git config"
echo "16) Exit"

# chmod +x git.sh

read -p "Enter your choice: " choice

case "$choice" in
    1)
        read -p "Enter repository URL: " repo_url
        git add .
        git commit -m "changes"
        git remote add origin "$repo_url"
        git push -u origin master
        ;;

    2)
        git add .
        git commit -m "changes"
        git push
        ;;

    3)
        git status
        ;;

    4)
        git pull
        ;;

    5)
        read -p "Branch name: " branch
        git checkout -b "$branch"
        ;;

    6)
        read -p "Branch name: " branch
        git checkout "$branch"
        ;;

    7)
        git branch -a
        ;;

    8)
        git log --oneline --graph --decorate --all
        ;;

    9)
        read -p "Repository URL: " repo_url
        git clone "$repo_url"
        ;;

    10)
        read -p "Stash message: " msg
        git stash push -m "$msg"
        ;;

    11)
        git stash pop
        ;;

    12)
        git reset --soft HEAD~1
        ;;

    13)
        git fetch --all
        ;;

    14)
        git remote -v
        ;;

    15)
        git config --list
        ;;

    16)
        echo "Goodbye!"
        exit 0
        ;;

    *)
        echo "Invalid choice."
        exit 1
        ;;
esac
