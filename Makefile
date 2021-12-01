SHELL=/bin/bash

YEAR ?= $(shell /bin/date +"%Y")
DAY ?= $(shell /bin/date +"%d")

AOC_SESSION_COOKIE ?= $(shell cat ~/.aoc-session-cookie)
AOC_INPUT_FILE = https://adventofcode.com/$(YEAR)/day/$(shell echo $(DAY) | sed 's/^0*//')/input

CURRENTDAY = $(YEAR)/day$(DAY)

all: $(CURRENTDAY)

$(CURRENTDAY):
	@mkdir -p ./$(YEAR)/
	@cp -r template ./$(CURRENTDAY)
	@sed -i 's/Day0/Day'$(DAY)'/g' ./$(CURRENTDAY)/day0_test.go
	@sed -i 's/Day0/Day'$(DAY)'/g' ./$(CURRENTDAY)/day0.go
	@sed -i 's/day0/day'$(DAY)'/g' ./$(CURRENTDAY)/day0_test.go
	@sed -i 's/day0/day'$(DAY)'/g' ./$(CURRENTDAY)/day0.go
	@mv ./$(CURRENTDAY)/day0.go ./$(CURRENTDAY)/day$(DAY).go
	@mv ./$(CURRENTDAY)/day0_test.go ./$(CURRENTDAY)/'day'$(DAY)'_test.go'
	@echo "folder ./$(CURRENTDAY) from template created"
	@curl --cookie $(AOC_SESSION_COOKIE) $(AOC_INPUT_FILE) -o ./$(CURRENTDAY)/input.txt

