mkdir auliani_$(Sun_Mar_19_11:33 '+%a_%b_%d_%H:%M:%S_%Y')

cd tegar_$(date '+%a_%b_%d_%H:%M:%S_%Y')

#!/bin/bash

# Buat folder utama
mkdir auliani_$(date '+%a_%b_%d_%H:%M:%S_%Y')

# Masuk ke dalam folder baru
cd auliani_$(date '+%a_%b_%d_%H:%M:%S_%Y')

# Buat subfolder "about_me"
mkdir about_me

# Buat subfolder "personal" di dalam "about_me"
mkdir about_me/personal

# Buat file "facebook.txt" di dalam "about_me/personal"
echo "URL Facebook: <URL Facebook>\nUsername: <aaaulianis>" > about_me/personal/facebook.txt

# Buat subfolder "professional" di dalam "about_me"
mkdir about_me/professional

# Buat file "linkedin.txt" di dalam "about_me/professional"
echo "URL LinkedIn: <URL LinkedIn>\nUsername: <aulianisukma>" > about_me/professional/linkedin.txt

# Buat subfolder "my_friends"
mkdir my_friends

# Unduh daftar nama teman dari file yang diberikan
curl -o my_friends/list_of_my_friends.txt https://example.com/list_of_my_friends.txt

# Buat subfolder "my_system_info"
mkdir my_system_info

# Buat file "about_this_laptop.txt" dengan nama user dan hasil dari perintah "uname -a"
echo "User: $(whoami)\n$(uname -a)" > my_system_info/about_this_laptop.txt

# Buat file "internet_connection.txt" dengan hasil dari perintah "ping google.com" sebanyak 3 kali
ping -c 3 google.com > my_system_info/internet_connection.txt

# Tampilkan struktur direktori
tree "auliani_$(date '+%a_%b_%d_%H:%M:%S_%Y')"
