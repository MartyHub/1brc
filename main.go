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

const (
	dot   = byte('.')
	eol   = byte('\n')
	minus = byte('-')
	sep   = byte(';')
	zero  = byte('0')

	sizeChunk    = 1024 * 512
	sizeStations = 512
)

func main() {
	run(os.Args[1], true, runtime.NumCPU()*2, sizeChunk)
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
		out[idx] = make(map[int]*station, sizeStations)

		go func() {
			defer wg.Done()

			parseLine(in, out[idx])
		}()
	}

	start := time.Now()

	for off := int64(0); ; {
		end := min(off+chunkSize, fileSize)

		for buf[end-1] != eol {
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
		fmt.Printf("Computed in %v, collected in %v, total in %v\n",
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

			for ; data[i] != sep; i++ {
				key = 31*key + int(data[i])
			}

			end = i
			i++

			if data[i] == minus {
				if data[i+2] == dot {
					val = -int(data[i+1]-zero)*10 + int(data[i+3]-zero)
					i += 5
				} else {
					val = -int(data[i+1]-zero)*100 + int(data[i+2]-zero)*10 + int(data[i+4]-zero)
					i += 6
				}
			} else if data[i+1] == dot {
				val = int(data[i]-zero)*10 + int(data[i+2]-zero)
				i += 4
			} else {
				val = int(data[i]-zero)*100 + int(data[i+1]-zero)*10 + int(data[i+3]-zero)
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
	stations := make(map[int]*station, sizeStations)

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
