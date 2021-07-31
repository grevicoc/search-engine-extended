Pagerank algorithm merupakan algoritma untuk menentukan "importance" sebuah page di dunia internet. Biasanya dalam sebuah page akan terdapat hyperlink-hyperlink yang menuju page lain dan hal inilah yang menjadi penentu "importance" sebuah page. Semakin banyak page yang mengacu (memiliki hyperlink) ke sebuah page, maka semakin tinggi pula "importance" page tersebut. Pagerank ada untuk membantu perhitungan tersebut dengan rumus dan contoh seperti ini:

Misal terdapat 4 page, yaitu A, B, C, dan D. Karena hanya terdapat 4 page maka masing-masing initial value page tersebut adalah 0.25 (1 dibagi 4). Selanjutnya, misal page B memiliki link ke page A dan C. Page C memiliki link ke page A. Page D memiliki link ke 3 page lainnya. Maka nantinya page A akan menerima semua initial value dari page C (0.25, karena page C hanya mempunyai link ke A) dan menerima 1/3 initial value dari page D (0.083, karena page D memiliki 3 link ke page lain) dan terakhir menerima 1/2 initial value dari page B (0.125, karena page B memiliki 2 link ke page lain). Jadi page A akan memiliki real value sebesar 0.458. Oleh karena itu rumus pagerank dapat ditulis menjadi:

PR(A) = PR(B)/2 + PR(D)/3 + PR(C)/1

atau secara general

PR(x) = jumlah(PR(y)/link(y))

dengan link(y) merupakan jumlah link di dalam page y.

Untuk pengimplementasian pada program ini, nilai pagerank akan dikalikan 0.1 dan ditambahkan ke hasil cosine-sim. Mengapa seperti itu, sebenarnya idealnya hasil pagerank dan cosine-sim dipisah sendiri-sendiri dan menjadi first dan second variabel perhitungan heapsort. Namun, agar lebih simpel maka dilakukan pengimplementasian seperti itu (nilai pagerank ikut diperhitungkan namun tidak merubah banyak nilai cosine-sim).

Kekurangan algoritma ini mungkin untuk website berita yang kadang link di dalam pagenya tidak berhubungan, lebih ke berita-berita terbaru, ter-hot, atau unik membuat fallacy importance. Karena menurutku pagerank ini kan niatnya membuat value sebuah page sebagai "sumber" dari page-page lain, namun kalo di website berita kadang page yang ga berhubungan aja bisa jadi punya link gara-gara hal tadi. Namun kelebihannya adalah bisa menjadi second metric dalam perhitungan relevansi. Ibaratnya kalo ada dua page yang relevansinya sama, namun pageranknya lebih besar yang satu maka yang akan diutamakan yang lebih besar tersebut.