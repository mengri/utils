package autowire

type OnCreate interface {
	OnCreate()
}
type OnInitialized interface { //初始化完成
	Initialized()
}
type Complete interface { // 注入完成
	OnComplete()
}
type PreComplete interface {
	OnPreComplete()
}
type PreCompletes []PreComplete

type PostComplete interface {
	OnPostComplete()
}
type PostCompletes []PostComplete

func (pos PostCompletes) OnPostComplete() {
	for _, po := range pos {
		po.OnPostComplete()
	}
}

func (pcs PreCompletes) OnPreComplete() {
	for _, pc := range pcs {
		pc.OnPreComplete()
	}
}

type CompleteList []Complete

func (cl CompleteList) OnComplete() {
	for _, c := range cl {
		c.OnComplete()
	}
}
