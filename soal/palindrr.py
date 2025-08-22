def palindrr(kata: str) -> bool:
    """
    Fungsi untuk memeriksa apakah sebuah kata adalah palindrom.
    
    Args:
    kata (str): Kata yang akan diperiksa.
    
    Returns:
    bool: True jika kata adalah palindrom, False jika tidak.
    """
    # Menghapus spasi dan mengubah huruf menjadi kecil
    kata = kata.replace(" ", "").lower()
    
    # Membandingkan kata dengan kebalikannya
    return kata == kata[::-1]

if __name__ == "__main__":
    # Contoh penggunaan
    print(palindrr("SalaS"))  # Output: True
    print(palindrr("Hello"))   # Output: False