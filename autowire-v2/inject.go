package autowire

import (
	"fmt"
	"maps"
	"strings"
)

func Check() {
	lock.Lock()
	defer lock.Unlock()
	if isChecked {
		return
	}
	doCheck()

}
func doCheck() {

	if len(requireBy) == 0 {
		return
	}

	for {
		temp := maps.Clone(requireBy)
		autoCreate(temp)
		if len(temp) == len(requireBy) { // 意味着没有新增依赖
			break
		}
	}
	beanOnInitialized() // 执行bean的 OnInitialized 事件
	lack := make(map[string][]string)
	for beanName, reqs := range requireBy {
		b, has := beans[beanName]
		if !has {
			for _, req := range reqs {
				if req.root != nil {
					lack[beanName] = append(lack[beanName], fmt.Sprintf("%s.%s", req.root.name, strings.Join(req.path, ".")))
				} else {
					lack[beanName] = append(lack[beanName], "autowired")
				}
			}
			continue
		}
		for _, req := range reqs {

			setValue(req.value, b.v, req.isExported)
			if req.root != nil {
				req.root.requireCount--
			}
		}
	}

	// 检查是否有未找到的bean
	if len(lack) > 0 {
		panicLack(lack)
		return
	}
	// 检查未执行Complete事件的bean,并执行
	cls := make(CompleteList, 0, len(beans))
	prcLis := make(PreCompletes, 0, len(beans))
	pocList := make(PostCompletes, 0, len(beans))
	for _, b := range beans {
		if !b.isComplete && b.requireCount == 0 {
			b.isComplete = true
			if ch, ok := b.v.Interface().(Complete); ok {
				cls = append(cls, ch)
			}
			if pch, ok := b.v.Interface().(PreComplete); ok {
				prcLis = append(prcLis, pch)
			}
			if poc, ok := b.v.Interface().(PostComplete); ok {
				pocList = append(pocList, poc)
			}
		}
	}
	isChecked = true
	//清空已经注入的依赖关系表
	requireBy = make(map[string][]*requireItem)

	prcLis.OnPreComplete()
	cls.OnComplete()
	pocList.OnPostComplete()
}

func autoCreate(reqs map[string][]*requireItem) {
	for beanName := range reqs {
		if !tryAuto(beanName) {
			panic("can not find bean:" + beanName)
		}
	}
}

func beanOnInitialized() {
	for _, bv := range beans {
		if !bv.isComplete {
			if ih, ok := bv.v.Interface().(OnInitialized); ok {
				ih.Initialized()
			}
		}
	}
}
func panicLack(lack map[string][]string) {
	b := &strings.Builder{}
	for beanName, reqs := range lack {
		_, _ = fmt.Fprintf(b, "not found bean [%s], require by [%s] ", beanName, strings.Join(reqs, ","))
	}
	panic(b.String())
}
