package main

type Set interface {
	Put(key string)
	Keys() []string
	Contains(key string) bool
	Remove(key string)
	// 如果之前已经有了，就返回旧的值，absent =false
	// 如果之前没有，就塞下去，返回 absent = true
	PutIfAbsent(key string) (old string, absent bool)
	Num()int
	ReturnMap() map[string]bool
}

type set struct{
	data map[string]bool
}

func (s *set) Put(key string){
	s.data[key] = true
}
func (s *set) Keys() []string{
	res:=make([]string,0)
	for k,_:=range(s.data){
		res=append(res,k)
	}
	return res
}
func (s *set)Contains(key string) bool{
	
	return s.data[key]
}
func (s *set)Remove(key string){
	if s.Contains(key){
		delete(s.data,key)
	}
}
func (s *set)PutIfAbsent(key string)(old string,absent bool){
	if s.Contains(key){
		return key,false
	}else{
		s.Put(key)
		return "",s.data[key]
	}
}
func (s *set)Num()int{
	return len(s.data)
}
func (s set)ReturnMap() map[string]bool{
	return s.data
}
func NewSet() Set{
	return &set{
		data: make(map[string]bool,0),
	}
}