package behaviors

const ()

var TravelChain *Chain

func init() {
	TravelChain = NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		return nil, nil
	}))
	LoginChain.AddNextChain(NewChain(NewFuncNode(LoginAdmin)), 1)
}
