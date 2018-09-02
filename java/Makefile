JAVAC = javac
JAR = jar
BUILD = build
PACKAGE = servepkx
SOURCE = src
TARGET = servepkx

.PHONY: clean all

all: build

build:
	@mkdir -p $(BUILD)
	@$(JAVAC) -d $(BUILD) $(SOURCE)/$(PACKAGE)/*.java
	@cd $(BUILD) && echo Manifest-Version: 1.0 > manifest.txt
	@cd $(BUILD) && echo Main-Class: $(PACKAGE).MainWindow > manifest.txt
	@cd $(BUILD) && $(JAR) cvfm $(TARGET).jar manifest.txt $(PACKAGE)/*.class

clean:
	@rm -fr $(BUILD)
	@echo Clean...