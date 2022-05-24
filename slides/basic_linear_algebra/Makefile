o: basic_linear_algebra.pdf
	open -a /Applications/Skim.app basic_linear_algebra.pdf

basic_linear_algebra.tex: basic_linear_algebra.lit
	lit -in basic_linear_algebra.lit -o slides -out basic_linear_algebra.tex

fmt:
	lit -in basic_linear_algebra.lit -out basic_linear_algebra.lit

basic_linear_algebra.pdf: ./basic_linear_algebra.tex ./macros.tex
	./latexrun basic_linear_algebra.tex

spell:
	aspell -c basic_linear_algebra.tex

reset:
	rm basic_linear_algebra.pdf basic_linear_algebra.tex
