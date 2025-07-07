package encode

import (
	"fmt"
	"gopkg.in/yaml.v3"
)

func Unmarshal[T any](data []byte) (t *T, err error) {
	t = new(T)
	err = yaml.Unmarshal(data, t)
	return
}

type Entry[T any] struct {
	Key   string
	Value T
}

func UnmarshalSortMap[T any, R any](parent *yaml.Node, fieldName string, handler func(v *Entry[T]) R) ([]R, error) {
	for i := 0; i < len(parent.Content); i += 2 {
		node := parent.Content[i]
		if node.Kind == yaml.ScalarNode && node.Value == fieldName {
			valueNode := parent.Content[i+1]
			if valueNode.Kind == yaml.MappingNode {

				return doUnmarshalSortMap(valueNode.Content, handler)
			} else {
				return nil, fmt.Errorf("unsuport type %s for field %s", valueNode.Tag, fieldName)
			}
		}
	}
	return nil, nil
}
func doUnmarshalSortMap[T any, R any](nodes []*yaml.Node, handler func(v *Entry[T]) R) ([]R, error) {
	l := len(nodes)
	rs := make([]R, l/2)
	for i := 0; i < l; i += 2 {
		e := new(Entry[T])
		e.Key = nodes[i].Value
		err := nodes[i+1].Decode(&e.Value)
		if err != nil {
			return nil, err
		}
		rs[i/2] = handler(e)
	}
	return rs, nil
}
