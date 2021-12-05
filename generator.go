package codename

type Pool []string

type Generator struct {
	Pools []Pool
	Separator string
	Seed string
	CodenameLength int
}

func NewGenerator() (g Generator) {
	g.Separator = " "
	g.CodenameLength = 2
	return
}

func (s *Generator) Generate() (code_name string) {
	sep := ""
	for i := 0; i < s.CodenameLength; i++ {
		pool := i % len(s.Pools)
		word := s.seedToPoolResult(pool)
		code_name = code_name + sep + s.Pools[pool][word]
		sep = s.Separator
	}
	return
}

func (s *Generator) AddPool(p Pool) {
	s.Pools = append(s.Pools, p)
}

func (s *Generator) seedToPoolResult(pool_number int) (result int) {
    pool_len := len(s.Pools[pool_number])

    for _, char := range s.Seed {
        result = result + int(char)
    }
    if result > pool_len {
        result = result % pool_len
    }
    return
}