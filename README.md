# Store Sensitive Data Securely on Users' Machines

This document outlines different methods for securely storing sensitive data on users' machines.

### 1. Store Data Using Platform-Specific APIs with the help of [Go-Keyring](https://github.com/zalando/go-keyring)

- **Windows**: Credential Manager
- **macOS**: Keychain
- **Linux**: GNOME Keyring

### 2. Store Data in the File System

- Encrypt the data using AES encryption before storing it in the file system.

### 3. Use LevelDB

- Store sensitive data using [LevelDB](https://github.com/google/leveldb), a fast key-value storage library.

### 4. Use SQLite3

- Store data in a local SQLite3 database.
