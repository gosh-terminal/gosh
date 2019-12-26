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
            if [ command -v apt-get 2>/dev/null ]; then
                sudo apt-get && sudo apt-get install -y golang
            elif [ command -v apt 2>/dev/null ]; then
                sudo apt && sudo apt install -y golang
            elif [ command -v dnf 2>/dev/null ]; then
                sudo dnf install -y golang
            elif [ command -v rpm 2>/dev/null && command -v yum 2>/dev/null && command -v curl 2>/dev/null ]; then
                sudo rpm --import https://mirror.go-repo.io/centos/RPM-GPG-KEY-GO-REPO
                curl -s https://mirror.go-repo.io/centos/go-repo.repo | sudo tee /etc/yum.repos.d/go-repo.repo
                sudo yum install golang
            elif [ command -v apk 2>/dev/null ]; then
                sudo apk add --no-cache --virtual .build-deps bash gcc musl-dev openssl go 
            elif [ command -v brew 2>/dev/null ]; then
                brew install golang
            fi
        fi
        go build -o gosh src/*.go
        mv gosh "${HOME}/.gosh/gosh"
        touch history.txt "${HOME}/.gosh"
        mv history.txt "${HOME}/.gosh"
        cd "${HOME}/.gosh"
        echo "export PATH=${PATH}:${PWD}" >> ~/.bashrc
    else
        echo "Directory: (${HOME}/.gosh) already exists!"
    fi
    echo ""
}
checks
