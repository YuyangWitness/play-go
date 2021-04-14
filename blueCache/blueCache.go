package blueCache

import (
	"container/list"
	"play-go/blueCache/utils"
)

type Cache struct {
	maxByte    int64                    // 最大支持存储
	usedByte   int64                    // 已使用存储
	doubleLink *list.List               // 双向链表存储Entry
	cache      map[string]*list.Element // 缓存字典[key]element
	callBack   utils.CallBackFunc       // 当一个缓存被删除运行回调
}

// 存储在链表里的数据
type Entry struct {
	key   string
	value utils.Value
}

func New(maxBytes int64, callBack utils.CallBackFunc) *Cache {
	return &Cache{
		maxByte:    maxBytes,
		callBack:   callBack,
		doubleLink: list.New(),
		cache:      make(map[string]*list.Element),
	}
}

func (c *Cache) RemoveOldElement() {
	el := c.doubleLink.Front()
	if el != nil {
		c.doubleLink.Remove(el)
		kv := el.Value.(*Entry)
		delete(c.cache, kv.key)
		c.usedByte -= int64(len(kv.key)) - int64(kv.value.Len())
		if c.callBack != nil {
			c.callBack(kv.key, kv.value)
		}
	}
}

func (c *Cache) Add(key string, value utils.Value) {
	if el, ok := c.cache[key]; ok {
		c.doubleLink.MoveToBack(el)
		kv := el.Value.(*Entry)
		c.usedByte += int64(value.Len()) - int64(kv.value.Len())
		kv.value = value
	} else {
		el := c.doubleLink.PushBack(&Entry{key, value})
		c.cache[key] = el
		c.usedByte += int64(len(key)) + int64(value.Len())
	}

	// 如果超过maxBytes则删除最前列的
	for c.maxByte != 0 && c.maxByte < c.usedByte {
		// 删除最前列的
		c.RemoveOldElement()
	}
}

func (c *Cache) Get(key string) (utils.Value, bool) {
	if el, ok := c.cache[key]; ok {
		c.doubleLink.MoveToBack(el)
		kv := el.Value.(*Entry)
		return kv.value, true
	}
	return nil, false
}

func (c *Cache) Len() int {
	return c.doubleLink.Len()
}
