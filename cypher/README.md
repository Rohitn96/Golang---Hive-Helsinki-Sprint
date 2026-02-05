An interactive command-line utility designed to encrypt and decrypt messages using various cryptographic methods. The tool ensures data integrity by preserving non-alphabetic  <br /> characters and provides a robust user interface with persistent input validation.


1. ROT13 (Rotate by 13)<br />
ROT13 is a simple substitution cipher that replaces a letter with the 13th letter after it in the alphabet.<br />
How it works: Since the Latin alphabet has 26 letters, shifting a letter by 13 twice results in a full 360-degree rotation back to the original letter.<br /
Identity: This cipher is its own inverse; the encryption and decryption processes are identical.
Example: A ↔ N, hello ↔ uryyb.<br />
2. Reverse (Atbash Cipher)<br />
The Reverse cipher is a "mirror" encryption method originally used for the Hebrew alphabet.<br />
How it works: It reverses the alphabet so that the first letter (a) is replaced by the last letter (z), the second letter (b) by the second-to-last (y), and so on.<br />
Example: abc ↔ zyx, hello ↔ svool.<br />
3. Caesar Cipher<br />
The Caesar cipher is a type of substitution cipher where each letter in the plaintext is replaced by a letter some fixed number of positions down the alphabet.<br />
How it works: In this implementation, the shift is typically fixed (often 3 positions). To decrypt, the tool simply shifts the letters in the opposite direction.<br />
Example (Shift of 3): A → D, hello → khoor.<br />



Usage <br />
Start: Run the application to receive the welcome greeting. <br />
Select Operation: Choose between encrypt or decrypt. <br />
Choose Cipher: Select one of the three available encryption methods. <br />
Enter Message: Provide the text you wish to transform. <br />



Example Usage:  <br />
    go run main.go <br />
    Select operation (1/2): <br />
    1.Encrypt <br />
    2.Decrypt <br />
    1 <br />
    Select cypher (1/2): <br />
    1.Rot13. <br />
    2.Reverse. <br />
    3.Ceaser <br />
    1 <br />
    Your cypher is: Rot13 <br />
    Enter the message: <br />
    hello <br />
    Encrypted message using Rot13: uryyb        <br />                                   
    c3r3p2% go run main.go <br />
    Select operation (1/2): <br />
    1.Encrypt <br />
    2.Decrypt <br />
    2 <br />
    Select cypher (1/2):
    1.Rot13. <br />
    2.Reverse. <br />
    3.Ceaser <br />
    1 <br />
    Your cypher is: Rot13 <br />
    Enter the message: <br />
    uryyb <br />
    Decrypted message using Rot13: hello <br />


Note: The tool automatically trims whitespace from the start and end of your input. <br />
Note: If you provide an invalid option, the tool will re-prompt you until valid data is entered.
