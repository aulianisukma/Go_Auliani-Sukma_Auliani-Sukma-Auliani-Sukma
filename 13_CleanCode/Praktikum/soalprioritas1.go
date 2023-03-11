package main

type user struct {
	id       int // Kekurangan yaitu tipe data username dan password sebaiknya menggunakan string.
	username int
	password int
}

type userservice struct {
	t []user // Kekurangannya yaitu nama variabel t sebaiknya lebih deskriptif, misalnya users atau dataUsers.
}

func (u userservice) getallusers() []user { // Kekurangannya yaitu penggunaan parameter receiver (u userservice) sebaiknya menggunakan pointer (*userservice) agar lebih efisien saat digunakan.
	return u.t
}

func (u userservice) getuserbyid(id int) user {
	for _, r := range u.t { //kekurangannya penggunaan variabel r sebaiknya diganti dengan nama yang lebih deskriptif, misalnya user.
		if id == r.id {
			return r
		}
	}

	return user{} //Seharusnya mengembalikan error atau pesan jika user tidak ditemukan.
}