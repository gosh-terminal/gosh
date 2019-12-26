#!/bin/bash

function checks() {
    echo -e "\nStarting checks..."
    if [ command -v sudo 2>/dev/null ]; then
        echo "sudo not installed on main system aborting"
        exit 1
    fi
    if [ ! -d ${HOME}/.gosh ]; then
        echo "Directory: ${HOME}/.gosh does not already exist. Creating..."
        mkdir ${HOME}/.gosh
        if [ command -v go 2>/dev/null ]; then
            echo "Go is not yet installed. Installing..."
            if which apt-get >/dev/null 2>&1; then
                sudo apt-get && sudo apt-get install -y golang
            if which apt >/dev/null 2>&1; then
                sudo apt && sudo apt install -y golang
            elif which dnf >/dev/null 2>&1; then
                sudo dnf install -y golang
            elif which yum >/dev/null 2>&1; then
                sudo rpm --import https://mirror.go-repo.io/centos/RPM-GPG-KEY-GO-REPO
                curl -s https://mirror.go-repo.io/centos/go-repo.repo | sudo tee /etc/yum.repos.d/go-repo.repo
                sudo yum install golang
            elif which apk >/dev/null 2>&1; then
                sudo apk add --no-cache --virtual .build-deps bash gcc musl-dev openssl go 
            elif which brew >/dev/null 2>&1; then
                brew install golang
            fi
        fi
    else
        echo "Directory: (${HOME}/.gosh) already exists!"
    fi
    echo ""
}
checks