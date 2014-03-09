wombat-desktop
================



What is Wombat?
---------------
Wombat is an end-to-end encrypted file sync service for everyone. This is the desktop application that watches directory, encrypts and synchronizes files with a remote server. Since the encryption is done client-side the server would never see your actual files.



Running:
--------
Built with [fsmonitor](https://github.com/howeyc/fsnotify) for cross platform support.


Currently only testing on linux.


change the values in testconf.json (although it has the WatchDirs in an array, it only currently accepts one value).

run the following after go install:
   
    wombat-desktop
