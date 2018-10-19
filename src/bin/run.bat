@echo off
start server.exe -logic
ping -n 2 127.0.0.1>nul
start server.exe -gate
