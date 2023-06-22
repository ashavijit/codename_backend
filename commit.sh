#!/bin/bash

check_git_installed() {
    git --version >/dev/null 2>&1
    if [ $? -ne 0 ]; then
        echo "Git is not installed. Please install Git and try again."
        exit 1
    fi
}

check_git_repository() {
    git rev-parse --is-inside-work-tree >/dev/null 2>&1
    if [ $? -ne 0 ]; then
        echo "Not a Git repository. Please navigate to a Git repository and try again."
        exit 1
    fi
}

check_git_changes() {
    git diff --quiet --exit-code
    if [ $? -eq 1 ]; then
        echo "There are changes in the repository."
        read -p "Please enter a commit message: " commit_message
        git commit -am "$commit_message"
        echo "Changes committed successfully."

        read -p "Do you want to push the changes? [y/n]: " push_option
        case "$push_option" in
            [Yy])
                git push
                echo "Changes pushed successfully."
                ;;
            *)
                echo "Changes are not pushed."
                ;;
        esac
    else
        echo "No changes detected in the repository."
    fi
}

check_git_installed
check_git_repository
check_git_changes
