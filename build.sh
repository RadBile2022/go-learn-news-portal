#!/bin/bash

# Buat file .netrc untuk autentikasi git supaya bisa akses private repo
echo "machine github.com login your_github_username password ${GITHUB_TOKEN}" > ~/.netrc
chmod 600 ~/.netrc

# Build aplikasi kamu, main.go sudah di root
go build -o app main.go
