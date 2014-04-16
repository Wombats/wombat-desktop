wombat-desktop
================


What is Wombat?
---------------
Wombat is an end-to-end encrypted file sync service for everyone. This is the desktop application that watches directory, encrypts and synchronizes files with a remote server. Since the encryption is done client-side the server would never see actual files.


Running:
--------
Built with [fsnotify](https://github.com/howeyc/fsnotify) for cross platform support.


Wombat is being developed on linux (Debian jessie x64) and is currently known to work on winows 7.


change the values in testconf.json.

compile it and run:

    wombat-desktop


API:
----
__This has not yet been implemented.__

For web and desktop clients there should be a singular route, '/api'. Data, specific methods, etc. should be called from within the request parameters.
  * Create
  * Delete
  * Modify
  * Move / Rename


        { 
          "method" : "create",
          "auth"   : "however we do auth (probably desktop only? Via the web a user would be auth'ed already.)",
          "data"   : "encrypedFileStuff",
        }
