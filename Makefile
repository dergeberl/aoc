SHELL=/bin/bash

YEAR ?= $(shell /bin/date +"%Y")
DAY ?= $(shell /bin/date +"%d")

AOC_SESSION_COOKIE ?= $(shell cat ~/.aoc-session-cookie)
AOC_INPUT_FILE = https://adventofcode.com/$(YEAR)/day/$(shell echo $(DAY) | sed 's/^0*//')/input

CURRENTDAYFOLDER = $(YEAR)/day$(DAY)
CURRENTDAYGOFILE = $(CURRENTDAYFOLDER)/day$(DAY).go
CURRENTDAYGOTESTFILE = $(CURRENTDAYFOLDER)/day$(DAY)_test.go
CURRENTDAYINPUT = $(CURRENTDAYFOLDER)/input.txt


all: $(CURRENTDAYFOLDER) $(CURRENTDAYGOFILE) $(CURRENTDAYGOTESTFILE) $(CURRENTDAYINPUT)

$(CURRENTDAYFOLDER):
	@mkdir -p ./$(CURRENTDAYFOLDER)

$(CURRENTDAYGOFILE):
	@cp -r template/day0.go ./$(CURRENTDAYGOFILE)
	@sed -i 's/Day0/Day'$(DAY)'/g' ./$(CURRENTDAYGOFILE)
	@sed -i 's/day0/day'$(DAY)'/g' ./$(CURRENTDAYGOFILE)

$(CURRENTDAYGOTESTFILE):
	@cp -r template/day0_test.go ./$(CURRENTDAYGOTESTFILE)
	@sed -i 's/Day0/Day'$(DAY)'/g' ./$(CURRENTDAYGOTESTFILE)
	@sed -i 's/day0/day'$(DAY)'/g' ./$(CURRENTDAYGOTESTFILE)

$(CURRENTDAYINPUT):
	@curl --cookie $(AOC_SESSION_COOKIE) $(AOC_INPUT_FILE) -o ./$(CURRENTDAYINPUT)