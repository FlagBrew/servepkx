SUBDIRS = go java
OUT = out

all:
	@for dir in $(SUBDIRS); do $(MAKE) -C $$dir; done
	@mkdir -p $(OUT)
	@mv go/servepkx.exe go/servepkx_linux go/servepkx_mac java/build/servepkx.jar $(OUT)
	@zip -r $(OUT)/servepkx.zip $(OUT)

clean:
	@rm -fr $(OUT)
	@for dir in $(SUBDIRS); do $(MAKE) clean -C $$dir; done