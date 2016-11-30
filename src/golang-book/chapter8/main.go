/*
Playing around with core Packages. Also contains solutions to all programming exercises.
*/

package main

import (
	"container/list"
	"crypto/sha1"
	"fmt"
	"hash/crc32"
	"io"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

func main() {

	//strings
	fmt.Println(strings.Contains("test", "es"))
	fmt.Println(strings.Count("test", "t"))
	fmt.Println(strings.HasPrefix("test", "te"))
	fmt.Println(strings.HasSuffix("test", "st"))
	fmt.Println(strings.Index("test", "e"))
	fmt.Println(strings.Join([]string{"a", "b"}, "-"))
	fmt.Println(strings.Repeat("a", 5))
	fmt.Println(strings.Replace("aaaaa", "a", "b", 3))
	fmt.Println(strings.Split("a-b-c-d", "-"))
	fmt.Println(strings.ToLower("TEST"))
	fmt.Println(strings.ToUpper("test"))
	arr := []byte("test")
	fmt.Println(arr)
	str := string([]byte{'t', 'e', 's', 't'})
	fmt.Println(str)

	//files and folders
	file, err := os.Create("test.txt")
	if err != nil {
		return
	}

	file.WriteString("Consistent hardwork and grit will make you invincible")
	file.Close()
	filePracticeOne()
	filePracticeTwo()
	filePracticeThree()

	//List
	var x list.List
	x.PushBack(1)
	x.PushBack(2)
	x.PushBack(3)

	for e := x.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(int))
	}

	//Sort   :-> len, Less, Swap
	kids := []Person{
		{"Jill", 9},
		{"Jack", 10},
	}
	sort.Sort(ByName(kids))
	fmt.Println(kids)

	sort.Sort(ByAge(kids))
	fmt.Println(kids)

	//Hashes and Cryptography
	//non-cryptographic hash functions: adler32, crc64, crc32, and fnv
	hashPracticeOne()
	hashPracticeTwo()

	crypto := sha1.New()
	crypto.Write([]byte("testing"))
	s := crypto.Sum([]byte{})
	fmt.Println(s)
}

//create hasher
func hashPracticeOne() {

	h := crc32.NewIEEE()
	//write data to hasher
	h.Write([]byte("test"))
	//Calculate crc32 checksum
	v := h.Sum32()
	fmt.Println(v)

	//Common use of crc32 is to compare two files
	//if two hashes are the same, files may be the same
	//if two hashes are different, files are definitely different
	file, err := os.Create("test2.txt")
	file2, err2 := os.Create("test3.txt")
	if err != nil || err2 != nil {
		return
	}
	file.WriteString("Consistent hardwork and grit will make you invincible")
	file2.WriteString("abracadabra")
	file.Close()
	file2.Close()
}

func hashPracticeTwo() {
	getHash := func(filename string) (uint32, error) {
		f, err := os.Open(filename)
		if err != nil {
			return uint32(0), nil
		}
		defer f.Close()
		h := crc32.NewIEEE()
		_, err1 := io.Copy(h, f)
		if err1 != nil {
			return uint32(0), nil
		}
		return h.Sum32(), nil
	}

	h1, err := getHash("test.txt")
	h2, err2 := getHash("test2.txt")
	h3, err3 := getHash("test3.txt")

	if err != nil || err2 != nil || err3 != nil {
		fmt.Println("Couldn't retrieve one of the files")
		return
	}

	fmt.Println("test and test2 are the same: {should be true}", h1 == h2)
	fmt.Println("test2 and test3 are the same: {should be false}", h2 == h3)
}

type Person struct {
	Name string
	Age  int
}

type ByName []Person

func (ps ByName) Len() int {
	return len(ps)
}

func (ps ByName) Less(i, j int) bool {
	return ps[i].Name < ps[j].Name
}

func (ps ByName) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

type ByAge []Person

func (this ByAge) Len() int {
	return len(this)
}

func (this ByAge) Less(i, j int) bool {
	return this[i].Age < this[j].Age
}

func (this ByAge) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func filePracticeOne() {
	file, err := os.Open("test.txt")
	if err != nil {
		return
	}
	defer file.Close()

	//get file size
	stat, err := file.Stat()
	if err != nil {
		return
	}

	//read file
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		return
	}

	str := string(bs)
	fmt.Println(str)
}

func filePracticeTwo() {
	bs, err := ioutil.ReadFile("test.txt")
	if err != nil {
		return
	}

	str := string(bs)
	fmt.Println(str)
}

func filePracticeThree() {
	dir, err := os.Open(".")
	if err != nil {
		return
	}
	defer dir.Close()

	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		return
	}

	for _, file := range fileInfos {
		fmt.Println(file.Name())
	}
}
