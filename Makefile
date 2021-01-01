PACKAGES := commander ika jwt middleware

test:
	for dir in $(PACKAGES); do \
	  $(MAKE) -C $${dir} test; \
	done
