type Solution struct{
}

func (s *Solution) Encode(strs []string) string {
	res := ""
	for _, val := range strs{
		lenght := len(val)
		res = res + strconv.Itoa(lenght) + "#" + val
	}
	return res
}

func (s *Solution) Decode(encoded string) []string {
	res := []string{}

	i := 0
	for i < len(encoded){
		j := i

		for encoded[j] != '#' {
			j++
		}
		
		length, _ := strconv.Atoi(encoded[i:j])
		j++

		str := encoded[j : j+length]
		res = append(res, str)

		i = j + length
	}
	return res
}