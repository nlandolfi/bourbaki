all.pdf: *.tex ../*.tex ../trademark.pdf ../sheets/*/*.tex

# doesn't work?!
	#../latexrun all.tex
	pdflatex all.tex
	pdflatex all.tex

inputs.tex: sheets.csv inputs.go inputs.tmpl
	go run inputs.go > inputs.tex

open: all.pdf
	open all.pdf

o: all.pdf
	make open

graph:
	cd .. && make open

