## 1.	Pada materi ini membahas mengenai clean and hexagonal architecture.  konsep clean architecture, setiap komponen yang ada bersifat independen dan tidak bergantung pada library external yang spesifik. Seperti tidak tergantung pada spesifik framework atau tidak bergantung pada spesifik database yang dipakai. 
## 2.	Context dalam clean and hexagonal architecture adalah suatu package yang membawa deadline, concellation signal, atau value lain pada request/permintaan API. Implementasian context kita dapat membuat root context dengan fungsi background. Fungsi background sendiri akan membalikan root context dimana kita bisa memakainya untuk operasi lain.  
## 3.	Dalam clean arsitektur terdapat beberapa layer pada setiap domainnya, berupa ; models, repository, usecase, dan delivery.

