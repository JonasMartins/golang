DONE="Task Done"

# GIT
status:
	@echo "Git status"
	git status -u

commit:
	@echo "Git add and commit"
	git add . && git commit
	@echo ${DONE}

push:
	@echo "Git push"
	git push -u origin development

reset:
	@echo "Git reset"
	git reset --hard && clean -fd
