package tag

type Repo string

func (r Repo) String() string {
	return string(r)
}
