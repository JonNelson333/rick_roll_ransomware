# rick_roll_ransomware

The way the current program is written, the user has the ability to chose encryption or decrytion, as well as specify the path to encrypt/decrypt. 

To use the program, import the main.go, go.mod, go.sum and index.html to the system and then run the main.go file. 

There is a known issue with "thekey2.key" not being properly replaced after a decryption has taken place and a new encryption starts. To overcome this, it is recommended to delete "thekey2.key" AFTER you have decrypted files that were encrypted with that key. 
