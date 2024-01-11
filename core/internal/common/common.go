package common

import (
	"sort"

	"github.com/pjimming/zustacm/core/internal/types"
	"github.com/pjimming/zustacm/core/model"

	"github.com/jinzhu/copier"
)

func GetResourceTree(resources []*model.Resource) []*types.Resource {
	resourcesMap := make(map[int64][]*model.Resource)
	for _, item := range resources {
		resourcesMap[item.ParentID] = append(resourcesMap[item.ParentID], item)
	}

	dummy := &types.Resource{ID: 0, Children: make([]*types.Resource, 0)}

	q := make([]*types.Resource, 0)
	q = append(q, dummy)

	for len(q) > 0 {
		node := q[0]
		q = q[1:]

		v, ok := resourcesMap[node.ID]
		if !ok {
			continue
		}

		node.Children = make([]*types.Resource, 0)

		for _, resource := range v {
			child := &types.Resource{}
			_ = copier.Copy(child, resource)
			node.Children = append(node.Children, child)
			q = append(q, child)
		}

		sort.Slice(node.Children, func(i, j int) bool {
			return node.Children[i].Order < node.Children[j].Order
		})
	}

	return dummy.Children
}
