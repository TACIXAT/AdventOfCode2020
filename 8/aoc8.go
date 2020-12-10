package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Operand int

const (
	NOP Operand = iota
	ACC
	JMP
)

type Instruction struct {
	Op  Operand
	Arg int
}

func getOp(op string) Operand {
	switch op {
	case "nop":
		return NOP
	case "acc":
		return ACC
	case "jmp":
		return JMP
	default:
		log.Fatal("unknown op", op)
	}

	return NOP
}

func (insn *Instruction) String() string {
	ops := map[Operand]string{
		NOP: "nop",
		ACC: "acc",
		JMP: "jmp",
	}

	return fmt.Sprintf("%s %d", ops[insn.Op], insn.Arg)
}

func NewInstruction(ln string) Instruction {
	toks := strings.Split(ln, " ")
	op := getOp(toks[0])
	arg, err := strconv.ParseInt(toks[1], 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	return Instruction{
		Op:  op,
		Arg: int(arg),
	}
}

func run(insns []Instruction) int {
	acc := 0
	idx := 0
	visited := make(map[int]struct{})
	empty := struct{}{}
	for {
		if _, ok := visited[idx]; ok {
			return acc
		}

		insn := insns[idx]
		log.Println(insn.String())

		visited[idx] = empty
		switch insn.Op {
		case NOP:
			idx += 1
		case ACC:
			acc += insn.Arg
			idx += 1
		case JMP:
			idx += insn.Arg
		}
	}

	return 0
}

func fork(visited map[int]struct{}) map[int]struct{} {
	state := make(map[int]struct{})
	for k, v := range visited {
		state[k] = v
	}

	return state
}

func debug(insns []Instruction) int {
	acc := 0
	idx := 0
	count := 0

	visited := make(map[int]struct{})
	empty := struct{}{}

	lastIdx := -1
	lastCount := -1
	lastAcc := 0
	active := false
	lastVisited := fork(visited)

	fmt.Println("len", len(insns))

	for {
		if idx == len(insns) {
			return acc
		}

		if _, ok := visited[idx]; ok || idx > len(insns) || idx < 0 {
			fmt.Println("# RESET", idx, ok, idx > len(insns))
			fmt.Println()
			idx = lastIdx
			acc = lastAcc
			visited = lastVisited
			count = lastCount
			active = false
		}

		insn := insns[idx]
		fmt.Println(count, idx, insn.String())

		visited[idx] = empty
		switch insn.Op {
		case NOP:
			if !active && count > lastCount {
				fmt.Println("# TREATED AS JMP")
				active = true
				lastIdx = idx
				lastAcc = acc
				lastVisited = fork(visited)
				lastCount = count
				idx += insn.Arg
			} else {
				idx += 1
			}
		case ACC:
			acc += insn.Arg
			idx += 1
		case JMP:
			if !active && count > lastCount {
				fmt.Println("# TREATED AS NOP")
				active = true
				lastIdx = idx
				lastAcc = acc
				lastVisited = fork(visited)
				lastCount = count
				idx += 1
			} else {
				idx += insn.Arg
			}
		}

		count++
	}

	return 0
}

func main() {
	file, err := os.Open("input8.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	insns := make([]Instruction, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ln := scanner.Text()
		if len(ln) == 0 {
			continue
		}

		insn := NewInstruction(ln)
		insns = append(insns, insn)
	}

	// LAST ONE??

	log.Println("out", run(insns))
	log.Println("out", debug(insns))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
