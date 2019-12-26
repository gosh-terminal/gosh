#!/bin/bash

function install() {
    echo -e "\nStarting checks..."
    if [ command -v sudo 2>/dev/null ]; then
        echo "sudo not installed on main system aborting"
        exit 1
    fi
    if [ ! -d ${HOME}/.gosh ]; then
        echo "Directory: ${HOME}/.gosh does not already exist. Creating..."
        mkdir ${HOME}/.gosh
    else
        echo "Directory: (${HOME}/.gosh) already exists!"
    fi
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
    if [ command -v git 2>/dev/null ]; then
        echo "Git is not yet installed. Installing..."
        if [ command -v apt-get 2>/dev/null ]; then
            sudo apt-get && sudo apt-get install -y git
        elif [ command -v apt 2>/dev/null ]; then
            sudo apt && sudo apt install -y git
        elif [ command -v dnf 2>/dev/null ]; then
            sudo dnf install -y git
        elif [ command -v rpm 2>/dev/null && command -v yum 2>/dev/null && command -v curl 2>/dev/null ]; then
            sudo yum install git
        elif [ command -v apk 2>/dev/null ]; then
            sudo apk add --no-cache --virtual .build-deps bash git
        elif [ command -v brew 2>/dev/null ]; then
            brew install git
        fi
    fi
    git clone https://github.com/gosh-terminal/gosh/tree/install.git
    cd gosh
    echo "Installing dependencies"
    go get -v -t -d ./...
    echo "building gosh"
    go build -o gosh src/*.go
    echo "Moving files"
    mv gosh "${HOME}/.gosh"
    touch history.txt "${HOME}/.gosh"
    mv history.txt "${HOME}/.gosh"
    echo "Removing old stuff"
    cd "${HOME}/.gosh"
    echo "Setting ENV Vars"
    echo "Setting \$PATH"
    echo "export PATH=${PATH}:${PWD}" >> ~/.bashrc
    echo "Setting \$GOSH_HOME"
    echo "export GOSH_HOME=${PWD}" >> ~/.bashrc
    echo -e "\nDone!!\n\nPlease open a new terminal, or run the following in the existing one:\n    source ~/.bashrc\n\n"
}
install
