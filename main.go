package main

import (
	"fmt"
	"os"
	"runtime"
	"slices"
	"sync"
	"syscall"
	"time"
)

func main() {
	run(os.Args[1], true, runtime.NumCPU()*2, 1024*512)
}

func run(fileName string, print bool, workers int, chunkSize int64) {
	if fileName == "" {
		fileName = "measurements_100_000_000.txt"
	}

	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	fi, err := file.Stat()
	if err != nil {
		panic(err)
	}

	fileSize := fi.Size()

	buf, err := syscall.Mmap(int(file.Fd()), 0, int(fileSize), syscall.PROT_READ, syscall.MAP_SHARED)
	if err != nil {
		panic(err)
	}

	defer func() {
		_ = syscall.Munmap(buf)
	}()

	in := make(chan []byte)
	out := make([]map[int]*station, workers)

	var wg sync.WaitGroup

	for i := 0; i < workers; i++ {
		idx := i

		wg.Add(1)
		out[idx] = make(map[int]*station, 512)

		go func() {
			defer wg.Done()

			parseLine(in, out[idx])
		}()
	}

	start := time.Now()

	for off := int64(0); ; {
		end := min(off+chunkSize, fileSize)

		for buf[end-1] != '\n' {
			end--
		}

		in <- buf[off:end]

		if end == fileSize {
			break
		}

		off = end
	}

	close(in)
	wg.Wait()

	computed := time.Now()

	output(out, print)

	if print {
		fmt.Printf("Computed in %v, printed in %v, done in %v\n",
			computed.Sub(start),
			time.Since(computed),
			time.Since(start),
		)
	}
}

func parseLine(in <-chan []byte, stations map[int]*station) {
	var key, start, end, val int

	for data := range in {
		for i := 0; i < len(data); {
			start = i
			key = int(data[i])
			i++

			for ; data[i] != ';'; i++ {
				key = 31*key + int(data[i])
			}

			end = i
			i++

			if data[i] == '-' {
				if data[i+2] == '.' {
					val = -int(data[i+1]-'0')*10 + int(data[i+3]-'0')
					i += 5
				} else {
					val = -int(data[i+1]-'0')*100 + int(data[i+2]-'0')*10 + int(data[i+4]-'0')
					i += 6
				}
			} else if data[i+1] == '.' {
				val = int(data[i]-'0')*10 + int(data[i+2]-'0')
				i += 4
			} else {
				val = int(data[i]-'0')*100 + int(data[i+1]-'0')*10 + int(data[i+3]-'0')
				i += 5
			}

			stn := stations[key]

			if stn == nil {
				stations[key] = &station{
					city:  string(data[start:end]),
					count: 1,
					sum:   val,
					min:   val,
					max:   val,
				}
			} else {
				stn.add(val)
			}
		}
	}
}

func output(outs []map[int]*station, print bool) {
	stations := make(map[int]*station)

	for _, out := range outs {
		for key, stn := range out {
			cur := stations[key]

			if cur == nil {
				stations[key] = stn
			} else {
				cur.merge(stn)
			}
		}
	}

	msgs := make([]string, 0, len(stations))

	for _, stn := range stations {
		msgs = append(msgs, stn.String())
	}

	slices.Sort(msgs)

	if !print {
		return
	}

	fmt.Print("{")

	for i, stn := range msgs {
		if i > 0 {
			fmt.Print(", ")
		}

		fmt.Print(stn)
	}

	fmt.Print("}\n")
	fmt.Printf("%d stations\n", len(msgs))
}
