.PHONY: git
git:
	git add .
	git commit -m"自动提交 git 代码"
	git push
.PHONY: tag
tag:
	git push --tags

.PHONY: proto
proto:
	protoc -I . --micro_out=. --gogofaster_out=. proto/event/event.proto