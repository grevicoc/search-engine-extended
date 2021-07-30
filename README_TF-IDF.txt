TF-IDF adalah teknik memvektorisasi kata-kata yang ada di dalam dokumen. TF sendiri merupakan kependekan dari Term Frequency yang memiliki arti menghitung banyaknya sebuah kata muncul di dalam dokumen. TF yang digunakan kali ini merupakan TF normalized atau artinya banyaknya sebuah kata akan dibagi dengan total jumlah kata di dokumen tersebut. 

TF(t,d) = countOf(t) / totalWord(d)

Selanjutnya sebelum masuk ke dalam IDF atau Inverse Document Frequency, perlu kita pahami Document Frequency terlebih dahulu. DF adalah banyaknya kemunculan kata t di dalam kumpulan dokumen. Kita tidak perlu mengetahui dalam sebuah dokumen kemunculan kata t ada berapa kali yang penting kata tersebut muncul di dokumen tersebut.

DF(t) = kemunculan t di dalam kumpulan dokumen

Lalu apa hubungannya dengan IDF? Menarik kesimpulan sebelumnya berarti semakin sering muncul kata di dalam kumpulan dokumen maka DF kata tersebut semakin besar. Dengan menghitung IDF atau inverse dari DF kata tersebut maka kita dapat mengetahui seberapa "pasaran" kata tersebut. Karena cara menghitung IDF adalah

IDF(t) = log(totalDocuments / DF(t)+1)

Sehingga semakin "pasaran" suatu kata akan semakin tidak bernilai kata tersebut. Hal ini cukup ampuh untuk menghilangkan stopwords atau kata-kata mainstream seperti "is", "are", dll dalam bahasa inggris. Mengapa perlu ditambah 1 pada pembaginya, hal ini untuk menghindari kasus tidak ada kata di dalam kumpulan dokumen sehingga pembagi bernilai 0.

Terakhir dari hasil TF dan IDF tersebut akan dikalikan sehingga mendapatkan nilai TF-IDF untuk sebuah kata. Nantinya setiap kata akan memiliki nilai sendiri-sendiri dan akan membentuk vektor. Vektor ini selanjutnya dapat diolah untuk mencari nilai cosine similarity.

Kelebihan dari teknik TF-IDF dalam vektorisasi kata adalah mereka secara tidak langsung meng-handle kasus stopwords sehingga jika querypun berisi stopwords, secara otomatis kata dari query tersebut tidak bernilai apa-apa  (nilai IDFnya 0 ya otomatis TF-IDFnya 0). Sedangkan kelemahan dari TF-IDF ini adalah tidak meng-handle kasus kata-kata yang memiliki sinonim. Misal kata "senang" pada query akan dianggap tidak ada pada dokumen jika di dokumen tersebut tidak ada kata "senang" (padahal ada kemungkinan terdapat kata yang mirip seperti "bahagia", "ceria", dll)