package vpp

// Lines はブロック順にセリフ行を返す
func (v *VPP) Lines() []Line {
	var lines []Line
	for _, block := range v.Project.Blocks {
		narrator := block.Narrator.Key
		for _, sentence := range block.SentenceList {
			if sentence.Text == "" {
				continue
			}
			lines = append(lines, Line{
				Narrator: narrator,
				Text:     sentence.Text,
			})
		}
	}
	return lines
}
