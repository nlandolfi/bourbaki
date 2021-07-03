all.pdf: *.tex ../*.tex ../trademark.pdf ../sheets/*/*.tex
# doesn't work?!
	#../latexrun all.tex
	pdflatex all.tex
	pdflatex all.tex

inputs.tex: sheets.csv inputs.tmpl # inputs.go
# 	go run inputs.go > inputs.tex
	bbk_all -sheets ../sheets/ > inputs.tex

open: all.pdf
	open all.pdf

o: all.pdf
	make open

graph:
	cd .. && make open

