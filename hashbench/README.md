# Results

~~~
$ go test -bench=.
 goos: darwin
 goarch: amd64
 pkg: github.com/8tomat8/sketches/hashbench
 BenchmarkMD5_16-12                           225           5357492 ns/op
 BenchmarkMD5_160-12                          216           5407450 ns/op
 BenchmarkMD5_1600-12                         217           5408127 ns/op
 BenchmarkMurmur_16-12                       1893            634950 ns/op
 BenchmarkMurmur_160-12                      1629            637577 ns/op
 BenchmarkMurmur_1600-12                     1887            669047 ns/op
 BenchmarkMinioBlake2b_16-12                  278           4023857 ns/op
 BenchmarkMinioBlake2b_160-12                 300           4140898 ns/op
 BenchmarkMinioBlake2b_1600-12                298           4015316 ns/op
 BenchmarkSHA256_16-12                        122           9699950 ns/op
 BenchmarkSHA256_160-12                       123           9890037 ns/op
 BenchmarkSHA256_1600-12                      121           9712701 ns/op
 PASS
 ok      github.com/8tomat8/sketches/hashbench   20.204s
~~~