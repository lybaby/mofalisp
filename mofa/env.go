
package mofa

type Env struct {
	data map[string]*AstNode
}

func NewEnv()*Env {
	env:=&Env{}
	env.data = make(map[string]*AstNode)
	return env
}

func (env *Env)Find(varname string)(*AstNode, bool) {
	val, ok:= env.data[varname]
	return val, ok
}

func (env *Env)Setq(varname string, val *AstNode) {
	env.data[varname] = val
}

