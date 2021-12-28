package aoc

import (
	"regexp"
	"strconv"

	"github.com/jaynak/aoc2021/pkg/util"
)

func Day24(path string) (int, int) {

	lines := util.ReadToStrings(path)

	inst := read_instruction_set(lines)

	max_model := find_valid_model(inst, []int{})

	// input := []int{1, 3, 5, 7, 9, 2, 4, 6, 8, 9, 9, 9, 9, 9}

	// reg := apply_monad_instructions(inst, input)

	// fmt.Println(reg)

	return max_model, 0
}

type monad_inst struct {
	op  int
	reg []int
	val int
}

func find_valid_model(instructions []*monad_inst, prev []int) int {

	if len(prev) == 14 {
		ret := apply_monad_instructions(instructions, prev)

		if ret[3] == 0 {
			// This is a valid model number
			str := ""
			for _, i := range prev {
				str += strconv.Itoa(i)
			}

			val, _ := strconv.Atoi(str)

			return val
		}

		return 0
	}

	// Not at the end yet
	nums := prev
	nums = append(nums, 0)
	for i := 9; i > 0; i-- {
		nums[len(nums)-1] = i
		ret := find_valid_model(instructions, nums)
		if ret != 0 {
			return ret
		}
	}

	// Fell through
	return 0
}

func apply_monad_instructions(instructions []*monad_inst, input []int) []int {

	if len(input) != 14 {
		panic("bad input")
	}

	registers := []int{0, 0, 0, 0}
	ptr := 0

	for _, inst := range instructions {
		num := inst.val

		if inst.op == 0 {
			num = input[ptr]
			ptr++
		} else {
			if len(inst.reg) == 2 {
				num = registers[inst.reg[1]]
			}
		}

		dest := inst.reg[0]

		// Do the op
		switch inst.op {
		case 0:
			registers[dest] = num
		case 1:
			registers[dest] += num
		case 2:
			registers[dest] *= num
		case 3:
			registers[dest] /= num
		case 4:
			registers[dest] = registers[dest] % num
		case 5:
			if registers[dest] == num {
				registers[dest] = 1
			} else {
				registers[dest] = 0
			}
		}
	}

	return registers
}

func read_instruction_set(lines []string) []*monad_inst {

	r := regexp.MustCompile("([a-z]+) ([w-z])( ([0-9a-z-]+))?")
	inst_map := map[string]int{
		"imp": 0, "add": 1, "mul": 2, "div": 3, "mod": 4, "eql": 5,
	}
	register_map := map[string]int{
		"w": 0, "x": 1, "y": 2, "z": 3,
	}

	instructions := []*monad_inst{}

	for _, line := range lines {
		m := r.FindAllStringSubmatch(line, -1)

		inst := &monad_inst{
			op:  inst_map[m[0][1]],
			reg: []int{register_map[m[0][2]]},
		}

		if inst.op != 0 {

			// Process the second variable
			switch m[0][4] {
			case "w", "x", "y", "z":
				inst.reg = append(inst.reg, register_map[m[0][4]])
			default:
				num, _ := strconv.Atoi(m[0][4])
				inst.val = num
			}

		}

		instructions = append(instructions, inst)
	}

	return instructions
}