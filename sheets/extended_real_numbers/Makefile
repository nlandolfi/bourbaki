# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make extended_real_numbers.pdf
	make graph.pdf

extended_real_numbers.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > extended_real_numbers.tex

extended_real_numbers.pdf: ../../*.tex ../../trademark.pdf *.tex extended_real_numbers.tex
	pdflatex --file-line-error -interaction=nonstopmode extended_real_numbers.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: extended_real_numbers.pdf
	open extended_real_numbers.pdf

o:
	make open

defs: extended_real_numbers.tex
	defs extended_real_numbers.tex

spell:
	aspell -c extended_real_numbers.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f extended_real_numbers.*
	make reset
	bbk_sheet -mode mk > Makefile