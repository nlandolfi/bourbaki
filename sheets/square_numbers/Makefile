# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make square_numbers.pdf
	make graph.pdf

square_numbers.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > square_numbers.tex

square_numbers.pdf: ../../*.tex ../../trademark.pdf *.tex square_numbers.tex
	# pdflatex --file-line-error -interaction=nonstopmode square_numbers.tex
	../../latexrun square_numbers.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: square_numbers.pdf
	open square_numbers.pdf

o:
	make open

terms: square_numbers.tex
	bbk_sheet -mode ts square_numbers.tex

spell:
	aspell -c square_numbers.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f square_numbers.*
	make reset
	bbk_sheet -mode mk > Makefile