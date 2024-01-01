# 赋值
# https://www.cnblogs.com/wanghuaijun/p/8092747.html	?= := = += 的区别
v1="v1"
v2:=${v1}
v1="v1.1"
v3:=${v1}
v4="v4"
v4?="v4.1"

.PHONY: t1
t1:
	@echo "v1 : ${v1}"
	@echo "v2 : ${v2}"
	@echo "v3 : ${v3}"
	@echo "v4 : ${v4}"

s1="s1"
ifdef GOOS
	s1+="GOOS"
endif

s2="s2"
ifeq ($(GOOS), darwin)
	s2+="GOOS"
endif

.PHONY: t2
t2:
	@echo "s1 : ${s1}"
	@echo "s2 : ${s2}"
