inputs.tex: sheets.csv inputs.go inputs.tmpl
	go run inputs.go > inputs.tex

all.toc: all.tex ../sheets/*/*.tex
	pdflatex all.tex

all.pdf: inputs.tex all.toc all.tex ../*.tex ../trademark.pdf ../sheets/*/*.tex
	pdflatex all.tex
	make defs

make:
	make all.pdf

open: all.pdf
	open all.pdf

o:
	make open

graph:
	cd .. && make open

defs: all.tex
	defs all.tex
