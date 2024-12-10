package main

import (
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/elisaado/aoc-2024/day"
	_ "github.com/elisaado/aoc-2024/day10"
	_ "github.com/elisaado/aoc-2024/day9"
	_ "github.com/elisaado/aoc-2024/day8"
	_ "github.com/elisaado/aoc-2024/day7"
	_ "github.com/elisaado/aoc-2024/day6"
	_ "github.com/elisaado/aoc-2024/day5"
	_ "github.com/elisaado/aoc-2024/day4"
	_ "github.com/elisaado/aoc-2024/day3"
	_ "github.com/elisaado/aoc-2024/day2"
	_ "github.com/elisaado/aoc-2024/day1"
)

func main() {
	if len(os.Args[1:]) < 2 {
		printHelp()
		return
	}
	dayArg := os.Args[2]
	dayNum, err := strconv.Atoi(dayArg)
	if err != nil {
		printHelp()
		return
	}
	args := os.Args[2:]
	switch action := os.Args[1]; action {
	case "download":
		{
			if len(args) != 1 {
				printHelp()
				return
			}
			downloadInput(dayNum)
		}
	case "test":
		{
			if len(args) != 3 {
				printHelp()
				return
			}
			partArg := args[1]
			partNum, err := strconv.Atoi(partArg)

			if err != nil {
				printHelp()
				return
			}

			testArg := args[2]
			if testArg == "all" {
				testDayAll(dayNum, partNum)
			} else {
				testNum, err := strconv.Atoi(testArg)
				if err != nil {
					printHelp()
					return
				}

				testDaySingle(dayNum, partNum, testNum)
			}
		}
	case "run":
		{
			if len(args) != 2 {
				printHelp()
				return
			}

			partArg := args[1]
			partNum, err := strconv.Atoi(partArg)
			if err != nil {
				printHelp()
				return
			}

			runDay(dayNum, partNum)
		}
	case "init":
		{
			if len(args) != 1 {
				printHelp()
				return
			}

			// download input
			downloadInput(dayNum)

			// copy contents of dayTemplate to dayN
			err := os.CopyFS("day"+strconv.Itoa(dayNum), os.DirFS("dayTemplate"))
			if err != nil {
				println("Error copying day template")
				panic(err)
			}

			// replace all instances of $ with N
			dir, err := os.ReadDir("day" + strconv.Itoa(dayNum))
			if err != nil {
				println("Error reading day directory")
				panic(err)
			}

			for _, file := range dir {
				if file.IsDir() {
					continue
				}
				handle, err := os.Open("day" + strconv.Itoa(dayNum) + "/" + file.Name())
				if err != nil {
					println("Error opening file")
					panic(err)
				}
				defer handle.Close()

				data, err := io.ReadAll(handle)
				if err != nil {
					println("Error reading file")
					panic(err)
				}

				newData := strings.ReplaceAll(string(data), "$", strconv.Itoa(dayNum))

				err = os.WriteFile("day"+strconv.Itoa(dayNum)+"/"+file.Name(), []byte(newData), 0644)
				if err != nil {
					println("Error writing file")
					panic(err)
				}
			}

			// create day folder in tests for both parts
			os.MkdirAll("tests/day"+strconv.Itoa(dayNum)+"/part1/test1", 0755)
			os.MkdirAll("tests/day"+strconv.Itoa(dayNum)+"/part2/test1", 0755)

			// create input.txt in test1 folder
			os.Create("tests/day" + strconv.Itoa(dayNum) + "/part1/test1/input.txt")
			os.Create("tests/day" + strconv.Itoa(dayNum) + "/part2/test1/input.txt")

			// create expected.txt in test1 folder
			os.Create("tests/day" + strconv.Itoa(dayNum) + "/part1/test1/expected.txt")
			os.Create("tests/day" + strconv.Itoa(dayNum) + "/part2/test1/expected.txt")

			// insert import in main.go at line 10
			// i love self modifying code
			mainFile, err := os.OpenFile("main.go", os.O_RDWR, 0644)
			if err != nil {
				println("Error opening main.go")
				panic(err)
			}
			defer mainFile.Close()

			mainData, err := io.ReadAll(mainFile)
			if err != nil {
				println("Error reading main.go")
				panic(err)
			}

			lines := strings.Split(string(mainData), "\n")
			lines = append(lines[:10], append([]string{"\t_ \"github.com/elisaado/aoc-2024/day" + strconv.Itoa(dayNum) + "\""}, lines[10:]...)...)

			err = mainFile.Truncate(0)
			if err != nil {
				println("Error truncating main.go")
				panic(err)
			}

			_, err = mainFile.Seek(0, 0)
			if err != nil {
				println("Error seeking main.go")
				panic(err)
			}

			_, err = mainFile.Write([]byte(strings.Join(lines, "\n")))
			if err != nil {
				println("Error writing to main.go")
				panic(err)
			}

			println("Day " + strconv.Itoa(dayNum) + " initialized")
		}
	case "help":
		printHelp()
	default:
		printHelp()
	}

	// COULD: add submission action
}

func printHelp() {
	println("Usage: aoc-2024 <action> <day> [<part>] [<test>]")
	println("Actions: download, test, run, init, help")
	println("Example: aoc-2024 run 1 1")
	println("Input file is automatically read from inputs/<day>.txt")
}

func runDay(dayNum int, partNum int) {
	inputFile := "inputs/" + strconv.Itoa(dayNum) + ".txt"
	dat, err := os.ReadFile(inputFile)
	if err != nil {
		println("Error reading input file")
		panic(err)
	}

	input := string(dat)

	day := day.GetDay(dayNum)
	if day.Number == 0 {
		println("Day not found")
		return
	}

	var result string
	if partNum == 1 {
		result = day.Part1(input)
	} else if partNum == 2 {
		result = day.Part2(input)
	} else {
		printHelp()
		return
	}

	println("Result:")
	println(result)
}

func testDaySingle(dayNum int, partNum int, testNum int) {
	inputFile := "tests/day" + strconv.Itoa(dayNum) + "/part" + strconv.Itoa(partNum) + "/test" + strconv.Itoa(testNum) + "/input.txt"
	dat, err := os.ReadFile(inputFile)
	if err != nil {
		println("Error reading input file")
		panic(err)
	}

	input := string(dat)

	expectedFile := "tests/day" + strconv.Itoa(dayNum) + "/part" + strconv.Itoa(partNum) + "/test" + strconv.Itoa(testNum) + "/expected.txt"
	dat, err = os.ReadFile(expectedFile)
	if err != nil {
		println("Error reading expected file")
		panic(err)
	}

	expected := string(dat)

	var result string

	if partNum == 1 {
		result = day.GetDay(dayNum).Part1(input)
	} else if partNum == 2 {
		result = day.GetDay(dayNum).Part2(input)
	} else {
		printHelp()
		return
	}

	if result == expected {
		println("Test " + strconv.Itoa(testNum) + " passed")
		println("Got: " + result)
	} else {
		println("Test " + strconv.Itoa(testNum) + " failed")
		println("Expected: " + expected)
		println("Got: " + result)
	}
}

func testDayAll(dayNum int, partNum int) {
	testDir := "tests/day" + strconv.Itoa(dayNum) + "/part" + strconv.Itoa(partNum)
	files, err := os.ReadDir(testDir)
	if err != nil {
		println("Error reading test directory")
		panic(err)
	}

	for _, file := range files {
		if file.IsDir() {
			testNum, err := strconv.Atoi(strings.Split(file.Name(), "test")[1])
			if err != nil {
				println("Error reading test directory")
				panic(err)
			}
			testDaySingle(dayNum, partNum, testNum)
		}
	}
}

func downloadInput(dayNum int) {
	os.OpenFile(".session", os.O_RDONLY, 0644)
	dat, err := os.ReadFile(".session")
	datStr := strings.Split(string(dat), "\n")[1]
	if err != nil {
		println("Error reading session file, does it exist?")
		panic(err)
	}

	if datStr == "" {
		println("No session (on second line) found in .session file")
		return
	}

	inputFile := "inputs/" + strconv.Itoa(dayNum) + ".txt"
	out, err := os.Create(inputFile)
	if err != nil {
		println("Error creating output file for input.txt")
		panic(err)
	}
	defer out.Close()

	cookie := &http.Cookie{
		Name:     "session",
		Value:    datStr,
		Secure:   true,
		HttpOnly: true,
	}

	req, err := http.NewRequest("GET", "https://adventofcode.com/2024/day/"+strconv.Itoa(dayNum)+"/input", nil)
	if err != nil {
		println("Error preparing request")
		panic(err)
	}
	req.AddCookie(cookie)
	req.Header.Set("User-Agent", "elisaado/aoc-2024")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		println("Error downloading input file")
		panic(err)
	}
	if resp.StatusCode != 200 {
		println("Error downloading input file, status code: " + resp.Status)
		if resp.StatusCode == 404 {
			println("Are you sure the input for this day is available?")
		}
		if resp.StatusCode == 500 {
			println("cookie might be malformed")
		}
		return
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		println("Error reading response body")
		panic(err)
	}
	out.Write(data)

}
