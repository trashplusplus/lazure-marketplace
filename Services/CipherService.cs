using System;
using System.IO;
using System.Security.Cryptography;
using System.Text;

public class CipherService
{
    private readonly byte[] encryptionKeyBytes;
    private readonly string _password;

    public CipherService(string encryptionKeyHex, string password)
    {
        // Convert the hexadecimal encryption key string to a byte array
        encryptionKeyBytes = Enumerable.Range(0, encryptionKeyHex.Length)
                             .Where(x => x % 2 == 0)
                             .Select(x => Convert.ToByte(encryptionKeyHex.Substring(x, 2), 16))
                             .ToArray();
        _password = password;

    }

    public string DecryptCipheredPassword(string cipherTextWithIvBase64)    //Base 64
    {
        var cipherTextWithIv = Convert.FromBase64String(cipherTextWithIvBase64);
        var iv = new byte[16];
        var cipherText = new byte[cipherTextWithIv.Length - iv.Length];

        Buffer.BlockCopy(cipherTextWithIv, 0, iv, 0, iv.Length);
        Buffer.BlockCopy(cipherTextWithIv, iv.Length, cipherText, 0, cipherText.Length);

        using (Aes aesAlg = Aes.Create())
        {
            aesAlg.Key = encryptionKeyBytes;
            aesAlg.IV = iv;
            aesAlg.Mode = CipherMode.CBC;
            aesAlg.Padding = PaddingMode.PKCS7;

            ICryptoTransform decryptor = aesAlg.CreateDecryptor(aesAlg.Key, aesAlg.IV);

            using (MemoryStream msDecrypt = new MemoryStream(cipherText))
            using (CryptoStream csDecrypt = new CryptoStream(msDecrypt, decryptor, CryptoStreamMode.Read))
            using (StreamReader srDecrypt = new StreamReader(csDecrypt))
            {
                return srDecrypt.ReadToEnd();
            }
        }
    }
    public bool ValidatePassword(string decryptedPassword)
    {
        return decryptedPassword.Equals(_password);
    }
}
